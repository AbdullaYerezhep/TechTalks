package main

import (
	"forum/pkg/controller"
	"forum/pkg/repository"
	"forum/pkg/server"
	"forum/pkg/service"
	"log"
	"os"
)

func main() {
	port := "8000"

	// infoLog - reports the program process
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)

	// errLog - reports errors
	errLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	// NewDB takes dbtype and dbname and returns *sql.DB
	db, err := repository.NewDB("sqlite3", "forum.db")
	if err != nil {
		errLog.Fatal(err.Error())
		return
	}
	infoLog.Println("Database creation: SUCSESS")

	// repository - is a layer of the project, which contains all database transactions
	// takes [data] from service and does transactions | repository[data] => transaction[data]
	repo := repository.New(db)

	// service - is use case layer of the project which contains all use cases scenaries for users
	// takes [data] from the business layer and transmits it to the lower layer | service[data] => repository[data]) => transaction[data]
	service := service.New(repo)

	// // handler - is business layer of the project which contains all business logic of the project
	// takes [data] from the client and transmits it to the lower layer | handler[data] => service[data] => repository[data] => transaction[data]
	handler := controller.New(infoLog, errLog, service)
	server := new(server.Server)
	if err := server.Run(port, handler.Router()); err != nil {
		log.Fatal(err.Error())
	}
}
