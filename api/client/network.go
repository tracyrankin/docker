// +build experimental

package client

import (
	nwclient "github.com/docker/libnetwork/client"
)

// CmdNetwork is used to create, display and configure network endpoints.
func (cli *DockerCli) CmdNetwork(args ...string) error {
	nCli := nwclient.NewNetworkCli(cli.out, cli.err, nwclient.CallFunc(cli.callWrapper))
	args = append([]string{"network"}, args...)
	return nCli.Cmd("docker", args...)
}
