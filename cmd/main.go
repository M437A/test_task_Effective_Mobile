package main

import (
	"fmt"
	"name_service/pkg/common"
	"name_service/pkg/routers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func main() {
	db := common.Init()
	defer db.Close()

	router := chi.NewRouter()
	routers.CreateRouters(router)

	err := http.ListenAndServe(viper.GetString("port"), router)
	if err != nil {
		panic(fmt.Sprintf("Server error: %v", err))
	}
}
