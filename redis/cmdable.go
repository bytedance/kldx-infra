package redis

import (
	"context"
	"time"
)

type IRedis interface {
	TTL(ctx context.Context, key string) *DurationCmd
	Type(ctx context.Context, key string) *StatusCmd
	Append(ctx context.Context, key, value string) *IntCmd
	GetRange(ctx context.Context, key string, start, end int64) *StringCmd
	GetSet(ctx context.Context, key string, value interface{}) *StringCmd
	Get(ctx context.Context, key string) *StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
	Del(ctx context.Context, keys ...string) *IntCmd
	Exists(ctx context.Context, keys ...string) *IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
	Persist(ctx context.Context, key string) *BoolCmd
	PExpire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	PExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
	PTTL(ctx context.Context, key string) *DurationCmd
	Incr(ctx context.Context, key string) *IntCmd
	Decr(ctx context.Context, key string) *IntCmd
	IncrBy(ctx context.Context, key string, value int64) *IntCmd
	DecrBy(ctx context.Context, key string, value int64) *IntCmd
	IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd
	MGet(ctx context.Context, keys ...string) *SliceCmd
	MSet(ctx context.Context, pairs ...interface{}) *StatusCmd
	SetNX(ctx context.Context, key string, value interface{}) *BoolCmd
	SetXX(ctx context.Context, key string, value interface{}) *BoolCmd
	SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd
	StrLen(ctx context.Context, key string) *IntCmd
	GetBit(ctx context.Context, key string, offset int64) *IntCmd
	SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd
	BitCount(ctx context.Context, key string, bitCount *BitCountArgs) *IntCmd
	HDel(ctx context.Context, key string, fields ...string) *IntCmd
	HExists(ctx context.Context, key, field string) *BoolCmd
	HGet(ctx context.Context, key, field string) *StringCmd
	HGetAll(ctx context.Context, key string) *StringStringMapCmd
	HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd
	HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd
	HKeys(ctx context.Context, key string) *StringSliceCmd
	HLen(ctx context.Context, key string) *IntCmd
	HMSet(ctx context.Context, key string, pairs ...interface{}) *StatusCmd
	HMGet(ctx context.Context, key string, fields ...string) *SliceCmd
	HSet(ctx context.Context, key string, field string, value interface{}) *BoolCmd
	HSetNX(ctx context.Context, key, field string, value interface{}) *BoolCmd
	HVals(ctx context.Context, key string) *StringSliceCmd
	LIndex(ctx context.Context, key string, index int64) *StringCmd
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) *IntCmd
	LLen(ctx context.Context, key string) *IntCmd
	LPop(ctx context.Context, key string) *StringCmd
	LPush(ctx context.Context, key string, values ...interface{}) *IntCmd
	LPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	LRem(ctx context.Context, key string, count int64, value interface{}) *IntCmd
	LSet(ctx context.Context, key string, index int64, value interface{}) *StatusCmd
	LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd
	RPop(ctx context.Context, key string) *StringCmd
	RPush(ctx context.Context, key string, values ...interface{}) *IntCmd
	RPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
	SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd
	SCard(ctx context.Context, key string) *IntCmd
	SDiff(ctx context.Context, keys ...string) *StringSliceCmd
	SDiffStore(destination string, ctx context.Context, keys ...string) *IntCmd
	SInter(ctx context.Context, keys ...string) *StringSliceCmd
	SInterStore(destination string, ctx context.Context, keys ...string) *IntCmd
	SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd
	SMembers(ctx context.Context, key string) *StringSliceCmd
	SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd
	SPop(ctx context.Context, key string) *StringCmd
	SPopN(ctx context.Context, key string, count int64) *StringSliceCmd
	SRandMember(ctx context.Context, key string) *StringCmd
	SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd
	SRem(ctx context.Context, key string, members ...interface{}) *IntCmd
	SUnion(ctx context.Context, keys ...string) *StringSliceCmd
	SUnionStore(destination string, ctx context.Context, keys ...string) *IntCmd
	ZAdd(ctx context.Context, key string, members ...*Z) *IntCmd
	ZAddNX(ctx context.Context, key string, members ...*Z) *IntCmd
	ZAddXX(ctx context.Context, key string, members ...*Z) *IntCmd
	ZAddCh(ctx context.Context, key string, members ...*Z) *IntCmd
	ZAddNXCh(ctx context.Context, key string, members ...*Z) *IntCmd
	ZAddXXCh(ctx context.Context, key string, members ...*Z) *IntCmd
	ZIncr(ctx context.Context, key string, member *Z) *FloatCmd
	ZIncrNX(ctx context.Context, key string, member *Z) *FloatCmd
	ZIncrXX(ctx context.Context, key string, member *Z) *FloatCmd
	ZCard(ctx context.Context, key string) *IntCmd
	ZCount(ctx context.Context, key, min, max string) *IntCmd
	ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd
	ZInterStore(ctx context.Context, destination string, store *ZStore) *IntCmd
	ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
	ZRank(ctx context.Context, key, member string) *IntCmd
	ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd
	ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd
	ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	ZRevRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
	ZRevRank(ctx context.Context, key, member string) *IntCmd
	ZScore(ctx context.Context, key, member string) *FloatCmd
	ZUnionStore(ctx context.Context, dest string, store *ZStore) *IntCmd
	PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd
	PFCount(ctx context.Context, keys ...string) *IntCmd
	PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd
}
