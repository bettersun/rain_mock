# rain mock

一个 Mock + 转发请求 的简易工具。

作为 Mock 时，可设置返回的HTTP状态码和响应体的内容。
也可以设置响应头部的信息。

转发请求时，可保存请求信息（请求头部和请求体）和响应信息（响应头部和响应体）。  
还可以自动记录代理转发的相关信息，再改为 Mock 项目的配置信息。

依赖的个人 github 仓库:  
https://github.com/bettersun/rain

## 配置

### config.yml

程序运行所需的配置，修改该文件的内容后，需要重新启动该工具以生效。  
格式为 yaml 格式，位于 config 目录下，文件位置不可更改。

```yml
# 此文件内容修改后需重启工具
# 端口
port: 51721
# 日志文件
logFile: mock.log
# 日志等级
# -1: DEBUG
# 0: INFO  默认为INFO
# 1: WARN
# 2: ERROR
# 3: FATAL
logLevel: 0
mock:
  # 目标主机
  destHost: http://127.0.0.1:952711
  # Mock 项目文件
  mockItemFile: config/item.json
  # 使用 Mock 通用响应头 仅当URL对应的响应头不存在时使用
  useCommonHeader: true
path:
  # 保存请求信息的相对位置（相对于工具所在目录）
  request: _request
  # 保存响应信息的相对位置（相对于工具所在目录）
  response: _response
  # 备份 Mock 项目的目录（相对于工具所在目录）
  # 当保存改变后的 Mock 项目文件内容时，会先备份当前的 Mock 项目文件
  backup: _backup
  # 通用响应头部文件
  commonHeaderFile: config/common_header.json
  # 所有 API 的响应头部信息文件
  # 文件位置为 response 配置的位置
  responseHeaderFile: header.json
```

1. port

   工具的监听端口。
   如果配置的端口已占用，则启动工具会报错。

   > 手机的模拟器不能连接 127.0.0.1 或者 localhost，手机需要连接工具运行所在机器的IP。

2. logFile

   保存工具运行时日志的文件，相对目录为工具所在目录，所以不能设置为绝对路径。  
   日志文件为该工具运行时日志保存的文件名。

3. logLevel

   保存工具运行时日志的等级，高于或等于该等级的日志会被保存。

   数字对应等级为：  
   - -1: DEBUG
   - 0: INFO
   - 1: WARN
   - 2: ERROR
   - 3: FATAL

   默认使用 0 (信息)。

4. mock

   Mock 相关配置

   1. destHost
   
      不使用 Mock 时，如果有真实的服务，可转发请求。  
      该配置为被代理的目标服务的主机信息。  
      转发请求时，如果没有为每个 API 指定单独的目标主机，则使用该配置的目标主机。

   2. mockItemFile
   
      各个 API 的 Mock 项目配置的文件。  
      前端通过该工具连接的对象主机列表的文件名。

   3. useCommonHeader

      使用 Mock 通用响应头部，该配置为共通配置，仅当 Mock 项目的 path + method 对应的响应头部信息不存在时使用。  
      默认为 false 。

      转发请求时，会保存真实主机返回的响应头部信息。  
      位置为 {path -> response}/{path -> responseHeaderFile}。

      不转发请求，使用 Mock 服务时，对于各个 path + method 的响应，需要对应的响应头部信息。  
      默认提供了一个通用响应头部，位置是 config/{path -> commonHeaderFile}。

      对于 path + method 的请求，首先会查找是否存在真实服务的响应头部信息（转发请求时保存）。  
      如果存在时，则会使用转发请求时保存的响应头部信息。这种情况下有可能不是最新的期望的响应头部信息。  
      如果不存在，并且 useCommonHeader 为 true 时，则会使用 config.yml 中 path -> commonHeaderFile 的内容作为 Mock 服务的响应头部信息。    
      如果不存在，并且 useCommonHeader 为 false 时，则不作特殊处理。这种情况可能会与期望的响应头部信息有出入。  

5. path

   相关目录的位置或文件名。

   1. request

      保存请求信息的相对位置(相对于工具所在目录)。

   2. response

      保存转发请求的响应信息的相对位置(相对于工具所在目录)。
   
      目录下的 path -> response 的文件保存所有转发的 API （真实API）的响应头部信息。

   3. backup

      备份 Mock 项目的目录（相对于工具所在目录）。  
      当保存改变后的 Mock 项目文件内容时，会先备份当前的 Mock 项目文件。  
      现在的程序在记录代理转发信息时会保存到 Mock 项目文件中。

   4. commonHeaderFile

      通用响应头部信息文件。
      ```json
      {
        "Access-Control-Allow-Origin":[
          "*"
        ],
        "Content-Encoding":[
          "gzip"
        ],
        "Content-Type":[
          "application/json;charset=UTF-8"
        ]
      }
      ```

   5. responseHeaderFile

      所有 API 的响应头部信息文件，代理转发时会把 API 的响应头部信息保存到该文件中。  
      文件位置为 response 配置的位置。  
      Mock 服务也会使用该文件中的响应头部信息，需要修改的话，可修改该文件中 path + method 对应的头部信息。

