package user

type UPlayerPool struct {
	* UStackPool
}
func NewPlayerPool(size TSize) *UPlayerPool {
	stack := &UPlayerPool{
		&UStackPool{
			Size: size,
		},
	}
	stack.Init()
	return stack
}
func (stack * UPlayerPool) Pop() * UPlayer{
	object := stack.UStackPool.Pop()
	if object == nil {
		return stack.New()
	}
	return object.(*UPlayer)
}
func (stack * UPlayerPool) New() * UPlayer{
	return &UPlayer{}
}