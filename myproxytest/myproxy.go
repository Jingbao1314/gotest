package myproxytest

type MyInvokeHandler interface {
	Invoke()
}

type MyProxy interface {
	NewProxy()
}

type ProxyHandler struct {
	MyProxy
}

func (*ProxyHandler) NewProxy(a interface{}, b uintptr) {

}
