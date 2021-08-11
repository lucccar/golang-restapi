package api

import (
	"api/router"
	"config"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	fmt.Printf("\n\tListening [::]:%d", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
