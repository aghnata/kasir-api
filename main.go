package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper" //untuk baca env
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DbConn string `mapstructure:"DB_CONN"`
}

func main() {

	viper.AutomaticEnv() //baca dari env secara otomatis
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DbConn: viper.GetString("DB_CONN"),
	}

	//setup db
	db, err := database.InitDB(config.DbConn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	//setup routes
	http.HandleFunc("/api/products", productHandler.HandleProducts)

	//GET localhost:8080/api/produk/{id}
	//PUT localhost:8080/api/produk/{id}
	//DELETE localhost:8080/api/produk/{id}
	// http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		getProdukByID(w, r)
	// 	} else if r.Method == "PUT" {
	// 		updateProdukByID(w, r)
	// 	} else if r.Method == "DELETE" {
	// 		deleteProdukByID(w, r)
	// 	}
	// })

	//GET localhost:8080/api/produk
	//POST localhost:8080/api/produk
	// http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(produk)
	// 	} else if r.Method == "POST" {
	// 		var produkBaru models.Produk
	// 		err := json.NewDecoder(r.Body).Decode(&produkBaru)
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusBadRequest)
	// 			return
	// 		}
	// 		produkBaru.ID = len(produk) + 1
	// 		produk = append(produk, produkBaru)
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(produkBaru)
	// 	}
	// })

	// /health endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})
	fmt.Println("Starting server on :" + config.Port)

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
