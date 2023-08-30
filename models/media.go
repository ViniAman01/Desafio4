package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Media struct {
	UUID             primitive.ObjectID  `bson:"_id,omitempty"`
	Data_type        string              `bson:"data_type"`
	Code_book        string              `bson:"code_book"`
	Description_page string              `bson:"description_page"`
	Data             string              `bson:"data"`
	Metadata         string              `bson:"metadata,omitempty"`
	Datetime         primitive.Timestamp `bson:"datetime,omitempty"`
}

var DescriptionsPage = []string{
	"cover_page",
	"description_page",
	"page1",
	"page2",
	"page3",
	"page4",
	"page5",
	"page6",
	"page7",
	"page8",
	"page9",
	"page10",
	"page11",
	"page12",
	"page13",
	"page14",
	"biography_page",
}
