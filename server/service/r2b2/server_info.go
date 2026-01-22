package r2b2

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/r2b2"
	r2b2Req "github.com/flipped-aurora/gin-vue-admin/server/model/r2b2/request"
)

type ServerInfoService struct {}

var srvService = new(ServerInfoService)

// CreateServerInfo 创建服务器信息记录
// Author [yourname](https://github.com/yourname)
func (srvService *ServerInfoService) CreateServerInfo(ctx context.Context, srv *r2b2.ServerInfo) (err error) {
	err = global.GVA_DB.Create(srv).Error
	return err
}

// DeleteServerInfo 删除服务器信息记录
// Author [yourname](https://github.com/yourname)
func (srvService *ServerInfoService)DeleteServerInfo(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&r2b2.ServerInfo{},"id = ?",ID).Error
	return err
}

// DeleteServerInfoByIds 批量删除服务器信息记录
// Author [yourname](https://github.com/yourname)
func (srvService *ServerInfoService)DeleteServerInfoByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]r2b2.ServerInfo{},"id in ?",IDs).Error
	return err
}

// UpdateServerInfo 更新服务器信息记录
// Author [yourname](https://github.com/yourname)
func (srvService *ServerInfoService)UpdateServerInfo(ctx context.Context, srv r2b2.ServerInfo) (err error) {
	err = global.GVA_DB.Model(&r2b2.ServerInfo{}).Where("id = ?",srv.ID).Updates(&srv).Error
	return err
}

// GetServerInfo 根据ID获取服务器信息记录
// Author [yourname](https://github.com/yourname)
func (srvService *ServerInfoService)GetServerInfo(ctx context.Context, ID string) (srv r2b2.ServerInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&srv).Error
	return srv, err
}

// GetServerInfoByServerId 根据ID获取服务器信息记录
// Author [yourname](https://github.com/yourname)
func (srvService *ServerInfoService)GetServerInfoByServerId(ctx context.Context, ID int64) (srv r2b2.ServerInfo, err error) {
	err = global.GVA_DB.Where(&r2b2.ServerInfo{ServerId: &ID}).First(&srv).Error
	return srv, err
}

// GetServerInfoInfoList 分页获取服务器信息记录
// Author [yourname](https://github.com/yourname)
func (srvService *ServerInfoService)GetServerInfoInfoList(ctx context.Context, info r2b2Req.ServerInfoSearch) (list []r2b2.ServerInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&r2b2.ServerInfo{})
    var srvs []r2b2.ServerInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["id"] = true
           orderMap["created_at"] = true
         	orderMap["server_group"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&srvs).Error
	return  srvs, total, err
}
func (srvService *ServerInfoService)GetServerInfoPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
