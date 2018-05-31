package client_cashdesk

import (
	golib_tools_courier_enumeration "golib/tools/courier/enumeration"
	golib_tools_timelib "golib/tools/timelib"
)

type AreaBank struct {
	//
	PabSubbranchBank
	// 银行编码
	BankID uint32 `json:"bankID"`
	// G7地区编码
	G7AreaCode string `json:"g7AreaCode"`
	// G7上级地区编码
	G7TopAreaCode1 string `json:"g7TopAreaCode1"`
	// G7上级地区编码
	G7TopAreaCode2 string `json:"g7TopAreaCode2"`
	// G7上级地区编码
	G7TopAreaCode3 string `json:"g7TopeAreaCode3"`
}

type AreaListByParentCodeDepthBody struct {
	// 地址编码
	AddCode string `json:"addCode"`
	// 创建时间
	CreateTime golib_tools_timelib.MySQLTimestamp `json:"createTime"`
	// 深度
	Depth int32 `json:"depth"`
	// 地址全称
	FullName string `json:"fullName"`
	// 主键ID
	ID uint64 `json:"id"`
	// 地址名称
	Name string `json:"name"`
	// 上级地区编码
	ParentCode string `json:"parentCode"`
	// 拼音
	PinYin string `json:"pinyin"`
	// 顶级地区编码
	TopCode string `json:"topCode"`
	// 修改时间
	UpdateTime golib_tools_timelib.MySQLTimestamp `json:"updateTime"`
	// 版本
	Version string `json:"version"`
}

type Bank struct {
	//
	BankInfo
	// 银行拼音名称首字母
	FirstPinyin string `json:"firstPinyin"`
	// 银行拼音
	Pinyin string `json:"pinyin"`
}

type BankInfo struct {
	//
	BankModel
	// 是否支持银联鉴权
	IsSupportUnionpayAuth golib_tools_courier_enumeration.Bool `json:"isSupportUnionpayAuth"`
	// 超级网银联行号
	SuperBankCode string `json:"superBankCode"`
}

type BankModel struct {
	// 银行logo
	BankLogo string `json:"bankLogo"`
	// 银行名
	BankName string `json:"bankName"`
	// 银行ID
	ID uint32 `json:"id"`
}

type CancelTransBody struct {
	// 注释
	Comments string `default:"" json:"comments" validate:"@string[0,1024]"`
}

type CreateTransReqBody struct {
	// 注释
	Comment string `default:"" json:"comment" validate:"@string[0,1024]"`
	// 币种
	Currency CashdeskCurrencyType `default:"1" json:"currency"`
	// 商品或服务名称
	GoodsName string `default:"" json:"goodsName" validate:"@string[1,32]"`
	// 商品或服务的url地址
	GoodsUrl string `default:"" json:"goodsUrl" validate:"@httpUrlOrEmpty"`
	// 支付成功后的通知地址
	NotifyUrl string `default:"" json:"notifyUrl" validate:"@httpUrl"`
	// 订单号
	OrderNo string `default:"" json:"orderNo" validate:"@string[1,64]"`
	// 用户终端类型
	PlatformType CashdeskPlatformType `json:"platformType"`
	// 随机字符串
	RandString string `default:"" json:"randString" validate:"@string[1,32]"`
	// 支付成功后的跳转地址
	ReturnUrl string `default:"" json:"returnUrl" validate:"@httpUrlOrEmpty"`
	// 用户终端ip
	SourceIP string `default:"" json:"sourceIP" validate:"@ipv4OrEmpty"`
	// 总金额(单位为分)
	TotalAmount int64 `default:"" json:"totalAmount" validate:"@int64[1,9007199254740991]"`
	// user-agent
	UserAgent string `default:"" json:"userAgent" validate:"@string[0,512]"`
	// 用户来源
	UserFrom CashdeskUserFrom `default:"" json:"userFrom" validate:"@string{,G7,ANONYMOUS}"`
	// 外部用户ID
	UserID string `default:"" json:"userID" validate:"@string[0,64]"`
}

type CreateTransRespBody struct {
	// 创建时间
	// Required : true
	CreateTime golib_tools_timelib.MySQLTimestamp `json:"createTime"`
	// 交易单id
	// Required : true
	TransID uint64 `json:"transID,string"`
}

