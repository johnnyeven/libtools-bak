package client_demo

import (
	"fmt"
	mime_multipart "mime/multipart"

	golib_tools_courier "golib/tools/courier"
	golib_tools_courier_status_error "golib/tools/courier/status_error"

	golib_tools_courier_client "golib/tools/courier/client"
)

type ClientDemo struct {
	golib_tools_courier_client.Client
}

func (c ClientDemo) MarshalDefaults(v interface{}) {
	if cl, ok := v.(*ClientDemo); ok {
		cl.Name = "demo"
		cl.Client.MarshalDefaults(&cl.Client)
	}
}

func (c ClientDemo) Init() {
	c.CheckService()
}

func (c ClientDemo) CheckService() {
	err := c.Request(c.Name+".Check", "HEAD", "/", nil).
		Do().
		Into(nil)
	statusErr := golib_tools_courier_status_error.FromError(err)
	if statusErr.Code == int64(golib_tools_courier_status_error.RequestTimeout) {
		panic(fmt.Errorf("service %s have some error %s", c.Name, statusErr))
	}
}

type CreateRequest struct {
	//
	Body Data `fmt:"json" in:"body"`
}

func (c ClientDemo) Create(req CreateRequest) (resp *CreateResponse, err error) {
	resp = &CreateResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".Create", "POST", "/demo/crud/", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type CreateResponse struct {
	Meta golib_tools_courier.Metadata
	Body []byte
}

func (c ClientDemo) FileDownload() (resp *FileDownloadResponse, err error) {
	resp = &FileDownloadResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".FileDownload", "GET", "/demo/files", nil).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type FileDownloadResponse struct {
	Meta golib_tools_courier.Metadata
	Body []byte
}

type FormMultipartWithFileRequest struct {
	//
	Body struct {
		//
		Data Data `default:"" json:"data"`
		//
		File *mime_multipart.FileHeader `json:"file"`
		//
		Slice []string `default:"" json:"slice"`
		//
		String string `default:"" json:"string"`
	} `in:"formData,multipart"`
}

func (c ClientDemo) FormMultipartWithFile(req FormMultipartWithFileRequest) (resp *FormMultipartWithFileResponse, err error) {
	resp = &FormMultipartWithFileResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".FormMultipartWithFile", "POST", "/demo/forms/multipart", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type FormMultipartWithFileResponse struct {
	Meta golib_tools_courier.Metadata
	Body []byte
}

type FormMultipartWithFilesRequest struct {
	//
	Body struct {
		//
		Files []*mime_multipart.FileHeader `json:"files"`
	} `in:"formData,multipart"`
}

func (c ClientDemo) FormMultipartWithFiles(req FormMultipartWithFilesRequest) (resp *FormMultipartWithFilesResponse, err error) {
	resp = &FormMultipartWithFilesResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".FormMultipartWithFiles", "POST", "/demo/forms/multipart-with-files", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type FormMultipartWithFilesResponse struct {
	Meta golib_tools_courier.Metadata
	Body []byte
}

type FormURLEncodedRequest struct {
	//
	Body struct {
		//
		Data Data `json:"data"`
		//
		Slice []string `json:"slice"`
		//
		String string `json:"string"`
	} `in:"formData"`
}

func (c ClientDemo) FormURLEncoded(req FormURLEncodedRequest) (resp *FormURLEncodedResponse, err error) {
	resp = &FormURLEncodedResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".FormURLEncoded", "POST", "/demo/forms/url-encoded", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type FormURLEncodedResponse struct {
	Meta golib_tools_courier.Metadata
	Body []byte
}

type GetByIDRequest struct {
	//
	Status ResourceStatus `in:"query" name:"status"`
	//
	ID string `in:"path" name:"id"`
	//
	Name string `default:"" in:"query" name:"name"`
	//
	Label []string `default:"" in:"query" name:"label"`
}

func (c ClientDemo) GetByID(req GetByIDRequest) (resp *GetByIDResponse, err error) {
	resp = &GetByIDResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".GetByID", "GET", "/demo/crud/:id", req).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetByIDResponse struct {
	Meta golib_tools_courier.Metadata
	Body GetByID
}

func (c ClientDemo) Swagger() (resp *SwaggerResponse, err error) {
	resp = &SwaggerResponse{}
	resp.Meta = golib_tools_courier.Metadata{}

	err = c.Request(c.Name+".Swagger", "GET", "/demo", nil).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type SwaggerResponse struct {
	Meta golib_tools_courier.Metadata
	Body JSONBytes
}
