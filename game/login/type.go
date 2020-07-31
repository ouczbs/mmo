package login

import (
	"github.com/ouczbs/Zmin/engine/znet"
	"mmo/engine/conf"
	"mmo/engine/model"
	"mmo/game/base"
	"mmo/game/login/user"
)

type (
	UClientProxy = znet.UClientProxy
	UMessage = znet.UMessage
	URequest = znet.URequest
	UService = base.UService
	FRequestHandle = znet.FRequestHandle

	TCmd = conf.TCmd
	TMessageType = conf.TMessageType
	TComponentId = conf.TComponentId

	UPlayer = user.UPlayer
	UAccount = model.UAccount
)

var (
	reqHandleMaps  = make(map[TCmd]FRequestHandle)
	centerProxy * UClientProxy

	dispatcherProxyMaps = make(map[TComponentId]*UClientProxy)

	clientPlayerMaps = make(map[TComponentId]*UPlayer)

	// -- id -> 找到对应角色？？？
)