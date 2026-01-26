package r2b2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/r2b2"
	r2b2Req "github.com/flipped-aurora/gin-vue-admin/server/model/r2b2/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ServerInfoApi struct {}



// CreateServerInfo 创建服务器信息
// @Tags ServerInfo
// @Summary 创建服务器信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body r2b2.ServerInfo true "创建服务器信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /srv/createServerInfo [post]
func (srvApi *ServerInfoApi) CreateServerInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var srv r2b2.ServerInfo
	err := c.ShouldBindJSON(&srv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = srvService.CreateServerInfo(ctx,&srv)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteServerInfo 删除服务器信息
// @Tags ServerInfo
// @Summary 删除服务器信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body r2b2.ServerInfo true "删除服务器信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /srv/deleteServerInfo [delete]
func (srvApi *ServerInfoApi) DeleteServerInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := srvService.DeleteServerInfo(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteServerInfoByIds 批量删除服务器信息
// @Tags ServerInfo
// @Summary 批量删除服务器信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /srv/deleteServerInfoByIds [delete]
func (srvApi *ServerInfoApi) DeleteServerInfoByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := srvService.DeleteServerInfoByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateServerInfo 更新服务器信息
// @Tags ServerInfo
// @Summary 更新服务器信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body r2b2.ServerInfo true "更新服务器信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /srv/updateServerInfo [put]
func (srvApi *ServerInfoApi) UpdateServerInfo(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var srv r2b2.ServerInfo
	err := c.ShouldBindJSON(&srv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = srvService.UpdateServerInfo(ctx,srv)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindServerInfo 用id查询服务器信息
// @Tags ServerInfo
// @Summary 用id查询服务器信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询服务器信息"
// @Success 200 {object} response.Response{data=r2b2.ServerInfo,msg=string} "查询成功"
// @Router /srv/findServerInfo [get]
func (srvApi *ServerInfoApi) FindServerInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	resrv, err := srvService.GetServerInfo(ctx, ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(resrv, c)
}

// GetServerInfoList 分页获取服务器信息列表
// @Tags ServerInfo
// @Summary 分页获取服务器信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query r2b2Req.ServerInfoSearch true "分页获取服务器信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /srv/getServerInfoList [get]
func (srvApi *ServerInfoApi) GetServerInfoList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo r2b2Req.ServerInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := srvService.GetServerInfoInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetServerGroup 获取服务器分组信息
// @Tags ServerInfo
// @Summary 获取服务器分组信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /srv/GetServerGroup [get]
func (srvApi *ServerInfoApi) GetServerGroup(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo r2b2Req.ServerInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := srvService.GetServerGroup(ctx)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
    }, "获取成功", c)
}

// GetServerInfoPublic 不需要鉴权的服务器信息接口
// @Tags ServerInfo
// @Summary 不需要鉴权的服务器信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /srv/getServerInfoPublic [get]
func (srvApi *ServerInfoApi) GetServerInfoPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    srvService.GetServerInfoPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的服务器信息接口信息",
    }, "获取成功", c)
}
