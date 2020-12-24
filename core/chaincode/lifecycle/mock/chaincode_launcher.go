// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/Randyshu2018/fabric/core/chaincode/lifecycle"
)

type ChaincodeLauncher struct {
	LaunchStub        func(string) error
	launchMutex       sync.RWMutex
	launchArgsForCall []struct {
		arg1 string
	}
	launchReturns struct {
		result1 error
	}
	launchReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func(string) error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		arg1 string
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChaincodeLauncher) Launch(arg1 string) error {
	fake.launchMutex.Lock()
	ret, specificReturn := fake.launchReturnsOnCall[len(fake.launchArgsForCall)]
	fake.launchArgsForCall = append(fake.launchArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Launch", []interface{}{arg1})
	fake.launchMutex.Unlock()
	if fake.LaunchStub != nil {
		return fake.LaunchStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.launchReturns
	return fakeReturns.result1
}

func (fake *ChaincodeLauncher) LaunchCallCount() int {
	fake.launchMutex.RLock()
	defer fake.launchMutex.RUnlock()
	return len(fake.launchArgsForCall)
}

func (fake *ChaincodeLauncher) LaunchCalls(stub func(string) error) {
	fake.launchMutex.Lock()
	defer fake.launchMutex.Unlock()
	fake.LaunchStub = stub
}

func (fake *ChaincodeLauncher) LaunchArgsForCall(i int) string {
	fake.launchMutex.RLock()
	defer fake.launchMutex.RUnlock()
	argsForCall := fake.launchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChaincodeLauncher) LaunchReturns(result1 error) {
	fake.launchMutex.Lock()
	defer fake.launchMutex.Unlock()
	fake.LaunchStub = nil
	fake.launchReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeLauncher) LaunchReturnsOnCall(i int, result1 error) {
	fake.launchMutex.Lock()
	defer fake.launchMutex.Unlock()
	fake.LaunchStub = nil
	if fake.launchReturnsOnCall == nil {
		fake.launchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.launchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeLauncher) Stop(arg1 string) error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Stop", []interface{}{arg1})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stopReturns
	return fakeReturns.result1
}

func (fake *ChaincodeLauncher) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *ChaincodeLauncher) StopCalls(stub func(string) error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *ChaincodeLauncher) StopArgsForCall(i int) string {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	argsForCall := fake.stopArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChaincodeLauncher) StopReturns(result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeLauncher) StopReturnsOnCall(i int, result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChaincodeLauncher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.launchMutex.RLock()
	defer fake.launchMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChaincodeLauncher) recordInvocation(key string, args []interface{}) {
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

var _ lifecycle.ChaincodeLauncher = new(ChaincodeLauncher)
