package controllers

import (
	"fmt"
	db "project/database"
	"project/models"

	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/bson"
)

var pdf = gofpdf.New("P", "mm", "A4", "")

var client = db.ConnectDB()
var coll = db.GetCollection("mainDB2", "collteste", client)

func CoverHandler() {
	imagemcapa := "./static/estante01.jpg"
	title := "ESTANTE MÁGICA"

	Capa(pdf, title, imagemcapa)
	PageIntroduction(pdf)
}

func PageHandler(code_book string) {

	for _, description_page := range models.DescriptionsPage {
		filter := bson.D{
			{"description_page", description_page},
			{"code_book", code_book},
		}
		media, err := db.FindDoc(coll, filter)

		if err == nil {
			fmt.Println(media.Data_type)

			if media.Data_type == "text" {
				NewPagTextSimple(pdf, media.Data)
			}

			if media.Data_type == "image" {
				fmt.Println(media.Data)
				NewPagImg(pdf, media.Data, "Uma legenda")
			}
		} else {
			fmt.Println(err)
		}
	}

	err := pdf.OutputFileAndClose("exemplo.pdf")

	if err != nil {
		fmt.Println(err)
	}
}
