package client_cashdesk

import (
	"bytes"
	"encoding"
	"errors"

	golib_tools_courier_enumeration "golib/tools/courier/enumeration"
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
	CASHDESK_CERTIFICATE_TYPE__ORG_INSITITUTE_CODE      CashdeskCertificateType = iota + 100 // 组织机构代码证
	CASHDESK_CERTIFICATE_TYPE__BUSINESS_LICENSE                                              // 营业执照
	CASHDESK_CERTIFICATE_TYPE__UNITY_SOCIAL_CREDIT_CODE                                      // 统一社会信用代码
	CASHDESK_CERTIFICATE_TYPE__LEGAL_PERSON_CODE                                             // 法人代码证
	CASHDESK_CERTIFICATE_TYPE__UNIT_UNITY_CODE                                               // 单位统一代码
	CASHDESK_CERTIFICATE_TYPE__FINANCIAL_ORG                                                 // 金融机构
	CASHDESK_CERTIFICATE_TYPE__COMPANY_OTHER                                                 // 其他公司证件类型
)

// swagger:enum
type CashdeskCurrencyType uint

const (
	CASHDESK_CURRENCY_TYPE_UNKNOWN CashdeskCurrencyType = iota
	CASHDESK_CURRENCY_TYPE__RMB                         // 人民币
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
type CashdeskPabAccountType uint

const (
	CASHDESK_PAB_ACCOUNT_TYPE_UNKNOWN  CashdeskPabAccountType = iota
	CASHDESK_PAB_ACCOUNT_TYPE__COMPANY                        // 企业开户
	CASHDESK_PAB_ACCOUNT_TYPE__PERSON                         // 个人开户
)

// swagger:enum
type CashdeskPayAdjustFlag uint

const (
	CASHDESK_PAY_ADJUST_FLAG_UNKNOWN     CashdeskPayAdjustFlag = iota
	CASHDESK_PAY_ADJUST_FLAG__UNADJUSTED                       // 未调帐
	CASHDESK_PAY_ADJUST_FLAG__ADJUSTED                         // 已调帐
)

// swagger:enum
type CashdeskPayCurrencyType uint

const (
	CASHDESK_PAY_CURRENCY_TYPE_UNKNOWN CashdeskPayCurrencyType = iota
	CASHDESK_PAY_CURRENCY_TYPE__RMB                            // 人民币
)

// swagger:enum
type CashdeskPayPayType uint

const (
	CASHDESK_PAY_PAY_TYPE_UNKNOWN   CashdeskPayPayType = iota
	CASHDESK_PAY_PAY_TYPE__BALANCE                     // 余额支付
	CASHDESK_PAY_PAY_TYPE__RECHARGE                    // 充值支付
)

// swagger:enum
type CashdeskPayPlatformType uint

const (
	CASHDESK_PAY_PLATFORM_TYPE_UNKNOWN       CashdeskPayPlatformType = iota
	CASHDESK_PAY_PLATFORM_TYPE__COMPUTER                             // 个人电脑
	CASHDESK_PAY_PLATFORM_TYPE__MULTI_MEDIA                          // 多媒体终端
	CASHDESK_PAY_PLATFORM_TYPE__MOBILE_PHONE                         // 手持智能设备（手机、MID等）
	CASHDESK_PAY_PLATFORM_TYPE__PAD                                  // 平板电脑
	CASHDESK_PAY_PLATFORM_TYPE__POS                                  // POS终端
	CASHDESK_PAY_PLATFORM_TYPE__MERCHANT                             // 商户系统
	CASHDESK_PAY_PLATFORM_TYPE__STB                                  // 数字机顶盒
	CASHDESK_PAY_PLATFORM_TYPE__TV                                   // 智能电视
	CASHDESK_PAY_PLATFORM_TYPE__VEM                                  // 自动柜员机（售货机等）
	CASHDESK_PAY_PLATFORM_TYPE__THIRD_SYSTEM                         // 第三方机构系统
)

// swagger:enum
type CashdeskPaySettlementType uint

const (
	CASHDESK_PAY_SETTLEMENT_TYPE_UNKNOWN       CashdeskPaySettlementType = iota
	CASHDESK_PAY_SETTLEMENT_TYPE__REAL_TIME                              // 实时结算
	CASHDESK_PAY_SETTLEMENT_TYPE__ASYNCHRONOUS                           // 异步结算
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
type CashdeskPaySubAccountType uint

const (
	CASHDESK_PAY_SUB_ACCOUNT_TYPE_UNKNOWN         CashdeskPaySubAccountType = iota
	CASHDESK_PAY_SUB_ACCOUNT_TYPE__CASH                                     // 现金子帐户
	CASHDESK_PAY_SUB_ACCOUNT_TYPE__PENDING_SETTLE                           // 待结算子帐户
	CASHDESK_PAY_SUB_ACCOUNT_TYPE__POINT                                    // 积分帐户
	CASHDESK_PAY_SUB_ACCOUNT_TYPE__DEPOSIT                                  // 保证金帐户
	CASHDESK_PAY_SUB_ACCOUNT_TYPE__FROZEN                                   // 冻结帐户
	CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB                                      // 平安现金子帐户
	CASHDESK_PAY_SUB_ACCOUNT_TYPE__OIL                                      // 油品子账户
	CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB_WITNESS                              // 见证宝现金子帐户
	CASHDESK_PAY_SUB_ACCOUNT_TYPE__TRANSIT                                  // 在途子账户
	CASHDESK_PAY_SUB_ACCOUNT_TYPE__PRIVATE                                  // 私有子账户
	CASHDESK_PAY_SUB_ACCOUNT_TYPE__GLP                                      // 普洛斯子账户
)

// swagger:enum
type CashdeskPayTransMode uint

const (
	CASHDESK_PAY_TRANS_MODE_UNKNOWN            CashdeskPayTransMode = iota
	CASHDESK_PAY_TRANS_MODE__INTERMEDIARY                           // 中介交易
	CASHDESK_PAY_TRANS_MODE__DIRECT_PAY                             // 直付交易
	CASHDESK_PAY_TRANS_MODE__PREPAY                                 // 预付交易
	CASHDESK_PAY_TRANS_MODE__INSURANCE_FINANCE                      // 保理交易
)

// swagger:enum
type CashdeskPayTransState uint

const (
	CASHDESK_PAY_TRANS_STATE_UNKNOWN      CashdeskPayTransState = iota
	CASHDESK_PAY_TRANS_STATE__TO_PAY                            // 等待付款
	CASHDESK_PAY_TRANS_STATE__PAY_OK                            // 付款成功
	CASHDESK_PAY_TRANS_STATE__TRANS_OK                          // 交易成功
	CASHDESK_PAY_TRANS_STATE__REFUND_OK                         // 退款成功
	CASHDESK_PAY_TRANS_STATE__CANCEL                            // 交易取消
	CASHDESK_PAY_TRANS_STATE__STOP                              // 交易终止
	CASHDESK_PAY_TRANS_STATE__PART_REFUND                       // 部分退款
	CASHDESK_PAY_TRANS_STATE__PRE_PAY                           // 预支付完成
)

// swagger:enum
type CashdeskPayTransType uint

const (
	CASHDESK_PAY_TRANS_TYPE_UNKNOWN       CashdeskPayTransType = iota
	CASHDESK_PAY_TRANS_TYPE__ETC                               // ETC
	CASHDESK_PAY_TRANS_TYPE__OIL_CARD                          // 油卡
	CASHDESK_PAY_TRANS_TYPE__TENDER                            // 招采
	CASHDESK_PAY_TRANS_TYPE__MALL                              // 商城
	CASHDESK_PAY_TRANS_TYPE__ZQX                               // 中启行(油)
	CASHDESK_PAY_TRANS_TYPE__SKID_MOUNTED                      // 撬装(油)
	CASHDESK_PAY_TRANS_TYPE__MAINTENANCE                       // 维保
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
type CashdeskSignAlgorithm uint

const (
	CASHDESK_SIGN_ALGORITHM_UNKNOWN CashdeskSignAlgorithm = iota
	CASHDESK_SIGN_ALGORITHM__MD5                          // MD5加密
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
type CashdeskUserFrom uint

const (
	CASHDESK_USER_FROM_UNKNOWN    CashdeskUserFrom = iota
	CASHDESK_USER_FROM__G7                         // G7平台
	CASHDESK_USER_FROM__ANONYMOUS                  // 匿名
	CASHDESK_USER_FROM__G7_ORG                     // G7机构
	CASHDESK_USER_FROM__WECHAT                     // 微信
)

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

func (CashdeskCertificateType) EnumType() string {
	return "CashdeskCertificateType"
}

func (CashdeskCertificateType) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_CERTIFICATE_TYPE__ID_CARD):                  {"ID_CARD", "身份证"},
		int(CASHDESK_CERTIFICATE_TYPE__PASSPORT):                 {"PASSPORT", "护照"},
		int(CASHDESK_CERTIFICATE_TYPE__MILITARY):                 {"MILITARY", "军官证"},
		int(CASHDESK_CERTIFICATE_TYPE__BOOKLET):                  {"BOOKLET", "户口薄"},
		int(CASHDESK_CERTIFICATE_TYPE__SOLIDER):                  {"SOLIDER", "士兵证"},
		int(CASHDESK_CERTIFICATE_TYPE__HK_MACAO_PASS):            {"HK_MACAO_PASS", "澳居台民来往内地通行证"},
		int(CASHDESK_CERTIFICATE_TYPE__TW_PASS):                  {"TW_PASS", "台湾居民来往内地通行证"},
		int(CASHDESK_CERTIFICATE_TYPE__TEMP_ID_CARD):             {"TEMP_ID_CARD", "临时身份证"},
		int(CASHDESK_CERTIFICATE_TYPE__FOREIGN_RESIDENCE):        {"FOREIGN_RESIDENCE", "外国人居住证"},
		int(CASHDESK_CERTIFICATE_TYPE__POLICE_CARD):              {"POLICE_CARD", "警官证"},
		int(CASHDESK_CERTIFICATE_TYPE__PERSON_OTHER):             {"PERSON_OTHER", "其他个人证件类型"},
		int(CASHDESK_CERTIFICATE_TYPE__ORG_INSITITUTE_CODE):      {"ORG_INSITITUTE_CODE", "组织机构代码证"},
		int(CASHDESK_CERTIFICATE_TYPE__BUSINESS_LICENSE):         {"BUSINESS_LICENSE", "营业执照"},
		int(CASHDESK_CERTIFICATE_TYPE__UNITY_SOCIAL_CREDIT_CODE): {"UNITY_SOCIAL_CREDIT_CODE", "统一社会信用代码"},
		int(CASHDESK_CERTIFICATE_TYPE__LEGAL_PERSON_CODE):        {"LEGAL_PERSON_CODE", "法人代码证"},
		int(CASHDESK_CERTIFICATE_TYPE__UNIT_UNITY_CODE):          {"UNIT_UNITY_CODE", "单位统一代码"},
		int(CASHDESK_CERTIFICATE_TYPE__FINANCIAL_ORG):            {"FINANCIAL_ORG", "金融机构"},
		int(CASHDESK_CERTIFICATE_TYPE__COMPANY_OTHER):            {"COMPANY_OTHER", "其他公司证件类型"},
	}
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

func (CashdeskCurrencyType) EnumType() string {
	return "CashdeskCurrencyType"
}

func (CashdeskCurrencyType) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_CURRENCY_TYPE__RMB): {"RMB", "人民币"},
	}
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

