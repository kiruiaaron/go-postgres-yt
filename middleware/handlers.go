package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kiruiaaron/go-postgres-yt/models"
	"github.com/vertica/vertica-sql-go/msgs"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres")

	return db

}

func CreateStock(w http.ResponseWriter, r *http.Request){

	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil{
		log.Fatal("Unable to decode the request body. %v", err)
	}

	insertID := insertStock(stock)

	res := response{
		ID: insertID,
		Message: "stock created successfully",
	}

	json.NewEncoder(w).Encode(res)


}

func GetStock(w http.ResponseWriter, r*http.Request){

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil{
		log.Fatal("Unable to convert the string into int. %v", err)
	}

	stock, err := getStock(int64(id))

	if err != nil{
		log.Fatal("Unable tp get the stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)

}

func GetAllStocks(w http.ResponseWriter, r*http.Request){
	stocks, err := getAllStocks()

	if err != nil{
		log.Fatalf("Unable to get all the stocks %v", err)
	}
	json.NewEncoder(w).Encode(stocks)
      
}

func UpdateStock(w http.ResponseWriter, r*http.Request){
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil{
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	var stock models.Stock


	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil{
		log.Fatalf("Unable to decode the request body. %v", err)

	}

	updateRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("Stock updated successfully. Total rows/records affected %v", updateRows)
	res:= response{
		ID:int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r*http.Request){

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"])

	if err != nil{
		log.Fatalf("Unable to convert string to int %v", err)
	}

	deleteRows := deleteStock(int64(id))
	msg := fmt.Sprintf("Stock deleted successfully. Total rows/records %v", deleteRows)
    
	res :=  response{
		ID: int64(id),
		Message: msg,

	}

	json.NewEncoder(w).Encode(res)
}