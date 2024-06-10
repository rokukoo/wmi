// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.TerminalServices
//
// ////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	cimv2 "github.com/microsoft/wmi/server2019/root/cimv2"
	"reflect"
)

// Win32_TSSessionDirectorySetting struct
type Win32_TSSessionDirectorySetting struct {
	*CIM_ElementSetting

	//
	Element cimv2.Win32_TerminalService

	//
	Setting Win32_TSSessionDirectory
}

func NewWin32_TSSessionDirectorySettingEx1(instance *cim.WmiInstance) (newInstance *Win32_TSSessionDirectorySetting, err error) {
	tmp, err := NewCIM_ElementSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSSessionDirectorySetting{
		CIM_ElementSetting: tmp,
	}
	return
}

func NewWin32_TSSessionDirectorySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSSessionDirectorySetting, err error) {
	tmp, err := NewCIM_ElementSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSSessionDirectorySetting{
		CIM_ElementSetting: tmp,
	}
	return
}

// SetElement sets the value of Element for the instance
func (instance *Win32_TSSessionDirectorySetting) SetPropertyElement(value cimv2.Win32_TerminalService) (err error) {
	return instance.SetProperty("Element", (value))
}

// GetElement gets the value of Element for the instance
func (instance *Win32_TSSessionDirectorySetting) GetPropertyElement() (value cimv2.Win32_TerminalService, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(cimv2.Win32_TerminalService)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " cimv2.Win32_TerminalService is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = cimv2.Win32_TerminalService(valuetmp)

	return
}

// SetSetting sets the value of Setting for the instance
func (instance *Win32_TSSessionDirectorySetting) SetPropertySetting(value Win32_TSSessionDirectory) (err error) {
	return instance.SetProperty("Setting", (value))
}

// GetSetting gets the value of Setting for the instance
func (instance *Win32_TSSessionDirectorySetting) GetPropertySetting() (value Win32_TSSessionDirectory, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Win32_TSSessionDirectory)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Win32_TSSessionDirectory is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Win32_TSSessionDirectory(valuetmp)

	return
}
