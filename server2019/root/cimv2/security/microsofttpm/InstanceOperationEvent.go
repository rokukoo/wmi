// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.Security.MicrosoftTpm
//////////////////////////////////////////////
package microsofttpm

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __InstanceOperationEvent struct
type __InstanceOperationEvent struct {
	*__Event

	//
	TargetInstance interface{}
}

func New__InstanceOperationEventEx1(instance *cim.WmiInstance) (newInstance *__InstanceOperationEvent, err error) {
	tmp, err := New__EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__InstanceOperationEvent{
		__Event: tmp,
	}
	return
}

func New__InstanceOperationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__InstanceOperationEvent, err error) {
	tmp, err := New__EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__InstanceOperationEvent{
		__Event: tmp,
	}
	return
}

// SetTargetInstance sets the value of TargetInstance for the instance
func (instance *__InstanceOperationEvent) SetPropertyTargetInstance(value interface{}) (err error) {
	return instance.SetProperty("TargetInstance", value)
}

// GetTargetInstance gets the value of TargetInstance for the instance
func (instance *__InstanceOperationEvent) GetPropertyTargetInstance() (value interface{}, err error) {
	retValue, err := instance.GetProperty("TargetInstance")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}