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

type TicketCache struct {
	rdb *RedisClient
}

func NewTicketCache(rdb *RedisClient) *TicketCache {
	return &TicketCache{rdb: rdb}
}

type cachedTicketList struct {
	Data  []model.Ticket
	Total int64
	Page  int
	Limit int
}

func (c *TicketCache) Key(
	status string,
	assignedAgentID uuid.UUID,
	page int,
	limit int,
) string {
	return fmt.Sprintf("ticket_list:%s:%s:%d:%d", status, assignedAgentID, page, limit)
}

func (c *TicketCache) Set(
	ctx context.Context,
	data []model.Ticket,
	total int64,
	page int,
	limit int,
	status string,
	assignedAgentID uuid.UUID,
) error {
	key := c.Key(status, assignedAgentID, page, limit)

	val, _ := json.Marshal(cachedTicketList{
		Data:  data,
		Total: total,
		Page:  page,
		Limit: limit,
	})

	return c.rdb.Client.Set(ctx, key, val, 10*time.Minute).Err()
}

func (c *TicketCache) Get(
	ctx context.Context,
	status string,
	assignedAgentID uuid.UUID,
	page int,
	limit int,
) ([]model.Ticket, int64, bool) {
	key := c.Key(status, assignedAgentID, page, limit)

	val, err := c.rdb.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, 0, false
	}
	if err != nil {
		return nil, 0, false
	}

	var cached cachedTicketList
	if err := json.Unmarshal([]byte(val), &cached); err != nil {
		return nil, 0, false
	}

	return cached.Data, cached.Total, true
}

func (c *TicketCache) InvalidateLists(ctx context.Context) error {
	iter := c.rdb.Client.Scan(ctx, 0, "ticket_list:*", 0).Iterator()
	var keysToDelete []string
	for iter.Next(ctx) {
		keysToDelete = append(keysToDelete, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return err
	}

	if len(keysToDelete) > 0 {
		return c.rdb.Client.Del(ctx, keysToDelete...).Err()
	}
	return nil
}
