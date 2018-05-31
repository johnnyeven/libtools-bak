package httplib

import (
	"io/ioutil"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"

	"golib/tools/courier/status_error"
	"golib/tools/courier/transport_http/transform"
	"golib/tools/httplib/testify"
	"golib/tools/timelib"
	"golib/tools/validate"
)

func v_ok(v interface{}) (bool, string) {
	return v.(bool), "not ok"
}

func v_data(v interface{}) (bool, string) {
	return v.(int32) == int32(1), "invalid data"
}

func init() {
	validate.AddValidateFunc("v_ok", v_ok)
	validate.AddValidateFunc("v_data", v_data)
}

type SimpleModel struct {
	Data int32 `json:"data" validate:"v_data"`
}

type SimpleReqBodyBasic struct {
	Name string `json:"name" default:""`
	ID   uint64 `json:"id"`
}

type SimpleReqBody struct {
	SimpleReqBodyBasic
	Price  float32                 `json:"price" default:"0.1"`
	OK     bool                    `json:"ok" validate:"v_ok" default:"false"`
	Models []SimpleModel           `json:"models"`
	V      int32                   `json:"v" default:"1"`
	D      string                  `json:"d" default:"abc"`
	G      uint32                  `json:"g" default:"2"`
	Time   *timelib.MySQLTimestamp `json:"time" default:""`
}

type CommonReq struct {
	ObjID   uint64 `name:"obj_id" in:"path"`
	ObjName string `name:"obj_name" in:"query"`
}

type Ctx struct {
	C string
}

type SimpleReq struct {
	CommonReq
	Ctx        *Ctx          `name:"ctx" in:"context"`
	Body       SimpleReqBody `in:"body"`
	ID         uint          `name:"id" in:"query"`
	IntSlice   []uint64      `name:"int_slice" in:"query" default:""`
	StrSlice   []string      `name:"str_slice" in:"query" default:""`
	FloatSlice []float64     `name:"float_slice" in:"query" default:""`
	ObjFloat   float32       `name:"obj_float" in:"path"`
	ObjBool    bool          `name:"obj_bool" in:"path"`
	HID        uint          `name:"h_id" in:"header" default:"0"`
	HName      string        `name:"h_name" in:"header" default:""`
	HFloat     float64       `name:"h_float" in:"header" default:"0.0"`
	HBool      bool          `name:"h_bool" in:"header" default:"false"`
}

func TestTrans(t *testing.T) {
	tt := assert.New(t)

	expectedReq := SimpleReq{
		Ctx: &Ctx{
			C: "12123",
		},
		Body: SimpleReqBody{
			SimpleReqBodyBasic: SimpleReqBodyBasic{
				Name: "a",
				ID:   1,
			},
			Price: 0.2,
			OK:    true,
			V:     1,
			D:     "abc",
			G:     2,
			Models: []SimpleModel{
				{
					Data: 1,
				},
				{
					Data: 1,
				},
				{
					Data: 1,
				},
			},
		},
		CommonReq: CommonReq{
			ObjID:   123,
			ObjName: "aaa",
		},
		ID:         1,
		IntSlice:   []uint64{1, 2, 3},
		StrSlice:   []string{"a", "b", "c"},
		FloatSlice: []float64{1.1, 1.2, 1.3},
		ObjFloat:   12.2,
		ObjBool:    true,
		HID:        1,
		HName:      "bc",
		HBool:      true,
		HFloat:     1.2,
	}

	{
		c := testify.NewContext("POST", "/:obj_id/:obj_float/:obj_bool/", expectedReq)

		req := SimpleReq{}
		err := TransToReq(c, &req)
		tt.NoError(err)
		tt.Equal(expectedReq, req)
	}

	{
		expectedReq.Body.ID = 0
		expectedReq.Body.Models[1].Data = 2
		expectedReq.ObjFloat = 0
		expectedReq.Body.OK = false

		c := testify.NewContext("POST", "/:obj_id/:obj_float/:obj_bool/", expectedReq)

		req := SimpleReq{}
		err := TransToReq(c, &req)

		tt.Equal(status_error.ErrorFields{
			{
				Field: "id",
				Msg:   "缺失必填字段",
				In:    "body",
			},
			{
				Field: "models[1].data",
				Msg:   "invalid data",
				In:    "body",
			},
			{
				Field: "obj_float",
				Msg:   "缺失必填字段",
				In:    "path",
			},
			{
				Field: "ok",
				Msg:   "not ok",
				In:    "body",
			},
		}, status_error.FromError(err).ErrorFields.Sort())
	}
}

