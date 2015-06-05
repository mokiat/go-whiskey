// This file was generated by counterfeiter
package client_fakes

import (
	"sync"

	"github.com/momchil-atanasov/go-whiskey/graphics/client"
)

type FakeElementClient struct {
	DrawTrianglesStub        func(indexCount, indexOffsetInBytes int) error
	drawTrianglesMutex       sync.RWMutex
	drawTrianglesArgsForCall []struct {
		indexCount         int
		indexOffsetInBytes int
	}
	drawTrianglesReturns struct {
		result1 error
	}
}

func (fake *FakeElementClient) DrawTriangles(indexCount int, indexOffsetInBytes int) error {
	fake.drawTrianglesMutex.Lock()
	fake.drawTrianglesArgsForCall = append(fake.drawTrianglesArgsForCall, struct {
		indexCount         int
		indexOffsetInBytes int
	}{indexCount, indexOffsetInBytes})
	fake.drawTrianglesMutex.Unlock()
	if fake.DrawTrianglesStub != nil {
		return fake.DrawTrianglesStub(indexCount, indexOffsetInBytes)
	} else {
		return fake.drawTrianglesReturns.result1
	}
}

func (fake *FakeElementClient) DrawTrianglesCallCount() int {
	fake.drawTrianglesMutex.RLock()
	defer fake.drawTrianglesMutex.RUnlock()
	return len(fake.drawTrianglesArgsForCall)
}

func (fake *FakeElementClient) DrawTrianglesArgsForCall(i int) (int, int) {
	fake.drawTrianglesMutex.RLock()
	defer fake.drawTrianglesMutex.RUnlock()
	return fake.drawTrianglesArgsForCall[i].indexCount, fake.drawTrianglesArgsForCall[i].indexOffsetInBytes
}

func (fake *FakeElementClient) DrawTrianglesReturns(result1 error) {
	fake.DrawTrianglesStub = nil
	fake.drawTrianglesReturns = struct {
		result1 error
	}{result1}
}

var _ client.ElementClient = new(FakeElementClient)
