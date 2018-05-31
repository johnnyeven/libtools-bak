package client_cashdesk

import (
	"bytes"
	"encoding"
	"errors"

	golib_tools_courier_enumeration "golib/tools/courier/enumeration"
)

// swagger:enum
type CashdeskTransState uint

const (
	CASHDESK_TRANS_STATE_UNKNOWN      CashdeskTransState = iota
	CASHDESK_TRANS_STATE__TO_PAY                         // 等待付款
	CASHDESK_TRANS_STATE__PAY_OK                         // 付款成功
	CASHDESK_TRANS_STATE__TRANS_OK                       // 交易成功
	CASHDESK_TRANS_STATE__REFUND_OK                      // 退款成功
	CASHDESK_TRANS_STATE__CANCEL                         // 交易取消
	CASHDESK_TRANS_STATE__STOP                           // 交易终止
	CASHDESK_TRANS_STATE__PART_REFUND                    // 部分退款
	CASHDESK_TRANS_STATE__PRE_PAY                        // 预支付完成
)

// swagger:enum
type CashdeskPabAccountType uint

const (
	CASHDESK_PAB_ACCOUNT_TYPE_UNKNOWN  CashdeskPabAccountType = iota
	CASHDESK_PAB_ACCOUNT_TYPE__COMPANY                        // 企业开户
	CASHDESK_PAB_ACCOUNT_TYPE__PERSON                         // 个人开户
)

// swagger:enum
type CashdeskPabAccountState uint

const (
	CASHDESK_PAB_ACCOUNT_STATE_UNKNOWN    CashdeskPabAccountState = iota
	CASHDESK_PAB_ACCOUNT_STATE__UNOPEN                            // 待开户
	CASHDESK_PAB_ACCOUNT_STATE__UNAUTH                            // 待验证
	CASHDESK_PAB_ACCOUNT_STATE__UNCONFIRM                         // 待确认
	CASHDESK_PAB_ACCOUNT_STATE__UNSIGN                            // 待签约
	CASHDESK_PAB_ACCOUNT_STATE__SIGNED                            // 已签约
	CASHDESK_PAB_ACCOUNT_STATE__UNCLOSE                           // 待销户
)

// swagger:enum
type CashdeskCardType uint

const (
	CASHDESK_CARD_TYPE_UNKNOWN CashdeskCardType = iota
	CASHDESK_CARD_TYPE__DEBIT                   // 借记卡
	CASHDESK_CARD_TYPE__CREDIT                  // 信用卡
)

// swagger:enum
type CashdeskPayMode uint

const (
	CASHDESK_PAY_MODE_UNKNOWN        CashdeskPayMode = iota
	CASHDESK_PAY_MODE__QUICK_PAY                     // 快捷支付
	CASHDESK_PAY_MODE__EBANK_PAY                     // 网银支付
	CASHDESK_PAY_MODE__WITHHOLD                      // 代扣
	CASHDESK_PAY_MODE__PAB                           // 平安易宝
	CASHDESK_PAY_MODE__WXPAY                         // 微信支付
	CASHDESK_PAY_MODE__ALIPAY                        // 支付宝
	CASHDESK_PAY_MODE__PAB_WITNESS                   // 平安见证宝
	CASHDESK_PAY_MODE__TRANSFER                      // 转账
	CASHDESK_PAY_MODE__WXPAY_QRCODE                  // 微信扫码支付
	CASHDESK_PAY_MODE__ALIPAY_QRCODE                 // 支付宝扫码支付
)

// swagger:enum
type CashdeskSettlementType uint

const (
	CASHDESK_SETTLEMENT_TYPE_UNKNOWN       CashdeskSettlementType = iota
	CASHDESK_SETTLEMENT_TYPE__REAL_TIME                           // 实时结算
	CASHDESK_SETTLEMENT_TYPE__ASYNCHRONOUS                        // 异步结算
)

// swagger:enum
type CashdeskTransMode uint

const (
	CASHDESK_TRANS_MODE_UNKNOWN            CashdeskTransMode = iota
	CASHDESK_TRANS_MODE__INTERMEDIARY                        // 中介交易
	CASHDESK_TRANS_MODE__DIRECT_PAY                          // 直付交易
	CASHDESK_TRANS_MODE__PREPAY                              // 预付交易
	CASHDESK_TRANS_MODE__INSURANCE_FINANCE                   // 保理交易
)

// swagger:enum
type CashdeskSignAlgorithm uint

const (
	CASHDESK_SIGN_ALGORITHM_UNKNOWN CashdeskSignAlgorithm = iota
	CASHDESK_SIGN_ALGORITHM__MD5                          // MD5
)

// swagger:enum
type CashdeskPayState uint

const (
	CASHDESK_PAY_STATE_UNKNOWN  CashdeskPayState = iota
	CASHDESK_PAY_STATE__UNPAY                    // 未支付
	CASHDESK_PAY_STATE__DEALING                  // 正在交易
	CASHDESK_PAY_STATE__SUCCESS                  // 支付成功
	CASHDESK_PAY_STATE__FAIL                     // 支付失败
)

// swagger:enum
type CashdeskUserFrom uint

const (
	CASHDESK_USER_FROM_UNKNOWN    CashdeskUserFrom = iota
	CASHDESK_USER_FROM__G7                         // G7平台
	CASHDESK_USER_FROM__ANONYMOUS                  // 匿名
	CASHDESK_USER_FROM__G7_ORG                     // G7机构
	CASHDESK_USER_FROM__WECHAT                     // 微信
)

// swagger:enum
type CashdeskAdjustFlag uint

const (
	CASHDESK_ADJUST_FLAG_UNKNOWN     CashdeskAdjustFlag = iota
	CASHDESK_ADJUST_FLAG__UNADJUSTED                    // 未调帐
	CASHDESK_ADJUST_FLAG__ADJUSTED                      // 已调帐
)

// swagger:enum
type CashdeskSubAccountType uint

const (
	CASHDESK_SUB_ACCOUNT_TYPE_UNKNOWN         CashdeskSubAccountType = iota
	CASHDESK_SUB_ACCOUNT_TYPE__CASH                                  // 现金子帐户
	CASHDESK_SUB_ACCOUNT_TYPE__PENDING_SETTLE                        // 待结算子帐户
	CASHDESK_SUB_ACCOUNT_TYPE__POINT                                 // 积分帐户
	CASHDESK_SUB_ACCOUNT_TYPE__DEPOSIT                               // 保证金帐户
	CASHDESK_SUB_ACCOUNT_TYPE__FROZEN                                // 冻结帐户
	CASHDESK_SUB_ACCOUNT_TYPE__PAB                                   // 平安现金子帐户
	CASHDESK_SUB_ACCOUNT_TYPE__OIL                                   // 油品子账户
	CASHDESK_SUB_ACCOUNT_TYPE__PAB_WITNESS                           // 见证宝现金子帐户
)

// swagger:enum
type CashdeskPayType uint

const (
	CASHDESK_PAY_TYPE_UNKNOWN   CashdeskPayType = iota
	CASHDESK_PAY_TYPE__BALANCE                  // 余额支付
	CASHDESK_PAY_TYPE__RECHARGE                 // 充值支付
)

// swagger:enum
type CashdeskCertificateType uint

const (
	CASHDESK_CERTIFICATE_TYPE_UNKNOWN            CashdeskCertificateType = iota
	CASHDESK_CERTIFICATE_TYPE__ID_CARD                                   // 身份证
	CASHDESK_CERTIFICATE_TYPE__PASSPORT                                  // 护照
	CASHDESK_CERTIFICATE_TYPE__MILITARY                                  // 军官证
	CASHDESK_CERTIFICATE_TYPE__BOOKLET                                   // 户口薄
	CASHDESK_CERTIFICATE_TYPE__SOLIDER                                   // 士兵证
	CASHDESK_CERTIFICATE_TYPE__HK_MACAO_PASS                             // 澳居台民来往内地通行证
	CASHDESK_CERTIFICATE_TYPE__TW_PASS                                   // 台湾居民来往内地通行证
	CASHDESK_CERTIFICATE_TYPE__TEMP_ID_CARD                              // 临时身份证
	CASHDESK_CERTIFICATE_TYPE__FOREIGN_RESIDENCE                         // 外国人居住证
	CASHDESK_CERTIFICATE_TYPE__POLICE_CARD                               // 警官证
	CASHDESK_CERTIFICATE_TYPE__PERSON_OTHER                              // 其他个人证件类型
)

const (
	CASHDESK_CERTIFICATE_TYPE__ORG_INSITITUTE_CODE      = iota + 100 // 组织机构代码证
	CASHDESK_CERTIFICATE_TYPE__BUSINESS_LICENSE                      // 营业执照
	CASHDESK_CERTIFICATE_TYPE__UNITY_SOCIAL_CREDIT_CODE              // 统一社会信用代码
	CASHDESK_CERTIFICATE_TYPE__LEGAL_PERSON_CODE                     // 法人代码证
	CASHDESK_CERTIFICATE_TYPE__UNIT_UNITY_CODE                       // 单位统一代码
	CASHDESK_CERTIFICATE_TYPE__FINANCIAL_ORG                         // 金融机构
	CASHDESK_CERTIFICATE_TYPE__COMPANY_OTHER                         // 其他公司证件类型
)

// swagger:enum
type CashdeskPlatformType uint

const (
	CASHDESK_PLATFORM_TYPE_UNKNOWN       CashdeskPlatformType = iota
	CASHDESK_PLATFORM_TYPE__COMPUTER                          // 个人电脑
	CASHDESK_PLATFORM_TYPE__MULTI_MEDIA                       // 多媒体终端
	CASHDESK_PLATFORM_TYPE__MOBILE_PHONE                      // 手持智能设备（手机、MID等）
	CASHDESK_PLATFORM_TYPE__PAD                               // 平板电脑
	CASHDESK_PLATFORM_TYPE__POS                               // POS终端
	CASHDESK_PLATFORM_TYPE__MERCHANT                          // 商户系统
	CASHDESK_PLATFORM_TYPE__STB                               // 数字机顶盒
	CASHDESK_PLATFORM_TYPE__TV                                // 智能电视
	CASHDESK_PLATFORM_TYPE__VEM                               // 自动柜员机（售货机等）
	CASHDESK_PLATFORM_TYPE__THIRD_SYSTEM                      // 第三方机构系统
)

