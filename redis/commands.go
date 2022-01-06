package redis

import (
	cExceptions "github.com/bytedance/kldx-common/exceptions"
	cHttp "github.com/bytedance/kldx-common/http"
	"github.com/bytedance/kldx-infra/http"
	"encoding/json"
	"strconv"
	"time"
)

const Nil = ErrorRedis("redis: nil")

type ErrorRedis string

func (e ErrorRedis) Error() string { return string(e) }

type result struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *result) bind(v interface{}) {
	r.Data = v
}

type baseCmd struct {
	client *Redis
	name   string
	args   []interface{}
	err    error
	result *result
}

func (c *baseCmd) Err() error {
	return c.err
}

type redisArgumentList struct {
	Cmd  string        `json:"cmd"`
	Args []interface{} `json:"args"`
}

// RedisCmdExecution Request
func (c *baseCmd) execute() {
	data, e := http.GetFaaSInfraClient().PostJson(http.GetFaaSInfraPathRedis(), nil, redisArgumentList{Cmd: c.name, Args: c.args}, cHttp.AppTokenMiddleware, http.FaaSInfraMiddleware)
	if e != nil {
		c.err = cExceptions.ErrorWrap(e)
		return
	}

	if e := json.Unmarshal(data, c.result); e != nil {
		c.err = cExceptions.InternalError("JsonUnmarshal failed, err: %s", e)
		return
	}

	if http.HasError(c.result.Code) {
		if http.IsSysError(c.result.Code) {
			c.err = cExceptions.InternalError("Request failed, code: %s, msg: %s", c.result.Code, c.result.Msg)
		} else {
			c.err = cExceptions.InvalidParamError("Request failed, code: %s, msg: %s", c.result.Code, c.result.Msg)
		}
		return
	}

	if c.result.Data == nil {
		c.err = Nil
		return
	}
}

type StringCmd struct {
	baseCmd
	val string
}

func NewStringCmd(client *Redis, name string, args ...interface{}) *StringCmd {
	cmd := &StringCmd{
		baseCmd: baseCmd{
			client: client,
			name:   name,
			args:   args,
			result: &result{},
		},
	}
	cmd.result.bind(&cmd.val)
	return cmd
}

func (c *StringCmd) Val() string {
	return c.val
}

func (c *StringCmd) Result() (string, error) {
	return c.val, c.err
}

func (c *StringCmd) Int() (int, error) {
	if c.err != nil {
		return 0, c.err
	}
	return strconv.Atoi(c.val)
}

func (c *StringCmd) Int64() (int64, error) {
	if c.err != nil {
		return 0, c.err
	}
	return strconv.ParseInt(c.val, 10, 64)
}

func (c *StringCmd) Uint64() (uint64, error) {
	if c.err != nil {
		return 0, c.err
	}
	return strconv.ParseUint(c.val, 10, 64)
}

func (c *StringCmd) Float64() (float64, error) {
	if c.err != nil {
		return 0, c.err
	}
	return strconv.ParseFloat(c.val, 64)
}

func (c *StringCmd) Time() (time.Time, error) {
	if c.err != nil {
		return time.Time{}, c.err
	}
	return time.Parse(time.RFC3339Nano, c.val)
}

type StatusCmd struct {
	baseCmd
	val string
}

func NewStatusCmd(client *Redis, name string, args ...interface{}) *StatusCmd {
	cmd := &StatusCmd{
		baseCmd: baseCmd{
			client: client,
			name:   name,
			args:   args,
			result: &result{},
		},
	}
	cmd.result.bind(&cmd.val)
	return cmd
}

func (c *StatusCmd) Val() string {
	return c.val
}

func (c *StatusCmd) Result() (string, error) {
	return c.val, c.err
}

type IntCmd struct {
	baseCmd
	val int64
}

func NewIntCmd(client *Redis, name string, args ...interface{}) *IntCmd {
	cmd := &IntCmd{
		baseCmd: baseCmd{
			client: client,
			name:   name,
			args:   args,
			result: &result{},
		},
	}
	cmd.result.bind(&cmd.val)
	return cmd
}

func (c *IntCmd) Val() int64 {
	return c.val
}

