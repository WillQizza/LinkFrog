package http

import (
	"net/http"

	"github.com/willqizza/linkfrog/http/routers"
)

func StartServer() {
	router := routers.NewRouter()
	http.ListenAndServe("127.0.0.1:3000", router)
}
