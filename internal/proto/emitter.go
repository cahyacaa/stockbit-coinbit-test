package proto_codec

import (
	deposit "github.com/cahyacaa/stockbit-coinbit-test/internal/proto_models"
	"github.com/golang/protobuf/proto"
)

type ProtoCodec struct{}

func (c *ProtoCodec) Encode(value interface{}) ([]byte, error) {
	var msg []byte
	var err error

	if byteMsg, ok := value.(deposit.Deposit); ok {
		msg, err = proto.Marshal(&byteMsg)
		if err != nil {
			return msg, nil
		}
	}

	return msg, err
}

func (c *ProtoCodec) Decode(data []byte) (interface{}, error) {
	var w deposit.Deposit

	err := proto.Unmarshal(data, &w)
	if err != nil {
		return []byte{}, err
	}

	return w, nil
}
