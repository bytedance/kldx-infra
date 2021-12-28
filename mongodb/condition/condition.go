package cond

import op "github/kldx/infra/mongodb/operator"

type M map[string]interface{}

type A []interface{}

func And(exps ...interface{}) M {
	if len(exps) == 0 {
		return nil
	}
	return M{op.And: exps}
}

func Or(exps ...interface{}) M {
	if len(exps) == 0 {
		return nil
	}

	return M{op.Or: exps}
}

func Eq(v interface{}) M {
	return M{op.Eq: v}
}

func Gt(v interface{}) M {
	return M{op.Gt: v}
}

func Gte(v interface{}) M {
	return M{op.Gte: v}
}

func Ne(v interface{}) M {
	return M{op.Ne: v}
}

func Lt(v interface{}) M {
	return M{op.Lt: v}
}

func Lte(v interface{}) M {
	return M{op.Lte: v}
}

func In(v interface{}) M {
	return M{op.In: v}
}

func Nin(v interface{}) M {
	return M{op.NotIn: v}
}

func Not(v interface{}) M {
	return M{op.Not: v}
}

func Regex(v interface{}) M {
	return M{op.Regex: v}
}
