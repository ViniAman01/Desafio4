package makePDF

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

var pageWidth = 210.0
var pageHeight = 297.0

var marginLeft = 30.0
var marginRight = 30.0
var marginTop = 30.0
var marginLow = 30.0

var numberPage = 1
var font = "times"

func Capa(pdf *gofpdf.Fpdf, title string, subtitle string, author string, img string) {

	pdf.AddPage()
	pdf.ImageOptions(img, 0, 0, pageWidth, pageHeight, false, gofpdf.ImageOptions{}, 0, "")

	utf_8 := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetTextColor(192, 192, 192)
	pdf.SetFont(font, "B", 36)
	titleWidth := pdf.GetStringWidth(title)
	Width, _ := pdf.GetPageSize()
	center := (Width - titleWidth) / 2.0
	pdf.Text(center, 80, utf_8(title))

	pdf.SetTextColor(200, 100, 30)
	pdf.SetFont(font, "B", 26)
	titleWidth = pdf.GetStringWidth(subtitle)
	center = (Width - titleWidth) / 2.0
	pdf.Text(center, 90, utf_8(subtitle))

	pdf.SetTextColor(170, 100, 10)
	pdf.SetFont(font, "I", 16)
	titleWidth = pdf.GetStringWidth(author)
	center = (Width - titleWidth) / 2.0
	pdf.Text(center, 250, utf_8(author))

}

func PageIntroduction(pdf *gofpdf.Fpdf) {

	title := "LIVRO CRIADO REFERENTE AO DESAFIO 4"
	introducao := `Este livro foi produzido pela equipe de estagiários da Estante Mágica. Agradecemos a todos pelo bom trabalho, esperamos que o livro esteja de acordo com o desafio proposto.`

	pdf.AddPage()
	pdf.SetMargins(marginLeft, marginTop, marginRight)

	numberPage++
	pdf.SetFont(font, "I", 20)
	pdf.SetTextColor(0, 0, 0)

	titleWidth := pdf.GetStringWidth(title)
	Width, _ := pdf.GetPageSize()
	center := (Width - titleWidth) / 2.0

	pdf.Text(center, 45, title)
	pdf.Text(center, 60, "_______________________________________")

	AddBlockText(pdf, introducao, 5.5, 90)

	pdf.Image("./img/logo.jpeg", 86.5, 240, 40, 40, false, "", 0, "")

	pdf.SetFont(font, "", 12)
	pdf.Text(pageWidth/2, 290, fmt.Sprint(numberPage))
}

func NewPagTextSimple(pdf *gofpdf.Fpdf, title string, text string) {

	pdf.AddPage()

	utf_8 := pdf.UnicodeTranslatorFromDescriptor("")

	numberPage++
	pdf.Text(105, 290, fmt.Sprint(numberPage))
	pdf.SetMargins(marginLeft, marginTop, marginRight)

	pdf.SetFont(font, "B", 32)
	pdf.SetTextColor(0, 0, 0)

	titleWidth := pdf.GetStringWidth(title)
	Width, _ := pdf.GetPageSize()

	center := (Width - titleWidth) / 2.0

	pdf.Text(center, 50, utf_8(title))

	pdf.SetFont(font, "", 12)
	pdf.SetY(70)
	pdf.MultiCell(0, 5.5, utf_8(text), "", "", false)
}

func NewPagImg(pdf *gofpdf.Fpdf, imgPath string, description string) {

	pdf.AddPage()

	utf_8 := pdf.UnicodeTranslatorFromDescriptor("")

	numberPage++
	pdf.Text(105, 290, fmt.Sprint(numberPage))
	pdf.SetMargins(marginLeft, marginTop, marginRight)

	pdf.Image(imgPath, 25, marginLeft+20, 160, 180, false, "", 0, "")
	pdf.SetFont(font, "I", 12)
	pdf.Text(25, 240, utf_8(description))
}

func AddBlockText(pdf *gofpdf.Fpdf, text string, y float64, yPage float64) {

	utf_8 := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetMargins(marginLeft, marginTop, marginRight)

	pdf.SetFont(font, "", 12)
	pdf.SetTextColor(10, 0, 0)
	pdf.SetY(yPage)
	pdf.MultiCell(0, y, utf_8(text), "", "L", false)

}
