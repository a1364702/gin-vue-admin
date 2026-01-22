package r2b2

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/r2b2"
	r2b2Req "github.com/flipped-aurora/gin-vue-admin/server/model/r2b2/request"
)

type GmCommandService struct {}

var gmService = new(GmCommandService)
// CreateGmCommand 创建GM命令记录
// Author [yourname](https://github.com/yourname)
func (gmService *GmCommandService) CreateGmCommand(ctx context.Context, gm *r2b2.GmCommand) (err error) {
	err = global.GVA_DB.Create(gm).Error
	return err
}

// DeleteGmCommand 删除GM命令记录
// Author [yourname](https://github.com/yourname)
func (gmService *GmCommandService)DeleteGmCommand(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&r2b2.GmCommand{},"id = ?",ID).Error
	return err
}

// DeleteGmCommandByIds 批量删除GM命令记录
// Author [yourname](https://github.com/yourname)
func (gmService *GmCommandService)DeleteGmCommandByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]r2b2.GmCommand{},"id in ?",IDs).Error
	return err
}

// UpdateGmCommand 更新GM命令记录
// Author [yourname](https://github.com/yourname)
func (gmService *GmCommandService)UpdateGmCommand(ctx context.Context, gm r2b2.GmCommand) (err error) {
	err = global.GVA_DB.Model(&r2b2.GmCommand{}).Where("id = ?",gm.ID).Updates(&gm).Error
	return err
}

// GetGmCommand 根据ID获取GM命令记录
// Author [yourname](https://github.com/yourname)
func (gmService *GmCommandService)GetGmCommand(ctx context.Context, ID string) (gm r2b2.GmCommand, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&gm).Error
	return
}
// GetGmCommandInfoList 分页获取GM命令记录
// Author [yourname](https://github.com/yourname)
func (gmService *GmCommandService)GetGmCommandInfoList(ctx context.Context, info r2b2Req.GmCommandSearch) (list []r2b2.GmCommand, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&r2b2.GmCommand{})
    var gms []r2b2.GmCommand
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Command != nil && *info.Command != "" {
        db = db.Where("command LIKE ?", "%"+ *info.Command+"%")
    }
			if len(info.ExecuteTimeRange) == 2 {
				db = db.Where("execute_time BETWEEN ? AND ? ", info.ExecuteTimeRange[0], info.ExecuteTimeRange[1])
			}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["id"] = true
           orderMap["created_at"] = true
         	orderMap["execute_time"] = true
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

	err = db.Find(&gms).Error
	return  gms, total, err
}
func (gmService *GmCommandService)GetGmCommandDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   executeServer := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("server_info").Where("deleted_at IS NULL").Select("server_name as label,id as value").Scan(&executeServer)
	   res["executeServer"] = executeServer
	   executor := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("sys_users").Where("deleted_at IS NULL").Select("nick_name as label,username as value").Scan(&executor)
	   res["executor"] = executor
	return
}
func (gmService *GmCommandService)GetGmCommandPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}

// ExcuteCommand 执行一个gm方法
// Author [yourname](https://github.com/yourname)
func (gmService *GmCommandService)ExcuteCommand(ctx context.Context, gm *r2b2.GmCommand) (result string, err error) {
	result = ""
	var serverInfos []r2b2.ServerInfo
	for _, zoneId := range gm.ExecuteServer{
		serverInfo, err := srvService.GetServerInfoByServerId(ctx,  zoneId)
		if err != nil {
			return result, err
		}
		serverInfos = append(serverInfos, serverInfo)
	}

	for _, serverInfo := range serverInfos{
		return "执行了命令" + *gm.Command + " 在" + *serverInfo.ServerName  + "中", nil
	}
    return "success", nil
}


