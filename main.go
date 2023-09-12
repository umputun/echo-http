package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
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

	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Printf("[DEBUG] received signal %v, exiting...", sig)
		cancel()
	}()

	if err := run(ctx); err != nil {
		log.Printf("[ERROR] server failed, %v", err)
	}
}

func run(ctx context.Context) error {

	router := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		echo := struct {
			Message    string            `json:"message"`
			Request    string            `json:"request"`
			Host       string            `json:"host"`
			Headers    map[string]string `json:"headers"`
			RemoteAddr string            `json:"remote_addr"`
		}{
			Message:    opts.Message,
			Request:    r.Method + " " + r.RequestURI,
			Host:       r.Host,
			Headers:    make(map[string]string),
			RemoteAddr: r.RemoteAddr,
		}

		for k, vv := range r.Header {
			echo.Headers[k] = strings.Join(vv, "; ")
		}

		rest.RenderJSON(w, &echo)
	})

	srv := http.Server{Addr: opts.Listen,
		Handler:           router,
		ReadHeaderTimeout: time.Second * 30,
		WriteTimeout:      time.Second * 30,
		IdleTimeout:       time.Second * 30,
	}

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("[WARN] shutdown failed, %v", err)
		}
	}()

	return srv.ListenAndServe()
}

func setupLog(dbg bool) {
	if dbg {
		log.Setup(log.Debug, log.CallerFile, log.CallerFunc, log.Msec, log.LevelBraces)
		return
	}
	log.Setup(log.Msec, log.LevelBraces)
}
