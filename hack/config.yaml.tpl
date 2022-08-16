
# CLI tool, only in development environment.
# https://goframe.org/pages/viewpage.action?pageId=3673173
gfcli:
  gen:
    dao:
      - link:     "mysql:admin:123456@tcp(127.0.0.1:3306)/db"
        tables:   "user,menu,role,operation_log,role_menu"
        jsonCase: "CamelLower
