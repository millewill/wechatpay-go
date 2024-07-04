// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付营销系统开放API
//
// 新增立减金api
//
// API version: 3.4.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package cashcoupons_test

import (
	"context"
	"log"

	"github.com/millewill/wechatpay-go/core"
	"github.com/millewill/wechatpay-go/core/option"
	"github.com/millewill/wechatpay-go/services/cashcoupons"
	"github.com/millewill/wechatpay-go/utils"
)

func ExampleCouponApiService_ListCouponsByFilter() {
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

	svc := cashcoupons.CouponApiService{Client: client}
	resp, result, err := svc.ListCouponsByFilter(ctx,
		cashcoupons.ListCouponsByFilterRequest{
			Openid:         core.String("Openid_example"),
			Appid:          core.String("Appid_example"),
			StockId:        core.String("9865000"),
			Status:         core.String("USED"),
			CreatorMchid:   core.String("9865002"),
			SenderMchid:    core.String("9865001"),
			AvailableMchid: core.String("9865000"),
			Offset:         core.Int64(0),
			Limit:          core.Int64(20),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListCouponsByFilter err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleCouponApiService_QueryCoupon() {
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

	svc := cashcoupons.CouponApiService{Client: client}
	resp, result, err := svc.QueryCoupon(ctx,
		cashcoupons.QueryCouponRequest{
			CouponId: core.String("9856888"),
			Appid:    core.String("Appid_example"),
			Openid:   core.String("Openid_example"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryCoupon err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleCouponApiService_SendCoupon() {
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

	svc := cashcoupons.CouponApiService{Client: client}
	resp, result, err := svc.SendCoupon(ctx,
		cashcoupons.SendCouponRequest{
			Openid:            core.String("Openid_example"),
			StockId:           core.String("example_stock_id"),
			OutRequestNo:      core.String("example_out_request_no"),
			Appid:             core.String("example_appid"),
			StockCreatorMchid: core.String("example_stock_creator_mchid"),
			CouponValue:       core.Int64(1),
			CouponMinimum:     core.Int64(1),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call SendCoupon err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
