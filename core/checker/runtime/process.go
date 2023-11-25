package runtime

import (
	"consecure/util/log"
	"consecure/util/process"
	"context"
	"errors"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type ContainerProcesInfo struct {
	Pid           int
	ContainerId   string
	ContainerName string
	ContainerType string
	EntryPointPid int
}

type RuntimeContainerProcess struct {
	client                *client.Client
	ContainerProcessInfos []ContainerProcesInfo
}

func NewRuntimeContainerProcess() *RuntimeContainerProcess {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal("Error creating docker client", err)
	}

	instance := &RuntimeContainerProcess{
		client: cli,
	}
	instance.Init()

	return instance
}

func (r *RuntimeContainerProcess) Init() {
	containers, err := r.client.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		return
	}

	for _, container := range containers {
		info, err := r.createContainerInfo(container)

		log.Debugln("Detected Container Process", info.ContainerName, info.ContainerId, info.Pid, info.EntryPointPid)

		if err != nil {
			break
		}

		r.ContainerProcessInfos = append(r.ContainerProcessInfos, info)
	}
}

func (r *RuntimeContainerProcess) AddProcessInfo(pid int, cmdLines []string) {
	containerId := cmdLines[4]

	container, err := r.GetContainerInfoFromId(pid, containerId)

	if err != nil {
		return
	}

	log.Debugln("Detected New Container Process", container.ContainerName, container.ContainerId, container.Pid)
	r.ContainerProcessInfos = append(r.ContainerProcessInfos, container)
}

func (r *RuntimeContainerProcess) GetContainerInfoFromId(pid int, containerId string) (ContainerProcesInfo, error) {
	containers, err := r.client.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		return ContainerProcesInfo{}, err
	}

	for _, container := range containers {
		if container.ID == containerId {
			info, err := r.createContainerInfo(container)

			if err != nil {
				break
			}

			return info, nil
		}
	}

	return ContainerProcesInfo{}, errors.New("Container not found")
}

func (r *RuntimeContainerProcess) createContainerInfo(container types.Container) (ContainerProcesInfo, error) {
	inspect, err := r.client.ContainerInspect(context.Background(), container.ID)

	if err != nil {
		return ContainerProcesInfo{}, err
	}

	entryPointPid := inspect.State.Pid
	containerPid := process.GetParentPid(entryPointPid)

	return ContainerProcesInfo{
		Pid:           containerPid,
		ContainerId:   container.ID,
		ContainerName: container.Names[0],
		ContainerType: "docker",
		EntryPointPid: inspect.State.Pid,
	}, nil
}

func (r *RuntimeContainerProcess) IsContainerPid(pid int, parentPid int) bool {
	for _, containerProcessInfo := range r.ContainerProcessInfos {
		if containerProcessInfo.Pid == parentPid && containerProcessInfo.EntryPointPid != pid {
			return true
		}
	}
	return false
}
