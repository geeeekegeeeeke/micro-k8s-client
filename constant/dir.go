package constant

import (
	"micro-k8s-client/global"
	"path"
)

var (
	DataDir              = global.CONF.System.DataDir
	ResourceDir          = path.Join(DataDir, "resource")
	AppResourceDir       = path.Join(ResourceDir, "apps")
	AppInstallDir        = path.Join(DataDir, "apps")
	LocalAppResourceDir  = path.Join(AppResourceDir, "local")
	LocalAppInstallDir   = path.Join(AppInstallDir, "local")
	RemoteAppResourceDir = path.Join(AppResourceDir, "remote")
	RuntimeDir           = path.Join(DataDir, "runtime")
)
