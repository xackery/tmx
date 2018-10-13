package tmx

import (
	"encoding/xml"
	"io"

	"github.com/pkg/errors"
	"github.com/xackery/tmx/pb"
)

// Decoder decodes TMX files to Protobuf
type Decoder struct {
	r io.Reader
}

// NewDecoder creates a new decoder
func NewDecoder(r io.Reader) (d *Decoder) {
	d = &Decoder{
		r: r,
	}
	return
}

// Decode reads a TMX xml and outputs to a protobuf file
func (d *Decoder) Decode(dst *pb.Map) (err error) {
	xd := xml.NewDecoder(d.r)
	n := Node{}
	err = xd.Decode(&n)
	if err != nil {
		err = errors.Wrap(err, "failed to decode")
		return
	}

	err = walkTMXRoot(dst, []Node{n})
	return
}
