package backend

import (
	"log"
	"net/http"

	"github.com/willqizza/linkfrog/backend/db"
	"github.com/willqizza/linkfrog/backend/routers"
)

func StartServer() {
	if err := db.Init(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	router := routers.NewRouter()
	http.ListenAndServe(":8080", router)
}
