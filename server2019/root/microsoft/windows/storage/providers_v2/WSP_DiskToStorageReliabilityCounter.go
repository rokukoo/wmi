// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_DiskToStorageReliabilityCounter struct
type WSP_DiskToStorageReliabilityCounter struct {
	*MSFT_DiskToStorageReliabilityCounter
}

func NewWSP_DiskToStorageReliabilityCounterEx1(instance *cim.WmiInstance) (newInstance *WSP_DiskToStorageReliabilityCounter, err error) {
	tmp, err := NewMSFT_DiskToStorageReliabilityCounterEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_DiskToStorageReliabilityCounter{
		MSFT_DiskToStorageReliabilityCounter: tmp,
	}
	return
}

func NewWSP_DiskToStorageReliabilityCounterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_DiskToStorageReliabilityCounter, err error) {
	tmp, err := NewMSFT_DiskToStorageReliabilityCounterEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_DiskToStorageReliabilityCounter{
		MSFT_DiskToStorageReliabilityCounter: tmp,
	}
	return
}
