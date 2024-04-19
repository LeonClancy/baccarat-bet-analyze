package main

import (
	"github.com/LeonClancy/baccarat-bet-analyze/controller"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		//ExposeHeaders:    []string{"Content-Length"},
	}))
	rc := &controller.RoadmapController{}
	web.Get("/roadmap", rc.GetRoadmap)
	web.Post("/roadmap", rc.NewRoadmap)
	web.Get("/roadmap/:id", rc.GetRoadmapById)
	web.Patch("/roadmap/:id", rc.Draw)
	web.Patch("/roadmap/:id/restore", rc.Restore)

	web.Get("/roadmap/patterns", rc.GetPatterns)
	web.Patch("/roadmap/:id/patterns", rc.SetPatterns)

	web.Run()
}
