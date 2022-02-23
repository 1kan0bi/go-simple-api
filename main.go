package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

const (
	UrlExample = "postgres://username:password@adress:port/db" //connect to Postgres credents
)

type Product struct { //productlist
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Count       int    `json:"count"`
	Description string `json:"description"`
}

type Productinfo struct { //info about product
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Count       int64  `json:"count"`
	Description string `json:"description"`
}

func connectToDB(urlExample string) *pgx.Conn { //connect to DB

	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func GetProductInfo(w http.ResponseWriter, r *http.Request) { //

	conn := connectToDB(UrlExample)
	var product []*Productinfo
	params := mux.Vars(r)
	id := params["id"]
	err := pgxscan.Select(context.Background(), conn, &product, `SELECT name, price, count, description FROM products where id=$1`, id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	product_json, _ := json.Marshal(product)
	byteError := []byte("null")

	if !bytes.Equal(product_json, byteError) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(product_json)
	} else {
		w.WriteHeader(404)
	}
}

func GetProductList(w http.ResponseWriter, r *http.Request) {

	conn := connectToDB(UrlExample)
	var products []*Product
	err := pgxscan.Select(context.Background(), conn, &products, `SELECT * FROM products`)

	if err != nil {
		//w.WriteHeader(500)
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	products_json, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json")
	w.Write(products_json)

	defer conn.Close(context.Background())
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/product/{id}", GetProductInfo).Methods("GET")
	router.HandleFunc("/v1/products", GetProductList).Methods("GET")
	fmt.Println("App started at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

