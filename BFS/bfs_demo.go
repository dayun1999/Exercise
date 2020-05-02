package BFS

//BFS(Breadth first search) is a traversing algorithm where you
//should start traversing from a selected node (source or starting node)
//and traverse the tree layerwise thus exploring the neighbor
//nodes (nodes which are directly connected to source node).
//You must then move towards the next-level neighbor nodes.
type Tree struct {
	Value  int
	Left   *Tree
	Right  *Tree
	Parent *Tree
	Repeat int
}

func BFS(tree *Tree) []int {
	queue := []*Tree{}
	queue = append(queue, tree)
	result := []int{}
	return BFSUtil(queue, result)
}

func BFSUtil(queue []*Tree, res []int) []int {
	//this appends the value of the root of subtree or tree to the result
	if len(queue) == 0 {
		return res
	}
	res = append(res, queue[0].Value)
	if queue[0].Right != nil {
		queue = append(queue, queue[0].Right)
	}
	if queue[0].Left != nil {
		queue = append(queue, queue[0].Left)
	}
	return BFSUtil(queue[1:], res)
}
