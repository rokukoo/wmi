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

// Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement struct
type Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement struct {
	*Win32_PerfRawData

	//
	ResourcesVMsrunningRemoteFX uint64

	//
	VRAMAvailableMBperGPU uint64

	//
	VRAMReservedPercentperGPU uint64

	//
	VRAMReservedPercentperGPU_Base uint64
}

func NewWin32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagementEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetResourcesVMsrunningRemoteFX sets the value of ResourcesVMsrunningRemoteFX for the instance
func (instance *Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) SetPropertyResourcesVMsrunningRemoteFX(value uint64) (err error) {
	return instance.SetProperty("ResourcesVMsrunningRemoteFX", value)
}

// GetResourcesVMsrunningRemoteFX gets the value of ResourcesVMsrunningRemoteFX for the instance
func (instance *Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) GetPropertyResourcesVMsrunningRemoteFX() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourcesVMsrunningRemoteFX")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVRAMAvailableMBperGPU sets the value of VRAMAvailableMBperGPU for the instance
func (instance *Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) SetPropertyVRAMAvailableMBperGPU(value uint64) (err error) {
	return instance.SetProperty("VRAMAvailableMBperGPU", value)
}

// GetVRAMAvailableMBperGPU gets the value of VRAMAvailableMBperGPU for the instance
func (instance *Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) GetPropertyVRAMAvailableMBperGPU() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRAMAvailableMBperGPU")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVRAMReservedPercentperGPU sets the value of VRAMReservedPercentperGPU for the instance
func (instance *Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) SetPropertyVRAMReservedPercentperGPU(value uint64) (err error) {
	return instance.SetProperty("VRAMReservedPercentperGPU", value)
}

// GetVRAMReservedPercentperGPU gets the value of VRAMReservedPercentperGPU for the instance
func (instance *Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) GetPropertyVRAMReservedPercentperGPU() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRAMReservedPercentperGPU")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVRAMReservedPercentperGPU_Base sets the value of VRAMReservedPercentperGPU_Base for the instance
func (instance *Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) SetPropertyVRAMReservedPercentperGPU_Base(value uint64) (err error) {
	return instance.SetProperty("VRAMReservedPercentperGPU_Base", value)
}

// GetVRAMReservedPercentperGPU_Base gets the value of VRAMReservedPercentperGPU_Base for the instance
func (instance *Win32_PerfRawData_Synth3dVideoPerfProvider_RemoteFXRootGPUManagement) GetPropertyVRAMReservedPercentperGPU_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("VRAMReservedPercentperGPU_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
