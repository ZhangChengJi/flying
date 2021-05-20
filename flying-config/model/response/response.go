package response

import (
	"flying-config/constant"
	proto "flying-config/proto/common"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 500
	SUCCESS = 200
	UPDATES = 304
)

func Result(code constant.PolicyType, data protobuf.Message, msg string) *proto.Response {
	any, _ := ptypes.MarshalAny(data)
	rsp := &proto.Response{
		Code: int32(code),
		Msg:  msg,
		Data: any,
	}
	return rsp
}

func Results(code int32, data []*anypb.Any, msg string) *proto.Responses {
	//any, _ := ptypes.MarshalAny(data)
	rsp := &proto.Responses{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	return rsp
}

func BuildResult(code constant.PolicyType) *proto.Response {
	return Result(code, nil, "")
}

func Ok() *proto.Response {
	return Result(SUCCESS, nil, "")
}

func OkWithMessage(message string) *proto.Response {
	return Result(SUCCESS, nil, message)
}

func OkWithData(data protobuf.Message) *proto.Response {
	return Result(SUCCESS, data, "操作成功")
}
func OkWithDatas(data []*anypb.Any) *proto.Responses {
	return Results(SUCCESS, data, "操作成功")
}
func OkWithDetailed(data protobuf.Message, message string) *proto.Response {
	return Result(SUCCESS, data, message)
}

func Fail() *proto.Response {
	return Result(ERROR, nil, "")
}

func FailWithMessage(message string) *proto.Response {
	return Result(ERROR, nil, message)
}

func FailWithDetailed(data protobuf.Message, message string) *proto.Response {
	return Result(ERROR, data, message)
}
func NotUpdated(message string) *proto.Response {
	return Result(UPDATES, nil, message)
}
