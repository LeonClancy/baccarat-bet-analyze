package controller

import (
	"github.com/LeonClancy/baccarat-bet-analyze/roadmap"
	"strconv"

	"github.com/LeonClancy/baccarat-bet-analyze/dealer"
	"github.com/LeonClancy/baccarat-bet-analyze/manager"
	"github.com/LeonClancy/baccarat-bet-analyze/service"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
)

type RoadmapController struct {
	web.Controller
}

type createRoadmapRequest struct {
	Name string `json:"name"`
}

func (r *RoadmapController) NewRoadmap(ctx *context.Context) {
	req := &createRoadmapRequest{}
	ctx.BindJSON(req)
	roadmapUuid := service.RoadmapService.NewRoadmap(req.Name)
	response := gmap.New()
	response.Set("uuid", roadmapUuid)
	err := ctx.Output.JSON(response, false, true)
	if err != nil {
		response := gmap.New()
		response.Set("error", err.Error())
		ctx.Output.JSON(response, false, true)
		return
	}
}

func (r *RoadmapController) GetRoadmap(ctx *context.Context) {
	managers := service.RoadmapService.RoadmapManagers
	managerData := garray.NewArray()
	for uuid := range managers {
		data := gmap.New()
		data.Set("name", managers[uuid].Name)
		data.Set("uuid", uuid)
		managerData.Append(data)
	}

	response := gmap.New()
	response.Set("roadmaps", managerData)
	err := ctx.Output.JSON(response, false, true)
	if err != nil {
		response := gmap.New()
		response.Set("error", err.Error())
		ctx.Output.JSON(response, false, true)
		return
	}
}

type drawRoadmapRequest struct {
	Result int `json:"result"`
}

func (r *RoadmapController) Draw(ctx *context.Context) {
	roadmapUuid := ctx.Input.Param(":id")
	manager, ok := service.RoadmapService.RoadmapManagers[roadmapUuid]
	if !ok {
		ctx.Output.SetStatus(404)
		response := gmap.New()
		response.Set("error", "roadmap not found")
		ctx.Output.JSON(response, false, true)
	}
	result := ctx.Request.FormValue("result")
	var results []dealer.Result
	switch result {
	case "1":
		results = append(results, dealer.Result_Banker)
	case "2":
		results = append(results, dealer.Result_Player)
	case "3":
		results = append(results, dealer.Result_Tie)
	}
	manager.Draw(results)
	bankerAskRoadResult := manager.AskRoad(roadmap.Symbol_Banker)
	playerAskRoadResult := manager.AskRoad(roadmap.Symbol_Player)
	manager.AnalyzeManager.Analyze(manager.Roadmaps, bankerAskRoadResult, playerAskRoadResult)

	response := gmap.New()
	response.Set("roadmaps", manager.Roadmaps)
	response.Set("result_counter", manager.ResultCounter)
	response.Set("predictions", manager.AnalyzeManager.Predictions)
	err := ctx.Output.JSON(response, false, true)
	if err != nil {
		response := gmap.New()
		response.Set("error", err.Error())
		ctx.Output.JSON(response, false, true)
		return
	}
}

func (r *RoadmapController) Restore(ctx *context.Context) {
	roadmapUuid := ctx.Input.Param(":id")
	roadmapManager, ok := service.RoadmapService.RoadmapManagers[roadmapUuid]
	if !ok {
		ctx.Output.SetStatus(404)
		response := gmap.New()
		response.Set("error", "roadmap not found")
		ctx.Output.JSON(response, false, true)
	}
	roadmapManager.Restore()
	bankerAskRoadResult := roadmapManager.AskRoad(roadmap.Symbol_Banker)
	playerAskRoadResult := roadmapManager.AskRoad(roadmap.Symbol_Player)
	roadmapManager.AnalyzeManager.Analyze(roadmapManager.Roadmaps, bankerAskRoadResult, playerAskRoadResult)
	response := gmap.New()
	response.Set("roadmaps", roadmapManager.Roadmaps)
	response.Set("result_counter", roadmapManager.ResultCounter)
	response.Set("predictions", roadmapManager.AnalyzeManager.Predictions)
	err := ctx.Output.JSON(response, false, true)
	if err != nil {
		response := gmap.New()
		response.Set("error", err.Error())
		ctx.Output.JSON(response, false, true)
		return
	}
}

func (r *RoadmapController) GetRoadmapById(ctx *context.Context) {
	roadmapUuid := ctx.Input.Param(":id")
	manager, ok := service.RoadmapService.RoadmapManagers[roadmapUuid]
	if !ok {
		ctx.Output.SetStatus(404)
		response := gmap.New()
		response.Set("error", "roadmap not found")
		ctx.Output.JSON(response, false, true)
	}
	response := gmap.New()
	response.Set("name", manager.Name)
	response.Set("roadmaps", manager.Roadmaps)
	ctx.Output.JSON(response, false, true)
}

func (r *RoadmapController) SetPatterns(ctx *context.Context) {
	roadmapUuid := ctx.Input.Param(":id")
	manager, ok := service.RoadmapService.RoadmapManagers[roadmapUuid]
	response := gmap.New()

	if !ok {
		ctx.Output.SetStatus(404)
		response.Set("error", "roadmap not found")
		ctx.Output.JSON(response, false, true)
	}

	pattern1Str := ctx.Request.FormValue("pattern1")
	if pattern1Str != "" {
		// pattern1Str to int
		pattern1, err := strconv.Atoi(pattern1Str)
		if err != nil {
			response.Set("error", err.Error())
			ctx.Output.JSON(response, false, true)
			return
		}
		manager.AnalyzeManager.Pattern1 = pattern1
	}

	pattern2Str := ctx.Request.FormValue("pattern2")
	if pattern2Str != "" {
		// pattern2Str to int
		pattern2, err := strconv.Atoi(pattern2Str)
		if err != nil {
			response.Set("error", err.Error())
			ctx.Output.JSON(response, false, true)
			return
		}
		manager.AnalyzeManager.Pattern2 = pattern2
	}

	//manager.AnalyzeManager.Analyze(manager.Roadmaps)

	response.Set("pattern1", pattern1Str)
	response.Set("pattern2", pattern2Str)
	response.Set("roadmaps", manager.Roadmaps)
	response.Set("predictions", manager.AnalyzeManager.Predictions)
	err := ctx.Output.JSON(response, false, true)
	if err != nil {
		response := gmap.New()
		response.Set("error", err.Error())
		ctx.Output.JSON(response, false, true)
		return
	}
}

func (r *RoadmapController) GetPatterns(ctx *context.Context) {
	response := gmap.New()
	response.Set("patterns", manager.Patterns)
	err := ctx.Output.JSON(response, false, true)
	if err != nil {
		response := gmap.New()
		response.Set("error", err.Error())
		ctx.Output.JSON(response, false, true)
		return
	}
}
