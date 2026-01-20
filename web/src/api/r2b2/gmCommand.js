import service from '@/utils/request'
// @Tags GmCommand
// @Summary 创建GM命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GmCommand true "创建GM命令"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /gm/createGmCommand [post]
export const createGmCommand = (data) => {
  return service({
    url: '/gm/createGmCommand',
    method: 'post',
    data
  })
}

// @Tags GmCommand
// @Summary 删除GM命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GmCommand true "删除GM命令"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gm/deleteGmCommand [delete]
export const deleteGmCommand = (params) => {
  return service({
    url: '/gm/deleteGmCommand',
    method: 'delete',
    params
  })
}

// @Tags GmCommand
// @Summary 批量删除GM命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GM命令"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gm/deleteGmCommand [delete]
export const deleteGmCommandByIds = (params) => {
  return service({
    url: '/gm/deleteGmCommandByIds',
    method: 'delete',
    params
  })
}

// @Tags GmCommand
// @Summary 更新GM命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GmCommand true "更新GM命令"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gm/updateGmCommand [put]
export const updateGmCommand = (data) => {
  return service({
    url: '/gm/updateGmCommand',
    method: 'put',
    data
  })
}

// @Tags GmCommand
// @Summary 用id查询GM命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.GmCommand true "用id查询GM命令"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gm/findGmCommand [get]
export const findGmCommand = (params) => {
  return service({
    url: '/gm/findGmCommand',
    method: 'get',
    params
  })
}

// @Tags GmCommand
// @Summary 分页获取GM命令列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GM命令列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gm/getGmCommandList [get]
export const getGmCommandList = (params) => {
  return service({
    url: '/gm/getGmCommandList',
    method: 'get',
    params
  })
}
// @Tags GmCommand
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gm/findGmCommandDataSource [get]
export const getGmCommandDataSource = () => {
  return service({
    url: '/gm/getGmCommandDataSource',
    method: 'get',
  })
}

// @Tags GmCommand
// @Summary 不需要鉴权的GM命令接口
// @Accept application/json
// @Produce application/json
// @Param data query r2b2Req.GmCommandSearch true "分页获取GM命令列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /gm/getGmCommandPublic [get]
export const getGmCommandPublic = () => {
  return service({
    url: '/gm/getGmCommandPublic',
    method: 'get',
  })
}
