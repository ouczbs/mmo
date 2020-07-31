package login

import (
	"github.com/ouczbs/Zmin/engine/zcache"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"go.mongodb.org/mongo-driver/bson"
	"mmo/engine/proto"
	"mmo/engine/proto/pb"
	"strconv"
	"strings"
)

func (service *ULoginService) LoginAccountAck(proxy *UClientProxy , request *URequest, roleList string , result string , success bool ) {
	message := &pb.LOGIN_ACCOUNT_ACK{Result: result,Success: success}
	for _,id := range strings.Split(roleList , ","){
		rid,_ := strconv.Atoi(id)
		message.RoleIdList = append(message.RoleIdList, int32(rid))
	}
	request.Cmd = zconf.TCmd(pb.CMD_MT_LOGIN_ACCOUNT_ACK)
	request.MessageType =  zconf.MT_TO_CLIENT
	request.ProtoMessage = message
	proto.ResponseMessage(proxy , request)
	request.Release()
}
func (service *ULoginService) LoginAccount(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*pb.LOGIN_ACCOUNT)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	account := &UAccount{}
	if err :=zcache.GetMongoClient().FindOne(account,bson.M{"account":message.Account}) ;err != nil {
		zlog.Debug("Account not exits")
		account.Account = message.Account
		account.Password = message.Password
		zcache.GetMongoClient().InsertOne(account)
	}
	if account.Password != message.Password {
		zlog.Debug("Password not right")
		service.LoginAccountAck(proxy ,request , "" , "login failed" , false )
		return
	}
	service.LoginAccountAck(proxy , request , account.RoleList , "" , true )
	zlog.Debug(account)

}
func (service *ULoginService) CreateRole(proxy *UClientProxy, request *URequest) {

}
func (service *ULoginService) InitGameDownHandles(){
	reqHandleMaps[TCmd(pb.CMD_MT_LOGIN_ACCOUNT)] = service.AddEngineComponent
}