package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/elastic/go-elasticsearch"
)

type ElasticsearchHelper struct {
	host     string
	conn     *elasticsearch.Client
	username string
	password string
}

func GetElasticsearchInstance(host string, username string, password string) ElasticsearchHelper {
	conn, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://" + host},
		Username:  username,
		Password:  password,
	})

	if err != nil {
		fmt.Println(err)
	}

	instance := ElasticsearchHelper{conn: conn, host: host, username: username, password: password}
	return instance
}

func (this *ElasticsearchHelper) Push(index string, obj interface{}) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
	}

	//	fmt.Println("es push:", string(jsonBytes))
	this.PushJson(index, jsonBytes)
}

func (this *ElasticsearchHelper) PushJson(index string, jsonBytes []byte) {
	res, err := this.conn.Index(
		index,                                // Index name
		strings.NewReader(string(jsonBytes)), // Document body
		this.conn.Index.WithRefresh("true"),  // Refresh
	)

	if err != nil {
		fmt.Println(err)
	} else {
		defer res.Body.Close()
	}
}

//获取索引列表
func (this *ElasticsearchHelper) GetIndexList() (indexMaps []map[string]string) {
	url := "http://" + this.username + ":" + this.password + "@" + this.host + "/_cat/indices?format=json&index=*"

	networkUtil := NetworkUtil{}
	var requestBody io.Reader
	_, responseBody, err, _ := networkUtil.HttpRequest("GET", url, map[string]string{}, requestBody)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(responseBody, &indexMaps)
	return indexMaps
}

type RecordCountResponse struct {
	Count int `json:"count"`
}

func (this *ElasticsearchHelper) GetCount(index string, keyword string) int {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"match_all": map[string]interface{}{}, //搜索全部
	}

	if keyword != "" {
		arr := strings.Split(keyword, ":")
		if len(arr) > 1 {
			query = map[string]interface{}{
				"match": map[string]interface{}{ //搜索关键字
					strings.TrimSpace(arr[0]): strings.TrimSpace(arr[1]), //按字段查询
				},
			}
		}
	}

	queryData := map[string]interface{}{
		"query": query,
	}

	recordCountResponse := RecordCountResponse{}

	if err := json.NewEncoder(&buf).Encode(queryData); err != nil {
		panic(err)
	}

	resp, err := this.conn.Count(
		this.conn.Count.WithContext(context.Background()),
		this.conn.Count.WithIndex(index),
		this.conn.Count.WithBody(&buf),
	)
	if err != nil {
		panic(err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&recordCountResponse); err != nil {
		panic(err)
	}

	return recordCountResponse.Count
}

func (this *ElasticsearchHelper) Search(index string, keyword string, order string, perPage int, page int) (returnData []interface{}) {
	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"match_all": map[string]interface{}{}, //搜索全部
	}

	if keyword != "" {
		arr := strings.Split(keyword, ":")
		if len(arr) > 1 {
			query = map[string]interface{}{
				"match": map[string]interface{}{ //搜索关键字
					strings.TrimSpace(arr[0]): strings.TrimSpace(arr[1]), //按字段查询
				},
			}
		}
	}

	sort := []map[string]map[string]string{
		map[string]map[string]string{
			"created_at": map[string]string{"order": order},
		},
	}

	queryData := map[string]interface{}{
		"query": query,
		"sort":  sort,
	}

	if err := json.NewEncoder(&buf).Encode(queryData); err != nil {
		panic(err)
	}

	resp, err := this.conn.Search(
		this.conn.Search.WithContext(context.Background()),
		this.conn.Search.WithIndex(index),
		this.conn.Search.WithBody(&buf),
		this.conn.Search.WithTrackTotalHits(true),
		this.conn.Search.WithPretty(),
		this.conn.Search.WithFrom((page-1)*perPage),
		this.conn.Search.WithSize(perPage),
	)

	if err != nil {
		panic(err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		panic(err)
	}

	if r["hits"] != nil {
		for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
			returnData = append(returnData, hit.(map[string]interface{})["_source"])
		}
	}

	return returnData
}
