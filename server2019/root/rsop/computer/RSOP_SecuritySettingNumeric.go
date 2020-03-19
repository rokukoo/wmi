// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_SecuritySettingNumeric struct
type RSOP_SecuritySettingNumeric struct {
	*RSOP_SecuritySettings

	//
	KeyName string

	//
	Setting uint32
}

func NewRSOP_SecuritySettingNumericEx1(instance *cim.WmiInstance) (newInstance *RSOP_SecuritySettingNumeric, err error) {
	tmp, err := NewRSOP_SecuritySettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_SecuritySettingNumeric{
		RSOP_SecuritySettings: tmp,
	}
	return
}

func NewRSOP_SecuritySettingNumericEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_SecuritySettingNumeric, err error) {
	tmp, err := NewRSOP_SecuritySettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_SecuritySettingNumeric{
		RSOP_SecuritySettings: tmp,
	}
	return
}

// SetKeyName sets the value of KeyName for the instance
func (instance *RSOP_SecuritySettingNumeric) SetPropertyKeyName(value string) (err error) {
	return instance.SetProperty("KeyName", value)
}

// GetKeyName gets the value of KeyName for the instance
func (instance *RSOP_SecuritySettingNumeric) GetPropertyKeyName() (value string, err error) {
	retValue, err := instance.GetProperty("KeyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetting sets the value of Setting for the instance
func (instance *RSOP_SecuritySettingNumeric) SetPropertySetting(value uint32) (err error) {
	return instance.SetProperty("Setting", value)
}

// GetSetting gets the value of Setting for the instance
func (instance *RSOP_SecuritySettingNumeric) GetPropertySetting() (value uint32, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}