func (CashdeskPabAccountState) EnumType() string {
	return "CashdeskPabAccountState"
}

func (CashdeskPabAccountState) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAB_ACCOUNT_STATE__UNOPEN):    {"UNOPEN", "待开户"},
		int(CASHDESK_PAB_ACCOUNT_STATE__UNAUTH):    {"UNAUTH", "待验证"},
		int(CASHDESK_PAB_ACCOUNT_STATE__UNCONFIRM): {"UNCONFIRM", "待确认"},
		int(CASHDESK_PAB_ACCOUNT_STATE__UNSIGN):    {"UNSIGN", "待签约"},
		int(CASHDESK_PAB_ACCOUNT_STATE__SIGNED):    {"SIGNED", "已签约"},
		int(CASHDESK_PAB_ACCOUNT_STATE__UNCLOSE):   {"UNCLOSE", "待销户"},
	}
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

func (CashdeskPabAccountType) EnumType() string {
	return "CashdeskPabAccountType"
}

func (CashdeskPabAccountType) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAB_ACCOUNT_TYPE__COMPANY): {"COMPANY", "企业开户"},
		int(CASHDESK_PAB_ACCOUNT_TYPE__PERSON):  {"PERSON", "个人开户"},
	}
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

var InvalidCashdeskPayAdjustFlag = errors.New("invalid CashdeskPayAdjustFlag")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPayAdjustFlag", map[string]string{
		"UNADJUSTED": "未调帐",
		"ADJUSTED":   "已调帐",
	})
}

func ParseCashdeskPayAdjustFlagFromString(s string) (CashdeskPayAdjustFlag, error) {
	switch s {
	case "":
		return CASHDESK_PAY_ADJUST_FLAG_UNKNOWN, nil
	case "UNADJUSTED":
		return CASHDESK_PAY_ADJUST_FLAG__UNADJUSTED, nil
	case "ADJUSTED":
		return CASHDESK_PAY_ADJUST_FLAG__ADJUSTED, nil
	}
	return CASHDESK_PAY_ADJUST_FLAG_UNKNOWN, InvalidCashdeskPayAdjustFlag
}

