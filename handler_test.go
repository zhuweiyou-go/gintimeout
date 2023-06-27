package gintimeout

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestHandler(t *testing.T) {
	r := gin.Default()
	r.Use(Handler(time.Minute * 2))
}
