package proto

import (
	"github.com/ouczbs/Zmin/engine/znet"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"mmo/engine/proto/pb"
	"strings"
)

var (
	pbMessageTypes  = make(map[TCmd]IReflectMessageType)
	CommandListName = pb.CMD_name
)

func PbMessageHandle(proxy *UClientProxy , packet * UPacket, cmd TCmd , id TComponentId){
	wrapBytes := packet.MessagePayload()
	wrapMessage := &UWrapMessage{}
	Unmarshal(wrapBytes , wrapMessage)
	handle , globalHandle := proxy.GetRequestHandles(wrapMessage.Response , cmd)
	if cmd == 0 || (handle == nil && globalHandle == nil){
		return
	}
	pbMessage,_ := newPbMessage(cmd)
	request := GetRequestMessage(wrapMessage ,pbMessage , id)
	proxy.Then(handle , request).Then(globalHandle , request)
}

func GetRequestMessage(wrap *UWrapMessage, message IReflectMessage , id TComponentId) *URequest {
	Unmarshal(wrap.Content, message)
	request := znet.NewRequest(_CMD_INVALID, _MT_INVALID ,message)
	request.Code = wrap.Code
	request.Next = true
	request.Id = id
	request.Request = wrap.Request
	return request
}
func newPbMessage(cmd TCmd) (IReflectMessage , error){
	messageType := pbMessageTypes[cmd]
	if messageType != nil {
		return messageType.New().Interface() ,nil
	}
	pbName := CommandListName[int32(cmd)]
	pbName = strings.Replace(pbName, "x" , "pb." , 1)
	messageType,err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(pbName))
	if err != nil {
		return nil , err
	}
	pbMessageTypes[cmd] = messageType
	return messageType.New().Interface() ,err
}