func ParseCashdeskPayAdjustFlagFromLabelString(s string) (CashdeskPayAdjustFlag, error) {
	switch s {
	case "":
		return CASHDESK_PAY_ADJUST_FLAG_UNKNOWN, nil
	case "未调帐":
		return CASHDESK_PAY_ADJUST_FLAG__UNADJUSTED, nil
	case "已调帐":
		return CASHDESK_PAY_ADJUST_FLAG__ADJUSTED, nil
	}
	return CASHDESK_PAY_ADJUST_FLAG_UNKNOWN, InvalidCashdeskPayAdjustFlag
}

func (CashdeskPayAdjustFlag) EnumType() string {
	return "CashdeskPayAdjustFlag"
}

func (CashdeskPayAdjustFlag) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAY_ADJUST_FLAG__UNADJUSTED): {"UNADJUSTED", "未调帐"},
		int(CASHDESK_PAY_ADJUST_FLAG__ADJUSTED):   {"ADJUSTED", "已调帐"},
	}
}
func (v CashdeskPayAdjustFlag) String() string {
	switch v {
	case CASHDESK_PAY_ADJUST_FLAG_UNKNOWN:
		return ""
	case CASHDESK_PAY_ADJUST_FLAG__UNADJUSTED:
		return "UNADJUSTED"
	case CASHDESK_PAY_ADJUST_FLAG__ADJUSTED:
		return "ADJUSTED"
	}
	return "UNKNOWN"
}

func (v CashdeskPayAdjustFlag) Label() string {
	switch v {
	case CASHDESK_PAY_ADJUST_FLAG_UNKNOWN:
		return ""
	case CASHDESK_PAY_ADJUST_FLAG__UNADJUSTED:
		return "未调帐"
	case CASHDESK_PAY_ADJUST_FLAG__ADJUSTED:
		return "已调帐"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPayAdjustFlag)(nil)

func (v CashdeskPayAdjustFlag) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPayAdjustFlag
	}
	return []byte(str), nil
}

func (v *CashdeskPayAdjustFlag) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPayAdjustFlagFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPayCurrencyType = errors.New("invalid CashdeskPayCurrencyType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPayCurrencyType", map[string]string{
		"RMB": "人民币",
	})
}

func ParseCashdeskPayCurrencyTypeFromString(s string) (CashdeskPayCurrencyType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_CURRENCY_TYPE_UNKNOWN, nil
	case "RMB":
		return CASHDESK_PAY_CURRENCY_TYPE__RMB, nil
	}
	return CASHDESK_PAY_CURRENCY_TYPE_UNKNOWN, InvalidCashdeskPayCurrencyType
}

func ParseCashdeskPayCurrencyTypeFromLabelString(s string) (CashdeskPayCurrencyType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_CURRENCY_TYPE_UNKNOWN, nil
	case "人民币":
		return CASHDESK_PAY_CURRENCY_TYPE__RMB, nil
	}
	return CASHDESK_PAY_CURRENCY_TYPE_UNKNOWN, InvalidCashdeskPayCurrencyType
}

func (CashdeskPayCurrencyType) EnumType() string {
	return "CashdeskPayCurrencyType"
}

func (CashdeskPayCurrencyType) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAY_CURRENCY_TYPE__RMB): {"RMB", "人民币"},
	}
}
func (v CashdeskPayCurrencyType) String() string {
	switch v {
	case CASHDESK_PAY_CURRENCY_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_CURRENCY_TYPE__RMB:
		return "RMB"
	}
	return "UNKNOWN"
}

func (v CashdeskPayCurrencyType) Label() string {
	switch v {
	case CASHDESK_PAY_CURRENCY_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_CURRENCY_TYPE__RMB:
		return "人民币"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPayCurrencyType)(nil)

func (v CashdeskPayCurrencyType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPayCurrencyType
	}
	return []byte(str), nil
}

func (v *CashdeskPayCurrencyType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPayCurrencyTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPayPayType = errors.New("invalid CashdeskPayPayType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPayPayType", map[string]string{
		"BALANCE":  "余额支付",
		"RECHARGE": "充值支付",
	})
}

func ParseCashdeskPayPayTypeFromString(s string) (CashdeskPayPayType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_PAY_TYPE_UNKNOWN, nil
	case "BALANCE":
		return CASHDESK_PAY_PAY_TYPE__BALANCE, nil
	case "RECHARGE":
		return CASHDESK_PAY_PAY_TYPE__RECHARGE, nil
	}
	return CASHDESK_PAY_PAY_TYPE_UNKNOWN, InvalidCashdeskPayPayType
}

func ParseCashdeskPayPayTypeFromLabelString(s string) (CashdeskPayPayType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_PAY_TYPE_UNKNOWN, nil
	case "余额支付":
		return CASHDESK_PAY_PAY_TYPE__BALANCE, nil
	case "充值支付":
		return CASHDESK_PAY_PAY_TYPE__RECHARGE, nil
	}
	return CASHDESK_PAY_PAY_TYPE_UNKNOWN, InvalidCashdeskPayPayType
}

func (CashdeskPayPayType) EnumType() string {
	return "CashdeskPayPayType"
}

func (CashdeskPayPayType) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAY_PAY_TYPE__BALANCE):  {"BALANCE", "余额支付"},
		int(CASHDESK_PAY_PAY_TYPE__RECHARGE): {"RECHARGE", "充值支付"},
	}
}
func (v CashdeskPayPayType) String() string {
	switch v {
	case CASHDESK_PAY_PAY_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_PAY_TYPE__BALANCE:
		return "BALANCE"
	case CASHDESK_PAY_PAY_TYPE__RECHARGE:
		return "RECHARGE"
	}
	return "UNKNOWN"
}

func (v CashdeskPayPayType) Label() string {
	switch v {
	case CASHDESK_PAY_PAY_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_PAY_TYPE__BALANCE:
		return "余额支付"
	case CASHDESK_PAY_PAY_TYPE__RECHARGE:
		return "充值支付"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPayPayType)(nil)

func (v CashdeskPayPayType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPayPayType
	}
	return []byte(str), nil
}

