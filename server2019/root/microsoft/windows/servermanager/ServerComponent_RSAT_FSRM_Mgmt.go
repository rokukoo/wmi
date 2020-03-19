// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ServerComponent_RSAT_FSRM_Mgmt struct
type ServerComponent_RSAT_FSRM_Mgmt struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_RSAT_FSRM_MgmtEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_RSAT_FSRM_Mgmt, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_RSAT_FSRM_Mgmt{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_RSAT_FSRM_MgmtEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_RSAT_FSRM_Mgmt, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_RSAT_FSRM_Mgmt{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
