package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FelipeMCassiano/Apostoli/cg"
	"github.com/FelipeMCassiano/Apostoli/uploader/deploy"
)

func main() {
	err := cg.LoadConfigs()
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()
	router.HandleFunc("POST /deploy", deploy.Deploy())

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("running uploader at port 8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
