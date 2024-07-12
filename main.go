package main

import (
	"main/server"

	_ "main/docs"
)


func main() {
	cm := server.CommandManager{}

	cm.AddCommand(server.Command{
		Script: "",
		Desc:   "Run backend server",
		ExecuteFunc: func(args []string) {
			apiServer := server.Initialize()
			apiServer.Start()
		},
	})

	cm.Execute()
}