### Mock 项目文件

json 格式。  
文件名是 config.yml 文件中 mock -> mockItemFile 设置的值，初始是 `item.json`。  
修改该文件内容时，工具会重新读取文件内容并重启服务。

转发请求时，会自动记录 path + method 对应的的 Mock 信息并保存到文件。  
前提是必须指定公共的目标主机，这样工具才能知道转发请求的目标主机。

1. url

   请求的目标 URL 。  
   和 method 两个项目决定一个 Mock 项目的信息。

2. method

   请求目标 URL 时的请求方法。  
   和 url 两个项目决定一个 Mock 项目的信息。

3. destHost
   
   不使用 Mock ，使用代理转发时，真实 API 所在的目标主机。
   
4. useMock

   设置为 `true` 时，使用 Mock 服务，返回配置的文件内容。  
   设置为 `false` 时，不使用 Mock 服务，对请求进行转发。

5. duration

   使用 Mock 服务时，响应需要等待的时间，用于模拟一些花时间的异步处理，如保存文件等。

6. statusCode

   使用 Mock 服务时，Mock 服务返回的 HTTP 状态码。

7. responseFile

   使用 Mock 服务时，返回配置文件的内容作为响应体的内容。

8. description
   
   Mock 项目的描述。

### 通用响应头部信息文件

common_header.json

可在 config.yml 中的 path -> commonHeaderFile 配置。

## 运行工具

根目录下运行 `go build` 命令后生成可执行文件，运行可执行文件即可。

可编译后单独运行。

```
go build
```

Mac 有可能需要 sudo ：
```
sudo go build
```

编译后文件名为 rainmock.exe(Mac下为 rainmock)。  
在终端或命令行启动工具。  

```
rainmock.exe
```

Mac：
```
./rainmock
```

启动后控制台会输出：
```
$ ./rainmock 
{"level":"info","msg":"=== 配置：[{Port:51721 LogFile:mock.log LogLevel:0 Mock:{DestHost:http://127.0.0.1:952711 MockItemFile:config/item.json UseConHeader:true} Path:{Request:_request Response:_response Backup:_backup CommonHeaderFile:config/common_header.json ResponseHeaderFile:header.json}}]","time":"2023-01-08T09:52:59+08:00"}
{"level":"info","msg":"Mock / 转发 服务运行中... 端口[51721]","time":"2023-01-08T09:52:59+08:00"}
```

item.json：
```
[
  {
    "path": "/bettersun/hello",
    "method": "POST",
    "useMock": false,
    "destHost": "http://127.0.0.1:9527",
    "duration": 500,
    "statusCode": 500,
    "responseFile": "_json/hello.json",
    "description": ""
  },
  {
    "path": "/bettersun/hello",
    "method": "GET",
    "useMock": true,
    "destHost": "http://127.0.0.1:9527",
    "duration": 200,
    "statusCode": 200,
    "responseFile": "_json/home.json",
    "description": ""
  }
]
```
item.json 中定义了两个 Mock 项目，

path + method 的组合分别是：
  1. /bettersun/hello POST
  2. /bettersun/hello POST

### 转发请求

其中 1 的 useMock 为 false ，对于客户端发送到工具的请求，不使用 Mock 服务，将请求转发到目标主机 127.0.0.0.1:9527 。  
如果该目标主机能处理对应的请求，则会将目标主机的响应再转发会工具的客户端。  
这时 duration / statusCode / responseFile 都不会起作用。

### 使用 Mock 服务

其中 2 的 useMock 为 true ，对于客户端发送到工具的请求，会使用 Mock 服务。  
duration 会是 Mock 服务响应的等待时间。  
statusCode 会是 Mock 服务响应的 HTTP 状态码。  
responseFile 的内容会是 Mock 服务响应的响应体内容。  

Mock 服务的响应头部，参考上面的 config.yml 的 mock -> useCommonHeader 和 path -> commonHeaderFile 。

### fakeapi/fakeapi.go

为测试用的假API（转发请求的目标服务）。

HOST: 127.0.0.1:9527

URL:
- /
- /bettersun
- /hello
- /bettersun/hello

每个 URL 都有 GET/PUT/POST/DELETE 的响应。

可编译后单独运行。

```
go build
```

Mac 有可能需要 sudo ：
```
sudo go build
```

编译后文件名为 fakeapi.exe(Mac下为 fakeapi)。

在终端或命令行运行：
```
fakeapi.exe
```

Mac：
```
./fakeapi
```

### 运行结果

上文中的 config.yml 和 item.json 的内容配置下，启动 fakeapi 的情况下， 运行结果为：

