package login

import (
	"github.com/ouczbs/Zmin/component/base"
	"github.com/ouczbs/Zmin/engine/zattr"
	"github.com/ouczbs/Zmin/engine/zcache"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/znet"
	"mmo/engine/conf"
	"net"
)

type ULoginService struct {
	*UService
}
func NewLoginService() *ULoginService {
	zcache.InitMongoClient("mongodb://111.229.54.9:27017" , "mmo")
	return &ULoginService{
		UService: base.NewService(reqHandleMaps),
	}
}
func (service *ULoginService)Run() {
	service.UService.Run()
	service.initService()
	go service.MessageLoop()
	service.ConnectToCenter()
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *ULoginService) initService() {
	logFile, ok := service.GetProperty(zattr.StringLogFile).(string)
	if !ok {
		logFile = zconf.LoginConfig.LogFile
	}
	zlog.SetOutput([]string{"stderr", logFile})
	service.InitDownHandles()
}
func (service *ULoginService) NewTcpConnection(conn net.Conn) {
	proxy := znet.NewClientProxy(service, conn)
	zattr.SetRemoteProperty(proxy,zattr.Int32MessageType , int32(conf.MT_Game_Login))
	proxy.Serve()
}