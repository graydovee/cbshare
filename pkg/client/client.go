package client

import (
	"context"
	"github.com/atotto/clipboard"
	"github.com/graydovee/clipboardshare/pkg/common"
	"github.com/graydovee/clipboardshare/pkg/logger"
	"github.com/graydovee/clipboardshare/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
	"time"
)

type ClipboardClient struct {
	cli proto.ClipboardClient

	lastUpdateTime time.Time
	data           string

	mu sync.RWMutex
}

func NewClipboardClient() *ClipboardClient {
	return &ClipboardClient{}
}

func (c *ClipboardClient) Start(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()
	c.cli = proto.NewClipboardClient(conn)

	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			c.requestClipboard()
			c.updateClipboard()
		}
	}
}

func (c *ClipboardClient) requestClipboard() {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := c.cli.GetClipboardMsg(ctx, &proto.GetClipboardMsgRequest{})
	if err != nil {
		logger.Sugar().Errorw("sync clipboard msg error", "err", err)
		return
	}

	lastUpdateTime := time.UnixMilli(resp.LastUpdateTime)
	if !c.IsNeedUpdate(resp.Data, lastUpdateTime) {
		return
	}
	if err = clipboard.WriteAll(resp.Data); err != nil {
		logger.Sugar().Errorw("update clipboard error", "err", err)
		return
	}
	if c.UpdateLocalClipboard(resp.Data, lastUpdateTime) {
		logger.Sugar().Infow("update local clipboard", "data", resp.Data)
	}
}

func (c *ClipboardClient) updateClipboard() {
	now := time.Now()
	data, err := clipboard.ReadAll()
	if err != nil {
		logger.Sugar().Errorw("read clipboard error", "err", err)
		return
	}
	if c.GetData() == data {
		return
	}
	if !c.IsNeedUpdate(data, now) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	request := &proto.UpdateClipboardMsgRequest{
		Data:      data,
		Timestamp: now.UnixMilli(),
	}
	resp, err := c.cli.UpdateClipboardMsg(ctx, request)
	if err != nil {
		logger.Sugar().Errorw("request of updating clipboard error", "err", err)
		return
	}

	if resp.Code != common.CodeOk {
		logger.Sugar().Errorw("update clipboard error", "code", resp.Code, "data", data, "timestamp")
		return
	}

	if c.UpdateLocalClipboard(data, now) {
		logger.Sugar().Infow("update remote clipboard", "data", data)
	}
}

func (c *ClipboardClient) unsafeIsNeedUpdate(data string, lastUpdateTime time.Time) bool {
	return lastUpdateTime.After(c.lastUpdateTime) && data != c.data
}

func (c *ClipboardClient) IsNeedUpdate(data string, lastUpdateTime time.Time) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.unsafeIsNeedUpdate(data, lastUpdateTime)
}

func (c *ClipboardClient) UpdateLocalClipboard(data string, lastUpdateTime time.Time) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.unsafeIsNeedUpdate(data, lastUpdateTime) {
		return false
	}
	c.data = data
	c.lastUpdateTime = lastUpdateTime
	return true
}

func (c *ClipboardClient) GetLastUpdateTime() time.Time {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lastUpdateTime
}

func (c *ClipboardClient) GetData() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data
}
