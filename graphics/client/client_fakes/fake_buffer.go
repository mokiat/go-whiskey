// This file was generated by counterfeiter
package client_fakes

import (
	"sync"

	"github.com/momchil-atanasov/go-whiskey/graphics/client"
)

type FakeBufferClient struct {
	BindVertexBufferStub        func(id client.BufferId) error
	bindVertexBufferMutex       sync.RWMutex
	bindVertexBufferArgsForCall []struct {
		id client.BufferId
	}
	bindVertexBufferReturns struct {
		result1 error
	}
	BindIndexBufferStub        func(id client.BufferId) error
	bindIndexBufferMutex       sync.RWMutex
	bindIndexBufferArgsForCall []struct {
		id client.BufferId
	}
	bindIndexBufferReturns struct {
		result1 error
	}
}

func (fake *FakeBufferClient) BindVertexBuffer(id client.BufferId) error {
	fake.bindVertexBufferMutex.Lock()
	fake.bindVertexBufferArgsForCall = append(fake.bindVertexBufferArgsForCall, struct {
		id client.BufferId
	}{id})
	fake.bindVertexBufferMutex.Unlock()
	if fake.BindVertexBufferStub != nil {
		return fake.BindVertexBufferStub(id)
	} else {
		return fake.bindVertexBufferReturns.result1
	}
}

func (fake *FakeBufferClient) BindVertexBufferCallCount() int {
	fake.bindVertexBufferMutex.RLock()
	defer fake.bindVertexBufferMutex.RUnlock()
	return len(fake.bindVertexBufferArgsForCall)
}

func (fake *FakeBufferClient) BindVertexBufferArgsForCall(i int) client.BufferId {
	fake.bindVertexBufferMutex.RLock()
	defer fake.bindVertexBufferMutex.RUnlock()
	return fake.bindVertexBufferArgsForCall[i].id
}

func (fake *FakeBufferClient) BindVertexBufferReturns(result1 error) {
	fake.BindVertexBufferStub = nil
	fake.bindVertexBufferReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBufferClient) BindIndexBuffer(id client.BufferId) error {
	fake.bindIndexBufferMutex.Lock()
	fake.bindIndexBufferArgsForCall = append(fake.bindIndexBufferArgsForCall, struct {
		id client.BufferId
	}{id})
	fake.bindIndexBufferMutex.Unlock()
	if fake.BindIndexBufferStub != nil {
		return fake.BindIndexBufferStub(id)
	} else {
		return fake.bindIndexBufferReturns.result1
	}
}

func (fake *FakeBufferClient) BindIndexBufferCallCount() int {
	fake.bindIndexBufferMutex.RLock()
	defer fake.bindIndexBufferMutex.RUnlock()
	return len(fake.bindIndexBufferArgsForCall)
}

func (fake *FakeBufferClient) BindIndexBufferArgsForCall(i int) client.BufferId {
	fake.bindIndexBufferMutex.RLock()
	defer fake.bindIndexBufferMutex.RUnlock()
	return fake.bindIndexBufferArgsForCall[i].id
}

func (fake *FakeBufferClient) BindIndexBufferReturns(result1 error) {
	fake.BindIndexBufferStub = nil
	fake.bindIndexBufferReturns = struct {
		result1 error
	}{result1}
}

var _ client.BufferClient = new(FakeBufferClient)