// swagger:enum
type CashdeskTransType uint

const (
	CASHDESK_TRANS_TYPE_UNKNOWN       CashdeskTransType = iota
	CASHDESK_TRANS_TYPE__ETC                            // ETC
	CASHDESK_TRANS_TYPE__OIL_CARD                       // 油卡
	CASHDESK_TRANS_TYPE__TENDER                         // 招采
	CASHDESK_TRANS_TYPE__MALL                           // 商城
	CASHDESK_TRANS_TYPE__ZQX                            // 中启行(油)
	CASHDESK_TRANS_TYPE__SKID_MOUNTED                   // 撬装(油)
	CASHDESK_TRANS_TYPE__MAINTENANCE                    // 维保
)

// swagger:enum
type CashdeskSignType uint

const (
	CASHDESK_SIGN_TYPE_UNKNOWN    CashdeskSignType = iota
	CASHDESK_SIGN_TYPE__QUICK_PAY                  // 快捷支付
	CASHDESK_SIGN_TYPE__WITHHOLD                   // 代扣
)

// swagger:enum
type CashdeskBindState uint

const (
	CASHDESK_BIND_STATE_UNKNOWN    CashdeskBindState = iota
	CASHDESK_BIND_STATE__VALID                       // 生效
	CASHDESK_BIND_STATE__INVALID                     // 无效
	CASHDESK_BIND_STATE__TO_VERIFY                   // 待验证
	CASHDESK_BIND_STATE__VERIFYING                   // 验证中
	CASHDESK_BIND_STATE__VERIFYERR                   // 验证失败
)

// swagger:enum
type CashdeskCurrencyType uint

const (
	CASHDESK_CURRENCY_TYPE_UNKNOWN CashdeskCurrencyType = iota
	CASHDESK_CURRENCY_TYPE__RMB                         // 人民币
)

var InvalidCashdeskPayState = errors.New("invalid CashdeskPayState")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPayState", map[string]string{
		"UNPAY":   "未支付",
		"DEALING": "正在交易",
		"SUCCESS": "支付成功",
		"FAIL":    "支付失败",
	})
}

func ParseCashdeskPayStateFromString(s string) (CashdeskPayState, error) {
	switch s {
	case "":
		return CASHDESK_PAY_STATE_UNKNOWN, nil
	case "UNPAY":
		return CASHDESK_PAY_STATE__UNPAY, nil
	case "DEALING":
		return CASHDESK_PAY_STATE__DEALING, nil
	case "SUCCESS":
		return CASHDESK_PAY_STATE__SUCCESS, nil
	case "FAIL":
		return CASHDESK_PAY_STATE__FAIL, nil
	}
	return CASHDESK_PAY_STATE_UNKNOWN, InvalidCashdeskPayState
}

func ParseCashdeskPayStateFromLabelString(s string) (CashdeskPayState, error) {
	switch s {
	case "":
		return CASHDESK_PAY_STATE_UNKNOWN, nil
	case "未支付":
		return CASHDESK_PAY_STATE__UNPAY, nil
	case "正在交易":
		return CASHDESK_PAY_STATE__DEALING, nil
	case "支付成功":
		return CASHDESK_PAY_STATE__SUCCESS, nil
	case "支付失败":
		return CASHDESK_PAY_STATE__FAIL, nil
	}
	return CASHDESK_PAY_STATE_UNKNOWN, InvalidCashdeskPayState
}

func (v CashdeskPayState) String() string {
	switch v {
	case CASHDESK_PAY_STATE_UNKNOWN:
		return ""
	case CASHDESK_PAY_STATE__UNPAY:
		return "UNPAY"
	case CASHDESK_PAY_STATE__DEALING:
		return "DEALING"
	case CASHDESK_PAY_STATE__SUCCESS:
		return "SUCCESS"
	case CASHDESK_PAY_STATE__FAIL:
		return "FAIL"
	}
	return "UNKNOWN"
}

func (v CashdeskPayState) Label() string {
	switch v {
	case CASHDESK_PAY_STATE_UNKNOWN:
		return ""
	case CASHDESK_PAY_STATE__UNPAY:
		return "未支付"
	case CASHDESK_PAY_STATE__DEALING:
		return "正在交易"
	case CASHDESK_PAY_STATE__SUCCESS:
		return "支付成功"
	case CASHDESK_PAY_STATE__FAIL:
		return "支付失败"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPayState)(nil)

func (v CashdeskPayState) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPayState
	}
	return []byte(str), nil
}

func (v *CashdeskPayState) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPayStateFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskUserFrom = errors.New("invalid CashdeskUserFrom")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskUserFrom", map[string]string{
		"G7":        "G7平台",
		"ANONYMOUS": "匿名",
		"G7_ORG":    "G7机构",
		"WECHAT":    "微信",
	})
}

func ParseCashdeskUserFromFromString(s string) (CashdeskUserFrom, error) {
	switch s {
	case "":
		return CASHDESK_USER_FROM_UNKNOWN, nil
	case "G7":
		return CASHDESK_USER_FROM__G7, nil
	case "ANONYMOUS":
		return CASHDESK_USER_FROM__ANONYMOUS, nil
	case "G7_ORG":
		return CASHDESK_USER_FROM__G7_ORG, nil
	case "WECHAT":
		return CASHDESK_USER_FROM__WECHAT, nil
	}
	return CASHDESK_USER_FROM_UNKNOWN, InvalidCashdeskUserFrom
}

func ParseCashdeskUserFromFromLabelString(s string) (CashdeskUserFrom, error) {
	switch s {
	case "":
		return CASHDESK_USER_FROM_UNKNOWN, nil
	case "G7平台":
		return CASHDESK_USER_FROM__G7, nil
	case "匿名":
		return CASHDESK_USER_FROM__ANONYMOUS, nil
	case "G7机构":
		return CASHDESK_USER_FROM__G7_ORG, nil
	case "微信":
		return CASHDESK_USER_FROM__WECHAT, nil
	}
	return CASHDESK_USER_FROM_UNKNOWN, InvalidCashdeskUserFrom
}

func (v CashdeskUserFrom) String() string {
	switch v {
	case CASHDESK_USER_FROM_UNKNOWN:
		return ""
	case CASHDESK_USER_FROM__G7:
		return "G7"
	case CASHDESK_USER_FROM__ANONYMOUS:
		return "ANONYMOUS"
	case CASHDESK_USER_FROM__G7_ORG:
		return "G7_ORG"
	case CASHDESK_USER_FROM__WECHAT:
		return "WECHAT"
	}
	return "UNKNOWN"
}

func (v CashdeskUserFrom) Label() string {
	switch v {
	case CASHDESK_USER_FROM_UNKNOWN:
		return ""
	case CASHDESK_USER_FROM__G7:
		return "G7平台"
	case CASHDESK_USER_FROM__ANONYMOUS:
		return "匿名"
	case CASHDESK_USER_FROM__G7_ORG:
		return "G7机构"
	case CASHDESK_USER_FROM__WECHAT:
		return "微信"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskUserFrom)(nil)

func (v CashdeskUserFrom) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskUserFrom
	}
	return []byte(str), nil
}

func (v *CashdeskUserFrom) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskUserFromFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskAdjustFlag = errors.New("invalid CashdeskAdjustFlag")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskAdjustFlag", map[string]string{
		"UNADJUSTED": "未调帐",
		"ADJUSTED":   "已调帐",
	})
}

func ParseCashdeskAdjustFlagFromString(s string) (CashdeskAdjustFlag, error) {
	switch s {
	case "":
		return CASHDESK_ADJUST_FLAG_UNKNOWN, nil
	case "UNADJUSTED":
		return CASHDESK_ADJUST_FLAG__UNADJUSTED, nil
	case "ADJUSTED":
		return CASHDESK_ADJUST_FLAG__ADJUSTED, nil
	}
	return CASHDESK_ADJUST_FLAG_UNKNOWN, InvalidCashdeskAdjustFlag
}

func ParseCashdeskAdjustFlagFromLabelString(s string) (CashdeskAdjustFlag, error) {
	switch s {
	case "":
		return CASHDESK_ADJUST_FLAG_UNKNOWN, nil
	case "未调帐":
		return CASHDESK_ADJUST_FLAG__UNADJUSTED, nil
	case "已调帐":
		return CASHDESK_ADJUST_FLAG__ADJUSTED, nil
	}
	return CASHDESK_ADJUST_FLAG_UNKNOWN, InvalidCashdeskAdjustFlag
}

func (v CashdeskAdjustFlag) String() string {
	switch v {
	case CASHDESK_ADJUST_FLAG_UNKNOWN:
		return ""
	case CASHDESK_ADJUST_FLAG__UNADJUSTED:
		return "UNADJUSTED"
	case CASHDESK_ADJUST_FLAG__ADJUSTED:
		return "ADJUSTED"
	}
	return "UNKNOWN"
}

func (v CashdeskAdjustFlag) Label() string {
	switch v {
	case CASHDESK_ADJUST_FLAG_UNKNOWN:
		return ""
	case CASHDESK_ADJUST_FLAG__UNADJUSTED:
		return "未调帐"
	case CASHDESK_ADJUST_FLAG__ADJUSTED:
		return "已调帐"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskAdjustFlag)(nil)

func (v CashdeskAdjustFlag) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskAdjustFlag
	}
	return []byte(str), nil
}

