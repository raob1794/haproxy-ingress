package global

import (
	"runtime"
	"strconv"

	"github.com/haproxytech/client-native/v5/models"

	"github.com/haproxytech/kubernetes-ingress/pkg/annotations/common"
	"github.com/haproxytech/kubernetes-ingress/pkg/store"
)

type Nbthread struct {
	global *models.Global
	name   string
}

func NewNbthread(n string, g *models.Global) *Nbthread {
	return &Nbthread{name: n, global: g}
}

func (a *Nbthread) GetName() string {
	return a.name
}

func (a *Nbthread) Process(k store.K8s, annotations ...map[string]string) error {
	input := common.GetValue(a.GetName(), annotations...)
	if input == "" {
		a.global.Nbthread = 0
		return nil
	}

	v, err := strconv.Atoi(input)
	if err != nil {
		return err
	}
	maxProcs := runtime.GOMAXPROCS(0)
	if v > maxProcs {
		v = maxProcs
	}
	a.global.Nbthread = int64(v)
	return nil
}
