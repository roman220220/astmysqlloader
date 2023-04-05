package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/roman220220/astmysqlloader/app/internal/db"
	"github.com/roman220220/astmysqlloader/app/internal/config"

	"net/http"
	"regexp"

	log "astmysqlloader/app/internal/logger"
)

type DialerParams struct {
        Action        string
        ProjectId     string
        Sql           db.DB
        Type          string
}

func main() {
	defer recoverAll()
	ws := &sync.WaitGroup{}
	var a Params
	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("Program killed !")
		a.Sql.DBClose()
		close(a.AgentsChannel)
		close(a.Channel)
		close(a.Ch)
		// do last actions and wait for all write operations to end
		os.Exit(0)
	}()
	a.Sql.ConnectDB()
	ws.Add(1)
	ws.Wait()
}
func recoverAll() {
	if err := recover(); err != nil {
		log.MakeLog(3, err)
	}
}
