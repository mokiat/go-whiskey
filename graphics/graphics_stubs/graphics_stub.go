// Generated by 'github.com/momchil-atanasov/gostub'

package graphics_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/go-whiskey/graphics"
)

type GraphicsStub struct {
	CreateFloat2AttributeArrayStub        func(arg1 int) (result1 alias1.Float2AttributeArray)
	createFloat2AttributeArrayMutex       sync.RWMutex
	createFloat2AttributeArrayArgsForCall []struct {
		arg1 int
	}
	createFloat2AttributeArrayReturns struct {
		result1 alias1.Float2AttributeArray
	}
	CreateFloat3AttributeArrayStub        func(arg1 int) (result1 alias1.Float3AttributeArray)
	createFloat3AttributeArrayMutex       sync.RWMutex
	createFloat3AttributeArrayArgsForCall []struct {
		arg1 int
	}
	createFloat3AttributeArrayReturns struct {
		result1 alias1.Float3AttributeArray
	}
}

var _ alias1.Graphics = new(GraphicsStub)

func (stub *GraphicsStub) CreateFloat2AttributeArray(arg1 int) alias1.Float2AttributeArray {
	stub.createFloat2AttributeArrayMutex.Lock()
	defer stub.createFloat2AttributeArrayMutex.Unlock()
	stub.createFloat2AttributeArrayArgsForCall = append(stub.createFloat2AttributeArrayArgsForCall, struct {
		arg1 int
	}{arg1})
	if stub.CreateFloat2AttributeArrayStub != nil {
		return stub.CreateFloat2AttributeArrayStub(arg1)
	} else {
		return stub.createFloat2AttributeArrayReturns.result1
	}
}
func (stub *GraphicsStub) CreateFloat2AttributeArrayCallCount() int {
	stub.createFloat2AttributeArrayMutex.RLock()
	defer stub.createFloat2AttributeArrayMutex.RUnlock()
	return len(stub.createFloat2AttributeArrayArgsForCall)
}
func (stub *GraphicsStub) CreateFloat2AttributeArrayArgsForCall(index int) int {
	stub.createFloat2AttributeArrayMutex.RLock()
	defer stub.createFloat2AttributeArrayMutex.RUnlock()
	return stub.createFloat2AttributeArrayArgsForCall[index].arg1
}
func (stub *GraphicsStub) CreateFloat2AttributeArrayReturns(result1 alias1.Float2AttributeArray) {
	stub.createFloat2AttributeArrayMutex.Lock()
	defer stub.createFloat2AttributeArrayMutex.Unlock()
	stub.createFloat2AttributeArrayReturns = struct {
		result1 alias1.Float2AttributeArray
	}{result1}
}
func (stub *GraphicsStub) CreateFloat3AttributeArray(arg1 int) alias1.Float3AttributeArray {
	stub.createFloat3AttributeArrayMutex.Lock()
	defer stub.createFloat3AttributeArrayMutex.Unlock()
	stub.createFloat3AttributeArrayArgsForCall = append(stub.createFloat3AttributeArrayArgsForCall, struct {
		arg1 int
	}{arg1})
	if stub.CreateFloat3AttributeArrayStub != nil {
		return stub.CreateFloat3AttributeArrayStub(arg1)
	} else {
		return stub.createFloat3AttributeArrayReturns.result1
	}
}
func (stub *GraphicsStub) CreateFloat3AttributeArrayCallCount() int {
	stub.createFloat3AttributeArrayMutex.RLock()
	defer stub.createFloat3AttributeArrayMutex.RUnlock()
	return len(stub.createFloat3AttributeArrayArgsForCall)
}
func (stub *GraphicsStub) CreateFloat3AttributeArrayArgsForCall(index int) int {
	stub.createFloat3AttributeArrayMutex.RLock()
	defer stub.createFloat3AttributeArrayMutex.RUnlock()
	return stub.createFloat3AttributeArrayArgsForCall[index].arg1
}
func (stub *GraphicsStub) CreateFloat3AttributeArrayReturns(result1 alias1.Float3AttributeArray) {
	stub.createFloat3AttributeArrayMutex.Lock()
	defer stub.createFloat3AttributeArrayMutex.Unlock()
	stub.createFloat3AttributeArrayReturns = struct {
		result1 alias1.Float3AttributeArray
	}{result1}
}
