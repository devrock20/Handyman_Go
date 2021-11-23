package entities

import (
	"fmt"
	"log"
	"net/http"
	"project/connections"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type user struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id"`
	First_name   string             `json:"first_name"`
	Last_name    string             `json:"last_name"`
	Phone_number int64              `json:"phone_no"`
	Email        string             `json:"email"`
	Address      string             `json:"address"`
	City         string             `json:"city"`
	State        string             `json:"state"`
	Password     string             `json:"password"`
}

func GetUsers(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, users)
	var users []*user
	client, ctx, cancel := connections.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	cursor, err := client.Database("MyProject").Collection("user").Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &users)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, users)
}

func AddUser(c *gin.Context) {
	var newUser *user

	if err := c.BindJSON(&newUser); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	client, ctx, cancel := connections.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	newUser.Id = primitive.NewObjectID()
	result, err := client.Database("MyProject").Collection("user").InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": result.InsertedID.(primitive.ObjectID)})
}

func GetUserByEmailAndPassword(c *gin.Context) {
	var getUser *user
	email := c.Param("email")
	password := c.Param("password")
	client, ctx, cancel := connections.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	result := client.Database("MyProject").Collection("user").FindOne(ctx, bson.M{"email": email, "password": password})
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not Found"})
		return
	}
	err := result.Decode(&getUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}
	c.JSON(http.StatusOK, getUser)
}

func UpdateUser(c *gin.Context) {
	var updateUser *user

	if err := c.BindJSON(&updateUser); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	client, ctx, cancel := connections.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"$set": updateUser,
	}

	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	err := client.Database("MyProject").Collection("user").FindOneAndUpdate(ctx, bson.M{"_id": updateUser.Id}, update, &opt).Decode(&updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, updateUser)
}

func DeleteUser(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	fmt.Println()
	client, ctx, cancel := connections.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	result, err := client.Database("MyProject").Collection("user").DeleteOne(ctx, bson.M{"_id": id})
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

func GetUserbyId(c *gin.Context) {
	var getUser *user
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	fmt.Println()
	client, ctx, cancel := connections.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	result := client.Database("MyProject").Collection("user").FindOne(ctx, bson.M{"_id": id})
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not Found"})
		return
	}
	error := result.Decode(&getUser)
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}
	c.JSON(http.StatusOK, getUser)
}
