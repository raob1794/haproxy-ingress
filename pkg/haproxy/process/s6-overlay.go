package process

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/haproxytech/kubernetes-ingress/pkg/haproxy/api"
	"github.com/haproxytech/kubernetes-ingress/pkg/haproxy/env"
	"github.com/haproxytech/kubernetes-ingress/pkg/utils"
)

type s6Control struct {
	API    api.HAProxyClient
	Env    env.Env
	OSArgs utils.OSArgs
}

func (d *s6Control) Service(action string) error {
	if d.OSArgs.Test {
		logger.Infof("HAProxy would be %sed now", action)
		return nil
	}
	var cmd *exec.Cmd

	switch action {
	case "start":
		// no need to start it is up already (s6)
		return nil
	case "stop":
		// no need to stop it (s6)
		return nil
	case "reload":
		cmd = exec.Command("s6-svc", "-2", "/run/service/haproxy")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "restart":
		cmd = exec.Command("s6-svc", "-t", "/run/service/haproxy")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unknown command '%s'", action)
	}
}

func (d *s6Control) UseAuxFile(useAuxFile bool) {
	// do nothing we always have it
}

func (d *s6Control) SetAPI(api api.HAProxyClient) {
	d.API = api
}
