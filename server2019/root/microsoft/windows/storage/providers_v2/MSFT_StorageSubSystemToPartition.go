// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageSubSystemToPartition struct
type MSFT_StorageSubSystemToPartition struct {
	*cim.WmiInstance

	//
	Partition MSFT_Partition

	//
	StorageSubSystem MSFT_StorageSubSystem
}

func NewMSFT_StorageSubSystemToPartitionEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageSubSystemToPartition, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageSubSystemToPartition{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_StorageSubSystemToPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageSubSystemToPartition, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageSubSystemToPartition{
		WmiInstance: tmp,
	}
	return
}

// SetPartition sets the value of Partition for the instance
func (instance *MSFT_StorageSubSystemToPartition) SetPropertyPartition(value MSFT_Partition) (err error) {
	return instance.SetProperty("Partition", value)
}

// GetPartition gets the value of Partition for the instance
func (instance *MSFT_StorageSubSystemToPartition) GetPropertyPartition() (value MSFT_Partition, err error) {
	retValue, err := instance.GetProperty("Partition")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Partition)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubSystem sets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToPartition) SetPropertyStorageSubSystem(value MSFT_StorageSubSystem) (err error) {
	return instance.SetProperty("StorageSubSystem", value)
}

// GetStorageSubSystem gets the value of StorageSubSystem for the instance
func (instance *MSFT_StorageSubSystemToPartition) GetPropertyStorageSubSystem() (value MSFT_StorageSubSystem, err error) {
	retValue, err := instance.GetProperty("StorageSubSystem")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageSubSystem)
	if !ok {
		// TODO: Set an error
	}
	return
}
