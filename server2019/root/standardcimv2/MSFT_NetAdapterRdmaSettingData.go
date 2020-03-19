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

// MSFT_NetAdapterRdmaSettingData struct
type MSFT_NetAdapterRdmaSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	Enabled bool

	//
	ETS uint32

	//
	MaxCompletionQueueCount uint32

	//
	MaxInboundReadLimit uint32

	//
	MaxMemoryRegionCount uint32

	//
	MaxMemoryWindowCount uint32

	//
	MaxOutboundReadLimit uint32

	//
	MaxProtectionDomainCount uint32

	//
	MaxQueuePairCount uint32

	//
	MaxSharedReceiveQueueCount uint32

	//
	PFC uint32

	//
	RdmaAdapterInfo MSFT_NetAdapter_RdmaAdapterInfo

	//
	RdmaMissingCounterInfo MSFT_NetAdapter_RdmaMissingCounterInfo
}

func NewMSFT_NetAdapterRdmaSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterRdmaSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterRdmaSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterRdmaSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterRdmaSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterRdmaSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetETS sets the value of ETS for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyETS(value uint32) (err error) {
	return instance.SetProperty("ETS", value)
}

// GetETS gets the value of ETS for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyETS() (value uint32, err error) {
	retValue, err := instance.GetProperty("ETS")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxCompletionQueueCount sets the value of MaxCompletionQueueCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyMaxCompletionQueueCount(value uint32) (err error) {
	return instance.SetProperty("MaxCompletionQueueCount", value)
}

// GetMaxCompletionQueueCount gets the value of MaxCompletionQueueCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyMaxCompletionQueueCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCompletionQueueCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxInboundReadLimit sets the value of MaxInboundReadLimit for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyMaxInboundReadLimit(value uint32) (err error) {
	return instance.SetProperty("MaxInboundReadLimit", value)
}

// GetMaxInboundReadLimit gets the value of MaxInboundReadLimit for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyMaxInboundReadLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInboundReadLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxMemoryRegionCount sets the value of MaxMemoryRegionCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyMaxMemoryRegionCount(value uint32) (err error) {
	return instance.SetProperty("MaxMemoryRegionCount", value)
}

// GetMaxMemoryRegionCount gets the value of MaxMemoryRegionCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyMaxMemoryRegionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxMemoryRegionCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxMemoryWindowCount sets the value of MaxMemoryWindowCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyMaxMemoryWindowCount(value uint32) (err error) {
	return instance.SetProperty("MaxMemoryWindowCount", value)
}

// GetMaxMemoryWindowCount gets the value of MaxMemoryWindowCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyMaxMemoryWindowCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxMemoryWindowCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxOutboundReadLimit sets the value of MaxOutboundReadLimit for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyMaxOutboundReadLimit(value uint32) (err error) {
	return instance.SetProperty("MaxOutboundReadLimit", value)
}

// GetMaxOutboundReadLimit gets the value of MaxOutboundReadLimit for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyMaxOutboundReadLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxOutboundReadLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxProtectionDomainCount sets the value of MaxProtectionDomainCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyMaxProtectionDomainCount(value uint32) (err error) {
	return instance.SetProperty("MaxProtectionDomainCount", value)
}

// GetMaxProtectionDomainCount gets the value of MaxProtectionDomainCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyMaxProtectionDomainCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxProtectionDomainCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxQueuePairCount sets the value of MaxQueuePairCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyMaxQueuePairCount(value uint32) (err error) {
	return instance.SetProperty("MaxQueuePairCount", value)
}

// GetMaxQueuePairCount gets the value of MaxQueuePairCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyMaxQueuePairCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxQueuePairCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxSharedReceiveQueueCount sets the value of MaxSharedReceiveQueueCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyMaxSharedReceiveQueueCount(value uint32) (err error) {
	return instance.SetProperty("MaxSharedReceiveQueueCount", value)
}

// GetMaxSharedReceiveQueueCount gets the value of MaxSharedReceiveQueueCount for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyMaxSharedReceiveQueueCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxSharedReceiveQueueCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPFC sets the value of PFC for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyPFC(value uint32) (err error) {
	return instance.SetProperty("PFC", value)
}

// GetPFC gets the value of PFC for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyPFC() (value uint32, err error) {
	retValue, err := instance.GetProperty("PFC")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRdmaAdapterInfo sets the value of RdmaAdapterInfo for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyRdmaAdapterInfo(value MSFT_NetAdapter_RdmaAdapterInfo) (err error) {
	return instance.SetProperty("RdmaAdapterInfo", value)
}

// GetRdmaAdapterInfo gets the value of RdmaAdapterInfo for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyRdmaAdapterInfo() (value MSFT_NetAdapter_RdmaAdapterInfo, err error) {
	retValue, err := instance.GetProperty("RdmaAdapterInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapter_RdmaAdapterInfo)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRdmaMissingCounterInfo sets the value of RdmaMissingCounterInfo for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) SetPropertyRdmaMissingCounterInfo(value MSFT_NetAdapter_RdmaMissingCounterInfo) (err error) {
	return instance.SetProperty("RdmaMissingCounterInfo", value)
}

// GetRdmaMissingCounterInfo gets the value of RdmaMissingCounterInfo for the instance
func (instance *MSFT_NetAdapterRdmaSettingData) GetPropertyRdmaMissingCounterInfo() (value MSFT_NetAdapter_RdmaMissingCounterInfo, err error) {
	retValue, err := instance.GetProperty("RdmaMissingCounterInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapter_RdmaMissingCounterInfo)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="cmdletOutput" type="MSFT_NetAdapterRdmaSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterRdmaSettingData) Enable( /* OUT */ cmdletOutput MSFT_NetAdapterRdmaSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="cmdletOutput" type="MSFT_NetAdapterRdmaSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterRdmaSettingData) Disable( /* OUT */ cmdletOutput MSFT_NetAdapterRdmaSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
