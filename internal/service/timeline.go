package service

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/gomicroim/go-timeline/api/timeline/v1"
	"github.com/gomicroim/go-timeline/internal/biz"
)

type TimelineService struct {
	pb.UnimplementedTimelineServer

	msgUseCase biz.MessageUseCase
	log        *log.Helper
}

func NewTimelineService(msgUseCase biz.MessageUseCase, logger log.Logger) *TimelineService {
	return &TimelineService{
		msgUseCase: msgUseCase,
		log:        log.NewHelper(logger),
	}
}

func (s *TimelineService) GetSyncMessage(ctx context.Context, req *pb.SyncMessageRequest) (*pb.SyncMessageReply, error) {
	msgArr, latestSeq, err := s.msgUseCase.GetSyncMessage(ctx, req.Member, req.LastRead, int(req.Count))
	if err != nil {
		return nil, err
	}

	reply := &pb.SyncMessageReply{
		EntrySet:  make([]*pb.TimelineEntry, len(msgArr)),
		LatestSeq: latestSeq,
	}

	if len(msgArr) > 0 {
		reply.EntrySetLastSeq = msgArr[len(msgArr)-1].Seq
	}

	for k, v := range msgArr {
		reply.EntrySet[k] = &pb.TimelineEntry{
			Sequence: v.Seq,
		}
		buffer, err := json.Marshal(v.Message)
		if err != nil {
			reply.EntrySet[k].Message = err.Error()
		} else {
			reply.EntrySet[k].Message = string(buffer)
		}
	}
	return reply, nil
}

func (s *TimelineService) Send(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageReply, error) {
	if req.From == "" || req.To == "" {
		return nil, errors.BadRequest(pb.ErrorReason_TIMELINE_INVALID_PARAM.String(), "miss from or to filed")
	}

	seq, err := s.msgUseCase.Send(ctx, req.From, req.To, req.Message)
	if err != nil {
		return nil, err
	}
	return &pb.SendMessageReply{Sequence: seq}, nil
}

func (s *TimelineService) SendGroup(ctx context.Context, req *pb.SendGroupRequest) (*pb.SendGroupReply, error) {
	if req.GroupName == "" || len(req.GroupMembers) <= 0 {
		return nil, errors.BadRequest(pb.ErrorReason_TIMELINE_INVALID_PARAM.String(), "miss group name or group members")
	}
	r, err := s.msgUseCase.SendGroup(ctx, req.GroupName, req.GroupMembers, req.Message)
	if err != nil {
		return nil, err
	}

	reply := &pb.SendGroupReply{
		FailedMembers: make([]string, 0),
	}
	for _, v := range r {
		if v.Err != nil {
			reply.FailedMembers = append(reply.FailedMembers, v.GroupMember)
		}
	}
	return reply, nil
}

func (s *TimelineService) GetSingleHistoryMessage(ctx context.Context, req *pb.GetSingleHistoryMessageRequest) (*pb.GetSingleHistoryMessageReply, error) {
	msgArr, err := s.msgUseCase.GetSingleHistoryMessage(ctx, req.From, req.To, req.LastRead, req.Count)
	if err != nil {
		return nil, err
	}

	reply := &pb.GetSingleHistoryMessageReply{
		EntrySet: make([]*pb.TimelineEntry, len(msgArr)),
	}
	for k, v := range msgArr {
		reply.EntrySet[k] = &pb.TimelineEntry{
			Sequence: v.Seq,
		}
		buffer, err := json.Marshal(v.Message)
		if err != nil {
			reply.EntrySet[k].Message = err.Error()
		} else {
			reply.EntrySet[k].Message = string(buffer)
		}
	}

	return reply, nil
}

func (s *TimelineService) GetGroupHistoryMessage(ctx context.Context, req *pb.GetGroupHistoryMessageRequest) (*pb.GetGroupHistoryMessageReply, error) {
	msgArr, err := s.msgUseCase.GetGroupHistoryMessage(ctx, req.Group, req.LastRead, req.Count)
	if err != nil {
		return nil, err
	}

	reply := &pb.GetGroupHistoryMessageReply{
		EntrySet: make([]*pb.TimelineEntry, len(msgArr)),
	}
	for k, v := range msgArr {
		reply.EntrySet[k] = &pb.TimelineEntry{
			Sequence: v.Seq,
		}
		buffer, err := json.Marshal(v.Message)
		if err != nil {
			reply.EntrySet[k].Message = err.Error()
		} else {
			reply.EntrySet[k].Message = string(buffer)
		}
	}
	return reply, nil
}
