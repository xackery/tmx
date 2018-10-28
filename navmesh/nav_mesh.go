package navmesh

// NavMesh represents a navigation mesh
type NavMesh struct {
}

// New creates a new navmesh
func New() (n *NavMesh, err error) {
	n = &NavMesh{}
	return
}
