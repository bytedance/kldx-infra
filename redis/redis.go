package redis

import (
	"context"
	"time"
)

type Redis struct{}

func NewRedis() *Redis {
	return &Redis{}
}

func (c *Redis) TTL(ctx context.Context, key string) *DurationCmd {
	cmd := NewDurationCmd(c, "ttl", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) Type(ctx context.Context, key string) *StatusCmd {
	cmd := NewStatusCmd(c, "type", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) Append(ctx context.Context, key, value string) *IntCmd {
	cmd := NewIntCmd(c, "append", key, value)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) GetRange(ctx context.Context, key string, start, end int64) *StringCmd {
	cmd := NewStringCmd(c, "getrange", key, start, end)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) GetSet(ctx context.Context, key string, value interface{}) *StringCmd {
	cmd := NewStringCmd(c, "getset", key, value)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) Get(ctx context.Context, key string) *StringCmd {
	cmd := NewStringCmd(c, "get", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd {
	args := []interface{}{key, value}
	if expiration > 0 {
		if usePrecise(expiration) {
			args = append(args, "px", formatMs(expiration))
		} else {
			args = append(args, "ex", formatSec(expiration))
		}
	}
	cmd := NewStatusCmd(c, "set", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) Del(ctx context.Context, keys ...string) *IntCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewIntCmd(c, "del", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) Exists(ctx context.Context, keys ...string) *IntCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewIntCmd(c, "exists", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	cmd := NewBoolCmd(c, "expire", key, formatSec(expiration))
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd {
	cmd := NewBoolCmd(c, "expireat", key, tm.Unix())
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) Persist(ctx context.Context, key string) *BoolCmd {
	cmd := NewBoolCmd(c, "persist", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) PExpire(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	cmd := NewBoolCmd(c, "pexpire", key, formatMs(expiration))
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) PExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd {
	cmd := NewBoolCmd(c, "pexpireat", key, tm.UnixNano()/int64(time.Millisecond))
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) PTTL(ctx context.Context, key string) *DurationCmd {
	cmd := NewDurationCmd(c, "pttl", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) Incr(ctx context.Context, key string) *IntCmd {
	cmd := NewIntCmd(c, "incr", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) Decr(ctx context.Context, key string) *IntCmd {
	cmd := NewIntCmd(c, "decr", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) IncrBy(ctx context.Context, key string, value int64) *IntCmd {
	cmd := NewIntCmd(c, "incrby", key, value)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) DecrBy(ctx context.Context, key string, value int64) *IntCmd {
	cmd := NewIntCmd(c, "decrby", key, value)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd {
	cmd := NewFloatCmd(c, "incrbyfloat", key, value)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) MGet(ctx context.Context, keys ...string) *SliceCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewSliceCmd(c, "mget", args...)
	cmd.execute(ctx)
	return cmd
}

// MSet is like Set but accepts multiple values:
//   - MSet("key1", "value1", "key2", "value2")
//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func (c *Redis) MSet(ctx context.Context, pairs ...interface{}) *StatusCmd {
	cmd := NewStatusCmd(c, "mset", pairs...)
	cmd.execute(ctx)
	return cmd
}

// SetNX is short for "SET if Not exists".
func (c *Redis) SetNX(ctx context.Context, key string, value interface{}) *BoolCmd {
	cmd := NewBoolCmd(c, "setnx", key, value)
	cmd.execute(ctx)
	return cmd
}

// SetXX is short for "SET if exists".
func (c *Redis) SetXX(ctx context.Context, key string, value interface{}) *BoolCmd {
	cmd := NewBoolCmd(c, "set", key, value, "xx")
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd {
	cmd := NewIntCmd(c, "setrange", key, offset, value)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) StrLen(ctx context.Context, key string) *IntCmd {
	cmd := NewIntCmd(c, "strlen", key)
	cmd.execute(ctx)
	return cmd
}

//------------------------------------------------------------------------------
// Bit

func (c *Redis) GetBit(ctx context.Context, key string, offset int64) *IntCmd {
	cmd := NewIntCmd(c, "getbit", key, offset)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd {
	cmd := NewIntCmd(c, "setbit", key, offset, value)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) BitCount(ctx context.Context, key string, bitCount *BitCountArgs) *IntCmd {
	args := []interface{}{key}
	if bitCount != nil {
		args = append(args, bitCount.Start, bitCount.End)
	}
	cmd := NewIntCmd(c, "bitcount", args...)
	cmd.execute(ctx)
	return cmd
}

//------------------------------------------------------------------------------
// Hash
func (c *Redis) HDel(ctx context.Context, key string, fields ...string) *IntCmd {
	args := make([]interface{}, 1+len(fields))
	args[0] = key
	for i, field := range fields {
		args[i+1] = field
	}
	cmd := NewIntCmd(c, "hdel", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) HExists(ctx context.Context, key, field string) *BoolCmd {
	cmd := NewBoolCmd(c, "hexists", key, field)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) HGet(ctx context.Context, key, field string) *StringCmd {
	cmd := NewStringCmd(c, "hget", key, field)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) HGetAll(ctx context.Context, key string) *StringStringMapCmd {
	cmd := NewStringStringMapCmd(c, "hgetall", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd {
	cmd := NewIntCmd(c, "hincrby", key, field, incr)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd {
	cmd := NewFloatCmd(c, "hincrbyfloat", key, field, incr)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) HKeys(ctx context.Context, key string) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "hkeys", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) HLen(ctx context.Context, key string) *IntCmd {
	cmd := NewIntCmd(c, "hlen", key)
	cmd.execute(ctx)
	return cmd
}

// HMGet returns the values for the specified fields in the hash stored at key.
// It returns an interface{} to distinguish between empty string and nil value.
func (c *Redis) HMGet(ctx context.Context, key string, fields ...string) *SliceCmd {
	args := make([]interface{}, 1+len(fields))
	args[0] = key
	for i, field := range fields {
		args[i+1] = field
	}
	cmd := NewSliceCmd(c, "hmget", args...)
	cmd.execute(ctx)
	return cmd
}

// HSet accepts values in following formats:
//   - HSet("myhash", "key1", "value1", "key2", "value2")
//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
//
// Note that it requires Redis v4 for multiple field/value pairs support.
func (c *Redis) HSet(ctx context.Context, key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = appendArgs(args, values)
	cmd := NewIntCmd(c, "hset", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) HSetNX(ctx context.Context, key, field string, value interface{}) *BoolCmd {
	cmd := NewBoolCmd(c, "hsetnx", key, field, value)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) HVals(ctx context.Context, key string) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "hvals", key)
	cmd.execute(ctx)
	return cmd
}

//------------------------------------------------------------------------------
// List
func (c *Redis) LIndex(ctx context.Context, key string, index int64) *StringCmd {
	cmd := NewStringCmd(c, "lindex", key, index)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *IntCmd {
	cmd := NewIntCmd(c, "linsert", key, op, pivot, value)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) LLen(ctx context.Context, key string) *IntCmd {
	cmd := NewIntCmd(c, "llen", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) LPop(ctx context.Context, key string) *StringCmd {
	cmd := NewStringCmd(c, "lpop", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) LPush(ctx context.Context, key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = appendArgs(args, values)
	cmd := NewIntCmd(c, "lpush", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) LPushX(ctx context.Context, key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = appendArgs(args, values)
	cmd := NewIntCmd(c, "lpushx", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "lrange", key, start, stop)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) LRem(ctx context.Context, key string, count int64, value interface{}) *IntCmd {
	cmd := NewIntCmd(c, "lrem", key, count, value)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) LSet(ctx context.Context, key string, index int64, value interface{}) *StatusCmd {
	cmd := NewStatusCmd(c, "lset", key, index, value)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd {
	cmd := NewStatusCmd(c, "ltrim", key, start, stop)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) RPop(ctx context.Context, key string) *StringCmd {
	cmd := NewStringCmd(c, "rpop", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) RPush(ctx context.Context, key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = append(args, values...)
	cmd := NewIntCmd(c, "rpush", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) RPushX(ctx context.Context, key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = append(args, values...)
	cmd := NewIntCmd(c, "rpushx", args...)
	cmd.execute(ctx)
	return cmd
}

//------------------------------------------------------------------------------
// Set

func (c *Redis) SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(members))
	args[0] = key
	args = append(args, members...)
	cmd := NewIntCmd(c, "sadd", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SCard(ctx context.Context, key string) *IntCmd {
	cmd := NewIntCmd(c, "scard", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SDiff(ctx context.Context, keys ...string) *StringSliceCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewStringSliceCmd(c, "sdiff", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SDiffStore(destination string, ctx context.Context, keys ...string) *IntCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = destination
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewIntCmd(c, "sdiffstore", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SInter(ctx context.Context, keys ...string) *StringSliceCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewStringSliceCmd(c, "sinter", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SInterStore(destination string, ctx context.Context, keys ...string) *IntCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = destination
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewIntCmd(c, "sinterstore", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd {
	cmd := NewBoolCmd(c, "sismember", key, member)
	cmd.execute(ctx)
	return cmd
}

// SMembers `SMEMBERS key` command output as a slice.
func (c *Redis) SMembers(ctx context.Context, key string) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "smembers", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd {
	cmd := NewBoolCmd(c, "smove", source, destination, member)
	cmd.execute(ctx)
	return cmd
}

// SPop `SPOP key` command.
func (c *Redis) SPop(ctx context.Context, key string) *StringCmd {
	cmd := NewStringCmd(c, "spop", key)
	cmd.execute(ctx)
	return cmd
}

// SPOP `SPOP ctx context.Context, key count` command.
func (c *Redis) SPopN(ctx context.Context, key string, count int64) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "spop", key, count)
	cmd.execute(ctx)
	return cmd
}

// SRandMember `SRANDMEMBER key` command.
func (c *Redis) SRandMember(ctx context.Context, key string) *StringCmd {
	cmd := NewStringCmd(c, "srandmember", key)
	cmd.execute(ctx)
	return cmd
}

// SRandMemberN `SRANDMEMBER ctx context.Context, key count` command.
func (c *Redis) SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "srandmember", key, count)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SRem(ctx context.Context, key string, members ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(members))
	args[0] = key
	args = append(args, members...)
	cmd := NewIntCmd(c, "srem", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SUnion(ctx context.Context, keys ...string) *StringSliceCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewStringSliceCmd(c, "sunion", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) SUnionStore(destination string, ctx context.Context, keys ...string) *IntCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = destination
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewIntCmd(c, "sunionstore", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) executeZSetIntCmd(ctx context.Context, name string, key string, members ...*Z) *IntCmd {
	l := make([]interface{}, 2*len(members)+1)
	l[0] = key
	for i := range members {
		l[2*i+1] = members[i].Score
		l[2*i+2] = members[i].Member
	}

	cmd := NewIntCmd(c, name, l...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) executeZSetFloatCmd(ctx context.Context, name string, key string, members ...*Z) *FloatCmd {
	l := make([]interface{}, 2*len(members)+1)
	l[0] = key
	for i := range members {
		l[i+1] = members[i].Score
		l[i+2] = members[i].Member
	}

	cmd := NewFloatCmd(c, name, l...)
	cmd.execute(ctx)
	return cmd
}

// ZAdd `ZADD ctx context.Context, key score member [score member ...]` command.
func (c *Redis) ZAdd(ctx context.Context, key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd(ctx, "zadd", key, members...)
}

// ZAddNX `ZADD ctx context.Context, key NX score member [score member ...]` command.
func (c *Redis) ZAddNX(ctx context.Context, key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd(ctx, "zaddnx", key, members...)
}

// ZAddXX `ZADD ctx context.Context, key XX score member [score member ...]` command.
func (c *Redis) ZAddXX(ctx context.Context, key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd(ctx, "zaddxx", key, members...)
}

// ZAddCh `ZADD ctx context.Context, key CH score member [score member ...]` command.
func (c *Redis) ZAddCh(ctx context.Context, key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd(ctx, "zaddch", key, members...)
}

// ZAddNXCh `ZADD ctx context.Context, key NX CH score member [score member ...]` command.
func (c *Redis) ZAddNXCh(ctx context.Context, key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd(ctx, "zaddnxch", key, members...)
}

// ZAddXXCh `ZADD ctx context.Context, key XX CH score member [score member ...]` command.
func (c *Redis) ZAddXXCh(ctx context.Context, key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd(ctx, "zaddxxch", key, members...)
}

// ZIncr `ZADD ctx context.Context, key INCR score member` command.
func (c *Redis) ZIncr(ctx context.Context, key string, member *Z) *FloatCmd {
	return c.executeZSetFloatCmd(ctx, "zincr", key, member)
}

// ZIncrNX `ZADD ctx context.Context, key NX INCR score member` command.
func (c *Redis) ZIncrNX(ctx context.Context, key string, member *Z) *FloatCmd {
	return c.executeZSetFloatCmd(ctx, "zincrnx", key, member)
}

// ZIncrXX `ZADD ctx context.Context, key XX INCR score member` command.
func (c *Redis) ZIncrXX(ctx context.Context, key string, member *Z) *FloatCmd {
	return c.executeZSetFloatCmd(ctx, "zincrxx", key, member)
}

func (c *Redis) ZCard(ctx context.Context, key string) *IntCmd {
	cmd := NewIntCmd(c, "zcard", key)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZCount(ctx context.Context, key, min, max string) *IntCmd {
	cmd := NewIntCmd(c, "zcount", key, min, max)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd {
	cmd := NewFloatCmd(c, "zincrby", key, increment, member)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZInterStore(ctx context.Context, destination string, store *ZStore) *IntCmd {
	args := make([]interface{}, 2+len(store.Keys))
	args[0] = destination
	args[1] = len(store.Keys)
	for i, key := range store.Keys {
		args[2+i] = key
	}
	if len(store.Weights) > 0 {
		args = append(args, "weights", len(store.Weights))
		for _, weight := range store.Weights {
			args = append(args, weight)
		}
	}
	if store.Aggregate != "" {
		args = append(args, "aggregate", store.Aggregate)
	}
	cmd := NewIntCmd(c, "zinterstore", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) zRange(ctx context.Context, key string, start, stop int64, withScores bool) *StringSliceCmd {
	args := []interface{}{key, start, stop}
	if withScores {
		args = append(args, "withscores")
	}
	cmd := NewStringSliceCmd(c, "zrange", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	return c.zRange(ctx, key, start, stop, false)
}

func (c *Redis) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	cmd := NewZSliceCmd(c, "zrange", key, start, stop, "withscores")
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) zRangeBy(ctx context.Context, zcmd string, key string, opt *ZRangeBy, withScores bool) *StringSliceCmd {
	args := []interface{}{key, opt.Min, opt.Max}
	if withScores {
		args = append(args, "withscores")
	}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, "limit", opt.Offset, opt.Count)
	}
	cmd := NewStringSliceCmd(c, zcmd, args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {
	return c.zRangeBy(ctx,"zrangebyscore", key, opt, false)
}

func (c *Redis) ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd {
	args := []interface{}{key, opt.Min, opt.Max, "withscores"}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, "limit", opt.Offset, opt.Count)
	}
	cmd := NewZSliceCmd(c, "zrangebyscore", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZRank(ctx context.Context, key, member string) *IntCmd {
	cmd := NewIntCmd(c, "zrank", key, member)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd {
	args := make([]interface{}, 2, 2+len(members))
	args[0] = key
	args = appendArgs(args, members)
	cmd := NewIntCmd(c, "zrem", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd {
	cmd := NewIntCmd(c, "zremrangebyrank", key, start, stop)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd {
	cmd := NewIntCmd(c, "zremrangebyscore", key, min, max)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "zrevrange", key, start, stop)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	cmd := NewZSliceCmd(c, "zrevrange", key, start, stop, "withscores")
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) zRevRangeBy(ctx context.Context, zcmd, key string, opt *ZRangeBy) *StringSliceCmd {
	args := []interface{}{key, opt.Max, opt.Min}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, "limit", opt.Offset, opt.Count)
	}
	cmd := NewStringSliceCmd(c, zcmd, args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZRevRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd {
	return c.zRevRangeBy(ctx,"zrevrangebyscore", key, opt)
}

func (c *Redis) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd {
	args := []interface{}{key, opt.Min, opt.Max, "withscores"}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, "limit", opt.Offset, opt.Count)
	}
	cmd := NewZSliceCmd(c, "zrevrangebyscore", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZRevRank(ctx context.Context, key, member string) *IntCmd {
	cmd := NewIntCmd(c, "zrevrank", key, member)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZScore(ctx context.Context, key, member string) *FloatCmd {
	cmd := NewFloatCmd(c, "zscore", key, member)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) ZUnionStore(ctx context.Context, dest string, store *ZStore) *IntCmd {
	args := make([]interface{}, 2+len(store.Keys))
	args[0] = dest
	args[1] = len(store.Keys)
	for i, key := range store.Keys {
		args[2+i] = key
	}
	if len(store.Weights) > 0 {
		args = append(args, "weights", len(store.Weights))
		for _, weight := range store.Weights {
			args = append(args, weight)
		}
	}
	if store.Aggregate != "" {
		args = append(args, "aggregate", store.Aggregate)
	}

	cmd := NewIntCmd(c, "zunionstore", args...)
	cmd.execute(ctx)
	return cmd
}

//------------------------------------------------------------------------------
// HyperLogLog

func (c *Redis) PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(els))
	args[0] = key
	args = appendArgs(args, els)
	cmd := NewIntCmd(c, "pfadd", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) PFCount(ctx context.Context, keys ...string) *IntCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewIntCmd(c, "pfcount", args...)
	cmd.execute(ctx)
	return cmd
}

func (c *Redis) PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = dest
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewStatusCmd(c, "pfmerge", args...)
	cmd.execute(ctx)
	return cmd
}
