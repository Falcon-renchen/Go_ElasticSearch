package AppInit

import (
	"github.com/olivere/elastic/v7"
)

func GetEsClient() *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetURL("http://172.16.17.157:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil
	}

	return client
}
