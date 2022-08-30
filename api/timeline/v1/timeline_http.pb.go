// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTimelineGetGroupHistoryMessage = "/timeline.v1.Timeline/GetGroupHistoryMessage"
const OperationTimelineGetSingleHistoryMessage = "/timeline.v1.Timeline/GetSingleHistoryMessage"
const OperationTimelineGetSyncMessage = "/timeline.v1.Timeline/GetSyncMessage"
const OperationTimelineSend = "/timeline.v1.Timeline/Send"
const OperationTimelineSendGroup = "/timeline.v1.Timeline/SendGroup"

type TimelineHTTPServer interface {
	GetGroupHistoryMessage(context.Context, *GetGroupHistoryMessageRequest) (*GetGroupHistoryMessageReply, error)
	GetSingleHistoryMessage(context.Context, *GetSingleHistoryMessageRequest) (*GetSingleHistoryMessageReply, error)
	GetSyncMessage(context.Context, *SyncMessageRequest) (*SyncMessageReply, error)
	Send(context.Context, *SendMessageRequest) (*SendMessageReply, error)
	SendGroup(context.Context, *SendGroupRequest) (*SendGroupReply, error)
}

func RegisterTimelineHTTPServer(s *http.Server, srv TimelineHTTPServer) {
	r := s.Route("/")
	r.POST("/timeline/send", _Timeline_Send0_HTTP_Handler(srv))
	r.POST("/timeline/sendGroup", _Timeline_SendGroup0_HTTP_Handler(srv))
	r.GET("/timeline/sync", _Timeline_GetSyncMessage0_HTTP_Handler(srv))
	r.GET("/timeline/history/single/{from}/{to}", _Timeline_GetSingleHistoryMessage0_HTTP_Handler(srv))
	r.GET("/timeline/history/group/{group}", _Timeline_GetGroupHistoryMessage0_HTTP_Handler(srv))
}

func _Timeline_Send0_HTTP_Handler(srv TimelineHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendMessageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTimelineSend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Send(ctx, req.(*SendMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendMessageReply)
		return ctx.Result(200, reply)
	}
}

func _Timeline_SendGroup0_HTTP_Handler(srv TimelineHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTimelineSendGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendGroup(ctx, req.(*SendGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Timeline_GetSyncMessage0_HTTP_Handler(srv TimelineHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SyncMessageRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTimelineGetSyncMessage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSyncMessage(ctx, req.(*SyncMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SyncMessageReply)
		return ctx.Result(200, reply)
	}
}

func _Timeline_GetSingleHistoryMessage0_HTTP_Handler(srv TimelineHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSingleHistoryMessageRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTimelineGetSingleHistoryMessage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSingleHistoryMessage(ctx, req.(*GetSingleHistoryMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSingleHistoryMessageReply)
		return ctx.Result(200, reply)
	}
}

func _Timeline_GetGroupHistoryMessage0_HTTP_Handler(srv TimelineHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGroupHistoryMessageRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTimelineGetGroupHistoryMessage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGroupHistoryMessage(ctx, req.(*GetGroupHistoryMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGroupHistoryMessageReply)
		return ctx.Result(200, reply)
	}
}

type TimelineHTTPClient interface {
	GetGroupHistoryMessage(ctx context.Context, req *GetGroupHistoryMessageRequest, opts ...http.CallOption) (rsp *GetGroupHistoryMessageReply, err error)
	GetSingleHistoryMessage(ctx context.Context, req *GetSingleHistoryMessageRequest, opts ...http.CallOption) (rsp *GetSingleHistoryMessageReply, err error)
	GetSyncMessage(ctx context.Context, req *SyncMessageRequest, opts ...http.CallOption) (rsp *SyncMessageReply, err error)
	Send(ctx context.Context, req *SendMessageRequest, opts ...http.CallOption) (rsp *SendMessageReply, err error)
	SendGroup(ctx context.Context, req *SendGroupRequest, opts ...http.CallOption) (rsp *SendGroupReply, err error)
}

type TimelineHTTPClientImpl struct {
	cc *http.Client
}

func NewTimelineHTTPClient(client *http.Client) TimelineHTTPClient {
	return &TimelineHTTPClientImpl{client}
}

func (c *TimelineHTTPClientImpl) GetGroupHistoryMessage(ctx context.Context, in *GetGroupHistoryMessageRequest, opts ...http.CallOption) (*GetGroupHistoryMessageReply, error) {
	var out GetGroupHistoryMessageReply
	pattern := "/timeline/history/group/{group}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTimelineGetGroupHistoryMessage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TimelineHTTPClientImpl) GetSingleHistoryMessage(ctx context.Context, in *GetSingleHistoryMessageRequest, opts ...http.CallOption) (*GetSingleHistoryMessageReply, error) {
	var out GetSingleHistoryMessageReply
	pattern := "/timeline/history/single/{from}/{to}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTimelineGetSingleHistoryMessage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TimelineHTTPClientImpl) GetSyncMessage(ctx context.Context, in *SyncMessageRequest, opts ...http.CallOption) (*SyncMessageReply, error) {
	var out SyncMessageReply
	pattern := "/timeline/sync"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTimelineGetSyncMessage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TimelineHTTPClientImpl) Send(ctx context.Context, in *SendMessageRequest, opts ...http.CallOption) (*SendMessageReply, error) {
	var out SendMessageReply
	pattern := "/timeline/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTimelineSend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TimelineHTTPClientImpl) SendGroup(ctx context.Context, in *SendGroupRequest, opts ...http.CallOption) (*SendGroupReply, error) {
	var out SendGroupReply
	pattern := "/timeline/sendGroup"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTimelineSendGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
