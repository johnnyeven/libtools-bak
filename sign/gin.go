package sign

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"

	"profzone/libtools/courier/status_error"
	"profzone/libtools/httplib"
)

func WithSignBy(exchangeSecret SecretExchanger) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := SignParams{}
		err := httplib.TransToReq(c, &req)
		if err != nil {
			httplib.WriteError(c, err)
			c.Abort()
			return
		}

		query := c.Request.URL.Query()

		expectSign, origin, err := getSign(c.Request, query, exchangeSecret)
		if err != nil {
			httplib.WriteError(c, err)
			c.Abort()
			return
		}

		if req.Sign != string(expectSign) {
			errForSign := status_error.SignFailed.StatusError().
				WithDesc(fmt.Sprintf("Origin %s&secret=***, Sign:[%s], Expected:[%s]", bytes.Split(origin, []byte("secret"))[0], req.Sign, expectSign))
			httplib.WriteError(c, errForSign)
			c.Abort()
			return
		}
	}
}
