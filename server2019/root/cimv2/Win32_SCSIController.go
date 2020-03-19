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

// Win32_SCSIController struct
type Win32_SCSIController struct {
	*CIM_SCSIController

	//
	DeviceMap string

	//
	DriverName string

	//
	HardwareVersion string

	//
	Index uint32

	//
	Manufacturer string
}

func NewWin32_SCSIControllerEx1(instance *cim.WmiInstance) (newInstance *Win32_SCSIController, err error) {
	tmp, err := NewCIM_SCSIControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_SCSIController{
		CIM_SCSIController: tmp,
	}
	return
}

func NewWin32_SCSIControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_SCSIController, err error) {
	tmp, err := NewCIM_SCSIControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_SCSIController{
		CIM_SCSIController: tmp,
	}
	return
}

// SetDeviceMap sets the value of DeviceMap for the instance
func (instance *Win32_SCSIController) SetPropertyDeviceMap(value string) (err error) {
	return instance.SetProperty("DeviceMap", value)
}

// GetDeviceMap gets the value of DeviceMap for the instance
func (instance *Win32_SCSIController) GetPropertyDeviceMap() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceMap")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverName sets the value of DriverName for the instance
func (instance *Win32_SCSIController) SetPropertyDriverName(value string) (err error) {
	return instance.SetProperty("DriverName", value)
}

// GetDriverName gets the value of DriverName for the instance
func (instance *Win32_SCSIController) GetPropertyDriverName() (value string, err error) {
	retValue, err := instance.GetProperty("DriverName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHardwareVersion sets the value of HardwareVersion for the instance
func (instance *Win32_SCSIController) SetPropertyHardwareVersion(value string) (err error) {
	return instance.SetProperty("HardwareVersion", value)
}

// GetHardwareVersion gets the value of HardwareVersion for the instance
func (instance *Win32_SCSIController) GetPropertyHardwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("HardwareVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIndex sets the value of Index for the instance
func (instance *Win32_SCSIController) SetPropertyIndex(value uint32) (err error) {
	return instance.SetProperty("Index", value)
}

// GetIndex gets the value of Index for the instance
func (instance *Win32_SCSIController) GetPropertyIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("Index")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *Win32_SCSIController) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *Win32_SCSIController) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
