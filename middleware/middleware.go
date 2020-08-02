package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shopping-list/models"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func CreateConnection() (*sql.DB, error) {

	//os.Remove("./productos.db")
	db, err := sql.Open("sqlite3", ":memory:")

	sqlStmt := `create table if not exists product (id integer not null primary key, name text, amount int);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
	insertDb(db)

	return db, err

}

func initConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./productos.db")
	if err != nil {
		log.Fatalf("%q: \n", err)
	}
	return db, err
}

func insertDb(db *sql.DB) {

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into product(id, name,amount) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	amount := 3
	for i := 0; i < 5; i++ {

		_, err = stmt.Exec(i, fmt.Sprintf("Producto %03d", i), amount*i)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		log.Fatalf("No se pudo obtener el Body del request.  %v", err)
	}

	createdMsg := createProduct(product)

	res := models.Response{
		ID:      product.ID,
		Message: createdMsg,
	}

	json.NewEncoder(w).Encode(res)

}

func createProduct(product models.Product) string {
	db, err := initConnection()

	tx, _ := db.Begin()
	stmt, err := tx.Prepare("insert into product(id, name,amount) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Amount)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	msg := fmt.Sprintf("Se creo un nuevo producto con el ID: %v", product.ID)
	return msg
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	products, err := getAllProducts()

	if err != nil {
		log.Fatalf("No es posible obtener la lista de productos %v", err)
	}

	json.NewEncoder(w).Encode(products)
}

func getAllProducts() ([]models.Product, error) {

	db, err := initConnection()

	rows, err := db.Query("select id, name, amount from product order by id")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product

		err = rows.Scan(&product.ID, &product.Name, &product.Amount)

		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return products, err
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("No se puede convertir string en int.  %v", err)
	}

	product, err := oneProduct(int(id))

	if err != nil {
		log.Fatalf("No es posible obtener el producto %v", err)
	}

	json.NewEncoder(w).Encode(product)
}

func oneProduct(id int) (models.Product, error) {

	db, err := initConnection()

	stmt, err := db.Prepare("select id, name, amount from product where id =?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	var product models.Product

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Amount)
	if err != nil {
		log.Fatal(err)
	}
	return product, err
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("No se puede convertir string en int.  %v", err)
	}

	msg := deleteProduct(int(id))

	res := models.Response{
		ID:      int(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func deleteProduct(id int) string {

	sId := strconv.Itoa(id)

	db, err := initConnection()

	tx, _ := db.Begin()

	_, err = tx.Exec("delete from product where id= ?", sId)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	msg := fmt.Sprintf("Se elimino con exito el producto con el ID: %v", id)
	return msg
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("No se puede convertir string en int.  %v", err)
	}

	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		log.Fatalf("No se pudo obtener el Body del request.  %v", err)
	}

	updatedMsg := updateProduct(id, product)

	res := models.Response{
		ID:      id,
		Message: updatedMsg,
	}

	json.NewEncoder(w).Encode(res)

}

func updateProduct(id int, product models.Product) string {

	sAmount := strconv.Itoa(product.Amount)
	sId := strconv.Itoa(id)

	db, err := initConnection()

	tx, err := db.Begin()

	sql := `UPDATE product SET name=?, amount=? WHERE id=?`

	stmt, err := tx.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(product.Name, sAmount, sId)

	if err2 != nil {
		panic(err2)
	}
	tx.Commit()

	msg := fmt.Sprintf("Producto con el ID: %v actualizado exitosamente", id)
	return msg
}
