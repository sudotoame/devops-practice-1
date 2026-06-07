package main

import (
	"github.com/sudotoame/devops-practice-1/logic"
	"github.com/sudotoame/devops-practice-1/server"
)

func main() {
	logics := logic.NewLogic()
	handlers := server.NewHandlers(logics)
	server := server.NewServer(handlers)

	server.Run()
}