func (v *CashdeskPayPayType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPayPayTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPayPlatformType = errors.New("invalid CashdeskPayPlatformType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPayPlatformType", map[string]string{
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

func ParseCashdeskPayPlatformTypeFromString(s string) (CashdeskPayPlatformType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_PLATFORM_TYPE_UNKNOWN, nil
	case "COMPUTER":
		return CASHDESK_PAY_PLATFORM_TYPE__COMPUTER, nil
	case "MULTI_MEDIA":
		return CASHDESK_PAY_PLATFORM_TYPE__MULTI_MEDIA, nil
	case "MOBILE_PHONE":
		return CASHDESK_PAY_PLATFORM_TYPE__MOBILE_PHONE, nil
	case "PAD":
		return CASHDESK_PAY_PLATFORM_TYPE__PAD, nil
	case "POS":
		return CASHDESK_PAY_PLATFORM_TYPE__POS, nil
	case "MERCHANT":
		return CASHDESK_PAY_PLATFORM_TYPE__MERCHANT, nil
	case "STB":
		return CASHDESK_PAY_PLATFORM_TYPE__STB, nil
	case "TV":
		return CASHDESK_PAY_PLATFORM_TYPE__TV, nil
	case "VEM":
		return CASHDESK_PAY_PLATFORM_TYPE__VEM, nil
	case "THIRD_SYSTEM":
		return CASHDESK_PAY_PLATFORM_TYPE__THIRD_SYSTEM, nil
	}
	return CASHDESK_PAY_PLATFORM_TYPE_UNKNOWN, InvalidCashdeskPayPlatformType
}

func ParseCashdeskPayPlatformTypeFromLabelString(s string) (CashdeskPayPlatformType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_PLATFORM_TYPE_UNKNOWN, nil
	case "个人电脑":
		return CASHDESK_PAY_PLATFORM_TYPE__COMPUTER, nil
	case "多媒体终端":
		return CASHDESK_PAY_PLATFORM_TYPE__MULTI_MEDIA, nil
	case "手持智能设备（手机、MID等）":
		return CASHDESK_PAY_PLATFORM_TYPE__MOBILE_PHONE, nil
	case "平板电脑":
		return CASHDESK_PAY_PLATFORM_TYPE__PAD, nil
	case "POS终端":
		return CASHDESK_PAY_PLATFORM_TYPE__POS, nil
	case "商户系统":
		return CASHDESK_PAY_PLATFORM_TYPE__MERCHANT, nil
	case "数字机顶盒":
		return CASHDESK_PAY_PLATFORM_TYPE__STB, nil
	case "智能电视":
		return CASHDESK_PAY_PLATFORM_TYPE__TV, nil
	case "自动柜员机（售货机等）":
		return CASHDESK_PAY_PLATFORM_TYPE__VEM, nil
	case "第三方机构系统":
		return CASHDESK_PAY_PLATFORM_TYPE__THIRD_SYSTEM, nil
	}
	return CASHDESK_PAY_PLATFORM_TYPE_UNKNOWN, InvalidCashdeskPayPlatformType
}

func (CashdeskPayPlatformType) EnumType() string {
	return "CashdeskPayPlatformType"
}

func (CashdeskPayPlatformType) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAY_PLATFORM_TYPE__COMPUTER):     {"COMPUTER", "个人电脑"},
		int(CASHDESK_PAY_PLATFORM_TYPE__MULTI_MEDIA):  {"MULTI_MEDIA", "多媒体终端"},
		int(CASHDESK_PAY_PLATFORM_TYPE__MOBILE_PHONE): {"MOBILE_PHONE", "手持智能设备（手机、MID等）"},
		int(CASHDESK_PAY_PLATFORM_TYPE__PAD):          {"PAD", "平板电脑"},
		int(CASHDESK_PAY_PLATFORM_TYPE__POS):          {"POS", "POS终端"},
		int(CASHDESK_PAY_PLATFORM_TYPE__MERCHANT):     {"MERCHANT", "商户系统"},
		int(CASHDESK_PAY_PLATFORM_TYPE__STB):          {"STB", "数字机顶盒"},
		int(CASHDESK_PAY_PLATFORM_TYPE__TV):           {"TV", "智能电视"},
		int(CASHDESK_PAY_PLATFORM_TYPE__VEM):          {"VEM", "自动柜员机（售货机等）"},
		int(CASHDESK_PAY_PLATFORM_TYPE__THIRD_SYSTEM): {"THIRD_SYSTEM", "第三方机构系统"},
	}
}
func (v CashdeskPayPlatformType) String() string {
	switch v {
	case CASHDESK_PAY_PLATFORM_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_PLATFORM_TYPE__COMPUTER:
		return "COMPUTER"
	case CASHDESK_PAY_PLATFORM_TYPE__MULTI_MEDIA:
		return "MULTI_MEDIA"
	case CASHDESK_PAY_PLATFORM_TYPE__MOBILE_PHONE:
		return "MOBILE_PHONE"
	case CASHDESK_PAY_PLATFORM_TYPE__PAD:
		return "PAD"
	case CASHDESK_PAY_PLATFORM_TYPE__POS:
		return "POS"
	case CASHDESK_PAY_PLATFORM_TYPE__MERCHANT:
		return "MERCHANT"
	case CASHDESK_PAY_PLATFORM_TYPE__STB:
		return "STB"
	case CASHDESK_PAY_PLATFORM_TYPE__TV:
		return "TV"
	case CASHDESK_PAY_PLATFORM_TYPE__VEM:
		return "VEM"
	case CASHDESK_PAY_PLATFORM_TYPE__THIRD_SYSTEM:
		return "THIRD_SYSTEM"
	}
	return "UNKNOWN"
}

func (v CashdeskPayPlatformType) Label() string {
	switch v {
	case CASHDESK_PAY_PLATFORM_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_PLATFORM_TYPE__COMPUTER:
		return "个人电脑"
	case CASHDESK_PAY_PLATFORM_TYPE__MULTI_MEDIA:
		return "多媒体终端"
	case CASHDESK_PAY_PLATFORM_TYPE__MOBILE_PHONE:
		return "手持智能设备（手机、MID等）"
	case CASHDESK_PAY_PLATFORM_TYPE__PAD:
		return "平板电脑"
	case CASHDESK_PAY_PLATFORM_TYPE__POS:
		return "POS终端"
	case CASHDESK_PAY_PLATFORM_TYPE__MERCHANT:
		return "商户系统"
	case CASHDESK_PAY_PLATFORM_TYPE__STB:
		return "数字机顶盒"
	case CASHDESK_PAY_PLATFORM_TYPE__TV:
		return "智能电视"
	case CASHDESK_PAY_PLATFORM_TYPE__VEM:
		return "自动柜员机（售货机等）"
	case CASHDESK_PAY_PLATFORM_TYPE__THIRD_SYSTEM:
		return "第三方机构系统"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPayPlatformType)(nil)

func (v CashdeskPayPlatformType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPayPlatformType
	}
	return []byte(str), nil
}

func (v *CashdeskPayPlatformType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPayPlatformTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPaySettlementType = errors.New("invalid CashdeskPaySettlementType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPaySettlementType", map[string]string{
		"REAL_TIME":    "实时结算",
		"ASYNCHRONOUS": "异步结算",
	})
}

func ParseCashdeskPaySettlementTypeFromString(s string) (CashdeskPaySettlementType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_SETTLEMENT_TYPE_UNKNOWN, nil
	case "REAL_TIME":
		return CASHDESK_PAY_SETTLEMENT_TYPE__REAL_TIME, nil
	case "ASYNCHRONOUS":
		return CASHDESK_PAY_SETTLEMENT_TYPE__ASYNCHRONOUS, nil
	}
	return CASHDESK_PAY_SETTLEMENT_TYPE_UNKNOWN, InvalidCashdeskPaySettlementType
}

