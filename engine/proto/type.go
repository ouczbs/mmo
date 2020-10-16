package proto

import (
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/znet"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"mmo/engine/conf"
	"mmo/engine/proto/pb"
)

type (
	TCode = conf.TCode
	TCmd = conf.TCmd
	TEnum = conf.TEnum
	TMessageType = conf.TMessageType
	TComponentId = conf.TComponentId
	IReflectMessage = protoreflect.ProtoMessage
	IReflectMessageType = protoreflect.MessageType
	IPbMessage = protoreflect.ProtoMessage

	UPacket = znet.UPacket
	URequest = znet.URequest
	UClientProxy = znet.UClientProxy

	UWrapMessage = pb.WrapMessage

	FRequestHandle = znet.FRequestHandle
)

var (
	Unmarshal = proto.Unmarshal
	Marshal = proto.Marshal
)

const (
	_MT_INVALID = zconf.MT_INVALID
	_CMD_INVALID = TCmd(_MT_INVALID)
)