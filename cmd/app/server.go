package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitServer(port string, router mux.Router) error {
	var (
		quitCh = make(chan os.Signal)
		srv    *http.Server
	)

	srv = &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      &router,
		ReadTimeout:  time.Second * 120,
		WriteTimeout: time.Second * 120,
	}

	signal.Notify(quitCh, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	go graceShutdown(quitCh, srv)

	log.Printf("server: starting on 127.0.0.1%s\n", port)
	return srv.ListenAndServe()

}

func graceShutdown(quitCh chan os.Signal, srv *http.Server) {
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		dur        time.Time
		err        error
	)

	s := <-quitCh
	log.Printf("server: received signal %+v\n", s)

	dur = time.Now().Add(30 * time.Second) // dummy deadline
	ctx, cancelFunc = context.WithDeadline(context.Background(), dur)
	defer cancelFunc()

	err = srv.Shutdown(ctx)
	if err != nil {
		log.Panicln("server: couldn't shutdown because of " + err.Error())
	}
}
