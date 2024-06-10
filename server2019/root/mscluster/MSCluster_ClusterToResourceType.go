// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_ClusterToResourceType struct
type MSCluster_ClusterToResourceType struct {
	*CIM_Component
}

func NewMSCluster_ClusterToResourceTypeEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ClusterToResourceType, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterToResourceType{
		CIM_Component: tmp,
	}
	return
}

func NewMSCluster_ClusterToResourceTypeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ClusterToResourceType, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterToResourceType{
		CIM_Component: tmp,
	}
	return
}
