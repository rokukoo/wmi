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

// CIM_Keyboard struct
type CIM_Keyboard struct {
	*CIM_UserDevice

	//
	Layout string

	//
	NumberOfFunctionKeys uint16

	//
	Password uint16
}

func NewCIM_KeyboardEx1(instance *cim.WmiInstance) (newInstance *CIM_Keyboard, err error) {
	tmp, err := NewCIM_UserDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Keyboard{
		CIM_UserDevice: tmp,
	}
	return
}

func NewCIM_KeyboardEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Keyboard, err error) {
	tmp, err := NewCIM_UserDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Keyboard{
		CIM_UserDevice: tmp,
	}
	return
}

// SetLayout sets the value of Layout for the instance
func (instance *CIM_Keyboard) SetPropertyLayout(value string) (err error) {
	return instance.SetProperty("Layout", value)
}

// GetLayout gets the value of Layout for the instance
func (instance *CIM_Keyboard) GetPropertyLayout() (value string, err error) {
	retValue, err := instance.GetProperty("Layout")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfFunctionKeys sets the value of NumberOfFunctionKeys for the instance
func (instance *CIM_Keyboard) SetPropertyNumberOfFunctionKeys(value uint16) (err error) {
	return instance.SetProperty("NumberOfFunctionKeys", value)
}

// GetNumberOfFunctionKeys gets the value of NumberOfFunctionKeys for the instance
func (instance *CIM_Keyboard) GetPropertyNumberOfFunctionKeys() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfFunctionKeys")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPassword sets the value of Password for the instance
func (instance *CIM_Keyboard) SetPropertyPassword(value uint16) (err error) {
	return instance.SetProperty("Password", value)
}

// GetPassword gets the value of Password for the instance
func (instance *CIM_Keyboard) GetPropertyPassword() (value uint16, err error) {
	retValue, err := instance.GetProperty("Password")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
