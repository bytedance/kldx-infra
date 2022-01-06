package redis

import (
	"time"
)

type Redis struct{}

func NewRedis() *Redis {
	return &Redis{}
}

func (c *Redis) TTL(key string) *DurationCmd {
	cmd := NewDurationCmd(c, "ttl", key)
	cmd.execute()
	return cmd
}

func (c *Redis) Type(key string) *StatusCmd {
	cmd := NewStatusCmd(c, "type", key)
	cmd.execute()
	return cmd
}

func (c *Redis) Append(key, value string) *IntCmd {
	cmd := NewIntCmd(c, "append", key, value)
	cmd.execute()
	return cmd
}

func (c *Redis) GetRange(key string, start, end int64) *StringCmd {
	cmd := NewStringCmd(c, "getrange", key, start, end)
	cmd.execute()
	return cmd
}

func (c *Redis) GetSet(key string, value interface{}) *StringCmd {
	cmd := NewStringCmd(c, "getset", key, value)
	cmd.execute()
	return cmd
}

func (c *Redis) Get(key string) *StringCmd {
	cmd := NewStringCmd(c, "get", key)
	cmd.execute()
	return cmd
}

func (c *Redis) Set(key string, value interface{}, expiration time.Duration) *StatusCmd {
	args := []interface{}{key, value}
	if expiration > 0 {
		if usePrecise(expiration) {
			args = append(args, "px", formatMs(expiration))
		} else {
			args = append(args, "ex", formatSec(expiration))
		}
	}
	cmd := NewStatusCmd(c, "set", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) Del(keys ...string) *IntCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewIntCmd(c, "del", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) Exists(keys ...string) *IntCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewIntCmd(c, "exists", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) Expire(key string, expiration time.Duration) *BoolCmd {
	cmd := NewBoolCmd(c, "expire", key, formatSec(expiration))
	cmd.execute()
	return cmd
}

func (c *Redis) ExpireAt(key string, tm time.Time) *BoolCmd {
	cmd := NewBoolCmd(c, "expireat", key, tm.Unix())
	cmd.execute()
	return cmd
}

func (c *Redis) Persist(key string) *BoolCmd {
	cmd := NewBoolCmd(c, "persist", key)
	cmd.execute()
	return cmd
}

func (c *Redis) PExpire(key string, expiration time.Duration) *BoolCmd {
	cmd := NewBoolCmd(c, "pexpire", key, formatMs(expiration))
	cmd.execute()
	return cmd
}

func (c *Redis) PExpireAt(key string, tm time.Time) *BoolCmd {
	cmd := NewBoolCmd(c, "pexpireat", key, tm.UnixNano()/int64(time.Millisecond))
	cmd.execute()
	return cmd
}

func (c *Redis) PTTL(key string) *DurationCmd {
	cmd := NewDurationCmd(c, "pttl", key)
	cmd.execute()
	return cmd
}

func (c *Redis) Incr(key string) *IntCmd {
	cmd := NewIntCmd(c, "incr", key)
	cmd.execute()
	return cmd
}

func (c *Redis) Decr(key string) *IntCmd {
	cmd := NewIntCmd(c, "decr", key)
	cmd.execute()
	return cmd
}

func (c *Redis) IncrBy(key string, value int64) *IntCmd {
	cmd := NewIntCmd(c, "incrby", key, value)
	cmd.execute()
	return cmd
}

func (c *Redis) DecrBy(key string, value int64) *IntCmd {
	cmd := NewIntCmd(c, "decrby", key, value)
	cmd.execute()
	return cmd
}

func (c *Redis) IncrByFloat(key string, value float64) *FloatCmd {
	cmd := NewFloatCmd(c, "incrbyfloat", key, value)
	cmd.execute()
	return cmd
}

func (c *Redis) MGet(keys ...string) *SliceCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewSliceCmd(c, "mget", args...)
	cmd.execute()
	return cmd
}

// MSet is like Set but accepts multiple values:
//   - MSet("key1", "value1", "key2", "value2")
//   - MSet([]string{"key1", "value1", "key2", "value2"})
//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func (c *Redis) MSet(values ...interface{}) *StatusCmd {
	cmd := NewStatusCmd(c, "mset", values...)
	cmd.execute()
	return cmd
}

