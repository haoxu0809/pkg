package render

import (
	"github.com/haoxu0809/pkg/errors"

	"github.com/gin-gonic/gin"
)

func Serializer(ctx *gin.Context, err error, data any) {
	var (
		httpCode = 200
		obj      = gin.H{"code": 100000, "data": data, "msg": "Success"}
	)

	if err != nil {
		coder := errors.ParseCoder(err)
		obj["code"] = coder.Code()
		obj["data"] = nil
		obj["msg"] = coder.String()
		obj["reference"] = coder.Reference()

		httpCode = coder.HTTPStatus()
	}

	ctx.JSON(httpCode, obj)
}
