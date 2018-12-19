package main

import (
	"fmt"
	"github.com/ivanberry/amy-blog/pkg"
	"github.com/ivanberry/amy-blog/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:router,
		ReadTimeout:setting.ReadTimeOut,
		WriteTimeout:setting.WriteTimeOut,
		MaxHeaderBytes:1 << 20,
	}

	s.ListenAndServe()
}