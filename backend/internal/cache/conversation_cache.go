package cache

import (
	"context"
	"encoding/json"
	"fmt"
	model "sociomile-apps/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type ConversationCache struct {
	rdb *RedisClient
}

func NewConversationCache(rdb *RedisClient) *ConversationCache {
	return &ConversationCache{rdb: rdb}
}

type cachedConversationList struct {
	Data  []model.Conversation
	Total int64
	Page  int
	Limit int
}

func (c *ConversationCache) key(
	tenantID uuid.UUID,
	status string,
	assignedAgentID uuid.UUID,
	page int,
	limit int,
) string {
	return fmt.Sprintf("conversation:%s:%s:%s:%d:%d", tenantID, status, assignedAgentID, page, limit)
}

func (c *ConversationCache) Get(
	ctx context.Context,
	tenantID uuid.UUID,
	status string,
	assignedAgentID uuid.UUID,
	page int,
	limit int,
) ([]model.Conversation, int64, bool) {

	key := c.key(tenantID, status, assignedAgentID, page, limit)

	val, err := c.rdb.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, 0, false
	}
	if err != nil {
		return nil, 0, false
	}

	var cached cachedConversationList
	if err := json.Unmarshal([]byte(val), &cached); err != nil {
		return nil, 0, false
	}

	return cached.Data, cached.Total, true

}

func (c *ConversationCache) Set(
	ctx context.Context,
	data []model.Conversation,
	total int64,
	page int,
	limit int,
	tenantID uuid.UUID,
	status string,
	assignedAgentID uuid.UUID,
) error {
	key := c.key(tenantID, status, assignedAgentID, page, limit)

	cached := cachedConversationList{
		Data:  data,
		Total: total,
		Page:  page,
		Limit: limit,
	}

	val, err := json.Marshal(cached)
	if err != nil {
		return err
	}

	return c.rdb.Client.Set(ctx, key, val, 10*time.Minute).Err()
}
