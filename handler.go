package gintimeout

import (
	timeout "github.com/vearne/gin-timeout"
	"time"

	"github.com/gin-gonic/gin"
)

func Handler(duration time.Duration) gin.HandlerFunc {
	return timeout.Timeout(timeout.WithTimeout(duration))
}
