package router

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	PostForm = iota // 0
	Param           // 1
	Query           // 2
)

func intParamCheck(cxt *gin.Context, ptype int8, pname string, num *int64) bool {
	value := ""
	switch ptype {
	case PostForm:
		value = cxt.PostForm(pname)
	case Param:
		value = cxt.Param(pname)
	case Query:
		value = cxt.Query(pname)
	default:
		return false
	}
	if value == "" {
		return false
	}
	if result, e := strconv.ParseInt(value, 10, 64); e == nil {
		*num = result
		return true
	}
	return false
}

func SetUpRouter(ser *gin.Engine) {
	userRouter(ser)
	postRouter(ser)
}
