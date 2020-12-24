// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/Randyshu2018/fabric/common/channelconfig"
)

type ChannelConfigGetter struct {
	GetChannelConfigStub        func(string) channelconfig.Resources
	getChannelConfigMutex       sync.RWMutex
	getChannelConfigArgsForCall []struct {
		arg1 string
	}
	getChannelConfigReturns struct {
		result1 channelconfig.Resources
	}
	getChannelConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Resources
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChannelConfigGetter) GetChannelConfig(arg1 string) channelconfig.Resources {
	fake.getChannelConfigMutex.Lock()
	ret, specificReturn := fake.getChannelConfigReturnsOnCall[len(fake.getChannelConfigArgsForCall)]
	fake.getChannelConfigArgsForCall = append(fake.getChannelConfigArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetChannelConfig", []interface{}{arg1})
	fake.getChannelConfigMutex.Unlock()
	if fake.GetChannelConfigStub != nil {
		return fake.GetChannelConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getChannelConfigReturns
	return fakeReturns.result1
}

func (fake *ChannelConfigGetter) GetChannelConfigCallCount() int {
	fake.getChannelConfigMutex.RLock()
	defer fake.getChannelConfigMutex.RUnlock()
	return len(fake.getChannelConfigArgsForCall)
}

func (fake *ChannelConfigGetter) GetChannelConfigCalls(stub func(string) channelconfig.Resources) {
	fake.getChannelConfigMutex.Lock()
	defer fake.getChannelConfigMutex.Unlock()
	fake.GetChannelConfigStub = stub
}

func (fake *ChannelConfigGetter) GetChannelConfigArgsForCall(i int) string {
	fake.getChannelConfigMutex.RLock()
	defer fake.getChannelConfigMutex.RUnlock()
	argsForCall := fake.getChannelConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChannelConfigGetter) GetChannelConfigReturns(result1 channelconfig.Resources) {
	fake.getChannelConfigMutex.Lock()
	defer fake.getChannelConfigMutex.Unlock()
	fake.GetChannelConfigStub = nil
	fake.getChannelConfigReturns = struct {
		result1 channelconfig.Resources
	}{result1}
}

func (fake *ChannelConfigGetter) GetChannelConfigReturnsOnCall(i int, result1 channelconfig.Resources) {
	fake.getChannelConfigMutex.Lock()
	defer fake.getChannelConfigMutex.Unlock()
	fake.GetChannelConfigStub = nil
	if fake.getChannelConfigReturnsOnCall == nil {
		fake.getChannelConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Resources
		})
	}
	fake.getChannelConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Resources
	}{result1}
}

func (fake *ChannelConfigGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getChannelConfigMutex.RLock()
	defer fake.getChannelConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChannelConfigGetter) recordInvocation(key string, args []interface{}) {
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
