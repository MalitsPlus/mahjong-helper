package api

import (
	"github.com/golang/protobuf/proto"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/proto/lq"
)

func WrapMessage(name string, message proto.Message) (data []byte, err error) {
	data, err = proto.Marshal(message)
	if err != nil {
		return
	}
	wrap := lq.Wrapper{
		Name: name,
		Data: data,
	}
	return proto.Marshal(&wrap)
}

func UnwrapData(rawData []byte) (methodName string, data []byte, err error) {
	wrapper := lq.Wrapper{}
	if err = proto.Unmarshal(rawData, &wrapper); err != nil {
		return
	}
	return wrapper.GetName(), wrapper.GetData(), nil
}

func UnwrapMessage(rawData []byte, message proto.Message) error {
	methodName, data, err := UnwrapData(rawData)
	if err != nil {
		return err
	}
	if methodName != "" {
		// TODO: assert equal
		// methodName ==? proto.MessageName(message)
	}
	return proto.Unmarshal(data, message)
}
