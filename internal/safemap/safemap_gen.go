// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package safemap

import (
	"context"
	"sync"
)

type ()

type StringToString struct {
	sync.RWMutex
	data map[string]string
}

func NewStringToString() StringToString {
	return StringToString{
		data: map[string]string{},
	}
}

func (safe *StringToString) Load(key string) (string, bool) {
	safe.RLock()
	defer safe.RUnlock()

	value, ok := safe.data[key]
	return value, ok
}

func (safe *StringToString) Store(key, value string) {
	safe.Lock()
	defer safe.Unlock()

	safe.data[key] = value
}

func (safe *StringToString) LoadOrStore(key, value string) string {
	safe.Lock()
	defer safe.Unlock()

	existing, ok := safe.data[key]
	if ok {
		return existing
	}

	safe.data[key] = value

	return value
}

func (safe *StringToString) Delete(key string) {
	safe.Lock()
	defer safe.Unlock()

	_, ok := safe.data[key]
	if ok {
		delete(safe.data, key)
	}
}

func (safe *StringToString) Range(fn func(key, value string) bool) {
	safe.Lock()
	defer safe.Unlock()

	for key, value := range safe.data {
		if !fn(key, value) {
			break
		}
	}
}

type ()

type StringToContextCancelFunc struct {
	sync.RWMutex
	data map[string]context.CancelFunc
}

func NewStringToContextCancelFunc() StringToContextCancelFunc {
	return StringToContextCancelFunc{
		data: map[string]context.CancelFunc{},
	}
}

func (safe *StringToContextCancelFunc) Load(key string) (context.CancelFunc, bool) {
	safe.RLock()
	defer safe.RUnlock()

	value, ok := safe.data[key]
	return value, ok
}

func (safe *StringToContextCancelFunc) Store(key string, value context.CancelFunc) {
	safe.Lock()
	defer safe.Unlock()

	safe.data[key] = value
}

func (safe *StringToContextCancelFunc) LoadOrStore(key string, value context.CancelFunc) context.CancelFunc {
	safe.Lock()
	defer safe.Unlock()

	existing, ok := safe.data[key]
	if ok {
		return existing
	}

	safe.data[key] = value

	return value
}

func (safe *StringToContextCancelFunc) Delete(key string) {
	safe.Lock()
	defer safe.Unlock()

	_, ok := safe.data[key]
	if ok {
		delete(safe.data, key)
	}
}

func (safe *StringToContextCancelFunc) Range(fn func(key string, value context.CancelFunc) bool) {
	safe.Lock()
	defer safe.Unlock()

	for key, value := range safe.data {
		if !fn(key, value) {
			break
		}
	}
}

type ()

type StringToAny struct {
	sync.RWMutex
	data map[string]Any
}

func NewStringToAny() StringToAny {
	return StringToAny{
		data: map[string]Any{},
	}
}

func (safe *StringToAny) Load(key string) (Any, bool) {
	safe.RLock()
	defer safe.RUnlock()

	value, ok := safe.data[key]
	return value, ok
}

func (safe *StringToAny) Store(key string, value Any) {
	safe.Lock()
	defer safe.Unlock()

	safe.data[key] = value
}

func (safe *StringToAny) LoadOrStore(key string, value Any) Any {
	safe.Lock()
	defer safe.Unlock()

	existing, ok := safe.data[key]
	if ok {
		return existing
	}

	safe.data[key] = value

	return value
}

func (safe *StringToAny) Delete(key string) {
	safe.Lock()
	defer safe.Unlock()

	_, ok := safe.data[key]
	if ok {
		delete(safe.data, key)
	}
}

func (safe *StringToAny) Range(fn func(key string, value Any) bool) {
	safe.Lock()
	defer safe.Unlock()

	for key, value := range safe.data {
		if !fn(key, value) {
			break
		}
	}
}

type ()

type IntToString struct {
	sync.RWMutex
	data map[int]string
}

func NewIntToString() IntToString {
	return IntToString{
		data: map[int]string{},
	}
}

func (safe *IntToString) Load(key int) (string, bool) {
	safe.RLock()
	defer safe.RUnlock()

	value, ok := safe.data[key]
	return value, ok
}

func (safe *IntToString) Store(key int, value string) {
	safe.Lock()
	defer safe.Unlock()

	safe.data[key] = value
}

func (safe *IntToString) LoadOrStore(key int, value string) string {
	safe.Lock()
	defer safe.Unlock()

	existing, ok := safe.data[key]
	if ok {
		return existing
	}

	safe.data[key] = value

	return value
}

func (safe *IntToString) Delete(key int) {
	safe.Lock()
	defer safe.Unlock()

	_, ok := safe.data[key]
	if ok {
		delete(safe.data, key)
	}
}

func (safe *IntToString) Range(fn func(key int, value string) bool) {
	safe.Lock()
	defer safe.Unlock()

	for key, value := range safe.data {
		if !fn(key, value) {
			break
		}
	}
}

type ()

type IntToContextCancelFunc struct {
	sync.RWMutex
	data map[int]context.CancelFunc
}

func NewIntToContextCancelFunc() IntToContextCancelFunc {
	return IntToContextCancelFunc{
		data: map[int]context.CancelFunc{},
	}
}

func (safe *IntToContextCancelFunc) Load(key int) (context.CancelFunc, bool) {
	safe.RLock()
	defer safe.RUnlock()

	value, ok := safe.data[key]
	return value, ok
}

func (safe *IntToContextCancelFunc) Store(key int, value context.CancelFunc) {
	safe.Lock()
	defer safe.Unlock()

	safe.data[key] = value
}

func (safe *IntToContextCancelFunc) LoadOrStore(key int, value context.CancelFunc) context.CancelFunc {
	safe.Lock()
	defer safe.Unlock()

	existing, ok := safe.data[key]
	if ok {
		return existing
	}

	safe.data[key] = value

	return value
}

func (safe *IntToContextCancelFunc) Delete(key int) {
	safe.Lock()
	defer safe.Unlock()

	_, ok := safe.data[key]
	if ok {
		delete(safe.data, key)
	}
}

func (safe *IntToContextCancelFunc) Range(fn func(key int, value context.CancelFunc) bool) {
	safe.Lock()
	defer safe.Unlock()

	for key, value := range safe.data {
		if !fn(key, value) {
			break
		}
	}
}

type ()

type IntToAny struct {
	sync.RWMutex
	data map[int]Any
}

func NewIntToAny() IntToAny {
	return IntToAny{
		data: map[int]Any{},
	}
}

func (safe *IntToAny) Load(key int) (Any, bool) {
	safe.RLock()
	defer safe.RUnlock()

	value, ok := safe.data[key]
	return value, ok
}

func (safe *IntToAny) Store(key int, value Any) {
	safe.Lock()
	defer safe.Unlock()

	safe.data[key] = value
}

func (safe *IntToAny) LoadOrStore(key int, value Any) Any {
	safe.Lock()
	defer safe.Unlock()

	existing, ok := safe.data[key]
	if ok {
		return existing
	}

	safe.data[key] = value

	return value
}

func (safe *IntToAny) Delete(key int) {
	safe.Lock()
	defer safe.Unlock()

	_, ok := safe.data[key]
	if ok {
		delete(safe.data, key)
	}
}

func (safe *IntToAny) Range(fn func(key int, value Any) bool) {
	safe.Lock()
	defer safe.Unlock()

	for key, value := range safe.data {
		if !fn(key, value) {
			break
		}
	}
}
