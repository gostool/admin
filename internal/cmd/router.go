package cmd

import (
	"admin/internal/controller"
	"admin/internal/middleware"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func RegisterRouter(s *ghttp.Server, ctx context.Context, in cMainHttpInput) *ghttp.Server {
	s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
		//403 : func(r *ghttp.Request){r.Response.Writeln("403")},
		404: func(r *ghttp.Request) {
			g.Log().Debugf(ctx, `r: %+v`, r.GetUrl())
			r.Response.CORSDefault()
		},
		500: func(r *ghttp.Request) {
			r.Response.CORSDefault()
			r.Response.Writeln("500")
		},
	})
	// todo: 批量注册
	// https://goframe.org/pages/viewpage.action?pageId=1114517#id-%E8%B7%AF%E7%94%B1%E6%B3%A8%E5%86%8C%E5%88%86%E7%BB%84%E8%B7%AF%E7%94%B1-%E6%89%B9%E9%87%8F%E6%B3%A8%E5%86%8C
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			middleware.HandlerResponse,
			ghttp.MiddlewareCORS,
		)
		group.GET("/doc/", func(r *ghttp.Request) {
			r.Response.Write(swaggerUIPageContent)
		})
		// api
		group.Group("/api", func(apiGroup *ghttp.RouterGroup) {
			apiGroup.ALLMap(g.Map{
				"/hello/": controller.Hello, // user
				"/user/":  controller.User,  // user
				"/tools/": controller.Tools, // tools
			})
		})
		// auth
		group.Group("/api", func(apiGroup *ghttp.RouterGroup) {
			apiGroup.Middleware(
				middleware.JwtAuth,
				middleware.Record,
			)
			apiGroup.ALLMap(g.Map{
				"/role/":       controller.Role,      // role
				"/menu/":       controller.Menu,      // menu
				"/log/":        controller.Log,       // log
				"/status":      controller.Status,    // status
				"/role_menu/":  controller.RoleMenu,  // role_menu
				"/admin_user/": controller.AdminUser, // admin user
			})
		})
	})
	return s
}
