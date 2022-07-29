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
					r.Response.CORSDefault()
				},
				500: func(r *ghttp.Request) {
					r.Response.CORSDefault()
					r.Response.Writeln("500")
				},
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					middleware.HandlerResponse,
					//ghttp.MiddlewareCORS,
				)
				group.GET("/doc/", func(r *ghttp.Request) {
					r.Response.Write(swaggerUIPageContent)
				})
				// api
				group.Group("/api", func(apiGroup *ghttp.RouterGroup) {
					apiGroup.Group("/v1", func(apiV1Group *ghttp.RouterGroup) {
						apiV1Group.Bind(
							controller.Hello,
						)
					})
					apiGroup.Group("/v2", func(apiV2Group *ghttp.RouterGroup) {
						apiV2Group.Bind(
							controller.Hello,
						)
					})
				})
			})
			s.SetOpenApiPath("/api.json")
			s.Run()
			return nil
		},
	}
)
