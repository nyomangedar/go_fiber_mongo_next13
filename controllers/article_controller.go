package controllers

import (
	"context"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var articleCollection *mongo.Collection = configs.GetCollection(configs.DB, "articles")
var validate = validator.New()

// Create an Article
func CreateArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var article models.Article
	defer cancel()

	//validate the req body
	if err := c.BodyParser(&article); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ArticleResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	newArticle := models.Article{
		Id:          primitive.NewObjectID(),
		Title:       article.Title,
		Author:      article.Author,
		Content:     article.Content,
		CreatedDate: time.Now(),
	}

	result, err := articleCollection.InsertOne(ctx, newArticle)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ArticleResponse{Status: http.StatusInternalServerError, Message: "Error creating new article", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.ArticleResponse{Status: http.StatusCreated, Message: "Article Succesfully created", Data: &fiber.Map{"data": result}})

}

// Get All articles
func GetArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var articles []models.Article
	defer cancel()

	results, err := articleCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ArticleResponse{Status: http.StatusInternalServerError, Message: "Error!", Data: &fiber.Map{"data": err}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleArticle models.Article
		if err = results.Decode(&singleArticle); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ArticleResponse{Status: http.StatusInternalServerError, Message: "Error validating in loop", Data: &fiber.Map{"data": err}})
		}
		articles = append(articles, singleArticle)
	}

	return c.Status(http.StatusOK).JSON(
		responses.ArticleResponse{Status: http.StatusOK, Message: "Found Articles", Data: &fiber.Map{"data": articles}},
	)

}

// Find one article
func FindArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	var article models.Article
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := articleCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&article)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ArticleResponse{Status: http.StatusInternalServerError, Message: "Error finding article", Data: &fiber.Map{"data": err}})
	}

	return c.Status(http.StatusOK).JSON(responses.ArticleResponse{Status: http.StatusOK, Message: "Found Article", Data: &fiber.Map{"data": article}})
}
