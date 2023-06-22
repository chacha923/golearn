package design_pattern

import "context"

/* 在观察者模式中，核心的角色包含三类：
• Observer：观察者. 指的是关注事物动态的角色
• Event：事物的变更事件. 其中 Topic 标识了事物的身份以及变更的类型，Val 是变更详情
• EventBus：事件总线. 位于观察者与事物之间承上启下的代理层. 负责维护管理观察者，并且在事物发生变更时，将情况同步给每个观察者.
*/

type Event struct {
	Topic string
	Val   interface{}
}

type Observer interface {
	OnChange(ctx context.Context, e *Event) error
}

type EventBus interface {
	Subscribe(topic string, o Observer)
	UnSubscribe(topic string, o Observer)
	Publish(ctx context.Context, e *Event)
}

/* 生产案例：
1. 发布订阅
• EventBus：对应的是消息队列组件，为整个通信架构提供了分布式解耦、流量削峰等能力

• Event：对应的是消息队列中的一条消息，有明确的主题 topic，由生产者 producer 提供

• Observer：对应的是消费者 consumer，对指定事物的动态（topic）进行订阅，并在消费到对应的变更事件后执行对应的处理逻辑


2. ETCD 监听回调
• EventBus：对应的是 etcd 服务端的 watchableStore 监听器存储模块，该模块会负责存储用户创建的一系列监听器 watcher，并建立由监听数据 key 到监听器集合 watcherGroup 之间的映射关系. 当任意存储数据发生变化时，etcd 的数据存储模块会在一个统一的切面中调用通知方法，将这一信息传达到 watchableStore 模块，watchableStore 则会将变更数据与监听数据 key 之间进行 join，最终得到一个需要执行回调操作的 watchers 组合，顺沿 watcher 中的路径，向订阅者发送通知消息

• Event：对应的是一条 etcd 状态机的数据变更事件，由 etcd 使用方在执行一条写数据操作时触发，在写操作真正生效后，变更事件会被传送到 watchableStore 模块执行回调处理

• Observer：使用 etcd watch 功能对指定范围数据建立监听回调机制的使用方，在 etcd 服务端 watchableStore 模块会建立监听器实体 watcher 作为自身的代理，当变更事件真的发生后，watchableStore 会以 watcher 作为起点，沿着返回路径一路将变更事件发送到使用方手中.

*/
