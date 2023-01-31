package main

func (t *Tree) Add(value int) {
    t.root = addNode(t.root, value)
}

func addNode(n *Node, value int) *Node {
    if n == nil {
        n = new(Node)
        n.value = value
        return n
    }
    
    if value < n.value {
        n.left = addNode(n.left, value)
    } else {
        n.right = addNode(n.right, value)
    }
    return n
}

/* Testing Code */
func main() {
    t := new(Tree)
    t.Add(6)
    t.Add(4)
    t.Add(2)
    t.Add(5)
    t.Add(1)
    t.PrintInOrder()
    t.Add(3)
    t.Add(8)
    t.Add(7)
    t.Add(9)
    t.Add(10)
    t.PrintInOrder()
}