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
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

var Indexes = make(map[string][]mongo.IndexModel)

type Database struct {
	URL     string        `yaml:"url" json:"url"`
	Name    string        `yaml:"name" json:"name"`
	Timeout time.Duration `yaml:"timeout" json:"timeout"`
}

// CreateIndex create index for collection
func CreateIndex(db string, indexes ...mongo.IndexModel) {
	if im, ok := Indexes[db]; ok {
		Indexes[db] = append(im, indexes...)
		return
	}
	Indexes[db] = indexes
}

func (e *Database) Init() {
	if e.URL == "" {
		e.URL = "mongodb://localhost:27017"
	}
	ctx, cancel := context.WithTimeout(
		context.Background(),
		e.Timeout*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(e.URL))
	if err != nil {
		log.Fatalln(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}
	DB = client.Database(e.Name)
	var indexes []string
	for table := range Indexes {
		indexes, err = DB.Collection(table).Indexes().CreateMany(ctx, Indexes[table])
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%s create index %s\n", table, strings.Join(indexes, ""))
	}
}
