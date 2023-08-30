package main

import (
	db "project/database"
	"project/models"
)

var media = models.Media{
	Data_type:        "text",
	Code_book:        "12345",
	Description_page: "page1",
	Data:             `Era uma vez uma pequena cidade chamada Serenidade, onde as pessoas viviam suas vidas tranquilas. No centro da cidade, havia uma loja de presentes chamada "Surpresas Mágicas", conhecida por suas ofertas especiais. Em uma manhã ensolarada, Maria, uma jovem moradora de Serenidade, entrou na loja em busca de um presente de aniversário para sua avó. Ela examinou os prateleiras, mas nada parecia especial o suficiente para a pessoa que tanto amava. O proprietário da loja, Sr. Higgins, percebeu a hesitação de Maria e se aproximou com um sorriso amigável. Ele sugeriu que ela desse à avó um livro de histórias antigas, cheio de memórias e aventuras compartilhadas. Maria seguiu o conselho de Sr. Higgins e escolheu um livro de histórias. Ela escreveu uma mensagem amorosa na primeira página e entregou o presente à sua avó no dia de seu aniversário. À noite, as duas se aconchegaram na sala de estar, lendo as histórias juntas. A avó de Maria sorriu e disse que era o presente mais especial que já recebera, pois continha as lembranças de sua família. E assim, a história de Maria e sua avó se tornou parte das histórias antigas do livro, um testemunho do amor e da conexão entre as gerações.`,
}

var media2 = models.Media{
	Data_type:        "image",
	Code_book:        "12345",
	Description_page: "page2",
	Data:             "./files/lendo_livro.jpg",
}

var media3 = models.Media{
	Data_type:        "text",
	Code_book:        "12345",
	Description_page: "page3",
	Data:             `Serenidade continuava a ser um refúgio de tranquilidade, envolto em suas paisagens pitorescas e no ritmo sereno de suas ruas calmas. A loja "Surpresas Mágicas" permanecia no coração da cidade, com suas janelas convidativas que exibiam uma miríade de tesouros encantados. Cada objeto na loja contava uma história única, uma história que esperava ser desvendada por aqueles que tinham olhos curiosos. As pessoas vinham de todas as partes para explorar suas prateleiras, em busca daquela surpresa que poderia iluminar os olhos de um ente querido. Uma manhã ensolarada trouxe consigo uma brisa suave que acariciava as folhas das árvores e as flores em flor. Era nesse cenário que Clara, uma jovem artista com um coração cheio de sonhos, decidiu visitar a loja. Ela estava à procura de algo especial para seu irmão mais novo, Mateus, que estava prestes a completar dez anos. Mateus sempre foi fascinado por histórias de aventuras e mundos mágicos. Clara se aventurou na loja, seus olhos dançando entre os objetos curiosos. Foi então que ela viu: um globo de neve com uma réplica encantadora de um navio pirata em seu interior, capturando a essência de um conto de fadas.`,
}

var media4 = models.Media{
	Data_type:        "image",
	Code_book:        "12345",
	Description_page: "page4",
	Data:             "./files/car.jpg",
}

var media5 = models.Media{
	Data_type:        "text",
	Code_book:        "12345",
	Description_page: "page5",
	Data:             "Sr. Higgins, o amável proprietário da loja, notou o brilho nos olhos de Clara e se aproximou. Com um aceno acolhedor, ele compartilhou a história por trás daquele globo de neve. Era uma história sobre coragem, descobertas e a jornada de um jovem pirata em busca de um tesouro perdido. Clara soube imediatamente que era o presente perfeito para Mateus, uma lembrança tangível de que a imaginação poderia ganhar vida. Ela agradeceu a Sr. Higgins e saiu da loja, sabendo que aquele presente traria alegria ao coração do irmão no dia de seu aniversário.",
}

var media6 = models.Media{
	Data_type:        "image",
	Code_book:        "12345",
	Description_page: "page6",
	Data:             "./files/3b1b.jpeg",
}

var media7 = models.Media{
	Data_type:        "text",
	Code_book:        "12345",
	Description_page: "page7",
	Data:             "O dia do aniversário de Mateus finalmente chegou, trazendo consigo risos contagiantes e a promessa de uma celebração repleta de felicidade. Clara entregou o presente a Mateus, cujos olhos se arregalaram de admiração ao ver o globo de neve. Ele girou o objeto entre suas mãos pequenas e logo se viu imerso na história que Sr. Higgins compartilhara.",
}

var medias = []models.Media{media, media2, media3, media4, media5, media6, media7}

func main() {
	client := db.ConnectDB()

	coll := db.GetCollection("mainDB2", "collteste", client)

	for _, data := range medias {
		db.InsertColl(coll, data)
	}
}