func ParseCashdeskPaySettlementTypeFromLabelString(s string) (CashdeskPaySettlementType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_SETTLEMENT_TYPE_UNKNOWN, nil
	case "实时结算":
		return CASHDESK_PAY_SETTLEMENT_TYPE__REAL_TIME, nil
	case "异步结算":
		return CASHDESK_PAY_SETTLEMENT_TYPE__ASYNCHRONOUS, nil
	}
	return CASHDESK_PAY_SETTLEMENT_TYPE_UNKNOWN, InvalidCashdeskPaySettlementType
}

func (CashdeskPaySettlementType) EnumType() string {
	return "CashdeskPaySettlementType"
}

func (CashdeskPaySettlementType) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAY_SETTLEMENT_TYPE__REAL_TIME):    {"REAL_TIME", "实时结算"},
		int(CASHDESK_PAY_SETTLEMENT_TYPE__ASYNCHRONOUS): {"ASYNCHRONOUS", "异步结算"},
	}
}
func (v CashdeskPaySettlementType) String() string {
	switch v {
	case CASHDESK_PAY_SETTLEMENT_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_SETTLEMENT_TYPE__REAL_TIME:
		return "REAL_TIME"
	case CASHDESK_PAY_SETTLEMENT_TYPE__ASYNCHRONOUS:
		return "ASYNCHRONOUS"
	}
	return "UNKNOWN"
}

func (v CashdeskPaySettlementType) Label() string {
	switch v {
	case CASHDESK_PAY_SETTLEMENT_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_SETTLEMENT_TYPE__REAL_TIME:
		return "实时结算"
	case CASHDESK_PAY_SETTLEMENT_TYPE__ASYNCHRONOUS:
		return "异步结算"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPaySettlementType)(nil)

func (v CashdeskPaySettlementType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPaySettlementType
	}
	return []byte(str), nil
}

func (v *CashdeskPaySettlementType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPaySettlementTypeFromString(string(bytes.ToUpper(data)))
	return
}

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

func (CashdeskPayState) EnumType() string {
	return "CashdeskPayState"
}

func (CashdeskPayState) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAY_STATE__UNPAY):   {"UNPAY", "未支付"},
		int(CASHDESK_PAY_STATE__DEALING): {"DEALING", "正在交易"},
		int(CASHDESK_PAY_STATE__SUCCESS): {"SUCCESS", "支付成功"},
		int(CASHDESK_PAY_STATE__FAIL):    {"FAIL", "支付失败"},
	}
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

var InvalidCashdeskPaySubAccountType = errors.New("invalid CashdeskPaySubAccountType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPaySubAccountType", map[string]string{
		"CASH":           "现金子帐户",
		"PENDING_SETTLE": "待结算子帐户",
		"POINT":          "积分帐户",
		"DEPOSIT":        "保证金帐户",
		"FROZEN":         "冻结帐户",
		"PAB":            "平安现金子帐户",
		"OIL":            "油品子账户",
		"PAB_WITNESS":    "见证宝现金子帐户",
		"TRANSIT":        "在途子账户",
		"PRIVATE":        "私有子账户",
		"GLP":            "普洛斯子账户",
	})
}

func ParseCashdeskPaySubAccountTypeFromString(s string) (CashdeskPaySubAccountType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE_UNKNOWN, nil
	case "CASH":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__CASH, nil
	case "PENDING_SETTLE":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__PENDING_SETTLE, nil
	case "POINT":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__POINT, nil
	case "DEPOSIT":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__DEPOSIT, nil
	case "FROZEN":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__FROZEN, nil
	case "PAB":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB, nil
	case "OIL":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__OIL, nil
	case "PAB_WITNESS":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB_WITNESS, nil
	case "TRANSIT":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__TRANSIT, nil
	case "PRIVATE":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__PRIVATE, nil
	case "GLP":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__GLP, nil
	}
	return CASHDESK_PAY_SUB_ACCOUNT_TYPE_UNKNOWN, InvalidCashdeskPaySubAccountType
}

func ParseCashdeskPaySubAccountTypeFromLabelString(s string) (CashdeskPaySubAccountType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE_UNKNOWN, nil
	case "现金子帐户":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__CASH, nil
	case "待结算子帐户":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__PENDING_SETTLE, nil
	case "积分帐户":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__POINT, nil
	case "保证金帐户":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__DEPOSIT, nil
	case "冻结帐户":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__FROZEN, nil
	case "平安现金子帐户":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB, nil
	case "油品子账户":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__OIL, nil
	case "见证宝现金子帐户":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB_WITNESS, nil
	case "在途子账户":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__TRANSIT, nil
	case "私有子账户":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__PRIVATE, nil
	case "普洛斯子账户":
		return CASHDESK_PAY_SUB_ACCOUNT_TYPE__GLP, nil
	}
	return CASHDESK_PAY_SUB_ACCOUNT_TYPE_UNKNOWN, InvalidCashdeskPaySubAccountType
}

func (CashdeskPaySubAccountType) EnumType() string {
	return "CashdeskPaySubAccountType"
}

func (CashdeskPaySubAccountType) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAY_SUB_ACCOUNT_TYPE__CASH):           {"CASH", "现金子帐户"},
		int(CASHDESK_PAY_SUB_ACCOUNT_TYPE__PENDING_SETTLE): {"PENDING_SETTLE", "待结算子帐户"},
		int(CASHDESK_PAY_SUB_ACCOUNT_TYPE__POINT):          {"POINT", "积分帐户"},
		int(CASHDESK_PAY_SUB_ACCOUNT_TYPE__DEPOSIT):        {"DEPOSIT", "保证金帐户"},
		int(CASHDESK_PAY_SUB_ACCOUNT_TYPE__FROZEN):         {"FROZEN", "冻结帐户"},
		int(CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB):            {"PAB", "平安现金子帐户"},
		int(CASHDESK_PAY_SUB_ACCOUNT_TYPE__OIL):            {"OIL", "油品子账户"},
		int(CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB_WITNESS):    {"PAB_WITNESS", "见证宝现金子帐户"},
		int(CASHDESK_PAY_SUB_ACCOUNT_TYPE__TRANSIT):        {"TRANSIT", "在途子账户"},
		int(CASHDESK_PAY_SUB_ACCOUNT_TYPE__PRIVATE):        {"PRIVATE", "私有子账户"},
		int(CASHDESK_PAY_SUB_ACCOUNT_TYPE__GLP):            {"GLP", "普洛斯子账户"},
	}
}
func (v CashdeskPaySubAccountType) String() string {
	switch v {
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__CASH:
		return "CASH"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__PENDING_SETTLE:
		return "PENDING_SETTLE"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__POINT:
		return "POINT"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__DEPOSIT:
		return "DEPOSIT"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__FROZEN:
		return "FROZEN"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB:
		return "PAB"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__OIL:
		return "OIL"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB_WITNESS:
		return "PAB_WITNESS"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__TRANSIT:
		return "TRANSIT"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__PRIVATE:
		return "PRIVATE"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__GLP:
		return "GLP"
	}
	return "UNKNOWN"
}

