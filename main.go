package main

import (
	"github.com/LeonClancy/baccarat-bet-analyze/controller"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	rc := &controller.RoadmapController{}
	web.Get("/roadmap", rc.GetRoadmap)
	web.Post("/roadmap", rc.NewRoadmap)
	web.Get("/roadmap/:id", rc.GetRoadmapById)
	web.Patch("/roadmap/:id", rc.Draw)
	web.Run()
}
