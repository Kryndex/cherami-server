// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package metadata

import "github.com/uber/cherami-thrift/.generated/go/shared"
import "github.com/uber/cherami-thrift/.generated/go/metadata"
import "github.com/stretchr/testify/mock"

// MetadataExposable is an autogenerated mock type for the MetadataExposable type
type MetadataExposable struct {
	mock.Mock
}

// HostAddrToUUID provides a mock function with given fields: hostAddr
func (_m *MetadataExposable) HostAddrToUUID(hostAddr string) (string, error) {
	ret := _m.Called(hostAddr)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(hostAddr)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hostAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExtentsStats provides a mock function with given fields: request
func (_m *MetadataExposable) ListExtentsStats(request *shared.ListExtentsStatsRequest) (*shared.ListExtentsStatsResult_, error) {
	ret := _m.Called(request)

	var r0 *shared.ListExtentsStatsResult_
	if rf, ok := ret.Get(0).(func(*shared.ListExtentsStatsRequest) *shared.ListExtentsStatsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListExtentsStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shared.ListExtentsStatsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHosts provides a mock function with given fields: request
func (_m *MetadataExposable) ListHosts(request *metadata.ListHostsRequest) (*metadata.ListHostsResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ListHostsResult_
	if rf, ok := ret.Get(0).(func(*metadata.ListHostsRequest) *metadata.ListHostsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListHostsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ListHostsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInputHostExtentsStats provides a mock function with given fields: request
func (_m *MetadataExposable) ListInputHostExtentsStats(request *metadata.ListInputHostExtentsStatsRequest) (*metadata.ListInputHostExtentsStatsResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ListInputHostExtentsStatsResult_
	if rf, ok := ret.Get(0).(func(*metadata.ListInputHostExtentsStatsRequest) *metadata.ListInputHostExtentsStatsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListInputHostExtentsStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ListInputHostExtentsStatsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStoreExtentsStats provides a mock function with given fields: request
func (_m *MetadataExposable) ListStoreExtentsStats(request *metadata.ListStoreExtentsStatsRequest) (*metadata.ListStoreExtentsStatsResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ListStoreExtentsStatsResult_
	if rf, ok := ret.Get(0).(func(*metadata.ListStoreExtentsStatsRequest) *metadata.ListStoreExtentsStatsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListStoreExtentsStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ListStoreExtentsStatsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadConsumerGroupExtent provides a mock function with given fields: request
func (_m *MetadataExposable) ReadConsumerGroupExtent(request *metadata.ReadConsumerGroupExtentRequest) (*metadata.ReadConsumerGroupExtentResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ReadConsumerGroupExtentResult_
	if rf, ok := ret.Get(0).(func(*metadata.ReadConsumerGroupExtentRequest) *metadata.ReadConsumerGroupExtentResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadConsumerGroupExtentResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ReadConsumerGroupExtentRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadConsumerGroupExtents provides a mock function with given fields: request
func (_m *MetadataExposable) ReadConsumerGroupExtents(request *shared.ReadConsumerGroupExtentsRequest) (*shared.ReadConsumerGroupExtentsResult_, error) {
	ret := _m.Called(request)

	var r0 *shared.ReadConsumerGroupExtentsResult_
	if rf, ok := ret.Get(0).(func(*shared.ReadConsumerGroupExtentsRequest) *shared.ReadConsumerGroupExtentsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ReadConsumerGroupExtentsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shared.ReadConsumerGroupExtentsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadExtentStats provides a mock function with given fields: request
func (_m *MetadataExposable) ReadExtentStats(request *metadata.ReadExtentStatsRequest) (*metadata.ReadExtentStatsResult_, error) {
	ret := _m.Called(request)

	var r0 *metadata.ReadExtentStatsResult_
	if rf, ok := ret.Get(0).(func(*metadata.ReadExtentStatsRequest) *metadata.ReadExtentStatsResult_); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadExtentStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*metadata.ReadExtentStatsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UUIDToHostAddr provides a mock function with given fields: hostUUID
func (_m *MetadataExposable) UUIDToHostAddr(hostUUID string) (string, error) {
	ret := _m.Called(hostUUID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(hostUUID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hostUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
