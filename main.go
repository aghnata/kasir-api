package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper" //untuk baca env
)

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

var produk = []Produk{
	{ID: 1, Nama: "Produk A", Harga: 10000, Stok: 50},
	{ID: 2, Nama: "Produk B", Harga: 20000, Stok: 30},
	{ID: 3, Nama: "Produk C", Harga: 15000, Stok: 20},
}

func getProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
}

func updateProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}
	var produkUpdate Produk
	err = json.NewDecoder(r.Body).Decode(&produkUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, p := range produk {
		if p.ID == id {
			produkUpdate.ID = id
			produk[i] = produkUpdate
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produkUpdate)
			return
		}
	}
	http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
}

func deleteProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}
	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Produk berhasil dihapus"})
			return
		}
	}
	http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
}

type Config struct {
	Port   string `mapstructure:"PORT"`
	DbConn string `mapstructure:"DB_CONN"`
}

func main() {
	// viper.SetConfigFile(".env")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	fmt.Println("Error reading config file", err)
	// 	return
	// }

	// var config Config
	// err = viper.Unmarshal(&config)
	// if err != nil {
	// 	fmt.Println("Error unmarshaling config", err)
	// 	return
	// }

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

	//GET localhost:8080/api/produk/{id}
	//PUT localhost:8080/api/produk/{id}
	//DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProdukByID(w, r)
		} else if r.Method == "DELETE" {
			deleteProdukByID(w, r)
		}
	})

	//GET localhost:8080/api/produk
	//POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

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
