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

// MSFT_NetDependOnLaterGroup struct
type MSFT_NetDependOnLaterGroup struct {
	*MSFT_SCMEventLogEvent

	//
	Service string
}

func NewMSFT_NetDependOnLaterGroupEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetDependOnLaterGroup, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetDependOnLaterGroup{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

func NewMSFT_NetDependOnLaterGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetDependOnLaterGroup, err error) {
	tmp, err := NewMSFT_SCMEventLogEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetDependOnLaterGroup{
		MSFT_SCMEventLogEvent: tmp,
	}
	return
}

// SetService sets the value of Service for the instance
func (instance *MSFT_NetDependOnLaterGroup) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetDependOnLaterGroup) GetPropertyService() (value string, err error) {
	retValue, err := instance.GetProperty("Service")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}