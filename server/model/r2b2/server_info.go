// 自动生成模板ServerInfo
package r2b2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 服务器信息 结构体  ServerInfo
type ServerInfo struct {
    global.GVA_MODEL
  ServerId  *int64 `json:"serverId" form:"serverId" gorm:"comment:服务器ID;column:server_id;" binding:"required"`  //ID
  ServerName  *string `json:"serverName" form:"serverName" gorm:"comment:服务器名称;column:server_name;" binding:"required"`  //名称
  ServerGroup  *string `json:"serverGroup" form:"serverGroup" gorm:"comment:服务器组别;column:server_group;" binding:"required"`  //组别
  ServerIp  *string `json:"serverIp" form:"serverIp" gorm:"comment:服务器IP地址;column:server_ip;" binding:"required"`  //IP地址
  ServerPort  *int64 `json:"serverPort" form:"serverPort" gorm:"comment:服务器端口;column:server_port;" binding:"required"`  //端口
  ServerDesc  *string `json:"serverDesc" form:"serverDesc" gorm:"comment:服务器描述;column:server_desc;"`  //描述
  ServerStatus  *int64 `json:"serverStatus" form:"serverStatus" gorm:"comment:服务器当前是什么状态;column:server_status;"`  //服务器状态
}


// TableName 服务器信息 ServerInfo自定义表名 server_info
func (ServerInfo) TableName() string {
    return "server_info"
}





