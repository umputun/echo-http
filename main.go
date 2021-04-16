package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/go-pkgz/lgr"
	"github.com/go-pkgz/rest"
	"github.com/umputun/go-flags"
)

var opts struct {
	Listen  string `short:"l" long:"listen" env:"LISTEN" default:"0.0.0.0:8080" description:"listen on host:port"`
	Message string `short:"m" long:"message" env:"MESSAGE" default:"echo" description:"response message"`
	Dbg     bool   `long:"dbg" env:"DEBUG" description:"debug mode"`
}

var revision = "unknown"

func main() {
	fmt.Printf("echo-http %s\n", revision)

	p := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash|flags.HelpFlag)
	p.SubcommandsOptional = true
	if _, err := p.Parse(); err != nil {
		if err.(*flags.Error).Type != flags.ErrHelp {
			log.Printf("[ERROR] cli error: %v", err)
		}
		os.Exit(1)
	}
	setupLog(opts.Dbg)
	if err := run(); err != nil {
		log.Printf("[ERROR] server failed, %v", err)
	}
}

func run() error {

	router := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		echo := struct {
			Message    string   `json:"message"`
			Request    string   `json:"request"`
			Host       string   `json:"host"`
			Headers    []string `json:"headers"`
			RemoteAddr string   `json:"remote_addr"`
		}{
			Message:    opts.Message,
			Request:    r.Method + " " + r.RequestURI,
			Host:       r.Host,
			RemoteAddr: r.RemoteAddr,
		}

		for k, vv := range r.Header {
			echo.Headers = append(echo.Headers, fmt.Sprintf("%s:%s", k, strings.Join(vv, ", ")))
		}

		rest.RenderJSON(w, &echo)
	})

	srv := http.Server{Addr: opts.Listen,
		Handler:           router,
		ReadHeaderTimeout: time.Second * 30,
		WriteTimeout:      time.Second * 30,
		IdleTimeout:       time.Second * 30,
	}

	return srv.ListenAndServe()
}

func setupLog(dbg bool) {
	if dbg {
		log.Setup(log.Debug, log.CallerFile, log.CallerFunc, log.Msec, log.LevelBraces)
		return
	}
	log.Setup(log.Msec, log.LevelBraces)
}
