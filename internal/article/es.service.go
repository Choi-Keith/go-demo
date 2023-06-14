package article

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

var ESClient *elastic.Client

func init() {
	var err error
	ESClient, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false))
	if err != nil {
		// 如果使用远程机器或者使用docker启的elasticsearch, 会报错：no elasticsearch node available
		// 解决的方法是设置elastic.SetSniff(false) 或者修改配置
		// cat elasticsearch.yaml

		// cluster.name: "docker-cluster"
		// network.host: 0.0.0.0
		// network.publish_host:"127.0.0.1"
		fmt.Printf("connect failed: %v", err)
		return
	}
}

func (a *service) Insert() {
	exists, err := ESClient.IndexExists("article").Do(context.Background())
	if err != nil {
		fmt.Printf("exists: %v", err)
	}
	if !exists {
		mappings := `
{
    "settings": {
        "number_of_shards":1,
		"number_of_replicas":0
    },
    "mappings": {
            "properties": {
                "id": {
                    "type": "keyword"
                },
				"title": {
					"type": "keyword"
				},
				"description": {
					"type": "text"
				},
				"content": {
					"type": "text"
				},
                "creatorID": {
                    "type": "keyword"
                },
                "createdAt": {
                    "type": "long"
                },
                "updatedAt": {
                    "type": "long"
                }
            }
    }
}`

		fmt.Println("no Acknowledged")
		createIndex, err := ESClient.CreateIndex("article").Body(mappings).Do(context.Background())
		if err != nil {
			// Handle error
			fmt.Printf("create index faild: %v", err)
			return
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			fmt.Println("no Acknowledged222222222222")
		}
	}
	article, err := Service.Get("1668790802597285888")
	if err != nil {
		fmt.Printf("get article failed: %v", article)
		return
	}

	put1, err := ESClient.Index().
		Index("article").
		Id("1").
		BodyJson(article).
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Printf("insert failed: %v", err)
		return
	}

	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

func (a *service) DeleteIndex() {
	index := ESClient.DeleteIndex("article")
	fmt.Printf("delete index: %v", index)
}