1. Postman 或其它客户端工具 GET 请求 http://127.0.0.1:51721/bettersun/hello

   该请求使用了 Mock 服务，所以运行结果为:
   
   HTTP 状态码：200

   响应体（item.json -> responseFile 指定的文件 _json/home.json 的内容）：
   ```json
   {
      "message": "welcome，世界"
   }
   ```
   
   控制台日志：
   ```
   {"level":"info","msg":"=== ################################################################################# ===","time":"2023-01-08T21:55:27+08:00"}
   {"level":"info","msg":"=== 响应处理开始 === 请求方法：[GET] 请求URL: [/bettersun/hello]","time":"2023-01-08T21:55:27+08:00"}
   {"level":"info","msg":"=== 使用Mock === Mock项目：[{/bettersun/hello GET http://127.0.0.1:9527 true 200 200 map[] _json/home.json }]","time":"2023-T21:55:27+08:00"}
   {"level":"info","msg":"=== 响应等待时间：[200]毫秒","time":"2023-01-08T21:55:27+08:00"}
   {"level":"warning","msg":"URL的响应头信息不存在，使用Mock 通用响应头","time":"2023-01-08T21:55:27+08:00"}
   {"level":"info","msg":"=== 使用已保存的响应头信息 ===","time":"2023-01-08T21:55:27+08:00"}
   {"level":"info","msg":"=== Mock响应正常结束 === ","time":"2023-01-08T21:55:27+08:00"}
   ```

2. Postman 或其它客户端工具 POST 请求 http://127.0.0.1:51721/bettersun/hello

   该请求未使用 Mock 服务，转发的请求，所以运行结果为:

   HTTP 状态码为 fakeapi (http://127.0.0.1:9527/bettersun/hello) 的 POST 的响应状态码。  
   fakeapi 正常运行的话，为 200 。

   响应头部和响应体是 fakeapi (http://127.0.0.1:9527/bettersun/hello) 的 POST 的响应头部和响应体：
   
   ```
   [Post] Hello, bettersun.
   ```

   控制台日志：

   ```
   {"level":"info","msg":"=== ################################################################################# ===","time":"2023-01-08T22:01:42+08:00"}
   {"level":"info","msg":"=== 响应处理开始 === 请求方法：[POST] 请求URL: [/bettersun/hello]","time":"2023-01-08T22:01:42+08:00"}
   {"level":"info","msg":"=== 代理转发 === 目标主机：[http://127.0.0.1:9527]","time":"2023-01-08T22:01:42+08:00"}
   {"level":"info","msg":"=== 完整URL: [http://127.0.0.1:9527/bettersun/hello]","time":"2023-01-08T22:01:42+08:00"}
   {"level":"info","msg":"=== 代理响应正常结束 === ","time":"2023-01-08T22:01:42+08:00"}
   {"level":"info","msg":"=== 响应保息已保存 === 文件: [_response/POST/bettersun_hello/body_0108220142.txt]","time":"2023-01-08T22:01:42+08:00"}
   {"level":"warning","msg":"响应信息转换JSON失败，保存为普通文件。","time":"2023-01-08T22:01:42+08:00"}
   ```
   
   其中：
   ```
   响应信息转换JSON失败，保存为普通文件。
   ```
   是由于程序会对转发请求的响应体内容是否是 json 的判断，不是错误。

3. item.json 修改后保存，工具会实时加载 item.json 保存后的内容，并重启工具的服务，再再次请求时会使用新的 Mock 项目进行响应。

   原来是 Mock 服务的可以改为转发请求，反之亦然。

   item.json 修改后保存时的控制台日志：

   ```
   {"level":"info","msg":"=== ################################################################################# ===","time":"2023-01-08T22:18:21+08:00"}
   {"level":"info","msg":"修改文件：[/Users/sunjiashu/Documents/Develop/github.com/bettersun/rainmock/config/item.json]\n","time":"2023-01-08T22:18:210"}
   {"level":"info","msg":"停止当前Mock服务","time":"2023-01-08T22:18:21+08:00"}
   {"level":"info","msg":"服务已关闭","time":"2023-01-08T22:18:21+08:00"}
   {"level":"info","msg":"重新加载文件信息","time":"2023-01-08T22:18:21+08:00"}
   {"level":"warning","msg":"http: Server closed","time":"2023-01-08T22:18:21+08:00"}
   {"level":"info","msg":"=== 配置：[{Port:51721 LogFile:mock.log LogLevel:0 Mock:{DestHost:http://127.0.0.1:952711 MockItemFile:config/item.json UseConHeader:true} Path:{Request:_request Response:_response Backup:_backup CommonHeaderFile:config/common_header.json ResponseHeaderFile:header.json}}]","time":"2023-01-08T22:18:21+08:00"}
   {"level":"info","msg":"重新启动Mock服务","time":"2023-01-08T22:18:21+08:00"}
   {"level":"info","msg":"Mock / 转发 服务运行中... 端口[51721]","time":"2023-01-08T22:18:21+08:00"}
   ```

## 后记

该工具现在非常简单，也有很多问题，总体来说，一般的 Mock + 转发请求 是够用的。  
有问题欢迎 issue 。

---