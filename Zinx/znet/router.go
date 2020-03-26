package znet

import "ZinxLearn/Zinx/ziface"

// BaseRouter 实现router时，先嵌入BaseRouter基类，然后对这个基类方法进行重写
type BaseRouter struct {
}

//这里之所以BaseRouter的方法都为空
//有的路由不希望实现Preehandle和Handle或者PostHandle方法
//所以Router全部继承BaseRouter的方法好处是，不需要实现PreHandle
// PreHandle 在处理conn业务之前的钩子方法hook
func (br *BaseRouter) PreHandle(request ziface.IRquest) {

}

// Handle 在处理hook业务的主方法
func (br *BaseRouter) Handle(request ziface.IRquest) {

}

// PostHandle 在处理conn业务知呼的钩子方法hook
func (br *BaseRouter) PostHandle(request ziface.IRquest) {

}
