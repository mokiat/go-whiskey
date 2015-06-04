// This file was generated by counterfeiter
package shader_fakes

import (
	"sync"

	"github.com/momchil-atanasov/go-whiskey/graphics/client"
	"github.com/momchil-atanasov/go-whiskey/graphics/shader"
)

type FakeRemoteShader struct {
	IdStub        func() client.ShaderId
	idMutex       sync.RWMutex
	idArgsForCall []struct{}
	idReturns struct {
		result1 client.ShaderId
	}
	CreatedStub        func() bool
	createdMutex       sync.RWMutex
	createdArgsForCall []struct{}
	createdReturns struct {
		result1 bool
	}
	CreateStub        func() error
	createMutex       sync.RWMutex
	createArgsForCall []struct{}
	createReturns struct {
		result1 error
	}
	DeleteStub        func() error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct{}
	deleteReturns struct {
		result1 error
	}
}

func (fake *FakeRemoteShader) Id() client.ShaderId {
	fake.idMutex.Lock()
	fake.idArgsForCall = append(fake.idArgsForCall, struct{}{})
	fake.idMutex.Unlock()
	if fake.IdStub != nil {
		return fake.IdStub()
	} else {
		return fake.idReturns.result1
	}
}

func (fake *FakeRemoteShader) IdCallCount() int {
	fake.idMutex.RLock()
	defer fake.idMutex.RUnlock()
	return len(fake.idArgsForCall)
}

func (fake *FakeRemoteShader) IdReturns(result1 client.ShaderId) {
	fake.IdStub = nil
	fake.idReturns = struct {
		result1 client.ShaderId
	}{result1}
}

func (fake *FakeRemoteShader) Created() bool {
	fake.createdMutex.Lock()
	fake.createdArgsForCall = append(fake.createdArgsForCall, struct{}{})
	fake.createdMutex.Unlock()
	if fake.CreatedStub != nil {
		return fake.CreatedStub()
	} else {
		return fake.createdReturns.result1
	}
}

func (fake *FakeRemoteShader) CreatedCallCount() int {
	fake.createdMutex.RLock()
	defer fake.createdMutex.RUnlock()
	return len(fake.createdArgsForCall)
}

func (fake *FakeRemoteShader) CreatedReturns(result1 bool) {
	fake.CreatedStub = nil
	fake.createdReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRemoteShader) Create() error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct{}{})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub()
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeRemoteShader) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeRemoteShader) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRemoteShader) Delete() error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct{}{})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub()
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeRemoteShader) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeRemoteShader) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ shader.RemoteShader = new(FakeRemoteShader)