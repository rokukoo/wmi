// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_RecordLog struct
type CIM_RecordLog struct {
	*CIM_Log
}

func NewCIM_RecordLogEx1(instance *cim.WmiInstance) (newInstance *CIM_RecordLog, err error) {
	tmp, err := NewCIM_LogEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_RecordLog{
		CIM_Log: tmp,
	}
	return
}

func NewCIM_RecordLogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_RecordLog, err error) {
	tmp, err := NewCIM_LogEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_RecordLog{
		CIM_Log: tmp,
	}
	return
}
