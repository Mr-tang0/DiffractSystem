package main

import (
	services "Diffract/services"
	components "Diffract/services/components"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx     context.Context
	updater *services.UpdateService

	// 设备实例
	Stage    *components.StageService
	Detector *components.DetectorService
	HVPS     *components.HVPSService
}

func NewApp(
	stage *components.StageService,
	detector *components.DetectorService,
	hvps *components.HVPSService) *App {
	return &App{
		Stage:    stage,
		Detector: detector,
		HVPS:     hvps,
	}
}

func (a *App) startup(ctx context.Context) {
	a.updater = &services.UpdateService{}
	a.ctx = ctx

	// 初始化Stage (DiffractService)
	a.Stage.Startup(a.ctx)
	// 初始化HVPS
	a.HVPS.SetContent(a.ctx)
}

func (a *App) APIUpdate() (services.GitHubRelease, error) {
	release, err := a.updater.GetUpdateInfo()
	if err != nil {
		fmt.Printf("获取更新信息失败: %v\n", err)
		return services.GitHubRelease{}, err
	}
	return release, nil
}

func (a *App) GetCachedRelease() services.GitHubRelease {
	return a.updater.GetCachedRelease()
}
