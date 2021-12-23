package cond

import (
	"fmt"
	"reflect"
)

type M map[string]interface{}

type A []interface{}

// Condition is a where condition or filter
type Condition struct {
	F interface{} // filter: struct map
	E error       // error in merging
}

// Or merge elems and current condition with 'or'
func (c Condition) Or(elems ...interface{}) Condition {
	a := c.merge(elems...)
	if len(a) == 0 {
		return c
	}
	if len(a) == 1 {
		c.F = a[0]
		return c
	}
	c.F = M{OPOr: a}
	return c
}

// And merge elems and current condition with 'and'
func (c Condition) And(elems ...interface{}) Condition {
	a := c.merge(elems...)
	if len(a) == 0 {
		return c
	}
	if len(a) == 1 {
		c.F = a[0]
		return c
	}
	c.F = M{OPAnd: a}
	return c
}

// Nor merge elems and current condition with 'nor'
func (c Condition) Nor(elems ...interface{}) Condition {
	a := c.merge(elems...)
	if len(a) == 0 {
		return c
	}
	if len(a) == 1 {
		c.F = a[0]
		return c
	}
	c.F = M{OPNor: a}
	return c
}

func (c Condition) merge(elems ...interface{}) A {
	a := A{}
	if c.F != nil {
		a = append(a, c.F)
	}

	for _, elem := range elems {
		typ := reflect.TypeOf(elem)
		val := reflect.ValueOf(elem)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
			val = val.Elem()
		}
		switch typ.Kind() {
		case reflect.Map:
			a = append(a, elem)
		case reflect.Struct:
			if typ == reflect.TypeOf(Condition{}) {
				anotherC := val.Interface().(Condition)
				if anotherC.E != nil {
					c.E = anotherC.E
					return a
				}
				a = append(a, val.Interface().(Condition).F)
			} else {
				a = append(a, elem)
			}
		default:
			c.E = fmt.Errorf("invalidate condition %v", elems)
			break
		}
	}
	return a
}

// Or merge elems with '$or' and create a Condition
func Or(elems ...interface{}) Condition {
	c := Condition{}
	return c.Or(elems...)
}

// And merge elems with '$and' and create a Condition
func And(elems ...interface{}) Condition {
	c := Condition{}
	return c.And(elems...)
}

// Nor merge elems with '$nor' and create a Condition
func Nor(elems ...interface{}) Condition {
	c := Condition{}
	return c.Nor(elems...)
}

// Gt '$gt'
func Gt(v interface{}) M {
	return M{OPGt: v}
}

// Gte '$gte'
func Gte(v interface{}) M {
	return M{OPGte: v}
}

// Eq '$eq'
func Eq(v interface{}) M {
	return M{OPEq: v}
}

// Ne '$eq'
func Ne(v interface{}) M {
	return M{OPNe: v}
}

// Lt '$lt'
func Lt(v interface{}) M {
	return M{OPLt: v}
}

// Lte '$lte'
func Lte(v interface{}) M {
	return M{OPLte: v}
}

// In '$in'
func In(v interface{}) M {
	return M{OPIn: v}
}

// Nin '$nin'
func Nin(v interface{}) M {
	return M{OPNotIn: v}
}

// Not 'not'
func Not(v interface{}) M {
	return M{OPNot: v}
}

// Regex $regex
func Regex(v interface{}) M {
	return M{OPRegex: v}
}