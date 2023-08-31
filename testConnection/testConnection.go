package main

import (
	db "project/database"
	"project/models"
	"strconv"
)

var text = []string{
	"Era uma vez, em um mundo repleto de magia e mistério, uma jovem chamada Elara. Elara tinha um desejo ardente de explorar além das fronteiras de sua pequena aldeia e desvendar os segredos escondidos na vastidão da floresta encantada.",
	"Elara se aventurou na floresta com determinação, empunhando um amuleto mágico que sua avó lhe dera. Dizia-se que o amuleto apontaria o caminho para a lendária Estrela Perdida, capaz de conceder um desejo a quem a encontrasse.",
	"À medida que Elara explorava a floresta, ela se deparou com criaturas mágicas e paisagens deslumbrantes. Mas também enfrentou desafios perigosos, provando sua coragem e inteligência.",
	"Em sua jornada, Elara conheceu um enigmático guardião da floresta, uma criatura sábia que a ajudou a decifrar os enigmas que protegiam o caminho para a Estrela Perdida.",
	"Finalmente, após muitas provações, Elara chegou ao local onde a Estrela Perdida deveria estar. No entanto, ela percebeu que a estrela não estava lá. Apenas um pergaminho antigo revelava a verdadeira natureza da busca.",
	"O pergaminho explicava que a verdadeira Estrela Perdida não era um objeto físico, mas sim a jornada em si. Elara percebeu que sua jornada a transformou, ensinando-a sobre coragem, amizade e autoconhecimento.",
	"Com um coração cheio de gratidão e sabedoria recém-encontrada, Elara retornou à sua aldeia. Ela compartilhou suas histórias e lições com os moradores, inspirando-os a também seguirem seus próprios sonhos.",
}

var path = []string{
	"./static/imagem1.jpg",
	"./static/imagem2.jpg",
	"./static/imagem3.jpg",
	"./static/imagem4.jpg",
	"./static/imagem5.jpg",
	"./static/imagem6.jpg",
	"./static/imagem7.jpg",
}

var media = models.Media{
	Data_type: "text",
	Code_book: "12345",
	// Description_page: "page1",
	// Data: ,
}

var media2 = models.Media{
	Data_type: "image",
	Code_book: "12345",
	// Description_page: "page2",
	// Data: ,
}

var medias = []models.Media{media, media2}

func main() {
	client := db.ConnectDB()

	coll := db.GetCollection("mainDB2", "collteste", client)
  indexText := 0
  indexImg := 0
	for i := 1; i <= 14; i++ {
		if i%2 != 0 {
      media.Description_page = "page" + strconv.Itoa(i)
			media.Data = text[indexText]
			db.InsertColl(coll, media)
      indexText++
		} else {
      media2.Description_page = "page" + strconv.Itoa(i)
			media2.Data = path[indexImg]
			db.InsertColl(coll, media2)
      indexImg++
		}
	}
}
