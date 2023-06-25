package concurrent

import (
	"fmt"
	"sync"
)

type ConcurrentMapShared[K comparable, V any] struct {
	items map[K]V
	// Read Write mutex, guards access to internal map.
	sync.RWMutex
}

// A "thread" safe map of type string:Anything.
// To avoid lock bottlenecks this map is dived to several (SHARD_COUNT) map shards.
type ConcurrentMap[K comparable, V any] struct {
	shards   []*ConcurrentMapShared[K, V]
	sharding func(key K) uint32 // 计算 key 分片的函数
}

// 用于打印
type Stringer interface {
	fmt.Stringer
	comparable
}

var SHARD_COUNT = 32

func create[K comparable, V any](sharding func(key K) uint32) ConcurrentMap[K, V] {
	var m = ConcurrentMap[K, V]{
		sharding: sharding,
		shards:   make([]*ConcurrentMapShared[K, V], SHARD_COUNT),
	}
	// 初始化每个 shard
	for i := 0; i < SHARD_COUNT; i++ {
		m.shards[i] = &ConcurrentMapShared[K, V]{items: make(map[K]V)}
	}
	return m
}

// Creates a new concurrent map.
func New[V any]() ConcurrentMap[string, V] {
	return create[string, V](fnv32)
}

// Creates a new concurrent map.
func NewStringer[K Stringer, V any]() ConcurrentMap[K, V] {
	return create[K, V](strfnv32[K])
}

// Creates a new concurrent map.
func NewWithCustomShardingFunction[K comparable, V any](sharding func(key K) uint32) ConcurrentMap[K, V] {
	return create[K, V](sharding)
}

func (m ConcurrentMap[K, V]) GetShard(key K) *ConcurrentMapShared[K, V] {
	return m.shards[m.sharding(key)%uint32(SHARD_COUNT)]
}

// Sets the given value under the specified key.
func (m ConcurrentMap[K, V]) Set(key K, value V) {
	var shard = m.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

// Sets the given value under the specified key if no value was associated with it.
func (m ConcurrentMap[K, V]) SetIfAbsent(key K, value V) bool {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	_, ok := shard.items[key]
	if !ok {
		shard.items[key] = value
	}
	shard.Unlock()
	return !ok
}

func (m ConcurrentMap[K, V]) Get(key K) (value V, ok bool) {
	shard := m.GetShard(key)
	shard.RLock()
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// Looks up an item under specified key
func (m ConcurrentMap[K, V]) Has(key K) bool {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// See if element is within shard.
	_, ok := shard.items[key]
	shard.RUnlock()
	return ok
}

// Count returns the number of elements within the map.
func (m ConcurrentMap[K, V]) Count() int {
	count := 0
	for i := 0; i < SHARD_COUNT; i++ {
		shard := m.shards[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

// Remove removes an element from the map.
func (m ConcurrentMap[K, V]) Remove(key K) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

func strfnv32[K fmt.Stringer](key K) uint32 {
	return fnv32(key.String())
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
