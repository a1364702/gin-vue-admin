package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ServerInfoSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
    ServerStatus  *int `json:"serverStatus" form:"serverStatus"` 
    ServerGroup []string `json:"serverGroup" form:"serverGroup"`
}
