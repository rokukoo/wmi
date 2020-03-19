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

// MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03 struct
type MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03 struct {
	*cim.WmiInstance

	//
	EnforcementMode string

	//
	InstanceID string

	//
	NonInteractiveProcessEnforcement string

	//
	ParentID string

	//
	Policy string
}

func NewMDM_AppLocker_ApplicationLaunchRestrictions01_EXE03Ex1(instance *cim.WmiInstance) (newInstance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_AppLocker_ApplicationLaunchRestrictions01_EXE03Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03{
		WmiInstance: tmp,
	}
	return
}

// SetEnforcementMode sets the value of EnforcementMode for the instance
func (instance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03) SetPropertyEnforcementMode(value string) (err error) {
	return instance.SetProperty("EnforcementMode", value)
}

// GetEnforcementMode gets the value of EnforcementMode for the instance
func (instance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03) GetPropertyEnforcementMode() (value string, err error) {
	retValue, err := instance.GetProperty("EnforcementMode")
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
func (instance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03) GetPropertyInstanceID() (value string, err error) {
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

// SetNonInteractiveProcessEnforcement sets the value of NonInteractiveProcessEnforcement for the instance
func (instance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03) SetPropertyNonInteractiveProcessEnforcement(value string) (err error) {
	return instance.SetProperty("NonInteractiveProcessEnforcement", value)
}

// GetNonInteractiveProcessEnforcement gets the value of NonInteractiveProcessEnforcement for the instance
func (instance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03) GetPropertyNonInteractiveProcessEnforcement() (value string, err error) {
	retValue, err := instance.GetProperty("NonInteractiveProcessEnforcement")
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
func (instance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03) GetPropertyParentID() (value string, err error) {
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

// SetPolicy sets the value of Policy for the instance
func (instance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03) SetPropertyPolicy(value string) (err error) {
	return instance.SetProperty("Policy", value)
}

// GetPolicy gets the value of Policy for the instance
func (instance *MDM_AppLocker_ApplicationLaunchRestrictions01_EXE03) GetPropertyPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("Policy")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}