// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付分账API
//
// 微信支付分账API
//
// API version: 0.0.3

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package profitsharing_test

import (
	"context"
	"log"

	"gitee.com/millewill/wechatpay-go/core"
	"gitee.com/millewill/wechatpay-go/core/option"
	"gitee.com/millewill/wechatpay-go/services/profitsharing"
	"gitee.com/millewill/wechatpay-go/utils"
)

func ExampleOrdersApiService_CreateOrder() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := profitsharing.OrdersApiService{Client: client}
	resp, result, err := svc.CreateOrder(ctx,
		profitsharing.CreateOrderRequest{
			Appid:      core.String("wx8888888888888888"),
			OutOrderNo: core.String("P20150806125346"),
			Receivers: []profitsharing.CreateOrderReceiver{profitsharing.CreateOrderReceiver{
				Account:     core.String("86693852"),
				Amount:      core.Int64(888),
				Description: core.String("分给商户A"),
				Name:        core.String("hu89ohu89ohu89o"),
				Type:        core.String("MERCHANT_ID"),
			}},
			SubAppid:        core.String("wx8888888888888889"),
			SubMchid:        core.String("1900000109"),
			TransactionId:   core.String("4208450740201411110007820472"),
			UnfreezeUnsplit: core.Bool(true),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateOrder err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleOrdersApiService_QueryOrder() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := profitsharing.OrdersApiService{Client: client}
	resp, result, err := svc.QueryOrder(ctx,
		profitsharing.QueryOrderRequest{
			TransactionId: core.String("4208450740201411110007820472"),
			OutOrderNo:    core.String("P20150806125346"),
			SubMchid:      core.String("1900000109"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryOrder err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleOrdersApiService_UnfreezeOrder() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := profitsharing.OrdersApiService{Client: client}
	resp, result, err := svc.UnfreezeOrder(ctx,
		profitsharing.UnfreezeOrderRequest{
			Description:   core.String("解冻全部剩余资金"),
			OutOrderNo:    core.String("P20150806125346"),
			SubMchid:      core.String("1900000109"),
			TransactionId: core.String("4208450740201411110007820472"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call UnfreezeOrder err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
