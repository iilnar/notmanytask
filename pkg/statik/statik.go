// Code generated by statik. DO NOT EDIT.

package statik

import (
	"github.com/rakyll/statik/fs"
)


func init() {
	data := "PK\x03\x04\x14\x00\x08\x00\x00\x00<c)T\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00	\x00	\x00flag.tmplUT\x05\x00\x01U\xd4\xdaa<!doctype html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css\" rel=\"stylesheet\">\n\n    <title>HSE Advanced C&#43;&#43;</title>\n    <style>\n.navbar-brand {\n  font-size: 3rem;\n  font-weight: 300\n}\n\n#floatingFlag {\n  font-family: monospace;\n}\n    </style>\n  </head>\n  <body>\n      <nav class=\"navbar navbar-light bg-light\">\n          <div class=\"container\">\n              <span class=\"navbar-brand mb-0 h1\"><a href=\"/\" class=\"text-decoration-none text-dark\">Advanced C++</a></span>\n              <div class=\"row\">\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Deadlines }}\"><h5>Tasks</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Standings }}\"><h5>Standings</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.SubmitFlag }}\"><h5>Submit flag</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Repository }}\"><h5>My Repo</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Submits }}\"><h5>Submits</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Logout }}\"><h5>Logout</h5></a>\n                  </div>\n              </div>\n          </div>\n          </div>\n      </nav>\n\n    <div class=\"container p-2 my-2\">\n      <div class=\"row p-2\">\n        <div class=\"col col-lg-6 offset-lg-3 col-md-10 offset-md-1\">\n          <div class=\"card\">\n            <div class=\"card-body\">\n              <form method=\"post\" action=\"{{ .Links.SubmitFlag }}\" class=\"needs-validation was-validated\">\n                <div class=\"form-floating mb-3\">\n                  <input type=\"text\" class=\"form-control\" id=\"floatingFlag\" placeholder=\"Flag\" name=\"flag\" required pattern=\"\\{FLAG(-[a-z0-9_]+)+(-[0-9a-f]+)+\\}\">\n                  <label for=\"floatingFlag\">Flag value</label>\n                  <div class=\"invalid-feedback\">\n                    Flag should be in form <code>{FLAG-crashme-d18736e-9287-ffaa-8bqe-4e6d891516ef}</code>\n                  </div>\n                </div>\n\n              {{ if .ErrorMessage }}\n              <div class=\"alert alert-danger\" role=\"alert\">\n                {{ .ErrorMessage }}\n              </div>\n              {{ end }}\n\n              {{ if .SuccessMessage }}\n              <div class=\"alert alert-success\" role=\"alert\">\n                {{ .SuccessMessage }}\n              </div>\n              {{ end }}\n\n                <div class=\"d-grid\">\n                  <button type=\"submit\" class=\"btn btn-outline-success\">Submit flag</button>\n                </div>\n              </form>\n\n            </div>\n          </div>\n        </div>\n      </div>\n    </div>\n\n  </body>\n</html>\n\n\nPK\x07\x08\x9c\xec\x95m2\x0c\x00\x002\x0c\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00<c)T\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00	\x00	\x00home.tmplUT\x05\x00\x01U\xd4\xdaa<!doctype html>\n<html lang=\"en\">\n    <head>\n        <meta charset=\"utf-8\">\n        <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n        <link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css\" rel=\"stylesheet\">\n\n        <title>{{ .Title }}</title>\n        <style>\n.shadow-hover:hover {\n    transition: all 0.1s ease;\n    box-shadow:0 .5rem 1rem rgba(0,0,0,.15)!important\n}\n.shadow-hover {\n    -webkit-transition: all 0.1s ease;\n    -moz-transition: all 0.1s ease;\n    -o-transition: all 0.1s ease;\n    transition: all 0.1s ease;\n    box-shadow:0 .125rem .25rem rgba(0,0,0,.075)!important\n}\n\n.task {\n    overflow: hidden;\n}\n\n.task-success {\n    background-color: #a6e9d5;\n    border-color: #4dd4ac;\n}\n\n.task-failed {\n    background-color: #f8d7da;\n    border-color: #f1aeb5;\n}\n\n.task-checking {\n    border-color: #0d6efd;\n    background-color:#9ec5fe;\n}\n\n.task-assigned {\n    background-color: #f8f9fa;\n}\n\n.navbar-brand {\n  font-size: 3rem;\n  font-weight: 300\n}\n\n.nav-link {\n  color: rgba(0, 0, 0, 0.9);\n}\n        </style>\n    </head>\n    <body>\n        <nav class=\"navbar navbar-light bg-light\">\n            <div class=\"container\">\n                <span class=\"navbar-brand mb-0 h1\"><a href=\"/\" class=\"text-decoration-none text-dark\">Advanced C++</a></span>\n                <div class=\"row\">\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Deadlines }}\"><h5>Tasks</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Standings }}\"><h5>Standings</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.SubmitFlag }}\"><h5>Submit flag</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Repository }}\"><h5>My Repo</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Submits }}\"><h5>Submits</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Logout }}\"><h5>Logout</h5></a>\n                  </div>\n                </div>\n            </div>\n            </div>\n        </nav>\n\n        {{ if .Scores }}\n            {{ range .Scores.Groups }}\n                <div class=\"container p-2 my-2\">\n                    <div class=\"p-2\">\n                        <a name=\"{{ .PrettyTitle }}\" href=\"#{{ .PrettyTitle }}\" class=\"text-decoration-none text-dark\">\n                            <h1>{{ .PrettyTitle }} <span class=\"text-muted\">{{ .Deadline.String }}</span></h1>\n                        </a>\n                    </div>\n                    <div class=\"row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 row-cols-xl-5 g-4 text-center\">\n                        {{ range .Tasks }}\n                            <div class=\"col\">\n                                <a href=\"{{ .TaskUrl }}\" class=\"text-decoration-none text-dark\">\n                                    <div class=\"card h-100 task task-{{ .Status }} shadow-hover\">\n                                        <div class=\"card-body\">\n                                            <h3 class=\"card-title text-nowrap text-dark\">{{ .ShortName }}</h3>\n                                            {{ if .PipelineUrl }}\n                                                <a href=\"{{ .PipelineUrl }}\" class=\"text-decoration-none\">\n                                            {{ end }}\n                                                <p class=\"card-text fs-1 text-decoration-none text-dark\">\n                                                    {{.Score}} / {{.MaxScore}}\n                                                </p>\n                                            {{ if .PipelineUrl }}\n                                                </a>\n                                            {{ end }}\n                                        </div>\n                                    </div>\n                                </a>\n                            </div>\n                        {{ end }}\n                    </div>\n\n                    <div class=\"p-2\">\n                        <h1>Total score: {{ .Score }} / {{ .MaxScore }}</h1>\n                    </div>\n                </div>\n            {{ end }}\n        {{ end}}\n    </body>\n</html>\nPK\x07\x08\xbb\xd2}R\x7f\x11\x00\x00\x7f\x11\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00<c)T\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x08\x00	\x00kek.htmlUT\x05\x00\x01U\xd4\xdaakek!\nPK\x07\x08Ln\xf0\x0c\x05\x00\x00\x00\x05\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xe9s)T\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0b\x00	\x00signup.tmplUT\x05\x00\x01\xb6\xf1\xdaa<!doctype html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css\" rel=\"stylesheet\">\n\n    <title>HSE Advanced C&#43;&#43;</title>\n    <style>\n.navbar-brand {\n  font-size: 3rem;\n  font-weight: 300\n}\n    </style>\n  </head>\n  <body>\n    <nav class=\"navbar navbar-light bg-light\">\n      <div class=\"container\">\n        <div class=\"col col-xxl-4 offset-xxl-4 col-lg-6 offset-lg-3 col-md-10 offset-md-1\">\n          <p class=\"navbar-brand mb-0 h1 text-center\">Advanced C++</p>\n        </div>\n      </div>\n    </nav>\n\n    <div class=\"container p-2 my-2\">\n      <div class=\"row p-2\">\n        <div class=\"col col-xxl-4 offset-xxl-4 col-lg-6 offset-lg-3 col-md-10 offset-md-1\">\n          <div class=\"card\">\n            <div class=\"card-body\">\n              <form method=\"post\" action=\"{{ .Config.Endpoints.Signup }}\" class=\"needs-validation was-validated\">\n                <div class=\"form-floating mb-3\">\n                  <input type=\"text\" class=\"form-control\" id=\"floatingFirstName\" placeholder=\"Ivan\" name=\"firstname\" required pattern=\"[A-Za-z-]+\">\n                  <label for=\"floatingFirstName\">First name</label>\n                  <div class=\"invalid-feedback\">\n                    Please use only Latin letters\n                  </div>\n                </div>\n                <div class=\"form-floating mb-3\">\n                  <input type=\"text\" class=\"form-control\" id=\"floatingLastName\" placeholder=\"Petrov\" name=\"lastname\" required pattern=\"[A-Za-z-]+\">\n                  <label for=\"floatingLastName\">Last name</label>\n                  <div class=\"invalid-feedback\">\n                    Please use only Latin letters\n                  </div>\n                </div>\n                <div class=\"form-floating mb-3\">\n                  <input type=\"text\" class=\"form-control\" id=\"floatingSecretCode\" placeholder=\"LolKekCheburek\" name=\"secret\" required pattern=\"[A-Za-z0-9-_]+\">\n                  <label for=\"floatingSecretCode\">Secret code</label>\n                  <div class=\"invalid-feedback\">\n                    Ask your teacher\n                  </div>\n                </div>\n\n                {{ if .ErrorMessage }}\n                <div class=\"alert alert-danger\" role=\"alert\">\n                    {{ .ErrorMessage }}\n                </div>\n                {{ end }}\n\n                <div class=\"d-grid mb-3\">\n                  <button type=\"submit\" class=\"btn btn-outline-success\">Sign up via GitLab</button>\n                </div>\n              </form>\n\n              <div class=\"d-grid\">\n                <a class=\"btn btn-outline-primary btn-block\" href=\"{{ .Config.Endpoints.Login }}\">Login via GitLab</a>\n              </div>\n            </div>\n          </div>\n        </div>\n      </div>\n    </div>\n\n  </body>\n</html>\n\nPK\x07\x08\x1dB\xacQO\x0b\x00\x00O\x0b\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xac\x90)T\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00standings.tmplUT\x05\x00\x01\xe4#\xdba<!doctype html>\n<html lang=\"en\">\n    <head>\n        <meta charset=\"utf-8\">\n        <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n        <link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css\" rel=\"stylesheet\">\n\n        <title>{{ .Title }}</title>\n        <style>\n.shadow-hover:hover {\n    transition: all 0.1s ease;\n    box-shadow:0 .5rem 1rem rgba(0,0,0,.15)!important\n}\n.shadow-hover {\n    -webkit-transition: all 0.1s ease;\n    -moz-transition: all 0.1s ease;\n    -o-transition: all 0.1s ease;\n    transition: all 0.1s ease;\n    box-shadow:0 .125rem .25rem rgba(0,0,0,.075)!important\n}\n\n.task-success {\n    background-color: #a6e9d5;\n    border-color: #4dd4ac;\n}\n\n.task-failed {\n    background-color: #f8d7da;\n    border-color: #f1aeb5;\n}\n\n.task-checking {\n    border-color: #0d6efd;\n    background-color:#9ec5fe;\n}\n\n.task-assigned {\n    background-color: #f8f9fa;\n}\n\n.navbar-brand {\n    font-size: 3rem;\n    font-weight: 300\n}\n\n.nav-link {\n    color: rgba(0, 0, 0, 0.9);\n}\n\n.task {\n    width: 120px;\n    max-width: 120px;\n    overflow: hidden;\n}\n        </style>\n    </head>\n    <body>\n      <nav class=\"navbar navbar-light bg-light\">\n          <div class=\"container\">\n              <span class=\"navbar-brand mb-0 h1\"><a href=\"/\" class=\"text-decoration-none text-dark\">Advanced C++</a></span>\n              <div class=\"row\">\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Deadlines }}\"><h5>Tasks</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Standings }}\"><h5>Standings</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.SubmitFlag }}\"><h5>Submit flag</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Repository }}\"><h5>My Repo</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Submits }}\"><h5>Submits</h5></a>\n                  </div>\n                  <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Links.Logout }}\"><h5>Logout</h5></a>\n                  </div>\n              </div>\n          </div>\n          </div>\n      </nav>\n\n        <div class=\"container p-2 my-2\">\n            <div class=\"container row\">\n                {{ range .Groups }}\n                    <div class=\"col-auto\">\n                      <a class=\"nav-link\" href=\"{{ .Link }}\"><h5>{{ .Name }}</h5></a>\n                    </div>\n                {{ end }}\n            </div>\n            <div class=\"table-responsive\">\n                <table class=\"table table-hover\">\n                    <thead>\n                        <tr>\n                            <th scope=\"col\" class=\"num\">#</th>\n                            <th scope=\"col\" class=\"name\">Student</th>\n                            <th scope=\"col\">Score</th>\n                            {{ range .Standings.Deadlines }}\n                                {{ range .Tasks }}\n                                    <th scope=\"col\" class=\"task\">{{ .Task }}</th>\n                                {{ end }}\n                            {{ end }}\n                        </tr>\n                    </thead>\n                    <tbody>\n                        {{ with index .Standings.Users 0 }}\n                            <tr>\n                                <th scope=\"row\" class=\"num\">0</th>\n                                <th scope=\"row\" class=\"name\">Chuck Norris</th>\n                                <td>{{ .MaxScore }}</td>\n                                {{ range .Groups }}\n                                    {{ range .Tasks }}\n                                        <td class=\"task table-success\"><a href=\"/private/solutions/{{ .Task }}\" class=\"text-decoration-none text-dark\">{{ .MaxScore }}</a></td>\n                                    {{ end }}\n                                {{ end }}\n                            </tr>\n                        {{ end }}\n                        {{ range $index, $user := .Standings.Users }}\n                            <tr>\n                                <th scope=\"row\" class=\"num\">{{ inc $index }}</th>\n                                <th scope=\"row\" class=\"name\">{{ $user.User.FirstName }} {{ $user.User.LastName }}</th>\n                                <td>{{ $user.Score }}</td>\n                                {{ range $user.Groups }}\n                                    {{ range .Tasks }}\n                                        {{ if eq .Status \"success\"}}\n                                            <td class=\"task table-success\">\n                                        {{ else if eq .Status \"failed\"}}\n                                            <td class=\"task table-danger\">\n                                        {{ else }}\n                                            <td class=\"task\">\n                                        {{ end }}\n                                        {{ if .PipelineUrl }}\n                                            <a href=\"{{ .PipelineUrl }}\" class=\"text-decoration-none text-dark\">\n                                        {{ end }}\n                                        {{ .Score }}\n                                        {{ if .PipelineUrl }}\n                                            </a>\n                                        {{ end }}\n                                        </td>\n                                    {{ end }}\n                                {{ end }}\n                            </tr>\n                        {{ end }}\n                    </tbody>\n                </table>\n            </div>\n        </div>\n    </body>\n</html>\nPK\x07\x08s2\x18\x8b\xcd\x16\x00\x00\xcd\x16\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00<c)T\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00	\x00	\x00style.cssUT\x05\x00\x01U\xd4\xdaabody {\n    margin: 0;\n    font-family: 'Source Code Pro', monospace;\n    display: flex;\n}\n\n.site {\n    max-width: 1200px;\n    width: 100%;\n\n    margin: 0 auto;\n    padding-left: 4em;\n    padding-right: 4em;\n\n    display: flex;\n    flex-direction: column;\n    align-items: center;\n}\n\n.header-container {\n    margin: 0 auto;\n    margin-top: 2em;\n\n    display: flex;\n}\n\n/* ========================================================================== */\n\n.main-menu {\n    padding: 0;\n    display: flex;\n    list-style: none;\n    color: #455a64;\n}\n\n.main-menu a {\n    text-decoration: none;\n    color: #455a64;\n}\n\n.main-menu li {\n    font-size: 1em;\n    text-transform: uppercase;\n    margin-left: 0.66em;\n}\n\n.main-menu li .current {\n    font-weight: bold;\n}\n\n/* ========================================================================== */\n\n.main {\n    width: 100%;\n    display: flex;\n    flex-direction: column;\n    align-items: center;\n}\n\n/* ========================================================================== */\n\n.flag-submit {\n    display: flex;\n    align-content: center;\n    margin: auto;\n}\n\n/* ========================================================================== */\n\n.group {\n    display: flex;\n    flex-direction: column;\n    width: 100%;\n}\n\n.group a {\n    text-decoration: none;\n}\n\n.group-header {\n    display: flex;\n}\n\n.group-header h1 {\n    white-space: pre;\n    margin: 0em;\n}\n\n.group-tasks {\n    display: flex;\n    flex-wrap: wrap;\n}\n\n.task {\n    width: 200px;\n    height: 120px;\n    margin: 10px;\n\n    display: flex;\n    flex-direction: column;\n    align-items: center;\n}\n\n.unsolved {\n    background-color: #1e3250;\n    color: white;\n}\n\n.solved {\n    background-color: #66cda3;\n    color: black;\n}\n\n.task .name {\n    margin: 0 auto;\n    margin-top: 0.33em;\n    font-size: 1.5em;\n    white-space: nowrap;\n}\n\n.task .score {\n    margin: 0 auto;\n    font-size: 3em;\n    font-weight: bold;\n}\n\n/* ========================================================================== */\n\n.signup {\n    width: 100%;\n    \n    display: flex;\n    flex-direction: column;\n    justify-content: center;\n    align-items: center;\n    margin: 2em;\n}\n\n.signup .login {\n    padding-top: 2em;\n    padding-bottom: 2em;\n\n    display: flex;\n}\n\n.login-button {\n    display: flex;\n\n    font-size: 2em;\n\n    margin: auto;\n    height: 80px;\n    width: 300px;\n\n    border: solid;\n    border-width: 1px;\n    border-color: #168f48;\n    background-color: #1aaa55;\n\n    text-decoration: none;\n}\n\n.login-button .text {\n    margin: auto;\n    color: white;\n}\n\n.signup .or {\n    display: flex;\n    min-width: 100px;\n}\n\n.or .text {\n    font-size: 1em;\n    margin: auto;\n}\n\n.signup .register {\n    display: flex;\n    padding-top: 2em;\n    padding-bottom: 2em;\n}\n\n.form {\n    width: 500px;\n\n    display: flex;\n    flex-direction: column;\n    \n    border: 1px solid #e5e5e5;\n}\n\n.form-header {\n    display: flex;\n    align-items: center;\n}\n\n.form-header h1 {\n    margin: 0 auto;\n    padding-top: 0.33em;\n    padding-bottom: 0.33em;\n    font-weight: normal;\n    font-size: 2em;\n}\n\n.form .form-element {\n    flex: 1;\n\n    margin: 0.33em;\n    margin-bottom: 0;\n\n    padding: 0.33em;\n    padding-bottom: 0;\n\n    display: flex;\n    flex-direction: column;\n}\n\n.form .form-element.last {\n    padding-bottom: 0.33em;\n    margin-bottom: 0.33em;\n}\n\n.form-element input {\n    flex: 1;\n    height: 40px;\n\n    font-size: 1.5em;\n    padding-left: 0.1em;\n    border: 1px solid #e5e5e5;\n}\n\n.form-element .button {\n    background-color: #1f78d1;\n    border-color: #1b69b6;\n    color: white;\n    cursor: pointer;\n    font-family: 'Source Code Pro', monospace;\n    font-size: 1em;\n}\n\n.form-element .name {\n    margin-left: 0.33em;\n    margin-bottom: 0.33em;\n    color: #555555;\n}\n\n.form .form-error {\n    background-color: #db3b21;\n}\n\n.form-error .error-message {\n    margin-left: 0.33em;\n    margin-bottom: 0.33em;\n    \n    color: white;\n}\n\n/* ========================================================================== */\n\n.status {\n    display: flex;\n    flex-direction: column;\n    width: 400px;\n    margin-right: 60px;\n}\n\n.status h1 {\n    margin-left: auto;\n    margin-right: auto;\n}\n\ntable {\n    border-spacing: 0.66em;\n}\n\ntable td {\n    text-align: center;\n}\n\ntable th {\n    text-align: center;\n}\nPK\x07\x08\xff\x8bCA\x9d\x10\x00\x00\x9d\x10\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00<c)T\x9c\xec\x95m2\x0c\x00\x002\x0c\x00\x00	\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81\x00\x00\x00\x00flag.tmplUT\x05\x00\x01U\xd4\xdaaPK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00<c)T\xbb\xd2}R\x7f\x11\x00\x00\x7f\x11\x00\x00	\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81r\x0c\x00\x00home.tmplUT\x05\x00\x01U\xd4\xdaaPK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00<c)TLn\xf0\x0c\x05\x00\x00\x00\x05\x00\x00\x00\x08\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x811\x1e\x00\x00kek.htmlUT\x05\x00\x01U\xd4\xdaaPK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xe9s)T\x1dB\xacQO\x0b\x00\x00O\x0b\x00\x00\x0b\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81u\x1e\x00\x00signup.tmplUT\x05\x00\x01\xb6\xf1\xdaaPK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xac\x90)Ts2\x18\x8b\xcd\x16\x00\x00\xcd\x16\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81\x06*\x00\x00standings.tmplUT\x05\x00\x01\xe4#\xdbaPK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00<c)T\xff\x8bCA\x9d\x10\x00\x00\x9d\x10\x00\x00	\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81\x18A\x00\x00style.cssUT\x05\x00\x01U\xd4\xdaaPK\x05\x06\x00\x00\x00\x00\x06\x00\x06\x00\x86\x01\x00\x00\xf5Q\x00\x00\x00\x00"
		fs.Register(data)
	}
	