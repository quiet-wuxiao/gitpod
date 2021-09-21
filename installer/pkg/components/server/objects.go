package server

import "github.com/gitpod-io/gitpod/installer/pkg/common"

var Objects = common.CompositeRenderFunc(
	configmap,
	deployment,
	ideconfigmap,
	networkpolicy,
	role,
	rolebinding,
	common.GenerateService(Component),
	common.DefaultServiceAccount(Component),
)
