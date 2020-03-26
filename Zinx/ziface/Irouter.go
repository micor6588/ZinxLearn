package ziface

/*
路由的抽象方法
路由里面的数据都是Irequest
*/
// IRouter 路由的接口
type IRouter interface {
	PreHandle(request IRquest)  //在处理conn业务之前的钩子方法hook
	Handle(request IRquest)     //在处理hook业务的主方法
	PostHandle(request IRquest) //在处理conn业务知呼的钩子方法hook
}
