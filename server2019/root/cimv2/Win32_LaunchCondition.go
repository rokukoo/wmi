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

// Win32_LaunchCondition struct
type Win32_LaunchCondition struct {
	*CIM_Check

	//
	Condition string
}

func NewWin32_LaunchConditionEx1(instance *cim.WmiInstance) (newInstance *Win32_LaunchCondition, err error) {
	tmp, err := NewCIM_CheckEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_LaunchCondition{
		CIM_Check: tmp,
	}
	return
}

func NewWin32_LaunchConditionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_LaunchCondition, err error) {
	tmp, err := NewCIM_CheckEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_LaunchCondition{
		CIM_Check: tmp,
	}
	return
}

// SetCondition sets the value of Condition for the instance
func (instance *Win32_LaunchCondition) SetPropertyCondition(value string) (err error) {
	return instance.SetProperty("Condition", value)
}

// GetCondition gets the value of Condition for the instance
func (instance *Win32_LaunchCondition) GetPropertyCondition() (value string, err error) {
	retValue, err := instance.GetProperty("Condition")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
