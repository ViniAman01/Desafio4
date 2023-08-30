package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"project/dbConnection"
	"project/makePDF"
	"project/models"
)

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")

	imagemcapa := "./files/estante01.jpg"
	title := "ESTANTE MÁGICA"
	subtitle := "Ler e Sonhar"
	author := "Equipe Estante Mágica"

	makePDF.Capa(pdf, title, subtitle, author, imagemcapa)
	makePDF.PageIntroduction(pdf)

	historyTitle := ""

	client := dbConnection.ConnectDB()
	coll := dbConnection.GetCollection("mainDB2", "collteste", client)

	for _, description_page := range models.DescriptionsPage {
		media, err := dbConnection.FindDoc(coll, "description_page", description_page)
		if err == nil {
			fmt.Println(media.Data_type)
			if media.Data_type == "text" {
				makePDF.NewPagTextSimple(pdf, historyTitle, media.Data)
			}
			if media.Data_type == "image" {
				fmt.Println(media.Data)
				makePDF.NewPagImg(pdf, media.Data, "Uma legenda")
			}
		}
	}

	pdf.OutputFileAndClose("exemplo.pdf")
}
