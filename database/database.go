package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PhilemonBrain/go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conectionString string = "mongodb+srv://brainphil:IFYEwcbmGYEMpTGl@mycluster.vuxlq.mongodb.net/?retryWrites=true&w=majority&appName=MyCluster"

type Db struct {
	client *mongo.Client
}

func Connect() Db {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(conectionString).SetServerAPIOptions(serverAPI)

	// send a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return Db{client}
}

func (db *Db) GetJob(id string) *model.JobListing {
	jobCollec := db.client.Database("mainDB").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobListing model.JobListing
	err := jobCollec.FindOne(ctx, filter).Decode(jobListing)
	if err != nil {
		log.Fatal(err)
	}

	return &jobListing
}

func (db *Db) GetAllJobs() []*model.JobListing {
	jobCollec := db.client.Database("mainDB").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var jobListings []*model.JobListing
	cursor, err := jobCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &jobListings); err != nil {
		panic(err)
	}

	return jobListings
}

func (db *Db) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	jobCollection := db.client.Database("mainDB").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	inserted, err := jobCollection.InsertOne(ctx, bson.M{
		"title":       jobInfo.Titile,
		"description": jobInfo.Description,
		"url":         jobInfo.URL,
		"company":     jobInfo.Company,
	})
	if err != nil {
		log.Fatal(err)
	}

	insertId := inserted.InsertedID.(primitive.ObjectID).Hex()
	returnjobListing := model.JobListing{
		ID:          insertId,
		Titile:      jobInfo.Titile,
		Description: jobInfo.Description,
		Company:     jobInfo.Company,
		URL:         *jobInfo.URL,
	}

	return &returnjobListing
}

func (db *Db) UpdateJobListing(id string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	jobCollection := db.client.Database("mainDB").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobInfo := bson.M{}

	if jobInfo.Company != nil {
		updateJobInfo["company"] = jobInfo.Company
	}
	if jobInfo.URL != nil {
		updateJobInfo["url"] = jobInfo.URL
	}
	if jobInfo.Description != nil {
		updateJobInfo["description"] = jobInfo.Description
	}
	if jobInfo.Titile != nil {
		updateJobInfo["title"] = jobInfo.Titile
	}

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := jobCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate())

	var jobListing model.JobListing
	if err := results.Decode(jobListing); err != nil {
		log.Fatal(err)
	}

	return &jobListing
}

func (db *Db) DeleteJobListing(jobId string) *model.DeleteJobResponse {

	// jobCollection := db.client.Database("mainDB").Collection("jobs")
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	var deletedJobResponse model.DeleteJobResponse
	return &deletedJobResponse
}
