package redis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var redisCli *Redis

func Init() {
	redisCli = &Redis{}
}

func TestMain(m *testing.M) {
	Init()
	m.Run()
}

func TestGetNil(t *testing.T) {
	cmd := redisCli.Get("NilKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, Nil, e)
}

func Test_AA_Set(t *testing.T) {
	cmd := redisCli.Set("TestKey", "Hello", 300*time.Second)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "OK", v)
	assert.Empty(t, e)
}

func Test_AB_Expire(t *testing.T) {
	cmd := redisCli.Expire("TestKey", 6*60*time.Minute)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AB_ExpireAt(t *testing.T) {
	cmd := redisCli.ExpireAt("TestKey", time.Now().Add(6*60*time.Minute))
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AB_PExpire(t *testing.T) {
	cmd := redisCli.PExpire("TestKey", 6*60*time.Minute)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AB_PExpireAt(t *testing.T) {
	cmd := redisCli.PExpireAt("TestKey", time.Now().Add(6*60*time.Minute))
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AB_Persist(t *testing.T) {
	cmd := redisCli.Persist("TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AC_TTL(t *testing.T) {
	cmd := redisCli.TTL("TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v s, Err: %v\n", v.Milliseconds()/1000, e)
	assert.Equal(t, true, v.Milliseconds()/1000 > 0)
	assert.Empty(t, e)
}

func Test_AC_PTTL(t *testing.T) {
	cmd := redisCli.PTTL("TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v ms, Err: %v\n", v.Milliseconds(), e)
	assert.Equal(t, true, v.Milliseconds() > 0)
	assert.Empty(t, e)
}

func Test_AD_Type(t *testing.T) {
	cmd := redisCli.Type("TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "string", v)
	assert.Empty(t, e)
}

func Test_AE_Append(t *testing.T) {
	cmd := redisCli.Append("TestKey", " World")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, 11, v)
	assert.Empty(t, e)
}

func Test_AF_Get(t *testing.T) {
	cmd := redisCli.Get("TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "Hello World", v)
	assert.Empty(t, e)
}

func Test_AF_GetRange(t *testing.T) {
	cmd := redisCli.GetRange("TestKey", 0, 5)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "Hello", v)
	assert.Empty(t, e)
}

func Test_AG_GetSet(t *testing.T) {
	cmd := redisCli.GetSet("TestKey", "Hello")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v != "")
	assert.Empty(t, e)
}

func Test_AG_Exists(t *testing.T) {
	cmd := redisCli.Exists("TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, int64(1), v)
	assert.Empty(t, e)
}

func Test_AH_SetNX(t *testing.T) {
	cmd := redisCli.SetNX("TestKey", "Value")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, false, v)
	assert.Empty(t, e)
}

func Test_AI_SetXX(t *testing.T) {
	cmd := redisCli.SetXX("TestKey", "Value")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_AJ_SetRange(t *testing.T) {
	cmd := redisCli.SetRange("TestKey", 0, "Hello")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, int64(len("Hello")), v)
	assert.Empty(t, e)
}

func Test_AK_StrLen(t *testing.T) {
	cmd := redisCli.StrLen("TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_AL_Del(t *testing.T) {
	cmd := redisCli.Del("TestKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, int64(1), v)
	assert.Empty(t, e)
}

func Test_BA_MSet(t *testing.T) {
	cmd := redisCli.MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "OK", v)
	assert.Empty(t, e)
}

func Test_BB_MGet(t *testing.T) {
	cmd := redisCli.MGet("key1", "key2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) == 2)
	assert.Empty(t, e)
}

func Test_CA_Incr(t *testing.T) {
	cmd := redisCli.Incr("IntKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_CB_IncrBy(t *testing.T) {
	cmd := redisCli.IncrBy("IntKey", 10)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_CB_IncrByFloat(t *testing.T) {
	cmd := redisCli.IncrByFloat("IntKey", 4)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_CC_Decr(t *testing.T) {
	cmd := redisCli.Decr("IntKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_CD_DecrBy(t *testing.T) {
	cmd := redisCli.DecrBy("IntKey", 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_DA_SetBit(t *testing.T) {
	cmd := redisCli.SetBit("BitKey", 10, 1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_DB_GetBit(t *testing.T) {
	cmd := redisCli.GetBit("BitKey", 10)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_DC_BitCount(t *testing.T) {
	cmd := redisCli.BitCount("BitKey", &BitCountArgs{Start: 0, End: 2})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_EA_HSet(t *testing.T) {
	cmd := redisCli.HSet("HashKey", map[string]interface{}{"field1": "value1", "field2": "value2", "field3": 1})
	redisCli.Expire("HashKey", 6*60*time.Minute)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_EB_HGet(t *testing.T) {
	cmd := redisCli.HGet("HashKey", "field1")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "value1", v)
	assert.Empty(t, e)
}

func Test_EB_HMGet(t *testing.T) {
	cmd := redisCli.HMGet("HashKey", "field1", "field2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_EB_HKeys(t *testing.T) {
	cmd := redisCli.HKeys("HashKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_EB_HVals(t *testing.T) {
	cmd := redisCli.HVals("HashKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_EB_HLen(t *testing.T) {
	cmd := redisCli.HLen("HashKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_EB_HGetAll(t *testing.T) {
	cmd := redisCli.HGetAll("HashKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_EC_HIncrBy(t *testing.T) {
	cmd := redisCli.HIncrBy("HashKey", "field3", 1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_EC_HIncrByFloat(t *testing.T) {
	cmd := redisCli.HIncrByFloat("HashKey", "field3", 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_ED_HExists(t *testing.T) {
	cmd := redisCli.HExists("HashKey", "field1")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_ED_HSetNX(t *testing.T) {
	cmd := redisCli.HSetNX("HashKey", "field1", "value1")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, false, v)
	assert.Empty(t, e)
}

func Test_EE_HDel(t *testing.T) {
	cmd := redisCli.HDel("HashKey", "field2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, int64(1), v)
	assert.Empty(t, e)
}

func Test_FA_LPushX(t *testing.T) {
	cmd := redisCli.LPushX("NotExistKey", "v1")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v == 0)
	assert.Empty(t, e)
}

func Test_FA_RPushX(t *testing.T) {
	cmd := redisCli.RPushX("NotExistKey", "v2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v == 0)
	assert.Empty(t, e)
}

func Test_FA_LPush(t *testing.T) {
	cmd := redisCli.LPush("ListKey", "v1", "V2", "v100")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_FA_RPush(t *testing.T) {
	cmd := redisCli.RPush("ListKey", "v3", "v4")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_FB_LInsert(t *testing.T) {
	cmd := redisCli.LInsert("ListKey", "BEFORE", "v3", "vx")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_FB_LTrim(t *testing.T) {
	cmd := redisCli.LTrim("ListKey", 1, -1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "OK", v)
	assert.Empty(t, e)
}

func Test_FB_LSet(t *testing.T) {
	cmd := redisCli.LSet("ListKey", 0, "v2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "OK", v)
	assert.Empty(t, e)
}

func Test_FB_LIndex(t *testing.T) {
	cmd := redisCli.LIndex("ListKey", 0)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_FB_LLen(t *testing.T) {
	cmd := redisCli.LLen("ListKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_FC_LRem(t *testing.T) {
	cmd := redisCli.LRem("ListKey", 0, "vx")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_FC_LPop(t *testing.T) {
	cmd := redisCli.LPop("ListKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_FC_RPop(t *testing.T) {
	cmd := redisCli.RPop("ListKey")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_FD_LRange(t *testing.T) {
	cmd := redisCli.LRange("ListKey", 0, -1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GA_SAdd(t *testing.T) {
	cmd := redisCli.SAdd("{Set}Key", 1, 2, 3, 4, 5)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GB_SIsMember(t *testing.T) {
	cmd := redisCli.SIsMember("{Set}Key", 1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_GB_SMembers(t *testing.T) {
	cmd := redisCli.SMembers("{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GB_SDiff(t *testing.T) {
	cmd := redisCli.SDiff("{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GB_SDiffStore(t *testing.T) {
	cmd := redisCli.SDiffStore("{Set}Key_2", "{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GB_SInter(t *testing.T) {
	cmd := redisCli.SInter("{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GB_SInterStore(t *testing.T) {
	cmd := redisCli.SInterStore("{Set}Key_3", "{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GB_SUnion(t *testing.T) {
	cmd := redisCli.SUnion("{Set}Key", "{Set}Key_4")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GB_SUnionStore(t *testing.T) {
	cmd := redisCli.SUnionStore("{Set}Key_5", "{Set}Key", "{Set}Key_4")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GB_SCard(t *testing.T) {
	cmd := redisCli.SCard("{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GD_SMove(t *testing.T) {
	cmd := redisCli.SMove("{Set}Key", "{Set}Key_6", 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v)
	assert.Empty(t, e)
}

func Test_GE_SRem(t *testing.T) {
	cmd := redisCli.SRem("{Set}Key", 3)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_GE_SPop(t *testing.T) {
	cmd := redisCli.SPop("{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GE_SPopN(t *testing.T) {
	cmd := redisCli.SPopN("{Set}Key", 1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GE_SRandMember(t *testing.T) {
	cmd := redisCli.SRandMember("{Set}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_GE_SRandMemberN(t *testing.T) {
	cmd := redisCli.SRandMemberN("{Set}Key", 1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, len(v) > 0)
	assert.Empty(t, e)
}

func Test_HA_ZAdd(t *testing.T) {
	cmd := redisCli.ZAdd("{ZSet}Key", &Z{Score: 1, Member: "V1"}, &Z{Score: 2, Member: "V2"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HA_ZAddNX(t *testing.T) {
	cmd := redisCli.ZAddNX("{ZSet}Key", &Z{Score: 3, Member: "V3"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HA_ZAddXX(t *testing.T) {
	cmd := redisCli.ZAddNX("{ZSet}Key", &Z{Score: 4, Member: "V4"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HA_ZAddCh(t *testing.T) {
	cmd := redisCli.ZAddCh("{ZSet}Key", &Z{Score: 5, Member: "V5"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HA_ZAddNXCh(t *testing.T) {
	cmd := redisCli.ZAddNXCh("{ZSet}Key", &Z{Score: 6, Member: "V6"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HA_ZAddXXCh(t *testing.T) {
	cmd := redisCli.ZAddXXCh("{ZSet}Key", &Z{Score: 7, Member: "V7"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HB_ZIncr(t *testing.T) {
	cmd := redisCli.ZIncr("{ZSet}Key", &Z{Score: 5, Member: "V5"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HB_ZIncrNX(t *testing.T) {
	cmd := redisCli.ZIncrNX("{ZSet}Key", &Z{Score: 5, Member: "V8"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HB_ZIncrXX(t *testing.T) {
	cmd := redisCli.ZIncrXX("{ZSet}Key", &Z{Score: 5, Member: "V5"})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HC_ZCard(t *testing.T) {
	cmd := redisCli.ZCard("{ZSet}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HC_ZCount(t *testing.T) {
	cmd := redisCli.ZCount("{ZSet}Key", "1", "5")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HD_ZInterStore(t *testing.T) {
	cmd := redisCli.ZInterStore("{ZSet}Key_4", &ZStore{
		Keys:      []string{"a", "b"},
		Weights:   []float64{1, 2},
		Aggregate: "",
	})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HD_ZRange(t *testing.T) {
	cmd := redisCli.ZRange("{ZSet}Key", 0, -1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HD_ZRangeWithScores(t *testing.T) {
	cmd := redisCli.ZRangeWithScores("{ZSet}Key", 0, -1)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HD_ZRangeByScore(t *testing.T) {
	cmd := redisCli.ZRangeByScore("{ZSet}Key", &ZRangeBy{
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
	cmd := redisCli.ZRangeByScoreWithScores("{ZSet}Key", &ZRangeBy{
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
	cmd := redisCli.ZRank("{ZSet}Key", "V1")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRem(t *testing.T) {
	cmd := redisCli.ZRem("{ZSet}Key", "V2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRemRangeByRank(t *testing.T) {
	cmd := redisCli.ZRemRangeByRank("{ZSet}Key", 1, 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRemRangeByScore(t *testing.T) {
	cmd := redisCli.ZRemRangeByScore("{ZSet}Key", "1", "2")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRevRange(t *testing.T) {
	cmd := redisCli.ZRevRange("{ZSet}Key", 1, 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRevRangeWithScores(t *testing.T) {
	cmd := redisCli.ZRevRangeWithScores("{ZSet}Key", 1, 2)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HE_ZRevRangeByScore(t *testing.T) {
	cmd := redisCli.ZRevRangeByScore("{ZSet}Key", &ZRangeBy{
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
	cmd := redisCli.ZRevRangeByScoreWithScores("{ZSet}Key", &ZRangeBy{
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
	cmd := redisCli.ZRevRank("{ZSet}Key", "V5")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HF_ZScore(t *testing.T) {
	cmd := redisCli.ZScore("{ZSet}Key", "V6")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_HF_ZUnionStore(t *testing.T) {
	cmd := redisCli.ZUnionStore("{ZSet}Key_6", &ZStore{
		Keys:      []string{"a", "b"},
		Weights:   []float64{1, 2},
		Aggregate: "",
	})
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Empty(t, e)
}

func Test_IA_PFAdd(t *testing.T) {
	cmd := redisCli.PFAdd("{PF}Key", 1, 2, 2, 3, 3, 4, 4, 5)
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_IB_PFCount(t *testing.T) {
	cmd := redisCli.PFCount("{PF}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, true, v > 0)
	assert.Empty(t, e)
}

func Test_IC_PFMerge(t *testing.T) {
	cmd := redisCli.PFMerge("{PF}Key_1", "{PF}Key")
	v, e := cmd.Result()
	fmt.Printf("Res: %v, Err: %v\n", v, e)
	assert.Equal(t, "OK", v)
	assert.Empty(t, e)
}
