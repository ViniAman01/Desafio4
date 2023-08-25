package main

import (
	"github.com/jung-kurt/gofpdf"
)

var larguraPagina = 210.0
var alturaPagina = 297.0

var margemEsquerda = 30.0
var margemDireita = 30.0
var margemSuperior = 30.0
var margemInferior = 30.0

func capa(pdf *gofpdf.Fpdf, title string, subtitle string, author string, img string) {

	pdf.AddPage()
	pdf.ImageOptions(img, 0, 0, larguraPagina, alturaPagina, false, gofpdf.ImageOptions{}, 0, "")

	pdf.SetTextColor(192, 192, 192)
	pdf.SetFont("Arial", "B", 36)
	pdf.Text(15, 80, title)

	pdf.SetTextColor(200, 100, 30)
	pdf.SetFont("Arial", "B", 18)
	pdf.Text(85, 90, subtitle)

	pdf.SetTextColor(170, 100, 10)
	pdf.SetFont("Arial", "I", 16)
	pdf.Text(85, 250, author)

	return
}

func newPagTextSimple(pdf *gofpdf.Fpdf, text string) {

	pdf.AddPage()
	pdf.SetMargins(margemEsquerda, margemSuperior, margemDireita)

	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.MultiCell(0, 5, text, "", "", false)
}

func newPagImg(pdf *gofpdf.Fpdf, imgPath string, description string) {

	pdf.AddPage()
	pdf.SetMargins(margemEsquerda, margemSuperior, margemDireita)

	imagePath := imgPath

	pdf.Image(imagePath, 25, margemSuperior+20, 160, 180, false, "", 0, "")
	pdf.SetFont("Arial", "I", 12)
	pdf.Text(25, 240, description)
}

func addBlockText(pdf *gofpdf.Fpdf, text string, y float64, yPage float64) {

	pdf.SetMargins(margemEsquerda, margemSuperior, margemDireita)

	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(10, 0, 0)
	pdf.SetY(yPage)
	pdf.MultiCell(0, y, text, "", "L", false)

}

func main() {
	text := `"Posso estar ficando meio confuso sobre tudo isso, mas muitas vezes sinto que ja nao te alcancarei. Peco desculpas para o vento quase todas as noites de lua minguante, sabe-se la o que aconteceu com voce, prefiro nao saber."
	
	Aqui fica minhas sinceras desculpas a mim mesmo..`

	novotext := `"A funcao pdf.Image na biblioteca gofpdf e usada para adicionar imagens a um documento PDF. Ela permite que voce insira imagens em uma pagina PDF, especificando o arquivo da imagem, suas dimensoes, e a posicao na pagina onde a imagem deve ser colocada."`

	pdf := gofpdf.New("P", "mm", "A4", "")

	imagemcapa := "./img/capa.jpg"
	title := "SABIA QUE NAO ERA ASSIM"
	subtitle := "Eu Me Avisei"
	author := "Alguem Qualquer"

	capa(pdf, title, subtitle, author, imagemcapa)

	pdf.AddPage()
	pdf.SetFont("Arial", "I", 16)
	pdf.SetTextColor(0, 0, 0)
	pdf.Text(65, 45, "SABIA QUE NAO ERA ASSIM")
	pdf.Text(90, 55, "MAS FOI")
	pdf.Text(50, 65, "___________________________________")

	addBlockText(pdf, text, 5, 100)
	addBlockText(pdf, novotext, 5, 130)
	newPagImg(pdf, "./img/jujutsu-kaisen.jpg", "1.1 - Guerra do vietna.")

	newPagTextSimple(pdf, novotext)
	newPagImg(pdf, "./img/baby-yoda.jpg", "1.2 - Aquele que dizimou a Terra.")
	newPagTextSimple(pdf, text)

	pdf.OutputFileAndClose("exemplo.pdf")
}
