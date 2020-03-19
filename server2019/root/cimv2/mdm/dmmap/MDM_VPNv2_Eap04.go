// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_VPNv2_Eap04 struct
type MDM_VPNv2_Eap04 struct {
	*cim.WmiInstance

	//
	Configuration string

	//
	InstanceID string

	//
	ParentID string

	//
	Type int32
}

func NewMDM_VPNv2_Eap04Ex1(instance *cim.WmiInstance) (newInstance *MDM_VPNv2_Eap04, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_Eap04{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_VPNv2_Eap04Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_VPNv2_Eap04, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_Eap04{
		WmiInstance: tmp,
	}
	return
}

// SetConfiguration sets the value of Configuration for the instance
func (instance *MDM_VPNv2_Eap04) SetPropertyConfiguration(value string) (err error) {
	return instance.SetProperty("Configuration", value)
}

// GetConfiguration gets the value of Configuration for the instance
func (instance *MDM_VPNv2_Eap04) GetPropertyConfiguration() (value string, err error) {
	retValue, err := instance.GetProperty("Configuration")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_VPNv2_Eap04) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_Eap04) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_VPNv2_Eap04) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_Eap04) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MDM_VPNv2_Eap04) SetPropertyType(value int32) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MDM_VPNv2_Eap04) GetPropertyType() (value int32, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
