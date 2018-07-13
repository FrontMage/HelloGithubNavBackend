package routers

import (
	"net/http"
	"strconv"

	"github.com/FrontMage/HelloGithubNavBackend/dao"
	"github.com/FrontMage/HelloGithubNavBackend/dao/content"
	"github.com/gin-gonic/gin"
)

// MountRouters 挂载路由
func MountRouters(r *gin.Engine) {
	c := r.Group("/content")
	{
		c.GET("/:id", func(ctx *gin.Context) {
			id, exists := ctx.Params.Get("id")
			if !exists {
				// TODO: return error
				return
			}
			idNum, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				// TODO: return error
				return
			}
			content, err := content.Get(idNum)
			switch err {
			case dao.ErrRecordNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": 404,
				})
			case nil:
				ctx.JSON(http.StatusOK, content)
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			}
		})
	}
}
