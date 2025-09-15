package search

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/spf13/cast"

	"github.com/xiusin/pinecms/src/config"
	client "github.com/zinclabs/sdk-go-zincsearch"
)

type PineZincSearch struct {
	client *client.APIClient
	ctx    context.Context
}

func (p *PineZincSearch) Search(index string, params SearchParams) (any, error) {
	// TODO: Fix this implementation
	return nil, nil
}

func (p *PineZincSearch) Update(index, id string, doc map[string]any) error {
	_, _, err := p.client.Document.Update(p.ctx, index, id).Document(doc).Execute()
	return err
}

func (p *PineZincSearch) Delete(index, id string) error {
	_, _, err := p.client.Document.Delete(p.ctx, index, id).Execute()
	return err
}

func (p *PineZincSearch) Index(index string, id string, doc map[string]any) (string, error) {
	if len(id) > 0 {
		doc["_id"] = id
	}
	resp, _, err := p.client.Document.Index(p.ctx, index).Document(doc).Execute()
	if err != nil {
		return "", err
	}
	return resp.GetId(), nil
}

func NewZincSearch() ISearch {
	cfg := config.DB()
	ctx := context.WithValue(context.Background(), client.ContextBasicAuth, client.BasicAuth{
		UserName: cfg.Elastic.UserName,
		Password: cfg.Elastic.Password,
	})
	configuration := client.NewConfiguration()
	configuration.Servers = client.ServerConfigurations{
		client.ServerConfiguration{URL: cfg.Elastic.Url},
	}
	return &PineZincSearch{client: client.NewAPIClient(configuration), ctx: ctx}
}
