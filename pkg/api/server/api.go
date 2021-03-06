package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"

	_ "github.com/KEVISONG/hello-go-web/statik"

	"github.com/KEVISONG/hello-go-web/pkg/api/controller"
	"github.com/KEVISONG/hello-go-web/pkg/config/server"
	"github.com/gin-gonic/gin"
)

func serveStatik(e *gin.Engine) {
	statikFS, err := fs.New()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	e.StaticFS("/index", statikFS)
}

func route(e *gin.Engine) {

	serveStatik(e)

	e.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := e.Group("/api").Group("/v1")
	v1.GET("/user/:id", controller.UserGet)
	v1.POST("/user", controller.UserPost)

}

// Run runs server
func Run(c server.Config) {

	engine := gin.Default()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	route(engine)

	engine.Run(fmt.Sprintf(":%d", c.Port))

}
