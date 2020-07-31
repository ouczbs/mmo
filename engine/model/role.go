package model

import "mmo/engine/conf"

type URole struct {
	Id conf.TComponentId
	ServerId conf.TComponentId
	Name string

	AvatarPath string
}

func (account * URole)Table()string{
	return "role"
}

func (account * URole) M()map[string]interface{}{
	return M(account)
}
var Role = &URole{}
//一个账号一个服可以创建一个角色，所以一个账号对应多个角色
func init(){
	Schema(Role)
}