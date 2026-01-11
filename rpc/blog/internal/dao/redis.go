package dao

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	// BlogLikedKey 博客点赞key前缀
	BlogLikedKey = "blog:liked:"

	// FeedKey 用户博客流key前缀
	FeedKey = "feed:"
)

// IsLiked 判断用户是否点赞了博客
func (d *Dao) IsLiked(ctx context.Context, blogId, userId uint64) (bool, error) {
	key := fmt.Sprintf("%s%d", BlogLikedKey, blogId)
	score, err := d.Redis.ZscoreCtx(ctx, key, strconv.FormatUint(userId, 10))
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	return score > 0, nil
}

// LikeBlog 点赞博客
func (d *Dao) LikeBlog(ctx context.Context, blogId, userId uint64) error {
	key := fmt.Sprintf("%s%d", BlogLikedKey, blogId)
	timestamp := time.Now().Unix()
	_, err := d.Redis.ZaddCtx(ctx, key, timestamp, strconv.FormatUint(userId, 10))
	return err
}

// UnlikeBlog 取消点赞
func (d *Dao) UnlikeBlog(ctx context.Context, blogId, userId uint64) error {
	key := fmt.Sprintf("%s%d", BlogLikedKey, blogId)
	_, err := d.Redis.ZremCtx(ctx, key, strconv.FormatUint(userId, 10))
	return err
}

// GetTopLikers 获取前N个点赞用户ID
func (d *Dao) GetTopLikers(ctx context.Context, blogId uint64, top int64) ([]uint64, error) {
	key := fmt.Sprintf("%s%d", BlogLikedKey, blogId)
	vals, err := d.Redis.ZrangeCtx(ctx, key, 0, top-1)
	if err != nil {
		return nil, err
	}

	userIds := make([]uint64, 0, len(vals))
	for _, val := range vals {
		if userId, err := strconv.ParseUint(val, 10, 64); err == nil {
			userIds = append(userIds, userId)
		}
	}
	return userIds, nil
}

// PushFeed 推送博客到粉丝的Feed流
func (d *Dao) PushFeed(ctx context.Context, userId, blogId uint64) error {
	key := fmt.Sprintf("%s%d", FeedKey, userId)
	timestamp := time.Now().UnixMilli()
	_, err := d.Redis.ZaddCtx(ctx, key, timestamp, strconv.FormatUint(blogId, 10))
	return err
}

// ScrollFeed 滚动分页查询Feed流
// maxScore: 上一次查询的最小时间戳（第一次为当前时间）
// offset: 偏移量（用于处理相同时间戳的情况）
// count: 每次获取的数量
func (d *Dao) ScrollFeed(ctx context.Context, userId uint64, maxScore, minScore float64, offset, count int64) ([]uint64, float64, int, error) {
	key := fmt.Sprintf("%s%d", FeedKey, userId)

	// ZREVRANGEBYSCORE key max min WITHSCORES LIMIT offset count
	pairs, err := d.Redis.ZrevrangebyscoreWithScoresAndLimitCtx(
		ctx, key, int64(maxScore), int64(minScore), int(offset), int(count),
	)
	if err != nil {
		return nil, 0, 0, err
	}

	if len(pairs) == 0 {
		return []uint64{}, 0, 0, nil
	}

	blogIds := make([]uint64, 0, len(pairs))
	var minTime float64
	offsetCount := 1

	for i, pair := range pairs {
		if blogId, err := strconv.ParseUint(pair.Key, 10, 64); err == nil {
			blogIds = append(blogIds, blogId)
		}

		// 记录最小时间戳
		if i == 0 {
			minTime = float64(pair.Score)
		}

		// 统计与最小时间戳相同的数量（作为下次的offset）
		if float64(pair.Score) == minTime {
			offsetCount = i + 1
		}
	}

	return blogIds, minTime, offsetCount, nil
}
