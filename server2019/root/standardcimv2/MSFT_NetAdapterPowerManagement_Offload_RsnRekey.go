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

// MSFT_NetAdapterPowerManagement_Offload_RsnRekey struct
type MSFT_NetAdapterPowerManagement_Offload_RsnRekey struct {
	*MSFT_NetAdapterPowerManagement_Offload

	//
	KCK []uint8

	//
	KEK []uint8

	//
	ReplayCounter uint64
}

func NewMSFT_NetAdapterPowerManagement_Offload_RsnRekeyEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterPowerManagement_Offload_RsnRekey, err error) {
	tmp, err := NewMSFT_NetAdapterPowerManagement_OffloadEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_Offload_RsnRekey{
		MSFT_NetAdapterPowerManagement_Offload: tmp,
	}
	return
}

func NewMSFT_NetAdapterPowerManagement_Offload_RsnRekeyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterPowerManagement_Offload_RsnRekey, err error) {
	tmp, err := NewMSFT_NetAdapterPowerManagement_OffloadEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_Offload_RsnRekey{
		MSFT_NetAdapterPowerManagement_Offload: tmp,
	}
	return
}

// SetKCK sets the value of KCK for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_RsnRekey) SetPropertyKCK(value []uint8) (err error) {
	return instance.SetProperty("KCK", value)
}

// GetKCK gets the value of KCK for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_RsnRekey) GetPropertyKCK() (value []uint8, err error) {
	retValue, err := instance.GetProperty("KCK")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKEK sets the value of KEK for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_RsnRekey) SetPropertyKEK(value []uint8) (err error) {
	return instance.SetProperty("KEK", value)
}

// GetKEK gets the value of KEK for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_RsnRekey) GetPropertyKEK() (value []uint8, err error) {
	retValue, err := instance.GetProperty("KEK")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplayCounter sets the value of ReplayCounter for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_RsnRekey) SetPropertyReplayCounter(value uint64) (err error) {
	return instance.SetProperty("ReplayCounter", value)
}

// GetReplayCounter gets the value of ReplayCounter for the instance
func (instance *MSFT_NetAdapterPowerManagement_Offload_RsnRekey) GetPropertyReplayCounter() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReplayCounter")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
