package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"server/api"
	"server/config"
	"server/db"
	"server/middleware"

	"server/controller"
	taskRepository "server/repository/task"
	"server/service"
)

const (
	// BLOG is a placeholder, to be used for "blog.orakem.ste"
	BLOG = "blog"

	// WIKI is a placeholder, to be used for "wiki.orakem.ste"
	WIKI = "wiki"
)

func main() {
	log.Println("Starting the server")

	r := mux.NewRouter()

	r.HandleFunc("/logout", api.LogoutFunc)

	createSubRouters(r)

	loggedRouter := middleware.Logger(r)
	srv := &http.Server{
		Handler: loggedRouter,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	if config.HTTPSMode() {
		srv.Addr = ":443"
		certFile := config.CertFile()
		keyFile := config.KeyFile()
		log.Fatal(srv.ListenAndServeTLS(certFile, keyFile))
	} else {
		if config.HTTPPort() != "" {
			srv.Addr = config.HTTPPort()
		} else {
			srv.Addr = ":80"
		}
		log.Fatal(srv.ListenAndServe())
	}

}

func mapStaticFiles(r *mux.Router) {
	// mathjax and js folders for js and css
	r.PathPrefix("/mathjax/").Handler(middleware.Middleware(
		http.StripPrefix("/mathjax/", http.FileServer(http.Dir(config.MathjaxDir()))),
		middleware.RedirectHTTPS,
	))

	r.PathPrefix("/js/").Handler(middleware.Middleware(
		http.StripPrefix("/js/", http.FileServer(http.Dir(config.JsCSSDir()))),
		middleware.RedirectHTTPS,
	))
}

// createSubRouters create routers for wiki and blogs
func createSubRouters(r *mux.Router) {

	// map wiki pages
	wikiSubRouter := r.Host(WIKI + "." + config.DomainName()).Subrouter()
	wikiSubRouter.HandleFunc("/login", api.LoginFunc)
	wikiSubRouter.HandleFunc("/logout", api.LogoutFunc)
	mapStaticFiles(wikiSubRouter)
	wikiSubRouter.PathPrefix("/files/").Handler(middleware.Middleware(
		http.StripPrefix("/files/", http.FileServer(http.Dir(config.FilesDir()))),
		middleware.RequiresLogin,
		middleware.RedirectHTTPS,
	))
	// html pages by default are private, except for blog pages.
	wikiSubRouter.PathPrefix("/").Handler(middleware.Middleware(
		http.StripPrefix("/", http.FileServer(http.Dir("../html"))),
		middleware.RequiresLogin,
		middleware.RedirectHTTPS,
	))

	taskConfig := config.TaskConfiguration()
	if taskConfig.IsNotEmpty() {
		log.Println("Creating task server...")
		taskSubRouter := r.Host("task." + config.DomainName()).Subrouter()
		mapTaskServer(taskSubRouter, taskConfig)
	}

	// map blog pages
	if config.DeployBlog() {
		blogSubrouter := r.Host(BLOG + "." + config.DomainName()).Subrouter()
		mapStaticFiles(blogSubrouter)

		blogSubrouter.PathPrefix("/files/").Handler(middleware.Middleware(
			http.StripPrefix("/files/", http.FileServer(http.Dir("../blog/files/"))),
			middleware.RedirectHTTPS,
		))

		blogSubrouter.PathPrefix("/").Handler(middleware.Middleware(
			http.StripPrefix("/", http.FileServer(http.Dir("../html/blog"))),
			middleware.RedirectHTTPS,
		))

		// map "/" to go to blogSbrouter by default
		if config.HTTPSMode() {
			r.Handle("/", http.RedirectHandler("https://"+BLOG+"."+config.DomainName()+"/", 301))
		} else {
			if config.HTTPPort() != "" {
				r.Handle("/", http.RedirectHandler("http://"+BLOG+"."+config.DomainName()+config.HTTPPort()+"/", 301))
			} else {
				r.Handle("/", http.RedirectHandler("http://"+BLOG+"."+config.DomainName()+"/", 301))
			}
		}

	}

}

func mapTaskServer(r *mux.Router, taskConfig config.TaskConfig) {
	var dbHandler db.Handler

	dbFile := taskConfig.DbURL

	log.Println("Opening db connection")

	switch taskConfig.DbType {
	case "SQLITE":
		dbHandler = db.NewSqliteHandler(dbFile)
		log.Printf("%v", dbHandler.Type())
	default:
		log.Fatalf("No handler registered for %s", taskConfig.DbType)
	}

	taskRepository.InitTaskRepo(dbHandler)
	err := service.InitializeTaskService(taskRepository.Repository())
	if err != nil {
		log.Fatalf("Could not start task service: %s", err.Error())
	}

	taskController := controller.TaskController{TaskService: service.TaskService}
	r.HandleFunc("/", HelloTask).Methods("GET")

	r.HandleFunc("/api/tasks", taskController.GetAllTasks).Methods("GET")
	r.HandleFunc("/api/task/{name}", taskController.GetTask).Methods("GET")
	r.HandleFunc("/api/task", taskController.CreateTask).Methods("POST")
	r.HandleFunc("/api/task", taskController.UpdateTask).Methods("PUT")
	r.HandleFunc("/api/task/{name}", taskController.DeleteTask).Methods("DELETE")

}

// HelloTask is a temp function. Delete it
func HelloTask(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Hello World", http.StatusMethodNotAllowed)
}
