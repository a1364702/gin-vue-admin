package r2b2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ServerInfoRouter struct {}

// InitServerInfoRouter 初始化 服务器信息 路由信息
func (s *ServerInfoRouter) InitServerInfoRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	srvRouter := Router.Group("srv").Use(middleware.OperationRecord())
	srvRouterWithoutRecord := Router.Group("srv")
	srvRouterWithoutAuth := PublicRouter.Group("srv")
	{
		srvRouter.POST("createServerInfo", srvApi.CreateServerInfo)   // 新建服务器信息
		srvRouter.DELETE("deleteServerInfo", srvApi.DeleteServerInfo) // 删除服务器信息
		srvRouter.DELETE("deleteServerInfoByIds", srvApi.DeleteServerInfoByIds) // 批量删除服务器信息
		srvRouter.PUT("updateServerInfo", srvApi.UpdateServerInfo)    // 更新服务器信息
	}
	{
		srvRouterWithoutRecord.GET("findServerInfo", srvApi.FindServerInfo)        // 根据ID获取服务器信息
		srvRouterWithoutRecord.GET("getServerInfoList", srvApi.GetServerInfoList)  // 获取服务器信息列表
	}
	{
	    srvRouterWithoutAuth.GET("getServerInfoPublic", srvApi.GetServerInfoPublic)  // 服务器信息开放接口
	}
}
