package main

import (
	"app/api/users"
	"app/api/v1"
	"log"
	"net/http"
)

var server http.Server

func init() {
	router := httprouter.New(&httprouter.NewParams{
		MaxConcurrentRequests: 1000,
	})

	v1.InitEndpoints(router)
	users.InitEndpoints(router)

	server = httpauth.New(&httpauth.NewParams{
		Authenticator: auth.HttpDelegate,
		Next:          router,
	})
}

func main() {
	for {
		err := http.ListenAndServe(":80", server)
		if err != nil {
			log.Println(err)
		}
	}
}
