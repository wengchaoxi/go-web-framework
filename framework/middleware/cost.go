package middleware

import (
	"log"
	"time"

	"goweb/framework"
)

// 计算 api 耗时
func Cost() framework.Handler {
	return func(c *framework.Context) error {
		start := time.Now()
		c.Next()
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost: %v", c.GetRequest().RequestURI, cost.Seconds())
		return nil
	}
}