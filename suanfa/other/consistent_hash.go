package other

import (
	"sync"
	"strconv"
	"hash/crc32"
	"sort"
	"fmt"
)

const (
	DefaultReplicas = 160 //节点副本, 与权重有关
)

type HashRing []uint32 //一个hashring包含 2^32 个节点, 索引为0~2^32-1

//返回hashring的长度
func (c HashRing) Len() int {
	return len(c)
}

//比较两个hashring上的value大小
func (c HashRing) Less(i, j int) bool {
	return c[i] < c[j]
}

//交换hashring上下标i,j的元素值
func (c HashRing) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type Node struct {
	Id       int	//id 必须唯一性
	Ip       string
	Port     int
	HostName string
	Weight   int
}

//创建一个节点
func NewNode(id int, ip string, port int, name string, weight int) *Node {
	return &Node{
		Id:       id,
		Ip:       ip,
		Port:     port,
		HostName: name,
		Weight:   weight,
	}
}

type Consistent struct {
	Nodes     map[uint32]Node		//保存全部节点
	numReps   int					//副本数量, 表示一个物理节点, 在consistent上虚拟为n个节点
	Resources map[int]bool			//记录物理节点是否存在的map, key为节点id
	ring      HashRing				//哈希环
	sync.RWMutex					//读写锁
}

func NewConsistent() *Consistent {
	return &Consistent{
		Nodes:     make(map[uint32]Node),
		numReps:   DefaultReplicas,
		Resources: make(map[int]bool),
		ring:      HashRing{},
	}
}

//向consistent添加一个节点
func (c *Consistent) Add(node *Node) bool {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.Resources[node.Id]; ok {	//如果节点id存在, 那么不再添加
		return false
	}

	count := c.numReps * node.Weight
	for i := 0; i < count; i++ {
		str := c.joinStr(i, node)
		c.Nodes[c.hashStr(str)] = *(node)	//每个虚拟节点生成hash值, 加入c.Nodes
	}

	c.Resources[node.Id] = true
	c.sortHashRing()
	return true
}

//根据虚拟节点的hash值排序
func (c *Consistent) sortHashRing() {
	c.ring = HashRing{}
	for k := range c.Nodes {
		c.ring = append(c.ring, k)
	}
	sort.Sort(c.ring)
}

//创建join字符串, 用于生成hash值
func (c *Consistent) joinStr(i int, node *Node) string {
	return node.Ip + "*" + strconv.Itoa(node.Weight) +
		"-" + strconv.Itoa(i) +
		"-" + strconv.Itoa(node.Id)
}

// MurMurHash算法 :https://github.com/spaolacci/murmur3
func (c *Consistent) hashStr(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

//根据key 求hash , 顺时针查找第一个虚拟节点, 那么key应该保存在这个虚拟节点对应的物理节点上
func (c *Consistent) Get(key string) Node {
	c.RLock()
	defer c.RLock()
	hash := c.hashStr(key)
	i := c.search(hash)

	return c.Nodes[c.ring[i]]
}

//根据hash值顺时针找到c.ring上第一个节点
func (c *Consistent) search(hash uint32) int {
	i := sort.Search(len(c.ring), func(i int) bool {
		return c.ring[i] >= hash
	})

	if i < len(c.ring) {
		if i == len(c.ring)-1 {
			return 0
		} else {
			return i
		}
	} else {
		return len(c.ring) - 1	//查找失败, hash值落到c.ring上最后一个节点
	}
}

//删除一个节点
func (c *Consistent) Remove(node *Node) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.Resources[node.Id]; !ok {
		return
	}

	delete(c.Resources, node.Id)
	count := c.numReps * node.Weight
	for i := 0; i < count; i++ {
		str := c.joinStr(i, node)
		delete(c.Nodes, c.hashStr(str))
	}
	c.sortHashRing()
}

func main() {
	cHashRing := NewConsistent()

	for i := 0; i < 10; i++ {
		si := fmt.Sprintf("%d", i)
		cHashRing.Add(NewNode(i, "172.18.1."+si, 8080, "host_"+si, 1))
	}

	for k, v := range cHashRing.Nodes {
		fmt.Println("Hash:", k, " IP:", v.Ip)
	}

	ipMap := make(map[string]int, 0)
	for i := 0; i < 1000; i++ {
		si := fmt.Sprintf("key%d", i)
		k := cHashRing.Get(si)
		if _, ok := ipMap[k.Ip]; ok {
			ipMap[k.Ip] += 1
		} else {
			ipMap[k.Ip] = 1
		}
	}

	for k, v := range ipMap {
		fmt.Println("Node IP:", k, " count:", v)
	}
}
