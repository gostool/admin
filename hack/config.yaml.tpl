
# CLI tool, only in development environment.
# https://goframe.org/pages/viewpage.action?pageId=3673173
gfcli:
  build:
    name: "main"
    arch: "amd64"
    system: "linux"
    mod: "none"
    cgo: 0
    pack: ""
    version: "0.0.1.0"
    extra: ""
  gen:
    dao:
      - link:     "mysql:admin:123456@tcp(127.0.0.1:3306)/db"
        tables:   "user,menu,role,operation_log,role_menu,api,casbin_rule"
        jsonCase: "CamelLower"