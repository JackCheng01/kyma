// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import spec "github.com/kyma-project/kyma/components/console-backend-service/internal/domain/assetstore/spec"

// specificationSvc is an autogenerated failing mock type for the specificationSvc type
type specificationSvc struct {
	err error
}

// NewSpecificationSvc creates a new specificationSvc type instance
func NewSpecificationSvc(err error) *specificationSvc {
	return &specificationSvc{err: err}
}

// AsyncAPI provides a failing mock function with given fields: baseURL, name
func (_m *specificationSvc) AsyncAPI(baseURL string, name string) (*spec.AsyncAPISpec, error) {
	var r0 *spec.AsyncAPISpec
	var r1 error
	r1 = _m.err

	return r0, r1
}