func (v *CashdeskAdjustFlag) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskAdjustFlagFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskSubAccountType = errors.New("invalid CashdeskSubAccountType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskSubAccountType", map[string]string{
		"CASH":           "现金子帐户",
		"PENDING_SETTLE": "待结算子帐户",
		"POINT":          "积分帐户",
		"DEPOSIT":        "保证金帐户",
		"FROZEN":         "冻结帐户",
		"PAB":            "平安现金子帐户",
		"OIL":            "油品子账户",
		"PAB_WITNESS":    "见证宝现金子帐户",
	})
}

func ParseCashdeskSubAccountTypeFromString(s string) (CashdeskSubAccountType, error) {
	switch s {
	case "":
		return CASHDESK_SUB_ACCOUNT_TYPE_UNKNOWN, nil
	case "CASH":
		return CASHDESK_SUB_ACCOUNT_TYPE__CASH, nil
	case "PENDING_SETTLE":
		return CASHDESK_SUB_ACCOUNT_TYPE__PENDING_SETTLE, nil
	case "POINT":
		return CASHDESK_SUB_ACCOUNT_TYPE__POINT, nil
	case "DEPOSIT":
		return CASHDESK_SUB_ACCOUNT_TYPE__DEPOSIT, nil
	case "FROZEN":
		return CASHDESK_SUB_ACCOUNT_TYPE__FROZEN, nil
	case "PAB":
		return CASHDESK_SUB_ACCOUNT_TYPE__PAB, nil
	case "OIL":
		return CASHDESK_SUB_ACCOUNT_TYPE__OIL, nil
	case "PAB_WITNESS":
		return CASHDESK_SUB_ACCOUNT_TYPE__PAB_WITNESS, nil
	}
	return CASHDESK_SUB_ACCOUNT_TYPE_UNKNOWN, InvalidCashdeskSubAccountType
}

func ParseCashdeskSubAccountTypeFromLabelString(s string) (CashdeskSubAccountType, error) {
	switch s {
	case "":
		return CASHDESK_SUB_ACCOUNT_TYPE_UNKNOWN, nil
	case "现金子帐户":
		return CASHDESK_SUB_ACCOUNT_TYPE__CASH, nil
	case "待结算子帐户":
		return CASHDESK_SUB_ACCOUNT_TYPE__PENDING_SETTLE, nil
	case "积分帐户":
		return CASHDESK_SUB_ACCOUNT_TYPE__POINT, nil
	case "保证金帐户":
		return CASHDESK_SUB_ACCOUNT_TYPE__DEPOSIT, nil
	case "冻结帐户":
		return CASHDESK_SUB_ACCOUNT_TYPE__FROZEN, nil
	case "平安现金子帐户":
		return CASHDESK_SUB_ACCOUNT_TYPE__PAB, nil
	case "油品子账户":
		return CASHDESK_SUB_ACCOUNT_TYPE__OIL, nil
	case "见证宝现金子帐户":
		return CASHDESK_SUB_ACCOUNT_TYPE__PAB_WITNESS, nil
	}
	return CASHDESK_SUB_ACCOUNT_TYPE_UNKNOWN, InvalidCashdeskSubAccountType
}

func (v CashdeskSubAccountType) String() string {
	switch v {
	case CASHDESK_SUB_ACCOUNT_TYPE_UNKNOWN:
		return ""
	case CASHDESK_SUB_ACCOUNT_TYPE__CASH:
		return "CASH"
	case CASHDESK_SUB_ACCOUNT_TYPE__PENDING_SETTLE:
		return "PENDING_SETTLE"
	case CASHDESK_SUB_ACCOUNT_TYPE__POINT:
		return "POINT"
	case CASHDESK_SUB_ACCOUNT_TYPE__DEPOSIT:
		return "DEPOSIT"
	case CASHDESK_SUB_ACCOUNT_TYPE__FROZEN:
		return "FROZEN"
	case CASHDESK_SUB_ACCOUNT_TYPE__PAB:
		return "PAB"
	case CASHDESK_SUB_ACCOUNT_TYPE__OIL:
		return "OIL"
	case CASHDESK_SUB_ACCOUNT_TYPE__PAB_WITNESS:
		return "PAB_WITNESS"
	}
	return "UNKNOWN"
}

func (v CashdeskSubAccountType) Label() string {
	switch v {
	case CASHDESK_SUB_ACCOUNT_TYPE_UNKNOWN:
		return ""
	case CASHDESK_SUB_ACCOUNT_TYPE__CASH:
		return "现金子帐户"
	case CASHDESK_SUB_ACCOUNT_TYPE__PENDING_SETTLE:
		return "待结算子帐户"
	case CASHDESK_SUB_ACCOUNT_TYPE__POINT:
		return "积分帐户"
	case CASHDESK_SUB_ACCOUNT_TYPE__DEPOSIT:
		return "保证金帐户"
	case CASHDESK_SUB_ACCOUNT_TYPE__FROZEN:
		return "冻结帐户"
	case CASHDESK_SUB_ACCOUNT_TYPE__PAB:
		return "平安现金子帐户"
	case CASHDESK_SUB_ACCOUNT_TYPE__OIL:
		return "油品子账户"
	case CASHDESK_SUB_ACCOUNT_TYPE__PAB_WITNESS:
		return "见证宝现金子帐户"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskSubAccountType)(nil)

func (v CashdeskSubAccountType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskSubAccountType
	}
	return []byte(str), nil
}

func (v *CashdeskSubAccountType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskSubAccountTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskSignAlgorithm = errors.New("invalid CashdeskSignAlgorithm")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskSignAlgorithm", map[string]string{
		"MD5": "MD5",
	})
}

func ParseCashdeskSignAlgorithmFromString(s string) (CashdeskSignAlgorithm, error) {
	switch s {
	case "":
		return CASHDESK_SIGN_ALGORITHM_UNKNOWN, nil
	case "MD5":
		return CASHDESK_SIGN_ALGORITHM__MD5, nil
	}
	return CASHDESK_SIGN_ALGORITHM_UNKNOWN, InvalidCashdeskSignAlgorithm
}

func ParseCashdeskSignAlgorithmFromLabelString(s string) (CashdeskSignAlgorithm, error) {
	switch s {
	case "":
		return CASHDESK_SIGN_ALGORITHM_UNKNOWN, nil
	case "MD5":
		return CASHDESK_SIGN_ALGORITHM__MD5, nil
	}
	return CASHDESK_SIGN_ALGORITHM_UNKNOWN, InvalidCashdeskSignAlgorithm
}

func (v CashdeskSignAlgorithm) String() string {
	switch v {
	case CASHDESK_SIGN_ALGORITHM_UNKNOWN:
		return ""
	case CASHDESK_SIGN_ALGORITHM__MD5:
		return "MD5"
	}
	return "UNKNOWN"
}

func (v CashdeskSignAlgorithm) Label() string {
	switch v {
	case CASHDESK_SIGN_ALGORITHM_UNKNOWN:
		return ""
	case CASHDESK_SIGN_ALGORITHM__MD5:
		return "MD5"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskSignAlgorithm)(nil)

func (v CashdeskSignAlgorithm) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskSignAlgorithm
	}
	return []byte(str), nil
}

func (v *CashdeskSignAlgorithm) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskSignAlgorithmFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPlatformType = errors.New("invalid CashdeskPlatformType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPlatformType", map[string]string{
		"COMPUTER":     "个人电脑",
		"MULTI_MEDIA":  "多媒体终端",
		"MOBILE_PHONE": "手持智能设备（手机、MID等）",
		"PAD":          "平板电脑",
		"POS":          "POS终端",
		"MERCHANT":     "商户系统",
		"STB":          "数字机顶盒",
		"TV":           "智能电视",
		"VEM":          "自动柜员机（售货机等）",
		"THIRD_SYSTEM": "第三方机构系统",
	})
}

func ParseCashdeskPlatformTypeFromString(s string) (CashdeskPlatformType, error) {
	switch s {
	case "":
		return CASHDESK_PLATFORM_TYPE_UNKNOWN, nil
	case "COMPUTER":
		return CASHDESK_PLATFORM_TYPE__COMPUTER, nil
	case "MULTI_MEDIA":
		return CASHDESK_PLATFORM_TYPE__MULTI_MEDIA, nil
	case "MOBILE_PHONE":
		return CASHDESK_PLATFORM_TYPE__MOBILE_PHONE, nil
	case "PAD":
		return CASHDESK_PLATFORM_TYPE__PAD, nil
	case "POS":
		return CASHDESK_PLATFORM_TYPE__POS, nil
	case "MERCHANT":
		return CASHDESK_PLATFORM_TYPE__MERCHANT, nil
	case "STB":
		return CASHDESK_PLATFORM_TYPE__STB, nil
	case "TV":
		return CASHDESK_PLATFORM_TYPE__TV, nil
	case "VEM":
		return CASHDESK_PLATFORM_TYPE__VEM, nil
	case "THIRD_SYSTEM":
		return CASHDESK_PLATFORM_TYPE__THIRD_SYSTEM, nil
	}
	return CASHDESK_PLATFORM_TYPE_UNKNOWN, InvalidCashdeskPlatformType
}

func ParseCashdeskPlatformTypeFromLabelString(s string) (CashdeskPlatformType, error) {
	switch s {
	case "":
		return CASHDESK_PLATFORM_TYPE_UNKNOWN, nil
	case "个人电脑":
		return CASHDESK_PLATFORM_TYPE__COMPUTER, nil
	case "多媒体终端":
		return CASHDESK_PLATFORM_TYPE__MULTI_MEDIA, nil
	case "手持智能设备（手机、MID等）":
		return CASHDESK_PLATFORM_TYPE__MOBILE_PHONE, nil
	case "平板电脑":
		return CASHDESK_PLATFORM_TYPE__PAD, nil
	case "POS终端":
		return CASHDESK_PLATFORM_TYPE__POS, nil
	case "商户系统":
		return CASHDESK_PLATFORM_TYPE__MERCHANT, nil
	case "数字机顶盒":
		return CASHDESK_PLATFORM_TYPE__STB, nil
	case "智能电视":
		return CASHDESK_PLATFORM_TYPE__TV, nil
	case "自动柜员机（售货机等）":
		return CASHDESK_PLATFORM_TYPE__VEM, nil
	case "第三方机构系统":
		return CASHDESK_PLATFORM_TYPE__THIRD_SYSTEM, nil
	}
	return CASHDESK_PLATFORM_TYPE_UNKNOWN, InvalidCashdeskPlatformType
}

