package server

import (
	"context"
	"github.com/graydovee/clipboardshare/pkg/common"
	"github.com/graydovee/clipboardshare/pkg/logger"
	"github.com/graydovee/clipboardshare/pkg/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

type ClipboardServer struct {
	proto.UnimplementedClipboardServer

	data           []byte
	lastUpdateTime time.Time

	mu sync.RWMutex
}

func NewClipboardServer() *ClipboardServer {
	return &ClipboardServer{
		lastUpdateTime: time.Now(),
	}
}

func (c *ClipboardServer) UpdateClipboardMsg(ctx context.Context, request *proto.UpdateClipboardMsgRequest) (*proto.UpdateClipboardMsgResponse, error) {
	if request == nil {
		return &proto.UpdateClipboardMsgResponse{
			Code: common.CodeEmpty,
		}, nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	updateTime := time.UnixMilli(request.Timestamp)
	if code := c.checkTimestamp(updateTime); code != common.CodeOk {
		logger.Sugar().Errorw("timestamp error", "data", string(request.Data), "updateTime", updateTime, "storedTimeStamp", c.lastUpdateTime, "code", code)
		return &proto.UpdateClipboardMsgResponse{
			Code: code,
		}, nil
	}

	c.lastUpdateTime = updateTime
	c.data = request.Data
	logger.Sugar().Infow("receive msg", "data", string(c.data), "lastUpdateTime", c.lastUpdateTime)
	return &proto.UpdateClipboardMsgResponse{
		Code: common.CodeOk,
	}, nil
}

func (c *ClipboardServer) GetClipboardMsg(ctx context.Context, request *proto.GetClipboardMsgRequest) (*proto.GetClipboardMsgResponse, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var data []byte
	if c.data != nil {
		data = make([]byte, len(c.data))
		copy(data, c.data)
	}
	return &proto.GetClipboardMsgResponse{
		Data:           data,
		LastUpdateTime: c.lastUpdateTime.UnixMilli(),
	}, nil
}

func (c *ClipboardServer) checkTimestamp(updateTime time.Time) int32 {
	// todo my vm server time incorrect
	now := time.Now().Add(time.Second * 5)
	if updateTime.After(now) {
		return common.CodeTimeIllegal
	}
	if updateTime.Before(c.lastUpdateTime) {
		return common.CodeNotLatest
	}
	return common.CodeOk
}

func (c *ClipboardServer) Start(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Sugar().Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterClipboardServer(s, c)
	logger.Sugar().Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Sugar().Errorf("failed to serve: %v", err)
	}
}
