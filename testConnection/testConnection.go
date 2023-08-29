package main

import (
	"fmt"
	"project/dbConnection"
	"project/models"
)


func main(){
	client := dbConnection.ConnectDB()

	db := dbConnection.GetDB("dbtest", client)

	fmt.Println(db)

	coll := dbConnection.GetCollection("dbtest","collteste",client)

	fmt.Println(coll)

	dbConnection.InsertColl(coll,models.Media{Data_type: "teste1",Code_book: "teste1",Description_page: "teste1", Data: "teste1"})
}