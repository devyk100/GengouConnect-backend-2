package main

import (
	"backend/cmd/db"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"net/http"
)

type Something struct {
	Student string `json:"student"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
	r.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		var something Something
		err := json.NewDecoder(request.Body).Decode(&something)
		fmt.Println(something.Student, "IS THE BODY")

		if err != nil {
			return
		}
		_, err = writer.Write([]byte(`{"message": "Hello World"}`))
		//c := request.Context().Done()
		//select {
		//case <-c:
		//	fmt.Println("DONE")
		//}
		if err != nil {
			return
		}
		//write, err := writer.Write([]byte("Hello World"))?
		//fmt.Printf("%v", write)
		//var p []byte
		//p := make([]byte, 100)
		//read, err := request.Body.Read(p)
		//if err != nil {
		//	return
		//}
		fmt.Printf("%v", "DONE")
		if err != nil {
			return
		}
	})
	db.Connect()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
