// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetTeredoState struct
type MSFT_NetTeredoState struct {
	*CIM_ElementSettingData

	//
	Error string

	//
	State string
}

func NewMSFT_NetTeredoStateEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetTeredoState, err error) {
	tmp, err := NewCIM_ElementSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetTeredoState{
		CIM_ElementSettingData: tmp,
	}
	return
}

func NewMSFT_NetTeredoStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetTeredoState, err error) {
	tmp, err := NewCIM_ElementSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetTeredoState{
		CIM_ElementSettingData: tmp,
	}
	return
}

// SetError sets the value of Error for the instance
func (instance *MSFT_NetTeredoState) SetPropertyError(value string) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *MSFT_NetTeredoState) GetPropertyError() (value string, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_NetTeredoState) SetPropertyState(value string) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_NetTeredoState) GetPropertyState() (value string, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
