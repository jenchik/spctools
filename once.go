package spctools

import (
	"sync"
)

type OnceString struct {
	val string
	ins sync.Once
}

func NewString() *OnceString {
	return &OnceString{}
}

func NewSetString(s string) *OnceString {
	return &OnceString{val: s}
}

func (o *OnceString) Do(f func() string) {
	o.ins.Do(func() {
		o.val = f()
	})
}

func (o *OnceString) Value() string {
	return o.val
}

type OnceInt struct {
	val int64
	ins sync.Once
}

func NewInt() *OnceInt {
	return &OnceInt{}
}

func NewSetInt(i int64) *OnceInt {
	return &OnceInt{val: i}
}

func (o *OnceInt) Do(f func() int64) {
	o.ins.Do(func() {
		o.val = f()
	})
}

func (o *OnceInt) Value() int64 {
	return o.val
}

type OnceFloat struct {
	val float64
	ins sync.Once
}

func NewFloat() *OnceFloat {
	return &OnceFloat{}
}

func NewSetFloat(f float64) *OnceFloat {
	return &OnceFloat{val: f}
}

func (o *OnceFloat) Do(f func() float64) {
	o.ins.Do(func() {
		o.val = f()
	})
}

func (o *OnceFloat) Value() float64 {
	return o.val
}

type OnceBool struct {
	val bool
	ins sync.Once
}

func NewBool() *OnceBool {
	return &OnceBool{}
}

func NewSetBool(b bool) *OnceBool {
	return &OnceBool{val: b}
}

func (o *OnceBool) Do(f func() bool) {
	o.ins.Do(func() {
		o.val = f()
	})
}

func (o *OnceBool) Value() bool {
	return o.val
}

type OnceValue struct {
	val interface{}
	ins sync.Once
}

func NewValue() *OnceValue {
	return &OnceValue{}
}

func NewSetValue(v interface{}) *OnceValue {
	return &OnceValue{val: v}
}

func (o *OnceValue) Do(f func() interface{}) {
	o.ins.Do(func() {
		o.val = f()
	})
}

func (o *OnceValue) Value() interface{} {
	return o.val
}
