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

// MDM_Policy_Result01_SystemServices02 struct
type MDM_Policy_Result01_SystemServices02 struct {
	*cim.WmiInstance

	//
	ConfigureHomeGroupListenerServiceStartupMode int32

	//
	ConfigureHomeGroupProviderServiceStartupMode int32

	//
	ConfigureXboxAccessoryManagementServiceStartupMode int32

	//
	ConfigureXboxLiveAuthManagerServiceStartupMode int32

	//
	ConfigureXboxLiveGameSaveServiceStartupMode int32

	//
	ConfigureXboxLiveNetworkingServiceStartupMode int32

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Result01_SystemServices02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_SystemServices02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_SystemServices02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_SystemServices02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_SystemServices02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_SystemServices02{
		WmiInstance: tmp,
	}
	return
}

// SetConfigureHomeGroupListenerServiceStartupMode sets the value of ConfigureHomeGroupListenerServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) SetPropertyConfigureHomeGroupListenerServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureHomeGroupListenerServiceStartupMode", value)
}

// GetConfigureHomeGroupListenerServiceStartupMode gets the value of ConfigureHomeGroupListenerServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) GetPropertyConfigureHomeGroupListenerServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureHomeGroupListenerServiceStartupMode")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigureHomeGroupProviderServiceStartupMode sets the value of ConfigureHomeGroupProviderServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) SetPropertyConfigureHomeGroupProviderServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureHomeGroupProviderServiceStartupMode", value)
}

// GetConfigureHomeGroupProviderServiceStartupMode gets the value of ConfigureHomeGroupProviderServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) GetPropertyConfigureHomeGroupProviderServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureHomeGroupProviderServiceStartupMode")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigureXboxAccessoryManagementServiceStartupMode sets the value of ConfigureXboxAccessoryManagementServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) SetPropertyConfigureXboxAccessoryManagementServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureXboxAccessoryManagementServiceStartupMode", value)
}

// GetConfigureXboxAccessoryManagementServiceStartupMode gets the value of ConfigureXboxAccessoryManagementServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) GetPropertyConfigureXboxAccessoryManagementServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureXboxAccessoryManagementServiceStartupMode")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigureXboxLiveAuthManagerServiceStartupMode sets the value of ConfigureXboxLiveAuthManagerServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) SetPropertyConfigureXboxLiveAuthManagerServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureXboxLiveAuthManagerServiceStartupMode", value)
}

// GetConfigureXboxLiveAuthManagerServiceStartupMode gets the value of ConfigureXboxLiveAuthManagerServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) GetPropertyConfigureXboxLiveAuthManagerServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureXboxLiveAuthManagerServiceStartupMode")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigureXboxLiveGameSaveServiceStartupMode sets the value of ConfigureXboxLiveGameSaveServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) SetPropertyConfigureXboxLiveGameSaveServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureXboxLiveGameSaveServiceStartupMode", value)
}

// GetConfigureXboxLiveGameSaveServiceStartupMode gets the value of ConfigureXboxLiveGameSaveServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) GetPropertyConfigureXboxLiveGameSaveServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureXboxLiveGameSaveServiceStartupMode")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigureXboxLiveNetworkingServiceStartupMode sets the value of ConfigureXboxLiveNetworkingServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) SetPropertyConfigureXboxLiveNetworkingServiceStartupMode(value int32) (err error) {
	return instance.SetProperty("ConfigureXboxLiveNetworkingServiceStartupMode", value)
}

// GetConfigureXboxLiveNetworkingServiceStartupMode gets the value of ConfigureXboxLiveNetworkingServiceStartupMode for the instance
func (instance *MDM_Policy_Result01_SystemServices02) GetPropertyConfigureXboxLiveNetworkingServiceStartupMode() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureXboxLiveNetworkingServiceStartupMode")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_SystemServices02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_SystemServices02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_SystemServices02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_SystemServices02) GetPropertyParentID() (value string, err error) {
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
