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

// Win32_PerfFormattedData struct
type Win32_PerfFormattedData struct {
	*Win32_Perf
}

func NewWin32_PerfFormattedDataEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData, err error) {
	tmp, err := NewWin32_PerfEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData{
		Win32_Perf: tmp,
	}
	return
}

func NewWin32_PerfFormattedDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData, err error) {
	tmp, err := NewWin32_PerfEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData{
		Win32_Perf: tmp,
	}
	return
}