type EBankPayReqBody struct {
	// 充值金额或交易单金额，如果有交易单，可不传
	Amount int64 `default:"0" json:"amount" validate:"@int64[0,1000000000]"`
	// 通道编号
	ChannelCode string `default:"" json:"channelCode"`
	// 外部用户id
	UserID string `default:"" json:"userID" validate:"@string[1,64]"`
}

type ErrorField struct {
	// 出错字段路径
	// 这个信息为一个 json 字符串,方便客户端进行定位错误原因
	// 例如输入中 {"name":{ "alias" : "test"}} 中的alias出错,则返回 "name.alias"
	// 如果alias是数组, 且第2个元素的a字段错误,则返回"name.alias[2].a"
	Field string `json:"field"`
	// 错误字段位置
	// body, query, header, path, formData
	In string `json:"in"`
	// 错误信息
	Msg string `json:"msg"`
}

type ErrorFields []ErrorField

type ExtQueryTransByIDRespBody struct {
	//
	Trans ExtTrans `json:"trans"`
}

type ExtTrans struct {
	//
	Trans
	//
	PayStatus CashdeskPayState `json:"payStatus"`
}

type GetAreaBankByBankCodeRespBody struct {
	// 分支行信息
	AreaBank AreaBank `json:"areaBank"`
	// 一级银行信息
	Bank Bank `json:"bank"`
}

type GetAreaBankListRespBody struct {
	// 分页数据
	Data []AreaBank `json:"data"`
	// 总数
	Total int32 `json:"total"`
}

type GetAreaListByParentCodeDepthRespBody struct {
	// 数据
	Data []AreaListByParentCodeDepthBody `json:"data"`
	// 总数
	Total int32 `json:"total"`
}

type GetSignRespBody struct {
	// 签名
	// Required : true
	Sign string `json:"sign"`
}

type KVPair struct {
	// 参数
	Params string `default:"" json:"params" validate:"@string[1,128]"`
	// 参数值
	ParamsValue string `default:"" json:"paramsValue" validate:"@string[0,1024]"`
}

type OldTransQueryRespBody struct {
	//
	ExtQueryTransByIDRespBody
	// 商户名
	MerchantName string `json:"merchantName"`
	// 外部用户ID
	UserID string `json:"userID"`
}

type PabLinkModel struct {
	// 客户名称
	AccountName string `json:"accountName"`
	// 开户类型
	AccountType CashdeskPabAccountType `json:"accountType"`
	// 证件号码
	CertificateCode string `json:"certificateCode"`
	// 证件类型
	CertificateType CashdeskCertificateType `default:"" json:"certificateType" validate:"@string{ID_CARD,ORG_INSITITUTE_CODE,UNITY_SOCIAL_CREDIT_CODE}"`
	// 见证宝账户手机号码
	Mobile string `json:"mobile"`
	// 见证宝账户余额
	PabAmount int64 `json:"pabAmount"`
	// 开户状态
	State CashdeskPabAccountState `json:"state"`
	// 易宝账户余额
	SubAccountAmount int64 `json:"subAccountAmount"`
	// 外部用户来源
	UserFrom CashdeskUserFrom `json:"userFrom"`
	// 外部用户ID
	UserID string `json:"userID"`
}

type PabPayReqBody struct {
	// 交易金额
	Amount int64 `default:"" json:"amount" validate:"@int64[0,9007199254740991]"`
	// 证件号码
	CertificateCode string `default:"" json:"certificateCode" validate:"@string[1,32]"`
	// 证件类型
	CertificateType CashdeskCertificateType `default:"" json:"certificateType" validate:"@string{ID_CARD,ORG_INSITITUTE_CODE,UNITY_SOCIAL_CREDIT_CODE}"`
	// 外部用户id
	UserID string `default:"" json:"userID" validate:"@string[1,64]"`
}

type PabPayRespBody struct {
	// 交易金额
	Amount int64 `json:"amount"`
	// 订单号
	OrderNo string `json:"orderNo"`
	// 交易状态
	PayStatus CashdeskPayState `json:"payStatus"`
	// 交易号
	TransID uint64 `json:"transID,string"`
	// 外部用户id
	UserID string `json:"userID"`
}

type PabSubbranchBank struct {
	// 地区编码
	CityCode string `json:"areaCode"`
	// 分支行编码
	BankCode string `json:"bankCode"`
	// 分支行名字
	BankName string `json:"bankName"`
	// 拼音
	PinYin string `json:"pinyin"`
}

