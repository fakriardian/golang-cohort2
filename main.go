package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db  *sql.DB
	err error
)

type Book struct {
	ID     int
	Title  string
	Author string
	Stock  int
}

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Variants  []Variant `json:"variants"`
}

type Variant struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	ProductID int       `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlInfo := os.Getenv("DB_URL")

	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	createProduct := Product{
		Name: "Baju",
	}

	updateProduct := Product{
		ID:   2,
		Name: "Celana di Update",
	}

	createVariant := Variant{
		Name:      "Warna Jingga",
		Quantity:  15,
		ProductID: 2,
	}

	updateVariant := Variant{
		ID:        2,
		Name:      "Warna Biru di Update",
		Quantity:  16,
		ProductID: 3,
	}

	CreateProduct(&createProduct)
	UpdateProduct(&updateProduct)
	GetProductById(2)
	CreateVariant(&createVariant)
	UpdateVariantById(&updateVariant)
	GetProductWithVariant()
	DeleteVariantById(4)
}

func CreateProduct(product *Product) {
	sqlStatement := `INSERT INTO products (name) VALUES (?)`
	result, err := db.Exec(sqlStatement, product.Name)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	var data = Product{}
	sqlRetrieve := `SELECT * FROM products WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, lastInsertID).Scan(&data.ID, &data.Name, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		panic(err)
	}

	response, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Printf("New product Data: %+v\n", string(response))
}

func UpdateProduct(product *Product) {
	sqlStatement := `UPDATE products SET name = ? WHERE id = ?;`
	res, err := db.Exec(sqlStatement, product.Name, product.ID)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Product Updated data amount: ", count)
}

func GetProductById(id int) {
	var data = Product{}
	sqlRetrieve := `SELECT * FROM products WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, id).Scan(&data.ID, &data.Name, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		panic(err)
	}

	response, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Printf("Product Data: %+v\n", string(response))
}

func CreateVariant(variant *Variant) {
	sqlStatement := `INSERT INTO variants (name, quantity, product_id) VALUES (?, ?, ?)`
	result, err := db.Exec(sqlStatement, variant.Name, variant.Quantity, variant.ProductID)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	var data = Variant{}
	sqlRetrieve := `SELECT * FROM variants WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, lastInsertID).Scan(&data.ID, &data.Name, &data.Quantity, &data.ProductID, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		panic(err)
	}

	response, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Printf("New variant Data: %+v\n", string(response))
}

func UpdateVariantById(variant *Variant) {
	sqlStatement := `UPDATE variants SET name = ?, quantity = ?, product_id = ? WHERE id = ?;`
	res, err := db.Exec(sqlStatement, variant.Name, variant.Quantity, variant.ProductID, variant.ID)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Variant Updated data amount: ", count)
}

func DeleteVariantById(id int) {
	sqlStatement := `DELETE from variants WHERE id = ?;`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted variant data amount:", count)
}

func GetProductWithVariant() {
	productRows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer productRows.Close()

	variantRows, err := db.Query("SELECT * FROM variants")
	if err != nil {
		log.Fatal(err)
	}
	defer variantRows.Close()

	productMap := make(map[int]*Product)

	for productRows.Next() {
		var product Product
		if err := productRows.Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		productMap[product.ID] = &product
	}

	for variantRows.Next() {
		var variant Variant
		if err := variantRows.Scan(&variant.ID, &variant.Name, &variant.Quantity, &variant.ProductID, &variant.CreatedAt, &variant.UpdatedAt); err != nil {
			log.Fatal(err)
		}

		product := productMap[variant.ProductID]
		if product != nil {
			product.Variants = append(product.Variants, variant)
		}
	}

	var products []Product
	for _, product := range productMap {
		products = append(products, *product)
	}

	response, err := json.Marshal(products)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Printf("Product Data wit variant: %+v\n", string(response))
}