type Counter struct {
	Total int32 `json:"total" in:"query" default:"0"`
}

func (c Counter) ValidateTotal() string {
	if c.Total > 1 {
		return "total too large"
	}
	return ""
}

type Req struct {
	Counter
	Pager
	CreateTimeRange
	Body Counter `in:"body"`
}

func (r Req) ValidateCreateStartTime() string {
	return "CreateStartTime Error"
}

func TestTransToReqWithValidateHook(t *testing.T) {
	tt := assert.New(t)

	expectedReq := Req{}
	expectedReq.Total = 4
	expectedReq.Body.Total = 4
	expectedReq.CreateStartTime, _ = timelib.ParseMySQLTimestampFromString("2017-01-01T20:00:00.000Z")
	expectedReq.CreateEndTime, _ = timelib.ParseMySQLTimestampFromString("2017-01-01T09:00:00.000Z")

	c := testify.NewContext("POST", "/", expectedReq)

	reqWithReqGroupForValid := Req{}

	err := TransToReq(c, &reqWithReqGroupForValid)

	if err != nil {
		tt.Equal(status_error.ErrorFields{
			{
				Field: "createEndTime",
				Msg:   "终止时间不得小于开始时间",
				In:    "query",
			},
			{
				Field: "createStartTime",
				Msg:   "CreateStartTime Error",
				In:    "query",
			},
			{
				Field: "total",
				Msg:   "total too large",
				In:    "query",
			},
			{
				Field: "total",
				Msg:   "total too large",
				In:    "body",
			},
		}, status_error.FromError(err).ErrorFields.Sort())
	}
}

type SingleFormFile struct {
	FormData struct {
		SingleFile *multipart.FileHeader `name:"singleFile"`
	} `in:"formData,multipart"`
}

func TestTransSingleFormFile(t *testing.T) {
	tt := assert.New(t)

	req := SingleFormFile{}
	req.FormData.SingleFile, _ = transform.NewFileHeader("singleFile", "SingleFile", []byte("1,2,3,4"))

	c := testify.NewContext("POST", "/", req)
	r := SingleFormFile{}
	err := TransToReq(c, &r)
	tt.Nil(err)

	actualFile, _ := r.FormData.SingleFile.Open()
	actualBytes, _ := ioutil.ReadAll(actualFile)
	tt.Equal([]byte("1,2,3,4"), actualBytes)
}

type MultiFormFile struct {
	FormData struct {
		FirstFile  *multipart.FileHeader `json:"firstFile"`
		SecondFile *multipart.FileHeader `json:"secondFile"`
	} `in:"formData,multipart"`
}

func TestTransMultiFormFile(t *testing.T) {
	tt := assert.New(t)

	req := MultiFormFile{}
	req.FormData.FirstFile, _ = transform.NewFileHeader("firstFile", "SingleFile", []byte("1"))
	req.FormData.SecondFile, _ = transform.NewFileHeader("secondFile", "SecondFile", []byte("2"))

	c := testify.NewContext("POST", "/", req)
	r := MultiFormFile{}
	err := TransToReq(c, &r)
	tt.Nil(err)

	{
		actualFile, _ := r.FormData.FirstFile.Open()
		actualBytes, _ := ioutil.ReadAll(actualFile)
		tt.Equal([]byte("1"), actualBytes)
	}

	{
		actualFile, _ := r.FormData.SecondFile.Open()
		actualBytes, _ := ioutil.ReadAll(actualFile)
		tt.Equal([]byte("2"), actualBytes)
	}
}
