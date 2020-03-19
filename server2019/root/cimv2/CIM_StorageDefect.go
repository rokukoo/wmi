// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_StorageDefect struct
type CIM_StorageDefect struct {
	*cim.WmiInstance

	//
	Error CIM_StorageError

	//
	Extent CIM_StorageExtent
}

func NewCIM_StorageDefectEx1(instance *cim.WmiInstance) (newInstance *CIM_StorageDefect, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_StorageDefect{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_StorageDefectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_StorageDefect, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_StorageDefect{
		WmiInstance: tmp,
	}
	return
}

// SetError sets the value of Error for the instance
func (instance *CIM_StorageDefect) SetPropertyError(value CIM_StorageError) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *CIM_StorageDefect) GetPropertyError() (value CIM_StorageError, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_StorageError)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExtent sets the value of Extent for the instance
func (instance *CIM_StorageDefect) SetPropertyExtent(value CIM_StorageExtent) (err error) {
	return instance.SetProperty("Extent", value)
}

// GetExtent gets the value of Extent for the instance
func (instance *CIM_StorageDefect) GetPropertyExtent() (value CIM_StorageExtent, err error) {
	retValue, err := instance.GetProperty("Extent")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_StorageExtent)
	if !ok {
		// TODO: Set an error
	}
	return
}
