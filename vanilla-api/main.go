package main

import "net/http"

func main() {

	db := NewDB("postgres", "user=root dbname=dev password=root sslmode=disable")

	router := NewRouter(db)

	server := http.Server{
		Addr:    "8080",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
