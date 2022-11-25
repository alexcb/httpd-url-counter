package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	counts := map[string]int{}

	e := echo.New()
	e.GET("*", func(c echo.Context) error {
		path := c.Request().URL.Path
		numGets := counts[path]
		counts[path] = numGets + 1
		remoteAddr := c.Request().RemoteAddr
		fmt.Printf("%v %s %s %d\n", time.Now().UTC().Format(time.RFC3339), remoteAddr, path, numGets)
		return c.String(http.StatusOK, fmt.Sprintf("%d\n", numGets))
	})

	err := http.ListenAndServe("0.0.0.0:1234", e)
	if err != nil {
		panic(err)
	}
}
