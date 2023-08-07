package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/joho/godotenv"
	"github.com/tranvannghia021/gocore/helpers"
	"log"
	"os"
	"strings"
)

var (
	keyMappingMsg = "ElasticSearch {variable} is required"
)

type elasticSearch interface {
	SetIndex(index string) elasticSearch
	GetIndex(index string) string
	SetType(typeEs string) elasticSearch
	GetType() string
	GetAllIndex() []string
	CreateIndex() bool
	DeleteIndex() bool
	DeleteAllIndex() bool
	GetDocument(id string) (string, bool)
	GetAllDocument() (string, bool)
	CreateDocument(document interface{}) (string, bool)
	UpdateDocument(id string, document string) bool
	DeleteDocument(id string) bool
	DeleteAllDocument() bool
	Search(query string) (string, bool)
}
type ElasticSearchStructCat []struct {
	Health       string `json:"health"`
	Status       string `json:"status"`
	Index        string `json:"index"`
	UUID         string `json:"uuid"`
	Pri          string `json:"pri"`
	Rep          string `json:"rep"`
	DocsCount    string `json:"docs.count"`
	DocsDeleted  string `json:"docs.deleted"`
	StoreSize    string `json:"store.size"`
	PriStoreSize string `json:"pri.store.size"`
}

type ElasticSearch struct {
	host   string
	port   string
	schema string
	user   string
	pass   string
	index  string
	typeEs string
	client *elasticsearch.Client
}

type handleResponseEs struct {
	Status bool
	Error  error
	Data   map[string]any
}

func NewElasticSearch(instance *ElasticSearch) elasticSearch {
	if godotenv.Load() == nil {
		host, _ := os.LookupEnv("ES_HOST")
		port, _ := os.LookupEnv("ES_PORT")
		schema, _ := os.LookupEnv("ES_SCHEMA")
		user, _ := os.LookupEnv("ES_USER")
		pass, _ := os.LookupEnv("ES_PASS")
		instance.host = host
		instance.port = port
		instance.schema = schema
		instance.user = user
		instance.pass = pass
	}
	instance.domainError()
	instance.portError()
	instance.schemaError()
	instance.userError()
	instance.passError()
	cfg := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("%s://%s:%s", instance.schema, instance.host, instance.port),
		},
		Username: instance.user,
		Password: instance.pass,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	instance.client = es
	return instance
}
func (e *ElasticSearch) domainError() {
	if e.host == "" {

		log.Fatal(strings.Replace(keyMappingMsg, "{variable}", "host", 1))
	}
}

func (e *ElasticSearch) portError() {
	if e.port == "" {

		log.Fatal(strings.Replace(keyMappingMsg, "{variable}", "port", 1))
	}
}

func (e *ElasticSearch) schemaError() {
	if e.schema == "" {

		log.Fatal(strings.Replace(keyMappingMsg, "{variable}", "schema", 1))
	}
}

func (e *ElasticSearch) userError() {
	if e.user == "" {

		log.Fatal(strings.Replace(keyMappingMsg, "{variable}", "user", 1))
	}
}

func (e *ElasticSearch) passError() {
	if e.pass == "" {

		log.Fatal(strings.Replace(keyMappingMsg, "{variable}", "password", 1))
	}
}

func (e *ElasticSearch) validateIndex() {
	if e.index == "" {
		log.Fatal(strings.Replace(keyMappingMsg, "{variable}", "index", 1))
	}

}

func (e *ElasticSearch) validateType() {
	if e.typeEs == "" {
		e.typeEs = "es_type"
	}

}
func (e *ElasticSearch) SetIndex(index string) elasticSearch {
	e.index = index
	return e
}
func (e *ElasticSearch) GetIndex(index string) string {
	return e.index
}
func (e *ElasticSearch) SetType(typeEs string) elasticSearch {
	e.typeEs = typeEs
	return e
}
func (e *ElasticSearch) GetType() string {
	e.validateType()
	return e.typeEs
}
func (e *ElasticSearch) CreateIndex() bool {
	e.validateIndex()
	e.client.Indices.Create(e.index)
	return true
}
func (e *ElasticSearch) DeleteIndex() bool {
	e.validateIndex()
	e.client.Indices.Delete([]string{e.index})
	return true
}
func (e *ElasticSearch) CreateDocument(document interface{}) (string, bool) {
	e.validateIndex()
	data, _ := json.Marshal(document)
	res, _ := e.client.Index(e.index, bytes.NewReader(data))
	return helpers.TrimResponseToStringJsonEs(res.String(), "[201 Created]"), !res.IsError()

}
func (e *ElasticSearch) UpdateDocument(id string, document string) bool {
	e.validateIndex()
	res, _ := e.client.Update(e.index, id, strings.NewReader(document))
	return !res.IsError()
}
func (e *ElasticSearch) DeleteDocument(id string) bool {
	e.validateIndex()
	res, _ := e.client.Delete(e.index, id)
	return !res.IsError()
}
func (e *ElasticSearch) DeleteAllDocument() bool {
	e.validateIndex()
	e.DeleteIndex()
	e.CreateIndex()
	return true
}

func (e *ElasticSearch) DeleteAllIndex() bool {
	listIndex := e.GetAllIndex()
	for _, index := range listIndex {
		e.SetIndex(index).DeleteIndex()
	}
	return true
}
func (e *ElasticSearch) GetAllIndex() []string {
	res, _ := esapi.CatIndicesRequest{Format: "json"}.Do(context.Background(), e.client)
	var listIndex ElasticSearchStructCat

	json.Unmarshal([]byte(helpers.TrimResponseToStringJsonEs(res.String(), "")), &listIndex)
	var result []string
	for _, index := range listIndex {
		result = append(result, index.Index)
	}
	return result
}

func (e *ElasticSearch) GetDocument(id string) (string, bool) {
	e.validateIndex()
	res, _ := e.client.Get(e.index, id)
	return helpers.TrimResponseToStringJsonEs(res.String(), ""), !res.IsError()
}
func (e *ElasticSearch) GetAllDocument() (string, bool) {
	e.validateIndex()
	res, _ := e.client.Search(
		e.client.Search.WithIndex(e.index),
		e.client.Search.WithBody(strings.NewReader(`{
												  "query": {
													"match_all": {}
												  }
												}`)),
	)
	return helpers.TrimResponseToStringJsonEs(res.String(), ""), !res.IsError()
}

func (e *ElasticSearch) Search(query string) (string, bool) {
	e.validateIndex()
	res, _ := e.client.Search(
		e.client.Search.WithIndex(e.index),
		e.client.Search.WithBody(strings.NewReader(query)),
	)
	return helpers.TrimResponseToStringJsonEs(res.String(), ""), !res.IsError()

}
