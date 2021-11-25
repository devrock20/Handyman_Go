package controller

import (
	"fmt"
	"log"
	"net/http"
	"project/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Worker struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id" form:"_id"`
	First_name   string             `json:"first_name" form:"first_name"`
	Last_name    string             `json:"last_name" form:"last_name"`
	Phone_number int64              `json:"phone_no" form:"phone_no"`
	Email        string             `json:"email" form:"email"`
	WorkType     string             `json:"workType" form:"workType"`
	City         string             `json:"city" form:"city"`
	State        string             `json:"state" form:"state"`
	Password     string             `json:"password" form:"password"`
}

func GetAllWorkers(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, users)
	var workers []*Worker
	client, ctx, cancel := model.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	cursor, err := client.Database("MyProject").Collection("handyman").Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &workers)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, workers)
}

func AddWorker(c *gin.Context) {
	var newWorker *Worker
	if err := c.Bind(&newWorker); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	// newWorker = c.MustGet("newWorker").(Worker)
	// hashedPassoword, error := bcrypt.GenerateFromPassword([]byte(newWorker.Password), bcrypt.DefaultCost)
	// if error != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to create hashed password for the provided password"})
	// }

	// hasherr := bcrypt.CompareHashAndPassword([]byte(hashedPassoword), []byte(newWorker.Password))
	// if hasherr != nil {
	// 	c.AbortWithStatusJSON(400, gin.H{"error": "Username or password is incorrect!"})
	// }
	// fmt.Println("hash", hasherr)

	// newWorker.Password = string(hashedPassoword)

	client, ctx, cancel := model.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	newWorker.Id = primitive.NewObjectID()
	result, err := client.Database("MyProject").Collection("handyman").InsertOne(ctx, newWorker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	// c.HTML(http.StatusOK, "login.tmpl", gin.H{
	// 	"message": "User Succesfully registered",
	// 	"id":      result.InsertedID.(primitive.ObjectID),
	// })
	c.JSON(http.StatusOK, gin.H{"id": result.InsertedID.(primitive.ObjectID)})
}

func GetWorkerById(c *gin.Context) {
	var getWorker *Worker
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	// fmt.Println()
	client, ctx, cancel := model.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	result := client.Database("MyProject").Collection("handyman").FindOne(ctx, bson.M{"_id": id})
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not Found"})
		return
	}
	error := result.Decode(&getWorker)
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}
	c.JSON(http.StatusOK, getWorker)
}

func UpdateWorker(c *gin.Context) {
	var updateWorker *Worker
	// c.Request.URL.Query()
	if err := c.Bind(&updateWorker); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	client, ctx, cancel := model.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"$set": updateWorker,
	}

	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	err := client.Database("MyProject").Collection("handyman").FindOneAndUpdate(ctx, bson.M{"_id": updateWorker.Id}, update, &opt).Decode(&updateWorker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, updateWorker)
}

func DeleteWorker(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	fmt.Println()
	client, ctx, cancel := model.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	result, err := client.Database("MyProject").Collection("handyman").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	fmt.Println(result)
	if result == nil {
		c.JSON(http.StatusNoContent, gin.H{"msg": "User not Deleted"})
		return
	}
	c.JSON(http.StatusNoContent, result)
}

func WorkerLogin(c *gin.Context) {
	var getWorker *Worker
	email := c.PostForm("email")
	password := c.PostForm("password")
	client, ctx, cancel := model.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	result := client.Database("MyProject").Collection("handyman").FindOne(ctx, bson.M{"email": email, "password": password})
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not Found"})
		return
	}
	err := result.Decode(&getWorker)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}
	c.JSON(http.StatusOK, getWorker)

}

func GetWorkerByEmailAndPassword(c *gin.Context) {
	var getWorker *Worker
	// email := c.Param("email")
	// password := c.Param("password")
	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println("email", email)
	fmt.Println("password", password)

	client, ctx, cancel := model.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	result := client.Database("MyProject").Collection("handyman").FindOne(ctx, bson.M{"email": email, "password": password})
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not Found"})
		return
	}
	err := result.Decode(&getWorker)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}
	c.JSON(http.StatusOK, getWorker)
}

func ViewWorkerLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "workerLogin.tmpl", gin.H{
		"title": "Main website",
	})
}
