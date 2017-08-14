package actor

import (
	"bytes"
	"encoding/gob"
	"io"
)

type binarySerializer struct {
	buf bytes.Buffer
	enc *gob.Encoder
	dec *gob.Decoder
}

func (self *binarySerializer) IsLoading() bool {
	return self.dec != nil
}

func (self *binarySerializer) Serialize(data interface{}) {

	if self.enc != nil {
		self.enc.Encode(data)
	} else if self.dec != nil {
		self.dec.Decode(data)
	}

}

func (self *binarySerializer) Bytes() []byte {
	return self.buf.Bytes()
}

func NewBinaryWriter() Serializer {

	self := &binarySerializer{}

	self.enc = gob.NewEncoder(&self.buf)

	return self
}

func NewBinaryReader(r io.Reader) Serializer {

	self := &binarySerializer{
		dec: gob.NewDecoder(r),
	}

	return self
}
