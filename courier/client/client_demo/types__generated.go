package client_demo

import (
	golib_tools_courier_transport_http "golib/tools/courier/transport_http"
)

type Data struct {
	//
	Label string `json:"label"`
	//
	Name string `json:"name"`
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

type File = golib_tools_courier_transport_http.File

type GetByID struct {
	//
	ID string `json:"id"`
	//
	Label []string `default:"" json:"label"`
	//
	Name string `default:"" json:"name"`
	//
	Status ResourceStatus `json:"status"`
}

type JSONBytes []uint8

type ResourceStatus = DemoResourceStatus

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
