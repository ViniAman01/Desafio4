package main

import (
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

	for _, dp := range models.DescriptionsPage {
		media, err := dbConnection.FindDoc(coll, "description_page", dp)
		if err == nil {
			if media.Data_type == "text" {
				makePDF.NewPagTextSimple(pdf, historyTitle, media.Data)
			}
		}
	}

	pdf.OutputFileAndClose("exemplo.pdf")
}
