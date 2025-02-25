package web

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/bigredeye/notmanytask/internal/config"
	"github.com/bigredeye/notmanytask/internal/database"
	"github.com/bigredeye/notmanytask/internal/deadlines"
	"github.com/bigredeye/notmanytask/internal/gitlab"
	"github.com/bigredeye/notmanytask/internal/scorer"
	zlog "github.com/bigredeye/notmanytask/pkg/log"
	"github.com/pkg/errors"
)

func Run() error {
	flag.Parse()
	config, err := config.ParseConfig()
	if err != nil {
		return err
	}
	log.Printf("Parsed config: %+v", config)

	logger, err := zlog.Init(config.Log)
	if err != nil {
		return errors.Wrap(err, "Failed to init logger")
	}
	defer func() {
		err = zlog.Sync()
	}()

	wg := sync.WaitGroup{}
	defer wg.Wait()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := database.OpenDataBase(logger.Named("database"), fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		config.DataBase.User,
		config.DataBase.Pass,
		config.DataBase.Host,
		config.DataBase.Port,
		config.DataBase.Name,
	))
	if err != nil {
		return errors.Wrap(err, "Failed to open database")
	}

	deadlinesCtx, deadlinesCancel := context.WithCancel(ctx)
	defer deadlinesCancel()
	deadlines, err := deadlines.NewFetcher(config, logger.Named("deadlines.fetcher"))
	if err != nil {
		return errors.Wrap(err, "Failed to create deadlines fetcher")
	}

	git, err := gitlab.NewClient(config, logger.Named("gitlab"))
	if err != nil {
		return errors.Wrap(err, "Failed to create gitlab client")
	}

	projectsCtx, projectsCancel := context.WithCancel(ctx)
	defer projectsCancel()
	projects, err := gitlab.NewProjectsMaker(git, db)
	if err != nil {
		return errors.Wrap(err, "Failed to create projects maker")
	}

	pipelinesCtx, pipelinesCancel := context.WithCancel(ctx)
	defer pipelinesCancel()
	pipelines, err := gitlab.NewPipelinesFetcher(git, db)
	if err != nil {
		return errors.Wrap(err, "Failed to create projects maker")
	}

	mergeRequestsCtx, mergeRequestsCancel := context.WithCancel(ctx)
	defer mergeRequestsCancel()
	mergeRequests, err := gitlab.NewMergeRequestsUpdater(git, db)
	if err != nil {
		return errors.Wrap(err, "Failed to create merge requests updater")
	}

	scorer := scorer.NewScorer(db, deadlines, git)

	wg.Add(4)
	go func() {
		defer wg.Done()
		deadlines.Run(deadlinesCtx)
	}()
	go func() {
		defer wg.Done()
		projects.Run(projectsCtx)
	}()
	go func() {
		defer wg.Done()
		pipelines.Run(pipelinesCtx)
	}()
	go func() {
		defer wg.Done()
		mergeRequests.Run(mergeRequestsCtx)
	}()

	s, err := newServer(config, logger.Named("server"), db, deadlines, projects, pipelines, scorer, git)
	if err != nil {
		return errors.Wrap(err, "Failed to start server")
	}

	return errors.Wrap(s.run(), "Server failed")
}
