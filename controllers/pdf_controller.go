package controllers

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
	introduction := `Este livro foi produzido pela equipe de estagiários da Estante Mágica. Agradecemos a todos pelo bom trabalho, esperamos que o livro esteja de acordo com o desafio proposto.`

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
	pdf.MultiCell(0, 6, utf_8(introduction), "", "", false)

	pdf.Image("./static/logo.jpeg", 80, 130, 40, 40, false, "", 0, "")

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