// SetNX is short for "SET if Not exists".
func (c *Redis) SetNX(key string, value interface{}) *BoolCmd {
	cmd := NewBoolCmd(c, "setnx", key, value)
	cmd.execute()
	return cmd
}

// SetXX is short for "SET if exists".
func (c *Redis) SetXX(key string, value interface{}) *BoolCmd {
	cmd := NewBoolCmd(c, "set", key, value, "xx")
	cmd.execute()
	return cmd
}

func (c *Redis) SetRange(key string, offset int64, value string) *IntCmd {
	cmd := NewIntCmd(c, "setrange", key, offset, value)
	cmd.execute()
	return cmd
}

func (c *Redis) StrLen(key string) *IntCmd {
	cmd := NewIntCmd(c, "strlen", key)
	cmd.execute()
	return cmd
}

//------------------------------------------------------------------------------
// Bit

func (c *Redis) GetBit(key string, offset int64) *IntCmd {
	cmd := NewIntCmd(c, "getbit", key, offset)
	cmd.execute()
	return cmd
}

func (c *Redis) SetBit(key string, offset int64, value int) *IntCmd {
	cmd := NewIntCmd(c, "setbit", key, offset, value)
	cmd.execute()
	return cmd
}

func (c *Redis) BitCount(key string, bitCount *BitCountArgs) *IntCmd {
	args := []interface{}{key}
	if bitCount != nil {
		args = append(args, bitCount.Start, bitCount.End)
	}
	cmd := NewIntCmd(c, "bitcount", args...)
	cmd.execute()
	return cmd
}

//------------------------------------------------------------------------------
// Hash
func (c *Redis) HDel(key string, fields ...string) *IntCmd {
	args := make([]interface{}, 1+len(fields))
	args[0] = key
	for i, field := range fields {
		args[i+1] = field
	}
	cmd := NewIntCmd(c, "hdel", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) HExists(key, field string) *BoolCmd {
	cmd := NewBoolCmd(c, "hexists", key, field)
	cmd.execute()
	return cmd
}

func (c *Redis) HGet(key, field string) *StringCmd {
	cmd := NewStringCmd(c, "hget", key, field)
	cmd.execute()
	return cmd
}

func (c *Redis) HGetAll(key string) *StringStringMapCmd {
	cmd := NewStringStringMapCmd(c, "hgetall", key)
	cmd.execute()
	return cmd
}

func (c *Redis) HIncrBy(key, field string, incr int64) *IntCmd {
	cmd := NewIntCmd(c, "hincrby", key, field, incr)
	cmd.execute()
	return cmd
}

func (c *Redis) HIncrByFloat(key, field string, incr float64) *FloatCmd {
	cmd := NewFloatCmd(c, "hincrbyfloat", key, field, incr)
	cmd.execute()
	return cmd
}

func (c *Redis) HKeys(key string) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "hkeys", key)
	cmd.execute()
	return cmd
}

func (c *Redis) HLen(key string) *IntCmd {
	cmd := NewIntCmd(c, "hlen", key)
	cmd.execute()
	return cmd
}

// HMGet returns the values for the specified fields in the hash stored at key.
// It returns an interface{} to distinguish between empty string and nil value.
func (c *Redis) HMGet(key string, fields ...string) *SliceCmd {
	args := make([]interface{}, 1+len(fields))
	args[0] = key
	for i, field := range fields {
		args[i+1] = field
	}
	cmd := NewSliceCmd(c, "hmget", args...)
	cmd.execute()
	return cmd
}

// HSet accepts values in following formats:
//   - HSet("myhash", "key1", "value1", "key2", "value2")
//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
//
// Note that it requires Redis v4 for multiple field/value pairs support.
func (c *Redis) HSet(key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = appendArgs(args, values)
	cmd := NewIntCmd(c, "hset", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) HSetNX(key, field string, value interface{}) *BoolCmd {
	cmd := NewBoolCmd(c, "hsetnx", key, field, value)
	cmd.execute()
	return cmd
}

func (c *Redis) HVals(key string) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "hvals", key)
	cmd.execute()
	return cmd
}

