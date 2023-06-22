package design_pattern

//单例模式

/*
场景：

• 一些只允许存在一个实例的类，比如全局统一的监控统计模块
• 一些实例化时很耗费资源的类，比如协程池、连接池、和第三方交互的客户端等
• 一些入参繁杂的系统模块组件，比如 controller、service、dao 等
*/

var s *singleton

func init() {
	s = newSingleton()
}

type singleton struct {
}

func (s *singleton) Work() {
}

func newSingleton() *singleton {
	return &singleton{}
}

func GetInstance() *singleton {
	return s
}
