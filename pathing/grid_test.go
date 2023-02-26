package pathing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	assert := assert.New(t)
	path := NewPathing()
	path.NodeCreate(0, 0, false)
	path.NodeCreate(1, 0, true)
	path.NodeCreate(0, 1, false)
	path.NodeCreate(1, 1, false)
	path.NodeCreate(2, 1, false)
	path.NodeCreate(2, 2, false)
	route, err := path.Path(0, 0, 1, 1)
	assert.NoError(err)
	fmt.Printf("path length %d\n", len(route))
	for _, r := range route {
		fmt.Printf("%s->", r.Action)
	}

	//assert.Empty(route)
	assert.NotEmpty(route)

}
