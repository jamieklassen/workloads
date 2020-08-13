// Code generated by counterfeiter. DO NOT EDIT.
package accountsfakes

import (
	"sync"

	"github.com/concourse/ft/accounts"
)

type FakeWorker struct {
	ContainersStub        func(...accounts.StatsOption) ([]accounts.Container, error)
	containersMutex       sync.RWMutex
	containersArgsForCall []struct {
		arg1 []accounts.StatsOption
	}
	containersReturns struct {
		result1 []accounts.Container
		result2 error
	}
	containersReturnsOnCall map[int]struct {
		result1 []accounts.Container
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorker) Containers(arg1 ...accounts.StatsOption) ([]accounts.Container, error) {
	fake.containersMutex.Lock()
	ret, specificReturn := fake.containersReturnsOnCall[len(fake.containersArgsForCall)]
	fake.containersArgsForCall = append(fake.containersArgsForCall, struct {
		arg1 []accounts.StatsOption
	}{arg1})
	fake.recordInvocation("Containers", []interface{}{arg1})
	fake.containersMutex.Unlock()
	if fake.ContainersStub != nil {
		return fake.ContainersStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.containersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorker) ContainersCallCount() int {
	fake.containersMutex.RLock()
	defer fake.containersMutex.RUnlock()
	return len(fake.containersArgsForCall)
}

func (fake *FakeWorker) ContainersCalls(stub func(...accounts.StatsOption) ([]accounts.Container, error)) {
	fake.containersMutex.Lock()
	defer fake.containersMutex.Unlock()
	fake.ContainersStub = stub
}

func (fake *FakeWorker) ContainersArgsForCall(i int) []accounts.StatsOption {
	fake.containersMutex.RLock()
	defer fake.containersMutex.RUnlock()
	argsForCall := fake.containersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWorker) ContainersReturns(result1 []accounts.Container, result2 error) {
	fake.containersMutex.Lock()
	defer fake.containersMutex.Unlock()
	fake.ContainersStub = nil
	fake.containersReturns = struct {
		result1 []accounts.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) ContainersReturnsOnCall(i int, result1 []accounts.Container, result2 error) {
	fake.containersMutex.Lock()
	defer fake.containersMutex.Unlock()
	fake.ContainersStub = nil
	if fake.containersReturnsOnCall == nil {
		fake.containersReturnsOnCall = make(map[int]struct {
			result1 []accounts.Container
			result2 error
		})
	}
	fake.containersReturnsOnCall[i] = struct {
		result1 []accounts.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.containersMutex.RLock()
	defer fake.containersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorker) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ accounts.Worker = new(FakeWorker)
