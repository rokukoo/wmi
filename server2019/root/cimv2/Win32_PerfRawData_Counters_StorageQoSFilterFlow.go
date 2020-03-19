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

// Win32_PerfRawData_Counters_StorageQoSFilterFlow struct
type Win32_PerfRawData_Counters_StorageQoSFilterFlow struct {
	*Win32_PerfRawData

	//
	AvgBandwidth uint64

	//
	AvgDeviceQueueLength uint64

	//
	AvgIOQuotaReplenishmentOperationsPersec uint64

	//
	AvgNormalizedIOPS uint64

	//
	AvgSchedulerQueueLength uint64

	//
	MaximumBandwidth uint64

	//
	NormalizedMaximumIORate uint64

	//
	NormalizedMinimumIORate uint64

	//
	TotalBandwidthquotaIncrementPersec uint64

	//
	TotalNormalizedIOQuotaIncrement uint64
}

func NewWin32_PerfRawData_Counters_StorageQoSFilterFlowEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_StorageQoSFilterFlow, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_StorageQoSFilterFlow{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_StorageQoSFilterFlowEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_StorageQoSFilterFlow, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_StorageQoSFilterFlow{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAvgBandwidth sets the value of AvgBandwidth for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) SetPropertyAvgBandwidth(value uint64) (err error) {
	return instance.SetProperty("AvgBandwidth", value)
}

// GetAvgBandwidth gets the value of AvgBandwidth for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) GetPropertyAvgBandwidth() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgBandwidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgDeviceQueueLength sets the value of AvgDeviceQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) SetPropertyAvgDeviceQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgDeviceQueueLength", value)
}

// GetAvgDeviceQueueLength gets the value of AvgDeviceQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) GetPropertyAvgDeviceQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgDeviceQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgIOQuotaReplenishmentOperationsPersec sets the value of AvgIOQuotaReplenishmentOperationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) SetPropertyAvgIOQuotaReplenishmentOperationsPersec(value uint64) (err error) {
	return instance.SetProperty("AvgIOQuotaReplenishmentOperationsPersec", value)
}

// GetAvgIOQuotaReplenishmentOperationsPersec gets the value of AvgIOQuotaReplenishmentOperationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) GetPropertyAvgIOQuotaReplenishmentOperationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgIOQuotaReplenishmentOperationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgNormalizedIOPS sets the value of AvgNormalizedIOPS for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) SetPropertyAvgNormalizedIOPS(value uint64) (err error) {
	return instance.SetProperty("AvgNormalizedIOPS", value)
}

// GetAvgNormalizedIOPS gets the value of AvgNormalizedIOPS for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) GetPropertyAvgNormalizedIOPS() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgNormalizedIOPS")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgSchedulerQueueLength sets the value of AvgSchedulerQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) SetPropertyAvgSchedulerQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgSchedulerQueueLength", value)
}

// GetAvgSchedulerQueueLength gets the value of AvgSchedulerQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) GetPropertyAvgSchedulerQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgSchedulerQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumBandwidth sets the value of MaximumBandwidth for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) SetPropertyMaximumBandwidth(value uint64) (err error) {
	return instance.SetProperty("MaximumBandwidth", value)
}

// GetMaximumBandwidth gets the value of MaximumBandwidth for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) GetPropertyMaximumBandwidth() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaximumBandwidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNormalizedMaximumIORate sets the value of NormalizedMaximumIORate for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) SetPropertyNormalizedMaximumIORate(value uint64) (err error) {
	return instance.SetProperty("NormalizedMaximumIORate", value)
}

// GetNormalizedMaximumIORate gets the value of NormalizedMaximumIORate for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) GetPropertyNormalizedMaximumIORate() (value uint64, err error) {
	retValue, err := instance.GetProperty("NormalizedMaximumIORate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNormalizedMinimumIORate sets the value of NormalizedMinimumIORate for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) SetPropertyNormalizedMinimumIORate(value uint64) (err error) {
	return instance.SetProperty("NormalizedMinimumIORate", value)
}

// GetNormalizedMinimumIORate gets the value of NormalizedMinimumIORate for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) GetPropertyNormalizedMinimumIORate() (value uint64, err error) {
	retValue, err := instance.GetProperty("NormalizedMinimumIORate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalBandwidthquotaIncrementPersec sets the value of TotalBandwidthquotaIncrementPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) SetPropertyTotalBandwidthquotaIncrementPersec(value uint64) (err error) {
	return instance.SetProperty("TotalBandwidthquotaIncrementPersec", value)
}

// GetTotalBandwidthquotaIncrementPersec gets the value of TotalBandwidthquotaIncrementPersec for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) GetPropertyTotalBandwidthquotaIncrementPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalBandwidthquotaIncrementPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalNormalizedIOQuotaIncrement sets the value of TotalNormalizedIOQuotaIncrement for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) SetPropertyTotalNormalizedIOQuotaIncrement(value uint64) (err error) {
	return instance.SetProperty("TotalNormalizedIOQuotaIncrement", value)
}

// GetTotalNormalizedIOQuotaIncrement gets the value of TotalNormalizedIOQuotaIncrement for the instance
func (instance *Win32_PerfRawData_Counters_StorageQoSFilterFlow) GetPropertyTotalNormalizedIOQuotaIncrement() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalNormalizedIOQuotaIncrement")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
