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

// MDM_eUICCs_01 struct
type MDM_eUICCs_01 struct {
	*cim.WmiInstance

	//
	Identifier string

	//
	InstanceID string

	//
	IsActive bool

	//
	ParentID string

	//
	PPR1Allowed bool

	//
	PPR1AlreadySet bool
}

func NewMDM_eUICCs_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_eUICCs_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_eUICCs_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_eUICCs_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_eUICCs_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_eUICCs_01{
		WmiInstance: tmp,
	}
	return
}

// SetIdentifier sets the value of Identifier for the instance
func (instance *MDM_eUICCs_01) SetPropertyIdentifier(value string) (err error) {
	return instance.SetProperty("Identifier", value)
}

// GetIdentifier gets the value of Identifier for the instance
func (instance *MDM_eUICCs_01) GetPropertyIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("Identifier")
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
func (instance *MDM_eUICCs_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_eUICCs_01) GetPropertyInstanceID() (value string, err error) {
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

// SetIsActive sets the value of IsActive for the instance
func (instance *MDM_eUICCs_01) SetPropertyIsActive(value bool) (err error) {
	return instance.SetProperty("IsActive", value)
}

// GetIsActive gets the value of IsActive for the instance
func (instance *MDM_eUICCs_01) GetPropertyIsActive() (value bool, err error) {
	retValue, err := instance.GetProperty("IsActive")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_eUICCs_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_eUICCs_01) GetPropertyParentID() (value string, err error) {
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

// SetPPR1Allowed sets the value of PPR1Allowed for the instance
func (instance *MDM_eUICCs_01) SetPropertyPPR1Allowed(value bool) (err error) {
	return instance.SetProperty("PPR1Allowed", value)
}

// GetPPR1Allowed gets the value of PPR1Allowed for the instance
func (instance *MDM_eUICCs_01) GetPropertyPPR1Allowed() (value bool, err error) {
	retValue, err := instance.GetProperty("PPR1Allowed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPPR1AlreadySet sets the value of PPR1AlreadySet for the instance
func (instance *MDM_eUICCs_01) SetPropertyPPR1AlreadySet(value bool) (err error) {
	return instance.SetProperty("PPR1AlreadySet", value)
}

// GetPPR1AlreadySet gets the value of PPR1AlreadySet for the instance
func (instance *MDM_eUICCs_01) GetPropertyPPR1AlreadySet() (value bool, err error) {
	retValue, err := instance.GetProperty("PPR1AlreadySet")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
