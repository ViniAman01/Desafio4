package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

var pageWidth = 200.0
var pageHeight = 200.0

var marginLeft = 30.0
var marginRight = 30.0
var marginTop = 30.0

var numberPage int
var font = "times"

func PageCover(pdf *gofpdf.Fpdf, title string, imgPath string) {

	pdf.AddPage()
	numberPage++

	pdf.ImageOptions(imgPath, 0, 0, pageWidth, pageHeight, false, gofpdf.ImageOptions{}, 0, "")

	utf_8 := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetFont(font, "B", 36)
	pdf.SetTextColor(192, 192, 192)

	titleWidth := pdf.GetStringWidth(title)
	Width, _ := pdf.GetPageSize()
	center := (Width - titleWidth) / 2.0

	pdf.Text(center, 80, utf_8(title))
}

func PageBiography(pdf *gofpdf.Fpdf, author string) {

	pdf.AddPage()
	numberPage++

	utf_8 := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetFont(font, "I", 24)
	pdf.SetTextColor(0, 0, 0)

	titleWidth := pdf.GetStringWidth(utf_8(author))
	Width, _ := pdf.GetPageSize()
	center := (Width - titleWidth) / 2.0

	pdf.Text(center, 90, utf_8(author))
}

func PageIntroduction(pdf *gofpdf.Fpdf) {

	title := "LIVRO CRIADO REFERENTE AO DESAFIO 4"
	introducao := `Este livro foi produzido pela equipe de estagiários da Estante Mágica. Agradecemos a todos pelo bom trabalho, esperamos que o livro esteja de acordo com o desafio proposto.`

	pdf.AddPage()
	numberPage++

	pdf.SetMargins(marginLeft, marginTop, marginRight)

	utf_8 := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetFont(font, "I", 20)
	pdf.SetTextColor(0, 0, 0)

	titleWidth := pdf.GetStringWidth(title)
	Width, _ := pdf.GetPageSize()
	center := (Width - titleWidth) / 2.0

	pdf.Text(center, marginTop, title)

	pdf.SetFont(font, "", 16)
	pdf.Text(center, marginTop+10, "________________________________________________")

	pdf.SetY(marginTop + 30)
	pdf.MultiCell(0, 6, utf_8(introducao), "", "", false)

	pdf.Image("./img/logo.jpeg", 80, 130, 40, 40, false, "", 0, "")

	pdf.SetFont(font, "", 14)
	pdf.Text(100, 190, fmt.Sprint(numberPage))
}

func NewPageText(pdf *gofpdf.Fpdf, text string) {

	pdf.AddPage()
	numberPage++

	pdf.SetMargins(marginLeft, marginTop, marginRight)

	utf_8 := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetFont(font, "", 16)
	pdf.SetTextColor(0, 0, 0)

	pdf.SetY(marginTop)
	pdf.MultiCell(0, 5.5, utf_8(text), "", "", false)

	pdf.SetFont(font, "", 14)
	pdf.Text(100, 190, fmt.Sprint(numberPage))
}

func NewPageImg(pdf *gofpdf.Fpdf, imgPath string) {
	pdf.AddPage()
	numberPage++
	pdf.ImageOptions(imgPath, 0, 0, pageWidth, pageHeight, false, gofpdf.ImageOptions{}, 0, "")

	pdf.SetFont(font, "", 14)
	pdf.Text(100, 190, fmt.Sprint(numberPage))
}

func main() {

	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		Size: gofpdf.SizeType{
			Wd: pageWidth,
			Ht: pageHeight,
		},
	})

	PageCover(pdf, "A fim de testes", "./img/estante01.jpg")
	PageBiography(pdf, "Alguem qualquer")
	PageIntroduction(pdf)
	NewPageText(pdf, "Ontem foi um dia daqueles que o mundo parecia se encontrar em perfeita harmonia, como eu poderia descrever a sensação de estar tudo dando certo. Para aqueles que um dia pensaram em desistir, saiba que vale a pena, mesmo se não valer, poderá dizer que aguentou até o fim.")
	NewPageImg(pdf, "./img/lendo_livro.jpg")

	err := pdf.OutputFileAndClose("exemplo.pdf")

	if err != nil {
		fmt.Println(err)
	}
}
