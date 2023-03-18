package main

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var Cli *client.Client

func init() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	Cli = cli
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ContainerList() ([]Container, error) {
	containers, err := Cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	containersRet := make([]Container, 0)
	for _, c := range containers {
		containersRet = append(containersRet, Container{
			ID:      c.ID[:12],
			Names:   c.Names,
			State:   c.State,
			Image:   c.Image,
			Command: c.Command,
			Created: time.Unix(c.Created, 0).Format("2006-01-02 15:04:05"),
			Status:  c.Status,
			Ports:   c.Ports,
		})
	}
	return containersRet, nil
}

func (a *App) ContainerStart(ID string) error {
	return Cli.ContainerStart(context.Background(), ID, types.ContainerStartOptions{})
}

func (a *App) ContainerStop(ID string) error {
	return Cli.ContainerStop(context.Background(), ID, container.StopOptions{})
}

func (a *App) ContainerDelete(ID string) error {
	return Cli.ContainerRemove(context.Background(), ID, types.ContainerRemoveOptions{Force: true})
}

type Container struct {
	ID      string       `json:"ID"`
	Names   []string     `json:"Names"`
	State   string       `json:"State"`
	Image   string       `json:"Image"`
	Command string       `json:"Command"`
	Created string       `json:"Created"`
	Status  string       `json:"Status"`
	Ports   []types.Port `json:"Ports"`
}