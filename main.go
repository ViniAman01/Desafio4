package main

import (
	"github.com/jung-kurt/gofpdf"
	"project/makePDF"
)

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")

	imagemcapa := "./img/estante01.jpg"
	title := "ESTANTE MÁGICA"
	subtitle := "Ler e Sonhar"
	author := "Equipe Estante Mágica"

	makePDF.Capa(pdf, title, subtitle, author, imagemcapa)
	makePDF.PageIntroduction(pdf)

	historyTitle := `Contos do GPT`
	history := `Era uma vez uma pequena cidade chamada Serenidade, onde as pessoas viviam suas vidas tranquilas. No centro da cidade, havia uma loja de presentes chamada "Surpresas Mágicas", conhecida por suas ofertas especiais. Em uma manhã ensolarada, Maria, uma jovem moradora de Serenidade, entrou na loja em busca de um presente de aniversário para sua avó. Ela examinou os prateleiras, mas nada parecia especial o suficiente para a pessoa que tanto amava. O proprietário da loja, Sr. Higgins, percebeu a hesitação de Maria e se aproximou com um sorriso amigável. Ele sugeriu que ela desse à avó um livro de histórias antigas, cheio de memórias e aventuras compartilhadas. Maria seguiu o conselho de Sr. Higgins e escolheu um livro de histórias. Ela escreveu uma mensagem amorosa na primeira página e entregou o presente à sua avó no dia de seu aniversário. À noite, as duas se aconchegaram na sala de estar, lendo as histórias juntas. A avó de Maria sorriu e disse que era o presente mais especial que já recebera, pois continha as lembranças de sua família. E assim, a história de Maria e sua avó se tornou parte das histórias antigas do livro, um testemunho do amor e da conexão entre as gerações.`

	makePDF.NewPagTextSimple(pdf, historyTitle, history)
	makePDF.NewPagImg(pdf, "./img/lendo_livro.jpg", "1.1 - Maria e sua Avó")

	pdf.OutputFileAndClose("exemplo.pdf")
}
