package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/markbates/pkger"
	"io"
	"log"
	"net/http"
	"os/user"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	remote "github.com/mkozhukh/go-remote"
)

var configPath string
var hub *remote.Hub

func main() {
	flag.StringVar(&configPath, "config", "", "path to the configuration file")
	flag.Parse()

	if configPath == "" {
		usr, _ := user.Current()
		configPath = usr.HomeDir + "/.dash.yml"
	}

	reloadConfig()
	initJWT()
	go trackChanges()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	if Config.Cors != "" {
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{Config.Cors},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
			MaxAge:           300,
		})
		r.Use(c.Handler)
	}

	r.Get("/api/v1", initApi())

	fs := http.FileServer(pkger.Dir("/public"))
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := pkger.Stat("/public"+r.RequestURI); err != nil {
			f, _ := pkger.Open("/public/index.html")
			defer f.Close()

			w.Header().Set("Content-type", "text/html")
			_,_ = io.Copy(w, f)
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	fmt.Println("Listen at port ", Config.Port)
	err := http.ListenAndServe(Config.Port, r)
	if err != nil {
		log.Println(err.Error())
	}
}

func reloadConfig() {
	Config.LoadFromFile(configPath)
}

func initApi() http.HandlerFunc {
	api := remote.NewServer(&remote.ServerConfig{
		WebSocket: true,
	})

	must(api.AddConstant("version", "1.0.0"))
	must(api.AddService("admin", &AdminAPI{}))

	must(api.Dependencies.AddProvider(func(ctx context.Context) *remote.Hub { return api.Events }))
	must(api.Dependencies.AddProvider(func(ctx context.Context) remote.ConnectionID {
		cid, _ := ctx.Value(remote.ConnectionValue).(remote.ConnectionID)
		return cid
	}))

	return api.ServeHTTP
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func trackChanges() {
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	err := watcher.Add(configPath)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	go debounce(200*time.Millisecond, watcher.Events, func(ev fsnotify.Event) {
		if ev.Op == fsnotify.Write {
			reloadConfig()
		}
	})
	<-done
}

func debounce(interval time.Duration, input chan fsnotify.Event, cb func(ev fsnotify.Event)) {
	var item fsnotify.Event
	data := false
	timer := time.NewTimer(interval)

	for {
		select {
		case item = <-input:
			data = true
			timer.Reset(interval)
		case <-timer.C:
			if data {
				cb(item)
			}
		}
	}
}
