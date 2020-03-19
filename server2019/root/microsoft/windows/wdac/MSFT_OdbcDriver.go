// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Wdac
//////////////////////////////////////////////
package wdac

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_OdbcDriver struct
type MSFT_OdbcDriver struct {
	*cim.WmiInstance

	//
	KeyValuePair []MSFT_OdbcKeyValuePair

	//
	Name string

	//
	Platform string
}

func NewMSFT_OdbcDriverEx1(instance *cim.WmiInstance) (newInstance *MSFT_OdbcDriver, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_OdbcDriver{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_OdbcDriverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_OdbcDriver, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_OdbcDriver{
		WmiInstance: tmp,
	}
	return
}

// SetKeyValuePair sets the value of KeyValuePair for the instance
func (instance *MSFT_OdbcDriver) SetPropertyKeyValuePair(value []MSFT_OdbcKeyValuePair) (err error) {
	return instance.SetProperty("KeyValuePair", value)
}

// GetKeyValuePair gets the value of KeyValuePair for the instance
func (instance *MSFT_OdbcDriver) GetPropertyKeyValuePair() (value []MSFT_OdbcKeyValuePair, err error) {
	retValue, err := instance.GetProperty("KeyValuePair")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_OdbcKeyValuePair)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_OdbcDriver) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_OdbcDriver) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlatform sets the value of Platform for the instance
func (instance *MSFT_OdbcDriver) SetPropertyPlatform(value string) (err error) {
	return instance.SetProperty("Platform", value)
}

// GetPlatform gets the value of Platform for the instance
func (instance *MSFT_OdbcDriver) GetPropertyPlatform() (value string, err error) {
	retValue, err := instance.GetProperty("Platform")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
