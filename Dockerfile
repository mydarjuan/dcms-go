FROM library/golang

## Install beego and the bee dev tool
RUN go get github.com/astaxie/beego && go get github.com/beego/bee

