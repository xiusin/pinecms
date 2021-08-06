package apidoc

import (
	apigateway "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apigateway/v20180808"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type tencentApiConf struct {
	SecretId  string `json:"-"`
	SecretKey string `json:"-"`
	Region    string `json:"-"`
	Endpoint  string `json:"-"`
}

func NewClient(cfg tencentApiConf) *apigateway.Client {
	credential := common.NewCredential(
		cfg.SecretId,
		cfg.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = cfg.Endpoint
	client, _ := apigateway.NewClient(credential, cfg.Region, cpf)
	return client
}
