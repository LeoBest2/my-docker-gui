package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
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

func (a *App) ImageList() ([]Image, error) {
	images, err := Cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return nil, err
	}
	var imagesRet []Image
	for _, i := range images {
		iID := strings.Index(i.ID, ":")
		iRepo := strings.LastIndex(i.RepoTags[0], ":")
		imagesRet = append(imagesRet, Image{
			ID:         i.ID[iID+1 : iID+13],
			Repository: i.RepoTags[0][:iRepo],
			Tag:        i.RepoTags[0][iRepo+1:],
			Created:    time.Unix(i.Created, 0).Format("2006-01-02 15:04:05"),
			Size:       fmt.Sprintf("%dMB", i.Size/1024/1024),
		})
	}
	return imagesRet, nil
}

func (a *App) ImageDelete(ID string) error {
	_, err := Cli.ImageRemove(context.Background(), ID, types.ImageRemoveOptions{Force: true})
	return err
}

func (a *App) Info() (types.Info, error) {
	return Cli.Info(context.Background())
}

func (a *App) Log(ID string) (string, error) {
	reader, err := Cli.ContainerLogs(context.Background(), ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(b), nil
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

type Image struct {
	ID         string `json:"ID"`
	Repository string `json:"Repository"`
	Tag        string `json:"Tag"`
	Created    string `json:"Created"`
	Size       string `json:"Size"`
}
