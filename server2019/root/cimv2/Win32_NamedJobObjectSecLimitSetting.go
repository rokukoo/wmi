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

// Win32_NamedJobObjectSecLimitSetting struct
type Win32_NamedJobObjectSecLimitSetting struct {
	*CIM_Setting

	//
	PrivilegesToDelete Win32_TokenPrivileges

	//
	RestrictedSIDs Win32_TokenGroups

	//
	SecurityLimitFlags uint32

	//
	SIDsToDisable Win32_TokenGroups
}

func NewWin32_NamedJobObjectSecLimitSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_NamedJobObjectSecLimitSetting, err error) {
	tmp, err := NewCIM_SettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_NamedJobObjectSecLimitSetting{
		CIM_Setting: tmp,
	}
	return
}

func NewWin32_NamedJobObjectSecLimitSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_NamedJobObjectSecLimitSetting, err error) {
	tmp, err := NewCIM_SettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_NamedJobObjectSecLimitSetting{
		CIM_Setting: tmp,
	}
	return
}

// SetPrivilegesToDelete sets the value of PrivilegesToDelete for the instance
func (instance *Win32_NamedJobObjectSecLimitSetting) SetPropertyPrivilegesToDelete(value Win32_TokenPrivileges) (err error) {
	return instance.SetProperty("PrivilegesToDelete", value)
}

// GetPrivilegesToDelete gets the value of PrivilegesToDelete for the instance
func (instance *Win32_NamedJobObjectSecLimitSetting) GetPropertyPrivilegesToDelete() (value Win32_TokenPrivileges, err error) {
	retValue, err := instance.GetProperty("PrivilegesToDelete")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_TokenPrivileges)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestrictedSIDs sets the value of RestrictedSIDs for the instance
func (instance *Win32_NamedJobObjectSecLimitSetting) SetPropertyRestrictedSIDs(value Win32_TokenGroups) (err error) {
	return instance.SetProperty("RestrictedSIDs", value)
}

// GetRestrictedSIDs gets the value of RestrictedSIDs for the instance
func (instance *Win32_NamedJobObjectSecLimitSetting) GetPropertyRestrictedSIDs() (value Win32_TokenGroups, err error) {
	retValue, err := instance.GetProperty("RestrictedSIDs")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_TokenGroups)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecurityLimitFlags sets the value of SecurityLimitFlags for the instance
func (instance *Win32_NamedJobObjectSecLimitSetting) SetPropertySecurityLimitFlags(value uint32) (err error) {
	return instance.SetProperty("SecurityLimitFlags", value)
}

// GetSecurityLimitFlags gets the value of SecurityLimitFlags for the instance
func (instance *Win32_NamedJobObjectSecLimitSetting) GetPropertySecurityLimitFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecurityLimitFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSIDsToDisable sets the value of SIDsToDisable for the instance
func (instance *Win32_NamedJobObjectSecLimitSetting) SetPropertySIDsToDisable(value Win32_TokenGroups) (err error) {
	return instance.SetProperty("SIDsToDisable", value)
}

// GetSIDsToDisable gets the value of SIDsToDisable for the instance
func (instance *Win32_NamedJobObjectSecLimitSetting) GetPropertySIDsToDisable() (value Win32_TokenGroups, err error) {
	retValue, err := instance.GetProperty("SIDsToDisable")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_TokenGroups)
	if !ok {
		// TODO: Set an error
	}
	return
}
