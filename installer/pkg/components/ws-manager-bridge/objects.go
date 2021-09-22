package wsmanagerbridge

import "github.com/gitpod-io/gitpod/installer/pkg/common"

var Objects = common.CompositeRenderFunc(
	deployment,
	rolebinding,
	common.DefaultServiceAccount(Component),
)
