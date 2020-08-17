package dockerTool

import (
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func NewClient(host string, port string, useLTS bool, tlsCaCert string, tlsCert string, tlsKey string) (context.Context, *client.Client, types.Info, error) {
	ctx := context.Background()
	var cli *client.Client
	err := errors.New("")
	if useLTS {
		cli, err = NewTlsClient(host, port, tlsCaCert, tlsCert, tlsKey)
		if err != nil {
			return ctx, &client.Client{}, types.Info{}, err
		}
	} else {
		cli, err = NewNoTlsClient(host, port)
		if err != nil {
			return ctx, &client.Client{}, types.Info{}, err
		}
	}
	info, err := cli.Info(ctx)
	if err != nil {
		return ctx, &client.Client{}, types.Info{}, err
	}
	return ctx, cli, info, nil
}

func NewNoTlsClient(Host string, port string) (*client.Client, error) {
	clientUrl := "tcp://" + Host + ":" + port
	cli, err := client.NewClientWithOpts(client.WithHost(clientUrl), client.WithAPIVersionNegotiation())
	if err != nil {
		return &client.Client{}, err
	}
	return cli, nil
}

func NewTlsClient(host string, port string, tlsCaCert string, tlsCert string, tlsKey string) (*client.Client, error) {
	var cli *client.Client
	return cli, nil
}
