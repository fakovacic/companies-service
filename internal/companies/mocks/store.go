// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/fakovacic/companies-service/internal/companies"
	"sync"
)

// Ensure, that StoreMock does implement companies.Store.
// If this is not the case, regenerate this file with moq.
var _ companies.Store = &StoreMock{}

// StoreMock is a mock implementation of companies.Store.
//
// 	func TestSomethingThatUsesStore(t *testing.T) {
//
// 		// make and configure a mocked companies.Store
// 		mockedStore := &StoreMock{
// 			CreateFunc: func(contextMoqParam context.Context, company *companies.Company) error {
// 				panic("mock out the Create method")
// 			},
// 			DeleteFunc: func(contextMoqParam context.Context, s string) error {
// 				panic("mock out the Delete method")
// 			},
// 			GetFunc: func(contextMoqParam context.Context, s string) (*companies.Company, error) {
// 				panic("mock out the Get method")
// 			},
// 			UpdateFunc: func(contextMoqParam context.Context, s string, company *companies.Company) error {
// 				panic("mock out the Update method")
// 			},
// 		}
//
// 		// use mockedStore in code that requires companies.Store
// 		// and then make assertions.
//
// 	}
type StoreMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(contextMoqParam context.Context, company *companies.Company) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(contextMoqParam context.Context, s string) error

	// GetFunc mocks the Get method.
	GetFunc func(contextMoqParam context.Context, s string) (*companies.Company, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(contextMoqParam context.Context, s string, company *companies.Company) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Company is the company argument value.
			Company *companies.Company
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// S is the s argument value.
			S string
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// S is the s argument value.
			S string
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// S is the s argument value.
			S string
			// Company is the company argument value.
			Company *companies.Company
		}
	}
	lockCreate sync.RWMutex
	lockDelete sync.RWMutex
	lockGet    sync.RWMutex
	lockUpdate sync.RWMutex
}

// Create calls CreateFunc.
func (mock *StoreMock) Create(contextMoqParam context.Context, company *companies.Company) error {
	if mock.CreateFunc == nil {
		panic("StoreMock.CreateFunc: method is nil but Store.Create was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Company         *companies.Company
	}{
		ContextMoqParam: contextMoqParam,
		Company:         company,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(contextMoqParam, company)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedStore.CreateCalls())
func (mock *StoreMock) CreateCalls() []struct {
	ContextMoqParam context.Context
	Company         *companies.Company
} {
	var calls []struct {
		ContextMoqParam context.Context
		Company         *companies.Company
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *StoreMock) Delete(contextMoqParam context.Context, s string) error {
	if mock.DeleteFunc == nil {
		panic("StoreMock.DeleteFunc: method is nil but Store.Delete was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		S               string
	}{
		ContextMoqParam: contextMoqParam,
		S:               s,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(contextMoqParam, s)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedStore.DeleteCalls())
func (mock *StoreMock) DeleteCalls() []struct {
	ContextMoqParam context.Context
	S               string
} {
	var calls []struct {
		ContextMoqParam context.Context
		S               string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *StoreMock) Get(contextMoqParam context.Context, s string) (*companies.Company, error) {
	if mock.GetFunc == nil {
		panic("StoreMock.GetFunc: method is nil but Store.Get was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		S               string
	}{
		ContextMoqParam: contextMoqParam,
		S:               s,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(contextMoqParam, s)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedStore.GetCalls())
func (mock *StoreMock) GetCalls() []struct {
	ContextMoqParam context.Context
	S               string
} {
	var calls []struct {
		ContextMoqParam context.Context
		S               string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *StoreMock) Update(contextMoqParam context.Context, s string, company *companies.Company) error {
	if mock.UpdateFunc == nil {
		panic("StoreMock.UpdateFunc: method is nil but Store.Update was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		S               string
		Company         *companies.Company
	}{
		ContextMoqParam: contextMoqParam,
		S:               s,
		Company:         company,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(contextMoqParam, s, company)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedStore.UpdateCalls())
func (mock *StoreMock) UpdateCalls() []struct {
	ContextMoqParam context.Context
	S               string
	Company         *companies.Company
} {
	var calls []struct {
		ContextMoqParam context.Context
		S               string
		Company         *companies.Company
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
