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
