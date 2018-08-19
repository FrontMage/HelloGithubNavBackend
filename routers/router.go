package routers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/FrontMage/HelloGithubNavBackend/dao/content"
	"github.com/gin-gonic/gin"
)

// GraphqlParams graphql params
type GraphqlParams struct {
	Query         string `json:"query,omitempty" form:"query"`
	OperationName string `json:"operationName,omitempty" form:"operationName"`
}

// Content single project content
type Content struct {
	Category    string `json:"category,omitempty"`
	CategoryID  uint64 `json:"categoryID,omitempty"`
	Description string `json:"description,omitempty"`
	ID          uint64 `json:"id,omitempty"`
	ImagePath   string `json:"imagePath,omitempty"`
	ProjectURL  string `json:"projectURL,omitempty"`
	Title       string `json:"title,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
}

// RecommendCategory recommend project by category id
type RecommendCategory struct {
	Category   string     `json:"category,omitempty"`
	CategoryID uint64     `json:"categoryID,omitempty"`
	Contents   []*Content `json:"contents,omitempty"`
}

// MountRouters 挂载路由
func MountRouters(r *gin.Engine) {
	c := r.Group("/content")
	{
		c.GET("/:id", func(ctx *gin.Context) {
			id, exists := ctx.Params.Get("id")
			if !exists {
				log.Printf("Invalid params, missing `id` field \n")
				ctx.JSON(http.StatusOK, gin.H{
					"code": 400,
					"msg":  "Invalid params, missing `id` field",
				})
				return
			}
			log.Printf("Got id=%s \n", id)
			idInt, _ := strconv.ParseUint(id, 10, 64)
			if c, e := content.Get(idInt); e != nil {
				fmt.Printf("Failed to get content with id=%s, err=%s\n", id, e.Error())
				ctx.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  "Internal Server Error",
				})
			} else {
				res := Content{
					ID:          c.ID,
					Category:    c.Category.Name,
					CategoryID:  c.Category.ID,
					Description: c.Description,
					ImagePath:   c.ImagePath,
					ProjectURL:  c.ProjectURL,
					Title:       c.Title,
					Avatar:      "", // TODO: avatar
				}
				ctx.JSON(http.StatusOK, res)
			}
		})
	}
	c = r.Group("/recommend")
	{
		c.GET("/", func(ctx *gin.Context) {
			// TODO: implement this
			if contents, e := content.BatchGet([]uint64{1, 5, 10}); e != nil {
				fmt.Printf("Failed to get content with id=%d, err=%s\n", 1, e.Error())
				ctx.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  "Internal Server Error",
				})
			} else {
				categories := []*RecommendCategory{}
				for _, c := range contents {
					res := Content{
						ID:          c.ID,
						Category:    c.Category.Name,
						CategoryID:  c.Category.ID,
						Description: c.Description,
						ImagePath:   c.ImagePath,
						ProjectURL:  c.ProjectURL,
						Title:       c.Title,
						Avatar:      "", // TODO: avatar
					}
					categories = append(categories, &RecommendCategory{
						Category:   res.Category,
						CategoryID: res.CategoryID,
						Contents:   []*Content{&res}})
				}
				ctx.JSON(http.StatusOK, categories)
			}
		})
		c.GET("/:category", func(ctx *gin.Context) {
			// TODO: implement this
			if c, e := content.Get(1); e != nil {
				fmt.Printf("Failed to get content with id=%d, err=%s\n", 1, e.Error())
				ctx.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  "Internal Server Error",
				})
			} else {
				res := Content{
					ID:          c.ID,
					Category:    c.Category.Name,
					CategoryID:  c.Category.ID,
					Description: c.Description,
					ImagePath:   c.ImagePath,
					ProjectURL:  c.ProjectURL,
					Title:       c.Title,
					Avatar:      "", // TODO: avatar
				}
				ctx.JSON(http.StatusOK, RecommendCategory{
					Category:   res.Category,
					CategoryID: res.CategoryID,
					Contents:   []*Content{&res}},
				)
			}
		})
	}
}