func (v CashdeskPlatformType) String() string {
	switch v {
	case CASHDESK_PLATFORM_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PLATFORM_TYPE__COMPUTER:
		return "COMPUTER"
	case CASHDESK_PLATFORM_TYPE__MULTI_MEDIA:
		return "MULTI_MEDIA"
	case CASHDESK_PLATFORM_TYPE__MOBILE_PHONE:
		return "MOBILE_PHONE"
	case CASHDESK_PLATFORM_TYPE__PAD:
		return "PAD"
	case CASHDESK_PLATFORM_TYPE__POS:
		return "POS"
	case CASHDESK_PLATFORM_TYPE__MERCHANT:
		return "MERCHANT"
	case CASHDESK_PLATFORM_TYPE__STB:
		return "STB"
	case CASHDESK_PLATFORM_TYPE__TV:
		return "TV"
	case CASHDESK_PLATFORM_TYPE__VEM:
		return "VEM"
	case CASHDESK_PLATFORM_TYPE__THIRD_SYSTEM:
		return "THIRD_SYSTEM"
	}
	return "UNKNOWN"
}

func (v CashdeskPlatformType) Label() string {
	switch v {
	case CASHDESK_PLATFORM_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PLATFORM_TYPE__COMPUTER:
		return "个人电脑"
	case CASHDESK_PLATFORM_TYPE__MULTI_MEDIA:
		return "多媒体终端"
	case CASHDESK_PLATFORM_TYPE__MOBILE_PHONE:
		return "手持智能设备（手机、MID等）"
	case CASHDESK_PLATFORM_TYPE__PAD:
		return "平板电脑"
	case CASHDESK_PLATFORM_TYPE__POS:
		return "POS终端"
	case CASHDESK_PLATFORM_TYPE__MERCHANT:
		return "商户系统"
	case CASHDESK_PLATFORM_TYPE__STB:
		return "数字机顶盒"
	case CASHDESK_PLATFORM_TYPE__TV:
		return "智能电视"
	case CASHDESK_PLATFORM_TYPE__VEM:
		return "自动柜员机（售货机等）"
	case CASHDESK_PLATFORM_TYPE__THIRD_SYSTEM:
		return "第三方机构系统"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPlatformType)(nil)

func (v CashdeskPlatformType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPlatformType
	}
	return []byte(str), nil
}

func (v *CashdeskPlatformType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPlatformTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskTransType = errors.New("invalid CashdeskTransType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskTransType", map[string]string{
		"ETC":          "ETC",
		"OIL_CARD":     "油卡",
		"TENDER":       "招采",
		"MALL":         "商城",
		"ZQX":          "中启行(油)",
		"SKID_MOUNTED": "撬装(油)",
		"MAINTENANCE":  "维保",
	})
}

func ParseCashdeskTransTypeFromString(s string) (CashdeskTransType, error) {
	switch s {
	case "":
		return CASHDESK_TRANS_TYPE_UNKNOWN, nil
	case "ETC":
		return CASHDESK_TRANS_TYPE__ETC, nil
	case "OIL_CARD":
		return CASHDESK_TRANS_TYPE__OIL_CARD, nil
	case "TENDER":
		return CASHDESK_TRANS_TYPE__TENDER, nil
	case "MALL":
		return CASHDESK_TRANS_TYPE__MALL, nil
	case "ZQX":
		return CASHDESK_TRANS_TYPE__ZQX, nil
	case "SKID_MOUNTED":
		return CASHDESK_TRANS_TYPE__SKID_MOUNTED, nil
	case "MAINTENANCE":
		return CASHDESK_TRANS_TYPE__MAINTENANCE, nil
	}
	return CASHDESK_TRANS_TYPE_UNKNOWN, InvalidCashdeskTransType
}

func ParseCashdeskTransTypeFromLabelString(s string) (CashdeskTransType, error) {
	switch s {
	case "":
		return CASHDESK_TRANS_TYPE_UNKNOWN, nil
	case "ETC":
		return CASHDESK_TRANS_TYPE__ETC, nil
	case "油卡":
		return CASHDESK_TRANS_TYPE__OIL_CARD, nil
	case "招采":
		return CASHDESK_TRANS_TYPE__TENDER, nil
	case "商城":
		return CASHDESK_TRANS_TYPE__MALL, nil
	case "中启行(油)":
		return CASHDESK_TRANS_TYPE__ZQX, nil
	case "撬装(油)":
		return CASHDESK_TRANS_TYPE__SKID_MOUNTED, nil
	case "维保":
		return CASHDESK_TRANS_TYPE__MAINTENANCE, nil
	}
	return CASHDESK_TRANS_TYPE_UNKNOWN, InvalidCashdeskTransType
}

func (v CashdeskTransType) String() string {
	switch v {
	case CASHDESK_TRANS_TYPE_UNKNOWN:
		return ""
	case CASHDESK_TRANS_TYPE__ETC:
		return "ETC"
	case CASHDESK_TRANS_TYPE__OIL_CARD:
		return "OIL_CARD"
	case CASHDESK_TRANS_TYPE__TENDER:
		return "TENDER"
	case CASHDESK_TRANS_TYPE__MALL:
		return "MALL"
	case CASHDESK_TRANS_TYPE__ZQX:
		return "ZQX"
	case CASHDESK_TRANS_TYPE__SKID_MOUNTED:
		return "SKID_MOUNTED"
	case CASHDESK_TRANS_TYPE__MAINTENANCE:
		return "MAINTENANCE"
	}
	return "UNKNOWN"
}

func (v CashdeskTransType) Label() string {
	switch v {
	case CASHDESK_TRANS_TYPE_UNKNOWN:
		return ""
	case CASHDESK_TRANS_TYPE__ETC:
		return "ETC"
	case CASHDESK_TRANS_TYPE__OIL_CARD:
		return "油卡"
	case CASHDESK_TRANS_TYPE__TENDER:
		return "招采"
	case CASHDESK_TRANS_TYPE__MALL:
		return "商城"
	case CASHDESK_TRANS_TYPE__ZQX:
		return "中启行(油)"
	case CASHDESK_TRANS_TYPE__SKID_MOUNTED:
		return "撬装(油)"
	case CASHDESK_TRANS_TYPE__MAINTENANCE:
		return "维保"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskTransType)(nil)

func (v CashdeskTransType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskTransType
	}
	return []byte(str), nil
}

func (v *CashdeskTransType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskTransTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskSignType = errors.New("invalid CashdeskSignType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskSignType", map[string]string{
		"QUICK_PAY": "快捷支付",
		"WITHHOLD":  "代扣",
	})
}

func ParseCashdeskSignTypeFromString(s string) (CashdeskSignType, error) {
	switch s {
	case "":
		return CASHDESK_SIGN_TYPE_UNKNOWN, nil
	case "QUICK_PAY":
		return CASHDESK_SIGN_TYPE__QUICK_PAY, nil
	case "WITHHOLD":
		return CASHDESK_SIGN_TYPE__WITHHOLD, nil
	}
	return CASHDESK_SIGN_TYPE_UNKNOWN, InvalidCashdeskSignType
}

func ParseCashdeskSignTypeFromLabelString(s string) (CashdeskSignType, error) {
	switch s {
	case "":
		return CASHDESK_SIGN_TYPE_UNKNOWN, nil
	case "快捷支付":
		return CASHDESK_SIGN_TYPE__QUICK_PAY, nil
	case "代扣":
		return CASHDESK_SIGN_TYPE__WITHHOLD, nil
	}
	return CASHDESK_SIGN_TYPE_UNKNOWN, InvalidCashdeskSignType
}

func (v CashdeskSignType) String() string {
	switch v {
	case CASHDESK_SIGN_TYPE_UNKNOWN:
		return ""
	case CASHDESK_SIGN_TYPE__QUICK_PAY:
		return "QUICK_PAY"
	case CASHDESK_SIGN_TYPE__WITHHOLD:
		return "WITHHOLD"
	}
	return "UNKNOWN"
}

func (v CashdeskSignType) Label() string {
	switch v {
	case CASHDESK_SIGN_TYPE_UNKNOWN:
		return ""
	case CASHDESK_SIGN_TYPE__QUICK_PAY:
		return "快捷支付"
	case CASHDESK_SIGN_TYPE__WITHHOLD:
		return "代扣"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskSignType)(nil)

func (v CashdeskSignType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskSignType
	}
	return []byte(str), nil
}

func (v *CashdeskSignType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskSignTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskBindState = errors.New("invalid CashdeskBindState")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskBindState", map[string]string{
		"VALID":     "生效",
		"INVALID":   "无效",
		"TO_VERIFY": "待验证",
		"VERIFYING": "验证中",
		"VERIFYERR": "验证失败",
	})
}

func ParseCashdeskBindStateFromString(s string) (CashdeskBindState, error) {
	switch s {
	case "":
		return CASHDESK_BIND_STATE_UNKNOWN, nil
	case "VALID":
		return CASHDESK_BIND_STATE__VALID, nil
	case "INVALID":
		return CASHDESK_BIND_STATE__INVALID, nil
	case "TO_VERIFY":
		return CASHDESK_BIND_STATE__TO_VERIFY, nil
	case "VERIFYING":
		return CASHDESK_BIND_STATE__VERIFYING, nil
	case "VERIFYERR":
		return CASHDESK_BIND_STATE__VERIFYERR, nil
	}
	return CASHDESK_BIND_STATE_UNKNOWN, InvalidCashdeskBindState
}

