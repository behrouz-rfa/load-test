### Golang [Load test]() is a load testing tool written in Go It currently supports GET,

#### 👨‍💻 Full list what has been used:
* [fasthttp](github.com/valyala/fasthttp/) - fasthttp
* [color](github.com/fatih/color/) - fatih color
* [Console Std Out](github.com/apoorvam/gotermina) - gotermina


#### Recommendation for local development most comfortable usage:
    make run   // run the application
    make race // run the application with  race flag
    make local // run all containers

#### 🙌👨‍💻🚀 Docker-compose files:
    docker-compose.yml

### Docker development usage:
    make docker

### Local development usage:
    make run
    make local

Demo:

[![Final video of fixing issues in your code in VS Code](https://i9.ytimg.com/vi_webp/6ewjpOUrGyU/mq2.webp?sqp=CPi_np4G-oaymwEmCMACELQB8quKqQMa8AEB-AH-CYAC0AWKAgwIABABGEEgUChlMA8=&rs=AOn4CLCYgNCMJ-aVRLP9f_bR4ti1vhPZAQ)](https://www.youtube.com/watch?v=6ewjpOUrGyU&list=PLc0PxFU2AtMRqWRMJUZChA8lqlR9UiE3h)

See also [The tutoiral on youtube](https://www.youtube.com/watch?v=6ewjpOUrGyU)

## Using 

supports the following options:

```
  -c int
        Number of concurrent requests (default 5)
  -d int
        Request duration (default 5)
  -h string
        header values separated with ';'
  -help
        know more command
  -m string
        Request method (default "GET")
  -to int
        Request time out in milli seconds (default 2000)
  -u string
        URL to run against
```


Below is A basic example which will use 10 workers and time out with 4000 millisec and duration is 5 Second: 

```sh
$ go run ./cmd/app/main.go -to 4000 -d 5 -c 5 -u https://jsonplaceholder.typicode.com/todos/1
```


