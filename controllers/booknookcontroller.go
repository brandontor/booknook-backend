package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/brandontor/booknookbackendgo/configs"
	"github.com/brandontor/booknookbackendgo/models"
	"github.com/brandontor/booknookbackendgo/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "booknooks")
var validate = validator.New()

func CreateBookNook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var booknook models.BookNook
	defer cancel()

	if err := c.BodyParser(&booknook); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.BookNookResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&booknook); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.BookNookResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newBookNook := models.BookNook{
		BookNook_ID:  primitive.NewObjectID(),
		BookNookName: booknook.BookNookName,
		BookTitle:    booknook.BookTitle,
		Description:  booknook.Description,
		Schedule:     booknook.Schedule,
		Genre:        booknook.Genre,
	}

	result, err := userCollection.InsertOne(ctx, newBookNook)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BookNookResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.BookNookResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetBookNookList(c *fiber.Ctx) error {
	fmt.Println("GetBookNookList")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.BookNook
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BookNookResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.BookNook
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.BookNookResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		users = append(users, singleUser)
	}

	return c.Status(http.StatusOK).JSON(
		responses.BookNookResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": users}},
	)

}

func GetBookNook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	bookNookId := c.Params("bookNookId")
	var booknook models.BookNook
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(bookNookId)

	err := userCollection.FindOne(ctx, bson.M{"booknook_id": objId}).Decode(&booknook)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BookNookResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.BookNookResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": booknook}})
}
