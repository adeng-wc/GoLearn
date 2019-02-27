package mock

// 可以不交 Receiver ，只有 Get 方法就可以
type Receiver struct {
	Contexts string
}


func (r Receiver) Get(url string) string {
	return r.Contexts
}
