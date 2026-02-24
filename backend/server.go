package backend

import (
	"net/http"

	"github.com/willqizza/linkfrog/backend/routers"
)

func StartServer() {
	router := routers.NewRouter()
	http.ListenAndServe(":8080", router)
}
