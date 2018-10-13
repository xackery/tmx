package tmx

import (
	"bytes"
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/xackery/tmx/pb"
)

// Encoder encodes a protobuf file
type Encoder struct {
	w io.Writer
	//UseJSON enables output to json
	UseJSON bool
}

// NewEncoder returns a new encoder that writes to w
func NewEncoder(w io.Writer) (e *Encoder) {
	e = &Encoder{
		w: w,
	}
	return
}

// Encode encodes protobuf to writer
func (e *Encoder) Encode(dst *pb.Map) (err error) {
	data, err := proto.Marshal(dst)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal protobuf")
		return
	}
	buf := bytes.NewBuffer(data)

	_, err = buf.WriteTo(e.w)
	if err != nil {
		err = errors.Wrap(err, "failed to write")
		return
	}
	return
}
