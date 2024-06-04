package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	r.POST(configInstance.Prefix+"/login", loginPost)
	r.NoRoute(proxyRequest)

	r.Run(configInstance.Host + ":" + strconv.Itoa(configInstance.Port))
}
