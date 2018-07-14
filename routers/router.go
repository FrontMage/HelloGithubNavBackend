package routers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"

	"log"

	"github.com/FrontMage/HelloGithubNavBackend/dao"
	"github.com/FrontMage/HelloGithubNavBackend/dao/content"
	"github.com/gin-gonic/gin"
)

// GraphqlParams graphql params
type GraphqlParams struct {
	Query         string `json:"query,omitempty" form:"query"`
	OperationName string `json:"operationName,omitempty" form:"operationName"`
}

// MountRouters 挂载路由
func MountRouters(r *gin.Engine) {
	c := r.Group("/content")
	{
		// TODO: doc this
		c.GET("/", func(ctx *gin.Context) {
			p := &GraphqlParams{}
			if err := ctx.Bind(p); err != nil {
				log.Printf("Bind params err=%s", err.Error())
				ctx.JSON(http.StatusOK, gin.H{
					"code":   400,
					"errors": []string{fmt.Sprintf("Invalid params %s", err.Error())},
				})
				return
			}
			graphqlParams := graphql.Params{
				Schema:        content.Schema,
				RequestString: p.Query,
				OperationName: p.OperationName,
			}
			result := graphql.Do(graphqlParams)
			if len(result.Errors) > 0 {
				errors := []string{}
				for _, e := range result.Errors {
					log.Printf("Query err=%s", e.Error())
					errors = append(errors, fmt.Sprintf("Query err: %s", e.Error()))
				}
				ctx.JSON(http.StatusOK, gin.H{
					"code":    500,
					"message": errors,
				})
				return
			}
			ctx.JSON(http.StatusOK, result)
		})
		// TODO: doc this
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
