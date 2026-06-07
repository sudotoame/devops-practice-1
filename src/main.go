package main

import (
	"devops-practice-1/logic"
	"devops-practice-1/server"
	"log"
)

func main() {
	logics := logic.NewLogic()
	handlers := server.NewHandlers(logics)
	server := server.NewServer(handlers)

	if err := server.Run(); err != nil {
		log.Println(err)
		return
	}
}
