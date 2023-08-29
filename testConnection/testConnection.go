package main

import (
	"project/dbConnection"
	"project/models"
)

var media = models.Media{
	Data_type: "tipoA",
	Code_book: "12345",
	Description_page: "cover_page",
	Data: "informações",
}

var media2 = models.Media{
	Data_type: "tipoB",
	Code_book: "12345",
	Description_page: "page2",
	Data: "informações",
}

var media3 = models.Media{
	Data_type: "tipoC",
	Code_book: "12345",
	Description_page: "page3",
	Data: "informações",
}

var media4 = models.Media{
	Data_type: "tipoD",
	Code_book: "12345",
	Description_page: "page4",
	Data: "informações",
}

var medias = []models.Media{media,media2,media3,media4}

func main(){
	client := dbConnection.ConnectDB()

	coll := dbConnection.GetCollection("mainDB2","collteste",client)

	for _,media := range medias {
		dbConnection.InsertColl(coll,media)
	}
}