func (v CashdeskPaySubAccountType) Label() string {
	switch v {
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__CASH:
		return "现金子帐户"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__PENDING_SETTLE:
		return "待结算子帐户"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__POINT:
		return "积分帐户"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__DEPOSIT:
		return "保证金帐户"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__FROZEN:
		return "冻结帐户"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB:
		return "平安现金子帐户"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__OIL:
		return "油品子账户"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__PAB_WITNESS:
		return "见证宝现金子帐户"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__TRANSIT:
		return "在途子账户"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__PRIVATE:
		return "私有子账户"
	case CASHDESK_PAY_SUB_ACCOUNT_TYPE__GLP:
		return "普洛斯子账户"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPaySubAccountType)(nil)

func (v CashdeskPaySubAccountType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPaySubAccountType
	}
	return []byte(str), nil
}

func (v *CashdeskPaySubAccountType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPaySubAccountTypeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPayTransMode = errors.New("invalid CashdeskPayTransMode")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPayTransMode", map[string]string{
		"INTERMEDIARY":      "中介交易",
		"DIRECT_PAY":        "直付交易",
		"PREPAY":            "预付交易",
		"INSURANCE_FINANCE": "保理交易",
	})
}

func ParseCashdeskPayTransModeFromString(s string) (CashdeskPayTransMode, error) {
	switch s {
	case "":
		return CASHDESK_PAY_TRANS_MODE_UNKNOWN, nil
	case "INTERMEDIARY":
		return CASHDESK_PAY_TRANS_MODE__INTERMEDIARY, nil
	case "DIRECT_PAY":
		return CASHDESK_PAY_TRANS_MODE__DIRECT_PAY, nil
	case "PREPAY":
		return CASHDESK_PAY_TRANS_MODE__PREPAY, nil
	case "INSURANCE_FINANCE":
		return CASHDESK_PAY_TRANS_MODE__INSURANCE_FINANCE, nil
	}
	return CASHDESK_PAY_TRANS_MODE_UNKNOWN, InvalidCashdeskPayTransMode
}

func ParseCashdeskPayTransModeFromLabelString(s string) (CashdeskPayTransMode, error) {
	switch s {
	case "":
		return CASHDESK_PAY_TRANS_MODE_UNKNOWN, nil
	case "中介交易":
		return CASHDESK_PAY_TRANS_MODE__INTERMEDIARY, nil
	case "直付交易":
		return CASHDESK_PAY_TRANS_MODE__DIRECT_PAY, nil
	case "预付交易":
		return CASHDESK_PAY_TRANS_MODE__PREPAY, nil
	case "保理交易":
		return CASHDESK_PAY_TRANS_MODE__INSURANCE_FINANCE, nil
	}
	return CASHDESK_PAY_TRANS_MODE_UNKNOWN, InvalidCashdeskPayTransMode
}

func (CashdeskPayTransMode) EnumType() string {
	return "CashdeskPayTransMode"
}

func (CashdeskPayTransMode) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAY_TRANS_MODE__INTERMEDIARY):      {"INTERMEDIARY", "中介交易"},
		int(CASHDESK_PAY_TRANS_MODE__DIRECT_PAY):        {"DIRECT_PAY", "直付交易"},
		int(CASHDESK_PAY_TRANS_MODE__PREPAY):            {"PREPAY", "预付交易"},
		int(CASHDESK_PAY_TRANS_MODE__INSURANCE_FINANCE): {"INSURANCE_FINANCE", "保理交易"},
	}
}
func (v CashdeskPayTransMode) String() string {
	switch v {
	case CASHDESK_PAY_TRANS_MODE_UNKNOWN:
		return ""
	case CASHDESK_PAY_TRANS_MODE__INTERMEDIARY:
		return "INTERMEDIARY"
	case CASHDESK_PAY_TRANS_MODE__DIRECT_PAY:
		return "DIRECT_PAY"
	case CASHDESK_PAY_TRANS_MODE__PREPAY:
		return "PREPAY"
	case CASHDESK_PAY_TRANS_MODE__INSURANCE_FINANCE:
		return "INSURANCE_FINANCE"
	}
	return "UNKNOWN"
}

func (v CashdeskPayTransMode) Label() string {
	switch v {
	case CASHDESK_PAY_TRANS_MODE_UNKNOWN:
		return ""
	case CASHDESK_PAY_TRANS_MODE__INTERMEDIARY:
		return "中介交易"
	case CASHDESK_PAY_TRANS_MODE__DIRECT_PAY:
		return "直付交易"
	case CASHDESK_PAY_TRANS_MODE__PREPAY:
		return "预付交易"
	case CASHDESK_PAY_TRANS_MODE__INSURANCE_FINANCE:
		return "保理交易"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPayTransMode)(nil)

func (v CashdeskPayTransMode) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPayTransMode
	}
	return []byte(str), nil
}

func (v *CashdeskPayTransMode) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPayTransModeFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPayTransState = errors.New("invalid CashdeskPayTransState")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPayTransState", map[string]string{
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

func ParseCashdeskPayTransStateFromString(s string) (CashdeskPayTransState, error) {
	switch s {
	case "":
		return CASHDESK_PAY_TRANS_STATE_UNKNOWN, nil
	case "TO_PAY":
		return CASHDESK_PAY_TRANS_STATE__TO_PAY, nil
	case "PAY_OK":
		return CASHDESK_PAY_TRANS_STATE__PAY_OK, nil
	case "TRANS_OK":
		return CASHDESK_PAY_TRANS_STATE__TRANS_OK, nil
	case "REFUND_OK":
		return CASHDESK_PAY_TRANS_STATE__REFUND_OK, nil
	case "CANCEL":
		return CASHDESK_PAY_TRANS_STATE__CANCEL, nil
	case "STOP":
		return CASHDESK_PAY_TRANS_STATE__STOP, nil
	case "PART_REFUND":
		return CASHDESK_PAY_TRANS_STATE__PART_REFUND, nil
	case "PRE_PAY":
		return CASHDESK_PAY_TRANS_STATE__PRE_PAY, nil
	}
	return CASHDESK_PAY_TRANS_STATE_UNKNOWN, InvalidCashdeskPayTransState
}

