package dao

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
)

var instance *elastic.Client

var ctx = context.Background()

func init() {
	host := os.Getenv("ES_HOST")
	fmt.Println(host)
	if host == "" {
		host = "http://localhost:9200"
	}

	client, err := elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)

	if err != nil {
		log.Fatal(err)
	}

	info, code, err := client.Ping(host).Do(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	instance = client
}

// GetESClient get an instance of elasticsearch client
func GetESClient() *elastic.Client {
	return instance
}
