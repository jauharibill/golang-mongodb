package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/treemux"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-mongodb/models"
	"golang-mongodb/presenter"
	"golang-mongodb/utilities"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type Controller struct {
	Mongodb *mongo.Client
	repo    *mongo.Collection
}

func InitController(mong *mongo.Client) Controller {
	collection := mong.Database("mongodb-golang").Collection("users")
	return Controller{
		Mongodb: mong,
		repo:    collection,
	}
}

// SHOW DATA
func (_r *Controller) Show(writer http.ResponseWriter, request treemux.Request) error {
	var response presenter.Response
	var user models.UserModel

	filter := bson.M{
		"id": utilities.StrToInt(request.Param("id")),
	}

	collection := _r.repo.FindOne(context.Background(), filter)

	err := collection.Decode(&user)

	if err != nil {
		log.Println(err.Error())
	}

	response.Data = user
	response.Message = "Success Show Data"

	return treemux.JSON(writer, response)
}

// STORE DATA
func (_r *Controller) Store(writer http.ResponseWriter, request treemux.Request) error {
	var user models.UserModel
	var response presenter.Response

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		response.Message = err.Error()
		return treemux.JSON(writer, response)
	}

	insert, err := _r.repo.InsertOne(context.Background(), user)

	if err != nil {
		response.Message = err.Error()
		return treemux.JSON(writer, response)
	}

	response.Data = user
	response.Message = fmt.Sprintf("Success Storing Data %s", insert.InsertedID)

	writer.WriteHeader(http.StatusCreated)
	return treemux.JSON(writer, response)
}

// UPDATE DATA
func (_r *Controller) Update(writer http.ResponseWriter, request treemux.Request) error {
	var user models.UserModel
	var response presenter.Response

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		response.Message = err.Error()
		treemux.JSON(writer, response)
	}

	filter := bson.M{
		"id": utilities.StrToInt(request.Param("id")),
	}

	_r.repo.UpdateOne(context.Background(), filter, user)

	response.Message = "Success update data"
	response.Data = nil

	return treemux.JSON(writer, response)
}

// DELETE DATA
func (_r *Controller) Delete(writer http.ResponseWriter, request treemux.Request) error {
	var response presenter.Response

	filter := bson.M{
		"id": utilities.StrToInt(request.Param("id")),
	}

	_, err := _r.repo.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err.Error())
	}

	response.Message = "Success Delete Data"
	response.Data = nil

	return treemux.JSON(writer, response)
}
