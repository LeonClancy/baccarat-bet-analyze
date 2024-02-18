package service

import (
	"github.com/LeonClancy/baccarat-bet-analyze/manager"
	"github.com/google/uuid"
	"sync"
)

type roadmapService struct {
	RoadmapManagers map[string]*manager.RoadmapManager
	Mutex           sync.Mutex
}

var RoadmapService *roadmapService

func init() {
	RoadmapService = &roadmapService{}
	RoadmapService.RoadmapManagers = make(map[string]*manager.RoadmapManager)
	RoadmapService.Mutex = sync.Mutex{}
}

func (service *roadmapService) NewRoadmap(name string) string {
	service.Mutex.Lock()
	defer service.Mutex.Unlock()
	newUUID := uuid.New().String()
	service.RoadmapManagers[newUUID] = manager.NewRoadmapManager(name)
	return newUUID
}
