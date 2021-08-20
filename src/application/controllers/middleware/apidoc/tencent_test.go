package apidoc

//func TestApiDoc(t *testing.T) {
//	cfg := tencentApiConf{
//		SecretId:  "AKIDRmCdPINEx6p13vfFdtiqlW6wd8RYUGuMBliT...",
//		SecretKey: "jHk4Dhz74pRuEPYWI4iibQ4YuPINE5966wb5...",
//		Region:    "ap-beijing",
//		Endpoint:  "apigateway.tencentcloudapi.com",
//	}
//
//	client := NewClient(cfg)
//
//	// 创建api
//	request := apigateway.NewCreateApiRequest()
//	request.ServiceId = common.StringPtr("service-q8ugahkx")
//	request.ApiName = common.StringPtr("会员列表")
//	request.ApiDesc = common.StringPtr("列举所有会员信息")
//	request.ApiType = common.StringPtr("NORMAL")
//
//	request.Protocol = common.StringPtr("HTTP")
//	request.EnableCORS = common.BoolPtr(true)
//
//
//	request.RequestConfig = &apigateway.ApiRequestConfig{
//		Path:   common.StringPtr("/user/list/detail"),
//		Method: common.StringPtr("GET"),
//	}
//
//	request.ServiceType = common.StringPtr("MOCK")
//	request.ServiceTimeout = common.Int64Ptr(60)
//	request.ServiceMockReturnMessage = common.StringPtr(`{"code": 200, "msg": "mock success", "list": []}`)
//
//	request.RequestParameters = []*apigateway.RequestParameter {
//		{
//			Name: common.StringPtr("page"),
//			Desc: common.StringPtr("页码"),
//			Position: common.StringPtr("QUERY"),
//			Type: common.StringPtr("int"),
//			DefaultValue: common.StringPtr("1"),
//			Required: common.BoolPtr(true),
//		},
//		{
//			Name: common.StringPtr("pageSize"),
//			Desc: common.StringPtr("分页条目数"),
//			Position: common.StringPtr("QUERY"),
//			Type: common.StringPtr("int"),
//			DefaultValue: common.StringPtr("10"),
//			Required: common.BoolPtr(true),
//		},
//	}
//
//	request.ResponseSuccessExample = common.StringPtr(`{"code": 200, "msg": "success", "list": []}`)
//	request.ResponseFailExample = common.StringPtr(`{"code": 500, "msg": "failed", "list": []}`)
//	request.ResponseType = common.StringPtr("JSON")
//
//	response, err := client.CreateApi(request)
//	if _, ok := err.(*errors.TencentCloudSDKError); ok {
//		fmt.Printf("An API error has returned: %s", err)
//		return
//	}
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("%s", response.ToJsonString())
//
//}
