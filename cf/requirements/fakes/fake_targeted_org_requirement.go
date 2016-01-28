// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
)

type FakeTargetedOrgRequirement struct {
	ExecuteStub        func() (success bool)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct{}
	executeReturns     struct {
		result1 bool
	}
	GetOrganizationFieldsStub        func() models.OrganizationFields
	getOrganizationFieldsMutex       sync.RWMutex
	getOrganizationFieldsArgsForCall []struct{}
	getOrganizationFieldsReturns     struct {
		result1 models.OrganizationFields
	}
}

func (fake *FakeTargetedOrgRequirement) Execute() (success bool) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub()
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeTargetedOrgRequirement) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeTargetedOrgRequirement) ExecuteReturns(result1 bool) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeTargetedOrgRequirement) GetOrganizationFields() models.OrganizationFields {
	fake.getOrganizationFieldsMutex.Lock()
	fake.getOrganizationFieldsArgsForCall = append(fake.getOrganizationFieldsArgsForCall, struct{}{})
	fake.getOrganizationFieldsMutex.Unlock()
	if fake.GetOrganizationFieldsStub != nil {
		return fake.GetOrganizationFieldsStub()
	} else {
		return fake.getOrganizationFieldsReturns.result1
	}
}

func (fake *FakeTargetedOrgRequirement) GetOrganizationFieldsCallCount() int {
	fake.getOrganizationFieldsMutex.RLock()
	defer fake.getOrganizationFieldsMutex.RUnlock()
	return len(fake.getOrganizationFieldsArgsForCall)
}

func (fake *FakeTargetedOrgRequirement) GetOrganizationFieldsReturns(result1 models.OrganizationFields) {
	fake.GetOrganizationFieldsStub = nil
	fake.getOrganizationFieldsReturns = struct {
		result1 models.OrganizationFields
	}{result1}
}

var _ requirements.TargetedOrgRequirement = new(FakeTargetedOrgRequirement)
