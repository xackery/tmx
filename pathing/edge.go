package pathing

// Edge represents an edge of an node
type Edge struct {
	Dest   *Node
	Action Direction
	score  float64
}
