package login

import (
	"github.com/ouczbs/Zmin/engine/zattr"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/znet"
	"github.com/ouczbs/Zmin/engine/zproto"
	"github.com/ouczbs/Zmin/engine/zproto/pb"
	"mmo/engine/proto"
)

func (service ULoginService) MessageLoop() {
	for {
		select {
		case message := <-service.MessageQueue:
			proxy := message.Proxy
			messageType := message.MessageType
			packet := message.Packet
			cmd := message.Cmd
			switch messageType {
			case zconf.MT_FROM_CENTER, zconf.MT_BROADCAST:
				zproto.PbMessageHandle(proxy, packet , cmd)
			default:
				id := packet.SubtractComponentId()
				proto.PbMessageHandle(proxy, packet , cmd , id)
			}
			//case <-service.ticker:
			//	post.Tick()
			//	service.sendEntitySyncInfosToGames()
			//	break
			//	default:
		}
	}
}
func (service *ULoginService) InitDownHandles() {
	service.UService.InitDownHandles()
	service.InitGameDownHandles()
	reqHandleMaps[TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT)] = service.AddEngineComponent
	reqHandleMaps[TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT_ACK)] = service.AddEngineComponentAck
}
func (service *ULoginService) ConnectToCenter() {
	centerProxy = service.MakeCenterProxy()
	if centerProxy == nil {
		service.Close()
	}
	message := &pb.ADD_ENGINE_COMPONENT{
		ComponentId: service.Config.ComponentId,
		Type:       pb.COMPONENT_TYPE_GAME,
		ListenAddr: service.Config.ListenAddr,
	}
	request := znet.NewRequest(TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT), zconf.MT_TO_CENTER , message)
	zproto.SendPbMessage(centerProxy, request)
	request.Release()
}
func (service *ULoginService) AddEngineComponentAck(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*pb.ADD_ENGINE_COMPONENT_ACK)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	proxy.SetProperty(zattr.Int32ComponentId, int32(message.ComponentId))
	proxy.SetProperty(zattr.Int32ComponentType, int32(pb.COMPONENT_TYPE_CENTER))
	for _, login := range message.ComponentList {
		dispatcherProxyMaps[login.ComponentId] = service.MakeClientProxy(string(login.ListenAddr),pb.COMPONENT_TYPE_DISPATCHER)
	}
}
func (service *ULoginService) AddEngineComponent(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*pb.ADD_ENGINE_COMPONENT)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	dispatcherProxyMaps[message.ComponentId] = service.MakeClientProxy(string(message.ListenAddr),pb.COMPONENT_TYPE_DISPATCHER)
	zlog.Debug("AddEngineComponent " , message.Type , message.ListenAddr)
}