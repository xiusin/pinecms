package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/esapi"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/spf13/cast"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
)

type ElasticSearch struct {
	client *elasticsearch8.Client
}

type SearchResult struct {
	Total int64
	Hits  []map[string]any
	Aggs  map[string]any
}

func (e *ElasticSearch) Search(index string, params SearchParams) (any, error) {
	var buf bytes.Buffer
	searchBody := map[string]any{
		"query": params.Query,
		"from":  params.From,
		"size":  params.Size,
		"sort":  params.Sort,
	}

	if params.Aggs != nil {
		searchBody["aggs"] = params.Aggs
	}

	if err := json.NewEncoder(&buf).Encode(searchBody); err != nil {
		return nil, fmt.Errorf("error encoding query: %w", err)
	}

	res, err := e.client.Search(
		e.client.Search.WithContext(context.Background()),
		e.client.Search.WithIndex(params.Index),
		e.client.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting response: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]any
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("error parsing the response body: %w", err)
		}
		return nil, fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]any)["type"], e["error"].(map[string]any)["reason"])
	}

	var esResponse struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
			Hits []map[string]any `json:"hits"`
		} `json:"hits"`
		Aggregations map[string]any `json:"aggregations"`
	}
	if err := json.NewDecoder(res.Body).Decode(&esResponse); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %w", err)
	}

	result := SearchResult{
		Total: esResponse.Hits.Total.Value,
		Hits:  esResponse.Hits.Hits,
		Aggs:  esResponse.Aggregations,
	}
	return &result, nil
}

func (e *ElasticSearch) Index(index string, id string, doc map[string]any) (string, error) {
	data, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}

	var resp *esapi.Response
	if id != "" {
		resp, err = e.client.Index(index, bytes.NewReader(data), e.client.Index.WithDocumentID(id))
	} else {
		resp, err = e.client.Index(index, bytes.NewReader(data))
	}

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		var errMap map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&errMap); err != nil {
			return "", fmt.Errorf("error parsing the response body: %w", err)
		}
		return "", fmt.Errorf("elasticsearch error: %s", errMap["error"].(map[string]any)["reason"])
	}

	var r map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", err
	}

	return cast.ToString(r["_id"]), nil
}

func (e *ElasticSearch) Update(index string, id string, doc map[string]any) error {
	updateData := map[string]any{
		"doc": doc,
	}
	byts, err := json.Marshal(updateData)
	if err != nil {
		return err
	}
	_, err = e.client.Update(index, id, bytes.NewBuffer(byts))
	return err
}

func (e *ElasticSearch) Delete(index string, id string) error {
	_, err := e.client.Delete(index, id)
	return err
}

func NewElasticSearch() (ISearch, error) {
	cfg := config.DB()
	es8, err := elasticsearch8.NewClient(elasticsearch8.Config{
		Addresses: []string{cfg.Elastic.Url},
		Username:  cfg.Elastic.UserName,
		Password:  cfg.Elastic.Password,
	})
	if err != nil {
		return nil, err
	}
	return &ElasticSearch{client: es8}, nil
}
