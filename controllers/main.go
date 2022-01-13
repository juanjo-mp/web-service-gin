package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"web-service-gin/config"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var paintingCollection *mongo.Collection = config.GetCollection(config.DB, "paintings")

func AddPainting(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	var painting models.Painting

	if err := c.BindJSON(&painting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(painting)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	painting.ID = primitive.NewObjectID()

	result, insertErr := paintingCollection.InsertOne(ctx, painting)
	if insertErr != nil {
		msg := fmt.Sprintf("painting was not added")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}

	defer cancel()
	fmt.Println(painting)
	c.JSON(http.StatusOK, result)
}

func GetPaintings(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	var painting []bson.M
	cursor, err := paintingCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &painting); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()
	fmt.Println(painting)
	c.JSON(http.StatusOK, painting)
}

func GetPaintingsByArtist(c *gin.Context) {
	artist := c.Params.ByName("artist")
	var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	var paintings []bson.M

	cursor, err := paintingCollection.Find(ctx, bson.M{"artist": artist})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err := cursor.All(ctx, &paintings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(paintings)
	c.JSON(http.StatusOK, paintings)
}

func GetPaintingByID(c *gin.Context) {
	paintingID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(paintingID)

	var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	var painting bson.M

	if err := paintingCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&painting); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(painting)
	c.JSON(http.StatusOK, painting)
}

func UpdateArtist(c *gin.Context) {
	paintingID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(paintingID)

	var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)

	type Artist struct {
		Name *string `json:"name"`
	}

	var artist Artist

	if err := c.BindJSON(&artist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := paintingCollection.UpdateOne(ctx, bson.M{"_id": docID},
		bson.D{
			{"$set", bson.D{{"artist", artist.Name}}},
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()
	fmt.Println(result.ModifiedCount)
	c.JSON(http.StatusOK, result.ModifiedCount)
}

func UpdatePainting(c *gin.Context) {
	paintingID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(paintingID)

	var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)

	var painting models.Painting
	if err := c.BindJSON(&painting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(painting)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	result, err := paintingCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"title":  painting.Title,
			"artist": painting.Artist,
			"image":  painting.Image,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()
	fmt.Println(result.ModifiedCount)
	c.JSON(http.StatusOK, result.ModifiedCount)
}

func DeleteOrder(c *gin.Context) {
	paintingID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(paintingID)

	var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)

	result, err := paintingCollection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

		defer cancel()
		c.JSON(http.StatusOK, result.DeletedCount)
	}
}
