package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	db "project/database"
	mpdf "project/controllers"
	"project/models"
)

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")

	imagemcapa := "./files/estante01.jpg"
	title := "ESTANTE MÁGICA"
	subtitle := "Ler e Sonhar"
	author := "Equipe Estante Mágica"

	mpdf.Capa(pdf, title, subtitle, author, imagemcapa)
	mpdf.PageIntroduction(pdf)

	historyTitle := ""

	client := db.ConnectDB()
	coll := db.GetCollection("mainDB2", "collteste", client)

	for _, description_page := range models.DescriptionsPage {
		media, err := db.FindDoc(coll, "description_page", description_page)
		if err == nil {
			fmt.Println(media.Data_type)
			if media.Data_type == "text" {
				mpdf.NewPagTextSimple(pdf, historyTitle, media.Data)
			}
			if media.Data_type == "image" {
				fmt.Println(media.Data)
				mpdf.NewPagImg(pdf, media.Data, "Uma legenda")
			}
		}
	}

	pdf.OutputFileAndClose("exemplo.pdf")
}
