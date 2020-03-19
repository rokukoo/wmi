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

// Win32_PnPDevicePropertyUint64 struct
type Win32_PnPDevicePropertyUint64 struct {
	*Win32_PnPDeviceProperty

	//
	Data uint64
}

func NewWin32_PnPDevicePropertyUint64Ex1(instance *cim.WmiInstance) (newInstance *Win32_PnPDevicePropertyUint64, err error) {
	tmp, err := NewWin32_PnPDevicePropertyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PnPDevicePropertyUint64{
		Win32_PnPDeviceProperty: tmp,
	}
	return
}

func NewWin32_PnPDevicePropertyUint64Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PnPDevicePropertyUint64, err error) {
	tmp, err := NewWin32_PnPDevicePropertyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PnPDevicePropertyUint64{
		Win32_PnPDeviceProperty: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyUint64) SetPropertyData(value uint64) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyUint64) GetPropertyData() (value uint64, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
