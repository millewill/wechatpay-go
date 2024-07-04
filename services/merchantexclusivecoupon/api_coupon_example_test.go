// Copyright 2021 Tencent Inc. All rights reserved.
//
// 营销商家券对外API
//
// No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
// API version: 0.0.11

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package merchantexclusivecoupon_test

import (
	"context"
	"log"

	"github.com/millewill/wechatpay-go/core"
	"github.com/millewill/wechatpay-go/core/option"
	"github.com/millewill/wechatpay-go/services/merchantexclusivecoupon"
	"github.com/millewill/wechatpay-go/utils"
)

func ExampleCouponApiService_AssociateTradeInfo() {
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

	svc := merchantexclusivecoupon.CouponApiService{Client: client}
	resp, result, err := svc.AssociateTradeInfo(ctx,
		merchantexclusivecoupon.AssociateTradeInfoRequest{
			StockId:      core.String("100088"),
			CouponCode:   core.String("sxxe34343434"),
			OutTradeNo:   core.String("MCH_102233445"),
			OutRequestNo: core.String("1002600620019090123143254435"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call AssociateTradeInfo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleCouponApiService_DeactivateCoupon() {
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

	svc := merchantexclusivecoupon.CouponApiService{Client: client}
	resp, result, err := svc.DeactivateCoupon(ctx,
		merchantexclusivecoupon.DeactivateCouponRequest{
			CouponCode:          core.String("sxxe34343434"),
			StockId:             core.String("1234567891"),
			DeactivateRequestNo: core.String("1002600620019090123143254436"),
			DeactivateReason:    core.String("此券使用时间设置错误"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call DeactivateCoupon err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleCouponApiService_DisassociateTradeInfo() {
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

	svc := merchantexclusivecoupon.CouponApiService{Client: client}
	resp, result, err := svc.DisassociateTradeInfo(ctx,
		merchantexclusivecoupon.DisassociateTradeInfoRequest{
			StockId:      core.String("100088"),
			CouponCode:   core.String("213dsadfsa"),
			OutTradeNo:   core.String("treads8a9f980"),
			OutRequestNo: core.String("fdsafdsafdsa231321"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call DisassociateTradeInfo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

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

	svc := merchantexclusivecoupon.CouponApiService{Client: client}
	resp, result, err := svc.ListCouponsByFilter(ctx,
		merchantexclusivecoupon.ListCouponsByFilterRequest{
			Openid:          core.String("2323dfsdf342342"),
			Appid:           core.String("wx233544546545989"),
			StockId:         core.String("100088"),
			CreatorMerchant: core.String("1000000001"),
			BelongMerchant:  core.String("1000000002"),
			SenderMerchant:  core.String("1000000003"),
			Offset:          core.Int64(0),
			Limit:           core.Int64(20),
			CouponState:     merchantexclusivecoupon.COUPONSTATUS_SENDED.Ptr(),
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

	svc := merchantexclusivecoupon.CouponApiService{Client: client}
	resp, result, err := svc.QueryCoupon(ctx,
		merchantexclusivecoupon.QueryCouponRequest{
			CouponCode: core.String("123446565767"),
			Appid:      core.String("wx233544546545989"),
			Openid:     core.String("2323dfsdf342342"),
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

func ExampleCouponApiService_ReturnCoupon() {
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

	svc := merchantexclusivecoupon.CouponApiService{Client: client}
	resp, result, err := svc.ReturnCoupon(ctx,
		merchantexclusivecoupon.ReturnCouponRequest{
			CouponCode:      core.String("sxxe34343434"),
			StockId:         core.String("1234567891"),
			ReturnRequestNo: core.String("1002600620019090123143254436"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ReturnCoupon err:%s", err)
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

	svc := merchantexclusivecoupon.CouponApiService{Client: client}
	resp, result, err := svc.SendCoupon(ctx,
		merchantexclusivecoupon.SendCouponRequest{
			Openid:       core.String("xsd3434454567676"),
			Appid:        core.String("wx1234567889999"),
			StockId:      core.String("12312354"),
			OutRequestNo: core.String("2335465"),
			CouponCode:   core.String("202007019999"),
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

func ExampleCouponApiService_SendGovCard() {
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

	svc := merchantexclusivecoupon.CouponApiService{Client: client}
	resp, result, err := svc.SendGovCard(ctx,
		merchantexclusivecoupon.SendGovCardRequest{
			CardId:       core.String("pIJMr5MMiIkO_93VtPyIiEk2DZ4w"),
			Appid:        core.String("wxc0b84a53ed8e8d29"),
			Openid:       core.String("obLatjhnqgy2syxrXVM3MJirbkdI"),
			OutRequestNo: core.String("oTYhjfdsahnssddj_0136"),
			SendTime:     core.String("2019-12-31T13:29:35+08:00"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call SendGovCard err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleCouponApiService_UseCoupon() {
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

	svc := merchantexclusivecoupon.CouponApiService{Client: client}
	resp, result, err := svc.UseCoupon(ctx,
		merchantexclusivecoupon.UseCouponRequest{
			CouponCode:   core.String("sxxe34343434"),
			StockId:      core.String("100088"),
			Appid:        core.String("wx1234567889999"),
			UseTime:      core.String("2015-05-20T13:29:35+08:00"),
			UseRequestNo: core.String("1002600620019090123143254435"),
			Openid:       core.String("xsd3434454567676"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call UseCoupon err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
