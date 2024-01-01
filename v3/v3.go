package rm

type layout string

// layout versions :
const (
	LayoutV1 layout = "v1"
	LayoutV2 layout = "v2"
	LayoutV3 layout = "v3"
	LayoutV4 layout = "v4"
)

// types :
type (
	PaymentType   string
	PaymentMethod string
	PaymentStatus string
)

const (
	// payment types :
	PaymentTypeWeb    PaymentType = "WEB_PAYMENT"
	PaymentTypeMobile PaymentType = "MOBILE_PAYMENT"

	// payment methods :
	PaymentMethodWeChatMalaysia    PaymentMethod = "WECHAT_MY"
	PaymentMethodWeChatChina       PaymentMethod = "WECHAT_CN"
	PaymentMethodBoostMalaysia     PaymentMethod = "BOOST_MY"
	PaymentMethodPrestoMalaysia    PaymentMethod = "PRESTO_MY"
	PaymentMethodAlipayChina       PaymentMethod = "ALIPAY_CN"
	PaymentMethodTnGMalaysia       PaymentMethod = "TNG_MY"
	PaymentMethodGrabMalaysia      PaymentMethod = "GRABPAY_MY"
	PaymentMethodMaybankMalaysia   PaymentMethod = "MAYBANK_MY"
	PaymentMethodRazerPayMalaysia  PaymentMethod = "RAZERPAY_MY"
	PaymentMethodMCashMalaysia     PaymentMethod = "MCASH_MY"
	PaymentMethodShopeePayMalaysia PaymentMethod = "SHOPEEPAY_MY"
	PaymentMethodFpxMalaysia       PaymentMethod = "FPX_MY"
	PaymentMethodGoBizMalaysia     PaymentMethod = "GOBIZ_MY"

	// payment status :
	PaymentStatusInProcess     PaymentStatus = "IN_PROCESS"
	PaymentStatusSuccess       PaymentStatus = "SUCCESS"
	PaymentStatusFailed        PaymentStatus = "FAILED"
	PaymentStatusFullyRefunded PaymentStatus = "FULL_REFUNDED"
	PaymentStatusReserved      PaymentStatus = "REVERSED"
	PaymentStatusCancelled     PaymentStatus = "CANCELLED"
)
