package r2b2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/r2b2"
	r2b2Req "github.com/flipped-aurora/gin-vue-admin/server/model/r2b2/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GmCommandApi struct {}



// CreateGmCommand 创建GM命令
// @Tags GmCommand
// @Summary 创建GM命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body r2b2.GmCommand true "创建GM命令"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /gm/createGmCommand [post]
func (gmApi *GmCommandApi) CreateGmCommand(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var gm r2b2.GmCommand
	err := c.ShouldBindJSON(&gm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = gmService.CreateGmCommand(ctx,&gm)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteGmCommand 删除GM命令
// @Tags GmCommand
// @Summary 删除GM命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body r2b2.GmCommand true "删除GM命令"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /gm/deleteGmCommand [delete]
func (gmApi *GmCommandApi) DeleteGmCommand(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := gmService.DeleteGmCommand(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteGmCommandByIds 批量删除GM命令
// @Tags GmCommand
// @Summary 批量删除GM命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /gm/deleteGmCommandByIds [delete]
func (gmApi *GmCommandApi) DeleteGmCommandByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := gmService.DeleteGmCommandByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateGmCommand 更新GM命令
// @Tags GmCommand
// @Summary 更新GM命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body r2b2.GmCommand true "更新GM命令"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /gm/updateGmCommand [put]
func (gmApi *GmCommandApi) UpdateGmCommand(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var gm r2b2.GmCommand
	err := c.ShouldBindJSON(&gm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = gmService.UpdateGmCommand(ctx,gm)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindGmCommand 用id查询GM命令
// @Tags GmCommand
// @Summary 用id查询GM命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询GM命令"
// @Success 200 {object} response.Response{data=r2b2.GmCommand,msg=string} "查询成功"
// @Router /gm/findGmCommand [get]
func (gmApi *GmCommandApi) FindGmCommand(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	regm, err := gmService.GetGmCommand(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(regm, c)
}
// GetGmCommandList 分页获取GM命令列表
// @Tags GmCommand
// @Summary 分页获取GM命令列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query r2b2Req.GmCommandSearch true "分页获取GM命令列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /gm/getGmCommandList [get]
func (gmApi *GmCommandApi) GetGmCommandList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo r2b2Req.GmCommandSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := gmService.GetGmCommandInfoList(ctx,pageInfo)
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
// GetGmCommandDataSource 获取GmCommand的数据源
// @Tags GmCommand
// @Summary 获取GmCommand的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /gm/getGmCommandDataSource [get]
func (gmApi *GmCommandApi) GetGmCommandDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
    dataSource, err := gmService.GetGmCommandDataSource(ctx)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
   		response.FailWithMessage("查询失败:" + err.Error(), c)
   		return
    }
   response.OkWithData(dataSource, c)
}

// GetGmCommandPublic 不需要鉴权的GM命令接口
// @Tags GmCommand
// @Summary 不需要鉴权的GM命令接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /gm/getGmCommandPublic [get]
func (gmApi *GmCommandApi) GetGmCommandPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    gmService.GetGmCommandPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的GM命令接口信息",
    }, "获取成功", c)
}
// ExcuteCommand 执行一个gm方法
// @Tags GmCommand
// @Summary 执行一个gm方法
// @Accept application/json
// @Produce application/json
// @Param data query r2b2Req.GmCommandSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /gm/excuteCommand [POST]
func (gmApi *GmCommandApi)ExcuteCommand(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()
    // 请添加自己的业务逻辑
	var gm r2b2.GmCommand
	err := c.ShouldBindJSON(&gm)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	var result string
    result, err = gmService.ExcuteCommand(ctx, &gm)
    if err != nil {
        global.GVA_LOG.Error("失败!", zap.Error(err))
   		response.FailWithMessage("失败", c)
   		return
   	}
   	response.OkWithData(result,c)
}