//------------------------------------------------------------------------------
// List
func (c *Redis) LIndex(key string, index int64) *StringCmd {
	cmd := NewStringCmd(c, "lindex", key, index)
	cmd.execute()
	return cmd
}

func (c *Redis) LInsert(key, op string, pivot, value interface{}) *IntCmd {
	cmd := NewIntCmd(c, "linsert", key, op, pivot, value)
	cmd.execute()
	return cmd
}

func (c *Redis) LLen(key string) *IntCmd {
	cmd := NewIntCmd(c, "llen", key)
	cmd.execute()
	return cmd
}

func (c *Redis) LPop(key string) *StringCmd {
	cmd := NewStringCmd(c, "lpop", key)
	cmd.execute()
	return cmd
}

func (c *Redis) LPush(key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = appendArgs(args, values)
	cmd := NewIntCmd(c, "lpush", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) LPushX(key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = appendArgs(args, values)
	cmd := NewIntCmd(c, "lpushx", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) LRange(key string, start, stop int64) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "lrange", key, start, stop)
	cmd.execute()
	return cmd
}

func (c *Redis) LRem(key string, count int64, value interface{}) *IntCmd {
	cmd := NewIntCmd(c, "lrem", key, count, value)
	cmd.execute()
	return cmd
}

func (c *Redis) LSet(key string, index int64, value interface{}) *StatusCmd {
	cmd := NewStatusCmd(c, "lset", key, index, value)
	cmd.execute()
	return cmd
}

func (c *Redis) LTrim(key string, start, stop int64) *StatusCmd {
	cmd := NewStatusCmd(c, "ltrim", key, start, stop)
	cmd.execute()
	return cmd
}

func (c *Redis) RPop(key string) *StringCmd {
	cmd := NewStringCmd(c, "rpop", key)
	cmd.execute()
	return cmd
}

func (c *Redis) RPush(key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = append(args, values...)
	cmd := NewIntCmd(c, "rpush", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) RPushX(key string, values ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = append(args, values...)
	cmd := NewIntCmd(c, "rpushx", args...)
	cmd.execute()
	return cmd
}

//------------------------------------------------------------------------------
// Set

func (c *Redis) SAdd(key string, members ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(members))
	args[0] = key
	args = append(args, members...)
	cmd := NewIntCmd(c, "sadd", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) SCard(key string) *IntCmd {
	cmd := NewIntCmd(c, "scard", key)
	cmd.execute()
	return cmd
}

func (c *Redis) SDiff(keys ...string) *StringSliceCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewStringSliceCmd(c, "sdiff", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) SDiffStore(destination string, keys ...string) *IntCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = destination
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewIntCmd(c, "sdiffstore", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) SInter(keys ...string) *StringSliceCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewStringSliceCmd(c, "sinter", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) SInterStore(destination string, keys ...string) *IntCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = destination
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewIntCmd(c, "sinterstore", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) SIsMember(key string, member interface{}) *BoolCmd {
	cmd := NewBoolCmd(c, "sismember", key, member)
	cmd.execute()
	return cmd
}

// SMembers `SMEMBERS key` command output as a slice.
func (c *Redis) SMembers(key string) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "smembers", key)
	cmd.execute()
	return cmd
}

func (c *Redis) SMove(source, destination string, member interface{}) *BoolCmd {
	cmd := NewBoolCmd(c, "smove", source, destination, member)
	cmd.execute()
	return cmd
}

// SPop `SPOP key` command.
func (c *Redis) SPop(key string) *StringCmd {
	cmd := NewStringCmd(c, "spop", key)
	cmd.execute()
	return cmd
}

// SPOP `SPOP key count` command.
func (c *Redis) SPopN(key string, count int64) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "spop", key, count)
	cmd.execute()
	return cmd
}

// SRandMember `SRANDMEMBER key` command.
func (c *Redis) SRandMember(key string) *StringCmd {
	cmd := NewStringCmd(c, "srandmember", key)
	cmd.execute()
	return cmd
}

// SRandMemberN `SRANDMEMBER key count` command.
func (c *Redis) SRandMemberN(key string, count int64) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "srandmember", key, count)
	cmd.execute()
	return cmd
}