func ParseCashdeskPayTransStateFromLabelString(s string) (CashdeskPayTransState, error) {
	switch s {
	case "":
		return CASHDESK_PAY_TRANS_STATE_UNKNOWN, nil
	case "等待付款":
		return CASHDESK_PAY_TRANS_STATE__TO_PAY, nil
	case "付款成功":
		return CASHDESK_PAY_TRANS_STATE__PAY_OK, nil
	case "交易成功":
		return CASHDESK_PAY_TRANS_STATE__TRANS_OK, nil
	case "退款成功":
		return CASHDESK_PAY_TRANS_STATE__REFUND_OK, nil
	case "交易取消":
		return CASHDESK_PAY_TRANS_STATE__CANCEL, nil
	case "交易终止":
		return CASHDESK_PAY_TRANS_STATE__STOP, nil
	case "部分退款":
		return CASHDESK_PAY_TRANS_STATE__PART_REFUND, nil
	case "预支付完成":
		return CASHDESK_PAY_TRANS_STATE__PRE_PAY, nil
	}
	return CASHDESK_PAY_TRANS_STATE_UNKNOWN, InvalidCashdeskPayTransState
}

func (CashdeskPayTransState) EnumType() string {
	return "CashdeskPayTransState"
}

func (CashdeskPayTransState) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAY_TRANS_STATE__TO_PAY):      {"TO_PAY", "等待付款"},
		int(CASHDESK_PAY_TRANS_STATE__PAY_OK):      {"PAY_OK", "付款成功"},
		int(CASHDESK_PAY_TRANS_STATE__TRANS_OK):    {"TRANS_OK", "交易成功"},
		int(CASHDESK_PAY_TRANS_STATE__REFUND_OK):   {"REFUND_OK", "退款成功"},
		int(CASHDESK_PAY_TRANS_STATE__CANCEL):      {"CANCEL", "交易取消"},
		int(CASHDESK_PAY_TRANS_STATE__STOP):        {"STOP", "交易终止"},
		int(CASHDESK_PAY_TRANS_STATE__PART_REFUND): {"PART_REFUND", "部分退款"},
		int(CASHDESK_PAY_TRANS_STATE__PRE_PAY):     {"PRE_PAY", "预支付完成"},
	}
}
func (v CashdeskPayTransState) String() string {
	switch v {
	case CASHDESK_PAY_TRANS_STATE_UNKNOWN:
		return ""
	case CASHDESK_PAY_TRANS_STATE__TO_PAY:
		return "TO_PAY"
	case CASHDESK_PAY_TRANS_STATE__PAY_OK:
		return "PAY_OK"
	case CASHDESK_PAY_TRANS_STATE__TRANS_OK:
		return "TRANS_OK"
	case CASHDESK_PAY_TRANS_STATE__REFUND_OK:
		return "REFUND_OK"
	case CASHDESK_PAY_TRANS_STATE__CANCEL:
		return "CANCEL"
	case CASHDESK_PAY_TRANS_STATE__STOP:
		return "STOP"
	case CASHDESK_PAY_TRANS_STATE__PART_REFUND:
		return "PART_REFUND"
	case CASHDESK_PAY_TRANS_STATE__PRE_PAY:
		return "PRE_PAY"
	}
	return "UNKNOWN"
}

func (v CashdeskPayTransState) Label() string {
	switch v {
	case CASHDESK_PAY_TRANS_STATE_UNKNOWN:
		return ""
	case CASHDESK_PAY_TRANS_STATE__TO_PAY:
		return "等待付款"
	case CASHDESK_PAY_TRANS_STATE__PAY_OK:
		return "付款成功"
	case CASHDESK_PAY_TRANS_STATE__TRANS_OK:
		return "交易成功"
	case CASHDESK_PAY_TRANS_STATE__REFUND_OK:
		return "退款成功"
	case CASHDESK_PAY_TRANS_STATE__CANCEL:
		return "交易取消"
	case CASHDESK_PAY_TRANS_STATE__STOP:
		return "交易终止"
	case CASHDESK_PAY_TRANS_STATE__PART_REFUND:
		return "部分退款"
	case CASHDESK_PAY_TRANS_STATE__PRE_PAY:
		return "预支付完成"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPayTransState)(nil)

func (v CashdeskPayTransState) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPayTransState
	}
	return []byte(str), nil
}

func (v *CashdeskPayTransState) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPayTransStateFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidCashdeskPayTransType = errors.New("invalid CashdeskPayTransType")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskPayTransType", map[string]string{
		"ETC":          "ETC",
		"OIL_CARD":     "油卡",
		"TENDER":       "招采",
		"MALL":         "商城",
		"ZQX":          "中启行(油)",
		"SKID_MOUNTED": "撬装(油)",
		"MAINTENANCE":  "维保",
	})
}

func ParseCashdeskPayTransTypeFromString(s string) (CashdeskPayTransType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_TRANS_TYPE_UNKNOWN, nil
	case "ETC":
		return CASHDESK_PAY_TRANS_TYPE__ETC, nil
	case "OIL_CARD":
		return CASHDESK_PAY_TRANS_TYPE__OIL_CARD, nil
	case "TENDER":
		return CASHDESK_PAY_TRANS_TYPE__TENDER, nil
	case "MALL":
		return CASHDESK_PAY_TRANS_TYPE__MALL, nil
	case "ZQX":
		return CASHDESK_PAY_TRANS_TYPE__ZQX, nil
	case "SKID_MOUNTED":
		return CASHDESK_PAY_TRANS_TYPE__SKID_MOUNTED, nil
	case "MAINTENANCE":
		return CASHDESK_PAY_TRANS_TYPE__MAINTENANCE, nil
	}
	return CASHDESK_PAY_TRANS_TYPE_UNKNOWN, InvalidCashdeskPayTransType
}

func ParseCashdeskPayTransTypeFromLabelString(s string) (CashdeskPayTransType, error) {
	switch s {
	case "":
		return CASHDESK_PAY_TRANS_TYPE_UNKNOWN, nil
	case "ETC":
		return CASHDESK_PAY_TRANS_TYPE__ETC, nil
	case "油卡":
		return CASHDESK_PAY_TRANS_TYPE__OIL_CARD, nil
	case "招采":
		return CASHDESK_PAY_TRANS_TYPE__TENDER, nil
	case "商城":
		return CASHDESK_PAY_TRANS_TYPE__MALL, nil
	case "中启行(油)":
		return CASHDESK_PAY_TRANS_TYPE__ZQX, nil
	case "撬装(油)":
		return CASHDESK_PAY_TRANS_TYPE__SKID_MOUNTED, nil
	case "维保":
		return CASHDESK_PAY_TRANS_TYPE__MAINTENANCE, nil
	}
	return CASHDESK_PAY_TRANS_TYPE_UNKNOWN, InvalidCashdeskPayTransType
}

func (CashdeskPayTransType) EnumType() string {
	return "CashdeskPayTransType"
}

