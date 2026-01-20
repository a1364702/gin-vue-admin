package r2b2

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ServerInfoApi
	GmCommandApi
}

var (
	srvService = service.ServiceGroupApp.R2b2ServiceGroup.ServerInfoService
	gmService  = service.ServiceGroupApp.R2b2ServiceGroup.GmCommandService
)
