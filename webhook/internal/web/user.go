package web

import (
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/domain"
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"time"
)

type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"

		passwordRegexPattern = `^[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern)
	passwordExp := regexp.MustCompile(passwordRegexPattern)
	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	var req SignUpReq
	//Bind 方法会根据Content-Type 来解析你的数据到req里面
	//解析错误，就会直接写回一个400的错误
	if err := ctx.Bind(&req); err != nil {
		return
	}
	ok := u.emailExp.MatchString(req.Email)
	if !ok {
		ctx.String(http.StatusOK, "你的邮箱格式不对")
		return
	}
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
	}
	ok = u.passwordExp.MatchString(req.Password)
	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}
	err := u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
	}

}
func (u *UserHandler) Login(c *gin.Context) {}
func (u *UserHandler) Edit(c *gin.Context) {
	type EditReq struct {
		Nickname    string `json:"nickname"`
		Birthday    string `json:"birthday"`
		Description string `json:"description"`
	}
	var req EditReq
	if err := c.ShouldBind(&req); err != nil {
		return
	}
	birthday, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		c.String(http.StatusOK, "生日日期有误")
	}
	u.svc.Edit(c, domain.EditUserInfo{
		Nickname:    req.Nickname,
		Birthday:    birthday,
		Description: req.Description,
	})
}
func (u *UserHandler) Profile(c *gin.Context) {}

/*
	func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
		server.POST("/users/signup", u.SignUp)
		server.POST("/users/login", u.Login)
		server.POST("/users/edit", u.Edit)
		server.GET("/users/profile", u.Profile)
	}
*/
func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit/:id", u.Edit)
	ug.GET("/profile", u.Profile)
}
func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}
