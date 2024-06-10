// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_Error struct
type Msvm_Error struct {
	*CIM_Error
}

func NewMsvm_ErrorEx1(instance *cim.WmiInstance) (newInstance *Msvm_Error, err error) {
	tmp, err := NewCIM_ErrorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_Error{
		CIM_Error: tmp,
	}
	return
}

func NewMsvm_ErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_Error, err error) {
	tmp, err := NewCIM_ErrorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_Error{
		CIM_Error: tmp,
	}
	return
}
