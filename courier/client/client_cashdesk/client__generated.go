package client_cashdesk

import (
	"fmt"

	golib_tools_courier "golib/tools/courier"
	golib_tools_courier_enumeration "golib/tools/courier/enumeration"
	golib_tools_courier_status_error "golib/tools/courier/status_error"
	golib_tools_timelib "golib/tools/timelib"

	golib_tools_courier_client "golib/tools/courier/client"
)

type ClientCashdesk struct {
	golib_tools_courier_client.Client
}

func (ClientCashdesk) MarshalDefaults(v interface{}) {
	if cl, ok := v.(*ClientCashdesk); ok {
		cl.Name = "cashdesk"
		cl.Client.MarshalDefaults(&cl.Client)
	}
}

func (c ClientCashdesk) Init() {
	c.CheckService()
}

func (c ClientCashdesk) CheckService() {
	err := c.Request(c.Name+".Check", "HEAD", "/", nil).
		Do().
		Into(nil)
	statusErr := golib_tools_courier_status_error.FromError(err)
	if statusErr.Code == int64(golib_tools_courier_status_error.RequestTimeout) {
		panic(fmt.Errorf("service %s have some error %s", c.Name, statusErr))
	}
}

type CancelTransRequest struct {
	// 交易单id
	TransID uint64 `in:"path" name:"transID" validate:"@uint64[1,]"`
	// 外部用户id
	UserID string `in:"path" name:"userID" validate:"@string[1,64]"`
	//
	Body CancelTransBody `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) CancelTrans(req CancelTransRequest) (resp *CancelTransResponse, err error) {
	resp = &CancelTransResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".CancelTrans", "PUT", "/cashdesk/v0/trans/:transID/user/:userID/cancel", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type CancelTransResponse struct {
	Meta golib_tools_courier.Metadata
	Body string
}

type CreateTransRequest struct {
	// 商户的access key
	AccessKey string `in:"header" name:"AccessKey" validate:"@string[1,128]"`
	// 签名算法
	SignAlgorithm CashdeskSignAlgorithm `in:"query" name:"signAlgorithm"`
	// 签名
	Sign string `in:"query" name:"sign" validate:"@string[1,32]"`
	//
	Body CreateTransReqBody `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) CreateTrans(req CreateTransRequest) (resp *CreateTransResponse, err error) {
	resp = &CreateTransResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".CreateTrans", "POST", "/cashdesk/v0/trans", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type CreateTransResponse struct {
	Meta golib_tools_courier.Metadata
	Body CreateTransRespBody
}

type EBankPayRequest struct {
	// 交易单号
	TransID uint64 `in:"path" name:"transID" validate:"@uint64[0,]"`
	//
	Body EBankPayReqBody `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) EBankPay(req EBankPayRequest) (resp *EBankPayResponse, err error) {
	resp = &EBankPayResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".EBankPay", "POST", "/cashdesk/v0/trans/:transID/ebankpay", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type EBankPayResponse struct {
	Meta golib_tools_courier.Metadata
	Body string
}

type FasptaySignVerifyRequest struct {
	//
	Body FastpaySignVerifyReqBody `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) FasptaySignVerify(req FasptaySignVerifyRequest) (resp *FasptaySignVerifyResponse, err error) {
	resp = &FasptaySignVerifyResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".FasptaySignVerify", "POST", "/cashdesk/v0/contract/fastpay/verify", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type FasptaySignVerifyResponse struct {
	Meta golib_tools_courier.Metadata
	Body string
}

type FastpayMessageRequest struct {
	// 交易号
	TransID uint64 `default:"0" in:"path" name:"transID" validate:"@uint64[0,]"`
	//
	Body FastpayMessageReqBody `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) FastpayMessage(req FastpayMessageRequest) (resp *FastpayMessageResponse, err error) {
	resp = &FastpayMessageResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".FastpayMessage", "POST", "/cashdesk/v0/trans/:transID/fastpay/message", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type FastpayMessageResponse struct {
	Meta golib_tools_courier.Metadata
	Body FastpayMessageRespBody
}

type FastpayResignMessageRequest struct {
	//
	Body FastpayResignMessageReqBody `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) FastpayResignMessage(req FastpayResignMessageRequest) (resp *FastpayResignMessageResponse, err error) {
	resp = &FastpayResignMessageResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".FastpayResignMessage", "POST", "/cashdesk/v0/contract/fastpay/remessage", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type FastpayResignMessageResponse struct {
	Meta golib_tools_courier.Metadata
	Body string
}

type FastpaySignMessageRequest struct {
	//
	Body FastpaySignMessageReqBody `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) FastpaySignMessage(req FastpaySignMessageRequest) (resp *FastpaySignMessageResponse, err error) {
	resp = &FastpaySignMessageResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".FastpaySignMessage", "POST", "/cashdesk/v0/contract/fastpay/message", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type FastpaySignMessageResponse struct {
	Meta golib_tools_courier.Metadata
	Body FastpaySignMessageRespBody
}

type FastpayVerifyRequest struct {
	// 交易号
	TransID uint64 `default:"0" in:"path" name:"transID" validate:"@uint64[0,]"`
	//
	Body FastpayVerifyReqBody `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) FastpayVerify(req FastpayVerifyRequest) (resp *FastpayVerifyResponse, err error) {
	resp = &FastpayVerifyResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".FastpayVerify", "POST", "/cashdesk/v0/trans/:transID/fastpay/verify", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type FastpayVerifyResponse struct {
	Meta golib_tools_courier.Metadata
	Body FastpayVerifyRespBody
}

type GetAreaBankByBankCodeRequest struct {
	// 分支行行号
	BankCode string `in:"path" name:"bankCode" validate:"@string[1,]"`
}

func (c ClientCashdesk) GetAreaBankByBankCode(req GetAreaBankByBankCodeRequest) (resp *GetAreaBankByBankCodeResponse, err error) {
	resp = &GetAreaBankByBankCodeResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".GetAreaBankByBankCode", "GET", "/cashdesk/v0/bank/bankCode/:bankCode", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetAreaBankByBankCodeResponse struct {
	Meta golib_tools_courier.Metadata
	Body GetAreaBankByBankCodeRespBody
}

type GetAreaBankListRequest struct {
	// 银行ID
	BankID uint32 `in:"query" name:"bankID" validate:"@uint32[1,9999]"`
	// G7地区编码
	AreaCode string `in:"query" name:"areaCode" validate:"@string[1,]"`
}

func (c ClientCashdesk) GetAreaBankList(req GetAreaBankListRequest) (resp *GetAreaBankListResponse, err error) {
	resp = &GetAreaBankListResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".GetAreaBankList", "GET", "/cashdesk/v0/bank/area-bank", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetAreaBankListResponse struct {
	Meta golib_tools_courier.Metadata
	Body GetAreaBankListRespBody
}

type GetAreaListRequest struct {
	// 上级地区编码
	ParentCode string `default:"" in:"query" name:"parentCode" validate:"@string[0,10]"`
	// 深度(需要往下几层)
	Depth int32 `default:"1" in:"query" name:"depth" validate:"@int32[1,4]"`
	// 当前深度(上级地区编码处在第几层)
	DepthNow int32 `default:"0" in:"query" name:"depthNow" validate:"@int32[0,4]"`
}

func (c ClientCashdesk) GetAreaList(req GetAreaListRequest) (resp *GetAreaListResponse, err error) {
	resp = &GetAreaListResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".GetAreaList", "GET", "/cashdesk/v0/bank/area", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetAreaListResponse struct {
	Meta golib_tools_courier.Metadata
	Body GetAreaListByParentCodeDepthRespBody
}

type GetBankByIDRequest struct {
	// 银行id
	BankID uint32 `in:"path" name:"bankID" validate:"@uint32[1,]"`
}

func (c ClientCashdesk) GetBankByID(req GetBankByIDRequest) (resp *GetBankByIDResponse, err error) {
	resp = &GetBankByIDResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".GetBankByID", "GET", "/cashdesk/v0/bank/bankID/:bankID", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetBankByIDResponse struct {
	Meta golib_tools_courier.Metadata
	Body Bank
}

type GetBankListRequest struct {
	// 是否支持银联鉴权
	SupportUnionpayAuth golib_tools_courier_enumeration.Bool `default:"" in:"query" name:"supportUnionpayAuth" validate:"@string{,TRUE,FALSE}"`
	// 是否有超级网银联行号
	HaveSuperBankCode golib_tools_courier_enumeration.Bool `default:"" in:"query" name:"haveSuperBankCode" validate:"@string{,TRUE,FALSE}"`
}

func (c ClientCashdesk) GetBankList(req GetBankListRequest) (resp *GetBankListResponse, err error) {
	resp = &GetBankListResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".GetBankList", "GET", "/cashdesk/v0/bank", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetBankListResponse struct {
	Meta golib_tools_courier.Metadata
	Body []Bank
}

type GetPabWhiteListTransRequest struct {
	// 创建起始时间筛选 eg:2016-01-12T00:00:00Z
	CreateStartDate golib_tools_timelib.MySQLTimestamp `default:"" in:"query" name:"createStartDate"`
	// 创建结束时间筛选
	CreateEndDate golib_tools_timelib.MySQLTimestamp `default:"" in:"query" name:"createEndDate"`
	// 分页大小，默认为10，-1为查询所有
	Size int32 `default:"10" in:"query" name:"size" validate:"@int32[-1,100]"`
	// 分页偏移，默认为0
	Offset int32 `default:"0" in:"query" name:"offset" validate:"@int32[0,]"`
	// 买家账户ID
	BuyerAccountID uint64 `default:"0" in:"query" name:"buyerAccountID" validate:"@uint64[0,]"`
	// 卖家账户ID
	SellerAccountID uint64 `default:"0" in:"query" name:"sellerAccountID" validate:"@uint64[0,]"`
	// 外部交易单号
	TransIDExt string `default:"" in:"query" name:"transIDExt" validate:"@string[0,64]"`
	// 内部交易单号
	TransID uint64 `default:"0" in:"path" name:"transID" validate:"@uint64[0,]"`
}

func (c ClientCashdesk) GetPabWhiteListTrans(req GetPabWhiteListTransRequest) (resp *GetPabWhiteListTransResponse, err error) {
	resp = &GetPabWhiteListTransResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".GetPabWhiteListTrans", "GET", "/cashdesk/v0/trans/:transID/pabWhiteListTrans", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetPabWhiteListTransResponse struct {
	Meta golib_tools_courier.Metadata
	Body PabWhiteListTransResp
}

type GetSignRequest struct {
	// 签名密钥
	Secret string `in:"query" name:"secret" validate:"@string[1,128]"`
	// 签名算法
	SignAlgorithm CashdeskSignAlgorithm `in:"query" name:"signAlgorithm"`
	//
	Body []KVPair `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) GetSign(req GetSignRequest) (resp *GetSignResponse, err error) {
	resp = &GetSignResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".GetSign", "POST", "/cashdesk/v0/sign", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetSignResponse struct {
	Meta golib_tools_courier.Metadata
	Body GetSignRespBody
}

type GetTransByOrderNoRequest struct {
	// 订单号
	OrderNo string `in:"path" name:"orderNo" validate:"@string[6,32]"`
	// 商户AccessKey
	AccessKey string `in:"header" name:"accessKey" validate:"@string[1,128]"`
	// 是否启用Mock(仅测试环境生效)
	UseMock bool `default:"false" in:"query" name:"useMock"`
	// 模拟银行订单查询结果(仅UserMock为true生效)
	MockState CashdeskPayState `default:"" in:"query" name:"mockState"`
}

func (c ClientCashdesk) GetTransByOrderNo(req GetTransByOrderNoRequest) (resp *GetTransByOrderNoResponse, err error) {
	resp = &GetTransByOrderNoResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".GetTransByOrderNo", "GET", "/cashdesk/v0/trans/0/orderNo/:orderNo", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetTransByOrderNoResponse struct {
	Meta golib_tools_courier.Metadata
	Body ReturnTransModel
}

type OldTransQueryRequest struct {
	// 每个商户独有的AccessKey
	AccessKey string `in:"header" name:"accessKey" validate:"@string[1,128]"`
	// 交易单id
	TransID uint64 `in:"path" name:"transID" validate:"@uint64[1,]"`
	// 是否启用Mock(仅测试环境生效)
	UseMock bool `default:"false" in:"query" name:"useMock"`
	// 买家身份证(UseMock为true时有效)
	BuyerID string `default:"" in:"query" name:"buyerID"`
	// 期望返回的支付状态(UseMock为true时有效)
	PayStatus CashdeskPayState `default:"" in:"query" name:"payStatus"`
}

func (c ClientCashdesk) OldTransQuery(req OldTransQueryRequest) (resp *OldTransQueryResponse, err error) {
	resp = &OldTransQueryResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".OldTransQuery", "GET", "/cashdesk/v0/user/0/oldtrans/:transID", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type OldTransQueryResponse struct {
	Meta golib_tools_courier.Metadata
	Body OldTransQueryRespBody
}

type PabPayRequest struct {
	// 交易号
	TransID uint64 `default:"0" in:"path" name:"transID" validate:"@uint64[1,]"`
	//
	Body PabPayReqBody `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) PabPay(req PabPayRequest) (resp *PabPayResponse, err error) {
	resp = &PabPayResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".PabPay", "POST", "/cashdesk/v0/trans/:transID/pabpay", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type PabPayResponse struct {
	Meta golib_tools_courier.Metadata
	Body PabPayRespBody
}

type PabQueryByCertificateRequest struct {
	// 证件类型
	CertificateType CashdeskCertificateType `in:"query" name:"certificateType" validate:"@string{ID_CARD,ORG_INSITITUTE_CODE,UNITY_SOCIAL_CREDIT_CODE}"`
	// 证件号码
	CertificateCode string `in:"query" name:"certificateCode" validate:"@string[1,32]"`
}

func (c ClientCashdesk) PabQueryByCertificate(req PabQueryByCertificateRequest) (resp *PabQueryByCertificateResponse, err error) {
	resp = &PabQueryByCertificateResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".PabQueryByCertificate", "GET", "/cashdesk/v0/user/0/pab/certificate", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type PabQueryByCertificateResponse struct {
	Meta golib_tools_courier.Metadata
	Body PabLinkModel
}

type PabQueryByUserRequest struct {
	// 外部用户ID
	UserID string `in:"path" name:"userID" validate:"@string[1,32]"`
}

func (c ClientCashdesk) PabQueryByUser(req PabQueryByUserRequest) (resp *PabQueryByUserResponse, err error) {
	resp = &PabQueryByUserResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".PabQueryByUser", "GET", "/cashdesk/v0/user/:userID/pab", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type PabQueryByUserResponse struct {
	Meta golib_tools_courier.Metadata
	Body PabLinkModel
}

type PayRouteRequest struct {
	// 终端类型
	PlatformType CashdeskPlatformType `in:"query" name:"platformType"`
	// user-agent
	UserAgent string `in:"header" name:"User-Agent"`
	// 外部用户id
	UserID string `in:"query" name:"userID" validate:"@string[1,64]"`
	// 交易单id,若无,则为纯充值
	TransID uint64 `default:"0" in:"query" name:"transID" validate:"@uint64[0,]"`
	// 充值金额,若为0,则使用交易单金额
	Amount int64 `default:"0" in:"query" name:"amount" validate:"@int64[0,9007199254740991]"`
}

func (c ClientCashdesk) PayRoute(req PayRouteRequest) (resp *PayRouteResponse, err error) {
	resp = &PayRouteResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".PayRoute", "GET", "/cashdesk/v0/route/pay", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type PayRouteResponse struct {
	Meta golib_tools_courier.Metadata
	Body PayRouteRespBody
}

func (c ClientCashdesk) Swagger() (resp *SwaggerResponse, err error) {
	resp = &SwaggerResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".Swagger", "GET", "/cashdesk", nil).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type SwaggerResponse struct {
	Meta golib_tools_courier.Metadata
	Body string
}

type TransListQueryRequest struct {
	// 外部用户id
	UserID string `in:"path" name:"userID" validate:"@string[1,64]"`
	// 起始时间
	StartTime golib_tools_timelib.MySQLTimestamp `default:"" in:"query" name:"startTime"`
	// 终止时间
	EndTime golib_tools_timelib.MySQLTimestamp `default:"" in:"query" name:"endTime"`
	// 交易状态
	TransState CashdeskTransState `default:"" in:"query" name:"transState"`
	// 交易类型
	TransType CashdeskTransType `default:"" in:"query" name:"transType"`
	// 偏移,默认为0
	Offset int32 `default:"0" in:"query" name:"offset" validate:"@int32[0,]"`
	// 查询数量
	Size int32 `default:"10" in:"query" name:"size" validate:"@int32[1,50]"`
}

func (c ClientCashdesk) TransListQuery(req TransListQueryRequest) (resp *TransListQueryResponse, err error) {
	resp = &TransListQueryResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".TransListQuery", "GET", "/cashdesk/v0/user/:userID", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type TransListQueryResponse struct {
	Meta golib_tools_courier.Metadata
	Body QueryTransResponseBody
}

type UnbindFastpayRequest struct {
	//
	Body UnbindFastpayReqBody `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) UnbindFastpay(req UnbindFastpayRequest) (resp *UnbindFastpayResponse, err error) {
	resp = &UnbindFastpayResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".UnbindFastpay", "DELETE", "/cashdesk/v0/contract/fastpay", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type UnbindFastpayResponse struct {
	Meta golib_tools_courier.Metadata
	Body string
}

type UpdateTransRequest struct {
	// 交易单号
	TransID uint64 `in:"path" name:"transID" validate:"@uint64[1,]"`
	// 商户的access key
	AccessKey string `in:"header" name:"AccessKey" validate:"@string[1,128]"`
	// 签名算法
	SignAlgorithm CashdeskSignAlgorithm `in:"query" name:"signAlgorithm"`
	// 是否启用Mock(仅测试环境生效)
	UseMock bool `default:"false" in:"query" name:"useMock"`
	// 模拟银行订单查询结果(仅UserMock为true生效)
	MockState CashdeskPayState `default:"" in:"query" name:"mockState"`
	// 签名
	Sign string `in:"query" name:"sign" validate:"@string[1,32]"`
	//
	Body UpdateTransReqBody `default:"" in:"body" name:"body"`
}

func (c ClientCashdesk) UpdateTrans(req UpdateTransRequest) (resp *UpdateTransResponse, err error) {
	resp = &UpdateTransResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".UpdateTrans", "PUT", "/cashdesk/v0/trans/:transID", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type UpdateTransResponse struct {
	Meta golib_tools_courier.Metadata
	Body string
}

type UserCardQueryRequest struct {
	// 外部用户id
	UserID string `in:"path" name:"userID" validate:"@string[1,64]"`
}

func (c ClientCashdesk) UserCardQuery(req UserCardQueryRequest) (resp *UserCardQueryResponse, err error) {
	resp = &UserCardQueryResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".UserCardQuery", "GET", "/cashdesk/v0/user/:userID/card", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type UserCardQueryResponse struct {
	Meta golib_tools_courier.Metadata
	Body UserCardQueryRespBody
}
