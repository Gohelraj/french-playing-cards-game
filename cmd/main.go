package main

import (
	"fmt"
	"github.com/Gohelraj/french-playing-cards-game/db"
	"log"
	"net/http"
	"time"

	"github.com/Gohelraj/french-playing-cards-game/config"
	"github.com/Gohelraj/french-playing-cards-game/restapiv1/routes"
)

func main() {
	var err error
	config.Conf, err = config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := db.Connect()
	defer db.Close()

	port := fmt.Sprintf(":%s", config.Conf.Server.Port)
	// Start the server
	srv := &http.Server{
		Addr:    port,
		Handler: routes.InitializeRouter(db),
		// IdleTimeout is the maximum amount of time to wait for the
		// next request when keep-alives are enabled.
		IdleTimeout: 10 * time.Minute,
	}

	log.Printf("Server listening on port: %s", port)
	log.Fatal(srv.ListenAndServe())
}
