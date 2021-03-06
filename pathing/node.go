package pathing

import (
	"fmt"
	"math"
)

// Node represents an element in the grid
type Node struct {
	X, Y       int8
	IsCollider bool
}

func (n *Node) String() string {
	return fmt.Sprintf("*Node(X: %d, Y: %d)", n.X, n.Y)
}

// Heuristic is a euclidean norm
func (n *Node) Heuristic(goalX, goalY int8) float64 {
	return math.Hypot(float64(goalX-n.X), float64(goalY-n.Y))
}

// Success returns true when we meet a goal
func (n *Node) Success(goalX, goalY int8) bool {
	return n.X == goalX && n.Y == goalY
}

// Copy creates a copy of node
func (n *Node) Copy() *Node {
	nn := &Node{
		X:          n.X,
		Y:          n.Y,
		IsCollider: n.IsCollider,
	}
	return nn
}
