// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_WmiFilterActivated struct
type MSFT_WmiFilterActivated struct {
	*MSFT_WmiFilterEvent
}

func NewMSFT_WmiFilterActivatedEx1(instance *cim.WmiInstance) (newInstance *MSFT_WmiFilterActivated, err error) {
	tmp, err := NewMSFT_WmiFilterEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_WmiFilterActivated{
		MSFT_WmiFilterEvent: tmp,
	}
	return
}

func NewMSFT_WmiFilterActivatedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_WmiFilterActivated, err error) {
	tmp, err := NewMSFT_WmiFilterEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_WmiFilterActivated{
		MSFT_WmiFilterEvent: tmp,
	}
	return
}
