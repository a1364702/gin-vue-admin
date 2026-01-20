package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/r2b2"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(r2b2.ServerInfo{}, r2b2.GmCommand{})
	if err != nil {
		return err
	}
	return nil
}