func ParseCashdeskBindStateFromLabelString(s string) (CashdeskBindState, error) {
	switch s {
	case "":
		return CASHDESK_BIND_STATE_UNKNOWN, nil
	case "生效":
		return CASHDESK_BIND_STATE__VALID, nil
	case "无效":
		return CASHDESK_BIND_STATE__INVALID, nil
	case "待验证":
		return CASHDESK_BIND_STATE__TO_VERIFY, nil
	case "验证中":
		return CASHDESK_BIND_STATE__VERIFYING, nil
	case "验证失败":
		return CASHDESK_BIND_STATE__VERIFYERR, nil
	}
	return CASHDESK_BIND_STATE_UNKNOWN, InvalidCashdeskBindState
}

func (v CashdeskBindState) String() string {
	switch v {
	case CASHDESK_BIND_STATE_UNKNOWN:
		return ""
	case CASHDESK_BIND_STATE__VALID:
		return "VALID"
	case CASHDESK_BIND_STATE__INVALID:
		return "INVALID"
	case CASHDESK_BIND_STATE__TO_VERIFY:
		return "TO_VERIFY"
	case CASHDESK_BIND_STATE__VERIFYING:
		return "VERIFYING"
	case CASHDESK_BIND_STATE__VERIFYERR:
		return "VERIFYERR"
	}
	return "UNKNOWN"
}

func (v CashdeskBindState) Label() string {
	switch v {
	case CASHDESK_BIND_STATE_UNKNOWN:
		return ""
	case CASHDESK_BIND_STATE__VALID:
		return "生效"
	case CASHDESK_BIND_STATE__INVALID:
		return "无效"
	case CASHDESK_BIND_STATE__TO_VERIFY:
		return "待验证"
	case CASHDESK_BIND_STATE__VERIFYING:
		return "验证中"
	case CASHDESK_BIND_STATE__VERIFYERR:
		return "验证失败"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskBindState)(nil)

func (v CashdeskBindState) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskBindState
	}
	return []byte(str), nil
}

func (v *CashdeskBindState) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskBindStateFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskCurrencyType = errors.New("invalid CashdeskCurrencyType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskCurrencyType", map[string]string{
		"RMB": "人民币",
	})
}

func ParseCashdeskCurrencyTypeFromString(s string) (CashdeskCurrencyType, error) {
	switch s {
	case "":
		return CASHDESK_CURRENCY_TYPE_UNKNOWN, nil
	case "RMB":
		return CASHDESK_CURRENCY_TYPE__RMB, nil
	}
	return CASHDESK_CURRENCY_TYPE_UNKNOWN, InvalidCashdeskCurrencyType
}

func ParseCashdeskCurrencyTypeFromLabelString(s string) (CashdeskCurrencyType, error) {
	switch s {
	case "":
		return CASHDESK_CURRENCY_TYPE_UNKNOWN, nil
	case "人民币":
		return CASHDESK_CURRENCY_TYPE__RMB, nil
	}
	return CASHDESK_CURRENCY_TYPE_UNKNOWN, InvalidCashdeskCurrencyType
}

func (v CashdeskCurrencyType) String() string {
	switch v {
	case CASHDESK_CURRENCY_TYPE_UNKNOWN:
		return ""
	case CASHDESK_CURRENCY_TYPE__RMB:
		return "RMB"
	}
	return "UNKNOWN"
}

func (v CashdeskCurrencyType) Label() string {
	switch v {
	case CASHDESK_CURRENCY_TYPE_UNKNOWN:
		return ""
	case CASHDESK_CURRENCY_TYPE__RMB:
		return "人民币"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskCurrencyType)(nil)

func (v CashdeskCurrencyType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskCurrencyType
	}
	return []byte(str), nil
}

func (v *CashdeskCurrencyType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskCurrencyTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPayType = errors.New("invalid CashdeskPayType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPayType", map[string]string{
		"BALANCE":  "余额支付",
		"RECHARGE": "充值支付",
	})
}

func ParseCashdeskPayTypeFromString(s string) (CashdeskPayType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_TYPE_UNKNOWN, nil
	case "BALANCE":
		return CASHDESK_PAY_TYPE__BALANCE, nil
	case "RECHARGE":
		return CASHDESK_PAY_TYPE__RECHARGE, nil
	}
	return CASHDESK_PAY_TYPE_UNKNOWN, InvalidCashdeskPayType
}

func ParseCashdeskPayTypeFromLabelString(s string) (CashdeskPayType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_TYPE_UNKNOWN, nil
	case "余额支付":
		return CASHDESK_PAY_TYPE__BALANCE, nil
	case "充值支付":
		return CASHDESK_PAY_TYPE__RECHARGE, nil
	}
	return CASHDESK_PAY_TYPE_UNKNOWN, InvalidCashdeskPayType
}

func (v CashdeskPayType) String() string {
	switch v {
	case CASHDESK_PAY_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_TYPE__BALANCE:
		return "BALANCE"
	case CASHDESK_PAY_TYPE__RECHARGE:
		return "RECHARGE"
	}
	return "UNKNOWN"
}

func (v CashdeskPayType) Label() string {
	switch v {
	case CASHDESK_PAY_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_TYPE__BALANCE:
		return "余额支付"
	case CASHDESK_PAY_TYPE__RECHARGE:
		return "充值支付"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPayType)(nil)

func (v CashdeskPayType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPayType
	}
	return []byte(str), nil
}

func (v *CashdeskPayType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPayTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskCertificateType = errors.New("invalid CashdeskCertificateType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskCertificateType", map[string]string{
		"ID_CARD":                  "身份证",
		"PASSPORT":                 "护照",
		"MILITARY":                 "军官证",
		"BOOKLET":                  "户口薄",
		"SOLIDER":                  "士兵证",
		"HK_MACAO_PASS":            "澳居台民来往内地通行证",
		"TW_PASS":                  "台湾居民来往内地通行证",
		"TEMP_ID_CARD":             "临时身份证",
		"FOREIGN_RESIDENCE":        "外国人居住证",
		"POLICE_CARD":              "警官证",
		"PERSON_OTHER":             "其他个人证件类型",
		"ORG_INSITITUTE_CODE":      "组织机构代码证",
		"BUSINESS_LICENSE":         "营业执照",
		"UNITY_SOCIAL_CREDIT_CODE": "统一社会信用代码",
		"LEGAL_PERSON_CODE":        "法人代码证",
		"UNIT_UNITY_CODE":          "单位统一代码",
		"FINANCIAL_ORG":            "金融机构",
		"COMPANY_OTHER":            "其他公司证件类型",
	})
}

func ParseCashdeskCertificateTypeFromString(s string) (CashdeskCertificateType, error) {
	switch s {
	case "":
		return CASHDESK_CERTIFICATE_TYPE_UNKNOWN, nil
	case "ID_CARD":
		return CASHDESK_CERTIFICATE_TYPE__ID_CARD, nil
	case "PASSPORT":
		return CASHDESK_CERTIFICATE_TYPE__PASSPORT, nil
	case "MILITARY":
		return CASHDESK_CERTIFICATE_TYPE__MILITARY, nil
	case "BOOKLET":
		return CASHDESK_CERTIFICATE_TYPE__BOOKLET, nil
	case "SOLIDER":
		return CASHDESK_CERTIFICATE_TYPE__SOLIDER, nil
	case "HK_MACAO_PASS":
		return CASHDESK_CERTIFICATE_TYPE__HK_MACAO_PASS, nil
	case "TW_PASS":
		return CASHDESK_CERTIFICATE_TYPE__TW_PASS, nil
	case "TEMP_ID_CARD":
		return CASHDESK_CERTIFICATE_TYPE__TEMP_ID_CARD, nil
	case "FOREIGN_RESIDENCE":
		return CASHDESK_CERTIFICATE_TYPE__FOREIGN_RESIDENCE, nil
	case "POLICE_CARD":
		return CASHDESK_CERTIFICATE_TYPE__POLICE_CARD, nil
	case "PERSON_OTHER":
		return CASHDESK_CERTIFICATE_TYPE__PERSON_OTHER, nil
	case "ORG_INSITITUTE_CODE":
		return CASHDESK_CERTIFICATE_TYPE__ORG_INSITITUTE_CODE, nil
	case "BUSINESS_LICENSE":
		return CASHDESK_CERTIFICATE_TYPE__BUSINESS_LICENSE, nil
	case "UNITY_SOCIAL_CREDIT_CODE":
		return CASHDESK_CERTIFICATE_TYPE__UNITY_SOCIAL_CREDIT_CODE, nil
	case "LEGAL_PERSON_CODE":
		return CASHDESK_CERTIFICATE_TYPE__LEGAL_PERSON_CODE, nil
	case "UNIT_UNITY_CODE":
		return CASHDESK_CERTIFICATE_TYPE__UNIT_UNITY_CODE, nil
	case "FINANCIAL_ORG":
		return CASHDESK_CERTIFICATE_TYPE__FINANCIAL_ORG, nil
	case "COMPANY_OTHER":
		return CASHDESK_CERTIFICATE_TYPE__COMPANY_OTHER, nil
	}
	return CASHDESK_CERTIFICATE_TYPE_UNKNOWN, InvalidCashdeskCertificateType
}