func (CashdeskPayTransType) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PAY_TRANS_TYPE__ETC):          {"ETC", "ETC"},
		int(CASHDESK_PAY_TRANS_TYPE__OIL_CARD):     {"OIL_CARD", "油卡"},
		int(CASHDESK_PAY_TRANS_TYPE__TENDER):       {"TENDER", "招采"},
		int(CASHDESK_PAY_TRANS_TYPE__MALL):         {"MALL", "商城"},
		int(CASHDESK_PAY_TRANS_TYPE__ZQX):          {"ZQX", "中启行(油)"},
		int(CASHDESK_PAY_TRANS_TYPE__SKID_MOUNTED): {"SKID_MOUNTED", "撬装(油)"},
		int(CASHDESK_PAY_TRANS_TYPE__MAINTENANCE):  {"MAINTENANCE", "维保"},
	}
}
func (v CashdeskPayTransType) String() string {
	switch v {
	case CASHDESK_PAY_TRANS_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_TRANS_TYPE__ETC:
		return "ETC"
	case CASHDESK_PAY_TRANS_TYPE__OIL_CARD:
		return "OIL_CARD"
	case CASHDESK_PAY_TRANS_TYPE__TENDER:
		return "TENDER"
	case CASHDESK_PAY_TRANS_TYPE__MALL:
		return "MALL"
	case CASHDESK_PAY_TRANS_TYPE__ZQX:
		return "ZQX"
	case CASHDESK_PAY_TRANS_TYPE__SKID_MOUNTED:
		return "SKID_MOUNTED"
	case CASHDESK_PAY_TRANS_TYPE__MAINTENANCE:
		return "MAINTENANCE"
	}
	return "UNKNOWN"
}

func (v CashdeskPayTransType) Label() string {
	switch v {
	case CASHDESK_PAY_TRANS_TYPE_UNKNOWN:
		return ""
	case CASHDESK_PAY_TRANS_TYPE__ETC:
		return "ETC"
	case CASHDESK_PAY_TRANS_TYPE__OIL_CARD:
		return "油卡"
	case CASHDESK_PAY_TRANS_TYPE__TENDER:
		return "招采"
	case CASHDESK_PAY_TRANS_TYPE__MALL:
		return "商城"
	case CASHDESK_PAY_TRANS_TYPE__ZQX:
		return "中启行(油)"
	case CASHDESK_PAY_TRANS_TYPE__SKID_MOUNTED:
		return "撬装(油)"
	case CASHDESK_PAY_TRANS_TYPE__MAINTENANCE:
		return "维保"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CashdeskPayTransType)(nil)

func (v CashdeskPayTransType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCashdeskPayTransType
	}
	return []byte(str), nil
}

func (v *CashdeskPayTransType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCashdeskPayTransTypeFromString(string(bytes.ToUpper(data)))
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

func (CashdeskPlatformType) EnumType() string {
	return "CashdeskPlatformType"
}

func (CashdeskPlatformType) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_PLATFORM_TYPE__COMPUTER):     {"COMPUTER", "个人电脑"},
		int(CASHDESK_PLATFORM_TYPE__MULTI_MEDIA):  {"MULTI_MEDIA", "多媒体终端"},
		int(CASHDESK_PLATFORM_TYPE__MOBILE_PHONE): {"MOBILE_PHONE", "手持智能设备（手机、MID等）"},
		int(CASHDESK_PLATFORM_TYPE__PAD):          {"PAD", "平板电脑"},
		int(CASHDESK_PLATFORM_TYPE__POS):          {"POS", "POS终端"},
		int(CASHDESK_PLATFORM_TYPE__MERCHANT):     {"MERCHANT", "商户系统"},
		int(CASHDESK_PLATFORM_TYPE__STB):          {"STB", "数字机顶盒"},
		int(CASHDESK_PLATFORM_TYPE__TV):           {"TV", "智能电视"},
		int(CASHDESK_PLATFORM_TYPE__VEM):          {"VEM", "自动柜员机（售货机等）"},
		int(CASHDESK_PLATFORM_TYPE__THIRD_SYSTEM): {"THIRD_SYSTEM", "第三方机构系统"},
	}
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

var InvalidCashdeskSignAlgorithm = errors.New("invalid CashdeskSignAlgorithm")

func init() {
	golib_tools_courier_enumeration.RegisterEnums("CashdeskSignAlgorithm", map[string]string{
		"MD5": "MD5加密",
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
	case "MD5加密":
		return CASHDESK_SIGN_ALGORITHM__MD5, nil
	}
	return CASHDESK_SIGN_ALGORITHM_UNKNOWN, InvalidCashdeskSignAlgorithm
}

func (CashdeskSignAlgorithm) EnumType() string {
	return "CashdeskSignAlgorithm"
}

func (CashdeskSignAlgorithm) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_SIGN_ALGORITHM__MD5): {"MD5", "MD5加密"},
	}
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
		return "MD5加密"
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

func (CashdeskTransState) EnumType() string {
	return "CashdeskTransState"
}

func (CashdeskTransState) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_TRANS_STATE__TO_PAY):      {"TO_PAY", "等待付款"},
		int(CASHDESK_TRANS_STATE__PAY_OK):      {"PAY_OK", "付款成功"},
		int(CASHDESK_TRANS_STATE__TRANS_OK):    {"TRANS_OK", "交易成功"},
		int(CASHDESK_TRANS_STATE__REFUND_OK):   {"REFUND_OK", "退款成功"},
		int(CASHDESK_TRANS_STATE__CANCEL):      {"CANCEL", "交易取消"},
		int(CASHDESK_TRANS_STATE__STOP):        {"STOP", "交易终止"},
		int(CASHDESK_TRANS_STATE__PART_REFUND): {"PART_REFUND", "部分退款"},
		int(CASHDESK_TRANS_STATE__PRE_PAY):     {"PRE_PAY", "预支付完成"},
	}
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

func (CashdeskTransType) EnumType() string {
	return "CashdeskTransType"
}

func (CashdeskTransType) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_TRANS_TYPE__ETC):          {"ETC", "ETC"},
		int(CASHDESK_TRANS_TYPE__OIL_CARD):     {"OIL_CARD", "油卡"},
		int(CASHDESK_TRANS_TYPE__TENDER):       {"TENDER", "招采"},
		int(CASHDESK_TRANS_TYPE__MALL):         {"MALL", "商城"},
		int(CASHDESK_TRANS_TYPE__ZQX):          {"ZQX", "中启行(油)"},
		int(CASHDESK_TRANS_TYPE__SKID_MOUNTED): {"SKID_MOUNTED", "撬装(油)"},
		int(CASHDESK_TRANS_TYPE__MAINTENANCE):  {"MAINTENANCE", "维保"},
	}
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

func (CashdeskUserFrom) EnumType() string {
	return "CashdeskUserFrom"
}

func (CashdeskUserFrom) Enums() map[int][]string {
	return map[int][]string{
		int(CASHDESK_USER_FROM__G7):        {"G7", "G7平台"},
		int(CASHDESK_USER_FROM__ANONYMOUS): {"ANONYMOUS", "匿名"},
		int(CASHDESK_USER_FROM__G7_ORG):    {"G7_ORG", "G7机构"},
		int(CASHDESK_USER_FROM__WECHAT):    {"WECHAT", "微信"},
	}
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
