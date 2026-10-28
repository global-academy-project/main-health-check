package handler

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func ContainerList() error {
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("create Docker client: %w", err)
	}
	defer docker.Close()

	ctx := context.Background()
	containers, err := docker.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		return fmt.Errorf("list Docker containers: %w", err)
	}

	for _, container := range containers {
		fmt.Println(container.ID, container.Names)
	}
	return nil
}

func ContainerUsageResource() {
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("create Docker client: %v\n", err)
		return
	}
	defer docker.Close()
}

func ContainerUsageCPU() {
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("create Docker client: %v\n", err)
		return
	}
	defer docker.Close()
}

func ContainerUsageMemory() {
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("create Docker client: %v\n", err)
		return
	}
	defer docker.Close()
}

func ContainerUsageNetwork() {
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("create Docker client: %v\n", err)
		return
	}
	defer docker.Close()
}

func ContainerUsageDisk() {
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("create Docker client: %v\n", err)
		return
	}
	defer docker.Close()
}
