package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sam0307/go-bookstore/pkg/routers"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterBookStoreRouter(r)
	http.Handle("/", r)
	fmt.Println("Welcome to Library...")
	log.Fatal(http.ListenAndServe("127.0.0.1:6060", r))
}
