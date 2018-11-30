package router

import (
	"github.com/LyricTian/gin-admin/src/service/mysql"
	"github.com/LyricTian/gin-admin/src/web/ctl"
	"github.com/gin-gonic/gin"
)

// APIV1Handler /api/v1路由
func APIV1Handler(r *gin.Engine, db *mysql.DB, c *ctl.Common) {
	relativePath := "/api/v1/"
	v1 := r.Group(relativePath,
		SessionMiddleware(db, relativePath),
		VerifySessionMiddleware(
			[]string{relativePath},
			[]string{
				relativePath + "login",
				relativePath + "logout",
			},
		))

	// 注册路由
	APILoginRouter(v1, c.LoginAPI)
	APIRoleRouter(v1, c.RoleAPI)
	APIDemoRouter(v1, c.DemoAPI)
	APIMenuRouter(v1, c.MenuAPI)
	APIUserRouter(v1, c.UserAPI)
}