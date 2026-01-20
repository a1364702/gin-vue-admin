package r2b2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GmCommandRouter struct {}

// InitGmCommandRouter 初始化 GM命令 路由信息
func (s *GmCommandRouter) InitGmCommandRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	gmRouter := Router.Group("gm").Use(middleware.OperationRecord())
	gmRouterWithoutRecord := Router.Group("gm")
	gmRouterWithoutAuth := PublicRouter.Group("gm")
	{
		gmRouter.POST("createGmCommand", gmApi.CreateGmCommand)   // 新建GM命令
		gmRouter.DELETE("deleteGmCommand", gmApi.DeleteGmCommand) // 删除GM命令
		gmRouter.DELETE("deleteGmCommandByIds", gmApi.DeleteGmCommandByIds) // 批量删除GM命令
		gmRouter.PUT("updateGmCommand", gmApi.UpdateGmCommand)    // 更新GM命令
	}
	{
		gmRouterWithoutRecord.GET("findGmCommand", gmApi.FindGmCommand)        // 根据ID获取GM命令
		gmRouterWithoutRecord.GET("getGmCommandList", gmApi.GetGmCommandList)  // 获取GM命令列表
	}
	{
	    gmRouterWithoutAuth.GET("getGmCommandDataSource", gmApi.GetGmCommandDataSource)  // 获取GM命令数据源
	    gmRouterWithoutAuth.GET("getGmCommandPublic", gmApi.GetGmCommandPublic)  // GM命令开放接口
	}
}
