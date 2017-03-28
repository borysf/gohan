// Copyright (C) 2015 NTT Innovation Institute, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mocks

import (
	"github.com/cloudwan/gohan/db/pagination"
	"github.com/cloudwan/gohan/db/transaction"
	"github.com/cloudwan/gohan/schema"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

// Transaction mock
type Transaction struct {
	mock.Mock
}

// Create mock
func (_m *Transaction) Create(_a0 *schema.Resource) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*schema.Resource) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update mock
func (_m *Transaction) Update(_a0 *schema.Resource) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*schema.Resource) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StateUpdate mock
func (_m *Transaction) StateUpdate(_a0 *schema.Resource, _a1 *transaction.ResourceState) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*schema.Resource, *transaction.ResourceState) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete mock
func (_m *Transaction) Delete(_a0 *schema.Schema, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*schema.Schema, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch mock
func (_m *Transaction) Fetch(_a0 *schema.Schema, _a1 transaction.Filter) (*schema.Resource, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *schema.Resource
	if rf, ok := ret.Get(0).(func(*schema.Schema, transaction.Filter) *schema.Resource); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*schema.Schema, transaction.Filter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LockFetch provides a mock function with given fields: _a0, _a1, _a2
func (_m *Transaction) LockFetch(_a0 *schema.Schema, _a1 transaction.Filter, _a2 schema.LockPolicy) (*schema.Resource, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *schema.Resource
	if rf, ok := ret.Get(0).(func(*schema.Schema, transaction.Filter, schema.LockPolicy) *schema.Resource); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*schema.Schema, transaction.Filter, schema.LockPolicy) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateFetch mock
func (_m *Transaction) StateFetch(_a0 *schema.Schema, _a1 transaction.Filter) (transaction.ResourceState, error) {
	ret := _m.Called(_a0, _a1)

	var r0 transaction.ResourceState
	if rf, ok := ret.Get(0).(func(*schema.Schema, transaction.Filter) transaction.ResourceState); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(transaction.ResourceState)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*schema.Schema, transaction.Filter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List mock
func (_m *Transaction) List(_a0 *schema.Schema, _a1 transaction.Filter, _a2 *pagination.Paginator) ([]*schema.Resource, uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []*schema.Resource
	if rf, ok := ret.Get(0).(func(*schema.Schema, transaction.Filter, *pagination.Paginator) []*schema.Resource); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*schema.Resource)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(*schema.Schema, transaction.Filter, *pagination.Paginator) uint64); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*schema.Schema, transaction.Filter, *pagination.Paginator) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LockList provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Transaction) LockList(_a0 *schema.Schema, _a1 transaction.Filter, _a2 *pagination.Paginator, _a3 schema.LockPolicy) ([]*schema.Resource, uint64, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []*schema.Resource
	if rf, ok := ret.Get(0).(func(*schema.Schema, transaction.Filter, *pagination.Paginator, schema.LockPolicy) []*schema.Resource); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*schema.Resource)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(*schema.Schema, transaction.Filter, *pagination.Paginator, schema.LockPolicy) uint64); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*schema.Schema, transaction.Filter, *pagination.Paginator, schema.LockPolicy) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RawTransaction mock
func (_m *Transaction) RawTransaction() *sqlx.Tx {
	ret := _m.Called()

	var r0 *sqlx.Tx
	if rf, ok := ret.Get(0).(func() *sqlx.Tx); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Tx)
		}
	}

	return r0
}

// Query mock
func (_m *Transaction) Query(_a0 *schema.Schema, _a1 string, _a2 []interface{}) ([]*schema.Resource, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []*schema.Resource
	if rf, ok := ret.Get(0).(func(*schema.Schema, string, []interface{}) []*schema.Resource); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*schema.Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*schema.Schema, string, []interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit mock
func (_m *Transaction) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close mock
func (_m *Transaction) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Closed mock
func (_m *Transaction) Closed() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//SetIsolationLevel specify transaction isolation level
func (_m *Transaction) SetIsolationLevel(_a0 transaction.Type) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(_a0 transaction.Type) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exec executes sql in transaction
func (_m *Transaction) Exec(sql string, args ...interface{}) error {
	return nil
}
