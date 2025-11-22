package main

import (
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/repostitory"
	dao "github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/repostitory/dao"
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/service"
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/web"
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {

	db := initDB()
	u := initUser(db)
	server := initWebServer()
	u.RegisterRoutesV1(server.Group("/users"))
	server.Run(":8080")
}
func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		//我只会在初始化过程中panic
		//panic相当于整个goroutine结束
		//一旦初始化错误，应用不在启动了
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initUser(db *gorm.DB) *web.UserHandler {

	ud := dao.NewUserDao(db)
	repo := repostitory.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	return web.NewUserHandler(svc)
}
func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"http://localhost:3000"},
		//AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"authorization,content-type"},
		//ExposeHeaders:    []string{"Content-Length"},
		//是否允许你带cookie之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return strings.Contains(origin, "localhost")
		},
		MaxAge: 12 * time.Hour,
	}))
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	server.Use(middleware.NewLoginMiddleWareBuilder().
		IgnorePath("/user/signup").
		IgnorePath("/users/login").
		Build())

	return server
}
