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

// MSFT_NetTransportConnection struct
type MSFT_NetTransportConnection struct {
	*CIM_NetworkPipe

	//
	CreationTime string

	//
	LocalAddress string

	//
	LocalPort uint16

	//
	OwningProcess uint32
}

func NewMSFT_NetTransportConnectionEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetTransportConnection, err error) {
	tmp, err := NewCIM_NetworkPipeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetTransportConnection{
		CIM_NetworkPipe: tmp,
	}
	return
}

func NewMSFT_NetTransportConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetTransportConnection, err error) {
	tmp, err := NewCIM_NetworkPipeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetTransportConnection{
		CIM_NetworkPipe: tmp,
	}
	return
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *MSFT_NetTransportConnection) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", value)
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *MSFT_NetTransportConnection) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalAddress sets the value of LocalAddress for the instance
func (instance *MSFT_NetTransportConnection) SetPropertyLocalAddress(value string) (err error) {
	return instance.SetProperty("LocalAddress", value)
}

// GetLocalAddress gets the value of LocalAddress for the instance
func (instance *MSFT_NetTransportConnection) GetPropertyLocalAddress() (value string, err error) {
	retValue, err := instance.GetProperty("LocalAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalPort sets the value of LocalPort for the instance
func (instance *MSFT_NetTransportConnection) SetPropertyLocalPort(value uint16) (err error) {
	return instance.SetProperty("LocalPort", value)
}

// GetLocalPort gets the value of LocalPort for the instance
func (instance *MSFT_NetTransportConnection) GetPropertyLocalPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("LocalPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOwningProcess sets the value of OwningProcess for the instance
func (instance *MSFT_NetTransportConnection) SetPropertyOwningProcess(value uint32) (err error) {
	return instance.SetProperty("OwningProcess", value)
}

// GetOwningProcess gets the value of OwningProcess for the instance
func (instance *MSFT_NetTransportConnection) GetPropertyOwningProcess() (value uint32, err error) {
	retValue, err := instance.GetProperty("OwningProcess")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
