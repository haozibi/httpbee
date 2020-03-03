# httpbee

![](logo.png)

[![CircleCI](https://circleci.com/gh/haozibi/httpbee.svg?style=svg&circle-token=459b435a854f30ff9ac58dcf43469a01ee369a23)](https://circleci.com/gh/haozibi/httpbee)

Fake HTTP Server

根据 JSON 配置文件做出响应

### 快速开始

```bash
$ httpbee -f=config.json
```

[config.json 示例](examples/config.json)

配置文件说明

|字段|类型|是否必需|默认值|说明|其他|
|:--|:--|:--|:--|:--|:--|
|router|string|是||匹配路由|*精准匹配，按照声明顺序匹配*|
|type|string|否|body|此路由响应方式|共支持三种（body、static、proxy）|
|resp.status|int|否|200|响应状态码|例如: 200、404、502|
|resp.headers|json|否||HTTP 响应头|例如: "content-type": "application/text" |
|resp.body|raw|否||type=body 时的响应输出|会按照 raw 形式输出|
|resp.file|string|否||type=static 时的响应内容的文件路径|例如: logo.png|
|resp.proxy|string|否||type=proxy 时代理的路径|例如: http://google.com|

```json
[
    {
        "router": "/hello",
        "resp": {
            "status": 200,
            "headers": {
                "version": "0.0.1",
                "content-type": "text/plain"
            },
            "body": "world"
        }
    },
    {
        "router": "/file",
        "type": "static",
        "resp": {
            "status": 200,
            "file": "logo.png"
        }
    },
    {
        "router": "/html",
        "type": "static",
        "resp": {
            "status": 200,
            "headers": {
                "content-type": "text/html"
            },
            "file": "1.html"
        }
    },
    {
        "router": "/json",
        "type": "body",
        "resp": {
            "status": 200,
            "headers": {
                "content-type": "application/json"
            },
            "body": {
                "err_code": 0,
                "err_msg": "success",
                "data": "ok"
            }
        }
    }
]
```

## todo

- [ ] 是否支持 Query 参数
- [ ] 是否支持 HTTP Method 方法过滤

## 已知 Bug

使用 vim 编辑 `config.json` 时会导致 [fsnotify](https://github.com/fsnotify/fsnotify) 监听错误（正常编辑会监听到 REMOVE 操作），导致 config.json 失效

- [https://github.com/fsnotify/fsnotify/issues/92](https://github.com/fsnotify/fsnotify/issues/92#issuecomment-262435215)

## 注意事项

- 路由精准匹配，并按照声明顺序尽心匹配

resp.type:

- static: 静态文件（文件具体形式根据 header 决定）
- body: 文件
