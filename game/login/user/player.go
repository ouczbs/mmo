package user

type UPlayer struct{
	Id      TComponentId
	ServerId TComponentId
	IsLogin bool

	isReleased bool
}

var (
	playerPool = NewPlayerPool(16)
)

func NewPlayer()*UPlayer{
	player := playerPool.Pop()
	player.isReleased = false
	return player
}
func (player * UPlayer) Release()  {
	if player.isReleased {
		return
	}
	player.isReleased = true
	playerPool.Push(player)
}