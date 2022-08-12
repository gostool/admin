package cmd

import (
	"admin/internal/controller"
	"admin/internal/middleware"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

const (
	swaggerUIPageContent = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <meta name="description" content="SwaggerUI"/>
  <title>SwaggerUI</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@latest/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@latest/swagger-ui-bundle.js" crossorigin></script>
<script>
	window.onload = () => {
		window.ui = SwaggerUIBundle({
			url:    '/api.json',
			dom_id: '#swagger-ui',
		});
	};
</script>
</body>
</html>
`
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
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
					apiGroup.ALLMap(g.Map{
						"/role/": controller.Role, // role
						"/menu/": controller.Menu, // menu
					})
				})
			})
			s.SetOpenApiPath("/api.json")
			s.Run()
			return nil
		},
	}
)
