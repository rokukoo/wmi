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

// Win32_Battery struct
type Win32_Battery struct {
	*CIM_Battery

	//
	BatteryRechargeTime uint32

	//
	ExpectedBatteryLife uint32
}

func NewWin32_BatteryEx1(instance *cim.WmiInstance) (newInstance *Win32_Battery, err error) {
	tmp, err := NewCIM_BatteryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_Battery{
		CIM_Battery: tmp,
	}
	return
}

func NewWin32_BatteryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_Battery, err error) {
	tmp, err := NewCIM_BatteryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_Battery{
		CIM_Battery: tmp,
	}
	return
}

// SetBatteryRechargeTime sets the value of BatteryRechargeTime for the instance
func (instance *Win32_Battery) SetPropertyBatteryRechargeTime(value uint32) (err error) {
	return instance.SetProperty("BatteryRechargeTime", value)
}

// GetBatteryRechargeTime gets the value of BatteryRechargeTime for the instance
func (instance *Win32_Battery) GetPropertyBatteryRechargeTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("BatteryRechargeTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpectedBatteryLife sets the value of ExpectedBatteryLife for the instance
func (instance *Win32_Battery) SetPropertyExpectedBatteryLife(value uint32) (err error) {
	return instance.SetProperty("ExpectedBatteryLife", value)
}

// GetExpectedBatteryLife gets the value of ExpectedBatteryLife for the instance
func (instance *Win32_Battery) GetPropertyExpectedBatteryLife() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExpectedBatteryLife")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
