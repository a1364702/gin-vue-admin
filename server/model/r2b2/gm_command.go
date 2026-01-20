
// 自动生成模板GmCommand
package r2b2
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// GM命令 结构体  GmCommand
type GmCommand struct {
    global.GVA_MODEL
  Command  *string `json:"command" form:"command" gorm:"comment:GM命令内容;column:command;" binding:"required"`  //命令内容
  Executor  *string `json:"executor" form:"executor" gorm:"primarykey;comment:命令执行人;column:executor;" binding:"required"`  //执行人
  Result  *string `json:"result" form:"result" gorm:"comment:命令执行结果;column:result;type:text;"`  //执行结果
  ExecuteTime  *time.Time `json:"executeTime" form:"executeTime" gorm:"primarykey;comment:命令执行时间;column:execute_time;" binding:"required"`  //执行时间
  ExecuteServer  *int64 `json:"executeServer" form:"executeServer" gorm:"comment:在哪个服务器执行;column:execute_server;" binding:"required"`  //服务器
}


// TableName GM命令 GmCommand自定义表名 gm_commands
func (GmCommand) TableName() string {
    return "gm_commands"
}





