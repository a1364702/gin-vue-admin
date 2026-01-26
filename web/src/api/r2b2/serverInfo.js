import service from '@/utils/request'
// @Tags ServerInfo
// @Summary 创建服务器信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerInfo true "创建服务器信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /srv/createServerInfo [post]
export const createServerInfo = (data) => {
  return service({
    url: '/srv/createServerInfo',
    method: 'post',
    data
  })
}

// @Tags ServerInfo
// @Summary 删除服务器信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerInfo true "删除服务器信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /srv/deleteServerInfo [delete]
export const deleteServerInfo = (params) => {
  return service({
    url: '/srv/deleteServerInfo',
    method: 'delete',
    params
  })
}

// @Tags ServerInfo
// @Summary 批量删除服务器信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除服务器信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /srv/deleteServerInfo [delete]
export const deleteServerInfoByIds = (params) => {
  return service({
    url: '/srv/deleteServerInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags ServerInfo
// @Summary 更新服务器信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerInfo true "更新服务器信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /srv/updateServerInfo [put]
export const updateServerInfo = (data) => {
  return service({
    url: '/srv/updateServerInfo',
    method: 'put',
    data
  })
}

// @Tags ServerInfo
// @Summary 用id查询服务器信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ServerInfo true "用id查询服务器信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /srv/findServerInfo [get]
export const findServerInfo = (params) => {
  return service({
    url: '/srv/findServerInfo',
    method: 'get',
    params
  })
}

// @Tags ServerInfo
// @Summary 分页获取服务器信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取服务器信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /srv/getServerInfoList [get]
export const getServerInfoList = (params) => {
  return service({
    url: '/srv/getServerInfoList',
    method: 'get',
    params
  })
}

// @Tags ServerInfo
// @Summary 获取服务器分组信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "获取服务器分组信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /srv/getServerGroup [get]
export const getServerGroup = (params) => {
  return service({
    url: '/srv/getServerGroup',
    method: 'get',
    params
  })
}

// @Tags ServerInfo
// @Summary 不需要鉴权的服务器信息接口
// @Accept application/json
// @Produce application/json
// @Param data query r2b2Req.ServerInfoSearch true "分页获取服务器信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /srv/getServerInfoPublic [get]
export const getServerInfoPublic = () => {
  return service({
    url: '/srv/getServerInfoPublic',
    method: 'get',
  })
}
