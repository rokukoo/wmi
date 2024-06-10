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

// Win32_ModuleTrace struct
type Win32_ModuleTrace struct {
	*Win32_SystemTrace
}

func NewWin32_ModuleTraceEx1(instance *cim.WmiInstance) (newInstance *Win32_ModuleTrace, err error) {
	tmp, err := NewWin32_SystemTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_ModuleTrace{
		Win32_SystemTrace: tmp,
	}
	return
}

func NewWin32_ModuleTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ModuleTrace, err error) {
	tmp, err := NewWin32_SystemTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ModuleTrace{
		Win32_SystemTrace: tmp,
	}
	return
}
