server:
  address:     ":8000"
  openapiPath: "/api.json"
  swaggerPath: "/swagger"
  serverRoot: "./resource"

logger:
  level : "all"
  stdout: true

database:
  default:
    link:  "mysql:root:mysecretpassword@tcp(127.0.0.1:3306)/godb"
  logger:
    Path: /var/log/sql/
    Level: all
    Stdout: true


