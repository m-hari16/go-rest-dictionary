package controller

import (
	"context"
	"go-rest-dictionary/config"
	"go-rest-dictionary/model"
	"go-rest-dictionary/response"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var vocabCollection *mongo.Collection = config.GetCollection(config.DB, "vocabularies")
var validate = validator.New()

func StoreVocab(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var vocab model.Vocabularies
	defer cancel()

	// validate the requesy body
	if err := c.BodyParser(&vocab); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseBuilder{Code: http.StatusBadRequest, Success: false, Message: "Please, check your data input", Data: err.Error()})
	}

	// use validation library to validation required fields
	if validationErr := validate.Struct(&vocab); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseBuilder{Code: http.StatusBadRequest, Success: false, Message: "Fill all fields", Data: validationErr.Error()})
	}

	newVocab := model.Vocabularies{
		Id:         primitive.NewObjectID(),
		Vocab_Type: vocab.Vocab_Type,
		En_Word:    vocab.En_Word,
		Id_Word:    vocab.Id_Word,
	}

	result, err := vocabCollection.InsertOne(ctx, newVocab)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseBuilder{Code: http.StatusInternalServerError, Success: false, Message: "Something wrong", Data: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(response.ResponseBuilder{Code: http.StatusCreated, Success: true, Message: "Data has been stored", Data: result})

}
