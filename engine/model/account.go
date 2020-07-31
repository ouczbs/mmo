package model

type UAccount struct {
	Account string
	Password string

	RoleList string // role id list
}

func (account * UAccount)Table()string{
	return "account"
}

func (account * UAccount) M()map[string]interface{}{
	return M(account)
}
var Account = &UAccount{}
//一个账号一个服可以创建一个角色，所以一个账号对应多个角色
func init(){
	Schema(Account)
}
