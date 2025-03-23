package main

import (
  "lore-fetcher/internal/adapters/handler"
  "lore-fetcher/internal/adapters/repository"
  "lore-fetcher/internal/core/services"
  "github.com/gin-gonic/gin"
  "lore-fetcher/internal/fetcher"
)

var (
  httpHandler *handler.HTTPHandler
  service *services.LFService
)

func main(){
  store := repository.NewLFPostgresRepository()
  service = services.NewLFService(store)
  fetcher := fetcher.NewFetcher(*service)
  go fetcher.FetchDaemon()

  InitRoutes()
}

func InitRoutes (){
  router := gin.Default()
  handler := handler.NewHTTPHandler(*service)
  router.GET("/patches", handler.ReadPatches)
  router.POST("/patches", handler.SavePatch)
  router.Run(":8080")
}
