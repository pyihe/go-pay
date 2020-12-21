package go_pay

import (
	"errors"
)

const (
	productionVerifyUrl = "https://buy.itunes.apple.com/verifyReceipt"
	sandboxVerifyUrl    = "https://sandbox.itunes.apple.com/verifyReceipt"
	googleBaseUrl       = "https://accounts.google.com/o/oauth2/token"
)

var (
	ErrInvalidToken = errors.New("invalid token data")
)

type (
	//内购事务
	LatestReceiptInfo struct {
		CancellationDate            string `json:"cancellation_date"`             //用户取消交易的日期，只有发生退款时才会出现本字段
		CancellationDateMs          string `json:"cancellation_date_ms"`          //用户取消交易或者自动续费订阅计划升级的时间，时间戳，精确到毫秒
		CancellationDatePst         string `json:"cancellation_date_pst"`         //苹果客户支持取消交易的时间，或自动续费订阅计划升级的时间，太平洋时区，退款事务才显示
		CancellationReason          string `json:"cancellation_reason"`           //退款原因，1-表示客户由于应用程序中的实际问题或感知到的问题而取消了其交易; 0-指示由于其他原因取消了交易；例如，如果客户意外地进行了购
		ExpiresDate                 string `json:"expires_date"`                  //订阅过期或者续订日期
		ExpiresDateMs               string `json:"expires_date_ms"`               //订阅过期或者续订时间，时间戳，精确到毫秒
		ExpiresDatePst              string `json:"expires_date_pst"`              //太平洋时区，订阅过期或者续订时间
		InAppOwnershipType          string `json:"in_app_ownership_type"`         //用户与他们可以访问的家庭共享购买的关系, FAMILY_SHARED-交易属于受益于服务的家庭成员, PURCHASED-交易属于购买者
		IsInIntroOfferPeriod        string `json:"is_in_intro_offer_period"`      //自动续费订阅是否处于介绍价格期
		IsTrialPeriod               string `json:"is_trial_period"`               //订阅是否处于免费试用期
		IsUpgraded                  string `json:"is_upgraded"`                   //订阅是否因为升级而取消
		OfferCodeRefName            string `json:"offer_code_ref_name"`           //用户挽回的订阅代码
		OriginalPurchaseDate        string `json:"original_purchase_date"`        //原始的应用内购买的时间, ISO 8601
		OriginalPurchaseDateMs      string `json:"original_purchase_date_ms"`     //原始的应用内购买的时间, 精确到毫秒的时间戳
		OriginalPurchaseDatePst     string `json:"original_purchase_date_pst"`    //原始的应用内购买的时间, 太平洋时区
		OriginalTransactionId       string `json:"original_transaction_id"`       //原始的交易ID
		ProductId                   string `json:"product_id"`                    //产品ID
		PromotionalOfferId          string `json:"promotional_offer_id"`          //用户兑换的自动续订订阅的促销优惠的标识符
		PurchaseDate                string `json:"purchase_date"`                 //消费日期
		PurchaseDateMs              string `json:"purchase_date_ms"`              //消费时间戳
		PurchaseDatePst             string `json:"purchase_date_pst"`             //消费时间
		Quantity                    string `json:"quantity"`                      //商品数量
		SubscriptionGroupIdentifier string `json:"subscription_group_identifier"` //订阅属于哪个订阅组
		WebOrderLineItemId          string `json:"web_order_line_item_id"`        //跨设备购买事件的唯一标识符，包括订阅续订事件
		TransactionId               string `json:"transaction_id"`                //交易ID
	}

	PendingRenewalInfo struct {
		AutoRenewProductId        string `json:"auto_renew_product_id"`
		AutoRenewStatus           string `json:"auto_renew_status"`
		ExpirationIntent          string `json:"expiration_intent"`
		GracePeriodExpiresDate    string `json:"grace_period_expires_date"`
		GracePeriodExpiresDateMs  string `json:"grace_period_expires_date_ms"`
		GracePeriodExpiresDatePst string `json:"grace_period_expires_date_pst"`
		IsInBillingRetryPeriod    string `json:"is_in_billing_retry_period"`
		OfferCodeRefName          string `json:"offer_code_ref_name"`
		OriginalTransactionId     string `json:"original_transaction_id"`
		PriceConsentStatus        string `json:"price_consent_status"`
		ProductId                 string `json:"product_id"`
	}

	InApp struct {
		CancellationDate        string `json:"cancellation_date"`
		CancellationDateMs      string `json:"cancellation_date_ms"`
		CancellationDatePst     string `json:"cancellation_date_pst"`
		CancellationReason      string `json:"cancellation_reason"`
		ExpiresDate             string `json:"expires_date"`
		ExpiresDateMs           string `json:"expires_date_ms"`
		ExpiresDatePst          string `json:"expires_date_pst"`
		IsInIntroOfferPeriod    string `json:"is_in_intro_offer_period"`
		IsTrialPeriod           string `json:"is_trial_period"`
		OriginalPurchaseDate    string `json:"original_purchase_date"`
		OriginalPurchaseDateMs  string `json:"original_purchase_date_ms"`
		OriginalPurchaseDatePst string `json:"original_purchase_date_pst"`
		OriginalTransactionId   string `json:"original_transaction_id"`
		ProductId               string `json:"product_id"`
		PromotionalOfferId      string `json:"promotional_offer_id"`
		PurchaseDate            string `json:"purchase_date"`
		PurchaseDateMs          string `json:"purchase_date_ms"`
		PurchaseDatePst         string `json:"purchase_date_pst"`
		Quantity                string `json:"quantity"`
		TransactionId           string `json:"transaction_id"`
		WebOrderLineItemId      string `json:"web_order_line_item_id"`
	}

	Receipt struct {
		AdamId                     int      `json:"adam_id"`
		AppItemId                  int      `json:"app_item_id"`
		ApplicationVersion         string   `json:"application_version"`
		BundleId                   string   `json:"bundle_id"`
		DownloadId                 int      `json:"download_id"`
		ExpirationDate             string   `json:"expiration_date"`
		ExpirationDateMs           string   `json:"expiration_date_ms"`
		ExpirationDatePst          string   `json:"expiration_date_pst"`
		InApp                      []*InApp `json:"in_app"`
		OriginalApplicationVersion string   `json:"original_application_version"`
		OriginalPurchaseDate       string   `json:"original_purchase_date"`
		OriginalPurchaseDateMs     string   `json:"original_purchase_date_ms"`
		OriginalPurchaseDatePst    string   `json:"original_purchase_date_pst"`
		PreorderDate               string   `json:"preorder_date"`
		PreorderDateMs             string   `json:"preorder_date_ms"`
		PreorderDatePst            string   `json:"preorder_date_pst"`
		ReceiptCreationDate        string   `json:"receipt_creation_date"`
		ReceiptCreationDateMs      string   `json:"receipt_creation_date_ms"`
		ReceiptCreationDatePst     string   `json:"receipt_creation_date_pst"`
		ReceiptType                string   `json:"receipt_type"`
		RequestDate                string   `json:"request_date"`
		RequestDateMs              string   `json:"request_date_ms"`
		RequestDatePst             string   `json:"request_date_pst"`
		VersionExternalIdentifier  int      `json:"version_external_identifier"`
	}

	ApplePayResponse struct {
		Environment        string              `json:"environment"`    //验证环境, 测试环境-Sandbox, 生产环境-Production
		IsRetryable        bool                `json:"is-retryable"`   //验证期间的错误指示器, 只有当status为21100-21199时适用, 值为1时表示临时问题，可以重新尝试验证，值为0时表示无法解决的问题, 不能重试
		LatestReceipt      []byte              `json:"latest_receipt"` //最新的Base64编码的请求回执, 只有当收据包含自动续订订阅时才返回
		LatestReceiptInfo  *LatestReceiptInfo  `json:"latest_receipt_info"`
		PendingRenewalInfo *PendingRenewalInfo `json:"pending_renewal_info"`
		Receipt            *Receipt            `json:"receipt"`
		Status             int                 `json:"status"`
	}

	GooglePayResponse struct {
		Kind                        string `json:"kind"`                        //服务类别
		PurchaseTimeMillis          string `json:"purchaseTimeMillis"`          //商品购买的时间戳，精确到毫秒
		PurchaseState               int    `json:"purchaseState"`               //商品（订单）购买状态: 0-已购买 1-取消 2-进行中
		ConsumptionState            int    `json:"consumptionState"`            //商品消耗状态: 0-未消费 1-已消费
		DeveloperPayload            string `json:"developerPayload"`            //开发人员指定的订单补充信息
		OrderId                     string `json:"orderId"`                     //商品关联的Google订单ID
		PurchaseType                int    `json:"purchaseType"`                //交易类型: 0-测试账号购买 1-使用促销码购买 2-来自于观看视频
		AcknowledgementState        int    `json:"acknowledgementState"`        //商品确认状态: 0-待确认 1-已确认
		PurchaseToken               string `json:"purchaseToken"`               //用于标识此次交易的token
		ProductId                   string `json:"productId"`                   //商品ID
		Quantity                    int    `json:"quantity"`                    //商品数量
		ObfuscatedExternalAccountId string `json:"obfuscatedExternalAccountId"` //与应用程序中的用户帐户唯一关联的id的模糊版本
		ObfuscatedExternalProfileId string `json:"obfuscatedExternalProfileId"` //与应用程序中的用户配置文件唯一关联的id的模糊版本
		RegionCode                  string `json:"regionCode"`                  //区域
	}
)

