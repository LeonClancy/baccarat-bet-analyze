package main

import (
	"github.com/LeonClancy/baccarat-bet-analyze/controller"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	rc := &controller.RoadmapController{}
	web.Get("/roadmap", rc.GetRoadmap)
	web.Post("/roadmap", rc.NewRoadmap)
	web.Get("/roadmap/:id", rc.GetRoadmapById)
	web.Patch("/roadmap/:id", rc.Draw)
	web.Run()
}
