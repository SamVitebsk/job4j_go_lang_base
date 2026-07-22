package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	size  int
	Head  *Node
	Tail  *Node
	items map[string]*Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		size:  size,
		items: make(map[string]*Node),
	}
}

func (l *LruCache) Put(key string, value string) {
	if l.size <= 0 {
		return
	}

	if node, ok := l.items[key]; ok {
		node.Value = value
		l.moveToHead(node)
		return
	}

	node := &Node{
		Key:   key,
		Value: value,
	}

	l.items[key] = node
	l.addToHead(node)

	if len(l.items) > l.size {
		l.removeTail()
	}
}

func (l *LruCache) Get(key string) *string {
	node, ok := l.items[key]
	if !ok {
		return nil
	}

	l.moveToHead(node)

	return &node.Value
}

func (l *LruCache) moveToHead(node *Node) {
	if l.Head == node {
		return
	}

	l.removeNode(node)
	l.addToHead(node)
}

func (l *LruCache) removeNode(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.Tail = node.Prev
	}

	node.Next = nil
	node.Prev = nil
}

func (l *LruCache) addToHead(node *Node) {
	node.Prev = nil
	node.Next = l.Head

	if l.Head != nil {
		l.Head.Prev = node
	}

	l.Head = node

	if l.Tail == nil {
		l.Tail = node
	}
}

func (l *LruCache) removeTail() {
	if l.Tail == nil {
		return
	}

	node := l.Tail
	l.removeNode(node)
	delete(l.items, node.Key)
}
