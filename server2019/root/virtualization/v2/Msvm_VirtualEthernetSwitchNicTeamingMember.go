// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualEthernetSwitchNicTeamingMember struct
type Msvm_VirtualEthernetSwitchNicTeamingMember struct {
	*CIM_Dependency
}

func NewMsvm_VirtualEthernetSwitchNicTeamingMemberEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualEthernetSwitchNicTeamingMember, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchNicTeamingMember{
		CIM_Dependency: tmp,
	}
	return
}

func NewMsvm_VirtualEthernetSwitchNicTeamingMemberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualEthernetSwitchNicTeamingMember, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchNicTeamingMember{
		CIM_Dependency: tmp,
	}
	return
}
