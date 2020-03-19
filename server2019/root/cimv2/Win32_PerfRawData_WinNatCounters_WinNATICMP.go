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

// Win32_PerfRawData_WinNatCounters_WinNATICMP struct
type Win32_PerfRawData_WinNatCounters_WinNATICMP struct {
	*Win32_PerfRawData

	//
	NumberOfBindings uint32

	//
	NumberOfSessions uint32

	//
	NumExtToIntTranslations uint32

	//
	NumIntToExtTranslations uint32

	//
	NumPacketsDropped uint32

	//
	NumSessionsTimedOut uint32
}

func NewWin32_PerfRawData_WinNatCounters_WinNATICMPEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_WinNatCounters_WinNATICMP, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_WinNatCounters_WinNATICMP{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_WinNatCounters_WinNATICMPEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_WinNatCounters_WinNATICMP, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_WinNatCounters_WinNATICMP{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetNumberOfBindings sets the value of NumberOfBindings for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) SetPropertyNumberOfBindings(value uint32) (err error) {
	return instance.SetProperty("NumberOfBindings", value)
}

// GetNumberOfBindings gets the value of NumberOfBindings for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) GetPropertyNumberOfBindings() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfBindings")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfSessions sets the value of NumberOfSessions for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) SetPropertyNumberOfSessions(value uint32) (err error) {
	return instance.SetProperty("NumberOfSessions", value)
}

// GetNumberOfSessions gets the value of NumberOfSessions for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) GetPropertyNumberOfSessions() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfSessions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumExtToIntTranslations sets the value of NumExtToIntTranslations for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) SetPropertyNumExtToIntTranslations(value uint32) (err error) {
	return instance.SetProperty("NumExtToIntTranslations", value)
}

// GetNumExtToIntTranslations gets the value of NumExtToIntTranslations for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) GetPropertyNumExtToIntTranslations() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumExtToIntTranslations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumIntToExtTranslations sets the value of NumIntToExtTranslations for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) SetPropertyNumIntToExtTranslations(value uint32) (err error) {
	return instance.SetProperty("NumIntToExtTranslations", value)
}

// GetNumIntToExtTranslations gets the value of NumIntToExtTranslations for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) GetPropertyNumIntToExtTranslations() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumIntToExtTranslations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumPacketsDropped sets the value of NumPacketsDropped for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) SetPropertyNumPacketsDropped(value uint32) (err error) {
	return instance.SetProperty("NumPacketsDropped", value)
}

// GetNumPacketsDropped gets the value of NumPacketsDropped for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) GetPropertyNumPacketsDropped() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumPacketsDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumSessionsTimedOut sets the value of NumSessionsTimedOut for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) SetPropertyNumSessionsTimedOut(value uint32) (err error) {
	return instance.SetProperty("NumSessionsTimedOut", value)
}

// GetNumSessionsTimedOut gets the value of NumSessionsTimedOut for the instance
func (instance *Win32_PerfRawData_WinNatCounters_WinNATICMP) GetPropertyNumSessionsTimedOut() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumSessionsTimedOut")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
