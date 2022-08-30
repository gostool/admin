---
swagger: '2.0'
info:
  title: App API
  description: Test description
  termsOfService: https://www.google.com/policies/terms/
  contact:
    email: admin@test.com
  license:
    name: BSD License
  version: v1
host: localhost:8000
schemes:
- https
- http
basePath: "/"
consumes:
- application/json
produces:
- application/json
securityDefinitions:
  basic:
    type: basic
security:
- basic: []
paths:
 "/api/login/":
    post:
      operationId: api_login_create
      summary: 用户名密码登录
      description: "---"
      parameters:
      - name: data
        in: body
        required: true
        schema:
          "$ref": "#/definitions/Login"
      responses:
        '201':
          description: ''
          schema:
            "$ref": "#/definitions/Login"
      tags:
      - api
    parameters: []
  "/api/logout/":
    get:
      operationId: api_logout_list
      summary: 用户退出
      description: "---"
      parameters:
      - name: page
        in: query
        description: A page number within the paginated result set.
        required: false
        type: integer
      - name: page_size
        in: query
        description: Number of results to return per page.
        required: false
        type: integer
      responses:
        '200':
          description: ''
          schema:
            required:
            - count
            - results
            type: object
            properties:
              count:
                type: integer
              next:
                type: string
                format: uri
                x-nullable: true
              previous:
                type: string
                format: uri
                x-nullable: true
              results:
                type: array
                items:
                  "$ref": "#/definitions/Logout"
      tags:
      - api
    post:
      operationId: api_logout_create
      summary: 用户退出
      description: "---"
      parameters:
      - name: data
        in: body
        required: true
        schema:
          "$ref": "#/definitions/Logout"
      responses:
        '201':
          description: ''
          schema:
            "$ref": "#/definitions/Logout"
      tags:
      - api
    parameters: []
definitions:
  Auth:
    required:
    - code
    type: object
    properties:
      code:
        title: 用户登录凭证
        type: string
        minLength: 1
      encrypted_data:
        title: 加密数据
        type: string
        minLength: 1
      iv:
        title: 初始向量
        type: string
        minLength: 1
      session_key:
        title: session_key
        type: string
        minLength: 1
      openid:
        title: openid
        type: string
        readOnly: true
        minLength: 1
  CodeToOpenId:
    required:
    - code
    type: object
    properties:
      code:
        title: code
        type: string
        minLength: 1
  Decrypt:
    required:
    - code
    type: object
    properties:
      code:
        title: 用户登录凭证
        type: string
        minLength: 1
      encrypted_data:
        title: 加密数据
        type: string
        minLength: 1
      iv:
        title: 初始向量
        type: string
        minLength: 1
      session_key:
        title: session_key
        type: string
        minLength: 1
  Login:
    required:
    - username
    - password
    type: object
    properties:
      username:
        title: 用户名
        type: string
        maxLength: 30
        minLength: 1
      password:
        title: 密码
        type: string
        minLength: 1
  Logout:
    type: object
    properties: {}