func (c *IntCmd) Result() (int64, error) {
	return c.val, c.err
}

func (c *IntCmd) Uint64() (uint64, error) {
	return uint64(c.val), c.err
}

type DurationCmd struct {
	baseCmd
	val time.Duration
}

func NewDurationCmd(client *Redis, name string, args ...interface{}) *DurationCmd {
	cmd := &DurationCmd{
		baseCmd: baseCmd{
			client: client,
			name:   name,
			args:   args,
			result: &result{},
		},
	}
	cmd.result.bind(&cmd.val)
	return cmd
}

func (c *DurationCmd) Val() time.Duration {
	return c.val
}

func (c *DurationCmd) Result() (time.Duration, error) {
	return c.val, c.err
}

type SliceCmd struct {
	baseCmd
	val []interface{}
}

func NewSliceCmd(client *Redis, name string, args ...interface{}) *SliceCmd {
	cmd := &SliceCmd{
		baseCmd: baseCmd{
			client: client,
			name:   name,
			args:   args,
			result: &result{},
		},
	}
	cmd.result.bind(&cmd.val)
	return cmd
}

func (c *SliceCmd) Val() []interface{} {
	return c.val
}

func (c *SliceCmd) Result() ([]interface{}, error) {
	return c.val, c.err
}

type FloatCmd struct {
	baseCmd
	val float64
}

func NewFloatCmd(client *Redis, name string, args ...interface{}) *FloatCmd {
	cmd := &FloatCmd{
		baseCmd: baseCmd{
			client: client,
			name:   name,
			args:   args,
			result: &result{},
		},
	}
	cmd.result.bind(&cmd.val)
	return cmd
}

func (c *FloatCmd) Val() float64 {
	val, err := c.Result()
	if err != nil {
		return 0
	}
	return val
}

func (c *FloatCmd) Result() (float64, error) {
	return c.val, c.err
}

type BoolCmd struct {
	baseCmd
	val bool
}

func NewBoolCmd(client *Redis, name string, args ...interface{}) *BoolCmd {
	cmd := &BoolCmd{
		baseCmd: baseCmd{
			client: client,
			name:   name,
			args:   args,
			result: &result{},
		},
	}
	cmd.result.bind(&cmd.val)
	return cmd
}

func (c *BoolCmd) Val() bool {
	return c.val
}

func (c *BoolCmd) Result() (bool, error) {
	return c.val, c.err
}

type StringStringMapCmd struct {
	baseCmd
	val map[string]string
}

func NewStringStringMapCmd(client *Redis, name string, args ...interface{}) *StringStringMapCmd {
	cmd := &StringStringMapCmd{
		baseCmd: baseCmd{
			client: client,
			name:   name,
			args:   args,
			result: &result{},
		},
	}
	cmd.result.bind(&cmd.val)
	return cmd
}

func (c *StringStringMapCmd) Val() map[string]string {
	return c.val
}

func (c *StringStringMapCmd) Result() (map[string]string, error) {
	return c.val, c.err
}

type StringSliceCmd struct {
	baseCmd
	val []string
}

func NewStringSliceCmd(client *Redis, name string, args ...interface{}) *StringSliceCmd {
	cmd := &StringSliceCmd{
		baseCmd: baseCmd{
			client: client,
			name:   name,
			args:   args,
			result: &result{},
		},
	}
	cmd.result.bind(&cmd.val)
	return cmd
}

func (c *StringSliceCmd) Val() []string {
	return c.val
}

func (c *StringSliceCmd) Result() ([]string, error) {
	return c.val, c.err
}

type ZSliceCmd struct {
	baseCmd
	val ZSlice
}

func NewZSliceCmd(client *Redis, name string, args ...interface{}) *ZSliceCmd {
	cmd := &ZSliceCmd{
		baseCmd: baseCmd{
			client: client,
			name:   name,
			args:   args,
			result: &result{},
		},
	}
	cmd.result.bind(&cmd.val)
	return cmd
}

func (c *ZSliceCmd) Val() []Z {
	return c.val
}

func (c *ZSliceCmd) Result() ([]Z, error) {
	return c.val, c.err
}
