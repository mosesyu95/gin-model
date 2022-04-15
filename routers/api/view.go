package api

import "github.com/gin-gonic/gin"

// Info godoc
// @Summary      默认页面
// @Description  默认的页面
// @Tags         info
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  ""
// @Failure      400  {string}  string  "ok"
// @Failure      404  {string}  string  "ok"
// @Failure      500  {string}  string  "ok"
// @Router       /info [get]
func Info(c *gin.Context) {
	c.JSON(200, "欢迎使用gin模板")
}
