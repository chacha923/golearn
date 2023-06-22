package design_pattern

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
工厂模式：
• 实现类和业务方法之间的解耦，如果类的构造过程发生变更，可以统一收口在工厂类中进行处理，从而对业务方法屏蔽相关细节
• 倘若有多个类都聚拢在工厂类中进行构造，这样各个类的构造流程中就天然形成了一个公共的切面，可以进行一些公共逻辑的执行
*/

// 简单工厂模式
type Fruit interface {
	Eat()
}

type Orange struct {
	name string
}

func NewOrange(name string) Fruit {
	return &Orange{
		name: name,
	}
}

func (o *Orange) Eat() {
	fmt.Printf("i am orange: %s, i am about to be eaten...", o.name)
}

type Strawberry struct {
	name string
}

func NewStrawberry(name string) Fruit {
	return &Strawberry{
		name: name,
	}
}

func (s *Strawberry) Eat() {
	fmt.Printf("i am strawberry: %s, i am about to be eaten...", s.name)
}

type Cherry struct {
	name string
}

func NewCherry(name string) Fruit {
	return &Cherry{
		name: name,
	}
}

func (c *Cherry) Eat() {
	fmt.Printf("i am cherry: %s, i am about to be eaten...", c.name)
}

type FruitFactory struct {
}

func NewFruitFactory() *FruitFactory {
	return &FruitFactory{}
}

func (f *FruitFactory) CreateFruit(typ string) (Fruit, error) {
	src := rand.NewSource(time.Now().UnixNano())
	rander := rand.New(src)
	name := strconv.Itoa(rander.Int())

	switch typ {
	case "orange":
		return NewOrange(name), nil
	case "strawberry":
		return NewStrawberry(name), nil
	case "cherry":
		return NewCherry(name), nil
	default:
		return nil, fmt.Errorf("fruit typ: %s is not supported yet", typ)
	}
}

/*
工厂模式可以进一步细分为:

• 简单工厂模式: 工厂模式中最简单直观的实现方式,有很好的切面效果,但是在组件类扩展时无法满足开闭原则

• 工厂方法模式: 一个组件类对应一个工厂类, 存在一定的代码冗余以及对公共切面的削弱，但是能够在组件类扩展时满足开闭原则

• 抽象工厂模式: 通过两个维度对组件类进行拆解. 需要保证易于扩展、灵活可变的维度需要定义为产品族；相对稳定、不易于扩展维度需要定义为产品等级. 这样能同时保证产品族维度的扩展灵活性以及产品等级维度的切面能力.

此外，本文还额外介绍了一种另类的容器工厂模式，底层需要基于依赖注入框架实现，让组件提供能够在各处方便地完成组件类的注入操作，而组件的使用方，则通过容器工厂的统一出口进行组件的获取.
*/
