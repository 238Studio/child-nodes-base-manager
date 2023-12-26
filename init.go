package child_nodes_base_manager

import (
	config "github.com/238Studio/child-nodes-config-service"
	database "github.com/238Studio/child-nodes-database-service"
	"github.com/238Studio/child-nodes-net-service/websocket/service"
	"strconv"
)

// Init 初始化一个管理器
// 传入：
// 传出：
func Init() *BaseManager {
	return &BaseManager{
		BaseModulesGroup:       make(map[string]*map[string]*BaseModulesGroup),
		DefaultBaseModuleGroup: nil,
	}
}

// InitDefaultBaseModules 初始化一套默认底层组
// 传入：无
// 传出：无
func (manager *BaseManager) InitDefaultBaseModules() error {
	baseModulesGroup := manager.DefaultBaseModuleGroup
	configService := config.InitConfigManager("/config")
	baseModulesGroup.ConfigService = configService
	var err error
	var databaseName string
	var databaseURL string
	databaseName, err = configService.ReadConfig("default", "databaseName")
	if err != nil {
		return err
	}
	databaseURL, err = configService.ReadConfig("default", "databaseURL")
	if err != nil {
		return err
	}
	baseModulesGroup.DatabaseService, err = database.InitSQLiteDatabase(databaseName, databaseURL)
	if err != nil {
		return err
	}
	var defaultWsURL string
	var pingTimerMargin int
	var pongWait int
	defaultWsURL, err = configService.ReadConfig("default", "WsURL")
	if err != nil {
		return err
	}
	pingTimerMargin_, err := configService.ReadConfig("default", "pingTimerMargin")
	if err != nil {
		return err
	}
	pingTimerMargin__, err := strconv.ParseInt(pingTimerMargin_, 10, 64)
	pingTimerMargin = (int)(pingTimerMargin__)
	if err != nil {
		return err
	}
	pongWait_, err := configService.ReadConfig("default", "pongWait")
	pongWait__, err := strconv.ParseInt(pongWait_, 10, 64)
	pongWait = (int)(pongWait__)
	baseModulesGroup.WebsocketService, err = service.Init(defaultWsURL, pingTimerMargin, pongWait)
	return err
}
