// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微工卡接口文档
//
// 服务商通过本API文档提供的接口，查询商户和微工卡的授权关系、生成预授权的token口令、核身预下单、核身结果的查询等。
//
// API version: 1.5.2

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package payrollcard_test

import (
	"context"
	"log"

	"gitee.com/millewill/wechatpay-go/core"
	"gitee.com/millewill/wechatpay-go/core/option"
	"gitee.com/millewill/wechatpay-go/services/payrollcard"
	"gitee.com/millewill/wechatpay-go/utils"
)

func ExampleAuthenticationsApiService_GetAuthentication() {
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

	svc := payrollcard.AuthenticationsApiService{Client: client}
	resp, result, err := svc.GetAuthentication(ctx,
		payrollcard.GetAuthenticationRequest{
			SubMchid:           core.String("1111111"),
			AuthenticateNumber: core.String("mcdhehfgisdhfjghed39384564i83"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetAuthentication err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleAuthenticationsApiService_ListAuthentications() {
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

	svc := payrollcard.AuthenticationsApiService{Client: client}
	resp, result, err := svc.ListAuthentications(ctx,
		payrollcard.ListAuthenticationsRequest{
			Openid:            core.String("onqOjjmo8wmTOOtSKwXtGjg9Gb58"),
			Appid:             core.String("wxa1111111"),
			SubAppid:          core.String("wxa1111111"),
			SubMchid:          core.String("1111111"),
			AuthenticateDate:  core.String("2020-12-25"),
			AuthenticateState: core.String("AUTHENTICATE_SUCCESS"),
			Offset:            core.Int64(0),
			Limit:             core.Int64(10),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListAuthentications err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleAuthenticationsApiService_PreOrderAuthentication() {
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

	svc := payrollcard.AuthenticationsApiService{Client: client}
	resp, result, err := svc.PreOrderAuthentication(ctx,
		payrollcard.PreOrderAuthenticationRequest{
			Openid:             core.String("onqOjjmo8wmTOOtSKwXtGjg9Gb58"),
			Appid:              core.String("wxa1111111"),
			SubMchid:           core.String("1111111"),
			SubAppid:           core.String("wxa1111111"),
			AuthenticateNumber: core.String("mcdhehfgisdhfjghed39384564i83"),
			ProjectName:        core.String("某项目"),
			EmployerName:       core.String("某用工企业"),
			AuthenticateType:   payrollcard.AUTHENTICATIONTYPE_NORMAL.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call PreOrderAuthentication err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleAuthenticationsApiService_PreOrderAuthenticationWithAuth() {
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

	svc := payrollcard.AuthenticationsApiService{Client: client}
	resp, result, err := svc.PreOrderAuthenticationWithAuth(ctx,
		payrollcard.PreOrderAuthenticationWithAuthRequest{
			Openid:             core.String("onqOjjmo8wmTOOtSKwXtGjg9Gb58"),
			Appid:              core.String("wxa1111111"),
			SubMchid:           core.String("1111111"),
			SubAppid:           core.String("wxa1111111"),
			AuthenticateNumber: core.String("mcdhehfgisdhfjghed39384564i83"),
			ProjectName:        core.String("某项目"),
			EmployerName:       core.String("某用工企业"),
			UserName:           core.String("LP7bT4hQXUsOZCEvK2YrSiqFsnP0oRMfeoLN0vBg"),
			IdCardNumber:       core.String("7FzH5XksJG3a8HLLsaaUV6K54y1OnPMY5"),
			EmploymentType:     payrollcard.EMPLOYMENTTYPE_LONG_TERM_EMPLOYMENT.Ptr(),
			AuthenticateType:   payrollcard.AUTHENTICATIONTYPE_NORMAL.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call PreOrderAuthenticationWithAuth err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
