package main

import (
	"fmt"

	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func runExporterContainer(ctx *pulumi.Context) error {
	exporter, err := docker.NewRemoteImage(ctx, "exporter", &docker.RemoteImageArgs{
		Name: pulumi.String(fmt.Sprintf("%s:%s", EXPORTER_IMAGE, EXPORTER_TAG)),
	})
	if err != nil {
		return err
	}

	_, err = docker.NewContainer(ctx, "exporter", &docker.ContainerArgs{
		Image: exporter.ImageId,
		Name:  pulumi.String("exporter"),
		Ports:  docker.ContainerPortArray{
			&docker.ContainerPortArgs{
				Internal: pulumi.Int(EXPORTER_PORT),
				External: pulumi.Int(EXPORTER_PORT),
			},
		},
		Envs: pulumi.StringArray{
			pulumi.Sprintf("GITLAB_API_TOKEN=%s", EXPORTER_API_TOKEN),
		},
		NetworksAdvanced: docker.ContainerNetworksAdvancedArray{
			&docker.ContainerNetworksAdvancedArgs{
				Name: pulumi.String(DOCKER_NETWORK_NAME),
			},
		},
		Rm: pulumi.BoolPtr(true),
	})
	return err
}
