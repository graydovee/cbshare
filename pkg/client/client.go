package client

import (
	"bytes"
	"context"
	"github.com/graydovee/clipboardshare/pkg/common"
	"github.com/graydovee/clipboardshare/pkg/logger"
	"github.com/graydovee/clipboardshare/pkg/proto"
	"golang.design/x/clipboard"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
	"time"
)

type ClipboardClient struct {
	cli proto.ClipboardClient

	lastUpdateTime time.Time
	data           []byte

	mu sync.RWMutex
}

func NewClipboardClient() *ClipboardClient {
	return &ClipboardClient{}
}

func (c *ClipboardClient) Start(addr string) {
	err := clipboard.Init()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()
	c.cli = proto.NewClipboardClient(conn)

	ticker := time.NewTicker(time.Second * 1)
	go c.watchClipboard()
	for {
		select {
		case <-ticker.C:
			c.requestClipboard()
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
	clipboard.Write(clipboard.FmtText, resp.Data)
	if c.UpdateLocalClipboard(resp.Data, lastUpdateTime) {
		logger.Sugar().Infow("update local clipboard", "data", string(resp.Data))
	}
}

func (c *ClipboardClient) watchClipboard() {
	ch := clipboard.Watch(context.Background(), clipboard.FmtText)
	for data := range ch {
		logger.Sugar().Infow("watched clipboard changed", "data", string(data))
		c.updateClipboard(data)
	}
}

func (c *ClipboardClient) updateClipboard(data []byte) {
	now := time.Now()
	if !c.UpdateLocalClipboard(data, now) {
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
		logger.Sugar().Errorw("update clipboard error", "code", resp.Code, "data", string(data), "timestamp")
		return
	}

	logger.Sugar().Infow("update remote clipboard", "data", string(data))
}

func (c *ClipboardClient) unsafeIsNeedUpdate(data []byte, lastUpdateTime time.Time) bool {
	return lastUpdateTime.After(c.lastUpdateTime) && !bytes.Equal(data, c.data)
}

func (c *ClipboardClient) IsNeedUpdate(data []byte, lastUpdateTime time.Time) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.unsafeIsNeedUpdate(data, lastUpdateTime)
}

func (c *ClipboardClient) UpdateLocalClipboard(data []byte, lastUpdateTime time.Time) bool {
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

func (c *ClipboardClient) GetData() []byte {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data
}
