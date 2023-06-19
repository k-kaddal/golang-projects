package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/k-kaddal/golang-projects/go-restaurant/database"
	"github.com/k-kaddal/golang-projects/go-restaurant/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		recordPerPage,err := strvconv.Atoi(c.Query("recordPerPage"))

		if err!=nil || recordPerPage<1{
			recordPerPage = 10
		}
		page, err := strcov.Atoi(c.Query("page"))
		
		if err != nil || page<1{
			page = 1
		}

		startIndex := (page-1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match"}, bson.D{{}}}
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum, 1"}}},{"data",bson.D{{"$push", "$$ROOT"}}} }}}
		projectStage := bson.D{
			{
				"$project", bson.D{
					{"_id",0},
					{"total_count", 1},
					{"food_items", bson.D{{"slice", []interface{}{"$data", startIndex, recordPerPage}}}},
				}
			}
		}

		result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage})
		defer cancel()

		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error:", "error occured while listing food items"})
		}

		var allFoods []bson.M
		if err = result.All(ctx, &allFoods); err != nil{
			log.Fatal(err)
		}
		
		c.JSON(http.StatusOK, allFoods[0])
	}
}

func GetFood() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while fetching food"})
		}

		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func round(num float64)int{}

func toFixed(num float64, precision int)float64{}

func UpdateFood() gin.HandlerFunc{
	return func(c *gin.Context){}
}