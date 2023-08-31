package controllers

import (
	"fmt"
	db "project/database"
	"project/models"

	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/bson"
)

var pdf = gofpdf.NewCustom(&gofpdf.InitType{
	Size: gofpdf.SizeType{
		Wd: pageWidth,
		Ht: pageHeight,
	},
})

var client = db.ConnectDB()
var coll = db.GetCollection("mainDB2", "collteste", client)

func CoverHandler() {
	imagemcapa := "./static/estante01.jpg"
	title := "ESTANTE MÁGICA"

	PageCover(pdf, title, imagemcapa)
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
				NewPageText(pdf, media.Data)
			}

			if media.Data_type == "image" {
				fmt.Println(media.Data)
				NewPageImg(pdf, media.Data)
			}
		} else {
			fmt.Println(err)
		}
	}

	PageBiography(pdf, "Autor desconhecido")

	err := pdf.OutputFileAndClose("exemplo.pdf")

	if err != nil {
		fmt.Println(err)
	}
}
