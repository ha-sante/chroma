// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	dbmodel "github.com/chroma-core/chroma/go/pkg/metastore/db/dbmodel"

	mock "github.com/stretchr/testify/mock"

	model "github.com/chroma-core/chroma/go/pkg/model"

	types "github.com/chroma-core/chroma/go/pkg/types"
)

// Catalog is an autogenerated mock type for the Catalog type
type Catalog struct {
	mock.Mock
}

// CreateCollection provides a mock function with given fields: ctx, createCollection, ts
func (_m *Catalog) CreateCollection(ctx context.Context, createCollection *model.CreateCollection, ts int64) (*model.Collection, error) {
	ret := _m.Called(ctx, createCollection, ts)

	if len(ret) == 0 {
		panic("no return value specified for CreateCollection")
	}

	var r0 *model.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateCollection, int64) (*model.Collection, error)); ok {
		return rf(ctx, createCollection, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateCollection, int64) *model.Collection); ok {
		r0 = rf(ctx, createCollection, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.CreateCollection, int64) error); ok {
		r1 = rf(ctx, createCollection, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDatabase provides a mock function with given fields: ctx, createDatabase, ts
func (_m *Catalog) CreateDatabase(ctx context.Context, createDatabase *model.CreateDatabase, ts int64) (*model.Database, error) {
	ret := _m.Called(ctx, createDatabase, ts)

	if len(ret) == 0 {
		panic("no return value specified for CreateDatabase")
	}

	var r0 *model.Database
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateDatabase, int64) (*model.Database, error)); ok {
		return rf(ctx, createDatabase, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateDatabase, int64) *model.Database); ok {
		r0 = rf(ctx, createDatabase, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Database)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.CreateDatabase, int64) error); ok {
		r1 = rf(ctx, createDatabase, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSegment provides a mock function with given fields: ctx, createSegment, ts
func (_m *Catalog) CreateSegment(ctx context.Context, createSegment *model.CreateSegment, ts int64) (*model.Segment, error) {
	ret := _m.Called(ctx, createSegment, ts)

	if len(ret) == 0 {
		panic("no return value specified for CreateSegment")
	}

	var r0 *model.Segment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateSegment, int64) (*model.Segment, error)); ok {
		return rf(ctx, createSegment, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateSegment, int64) *model.Segment); ok {
		r0 = rf(ctx, createSegment, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Segment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.CreateSegment, int64) error); ok {
		r1 = rf(ctx, createSegment, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTenant provides a mock function with given fields: ctx, createTenant, ts
func (_m *Catalog) CreateTenant(ctx context.Context, createTenant *model.CreateTenant, ts int64) (*model.Tenant, error) {
	ret := _m.Called(ctx, createTenant, ts)

	if len(ret) == 0 {
		panic("no return value specified for CreateTenant")
	}

	var r0 *model.Tenant
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateTenant, int64) (*model.Tenant, error)); ok {
		return rf(ctx, createTenant, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateTenant, int64) *model.Tenant); ok {
		r0 = rf(ctx, createTenant, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Tenant)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.CreateTenant, int64) error); ok {
		r1 = rf(ctx, createTenant, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCollection provides a mock function with given fields: ctx, deleteCollection
func (_m *Catalog) DeleteCollection(ctx context.Context, deleteCollection *model.DeleteCollection) error {
	ret := _m.Called(ctx, deleteCollection)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCollection")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.DeleteCollection) error); ok {
		r0 = rf(ctx, deleteCollection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSegment provides a mock function with given fields: ctx, segmentID
func (_m *Catalog) DeleteSegment(ctx context.Context, segmentID types.UniqueID) error {
	ret := _m.Called(ctx, segmentID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSegment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.UniqueID) error); ok {
		r0 = rf(ctx, segmentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlushCollectionCompaction provides a mock function with given fields: ctx, flushCollectionCompaction
func (_m *Catalog) FlushCollectionCompaction(ctx context.Context, flushCollectionCompaction *model.FlushCollectionCompaction) (*model.FlushCollectionInfo, error) {
	ret := _m.Called(ctx, flushCollectionCompaction)

	if len(ret) == 0 {
		panic("no return value specified for FlushCollectionCompaction")
	}

	var r0 *model.FlushCollectionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.FlushCollectionCompaction) (*model.FlushCollectionInfo, error)); ok {
		return rf(ctx, flushCollectionCompaction)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.FlushCollectionCompaction) *model.FlushCollectionInfo); ok {
		r0 = rf(ctx, flushCollectionCompaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FlushCollectionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.FlushCollectionCompaction) error); ok {
		r1 = rf(ctx, flushCollectionCompaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllDatabases provides a mock function with given fields: ctx, ts
func (_m *Catalog) GetAllDatabases(ctx context.Context, ts int64) ([]*model.Database, error) {
	ret := _m.Called(ctx, ts)

	if len(ret) == 0 {
		panic("no return value specified for GetAllDatabases")
	}

	var r0 []*model.Database
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*model.Database, error)); ok {
		return rf(ctx, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*model.Database); ok {
		r0 = rf(ctx, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Database)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTenants provides a mock function with given fields: ctx, ts
func (_m *Catalog) GetAllTenants(ctx context.Context, ts int64) ([]*model.Tenant, error) {
	ret := _m.Called(ctx, ts)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTenants")
	}

	var r0 []*model.Tenant
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*model.Tenant, error)); ok {
		return rf(ctx, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*model.Tenant); ok {
		r0 = rf(ctx, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Tenant)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollections provides a mock function with given fields: ctx, collectionID, collectionName, tenantID, databaseName
func (_m *Catalog) GetCollections(ctx context.Context, collectionID types.UniqueID, collectionName *string, tenantID string, databaseName string) ([]*model.Collection, error) {
	ret := _m.Called(ctx, collectionID, collectionName, tenantID, databaseName)

	if len(ret) == 0 {
		panic("no return value specified for GetCollections")
	}

	var r0 []*model.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.UniqueID, *string, string, string) ([]*model.Collection, error)); ok {
		return rf(ctx, collectionID, collectionName, tenantID, databaseName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.UniqueID, *string, string, string) []*model.Collection); ok {
		r0 = rf(ctx, collectionID, collectionName, tenantID, databaseName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.UniqueID, *string, string, string) error); ok {
		r1 = rf(ctx, collectionID, collectionName, tenantID, databaseName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDatabases provides a mock function with given fields: ctx, getDatabase, ts
func (_m *Catalog) GetDatabases(ctx context.Context, getDatabase *model.GetDatabase, ts int64) (*model.Database, error) {
	ret := _m.Called(ctx, getDatabase, ts)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabases")
	}

	var r0 *model.Database
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.GetDatabase, int64) (*model.Database, error)); ok {
		return rf(ctx, getDatabase, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.GetDatabase, int64) *model.Database); ok {
		r0 = rf(ctx, getDatabase, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Database)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.GetDatabase, int64) error); ok {
		r1 = rf(ctx, getDatabase, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSegments provides a mock function with given fields: ctx, segmentID, segmentType, scope, collectionID
func (_m *Catalog) GetSegments(ctx context.Context, segmentID types.UniqueID, segmentType *string, scope *string, collectionID types.UniqueID) ([]*model.Segment, error) {
	ret := _m.Called(ctx, segmentID, segmentType, scope, collectionID)

	if len(ret) == 0 {
		panic("no return value specified for GetSegments")
	}

	var r0 []*model.Segment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.UniqueID, *string, *string, types.UniqueID) ([]*model.Segment, error)); ok {
		return rf(ctx, segmentID, segmentType, scope, collectionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.UniqueID, *string, *string, types.UniqueID) []*model.Segment); ok {
		r0 = rf(ctx, segmentID, segmentType, scope, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Segment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.UniqueID, *string, *string, types.UniqueID) error); ok {
		r1 = rf(ctx, segmentID, segmentType, scope, collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTenants provides a mock function with given fields: ctx, getTenant, ts
func (_m *Catalog) GetTenants(ctx context.Context, getTenant *model.GetTenant, ts int64) (*model.Tenant, error) {
	ret := _m.Called(ctx, getTenant, ts)

	if len(ret) == 0 {
		panic("no return value specified for GetTenants")
	}

	var r0 *model.Tenant
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.GetTenant, int64) (*model.Tenant, error)); ok {
		return rf(ctx, getTenant, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.GetTenant, int64) *model.Tenant); ok {
		r0 = rf(ctx, getTenant, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Tenant)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.GetTenant, int64) error); ok {
		r1 = rf(ctx, getTenant, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTenantsLastCompactionTime provides a mock function with given fields: ctx, tenantIDs
func (_m *Catalog) GetTenantsLastCompactionTime(ctx context.Context, tenantIDs []string) ([]*dbmodel.Tenant, error) {
	ret := _m.Called(ctx, tenantIDs)

	if len(ret) == 0 {
		panic("no return value specified for GetTenantsLastCompactionTime")
	}

	var r0 []*dbmodel.Tenant
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]*dbmodel.Tenant, error)); ok {
		return rf(ctx, tenantIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*dbmodel.Tenant); ok {
		r0 = rf(ctx, tenantIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.Tenant)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, tenantIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetState provides a mock function with given fields: ctx
func (_m *Catalog) ResetState(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ResetState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTenantLastCompactionTime provides a mock function with given fields: ctx, tenantID, lastCompactionTime
func (_m *Catalog) SetTenantLastCompactionTime(ctx context.Context, tenantID string, lastCompactionTime int64) error {
	ret := _m.Called(ctx, tenantID, lastCompactionTime)

	if len(ret) == 0 {
		panic("no return value specified for SetTenantLastCompactionTime")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, tenantID, lastCompactionTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCollection provides a mock function with given fields: ctx, updateCollection, ts
func (_m *Catalog) UpdateCollection(ctx context.Context, updateCollection *model.UpdateCollection, ts int64) (*model.Collection, error) {
	ret := _m.Called(ctx, updateCollection, ts)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCollection")
	}

	var r0 *model.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UpdateCollection, int64) (*model.Collection, error)); ok {
		return rf(ctx, updateCollection, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.UpdateCollection, int64) *model.Collection); ok {
		r0 = rf(ctx, updateCollection, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.UpdateCollection, int64) error); ok {
		r1 = rf(ctx, updateCollection, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSegment provides a mock function with given fields: ctx, segmentInfo, ts
func (_m *Catalog) UpdateSegment(ctx context.Context, segmentInfo *model.UpdateSegment, ts int64) (*model.Segment, error) {
	ret := _m.Called(ctx, segmentInfo, ts)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSegment")
	}

	var r0 *model.Segment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UpdateSegment, int64) (*model.Segment, error)); ok {
		return rf(ctx, segmentInfo, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.UpdateSegment, int64) *model.Segment); ok {
		r0 = rf(ctx, segmentInfo, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Segment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.UpdateSegment, int64) error); ok {
		r1 = rf(ctx, segmentInfo, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCatalog creates a new instance of Catalog. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCatalog(t interface {
	mock.TestingT
	Cleanup(func())
}) *Catalog {
	mock := &Catalog{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
