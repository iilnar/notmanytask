package database

import (
	goerrors "errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"moul.io/zapgorm2"

	"github.com/bigredeye/notmanytask/internal/models"
)

type DataBase struct {
	*gorm.DB
}

type DuplicateKey struct {
	nested error
}

func (e *DuplicateKey) Error() string {
	return e.nested.Error()
}

func (e *DuplicateKey) Unwrap() error {
	return e.nested
}

func IsDuplicateKey(err error) bool {
	duplicateKey := &DuplicateKey{}
	return goerrors.As(err, &duplicateKey)
}

// gorm sucks huge balls:(
// https://github.com/go-gorm/gorm/issues/4037
func isUnqiueViolation(err error) bool {
	perr, ok := err.(*pgconn.PgError)
	if ok {
		return perr.Code == "23505"
	}
	return false
}

func OpenDataBase(logger *zap.Logger, dsn string) (*DataBase, error) {
	zapLogger := zapgorm2.New(logger.Named("gorm"))
	zapLogger.SetAsDefault()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: zapLogger,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Pipeline{}, &models.Session{}, &models.Flag{}, &models.MergeRequest{})
	if err != nil {
		return nil, err
	}

	return &DataBase{db}, nil
}

func (db *DataBase) AddUser(user *models.User) (*models.User, error) {
	var res models.User
	err := db.FirstOrCreate(&res, user).Error
	if err != nil {
		if isUnqiueViolation(err) {
			return nil, &DuplicateKey{err}
		}
		return nil, err
	}
	return &res, nil
}

func (db *DataBase) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	err := db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DataBase) FindUserByGitlabLogin(login string) (*models.User, error) {
	var user models.User
	err := db.First(&user, "gitlab_login = ?", login).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DataBase) FindUserByGitlabID(id int) (*models.User, error) {
	var user models.User
	err := db.First(&user, "gitlab_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DataBase) ListUsersWithoutRepos() ([]*models.User, error) {
	var users []*models.User
	err := db.Find(&users, "repository IS NULL AND gitlab_id IS NOT NULL AND gitlab_login IS NOT NULL").Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *DataBase) ListGroupUsers(groupName string, subgroupName string) ([]*models.User, error) {
	var users []*models.User
	if subgroupName != "" {
		err := db.Find(&users, "repository IS NOT NULL AND group_name = ? AND subgroup_name = ?", groupName, subgroupName).Order("created_at").Error
		if err != nil {
			return nil, err
		}
	} else {
		err := db.Find(&users, "repository IS NOT NULL AND group_name = ?", groupName).Order("created_at").Error
		if err != nil {
			return nil, err
		}
	}
	
	return users, nil
}

func (db *DataBase) SetUserGitlabAccount(uid uint, user *models.GitlabUser) error {
	res := db.Model(&models.User{}).
		Where("id = ? AND (gitlab_id IS NULL OR gitlab_login IS NULL)", uid).
		Updates(map[string]interface{}{
			"gitlab_id":    user.GitlabID,
			"gitlab_login": user.GitlabLogin,
		})

	if res.Error != nil {
		if isUnqiueViolation(res.Error) {
			return &DuplicateKey{res.Error}
		}
		return res.Error
	}

	if res.RowsAffected < 1 {
		return errors.Errorf("Unknown user %d", uid)
	}
	return nil
}

func (db *DataBase) SetUserRepository(user *models.User) error {
	res := db.Model(user).Update("repository", user.Repository)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected < 1 {
		return errors.Errorf("Unknown user %d", user.ID)
	}
	return nil
}

func (db *DataBase) AddPipeline(pipeline *models.Pipeline) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status"}),
	}).Create(pipeline).Error
}

func (db *DataBase) ListProjectPipelines(project string) (pipelines []models.Pipeline, err error) {
	pipelines = make([]models.Pipeline, 0)
	err = db.Find(&pipelines, "project = ?", project).Error
	if err != nil {
		pipelines = nil
	}
	return
}

func (db *DataBase) ListAllPipelines() (pipelines []models.Pipeline, err error) {
	pipelines = make([]models.Pipeline, 0)
	err = db.Find(&pipelines).Error
	if err != nil {
		pipelines = nil
	}
	return
}

func (db *DataBase) FindLatestPipeline(project string, task string) (*models.Pipeline, error) {
	pipelines := make([]models.Pipeline, 0)
	err := db.Find(&pipelines, "project = ? AND task = ?", project, task).Order("started_at desc").Error
	if err != nil {
		pipelines = nil
	}
	return &pipelines[0], err
}

func (db *DataBase) CreateSession(user uint) (*models.Session, error) {
	session := &models.Session{
		Token:  uuid.Must(uuid.NewUUID()).String(),
		UserID: user,
	}
	res := db.DB.Create(session)
	if res.Error != nil {
		return nil, res.Error
	}
	return session, nil
}

func (db *DataBase) FindSession(token string) (*models.Session, error) {
	var session models.Session
	res := db.DB.Where("token", token).Take(&session)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected < 1 {
		return nil, errors.New("Unknown session")
	}
	return &session, nil
}

func (db *DataBase) FindUserBySession(token string) (*models.User, *models.Session, error) {
	session, err := db.FindSession(token)
	if err != nil {
		return nil, nil, err
	}
	user, err := db.FindUserByID(session.UserID)
	if err != nil {
		return nil, session, err
	}
	return user, session, nil
}

func (db *DataBase) CreateFlag(task string) (*models.Flag, error) {
	flag := &models.Flag{
		ID:        fmt.Sprintf("{FLAG-%s-%s}", task, uuid.New().String()),
		Task:      task,
		CreatedAt: time.Now(),
	}
	err := db.Create(flag).Error
	if err != nil {
		return nil, err
	}
	return flag, nil
}

func (db *DataBase) SubmitFlag(id string, gitlabLogin string) error {
	result := db.Model(&models.Flag{}).Where("id = ? AND gitlab_login IS NULL", id).Update("gitlab_login", gitlabLogin)
	if goerrors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("Unknown flag")
	}
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Unknown flag")
	}
	return nil
}

func (db *DataBase) ListUserFlags(gitlabLogin string) (flags []models.Flag, err error) {
	flags = make([]models.Flag, 0)
	err = db.Find(&flags, "gitlab_login = ?", gitlabLogin).Error
	if err != nil {
		flags = nil
	}
	return
}

func (db *DataBase) ListSubmittedFlags() (flags []models.Flag, err error) {
	flags = make([]models.Flag, 0)
	err = db.Find(&flags, "gitlab_login IS NOT NULL").Error
	if err != nil {
		flags = nil
	}
	return
}

func (db *DataBase) AddMergeRequest(mergeRequest *models.MergeRequest) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status"}),
	}).Create(mergeRequest).Error
}

func (db *DataBase) FindMergeRequest(project string, task string) (*models.MergeRequest, error) {
	var mergeRequest models.MergeRequest
	res := db.DB.Where("project = ? AND task = ?", project, task).Take(&mergeRequest)
	if res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &mergeRequest, nil
}

func (db *DataBase) ListProjectMergeRequests(project string) (mergeRequests []models.MergeRequest, err error) {
	mergeRequests = make([]models.MergeRequest, 0)
	err = db.Find(&mergeRequests, "project = ?", project).Error
	if err != nil {
		mergeRequests = nil
	}
	return
}

func (db *DataBase) ListAllMergeRequests() (mergeRequests []models.MergeRequest, err error) {
	mergeRequests = make([]models.MergeRequest, 0)
	err = db.Find(&mergeRequests).Error
	if err != nil {
		mergeRequests = nil
	}
	return
}
