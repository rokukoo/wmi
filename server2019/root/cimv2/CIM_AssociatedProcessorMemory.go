// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_AssociatedProcessorMemory struct
type CIM_AssociatedProcessorMemory struct {
	*CIM_AssociatedMemory

	//
	BusSpeed uint32
}

func NewCIM_AssociatedProcessorMemoryEx1(instance *cim.WmiInstance) (newInstance *CIM_AssociatedProcessorMemory, err error) {
	tmp, err := NewCIM_AssociatedMemoryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_AssociatedProcessorMemory{
		CIM_AssociatedMemory: tmp,
	}
	return
}

func NewCIM_AssociatedProcessorMemoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_AssociatedProcessorMemory, err error) {
	tmp, err := NewCIM_AssociatedMemoryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_AssociatedProcessorMemory{
		CIM_AssociatedMemory: tmp,
	}
	return
}

// SetBusSpeed sets the value of BusSpeed for the instance
func (instance *CIM_AssociatedProcessorMemory) SetPropertyBusSpeed(value uint32) (err error) {
	return instance.SetProperty("BusSpeed", value)
}

// GetBusSpeed gets the value of BusSpeed for the instance
func (instance *CIM_AssociatedProcessorMemory) GetPropertyBusSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
