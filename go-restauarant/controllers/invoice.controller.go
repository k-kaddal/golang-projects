package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/k-kaddal/golang-projects/go-restaurant/database"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type InvoiceViewFormat struct{
	Invoice_Id				string
	Payment_method			string
	Order_Id				string
	Payment_status			*string
	Payment_due				interface{}
	Table_number			interface{}
	Payment_due_date		time.Time
	Order_details			interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listing invoices"})
		}

		var allInvoices []bson.M
		if err = result.All(ctx, &allInvoices); err != nil{
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allInvoices)
	}
}
func GetInvoice() gin.HandlerFunc{
	return func(c *gin.Context){}
}
func CreateInvoice() gin.HandlerFunc{
	return func(c *gin.Context){}
}
func UpdateInvoice() gin.HandlerFunc{
	return func(c *gin.Context){}
}