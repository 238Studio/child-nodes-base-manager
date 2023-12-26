package child_nodes_base_manager

import (
	"github.com/238Studio/child-nodes-config-service"
	"github.com/238Studio/child-nodes-database-service"
	"github.com/238Studio/child-nodes-net-service/websocket/service"
)

// BaseModulesGroup 底层模块组
type BaseModulesGroup struct {
	ConfigService    *config.ConfigManager
	DatabaseService  *database.DatabaseAPP
	WebsocketService *service.WebsocketServiceApp
}

// BaseManager 管理器
type BaseManager struct {
	// 存储的底层模块组 packageID->key->BaseModulesGroup
	BaseModulesGroup map[string]*map[string]*BaseModulesGroup
	// 默认底层组
	DefaultBaseModuleGroup *BaseModulesGroup
}

// GetAndRegisterBaseModulesGroup 获取并在同时注册一组底层模块到管理器
// 传入：模块名，key
// 传出：底层模块组
func (manager *BaseManager) GetAndRegisterBaseModulesGroup(moduleName string, key string) *BaseModulesGroup {
	_, isOk := manager.BaseModulesGroup[moduleName]
	if !isOk {
		m := make(map[string]*BaseModulesGroup)
		manager.BaseModulesGroup[moduleName] = &m
	}
	re := BaseModulesGroup{
		ConfigService:    nil,
		DatabaseService:  nil,
		WebsocketService: nil,
	}
	(*manager.BaseModulesGroup[moduleName])[key] = &re
	return &re
}