type PabWhiteListTrans struct {
	// 买家账号ID
	BuyerAccountID uint64 `json:"buyerAccountID,string"`
	// 交易创建时间
	CreateTime golib_tools_timelib.MySQLTimestamp `json:"createTime"`
	// 卖家账号ID
	SellerAccountID uint64 `json:"sellerAccountID,string"`
	// 交易额
	TransAmount int64 `json:"transAmount"`
	// 交易单ID
	TransID uint64 `json:"transID,string"`
	// 外部订单/交易号
	TransIDExt string `json:"transIDExt"`
}

type PabWhiteListTransList []PabWhiteListTrans

type PabWhiteListTransResp struct {
	//
	Data PabWhiteListTransList `json:"data"`
	//
	Total int32 `json:"total"`
}

type QueryTransResponseBody struct {
	//
	Data []TransQuery `json:"data"`
	//
	Total int32 `json:"total"`
}

type ReturnTransModel struct {
	//
	Trans
	// 商户名称
	MerchantName string `json:"merchantName"`
	// 支付状态
	PayStatus CashdeskPayState `json:"payStatus"`
	// 用户来源
	UserFrom CashdeskUserFrom `json:"userFrom"`
	// 外部用户ID
	UserID string `json:"userID"`
}

type StatusError struct {
	// 是否能作为错误话术
	CanBeErrorTalk bool `json:"canBeTalkError"`
	// 错误代码
	Code int64 `json:"code"`
	// 详细描述
	Desc string `json:"desc"`
	// 出错字段
	ErrorFields ErrorFields `json:"errorFields"`
	// 请求 ID
	ID string `json:"id"`
	// 错误 Key
	Key string `json:"key"`
	// 错误信息
	Msg string `json:"msg"`
	// 错误溯源
	Source []string `json:"source"`
}

type Trans struct {
	// 调账标志 1:非调账 2：调账记录
	AdjustFlag CashdeskPayAdjustFlag `json:"adjustFlag"`
	// 银行端付款时间
	BankPayTime golib_tools_timelib.MySQLTimestamp `json:"bankPayTime"`
	// 买家账号ID
	BuyerAccountID uint64 `json:"buyerAccountID,string"`
	// 买家外部通道账户ID
	BuyerChannelAccountID uint64 `json:"buyerChannelAccountID,string"`
	// 买家子账户类型
	BuyerSubAccountType CashdeskPaySubAccountType `json:"buyerSubAccountType"`
	// 交易是否可修改 1为可以 2为不可以
	CanModify golib_tools_courier_enumeration.Bool `json:"canModify"`
	// 实际支付金额
	CashAmount int64 `json:"cashAmount"`
	// 备注
	Comments string `json:"comments"`
	// 优惠券抵扣金额
	CouponAmount int64 `json:"couponAmount"`
	// 交易创建时间
	CreateTime golib_tools_timelib.MySQLTimestamp `json:"createTime"`
	// 币种
	Currency CashdeskPayCurrencyType `json:"currency"`
	// 交易结束时间
	EndTime golib_tools_timelib.MySQLTimestamp `json:"endTime"`
	// 结算账户ID,如果金额为0 此字段可为0
	FeeAccountID uint64 `json:"feeAccountID,string"`
	// 交易手续费
	FeeAmount int64 `json:"feeAmount"`
	// 结算子账户类型
	FeeSubAccountType CashdeskPaySubAccountType `json:"feeSubAccountType"`
	// 商品名称
	GoodsName string `json:"goodsName"`
	// 商品URL地址
	GoodsUrl string `json:"goodsUrl"`
	// 支付成功后的通知地址
	NotifyUrl string `json:"notifyUrl"`
	// 交易对应的收款单号
	PayRecvID uint64 `json:"payRecvID,string"`
	// 交易付款时间
	PayTime golib_tools_timelib.MySQLTimestamp `json:"payTime"`
	// 付款类型 1:余额支付 2：充值支付
	PayType CashdeskPayPayType `json:"payType"`
	// 平台类型 参看银行网关的的终端标识
	PlatformType CashdeskPayPlatformType `json:"platformType"`
	// 支付使用积分
	Points int64 `json:"points"`
	// 积分抵扣金额
	PointsAmount int64 `json:"pointsAmount"`
	// 支付成功后的跳转地址
	ReturnUrl string `json:"returnUrl"`
	// 卖家账号ID
	SellerAccountID uint64 `json:"sellerAccountID,string"`
	// 卖家外部通道账户ID
	SellerChannelAccountID uint64 `json:"sellerChannelAccountID,string"`
	// 卖家子账户类型
	SellerSubAccountType CashdeskPaySubAccountType `json:"sellerSubAccountType"`
	// 结算模式
	SettlementType CashdeskPaySettlementType `json:"settlementType"`
	// 商户ID
	SpID uint64 `json:"spID,string"`
	// 交易状态
	State CashdeskPayTransState `json:"state"`
	// 交易总金额
	TotalAmount int64 `json:"totalAmount"`
	// 交易单ID
	TransID uint64 `json:"transID,string"`
	// 外部订单/交易号
	TransIDExt string `json:"transIDExt"`
	// F_trans_mode 1: 中介 2：直付 3：预付 4：保理交易
	TransMode CashdeskPayTransMode `json:"transMode"`
	// 交易类型 1：ETC 2:油卡 3：招采
	TransType CashdeskPayTransType `json:"transType"`
	// user_agent
	UserAgent string `json:"userAgent"`
}

