# API Recorder
This is simple server that keeps track of every API call made to it and then exposes them.

## Docker image
[![](https://images.microbadger.com/badges/version/intraway/api_recorder.svg)](https://microbadger.com/images/intraway/api_recorder "Get your own version badge on microbadger.com")
[![](https://images.microbadger.com/badges/image/intraway/api_recorder.svg)](https://microbadger.com/images/intraway/api_recorder "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/commit/intraway/api_recorder.svg)](https://microbadger.com/images/intraway/api_recorder "Get your own commit badge on microbadger.com")
[![](https://images.microbadger.com/badges/license/intraway/api_recorder.svg)](https://microbadger.com/images/intraway/api_recorder "Get your own license badge on microbadger.com")

## Configuration
There are a few things that can be changed with environment variables:
* `PORT`: default is 8080.
* `HOST`: default is 0.0.0.0
* `SHOW_URL`: default is 'showmewhatyougot'
* `RESET_URL`: default is 'resetwhatyougot'

## Manage
### Get calls
Get the API calls with the `/showmewhatyougot` endpoint

![showme](https://cloud.githubusercontent.com/assets/1962934/26114632/2a98048c-3a34-11e7-88ba-94f5f5a5b9e4.png)

Available filters:
* `url=<url>`: only get call to `<url>`
* `n=<number>`: retrieve the latest `n` calls (only works with `url`)

### Reset calls
Reset the API calls with the `/resetwhatyougot` endpoint

Available filters:
* `url=<url>`: only reset `<url>` calls

## Example
```
$ sudo docker run -t --rm -e PORT=9000 intraway/api_recorder
```

Once the server is running, you can try making some calls:
```
$ curl 'http://127.0.0.1:9000/aaa'
$ curl 'http://127.0.0.1:9000/aaa?q1=10'
$ curl -X POST 'http://127.0.0.1:9000/bbb' -d 'hello'
$ curl -X POST 'http://127.0.0.1:9000/bbb' -H "Content-Type: application/json" -d '{"one": 1}'
$ curl -X PUT 'http://127.0.0.1:9000/ccc' -H "Content-Type: application/json" -d '{"two": 2}'
```

### Get all calls
![showme](https://cloud.githubusercontent.com/assets/1962934/26114632/2a98048c-3a34-11e7-88ba-94f5f5a5b9e4.png)
```
$ curl 'http://127.0.0.1:9000/showmewhatyougot'
{
  "/aaa": [
    {
      "URL": "/aaa",
      "Method": "GET",
      "URI": "/aaa",
      "Header": {
        "Accept": [
          "*/*"
        ],
        "User-Agent": [
          "curl/7.29.0"
        ]
      },
      "Proto": "HTTP/1.1",
      "Host": "127.0.0.1:9000",
      "Body": "",
      "RemoteAddress": "10.66.35.10:46004",
      "ContentType": "",
      "Timestamp": "2017-05-16T15:14:56.010454376Z"
    },
    {
      "URL": "/aaa",
      "Method": "GET",
      "URI": "/aaa?q1=10",
      "Header": {
        "Accept": [
          "*/*"
        ],
        "User-Agent": [
          "curl/7.29.0"
        ]
      },
      "Proto": "HTTP/1.1",
      "Host": "127.0.0.1:9000",
      "Body": "",
      "RemoteAddress": "10.66.35.10:46008",
      "ContentType": "",
      "Timestamp": "2017-05-16T15:14:56.019024735Z"
    }
  ],
  "/bbb": [
    {
      "URL": "/bbb",
      "Method": "POST",
      "URI": "/bbb",
      "Header": {
        "Accept": [
          "*/*"
        ],
        "Content-Length": [
          "5"
        ],
        "Content-Type": [
          "application/x-www-form-urlencoded"
        ],
        "User-Agent": [
          "curl/7.29.0"
        ]
      },
      "Proto": "HTTP/1.1",
      "Host": "127.0.0.1:9000",
      "Body": "hello",
      "RemoteAddress": "10.66.35.10:46012",
      "ContentType": "application/x-www-form-urlencoded",
      "Timestamp": "2017-05-16T15:14:56.029243233Z"
    },
    {
      "URL": "/bbb",
      "Method": "POST",
      "URI": "/bbb",
      "Header": {
        "Accept": [
          "*/*"
        ],
        "Content-Length": [
          "10"
        ],
        "Content-Type": [
          "application/json"
        ],
        "User-Agent": [
          "curl/7.29.0"
        ]
      },
      "Proto": "HTTP/1.1",
      "Host": "127.0.0.1:9000",
      "Body": "{\"one\": 1}",
      "RemoteAddress": "10.66.35.10:46016",
      "ContentType": "application/json",
      "Timestamp": "2017-05-16T15:14:56.041657195Z"
    }
  ],
  "/ccc": [
    {
      "URL": "/ccc",
      "Method": "PUT",
      "URI": "/ccc",
      "Header": {
        "Accept": [
          "*/*"
        ],
        "Content-Length": [
          "10"
        ],
        "Content-Type": [
          "application/json"
        ],
        "User-Agent": [
          "curl/7.29.0"
        ]
      },
      "Proto": "HTTP/1.1",
      "Host": "127.0.0.1:9000",
      "Body": "{\"two\": 2}",
      "RemoteAddress": "10.66.35.10:46020",
      "ContentType": "application/json",
      "Timestamp": "2017-05-16T15:14:56.598073152Z"
    }
  ]
}
```

### Get calls with filters
```
$ curl 'http://127.0.0.1:9000/showmewhatyougot?url=/aaa&n=1'
{
  "/aaa": [
    {
      "URL": "/aaa",
      "Method": "GET",
      "URI": "/aaa?q1=10",
      "Header": {
        "Accept": [
          "*/*"
        ],
        "User-Agent": [
          "curl/7.29.0"
        ]
      },
      "Proto": "HTTP/1.1",
      "Host": "127.0.0.1:9000",
      "Body": "",
      "RemoteAddress": "10.66.35.10:46008",
      "ContentType": "",
      "Timestamp": "2017-05-16T15:14:56.019024735Z"
    }
  ]
}
```

### Reset calls with filters
```
$ curl 'http://127.0.0.1:9000/resetwhatyougot?url=/aaa&url=/bbb'
$ curl 'http://127.0.0.1:9000/showmewhatyougot'
{
  "/ccc": [
    {
      "URL": "/ccc",
      "Method": "PUT",
      "URI": "/ccc",
      "Header": {
        "Accept": [
          "*/*"
        ],
        "Content-Length": [
          "10"
        ],
        "Content-Type": [
          "application/json"
        ],
        "User-Agent": [
          "curl/7.29.0"
        ]
      },
      "Proto": "HTTP/1.1",
      "Host": "127.0.0.1:9000",
      "Body": "{\"two\": 2}",
      "RemoteAddress": "10.66.35.10:46020",
      "ContentType": "application/json",
      "Timestamp": "2017-05-16T15:14:56.598073152Z"
    }
  ]
}
```

### Reset all calls
```
$ curl 'http://127.0.0.1:9000/resetwhatyougot'
$ curl 'http://127.0.0.1:9000/showmewhatyougot'
{}
```


## docker-compose
```
version: 2

service:
    api_recorder:
        image: intraway/api_recorder:latest
        environment:
            PORT=8080
            HOST=0.0.0.0
            SHOW_URL=showmewhatyougot
            RESET_URL=resetwhatyougot

```

