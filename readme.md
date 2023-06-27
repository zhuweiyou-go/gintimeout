# gintimeout

```bash
go get -u github.com/zhuweiyou-go/gintimeout
```

```go
package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhuweiyou-go/gintimeout"
)

func main() {
	r := gin.Default()
	r.Use(gintimeout.Handler(time.Minute * 2))
}
```
