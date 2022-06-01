package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vocabularies struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Vocab_Type string             `json:"vocab_type" validate:"required"`
	En_Word    string             `json:"en_word" validate:"required"`
	Id_Word    []string           `json:"id_word" validate:"required"`
}
