package main

import (
	"flag"

	"github.com/FrontMage/HelloGithubNavBackend/dao"
	"github.com/FrontMage/HelloGithubNavBackend/routers"
	"github.com/gin-gonic/gin"
)

var httpBindAddr = flag.String("http_bind_addr", "0.0.0.0:8080", "http listening address")

func init() {
	flag.Parse()
}

func main() {
	r := gin.Default()
	routers.MountRouters(r)
	defer func() {
		if dao.DB != nil {
			dao.DB.Close()
		}
	}()
	r.Run(*httpBindAddr)
}
