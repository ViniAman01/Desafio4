package controllers

import (
	"github.com/jung-kurt/gofpdf"
  "fmt"
  "project/models"
	db "project/database"
)

var pdf = gofpdf.New("P", "mm", "A4", "")

var client = db.ConnectDB()
var coll = db.GetCollection("mainDB2", "collteste", client)

func CoverHandler() {
	imagemcapa := "./static/estante01.jpg"
	title := "ESTANTE MÁGICA"
	subtitle := "Ler e Sonhar"
	author := "Equipe Estante Mágica"

	Capa(pdf, title, subtitle, author, imagemcapa)
	PageIntroduction(pdf)
}

func PageHandler() {
	for _, description_page := range models.DescriptionsPage {
		media, err := db.FindDoc(coll, "description_page", description_page)

		if err == nil {
			fmt.Println(media.Data_type)

			if media.Data_type == "text" {
				NewPagTextSimple(pdf, " ", media.Data)
			}

			if media.Data_type == "image" {
				fmt.Println(media.Data)
				NewPagImg(pdf, media.Data, "Uma legenda")
			}
		}
	}

  err := pdf.OutputFileAndClose("exemplo.pdf")

	if err != nil {
		fmt.Println(err)
	}
}
