package main

import (
	"fmt"
	"net/http"
	"project/pkg/setting"
	"project/routers"
)

func main() {
	route := routers.InitRoute()
	s := &http.Server{
		Addr: 			fmt.Sprintf(":%d", setting.HTTPort),
		Handler: 		route,
		ReadTimeout: 	setting.ReadTimeout,
		WriteTimeout: 	setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
