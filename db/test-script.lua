### 请求方式

wrk.method = "GET"



### 设置 请求类型

wrk.headers["Content-Type"] = "application/json"

wrk.headers["Token"] = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ5eWsiLCJleHAiOjE2OTQzMDkyNjksImp0aSI6IkJnNDA3NjgwMSIsImlhdCI6MTY5MzQ0NTI2OSwiaXNzIjoiYWRtaW4wMDAxIiwibmJmIjoxNjkzNDQ1MjY5LCJzdWIiOiJsb2dpbiJ9.CN3ArttY93s9kim3rk6QXvaoIXacupAWTFwT_xuF4ro"

### POST 请求参数

wrk.body = '{"val": "13999999999","username": "13999999999"}'