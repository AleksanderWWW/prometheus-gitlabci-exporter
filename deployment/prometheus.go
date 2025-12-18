package main

import (
	"fmt"
	"os"
	"path"

	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func runPrometheusContainer(ctx *pulumi.Context) error {
	prometheus, err := docker.NewRemoteImage(ctx, "prometheus", &docker.RemoteImageArgs{
		Name: pulumi.String(fmt.Sprintf("%s:%s", PROMETHEUS_IMAGE, PROMETHEUS_TAG)),
	})
	if err != nil {
		return err
	}

	volume, err := docker.NewVolume(ctx, "prom-vol", &docker.VolumeArgs{Name: pulumi.String("prom-vol")})
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = docker.NewContainer(ctx, "prometheus", &docker.ContainerArgs{
		Image: prometheus.ImageId,
		Name:  pulumi.String("prometheus"),
		Ports: docker.ContainerPortArray{
			&docker.ContainerPortArgs{
				Internal: pulumi.Int(PROMETHEUS_PORT),
				External: pulumi.Int(PROMETHEUS_PORT),
			},
		},
		NetworksAdvanced: docker.ContainerNetworksAdvancedArray{
			&docker.ContainerNetworksAdvancedArgs{
				Name: pulumi.String(DOCKER_NETWORK_NAME),
			},
		},
		Rm: pulumi.BoolPtr(false),
		Mounts: docker.ContainerMountArray{
			&docker.ContainerMountArgs{
				Target:   pulumi.String("/etc/prometheus/prometheus.yml"),
				Type:     pulumi.String("bind"),
				ReadOnly: pulumi.Bool(true),
				Source:   pulumi.String(path.Join(cwd, "config", "prometheus.yml")),
			},
		},
		Volumes: docker.ContainerVolumeArray{
			&docker.ContainerVolumeArgs{
				ContainerPath: pulumi.String("/prometheus"),
				ReadOnly:      pulumi.Bool(false),
				VolumeName:    volume.Name,
			},
		},
	})
	return err
}
