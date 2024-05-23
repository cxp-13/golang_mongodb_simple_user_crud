package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang_mongoDB/models"
	"net/http"
)

type UserController struct {
	mongoDBClient *mongo.Client
}

func NewUserController(client *mongo.Client) *UserController {
	return &UserController{mongoDBClient: client}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	collection := uc.mongoDBClient.Database("youtube_test").Collection("users")
	filter := bson.M{"_id": objectId}

	var user models.User
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// CreateUser creates a new user in the MongoDB database
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	collection := uc.mongoDBClient.Database("youtube_test").Collection("users")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result.InsertedID})
}

// DeleteUser deletes a user by ID from the MongoDB database
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	collection := uc.mongoDBClient.Database("youtube_test").Collection("users")
	filter := bson.M{"_id": objectId}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}
	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": "User deleted"})
}

// UpdateUser updates an existing user in the MongoDB database
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	collection := uc.mongoDBClient.Database("youtube_test").Collection("users")
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": user}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}
	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": "User updated"})
}
