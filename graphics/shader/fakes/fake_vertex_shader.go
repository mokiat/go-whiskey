// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/momchil-atanasov/go-whiskey/graphics"
	"github.com/momchil-atanasov/go-whiskey/graphics/shader"
)

type FakeVertexShader struct {
	IdStub        func() graphics.ResourceId
	idMutex       sync.RWMutex
	idArgsForCall []struct{}
	idReturns struct {
		result1 graphics.ResourceId
	}
	SourceCodeStub        func() string
	sourceCodeMutex       sync.RWMutex
	sourceCodeArgsForCall []struct{}
	sourceCodeReturns struct {
		result1 string
	}
	CreateRemotelyStub        func() error
	createRemotelyMutex       sync.RWMutex
	createRemotelyArgsForCall []struct{}
	createRemotelyReturns struct {
		result1 error
	}
	DeleteRemotelyStub        func()
	deleteRemotelyMutex       sync.RWMutex
	deleteRemotelyArgsForCall []struct{}
	CreatedRemotelyStub        func() bool
	createdRemotelyMutex       sync.RWMutex
	createdRemotelyArgsForCall []struct{}
	createdRemotelyReturns struct {
		result1 bool
	}
}

func (fake *FakeVertexShader) Id() graphics.ResourceId {
	fake.idMutex.Lock()
	fake.idArgsForCall = append(fake.idArgsForCall, struct{}{})
	fake.idMutex.Unlock()
	if fake.IdStub != nil {
		return fake.IdStub()
	} else {
		return fake.idReturns.result1
	}
}

func (fake *FakeVertexShader) IdCallCount() int {
	fake.idMutex.RLock()
	defer fake.idMutex.RUnlock()
	return len(fake.idArgsForCall)
}

func (fake *FakeVertexShader) IdReturns(result1 graphics.ResourceId) {
	fake.IdStub = nil
	fake.idReturns = struct {
		result1 graphics.ResourceId
	}{result1}
}

func (fake *FakeVertexShader) SourceCode() string {
	fake.sourceCodeMutex.Lock()
	fake.sourceCodeArgsForCall = append(fake.sourceCodeArgsForCall, struct{}{})
	fake.sourceCodeMutex.Unlock()
	if fake.SourceCodeStub != nil {
		return fake.SourceCodeStub()
	} else {
		return fake.sourceCodeReturns.result1
	}
}

func (fake *FakeVertexShader) SourceCodeCallCount() int {
	fake.sourceCodeMutex.RLock()
	defer fake.sourceCodeMutex.RUnlock()
	return len(fake.sourceCodeArgsForCall)
}

func (fake *FakeVertexShader) SourceCodeReturns(result1 string) {
	fake.SourceCodeStub = nil
	fake.sourceCodeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVertexShader) CreateRemotely() error {
	fake.createRemotelyMutex.Lock()
	fake.createRemotelyArgsForCall = append(fake.createRemotelyArgsForCall, struct{}{})
	fake.createRemotelyMutex.Unlock()
	if fake.CreateRemotelyStub != nil {
		return fake.CreateRemotelyStub()
	} else {
		return fake.createRemotelyReturns.result1
	}
}

func (fake *FakeVertexShader) CreateRemotelyCallCount() int {
	fake.createRemotelyMutex.RLock()
	defer fake.createRemotelyMutex.RUnlock()
	return len(fake.createRemotelyArgsForCall)
}

func (fake *FakeVertexShader) CreateRemotelyReturns(result1 error) {
	fake.CreateRemotelyStub = nil
	fake.createRemotelyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVertexShader) DeleteRemotely() {
	fake.deleteRemotelyMutex.Lock()
	fake.deleteRemotelyArgsForCall = append(fake.deleteRemotelyArgsForCall, struct{}{})
	fake.deleteRemotelyMutex.Unlock()
	if fake.DeleteRemotelyStub != nil {
		fake.DeleteRemotelyStub()
	}
}

func (fake *FakeVertexShader) DeleteRemotelyCallCount() int {
	fake.deleteRemotelyMutex.RLock()
	defer fake.deleteRemotelyMutex.RUnlock()
	return len(fake.deleteRemotelyArgsForCall)
}

func (fake *FakeVertexShader) CreatedRemotely() bool {
	fake.createdRemotelyMutex.Lock()
	fake.createdRemotelyArgsForCall = append(fake.createdRemotelyArgsForCall, struct{}{})
	fake.createdRemotelyMutex.Unlock()
	if fake.CreatedRemotelyStub != nil {
		return fake.CreatedRemotelyStub()
	} else {
		return fake.createdRemotelyReturns.result1
	}
}

func (fake *FakeVertexShader) CreatedRemotelyCallCount() int {
	fake.createdRemotelyMutex.RLock()
	defer fake.createdRemotelyMutex.RUnlock()
	return len(fake.createdRemotelyArgsForCall)
}

func (fake *FakeVertexShader) CreatedRemotelyReturns(result1 bool) {
	fake.CreatedRemotelyStub = nil
	fake.createdRemotelyReturns = struct {
		result1 bool
	}{result1}
}

var _ shader.VertexShader = new(FakeVertexShader)
