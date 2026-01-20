package r2b2

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ServerInfoRouter
	GmCommandRouter
}

var (
	srvApi = api.ApiGroupApp.R2b2ApiGroup.ServerInfoApi
	gmApi  = api.ApiGroupApp.R2b2ApiGroup.GmCommandApi
)
