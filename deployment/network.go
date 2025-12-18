package main

import (

	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createNetwork(ctx *pulumi.Context) error {
	_, err := docker.NewNetwork(ctx, DOCKER_NETWORK_NAME, &docker.NetworkArgs{
			Name: pulumi.String(DOCKER_NETWORK_NAME),
		})
	
	return err
}
