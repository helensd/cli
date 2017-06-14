// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeDeleteOrganizationActor struct {
	DeleteOrganizationStub        func(orgName string) ([]string, error)
	deleteOrganizationMutex       sync.RWMutex
	deleteOrganizationArgsForCall []struct {
		orgName string
	}
	deleteOrganizationReturns struct {
		result1 []string
		result2 error
	}
	deleteOrganizationReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ClearOrganizationAndSpaceStub        func(config v2action.Config)
	clearOrganizationAndSpaceMutex       sync.RWMutex
	clearOrganizationAndSpaceArgsForCall []struct {
		config v2action.Config
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganization(orgName string) ([]string, error) {
	fake.deleteOrganizationMutex.Lock()
	ret, specificReturn := fake.deleteOrganizationReturnsOnCall[len(fake.deleteOrganizationArgsForCall)]
	fake.deleteOrganizationArgsForCall = append(fake.deleteOrganizationArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("DeleteOrganization", []interface{}{orgName})
	fake.deleteOrganizationMutex.Unlock()
	if fake.DeleteOrganizationStub != nil {
		return fake.DeleteOrganizationStub(orgName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteOrganizationReturns.result1, fake.deleteOrganizationReturns.result2
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationCallCount() int {
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	return len(fake.deleteOrganizationArgsForCall)
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationArgsForCall(i int) string {
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	return fake.deleteOrganizationArgsForCall[i].orgName
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationReturns(result1 []string, result2 error) {
	fake.DeleteOrganizationStub = nil
	fake.deleteOrganizationReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationReturnsOnCall(i int, result1 []string, result2 error) {
	fake.DeleteOrganizationStub = nil
	if fake.deleteOrganizationReturnsOnCall == nil {
		fake.deleteOrganizationReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.deleteOrganizationReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrganizationActor) ClearOrganizationAndSpace(config v2action.Config) {
	fake.clearOrganizationAndSpaceMutex.Lock()
	fake.clearOrganizationAndSpaceArgsForCall = append(fake.clearOrganizationAndSpaceArgsForCall, struct {
		config v2action.Config
	}{config})
	fake.recordInvocation("ClearOrganizationAndSpace", []interface{}{config})
	fake.clearOrganizationAndSpaceMutex.Unlock()
	if fake.ClearOrganizationAndSpaceStub != nil {
		fake.ClearOrganizationAndSpaceStub(config)
	}
}

func (fake *FakeDeleteOrganizationActor) ClearOrganizationAndSpaceCallCount() int {
	fake.clearOrganizationAndSpaceMutex.RLock()
	defer fake.clearOrganizationAndSpaceMutex.RUnlock()
	return len(fake.clearOrganizationAndSpaceArgsForCall)
}

func (fake *FakeDeleteOrganizationActor) ClearOrganizationAndSpaceArgsForCall(i int) v2action.Config {
	fake.clearOrganizationAndSpaceMutex.RLock()
	defer fake.clearOrganizationAndSpaceMutex.RUnlock()
	return fake.clearOrganizationAndSpaceArgsForCall[i].config
}

func (fake *FakeDeleteOrganizationActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	fake.clearOrganizationAndSpaceMutex.RLock()
	defer fake.clearOrganizationAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteOrganizationActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.DeleteOrganizationActor = new(FakeDeleteOrganizationActor)
