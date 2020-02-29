# httpbee

[![CircleCI](https://circleci.com/gh/haozibi/httpbee.svg?style=svg&circle-token=459b435a854f30ff9ac58dcf43469a01ee369a23)](https://circleci.com/gh/haozibi/httpbee)

Fake HTTP Server

根据 file 做出相应的响应

路由是有顺序的

```json
[
    {
        "router": "/hello",
        "resp": {
            "status": 200,
            "headers": {
                "content-type": "application/text"
            },
            "body": "world"
        }
    },
    {
        "router": "/json",
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