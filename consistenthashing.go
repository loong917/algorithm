package algorithm

import (
	"sort"
	"strconv"
	"sync"

	"github.com/spaolacci/murmur3"
)

const (
	// 最小虚拟节点个数
	MIN_REPLICAS = 10
)

// Func defines the hash method.
type Func func(data []byte) uint64

// A Consistent Hashing is a ring hash implementation.
type ConsistentHashing struct {
	// hash function
	hashFunc Func
	// determine the number of virtual nodes of the node
	replicas int
	// list of virtual nodes
	keys []uint64
	// mapping of virtual nodes to physical nodes
	ring map[uint64][]string
	// physical node map, quickly determine if a node exists
	nodes map[string]struct{}
	// read and write lock
	lock sync.RWMutex
}

// returns a Consistent Hashing with given replicas and hash func.
func NewConsistentHashing(replicas int, fn Func) *ConsistentHashing {
	if replicas < MIN_REPLICAS {
		replicas = MIN_REPLICAS
	}
	if fn == nil {
		fn = hashDefault
	}
	return &ConsistentHashing{
		hashFunc: fn,
		replicas: replicas,
		ring:     make(map[uint64][]string),
		nodes:    make(map[string]struct{}),
	}
}

// adds the node with the number of replicas.
// replicas will be truncated to h.replicas if it's larger than h.replicas.
// the later call will overwrite the replicas of the former calls.
func (h *ConsistentHashing) AddWithReplicas(node string, replicas int) {
	// support repeatable additions, perform the delete operation first
	h.Remove(node)
	// cannot exceed the upper limit of the magnification factor
	if replicas > h.replicas {
		replicas = h.replicas
	}
	h.lock.Lock()
	defer h.lock.Unlock()
	// add node
	h.nodes[node] = struct{}{}
	for i := 0; i < replicas; i++ {
		// create virtual node
		hash := h.hashFunc([]byte(node + strconv.Itoa(i)))
		// note that hashFunc may have hash conflicts
		h.keys = append(h.keys, hash)
		// hash conflicts means that virtual node will correspond to more than one real node
		h.ring[hash] = append(h.ring[hash], node)
	}
	// later, we'll use a binary search lookup of the virtual nodes
	sort.Slice(h.keys, func(i, j int) bool {
		return h.keys[i] < h.keys[j]
	})
}

// removes the given node from h.
func (h *ConsistentHashing) Remove(node string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	// node does not exist
	_, ok := h.nodes[node]
	if !ok {
		return
	}
	// remove the virtual node mapping
	for i := 0; i < h.replicas; i++ {
		hash := h.hashFunc([]byte(node + strconv.Itoa(i)))
		// binary search
		// if index < len(n.keys), FOUND , otherwise NOT FOUND
		index := sort.Search(len(h.keys), func(i int) bool {
			return h.keys[i] >= hash
		})
		// locate the element before the index of the slice
		// move the element after index forward to cover index
		if index < len(h.keys) && h.keys[index] == hash {
			h.keys = append(h.keys[:index], h.keys[index+1:]...)
		}
		// virtual node removal mapping
		h.removeRingNode(hash, node)
	}
	// remove the real node
	delete(h.nodes, node)
}

func (h *ConsistentHashing) removeRingNode(hash uint64, node string) {
	// map should be checked when used
	if nodes, ok := h.ring[hash]; ok {
		// create a new empty slice, with the same capacity as nodes
		newNodes := nodes[:0]
		for _, x := range nodes {
			if x != node {
				newNodes = append(newNodes, x)
			}
		}
		if len(newNodes) > 0 {
			// rebind the mapping relationship if the remaining nodes are not empty
			h.ring[hash] = newNodes
		} else {
			// otherwise just delete
			delete(h.ring, hash)
		}
	}
}

// returns the corresponding node from h base on the given value.
func (h *ConsistentHashing) Get(value string) (interface{}, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if len(h.ring) == 0 {
		return nil, false
	}
	hash := h.hashFunc([]byte(value))
	// the first node queried is our target node
	index := sort.Search(len(h.keys), func(i int) bool {
		return h.keys[i] >= hash
	})
	// mod operation give us a circular list effect, finding nodes clockwise
	hashRing := h.keys[index%len(h.keys)]
	// virtual nodes -> physical nodes mapping
	nodes := h.ring[hashRing]
	switch len(nodes) {
	case 0:
		return nil, false
	case 1:
		return nodes[0], true
	default:
		// at this point we re-hash v
		// we get a new index by taking the remainder of the nodes length
		newIndex := hashDefault([]byte(value))
		pos := int(newIndex % uint64(len(nodes)))
		return nodes[pos], true
	}
}

func hashDefault(data []byte) uint64 {
	return murmur3.Sum64(data)
}
