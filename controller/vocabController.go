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
	"go.mongodb.org/mongo-driver/bson"
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

func IndexVocab(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	var vocab []model.Vocabularies
	defer cancel()

	result, err := vocabCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseBuilder{Code: http.StatusInternalServerError, Success: false, Message: "Something wrong", Data: err.Error()})
	}

	defer result.Close(ctx)
	for result.Next(ctx) {
		var itemVocab model.Vocabularies
		if err := result.Decode(&itemVocab); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(response.ResponseBuilder{Code: http.StatusInternalServerError, Success: false, Message: "Something wrong", Data: err.Error()})
		}

		vocab = append(vocab, itemVocab)
	}

	return c.Status(http.StatusOK).JSON(response.ResponseBuilder{
		Code:    http.StatusOK,
		Success: true,
		Message: "Ok",
		Data:    vocab,
	})
}

func Translate(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	word := c.Params("word")
	var vocab model.Vocabularies
	defer cancel()

	err := vocabCollection.FindOne(ctx, bson.M{"en_word": word}).Decode(&vocab)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseBuilder{Code: http.StatusInternalServerError, Success: false, Message: "Something wrong", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(response.ResponseBuilder{
		Code:    http.StatusOK,
		Success: true,
		Message: "Ok",
		Data:    vocab,
	})
}

func UpdateVocab(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	var vocab model.Vocabularies
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	// validate request body
	if err := c.BodyParser(&vocab); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseBuilder{Code: http.StatusBadRequest, Success: false, Message: "Please Check your data input", Data: err.Error()})
	}

	// use validator library to validate required fields
	if validationErr := validate.Struct(&vocab); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseBuilder{Code: http.StatusBadRequest, Success: false, Message: "Please Check your data input", Data: validationErr.Error()})
	}

	tmp := bson.M{"vocab_type": vocab.Vocab_Type, "en_word": vocab.En_Word, "id_word": vocab.Id_Word}

	result, err := vocabCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": tmp})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseBuilder{Code: http.StatusInternalServerError, Success: false, Message: "Something wrong", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(response.ResponseBuilder{Code: http.StatusOK, Success: true, Message: "Ok", Data: result.MatchedCount})
}
