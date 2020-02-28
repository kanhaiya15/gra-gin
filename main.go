package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	LoginN
}

type LoginN struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBind(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.Run(":8080")
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"github.com/gin-gonic/gin"
// 	api "github.com/kanhaiya15/gra-gin/apis/v1"
// 	"github.com/kanhaiya15/gra-gin/cfg"
// 	"github.com/kanhaiya15/gra-gin/cfg/dbs/kmysql"
// 	"github.com/kanhaiya15/gra-gin/pkg/klog"
// )

// func main() {
// 	klog.NewLogger()
// 	cfg.NewConfig("./cfg", "pre-live")
// 	kmysql.NewConfig()
// 	cleanupHook()
// 	router := gin.Default()
// 	v1 := router.Group("/gin/api/v1/todos")
// 	{
// 		v1.POST("/", api.CreateTodo)
// 		v1.GET("/", api.FetchAllTodo)
// 		v1.GET("/:id", api.FetchSingleTodo)
// 		v1.PUT("/:id", api.UpdateTodo)
// 		v1.DELETE("/:id", api.DeleteTodo)
// 	}
// 	router.Run(":8000")
// }

// func cleanupHook() {
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)
// 	signal.Notify(c, syscall.SIGTERM)
// 	signal.Notify(c, syscall.SIGKILL)
// 	go func() {
// 		<-c
// 		klog.SLogger.Sync()
// 		kmysql.DB.Close()
// 		os.Exit(0)
// 	}()
// }

// func loggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Do stuff here
// 		log.Println(r.RequestURI)
// 		// Call the next handler, which can be another middleware in the chain, or the final handler.
// 		next.ServeHTTP(w, r)
// 	})
// }
