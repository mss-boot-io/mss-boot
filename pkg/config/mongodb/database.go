/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 13:50
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 13:50
 */

package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

var tables = make([]Tabler, 0)

type Database struct {
	URL     string        `yaml:"url" json:"url"`
	Name    string        `yaml:"name" json:"name"`
	Timeout time.Duration `yaml:"timeout" json:"timeout"`
}

// AppendTable append table
func AppendTable(t Tabler) {
	tables = append(tables, t)
}

func (e *Database) Init() {
	if e.URL == "" {
		e.URL = "mongodb://localhost:27017"
	}
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(e.URL).
		SetServerAPIOptions(serverAPIOptions)
	if e.Timeout == 0 {
		//set default timeout
		e.Timeout = 10 * time.Second
	}
	ctx, cancel := context.WithTimeout(
		context.TODO(),
		e.Timeout)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalln(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}
	DB = client.Database(e.Name)
	for i := range tables {
		tables[i].Make()
	}
}

// C get table's Collection
func (e *Database) C(t Tabler) *mongo.Collection {
	return DB.Collection(t.TableName())
}
