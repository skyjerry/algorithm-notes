// Leetcode 146. LRU 缓存机制

type Node struct {
    key, val int
    prev, next *Node
}

type LRUCache struct {
    hashMap map[int]*Node
    head, tail *Node
    cap int
}

func initNode(key, val int) *Node {
    return &Node{
        key: key,
        val: val,
    }
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
        hashMap: make(map[int]*Node),
        cap: capacity,
        head: initNode(0, 0),
        tail: initNode(0, 0),
    }

    lru.head.next = lru.tail
    lru.tail.prev = lru.head

    return lru
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.hashMap[key]
    if !ok {
        return -1
    }

    this.moveToHead(node)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    node, ok := this.hashMap[key]
    if ok {
        if node.val != value {
            node.val = value
        }
        this.moveToHead(node)
        return
    }

    newHead := initNode(key, value)

    this.hashMap[key] = newHead
    this.addHead(newHead)

    if len(this.hashMap) > this.cap {
        removeNode := this.removeTail()
        delete(this.hashMap, removeNode.key)
    }
}

func (this *LRUCache) removeNode(node *Node) *Node {
    node.next.prev = node.prev
    node.prev.next = node.next

    return node
}

func (this *LRUCache) removeTail() *Node {
    node := this.tail.prev
    this.removeNode(node)
    return node
}

func (this *LRUCache) moveToHead(node *Node) {
    this.removeNode(node)
    this.addHead(node)
}

func (this *LRUCache) addHead(node *Node) {
    node.next = this.head.next
    node.prev = this.head
    this.head.next.prev = node
    this.head.next = node
}