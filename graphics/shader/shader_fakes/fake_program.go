// This file was generated by counterfeiter
package shader_fakes

import (
	"sync"

	"github.com/momchil-atanasov/go-whiskey/graphics/client"
	"github.com/momchil-atanasov/go-whiskey/graphics/shader"
)

type FakeProgram struct {
	VertexShaderStub        func() shader.Shader
	vertexShaderMutex       sync.RWMutex
	vertexShaderArgsForCall []struct{}
	vertexShaderReturns struct {
		result1 shader.Shader
	}
	FragmentShaderStub        func() shader.Shader
	fragmentShaderMutex       sync.RWMutex
	fragmentShaderArgsForCall []struct{}
	fragmentShaderReturns struct {
		result1 shader.Shader
	}
	IdStub        func() client.ProgramId
	idMutex       sync.RWMutex
	idArgsForCall []struct{}
	idReturns struct {
		result1 client.ProgramId
	}
	CreatedStub        func() bool
	createdMutex       sync.RWMutex
	createdArgsForCall []struct{}
	createdReturns struct {
		result1 bool
	}
	CreateStub        func(client.ShaderClient) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 client.ShaderClient
	}
	createReturns struct {
		result1 error
	}
	DeleteStub        func(client.ShaderClient) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 client.ShaderClient
	}
	deleteReturns struct {
		result1 error
	}
}

func (fake *FakeProgram) VertexShader() shader.Shader {
	fake.vertexShaderMutex.Lock()
	fake.vertexShaderArgsForCall = append(fake.vertexShaderArgsForCall, struct{}{})
	fake.vertexShaderMutex.Unlock()
	if fake.VertexShaderStub != nil {
		return fake.VertexShaderStub()
	} else {
		return fake.vertexShaderReturns.result1
	}
}

func (fake *FakeProgram) VertexShaderCallCount() int {
	fake.vertexShaderMutex.RLock()
	defer fake.vertexShaderMutex.RUnlock()
	return len(fake.vertexShaderArgsForCall)
}

func (fake *FakeProgram) VertexShaderReturns(result1 shader.Shader) {
	fake.VertexShaderStub = nil
	fake.vertexShaderReturns = struct {
		result1 shader.Shader
	}{result1}
}

func (fake *FakeProgram) FragmentShader() shader.Shader {
	fake.fragmentShaderMutex.Lock()
	fake.fragmentShaderArgsForCall = append(fake.fragmentShaderArgsForCall, struct{}{})
	fake.fragmentShaderMutex.Unlock()
	if fake.FragmentShaderStub != nil {
		return fake.FragmentShaderStub()
	} else {
		return fake.fragmentShaderReturns.result1
	}
}

func (fake *FakeProgram) FragmentShaderCallCount() int {
	fake.fragmentShaderMutex.RLock()
	defer fake.fragmentShaderMutex.RUnlock()
	return len(fake.fragmentShaderArgsForCall)
}

func (fake *FakeProgram) FragmentShaderReturns(result1 shader.Shader) {
	fake.FragmentShaderStub = nil
	fake.fragmentShaderReturns = struct {
		result1 shader.Shader
	}{result1}
}

func (fake *FakeProgram) Id() client.ProgramId {
	fake.idMutex.Lock()
	fake.idArgsForCall = append(fake.idArgsForCall, struct{}{})
	fake.idMutex.Unlock()
	if fake.IdStub != nil {
		return fake.IdStub()
	} else {
		return fake.idReturns.result1
	}
}

func (fake *FakeProgram) IdCallCount() int {
	fake.idMutex.RLock()
	defer fake.idMutex.RUnlock()
	return len(fake.idArgsForCall)
}

func (fake *FakeProgram) IdReturns(result1 client.ProgramId) {
	fake.IdStub = nil
	fake.idReturns = struct {
		result1 client.ProgramId
	}{result1}
}

func (fake *FakeProgram) Created() bool {
	fake.createdMutex.Lock()
	fake.createdArgsForCall = append(fake.createdArgsForCall, struct{}{})
	fake.createdMutex.Unlock()
	if fake.CreatedStub != nil {
		return fake.CreatedStub()
	} else {
		return fake.createdReturns.result1
	}
}

func (fake *FakeProgram) CreatedCallCount() int {
	fake.createdMutex.RLock()
	defer fake.createdMutex.RUnlock()
	return len(fake.createdArgsForCall)
}

func (fake *FakeProgram) CreatedReturns(result1 bool) {
	fake.CreatedStub = nil
	fake.createdReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeProgram) Create(arg1 client.ShaderClient) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 client.ShaderClient
	}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeProgram) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeProgram) CreateArgsForCall(i int) client.ShaderClient {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *FakeProgram) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProgram) Delete(arg1 client.ShaderClient) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 client.ShaderClient
	}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeProgram) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeProgram) DeleteArgsForCall(i int) client.ShaderClient {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1
}

func (fake *FakeProgram) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ shader.Program = new(FakeProgram)