func ParseCashdeskCertificateTypeFromLabelString(s string) (CashdeskCertificateType, error) {
	switch s {
	case "":
		return CASHDESK_CERTIFICATE_TYPE_UNKNOWN, nil
	case "身份证":
		return CASHDESK_CERTIFICATE_TYPE__ID_CARD, nil
	case "护照":
		return CASHDESK_CERTIFICATE_TYPE__PASSPORT, nil
	case "军官证":
		return CASHDESK_CERTIFICATE_TYPE__MILITARY, nil
	case "户口薄":
		return CASHDESK_CERTIFICATE_TYPE__BOOKLET, nil
	case "士兵证":
		return CASHDESK_CERTIFICATE_TYPE__SOLIDER, nil
	case "澳居台民来往内地通行证":
		return CASHDESK_CERTIFICATE_TYPE__HK_MACAO_PASS, nil
	case "台湾居民来往内地通行证":
		return CASHDESK_CERTIFICATE_TYPE__TW_PASS, nil
	case "临时身份证":
		return CASHDESK_CERTIFICATE_TYPE__TEMP_ID_CARD, nil
	case "外国人居住证":
		return CASHDESK_CERTIFICATE_TYPE__FOREIGN_RESIDENCE, nil
	case "警官证":
		return CASHDESK_CERTIFICATE_TYPE__POLICE_CARD, nil
	case "其他个人证件类型":
		return CASHDESK_CERTIFICATE_TYPE__PERSON_OTHER, nil
	case "组织机构代码证":
		return CASHDESK_CERTIFICATE_TYPE__ORG_INSITITUTE_CODE, nil
	case "营业执照":
		return CASHDESK_CERTIFICATE_TYPE__BUSINESS_LICENSE, nil
	case "统一社会信用代码":
		return CASHDESK_CERTIFICATE_TYPE__UNITY_SOCIAL_CREDIT_CODE, nil
	case "法人代码证":
		return CASHDESK_CERTIFICATE_TYPE__LEGAL_PERSON_CODE, nil
	case "单位统一代码":
		return CASHDESK_CERTIFICATE_TYPE__UNIT_UNITY_CODE, nil
	case "金融机构":
		return CASHDESK_CERTIFICATE_TYPE__FINANCIAL_ORG, nil
	case "其他公司证件类型":
		return CASHDESK_CERTIFICATE_TYPE__COMPANY_OTHER, nil
	}
	return CASHDESK_CERTIFICATE_TYPE_UNKNOWN, InvalidCashdeskCertificateType
}

func (v CashdeskCertificateType) String() string {
	switch v {
	case CASHDESK_CERTIFICATE_TYPE_UNKNOWN:
		return ""
	case CASHDESK_CERTIFICATE_TYPE__ID_CARD:
		return "ID_CARD"
	case CASHDESK_CERTIFICATE_TYPE__PASSPORT:
		return "PASSPORT"
	case CASHDESK_CERTIFICATE_TYPE__MILITARY:
		return "MILITARY"
	case CASHDESK_CERTIFICATE_TYPE__BOOKLET:
		return "BOOKLET"
	case CASHDESK_CERTIFICATE_TYPE__SOLIDER:
		return "SOLIDER"
	case CASHDESK_CERTIFICATE_TYPE__HK_MACAO_PASS:
		return "HK_MACAO_PASS"
	case CASHDESK_CERTIFICATE_TYPE__TW_PASS:
		return "TW_PASS"
	case CASHDESK_CERTIFICATE_TYPE__TEMP_ID_CARD:
		return "TEMP_ID_CARD"
	case CASHDESK_CERTIFICATE_TYPE__FOREIGN_RESIDENCE:
		return "FOREIGN_RESIDENCE"
	case CASHDESK_CERTIFICATE_TYPE__POLICE_CARD:
		return "POLICE_CARD"
	case CASHDESK_CERTIFICATE_TYPE__PERSON_OTHER:
		return "PERSON_OTHER"
	case CASHDESK_CERTIFICATE_TYPE__ORG_INSITITUTE_CODE:
		return "ORG_INSITITUTE_CODE"
	case CASHDESK_CERTIFICATE_TYPE__BUSINESS_LICENSE:
		return "BUSINESS_LICENSE"
	case CASHDESK_CERTIFICATE_TYPE__UNITY_SOCIAL_CREDIT_CODE:
		return "UNITY_SOCIAL_CREDIT_CODE"
	case CASHDESK_CERTIFICATE_TYPE__LEGAL_PERSON_CODE:
		return "LEGAL_PERSON_CODE"
	case CASHDESK_CERTIFICATE_TYPE__UNIT_UNITY_CODE:
		return "UNIT_UNITY_CODE"
	case CASHDESK_CERTIFICATE_TYPE__FINANCIAL_ORG:
		return "FINANCIAL_ORG"
	case CASHDESK_CERTIFICATE_TYPE__COMPANY_OTHER:
		return "COMPANY_OTHER"
	}
	return "UNKNOWN"
}

func (v CashdeskCertificateType) Label() string {
	switch v {
	case CASHDESK_CERTIFICATE_TYPE_UNKNOWN:
		return ""
	case CASHDESK_CERTIFICATE_TYPE__ID_CARD:
		return "身份证"
	case CASHDESK_CERTIFICATE_TYPE__PASSPORT:
		return "护照"
	case CASHDESK_CERTIFICATE_TYPE__MILITARY:
		return "军官证"
	case CASHDESK_CERTIFICATE_TYPE__BOOKLET:
		return "户口薄"
	case CASHDESK_CERTIFICATE_TYPE__SOLIDER:
		return "士兵证"
	case CASHDESK_CERTIFICATE_TYPE__HK_MACAO_PASS:
		return "澳居台民来往内地通行证"
	case CASHDESK_CERTIFICATE_TYPE__TW_PASS:
		return "台湾居民来往内地通行证"
	case CASHDESK_CERTIFICATE_TYPE__TEMP_ID_CARD:
		return "临时身份证"
	case CASHDESK_CERTIFICATE_TYPE__FOREIGN_RESIDENCE:
		return "外国人居住证"
	case CASHDESK_CERTIFICATE_TYPE__POLICE_CARD:
		return "警官证"
	case CASHDESK_CERTIFICATE_TYPE__PERSON_OTHER:
		return "其他个人证件类型"
	case CASHDESK_CERTIFICATE_TYPE__ORG_INSITITUTE_CODE:
		return "组织机构代码证"
	case CASHDESK_CERTIFICATE_TYPE__BUSINESS_LICENSE:
		return "营业执照"
	case CASHDESK_CERTIFICATE_TYPE__UNITY_SOCIAL_CREDIT_CODE:
		return "统一社会信用代码"
	case CASHDESK_CERTIFICATE_TYPE__LEGAL_PERSON_CODE:
		return "法人代码证"
	case CASHDESK_CERTIFICATE_TYPE__UNIT_UNITY_CODE:
		return "单位统一代码"
	case CASHDESK_CERTIFICATE_TYPE__FINANCIAL_ORG:
		return "金融机构"
	case CASHDESK_CERTIFICATE_TYPE__COMPANY_OTHER:
		return "其他公司证件类型"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskCertificateType)(nil)

func (v CashdeskCertificateType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskCertificateType
	}
	return []byte(str), nil
}

func (v *CashdeskCertificateType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskCertificateTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPabAccountType = errors.New("invalid CashdeskPabAccountType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPabAccountType", map[string]string{
		"COMPANY": "企业开户",
		"PERSON":  "个人开户",
	})
}

func ParseCashdeskPabAccountTypeFromString(s string) (CashdeskPabAccountType, error) {
	switch s {
	case "":
		return CASHDESK_PAB_ACCOUNT_TYPE_UNKNOWN, nil
	case "COMPANY":
		return CASHDESK_PAB_ACCOUNT_TYPE__COMPANY, nil
	case "PERSON":
		return CASHDESK_PAB_ACCOUNT_TYPE__PERSON, nil
	}
	return CASHDESK_PAB_ACCOUNT_TYPE_UNKNOWN, InvalidCashdeskPabAccountType
}

func ParseCashdeskPabAccountTypeFromLabelString(s string) (CashdeskPabAccountType, error) {
	switch s {
	case "":
		return CASHDESK_PAB_ACCOUNT_TYPE_UNKNOWN, nil
	case "企业开户":
		return CASHDESK_PAB_ACCOUNT_TYPE__COMPANY, nil
	case "个人开户":
		return CASHDESK_PAB_ACCOUNT_TYPE__PERSON, nil
	}
	return CASHDESK_PAB_ACCOUNT_TYPE_UNKNOWN, InvalidCashdeskPabAccountType
}

func (v CashdeskPabAccountType) String() string {
	switch v {
	case CASHDESK_PAB_ACCOUNT_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAB_ACCOUNT_TYPE__COMPANY:
		return "COMPANY"
	case CASHDESK_PAB_ACCOUNT_TYPE__PERSON:
		return "PERSON"
	}
	return "UNKNOWN"
}

func (v CashdeskPabAccountType) Label() string {
	switch v {
	case CASHDESK_PAB_ACCOUNT_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAB_ACCOUNT_TYPE__COMPANY:
		return "企业开户"
	case CASHDESK_PAB_ACCOUNT_TYPE__PERSON:
		return "个人开户"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPabAccountType)(nil)

func (v CashdeskPabAccountType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPabAccountType
	}
	return []byte(str), nil
}

func (v *CashdeskPabAccountType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPabAccountTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPabAccountState = errors.New("invalid CashdeskPabAccountState")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPabAccountState", map[string]string{
		"UNOPEN":    "待开户",
		"UNAUTH":    "待验证",
		"UNCONFIRM": "待确认",
		"UNSIGN":    "待签约",
		"SIGNED":    "已签约",
		"UNCLOSE":   "待销户",
	})
}

func ParseCashdeskPabAccountStateFromString(s string) (CashdeskPabAccountState, error) {
	switch s {
	case "":
		return CASHDESK_PAB_ACCOUNT_STATE_UNKNOWN, nil
	case "UNOPEN":
		return CASHDESK_PAB_ACCOUNT_STATE__UNOPEN, nil
	case "UNAUTH":
		return CASHDESK_PAB_ACCOUNT_STATE__UNAUTH, nil
	case "UNCONFIRM":
		return CASHDESK_PAB_ACCOUNT_STATE__UNCONFIRM, nil
	case "UNSIGN":
		return CASHDESK_PAB_ACCOUNT_STATE__UNSIGN, nil
	case "SIGNED":
		return CASHDESK_PAB_ACCOUNT_STATE__SIGNED, nil
	case "UNCLOSE":
		return CASHDESK_PAB_ACCOUNT_STATE__UNCLOSE, nil
	}
	return CASHDESK_PAB_ACCOUNT_STATE_UNKNOWN, InvalidCashdeskPabAccountState
}

