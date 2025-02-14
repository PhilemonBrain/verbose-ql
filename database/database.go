package database

import (
	"context"
	"fmt"

	"github.com/PhilemonBrain/go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
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

	var jobListing model.JobListing

	return &jobListing
}

func (db *Db) GetAllJobs() []*model.JobListing {

	var jobListings []*model.JobListing

	return jobListings
}

func (db *Db) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {

	var jobListing model.JobListing
	return &jobListing
}

func (db *Db) UpdateJobListing(id string, jobInfo model.UpdateJobListingInput) *model.JobListing {

	var jobListing model.JobListing
	return &jobListing
}

func (db *Db) DeleteJobListing(jobId string) *model.DeleteJobResponse {

	var deletedJobResponse model.DeleteJobResponse
	return &deletedJobResponse
}
