// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerPerformanceCounterSamples struct
type MSFT_ServerPerformanceCounterSamples struct {
	*cim.WmiInstance

	//
	CounterPaths []string

	//
	Timestamps []string

	//
	Values []string
}

func NewMSFT_ServerPerformanceCounterSamplesEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerPerformanceCounterSamples, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerPerformanceCounterSamples{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerPerformanceCounterSamplesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerPerformanceCounterSamples, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerPerformanceCounterSamples{
		WmiInstance: tmp,
	}
	return
}

// SetCounterPaths sets the value of CounterPaths for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) SetPropertyCounterPaths(value []string) (err error) {
	return instance.SetProperty("CounterPaths", value)
}

// GetCounterPaths gets the value of CounterPaths for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) GetPropertyCounterPaths() (value []string, err error) {
	retValue, err := instance.GetProperty("CounterPaths")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimestamps sets the value of Timestamps for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) SetPropertyTimestamps(value []string) (err error) {
	return instance.SetProperty("Timestamps", value)
}

// GetTimestamps gets the value of Timestamps for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) GetPropertyTimestamps() (value []string, err error) {
	retValue, err := instance.GetProperty("Timestamps")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValues sets the value of Values for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) SetPropertyValues(value []string) (err error) {
	return instance.SetProperty("Values", value)
}

// GetValues gets the value of Values for the instance
func (instance *MSFT_ServerPerformanceCounterSamples) GetPropertyValues() (value []string, err error) {
	retValue, err := instance.GetProperty("Values")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