func (c *Redis) SRem(key string, members ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(members))
	args[0] = key
	args = append(args, members...)
	cmd := NewIntCmd(c, "srem", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) SUnion(keys ...string) *StringSliceCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewStringSliceCmd(c, "sunion", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) SUnionStore(destination string, keys ...string) *IntCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = destination
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewIntCmd(c, "sunionstore", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) executeZSetIntCmd(name string, key string, members ...*Z) *IntCmd {
	l := make([]interface{}, 2*len(members)+1)
	l[0] = key
	for i := range members {
		l[2*i+1] = members[i].Score
		l[2*i+2] = members[i].Member
	}

	cmd := NewIntCmd(c, name, l...)
	cmd.execute()
	return cmd
}

func (c *Redis) executeZSetFloatCmd(name string, key string, members ...*Z) *FloatCmd {
	l := make([]interface{}, 2*len(members)+1)
	l[0] = key
	for i := range members {
		l[i+1] = members[i].Score
		l[i+2] = members[i].Member
	}

	cmd := NewFloatCmd(c, name, l...)
	cmd.execute()
	return cmd
}

// ZAdd `ZADD key score member [score member ...]` command.
func (c *Redis) ZAdd(key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd("zadd", key, members...)
}

// ZAddNX `ZADD key NX score member [score member ...]` command.
func (c *Redis) ZAddNX(key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd("zaddnx", key, members...)
}

// ZAddXX `ZADD key XX score member [score member ...]` command.
func (c *Redis) ZAddXX(key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd("zaddxx", key, members...)
}

// ZAddCh `ZADD key CH score member [score member ...]` command.
func (c *Redis) ZAddCh(key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd("zaddch", key, members...)
}

// ZAddNXCh `ZADD key NX CH score member [score member ...]` command.
func (c *Redis) ZAddNXCh(key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd("zaddnxch", key, members...)
}

// ZAddXXCh `ZADD key XX CH score member [score member ...]` command.
func (c *Redis) ZAddXXCh(key string, members ...*Z) *IntCmd {
	return c.executeZSetIntCmd("zaddxxch", key, members...)
}

// ZIncr `ZADD key INCR score member` command.
func (c *Redis) ZIncr(key string, member *Z) *FloatCmd {
	return c.executeZSetFloatCmd("zincr", key, member)
}

// ZIncrNX `ZADD key NX INCR score member` command.
func (c *Redis) ZIncrNX(key string, member *Z) *FloatCmd {
	return c.executeZSetFloatCmd("zincrnx", key, member)
}

// ZIncrXX `ZADD key XX INCR score member` command.
func (c *Redis) ZIncrXX(key string, member *Z) *FloatCmd {
	return c.executeZSetFloatCmd("zincrxx", key, member)
}

func (c *Redis) ZCard(key string) *IntCmd {
	cmd := NewIntCmd(c, "zcard", key)
	cmd.execute()
	return cmd
}

func (c *Redis) ZCount(key, min, max string) *IntCmd {
	cmd := NewIntCmd(c, "zcount", key, min, max)
	cmd.execute()
	return cmd
}

func (c *Redis) ZIncrBy(key string, increment float64, member string) *FloatCmd {
	cmd := NewFloatCmd(c, "zincrby", key, increment, member)
	cmd.execute()
	return cmd
}

func (c *Redis) ZInterStore(destination string, store *ZStore) *IntCmd {
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
	cmd.execute()
	return cmd
}

func (c *Redis) zRange(key string, start, stop int64, withScores bool) *StringSliceCmd {
	args := []interface{}{key, start, stop}
	if withScores {
		args = append(args, "withscores")
	}
	cmd := NewStringSliceCmd(c, "zrange", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) ZRange(key string, start, stop int64) *StringSliceCmd {
	return c.zRange(key, start, stop, false)
}

func (c *Redis) ZRangeWithScores(key string, start, stop int64) *ZSliceCmd {
	cmd := NewZSliceCmd(c, "zrange", key, start, stop, "withscores")
	cmd.execute()
	return cmd
}

func (c *Redis) zRangeBy(zcmd string, key string, opt *ZRangeBy, withScores bool) *StringSliceCmd {
	args := []interface{}{key, opt.Min, opt.Max}
	if withScores {
		args = append(args, "withscores")
	}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, "limit", opt.Offset, opt.Count)
	}
	cmd := NewStringSliceCmd(c, zcmd, args...)
	cmd.execute()
	return cmd
}

