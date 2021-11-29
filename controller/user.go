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
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id" form:"_id"`
	First_name   string             `json:"first_name" form:"first_name"`
	Last_name    string             `json:"last_name" form:"last_name"`
	Phone_number int64              `json:"phone_no" form:"phone_no"`
	Email        string             `json:"email" form:"email"`
	Address      string             `json:"address" form:"address"`
	City         string             `json:"city" form:"city"`
	State        string             `json:"state" form:"state"`
	Password     string             `json:"password" form:"password"`
}

func GetUsers(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, users)
	var users []*user
	client, ctx, cancel := model.GetConnection()
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
	if err := c.Bind(&newUser); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to create hashed password for the provided password"})
	}
	newUser.Password = string(hashedPassword)

	client, ctx, cancel := model.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	newUser.Id = primitive.NewObjectID()
	result, err := client.Database("MyProject").Collection("user").InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"message": "User Succesfully registered",
		"id":      result.InsertedID.(primitive.ObjectID),
	})
	//c.JSON(http.StatusOK, gin.H{"id": result.InsertedID.(primitive.ObjectID)})
}

func GetUserByEmailAndPassword(c *gin.Context) {
	var getUser *user
	email := c.PostForm("email")
	password := c.PostForm("password")

	client, ctx, cancel := model.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	result := client.Database("MyProject").Collection("user").FindOne(ctx, bson.M{"email": email})
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not Found"})
		return
	}
	err := result.Decode(&getUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}
	error := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(password))
	if error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Username or password is incorrect!"})
		return
	}
	c.Set("id", getUser.Id.Hex())
	// c.JSON(http.StatusOK, getUser)
	// log.Print(result)
	log.Print(err)
	c.HTML(http.StatusOK, "user/profile", gin.H{
		"message": "User Succesfully logged",
		"user":    getUser.First_name,
		"id":      c.MustGet("id"),
	})
	//c.JSON(http.StatusOK, getUser)
}

func UpdateUser(c *gin.Context) {
	var updateUser *user

	if err := c.Bind(&updateUser); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	client, ctx, cancel := model.GetConnection()
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
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	fmt.Println()
	client, ctx, cancel := model.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	result, err := client.Database("MyProject").Collection("user").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	// fmt.Println(result)
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
	client, ctx, cancel := model.GetConnection()
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

func ViewLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login.tmpl", gin.H{
		"title": "Main website",
	})
}

func NewUser(c *gin.Context) {
	c.HTML(http.StatusOK, "user/new.tmpl", gin.H{
		"title": "Main website",
	})
}

func ViewProfile(c *gin.Context) {
	c.HTML(http.StatusOK, "user/profile.tmpl", gin.H{
		"title": "Main website",
	})
}
