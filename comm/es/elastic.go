package es

import (
	"comm/tool"
	"fmt"
	"github.com/olivere/elastic"
)

func GetESClient() (*elastic.Client, error) {
	es := tool.GetEnvDefault("ELASTICSEARCH", "127.0.0.1:9200")
	client, err := elastic.NewClient(elastic.SetURL(es),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	fmt.Println("ES initialized...")
	return client, err

}
