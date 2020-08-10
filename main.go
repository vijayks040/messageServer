package main

import (
	"context"
	"database/sql"
	con "messageServer/db"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

var db *sql.DB
var message con.DbProp

//initiating DB connection
// func init() {
// 	log.Println("initalizing DB...")
// 	message.DbName = "postgres"
// 	message.Host = "localhost"
// 	message.Pass = "admin"
// 	message.Port = 5432
// 	message.User = "postgres"
// 	db, _ = con.GetConn(message)
// 	log.Println("initialized...")
// }
// func pingDb() {
// 	err := db.Ping()
// 	if err != nil {
// 		log.Println("DB Ping failed...")
// 	}
// 	log.Println("Db ping successfull.....")
// }

//starting the http server...
func StartHTTPServer() {
	log.Info("Start HTTP Server...")

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/addMessage", AddMessage).Methods("POST")
	r.HandleFunc("/listMessage", ListMessage).Methods("GET")
	r.HandleFunc("/getOneMessage", GetOneMessage).Methods("GET")
	r.HandleFunc("/deleteOneMessage", DeleteOneMessage).Methods("DELETE")

	// Bind to a port and pass our router in
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
func main() {
	StartHTTPServer()
	//go pingDb()
}
