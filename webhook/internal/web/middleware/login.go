package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddleWareBuilder struct {
	path []string
}

func (l *LoginMiddleWareBuilder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, path := range l.path {
			if c.Request.URL.Path == path {
				return
			}
		}
		session := sessions.Default(c)
		if session.Get("userId") == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		id := session.Get("userId")
		if id == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
func NewLoginMiddleWareBuilder() *LoginMiddleWareBuilder {
	return &LoginMiddleWareBuilder{}
}
func (l *LoginMiddleWareBuilder) IgnorePath(path string) *LoginMiddleWareBuilder {
	l.path = append(l.path, path)
	return l
}
