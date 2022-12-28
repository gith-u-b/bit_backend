package v1

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/sai1024/bit_backend/models"
	"github.com/sai1024/bit_backend/pkg/e"
	"github.com/sai1024/bit_backend/pkg/setting"
	"github.com/sai1024/bit_backend/pkg/util"
	"github.com/unknwon/com"
)

func GetRedeemCodes(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["lists"] = models.GetRedeemCodes(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetRedeemCodeTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddRedeemCode(c *gin.Context) {
	number := c.PostForm("number")
	comment := c.PostForm("comment")

	valid := validation.Validation{}
	valid.Required(number, "number").Message("生成数量不能为空")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 1; i++ {
		models.AddRedeemCode(randomString(8), comment)
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func DeleteRedeemCode(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		models.DeleteRedeemCode(id)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
