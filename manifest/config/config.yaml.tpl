server:
  address:     ":8000"
  openapiPath: "/api.json"
  swaggerPath: "/swagger"
  serverRoot: "./resource"
  LogPath: /tmp/log/admin/server
logger:
  Path: /tmp/log/admin
  Level: all
  Stdout: true
  debug:
    path: /tmp/log/admin/debug
    level: dev
    stdout: true

database:
  default:
    link:  "mysql:root:mysecretpassword@tcp(127.0.0.1:3306)/godb"
    debug:   true
  logger:
    Path:  /tmp/log/admin/sq
    Level: all
    Stdout: true


