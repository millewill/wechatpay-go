// Copyright 2021 Tencent Inc. All rights reserved.
//
// 点金计划对外API
//
// 特约商户点金计划管理API
//
// API version: 0.3.3

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package goldplan_test

import (
	"context"
	"log"

	"gitee.com/millewill/wechatpay-go/core"
	"gitee.com/millewill/wechatpay-go/core/option"
	"gitee.com/millewill/wechatpay-go/services/goldplan"
	"gitee.com/millewill/wechatpay-go/utils"
)

func ExampleStatusApiService_ChangeCustomPageStatus() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := goldplan.StatusApiService{Client: client}
	resp, result, err := svc.ChangeCustomPageStatus(ctx,
		goldplan.ChangeCustomPageStatusRequest{
			SubMchid:      core.String("1900000109"),
			OperationType: goldplan.OPERATIONTYPE_OPEN.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ChangeCustomPageStatus err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleStatusApiService_ChangeGoldPlanStatus() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := goldplan.StatusApiService{Client: client}
	resp, result, err := svc.ChangeGoldPlanStatus(ctx,
		goldplan.ChangeGoldPlanStatusRequest{
			SubMchid:          core.String("1900000109"),
			OperationType:     goldplan.OPERATIONTYPE_OPEN.Ptr(),
			OperationPayScene: goldplan.OPERATIONPAYSCENE_JSAPI_AND_MINIPROGRAM.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ChangeGoldPlanStatus err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
