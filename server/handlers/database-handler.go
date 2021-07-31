package handlers

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/LeNgocPhuc99/private-chat/server/database"
	"github.com/LeNgocPhuc99/private-chat/server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user collection

func GetUserByUsername(username string) (User, error) {
	var user User

	// get collection
	collection := database.DBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// find user
	queryErr := collection.FindOne(ctx, bson.M{
		"username": username,
	}).Decode(&user)

	return user, queryErr
}

func GetUserByUserID(userID string) (User, error) {
	var user User

	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return User{}, err
	}

	collection := database.DBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	queryErr := collection.FindOne(ctx, bson.M{
		"_id": docID,
	}).Decode(&user)

	return user, queryErr
}

func UpdateUserStatus(userID string, status string) error {
	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	// get collection
	collection := database.DBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, queryErr := collection.UpdateOne(ctx, bson.M{"_id": docID}, bson.M{"$set": bson.M{"online": status}})
	if queryErr != nil {
		return errors.New("update status err:" + queryErr.Error())
	}

	return nil
}

func LoginQuery(requestPayload UserLoginRequest) (UserResponse, error) {
	user, err := GetUserByUsername(requestPayload.Username)
	if err != nil {
		return UserResponse{}, errors.New(UserIsNotRegisteredWithUs)
	}

	if err := utils.CommparePassword(requestPayload.Password, user.Password); err != nil {
		return UserResponse{}, errors.New(LoginPasswordIsInCorrect)
	}

	if err := UpdateUserStatus(user.ID, "Y"); err != nil {
		return UserResponse{}, errors.New(UpdateStatusFail)
	}

	return UserResponse{
		UserID:   user.ID,
		Username: user.Username,
	}, nil
}

// return UserID
func RegisterQuery(requestPayload UserRegistrationRequest) (string, error) {
	// hash password
	passwordHash, err := utils.CreatePassword(requestPayload.Password)
	if err != nil {
		return "", errors.New(ServerFailedResponse)
	}

	// get collection
	collection := database.DBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	queryResponse, queryErr := collection.InsertOne(ctx, bson.M{
		"username": requestPayload.Username,
		"password": passwordHash,
		"online":   "Y",
	})

	if queryErr != nil {
		return "", errors.New("insert err:" + queryErr.Error())
	}

	objectID, queryObjectIDErr := queryResponse.InsertedID.(primitive.ObjectID)

	if !queryObjectIDErr {
		return "", errors.New("query Object's ID error")
	}

	return objectID.Hex(), nil
}

// except userID
func GetAllOnlineUser(userID string) ([]UserResponse, error) {
	var onlineUser []UserResponse

	// get document ID
	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	// get collection
	collection := database.DBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// query database
	cursor, queryErr := collection.Find(ctx, bson.M{
		"online": "Y",
		"_id": bson.M{
			"$ne": docID,
		},
	})

	if queryErr != nil {
		return nil, queryErr
	}

	// get document
	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		if err == nil {
			onlineUser = append(onlineUser, UserResponse{
				Username: user.Username,
				UserID:   user.ID,
				Online:   user.Online,
			})
		}
	}

	return onlineUser, nil
}

// message collection

func StoreMessage(messagePayload MessagePayload) bool {

	collection := database.DBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, queryErr := collection.InsertOne(ctx, bson.M{
		"message":    messagePayload.Message,
		"toUserID":   messagePayload.ToUserID,
		"fromUserID": messagePayload.FromUserID,
	})

	return queryErr == nil
}

func GetConversationBetweenTwoUsers(fromUserID, toUserID string) ([]Conversation, error) {
	var conversations []Conversation

	collection := database.DBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	queryCondition := bson.M{

		"$or": []bson.M{
			{
				"$and": []bson.M{
					{"fromUserID": fromUserID},
					{"toUserID": toUserID},
				},
			},
			{
				"$and": []bson.M{
					{"fromUserID": toUserID},
					{"toUserID": fromUserID},
				},
			},
		},
	}

	cursor, queryError := collection.Find(ctx, queryCondition)

	if queryError != nil {
		return nil, queryError
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var conversation Conversation
		err := cursor.Decode(&conversation)

		if err == nil {
			conversations = append(conversations, Conversation{
				ID:         conversation.ID,
				FromUserID: conversation.FromUserID,
				ToUserID:   conversation.ToUserID,
				Message:    conversation.Message,
			})
		}
	}
	return conversations, nil
}