func ParseCashdeskPabAccountStateFromLabelString(s string) (CashdeskPabAccountState, error) {
	switch s {
	case "":
		return CASHDESK_PAB_ACCOUNT_STATE_UNKNOWN, nil
	case "待开户":
		return CASHDESK_PAB_ACCOUNT_STATE__UNOPEN, nil
	case "待验证":
		return CASHDESK_PAB_ACCOUNT_STATE__UNAUTH, nil
	case "待确认":
		return CASHDESK_PAB_ACCOUNT_STATE__UNCONFIRM, nil
	case "待签约":
		return CASHDESK_PAB_ACCOUNT_STATE__UNSIGN, nil
	case "已签约":
		return CASHDESK_PAB_ACCOUNT_STATE__SIGNED, nil
	case "待销户":
		return CASHDESK_PAB_ACCOUNT_STATE__UNCLOSE, nil
	}
	return CASHDESK_PAB_ACCOUNT_STATE_UNKNOWN, InvalidCashdeskPabAccountState
}

func (v CashdeskPabAccountState) String() string {
	switch v {
	case CASHDESK_PAB_ACCOUNT_STATE_UNKNOWN:
		return ""
	case CASHDESK_PAB_ACCOUNT_STATE__UNOPEN:
		return "UNOPEN"
	case CASHDESK_PAB_ACCOUNT_STATE__UNAUTH:
		return "UNAUTH"
	case CASHDESK_PAB_ACCOUNT_STATE__UNCONFIRM:
		return "UNCONFIRM"
	case CASHDESK_PAB_ACCOUNT_STATE__UNSIGN:
		return "UNSIGN"
	case CASHDESK_PAB_ACCOUNT_STATE__SIGNED:
		return "SIGNED"
	case CASHDESK_PAB_ACCOUNT_STATE__UNCLOSE:
		return "UNCLOSE"
	}
	return "UNKNOWN"
}

func (v CashdeskPabAccountState) Label() string {
	switch v {
	case CASHDESK_PAB_ACCOUNT_STATE_UNKNOWN:
		return ""
	case CASHDESK_PAB_ACCOUNT_STATE__UNOPEN:
		return "待开户"
	case CASHDESK_PAB_ACCOUNT_STATE__UNAUTH:
		return "待验证"
	case CASHDESK_PAB_ACCOUNT_STATE__UNCONFIRM:
		return "待确认"
	case CASHDESK_PAB_ACCOUNT_STATE__UNSIGN:
		return "待签约"
	case CASHDESK_PAB_ACCOUNT_STATE__SIGNED:
		return "已签约"
	case CASHDESK_PAB_ACCOUNT_STATE__UNCLOSE:
		return "待销户"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPabAccountState)(nil)

func (v CashdeskPabAccountState) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPabAccountState
	}
	return []byte(str), nil
}

func (v *CashdeskPabAccountState) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPabAccountStateFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskTransState = errors.New("invalid CashdeskTransState")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskTransState", map[string]string{
		"TO_PAY":      "等待付款",
		"PAY_OK":      "付款成功",
		"TRANS_OK":    "交易成功",
		"REFUND_OK":   "退款成功",
		"CANCEL":      "交易取消",
		"STOP":        "交易终止",
		"PART_REFUND": "部分退款",
		"PRE_PAY":     "预支付完成",
	})
}

func ParseCashdeskTransStateFromString(s string) (CashdeskTransState, error) {
	switch s {
	case "":
		return CASHDESK_TRANS_STATE_UNKNOWN, nil
	case "TO_PAY":
		return CASHDESK_TRANS_STATE__TO_PAY, nil
	case "PAY_OK":
		return CASHDESK_TRANS_STATE__PAY_OK, nil
	case "TRANS_OK":
		return CASHDESK_TRANS_STATE__TRANS_OK, nil
	case "REFUND_OK":
		return CASHDESK_TRANS_STATE__REFUND_OK, nil
	case "CANCEL":
		return CASHDESK_TRANS_STATE__CANCEL, nil
	case "STOP":
		return CASHDESK_TRANS_STATE__STOP, nil
	case "PART_REFUND":
		return CASHDESK_TRANS_STATE__PART_REFUND, nil
	case "PRE_PAY":
		return CASHDESK_TRANS_STATE__PRE_PAY, nil
	}
	return CASHDESK_TRANS_STATE_UNKNOWN, InvalidCashdeskTransState
}

func ParseCashdeskTransStateFromLabelString(s string) (CashdeskTransState, error) {
	switch s {
	case "":
		return CASHDESK_TRANS_STATE_UNKNOWN, nil
	case "等待付款":
		return CASHDESK_TRANS_STATE__TO_PAY, nil
	case "付款成功":
		return CASHDESK_TRANS_STATE__PAY_OK, nil
	case "交易成功":
		return CASHDESK_TRANS_STATE__TRANS_OK, nil
	case "退款成功":
		return CASHDESK_TRANS_STATE__REFUND_OK, nil
	case "交易取消":
		return CASHDESK_TRANS_STATE__CANCEL, nil
	case "交易终止":
		return CASHDESK_TRANS_STATE__STOP, nil
	case "部分退款":
		return CASHDESK_TRANS_STATE__PART_REFUND, nil
	case "预支付完成":
		return CASHDESK_TRANS_STATE__PRE_PAY, nil
	}
	return CASHDESK_TRANS_STATE_UNKNOWN, InvalidCashdeskTransState
}

func (v CashdeskTransState) String() string {
	switch v {
	case CASHDESK_TRANS_STATE_UNKNOWN:
		return ""
	case CASHDESK_TRANS_STATE__TO_PAY:
		return "TO_PAY"
	case CASHDESK_TRANS_STATE__PAY_OK:
		return "PAY_OK"
	case CASHDESK_TRANS_STATE__TRANS_OK:
		return "TRANS_OK"
	case CASHDESK_TRANS_STATE__REFUND_OK:
		return "REFUND_OK"
	case CASHDESK_TRANS_STATE__CANCEL:
		return "CANCEL"
	case CASHDESK_TRANS_STATE__STOP:
		return "STOP"
	case CASHDESK_TRANS_STATE__PART_REFUND:
		return "PART_REFUND"
	case CASHDESK_TRANS_STATE__PRE_PAY:
		return "PRE_PAY"
	}
	return "UNKNOWN"
}

func (v CashdeskTransState) Label() string {
	switch v {
	case CASHDESK_TRANS_STATE_UNKNOWN:
		return ""
	case CASHDESK_TRANS_STATE__TO_PAY:
		return "等待付款"
	case CASHDESK_TRANS_STATE__PAY_OK:
		return "付款成功"
	case CASHDESK_TRANS_STATE__TRANS_OK:
		return "交易成功"
	case CASHDESK_TRANS_STATE__REFUND_OK:
		return "退款成功"
	case CASHDESK_TRANS_STATE__CANCEL:
		return "交易取消"
	case CASHDESK_TRANS_STATE__STOP:
		return "交易终止"
	case CASHDESK_TRANS_STATE__PART_REFUND:
		return "部分退款"
	case CASHDESK_TRANS_STATE__PRE_PAY:
		return "预支付完成"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskTransState)(nil)

func (v CashdeskTransState) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskTransState
	}
	return []byte(str), nil
}

func (v *CashdeskTransState) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskTransStateFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPayMode = errors.New("invalid CashdeskPayMode")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPayMode", map[string]string{
		"QUICK_PAY":     "快捷支付",
		"EBANK_PAY":     "网银支付",
		"WITHHOLD":      "代扣",
		"PAB":           "平安易宝",
		"WXPAY":         "微信支付",
		"ALIPAY":        "支付宝",
		"PAB_WITNESS":   "平安见证宝",
		"TRANSFER":      "转账",
		"WXPAY_QRCODE":  "微信扫码支付",
		"ALIPAY_QRCODE": "支付宝扫码支付",
	})
}

func ParseCashdeskPayModeFromString(s string) (CashdeskPayMode, error) {
	switch s {
	case "":
		return CASHDESK_PAY_MODE_UNKNOWN, nil
	case "QUICK_PAY":
		return CASHDESK_PAY_MODE__QUICK_PAY, nil
	case "EBANK_PAY":
		return CASHDESK_PAY_MODE__EBANK_PAY, nil
	case "WITHHOLD":
		return CASHDESK_PAY_MODE__WITHHOLD, nil
	case "PAB":
		return CASHDESK_PAY_MODE__PAB, nil
	case "WXPAY":
		return CASHDESK_PAY_MODE__WXPAY, nil
	case "ALIPAY":
		return CASHDESK_PAY_MODE__ALIPAY, nil
	case "PAB_WITNESS":
		return CASHDESK_PAY_MODE__PAB_WITNESS, nil
	case "TRANSFER":
		return CASHDESK_PAY_MODE__TRANSFER, nil
	case "WXPAY_QRCODE":
		return CASHDESK_PAY_MODE__WXPAY_QRCODE, nil
	case "ALIPAY_QRCODE":
		return CASHDESK_PAY_MODE__ALIPAY_QRCODE, nil
	}
	return CASHDESK_PAY_MODE_UNKNOWN, InvalidCashdeskPayMode
}

func ParseCashdeskPayModeFromLabelString(s string) (CashdeskPayMode, error) {
	switch s {
	case "":
		return CASHDESK_PAY_MODE_UNKNOWN, nil
	case "快捷支付":
		return CASHDESK_PAY_MODE__QUICK_PAY, nil
	case "网银支付":
		return CASHDESK_PAY_MODE__EBANK_PAY, nil
	case "代扣":
		return CASHDESK_PAY_MODE__WITHHOLD, nil
	case "平安易宝":
		return CASHDESK_PAY_MODE__PAB, nil
	case "微信支付":
		return CASHDESK_PAY_MODE__WXPAY, nil
	case "支付宝":
		return CASHDESK_PAY_MODE__ALIPAY, nil
	case "平安见证宝":
		return CASHDESK_PAY_MODE__PAB_WITNESS, nil
	case "转账":
		return CASHDESK_PAY_MODE__TRANSFER, nil
	case "微信扫码支付":
		return CASHDESK_PAY_MODE__WXPAY_QRCODE, nil
	case "支付宝扫码支付":
		return CASHDESK_PAY_MODE__ALIPAY_QRCODE, nil
	}
	return CASHDESK_PAY_MODE_UNKNOWN, InvalidCashdeskPayMode
}

