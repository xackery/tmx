package pathing

import (
	"container/heap"
	"fmt"
)

// Pathing represents a navigation mesh
type Pathing struct {
	nodes map[int8]map[int8]*Node
}

// NewPathing returns a new pathing system
func NewPathing() (p *Pathing) {
	p = &Pathing{
		nodes: make(map[int8]map[int8]*Node),
	}
	return
}

// NodeRead returns a node at a coordinate
func (p *Pathing) NodeRead(x int8, y int8) *Node {
	nY, ok := p.nodes[x]
	if !ok {
		return nil
	}
	return nY[y]
}

// NodeCreate adds a new node
func (p *Pathing) NodeCreate(x int8, y int8, isCollider bool) {
	n := p.NodeRead(x, y)
	if n != nil {
		n.IsCollider = isCollider
		return
	}
	_, ok := p.nodes[x]
	if !ok {
		p.nodes[x] = make(map[int8]*Node)
	}
	p.nodes[x][y] = &Node{X: x, Y: y, IsCollider: isCollider}
	return
}

// Path using astar
func (p *Pathing) Path(startX int8, startY int8, goalX int8, goalY int8) (path []*Edge, err error) {
	var ok bool
	seen := make(map[int8]map[int8]bool)
	seen[startX] = make(map[int8]bool)

	openHeap := make(PriorityQueue, 0)
	heap.Init(&openHeap)
	cameFrom := make(map[int8]map[int8]*Edge)

	gScore := make(map[int8]map[int8]float64)
	gScore[startX] = make(map[int8]float64)
	fScore := make(map[int8]map[int8]float64)
	fScore[startX] = make(map[int8]float64)

	node := p.NodeRead(startX, startY)
	if node == nil {
		err = fmt.Errorf("invalid starting point")
		return
	}
	gScore[startX][startY] = 0
	fScore[startX][startY] = gScore[startX][startY] + node.Heuristic(goalX, goalY)
	heap.Push(&openHeap, &Item{node: node, priority: 0})

	seen[startX][startY] = true
	var item interface{}
	for {
		item = heap.Pop(&openHeap)
		if item == nil {
			err = fmt.Errorf("out of items")
			return
		}
		node = item.(*Item).node
		if node == nil {
			err = fmt.Errorf("out of items")
			return
		}
		fmt.Println("testing", node)

		if node.IsCollider {
			continue
		}
		if node.Success(goalX, goalY) {
			fmt.Println("found")
			return reconstructPath(cameFrom, node), nil
		}
		for _, edge := range p.Neighbors(node.X, node.Y) {
			if edge == nil {
				continue
			}
			if edge.Dest == nil {
				continue
			}
			adj := edge.Dest
			action := edge.Action
			_, ok = seen[adj.X]
			if ok && seen[adj.X][adj.Y] {
				continue
			}
			if !ok {
				seen[adj.X] = make(map[int8]bool)
			}
			seen[adj.X][adj.Y] = true

			_, ok = gScore[adj.X]
			if !ok {
				gScore[adj.X] = make(map[int8]float64)
			}
			_, ok = fScore[adj.X]
			if !ok {
				fScore[adj.X] = make(map[int8]float64)
			}
			_, ok = cameFrom[adj.X]
			if !ok {
				cameFrom[adj.X] = make(map[int8]*Edge)
			}
			// adjacency cost is based on a constant step
			gScore[adj.X][adj.Y] = gScore[node.X][node.Y] + 1
			hScore := adj.Heuristic(goalX, goalY)
			fScore[adj.X][adj.Y] = gScore[adj.X][adj.Y] + hScore
			fmt.Println("Push", adj)
			heap.Push(&openHeap, &Item{node: adj, priority: fScore[adj.X][adj.Y]})
			// reverse the edge for reconstruction
			cameFrom[adj.X][adj.Y] = &Edge{
				Dest:   node,
				Action: action,
				score:  fScore[adj.X][adj.Y],
			}
		}
	}
}

// Neighbors returns edge neighbors of a position
func (p *Pathing) Neighbors(nodeX int8, nodeY int8) (neighbors []*Edge) {
	neighbors = []*Edge{
		&Edge{Dest: p.NodeRead(nodeX+1, nodeY), Action: Right},
		&Edge{Dest: p.NodeRead(nodeX-1, nodeY), Action: Left},
		&Edge{Dest: p.NodeRead(nodeX, nodeY+1), Action: Up},
		&Edge{Dest: p.NodeRead(nodeX, nodeY-1), Action: Down},
	}
	return
}

/*
// AStar implements an astar pathfinding
func AStar(start Node, goal Node) []Edge {
	seen := make(map[Node]bool)
	openHeap := make(PriorityQueue, 0)
	heap.Init(&openHeap)
	cameFrom := make(map[Node]Edge)
	gScore := make(map[Node]float64)
	fScore := make(map[Node]float64)
	gScore[start] = 0
	fScore[start] = gScore[start] + start.Heuristic(goal)
	heap.Push(&openHeap, &Item{node: start, priority: fScore[start]})
	seen[start] = true
	for {
		node := heap.Pop(&openHeap).(*Item).node
		if node.Success(goal) {
			return reconstructPath(cameFrom, node)
		}
		for _, edge := range node.Neighbors() {
			adj := edge.Dest
			action := edge.Action
			if seen[adj] {
				continue
			}
			seen[adj] = true
			// adjacency cost is based on a constant step
			gScore[adj] = gScore[node] + 1
			hScore := adj.Heuristic(goal)
			fScore[adj] = gScore[adj] + hScore
			heap.Push(&openHeap, &Item{node: adj, priority: fScore[adj]})
			// reverse the edge for reconstruction
			cameFrom[adj] = Edge{
				Dest:   node,
				Action: action,
				score:  fScore[adj],
			}
		}
	}
}
*/

func reconstructPath(cameFrom map[int8]map[int8]*Edge, node *Node) []*Edge {
	if edge, ok := cameFrom[node.X][node.Y]; ok {
		return append(reconstructPath(cameFrom, edge.Dest), edge)
	}
	return make([]*Edge, 0)
}
