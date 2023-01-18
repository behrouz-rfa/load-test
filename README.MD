### Golang [Load test]() is a load testing tool written in Go It currently supports GET,

#### 👨‍💻 Full list what has been used:
* [fasthttp](github.com/valyala/fasthttp/) - fasthttp
* [color](github.com/fatih/color/) - fatih color
* [Console Std Out](github.com/apoorvam/gotermina) - gotermina


#### Recommendation for local development most comfortable usage:
    make local // run all containers
    make run // run the application

#### 🙌👨‍💻🚀 Docker-compose files:
    docker-compose.yml

### Docker development usage:
    make docker

### Local development usage:
    make local
    make run

Demo:

[//]: # (![]&#40;http://damien.pobel.fr/images/youtube-video-github.gif&#41;)

[//]: # ()
[//]: # (See also [Embed a Youtube video in Github markdown]&#40;http://damien.pobel.fr/post/youtube-video-github/&#41;)

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