type TransQuery struct {
	// 银行端付款时间
	BankPayTime golib_tools_timelib.MySQLTimestamp `json:"bankPayTime"`
	// 买家账号ID
	BuyerAccountID uint64 `json:"buyerAccountID,string"`
	// 买家外部通道账户ID
	BuyerChannelAccountID uint64 `json:"buyerChannelAccountID,string"`
	// 卖家子账户类型
	BuyerSubAccountType CashdeskPaySubAccountType `json:"buyerSubAccountType"`
	// 实际支付金额
	CashAmount int64 `json:"cashAmount"`
	// 备注
	Comments string `json:"comments"`
	// 优惠券抵扣金额
	CouponAmount int64 `json:"couponAmount"`
	// 交易创建时间
	CreateTime golib_tools_timelib.MySQLTimestamp `json:"createTime"`
	// 币种
	Currency CashdeskPayCurrencyType `json:"currency"`
	// 交易结束时间
	EndTime golib_tools_timelib.MySQLTimestamp `json:"endTime"`
	// 交易手续费
	FeeAmount int64 `json:"feeAmount"`
	// 商品名称
	GoodsName string `json:"goodsName"`
	// 商品URL地址
	GoodsUrl string `json:"goodsUrl"`
	// 交易付款时间
	PayTime golib_tools_timelib.MySQLTimestamp `json:"payTime"`
	// 付款类型 1:余额支付 2：充值支付
	PayType CashdeskPayPayType `json:"payType"`
	// 支付使用积分
	Points int64 `json:"points"`
	// 积分抵扣金额
	PointsAmount int64 `json:"pointsAmount"`
	// 卖家账号ID
	SellerAccountID uint64 `json:"sellerAccountID,string"`
	// 卖家外部通道账户ID
	SellerChannelAccountID uint64 `json:"sellerChannelAccountID,string"`
	// 卖家子账户类型
	SellerSubAccountType CashdeskPaySubAccountType `json:"sellerSubAccountType"`
	// 结算模式
	SettlementType CashdeskPaySettlementType `json:"settlementType"`
	// 商户ID
	SpID uint64 `json:"spID,string"`
	// 交易状态
	State CashdeskPayTransState `json:"state"`
	// 交易总金额
	TotalAmount int64 `json:"totalAmount"`
	// 交易单ID
	TransID uint64 `json:"transID,string"`
	// 外部交易单号
	TransIDExt string `json:"transIDExt"`
	// F_trans_mode 1: 中介 2：直付 3：预付 4：保理交易
	TransMode CashdeskPayTransMode `json:"transMode"`
	// 交易类型 1：ETC 2:油卡 3：招采
	TransType CashdeskPayTransType `json:"transType"`
}

type UpdateTransReqBody struct {
	// 随机字符串
	RandString string `default:"" json:"randString" validate:"@string[1,32]"`
	// 更新订单金额(单位为分)
	TotalAmount int64 `default:"" json:"totalAmount" validate:"@int64[1,9007199254740991]"`
}
