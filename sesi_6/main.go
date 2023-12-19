package main

import (
	"log"
	"os"
	"sesi_6/product"
	"sesi_6/repository"
	"sesi_6/server"
	"sesi_6/warehouse"

	"github.com/joho/godotenv"
)

/*
Challenge Asset Management Part I

Kamu ditugaskan oleh sebuah perusahaan retail untuk membuat sebuah REST API untuk aplikasi manajemen aset,
dimana kamu harus handle stock inventory disetiap gudang yang perusahaan tersebut miliki.
Aplikasi yang akan kamu buat harus memiliki fitur untuk CRUD data gudang,
dan juga CRUD data produk yang ada pada setiap gudang nya,
Pada tahap ini kamu tidak perlu langsung mengintegrasikan aplikasi dengan database,
kamu cukup menyimpan data sementara pada sebuah variable saja.
Data gudang yang diperlukan adalah, nama gudang, alamat
Berikut adalah beberapa list endpoint warehouse yang perlu kamu buat:

POST /warehouses -> membuat data gudang baru
GET /warehouses -> menampilkan semua list gudang
GET /warehouses/:id -> menampilkan data detail gudang, termasuk jumlah barang yang tersimpan di gudang tersebut
PUT /warehouses/:id -> memperbaharui data gudang
DELETE /warehouses/:id -> menghapus data gudang

Data produk yang diperlukan adalah, nama, jenis produk(makanan/minuman/baju/dll), jumlah stock, harga,
dan di gudang mana produk tersebut disimpan
Berikut adalah beberapa list endpoint warehouse yang perlu kamu buat:

POST /products -> membuat data produk baru
GET /products -> menampilkan semua list produk
GET /products/:id -> menampilkan data detail produk
PUT /products/:id -> memperbaharui data produk
DELETE /products/:id -> menghapus data produk
*/

func getEnv() {
	err := godotenv.Load()
	if err == nil {
		log.Println("Reading from .env file")
	}
}

// @title Warehouse APIs
// @version 1.0
// @description Warehouse APIs.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @securityDefinitions.apiKey JWT
// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @schemes http
func main() {
	getEnv()

	repo, err := repository.New()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	product := product.New(repo.ProductRepo)
	warehouses := warehouse.New(repo.WarehouseRepo)
	server.Start(warehouses, product)
}