func (v CashdeskPayMode) String() string {
	switch v {
	case CASHDESK_PAY_MODE_UNKNOWN:
		return ""
	case CASHDESK_PAY_MODE__QUICK_PAY:
		return "QUICK_PAY"
	case CASHDESK_PAY_MODE__EBANK_PAY:
		return "EBANK_PAY"
	case CASHDESK_PAY_MODE__WITHHOLD:
		return "WITHHOLD"
	case CASHDESK_PAY_MODE__PAB:
		return "PAB"
	case CASHDESK_PAY_MODE__WXPAY:
		return "WXPAY"
	case CASHDESK_PAY_MODE__ALIPAY:
		return "ALIPAY"
	case CASHDESK_PAY_MODE__PAB_WITNESS:
		return "PAB_WITNESS"
	case CASHDESK_PAY_MODE__TRANSFER:
		return "TRANSFER"
	case CASHDESK_PAY_MODE__WXPAY_QRCODE:
		return "WXPAY_QRCODE"
	case CASHDESK_PAY_MODE__ALIPAY_QRCODE:
		return "ALIPAY_QRCODE"
	}
	return "UNKNOWN"
}

func (v CashdeskPayMode) Label() string {
	switch v {
	case CASHDESK_PAY_MODE_UNKNOWN:
		return ""
	case CASHDESK_PAY_MODE__QUICK_PAY:
		return "快捷支付"
	case CASHDESK_PAY_MODE__EBANK_PAY:
		return "网银支付"
	case CASHDESK_PAY_MODE__WITHHOLD:
		return "代扣"
	case CASHDESK_PAY_MODE__PAB:
		return "平安易宝"
	case CASHDESK_PAY_MODE__WXPAY:
		return "微信支付"
	case CASHDESK_PAY_MODE__ALIPAY:
		return "支付宝"
	case CASHDESK_PAY_MODE__PAB_WITNESS:
		return "平安见证宝"
	case CASHDESK_PAY_MODE__TRANSFER:
		return "转账"
	case CASHDESK_PAY_MODE__WXPAY_QRCODE:
		return "微信扫码支付"
	case CASHDESK_PAY_MODE__ALIPAY_QRCODE:
		return "支付宝扫码支付"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPayMode)(nil)

func (v CashdeskPayMode) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPayMode
	}
	return []byte(str), nil
}

func (v *CashdeskPayMode) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPayModeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskSettlementType = errors.New("invalid CashdeskSettlementType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskSettlementType", map[string]string{
		"REAL_TIME":    "实时结算",
		"ASYNCHRONOUS": "异步结算",
	})
}

func ParseCashdeskSettlementTypeFromString(s string) (CashdeskSettlementType, error) {
	switch s {
	case "":
		return CASHDESK_SETTLEMENT_TYPE_UNKNOWN, nil
	case "REAL_TIME":
		return CASHDESK_SETTLEMENT_TYPE__REAL_TIME, nil
	case "ASYNCHRONOUS":
		return CASHDESK_SETTLEMENT_TYPE__ASYNCHRONOUS, nil
	}
	return CASHDESK_SETTLEMENT_TYPE_UNKNOWN, InvalidCashdeskSettlementType
}

func ParseCashdeskSettlementTypeFromLabelString(s string) (CashdeskSettlementType, error) {
	switch s {
	case "":
		return CASHDESK_SETTLEMENT_TYPE_UNKNOWN, nil
	case "实时结算":
		return CASHDESK_SETTLEMENT_TYPE__REAL_TIME, nil
	case "异步结算":
		return CASHDESK_SETTLEMENT_TYPE__ASYNCHRONOUS, nil
	}
	return CASHDESK_SETTLEMENT_TYPE_UNKNOWN, InvalidCashdeskSettlementType
}

func (v CashdeskSettlementType) String() string {
	switch v {
	case CASHDESK_SETTLEMENT_TYPE_UNKNOWN:
		return ""
	case CASHDESK_SETTLEMENT_TYPE__REAL_TIME:
		return "REAL_TIME"
	case CASHDESK_SETTLEMENT_TYPE__ASYNCHRONOUS:
		return "ASYNCHRONOUS"
	}
	return "UNKNOWN"
}

func (v CashdeskSettlementType) Label() string {
	switch v {
	case CASHDESK_SETTLEMENT_TYPE_UNKNOWN:
		return ""
	case CASHDESK_SETTLEMENT_TYPE__REAL_TIME:
		return "实时结算"
	case CASHDESK_SETTLEMENT_TYPE__ASYNCHRONOUS:
		return "异步结算"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskSettlementType)(nil)

func (v CashdeskSettlementType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskSettlementType
	}
	return []byte(str), nil
}

func (v *CashdeskSettlementType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskSettlementTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskTransMode = errors.New("invalid CashdeskTransMode")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskTransMode", map[string]string{
		"INTERMEDIARY":      "中介交易",
		"DIRECT_PAY":        "直付交易",
		"PREPAY":            "预付交易",
		"INSURANCE_FINANCE": "保理交易",
	})
}

func ParseCashdeskTransModeFromString(s string) (CashdeskTransMode, error) {
	switch s {
	case "":
		return CASHDESK_TRANS_MODE_UNKNOWN, nil
	case "INTERMEDIARY":
		return CASHDESK_TRANS_MODE__INTERMEDIARY, nil
	case "DIRECT_PAY":
		return CASHDESK_TRANS_MODE__DIRECT_PAY, nil
	case "PREPAY":
		return CASHDESK_TRANS_MODE__PREPAY, nil
	case "INSURANCE_FINANCE":
		return CASHDESK_TRANS_MODE__INSURANCE_FINANCE, nil
	}
	return CASHDESK_TRANS_MODE_UNKNOWN, InvalidCashdeskTransMode
}

func ParseCashdeskTransModeFromLabelString(s string) (CashdeskTransMode, error) {
	switch s {
	case "":
		return CASHDESK_TRANS_MODE_UNKNOWN, nil
	case "中介交易":
		return CASHDESK_TRANS_MODE__INTERMEDIARY, nil
	case "直付交易":
		return CASHDESK_TRANS_MODE__DIRECT_PAY, nil
	case "预付交易":
		return CASHDESK_TRANS_MODE__PREPAY, nil
	case "保理交易":
		return CASHDESK_TRANS_MODE__INSURANCE_FINANCE, nil
	}
	return CASHDESK_TRANS_MODE_UNKNOWN, InvalidCashdeskTransMode
}

func (v CashdeskTransMode) String() string {
	switch v {
	case CASHDESK_TRANS_MODE_UNKNOWN:
		return ""
	case CASHDESK_TRANS_MODE__INTERMEDIARY:
		return "INTERMEDIARY"
	case CASHDESK_TRANS_MODE__DIRECT_PAY:
		return "DIRECT_PAY"
	case CASHDESK_TRANS_MODE__PREPAY:
		return "PREPAY"
	case CASHDESK_TRANS_MODE__INSURANCE_FINANCE:
		return "INSURANCE_FINANCE"
	}
	return "UNKNOWN"
}

func (v CashdeskTransMode) Label() string {
	switch v {
	case CASHDESK_TRANS_MODE_UNKNOWN:
		return ""
	case CASHDESK_TRANS_MODE__INTERMEDIARY:
		return "中介交易"
	case CASHDESK_TRANS_MODE__DIRECT_PAY:
		return "直付交易"
	case CASHDESK_TRANS_MODE__PREPAY:
		return "预付交易"
	case CASHDESK_TRANS_MODE__INSURANCE_FINANCE:
		return "保理交易"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskTransMode)(nil)

func (v CashdeskTransMode) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskTransMode
	}
	return []byte(str), nil
}

func (v *CashdeskTransMode) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskTransModeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskCardType = errors.New("invalid CashdeskCardType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskCardType", map[string]string{
		"DEBIT":  "借记卡",
		"CREDIT": "信用卡",
	})
}

func ParseCashdeskCardTypeFromString(s string) (CashdeskCardType, error) {
	switch s {
	case "":
		return CASHDESK_CARD_TYPE_UNKNOWN, nil
	case "DEBIT":
		return CASHDESK_CARD_TYPE__DEBIT, nil
	case "CREDIT":
		return CASHDESK_CARD_TYPE__CREDIT, nil
	}
	return CASHDESK_CARD_TYPE_UNKNOWN, InvalidCashdeskCardType
}

func ParseCashdeskCardTypeFromLabelString(s string) (CashdeskCardType, error) {
	switch s {
	case "":
		return CASHDESK_CARD_TYPE_UNKNOWN, nil
	case "借记卡":
		return CASHDESK_CARD_TYPE__DEBIT, nil
	case "信用卡":
		return CASHDESK_CARD_TYPE__CREDIT, nil
	}
	return CASHDESK_CARD_TYPE_UNKNOWN, InvalidCashdeskCardType
}

func (v CashdeskCardType) String() string {
	switch v {
	case CASHDESK_CARD_TYPE_UNKNOWN:
		return ""
	case CASHDESK_CARD_TYPE__DEBIT:
		return "DEBIT"
	case CASHDESK_CARD_TYPE__CREDIT:
		return "CREDIT"
	}
	return "UNKNOWN"
}

func (v CashdeskCardType) Label() string {
	switch v {
	case CASHDESK_CARD_TYPE_UNKNOWN:
		return ""
	case CASHDESK_CARD_TYPE__DEBIT:
		return "借记卡"
	case CASHDESK_CARD_TYPE__CREDIT:
		return "信用卡"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskCardType)(nil)

func (v CashdeskCardType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskCardType
	}
	return []byte(str), nil
}

func (v *CashdeskCardType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskCardTypeFromString(string(bytes.ToUpper(data)))
	return
}
