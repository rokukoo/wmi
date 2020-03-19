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

// MSFT_PrinterPort struct
type MSFT_PrinterPort struct {
	*CIM_ManagedSystemElement

	//
	ComputerName string

	//
	PortMonitor string
}

func NewMSFT_PrinterPortEx1(instance *cim.WmiInstance) (newInstance *MSFT_PrinterPort, err error) {
	tmp, err := NewCIM_ManagedSystemElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterPort{
		CIM_ManagedSystemElement: tmp,
	}
	return
}

func NewMSFT_PrinterPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_PrinterPort, err error) {
	tmp, err := NewCIM_ManagedSystemElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterPort{
		CIM_ManagedSystemElement: tmp,
	}
	return
}

// SetComputerName sets the value of ComputerName for the instance
func (instance *MSFT_PrinterPort) SetPropertyComputerName(value string) (err error) {
	return instance.SetProperty("ComputerName", value)
}

// GetComputerName gets the value of ComputerName for the instance
func (instance *MSFT_PrinterPort) GetPropertyComputerName() (value string, err error) {
	retValue, err := instance.GetProperty("ComputerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortMonitor sets the value of PortMonitor for the instance
func (instance *MSFT_PrinterPort) SetPropertyPortMonitor(value string) (err error) {
	return instance.SetProperty("PortMonitor", value)
}

// GetPortMonitor gets the value of PortMonitor for the instance
func (instance *MSFT_PrinterPort) GetPropertyPortMonitor() (value string, err error) {
	retValue, err := instance.GetProperty("PortMonitor")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
