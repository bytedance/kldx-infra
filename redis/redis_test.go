package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	redisCli *Redis
	ctx = context.Background()
)

func Init() {
	redisCli = &Redis{}
}

func TestMain(m *testing.M) {
	Init()
	m.Run()
}

func TestGetNil(t *testing.T) {
	cmd := redisCli.Get(ctx,"NilKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, Nil, e)
}

func Test_AA_Set(t *testing.T) {
	cmd := redisCli.Set(ctx, "TestKey", "Hello", 300*time.Second)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "OK", v)
	assert.Empty(t, e)
}

func Test_AB_Expire(t *testing.T) {
	cmd := redisCli.Expire(ctx, "TestKey", 6*60*time.Minute)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AB_ExpireAt(t *testing.T) {
	cmd := redisCli.ExpireAt(ctx, "TestKey", time.Now().Add(6*60*time.Minute))
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AB_PExpire(t *testing.T) {
	cmd := redisCli.PExpire(ctx, "TestKey", 6*60*time.Minute)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AB_PExpireAt(t *testing.T) {
	cmd := redisCli.PExpireAt(ctx, "TestKey", time.Now().Add(6*60*time.Minute))
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AB_Persist(t *testing.T) {
	cmd := redisCli.Persist(ctx, "TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AC_TTL(t *testing.T) {
	cmd := redisCli.TTL(ctx, "TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v s, Err: %v\n", v.Milliseconds()/1000, e)
	assert.Equal(t, true, v.Milliseconds()/1000 > 0)
	assert.Empty(t, e)
}

func Test_AC_PTTL(t *testing.T) {
	cmd := redisCli.PTTL(ctx, "TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v ms, Err: %v\n", v.Milliseconds(), e)
	assert.Equal(t, true, v.Milliseconds() > 0)
	assert.Empty(t, e)
}

func Test_AD_Type(t *testing.T) {
	cmd := redisCli.Type(ctx, "TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "string", v)
	assert.Empty(t, e)
}

func Test_AE_Append(t *testing.T) {
	cmd := redisCli.Append(ctx, "TestKey", " World")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, 11, v)
	assert.Empty(t, e)
}

func Test_AF_Get(t *testing.T) {
	cmd := redisCli.Get(ctx, "TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "Hello World", v)
	assert.Empty(t, e)
}

func Test_AF_GetRange(t *testing.T) {
	cmd := redisCli.GetRange(ctx, "TestKey", 0, 5)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "Hello", v)
	assert.Empty(t, e)
}

func Test_AG_GetSet(t *testing.T) {
	cmd := redisCli.GetSet(ctx, "TestKey", "Hello")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v != "")
	assert.Empty(t, e)
}

func Test_AG_Exists(t *testing.T) {
	cmd := redisCli.Exists(ctx, "TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, int64(1), v)
	assert.Empty(t, e)
}

func Test_AH_SetNX(t *testing.T) {
	cmd := redisCli.SetNX(ctx, "TestKey", "Value")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, false, v)
	assert.Empty(t, e)
}

func Test_AI_SetXX(t *testing.T) {
	cmd := redisCli.SetXX(ctx, "TestKey", "Value")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AJ_SetRange(t *testing.T) {
	cmd := redisCli.SetRange(ctx, "TestKey", 0, "Hello")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, int64(len("Hello")), v)
	assert.Empty(t, e)
}

func Test_AK_StrLen(t *testing.T) {
	cmd := redisCli.StrLen(ctx, "TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_AL_Del(t *testing.T) {
	cmd := redisCli.Del(ctx, "TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, int64(1), v)
	assert.Empty(t, e)
}

func Test_BA_MSet(t *testing.T) {
	cmd := redisCli.MSet(ctx, map[string]interface{}{"key1": "value1", "key2": "value2"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "OK", v)
	assert.Empty(t, e)
}

func Test_BB_MGet(t *testing.T) {
	cmd := redisCli.MGet(ctx, "key1", "key2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) == 2)
	assert.Empty(t, e)
}

func Test_CA_Incr(t *testing.T) {
	cmd := redisCli.Incr(ctx, "IntKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_CB_IncrBy(t *testing.T) {
	cmd := redisCli.IncrBy(ctx, "IntKey", 10)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_CB_IncrByFloat(t *testing.T) {
	cmd := redisCli.IncrByFloat(ctx, "IntKey", 4)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_CC_Decr(t *testing.T) {
	cmd := redisCli.Decr(ctx, "IntKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_CD_DecrBy(t *testing.T) {
	cmd := redisCli.DecrBy(ctx, "IntKey", 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_DA_SetBit(t *testing.T) {
	cmd := redisCli.SetBit(ctx, "BitKey", 10, 1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_DB_GetBit(t *testing.T) {
	cmd := redisCli.GetBit(ctx, "BitKey", 10)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_DC_BitCount(t *testing.T) {
	cmd := redisCli.BitCount(ctx, "BitKey", &BitCountArgs{Start: 0, End: 2})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_EA_HSet(t *testing.T) {
	cmd := redisCli.HSet(ctx, "HashKey", map[string]interface{}{"field1": "value1", "field2": "value2", "field3": 1})
	redisCli.Expire(ctx, "HashKey", 6*60*time.Minute)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_EB_HGet(t *testing.T) {
	cmd := redisCli.HGet(ctx, "HashKey", "field1")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "value1", v)
	assert.Empty(t, e)
}

func Test_EB_HMGet(t *testing.T) {
	cmd := redisCli.HMGet(ctx, "HashKey", "field1", "field2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_EB_HKeys(t *testing.T) {
	cmd := redisCli.HKeys(ctx, "HashKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_EB_HVals(t *testing.T) {
	cmd := redisCli.HVals(ctx, "HashKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_EB_HLen(t *testing.T) {
	cmd := redisCli.HLen(ctx, "HashKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_EB_HGetAll(t *testing.T) {
	cmd := redisCli.HGetAll(ctx, "HashKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_EC_HIncrBy(t *testing.T) {
	cmd := redisCli.HIncrBy(ctx, "HashKey", "field3", 1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_EC_HIncrByFloat(t *testing.T) {
	cmd := redisCli.HIncrByFloat(ctx, "HashKey", "field3", 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_ED_HExists(t *testing.T) {
	cmd := redisCli.HExists(ctx, "HashKey", "field1")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_ED_HSetNX(t *testing.T) {
	cmd := redisCli.HSetNX(ctx, "HashKey", "field1", "value1")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, false, v)
	assert.Empty(t, e)
}

func Test_EE_HDel(t *testing.T) {
	cmd := redisCli.HDel(ctx, "HashKey", "field2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, int64(1), v)
	assert.Empty(t, e)
}

func Test_FA_LPushX(t *testing.T) {
	cmd := redisCli.LPushX(ctx, "NotExistKey", "v1")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v == 0)
	assert.Empty(t, e)
}

func Test_FA_RPushX(t *testing.T) {
	cmd := redisCli.RPushX(ctx, "NotExistKey", "v2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v == 0)
	assert.Empty(t, e)
}

func Test_FA_LPush(t *testing.T) {
	cmd := redisCli.LPush(ctx, "ListKey", "v1", "V2", "v100")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_FA_RPush(t *testing.T) {
	cmd := redisCli.RPush(ctx, "ListKey", "v3", "v4")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_FB_LInsert(t *testing.T) {
	cmd := redisCli.LInsert(ctx, "ListKey", "BEFORE", "v3", "vx")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_FB_LTrim(t *testing.T) {
	cmd := redisCli.LTrim(ctx, "ListKey", 1, -1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "OK", v)
	assert.Empty(t, e)
}

func Test_FB_LSet(t *testing.T) {
	cmd := redisCli.LSet(ctx, "ListKey", 0, "v2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "OK", v)
	assert.Empty(t, e)
}

func Test_FB_LIndex(t *testing.T) {
	cmd := redisCli.LIndex(ctx, "ListKey", 0)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_FB_LLen(t *testing.T) {
	cmd := redisCli.LLen(ctx, "ListKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_FC_LRem(t *testing.T) {
	cmd := redisCli.LRem(ctx, "ListKey", 0, "vx")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_FC_LPop(t *testing.T) {
	cmd := redisCli.LPop(ctx, "ListKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_FC_RPop(t *testing.T) {
	cmd := redisCli.RPop(ctx, "ListKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_FD_LRange(t *testing.T) {
	cmd := redisCli.LRange(ctx, "ListKey", 0, -1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GA_SAdd(t *testing.T) {
	cmd := redisCli.SAdd(ctx, "{Set}Key", 1, 2, 3, 4, 5)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GB_SIsMember(t *testing.T) {
	cmd := redisCli.SIsMember(ctx, "{Set}Key", 1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_GB_SMembers(t *testing.T) {
	cmd := redisCli.SMembers(ctx, "{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GB_SDiff(t *testing.T) {
	cmd := redisCli.SDiff(ctx, "{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GB_SDiffStore(t *testing.T) {
	cmd := redisCli.SDiffStore("{Set}Key_2", ctx, "{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GB_SInter(t *testing.T) {
	cmd := redisCli.SInter(ctx, "{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GB_SInterStore(t *testing.T) {
	cmd := redisCli.SInterStore("{Set}Key_3", ctx, "{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GB_SUnion(t *testing.T) {
	cmd := redisCli.SUnion(ctx, "{Set}Key", "{Set}Key_4")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GB_SUnionStore(t *testing.T) {
	cmd := redisCli.SUnionStore("{Set}Key_5", ctx, "{Set}Key", "{Set}Key_4")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GB_SCard(t *testing.T) {
	cmd := redisCli.SCard(ctx, "{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GD_SMove(t *testing.T) {
	cmd := redisCli.SMove(ctx, "{Set}Key", "{Set}Key_6", 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_GE_SRem(t *testing.T) {
	cmd := redisCli.SRem(ctx, "{Set}Key", 3)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GE_SPop(t *testing.T) {
	cmd := redisCli.SPop(ctx, "{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GE_SPopN(t *testing.T) {
	cmd := redisCli.SPopN(ctx, "{Set}Key", 1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GE_SRandMember(t *testing.T) {
	cmd := redisCli.SRandMember(ctx, "{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GE_SRandMemberN(t *testing.T) {
	cmd := redisCli.SRandMemberN(ctx, "{Set}Key", 1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_HA_ZAdd(t *testing.T) {
	cmd := redisCli.ZAdd(ctx, "{ZSet}Key", &Z{Score: 1, Member: "V1"}, &Z{Score: 2, Member: "V2"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HA_ZAddNX(t *testing.T) {
	cmd := redisCli.ZAddNX(ctx, "{ZSet}Key", &Z{Score: 3, Member: "V3"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HA_ZAddXX(t *testing.T) {
	cmd := redisCli.ZAddNX(ctx, "{ZSet}Key", &Z{Score: 4, Member: "V4"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HA_ZAddCh(t *testing.T) {
	cmd := redisCli.ZAddCh(ctx, "{ZSet}Key", &Z{Score: 5, Member: "V5"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HA_ZAddNXCh(t *testing.T) {
	cmd := redisCli.ZAddNXCh(ctx, "{ZSet}Key", &Z{Score: 6, Member: "V6"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HA_ZAddXXCh(t *testing.T) {
	cmd := redisCli.ZAddXXCh(ctx, "{ZSet}Key", &Z{Score: 7, Member: "V7"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HB_ZIncr(t *testing.T) {
	cmd := redisCli.ZIncr(ctx, "{ZSet}Key", &Z{Score: 5, Member: "V5"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HB_ZIncrNX(t *testing.T) {
	cmd := redisCli.ZIncrNX(ctx, "{ZSet}Key", &Z{Score: 5, Member: "V8"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HB_ZIncrXX(t *testing.T) {
	cmd := redisCli.ZIncrXX(ctx, "{ZSet}Key", &Z{Score: 5, Member: "V5"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HC_ZCard(t *testing.T) {
	cmd := redisCli.ZCard(ctx, "{ZSet}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HC_ZCount(t *testing.T) {
	cmd := redisCli.ZCount(ctx, "{ZSet}Key", "1", "5")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HD_ZInterStore(t *testing.T) {
	cmd := redisCli.ZInterStore(ctx, "{ZSet}Key_4", &ZStore{
		Keys:      []string{"a", "b"},
		Weights:   []float64{1, 2},
		Aggregate: "",
	})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HD_ZRange(t *testing.T) {
	cmd := redisCli.ZRange(ctx, "{ZSet}Key", 0, -1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HD_ZRangeWithScores(t *testing.T) {
	cmd := redisCli.ZRangeWithScores(ctx, "{ZSet}Key", 0, -1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HD_ZRangeByScore(t *testing.T) {
	cmd := redisCli.ZRangeByScore(ctx, "{ZSet}Key", &ZRangeBy{
		Min:    "0",
		Max:    "10",
		Offset: 1,
		Count:  2,
	})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HD_ZRangeByScoreWithScores(t *testing.T) {
	cmd := redisCli.ZRangeByScoreWithScores(ctx, "{ZSet}Key", &ZRangeBy{
		Min:    "0",
		Max:    "10",
		Offset: 1,
		Count:  2,
	})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HD_ZRank(t *testing.T) {
	cmd := redisCli.ZRank(ctx, "{ZSet}Key", "V1")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRem(t *testing.T) {
	cmd := redisCli.ZRem(ctx, "{ZSet}Key", "V2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRemRangeByRank(t *testing.T) {
	cmd := redisCli.ZRemRangeByRank(ctx, "{ZSet}Key", 1, 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRemRangeByScore(t *testing.T) {
	cmd := redisCli.ZRemRangeByScore(ctx, "{ZSet}Key", "1", "2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRevRange(t *testing.T) {
	cmd := redisCli.ZRevRange(ctx, "{ZSet}Key", 1, 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRevRangeWithScores(t *testing.T) {
	cmd := redisCli.ZRevRangeWithScores(ctx, "{ZSet}Key", 1, 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRevRangeByScore(t *testing.T) {
	cmd := redisCli.ZRevRangeByScore(ctx, "{ZSet}Key", &ZRangeBy{
		Min:    "0",
		Max:    "10",
		Offset: 1,
		Count:  2,
	})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRevRangeByScoreWithScores(t *testing.T) {
	cmd := redisCli.ZRevRangeByScoreWithScores(ctx, "{ZSet}Key", &ZRangeBy{
		Min:    "0",
		Max:    "10",
		Offset: 1,
		Count:  2,
	})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRevRank(t *testing.T) {
	cmd := redisCli.ZRevRank(ctx, "{ZSet}Key", "V5")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HF_ZScore(t *testing.T) {
	cmd := redisCli.ZScore(ctx, "{ZSet}Key", "V6")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HF_ZUnionStore(t *testing.T) {
	cmd := redisCli.ZUnionStore(ctx, "{ZSet}Key_6", &ZStore{
		Keys:      []string{"a", "b"},
		Weights:   []float64{1, 2},
		Aggregate: "",
	})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_IA_PFAdd(t *testing.T) {
	cmd := redisCli.PFAdd(ctx, "{PF}Key", 1, 2, 2, 3, 3, 4, 4, 5)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_IB_PFCount(t *testing.T) {
	cmd := redisCli.PFCount(ctx, "{PF}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_IC_PFMerge(t *testing.T) {
	cmd := redisCli.PFMerge(ctx, "{PF}Key_1", "{PF}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "OK", v)
	assert.Empty(t, e)
}
