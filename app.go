package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"training-backend/package/log"
	webserver "training-backend/server"
	"training-backend/services/database"

	"training-backend/package/config"
)

func init() {
	// set timezone
	tzZone, _ := time.LoadLocation("Africa/Dar_es_Salaam") //10800
	time.Local = tzZone

	path := config.LoggerPath()

	log.SetOptions(
		log.Development(),
		log.WithCaller(false),
		log.WithLogDirs(path),
	)

}
func main() {
	//postgresql database
	database.Connect()
	defer database.Close() //close database pool TODO: add this into graceful app close function

	//start webserver
	go webserver.StartWebserver()

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)
	<-stop
	log.Infoln("notification is shutting down...  👋 !")
	fmt.Println("notification is shutting down .... 👋 !")
	database.Close()

	go func() {
		<-stop
		log.Fatalln("notification is terminating...")
	}()

	defer os.Exit(0)
}