func (c *Redis) ZRangeByScore(key string, opt *ZRangeBy) *StringSliceCmd {
	return c.zRangeBy("zrangebyscore", key, opt, false)
}

func (c *Redis) ZRangeByScoreWithScores(key string, opt *ZRangeBy) *ZSliceCmd {
	args := []interface{}{key, opt.Min, opt.Max, "withscores"}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, "limit", opt.Offset, opt.Count)
	}
	cmd := NewZSliceCmd(c, "zrangebyscore", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) ZRank(key, member string) *IntCmd {
	cmd := NewIntCmd(c, "zrank", key, member)
	cmd.execute()
	return cmd
}

func (c *Redis) ZRem(key string, members ...interface{}) *IntCmd {
	args := make([]interface{}, 2, 2+len(members))
	args[0] = key
	args = appendArgs(args, members)
	cmd := NewIntCmd(c, "zrem", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) ZRemRangeByRank(key string, start, stop int64) *IntCmd {
	cmd := NewIntCmd(c, "zremrangebyrank", key, start, stop)
	cmd.execute()
	return cmd
}

func (c *Redis) ZRemRangeByScore(key, min, max string) *IntCmd {
	cmd := NewIntCmd(c, "zremrangebyscore", key, min, max)
	cmd.execute()
	return cmd
}

func (c *Redis) ZRevRange(key string, start, stop int64) *StringSliceCmd {
	cmd := NewStringSliceCmd(c, "zrevrange", key, start, stop)
	cmd.execute()
	return cmd
}

func (c *Redis) ZRevRangeWithScores(key string, start, stop int64) *ZSliceCmd {
	cmd := NewZSliceCmd(c, "zrevrange", key, start, stop, "withscores")
	cmd.execute()
	return cmd
}

func (c *Redis) zRevRangeBy(zcmd, key string, opt *ZRangeBy) *StringSliceCmd {
	args := []interface{}{key, opt.Max, opt.Min}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, "limit", opt.Offset, opt.Count)
	}
	cmd := NewStringSliceCmd(c, zcmd, args...)
	cmd.execute()
	return cmd
}

func (c *Redis) ZRevRangeByScore(key string, opt *ZRangeBy) *StringSliceCmd {
	return c.zRevRangeBy("zrevrangebyscore", key, opt)
}

func (c *Redis) ZRevRangeByScoreWithScores(key string, opt *ZRangeBy) *ZSliceCmd {
	args := []interface{}{key, opt.Min, opt.Max, "withscores"}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, "limit", opt.Offset, opt.Count)
	}
	cmd := NewZSliceCmd(c, "zrevrangebyscore", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) ZRevRank(key, member string) *IntCmd {
	cmd := NewIntCmd(c, "zrevrank", key, member)
	cmd.execute()
	return cmd
}

func (c *Redis) ZScore(key, member string) *FloatCmd {
	cmd := NewFloatCmd(c, "zscore", key, member)
	cmd.execute()
	return cmd
}

func (c *Redis) ZUnionStore(dest string, store *ZStore) *IntCmd {
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
	cmd.execute()
	return cmd
}

//------------------------------------------------------------------------------
// HyperLogLog

func (c *Redis) PFAdd(key string, els ...interface{}) *IntCmd {
	args := make([]interface{}, 1, 1+len(els))
	args[0] = key
	args = appendArgs(args, els)
	cmd := NewIntCmd(c, "pfadd", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) PFCount(keys ...string) *IntCmd {
	args := make([]interface{}, len(keys))
	for i, key := range keys {
		args[i] = key
	}
	cmd := NewIntCmd(c, "pfcount", args...)
	cmd.execute()
	return cmd
}

func (c *Redis) PFMerge(dest string, keys ...string) *StatusCmd {
	args := make([]interface{}, 1+len(keys))
	args[0] = dest
	for i, key := range keys {
		args[1+i] = key
	}
	cmd := NewStatusCmd(c, "pfmerge", args...)
	cmd.execute()
	return cmd
}
