# Application using Go Echo framework

`sudo docker build -t rin-echo:system-development .`

`go clearn -modcache`

`go mod tidy`

sudo docker run -it --rm -p 8090:8090 -v $PWD/src:/go/src/github.com/rinnguyen1614/rin-echo/internal/system// rin-echo:system-development
