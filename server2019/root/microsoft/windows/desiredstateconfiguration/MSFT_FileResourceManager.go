// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_FileResourceManager struct
type MSFT_FileResourceManager struct {
	*OMI_ResourceModuleManager

	//
	Credential MSFT_Credential

	//
	SourcePath string
}

func NewMSFT_FileResourceManagerEx1(instance *cim.WmiInstance) (newInstance *MSFT_FileResourceManager, err error) {
	tmp, err := NewOMI_ResourceModuleManagerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_FileResourceManager{
		OMI_ResourceModuleManager: tmp,
	}
	return
}

func NewMSFT_FileResourceManagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_FileResourceManager, err error) {
	tmp, err := NewOMI_ResourceModuleManagerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_FileResourceManager{
		OMI_ResourceModuleManager: tmp,
	}
	return
}

// SetCredential sets the value of Credential for the instance
func (instance *MSFT_FileResourceManager) SetPropertyCredential(value MSFT_Credential) (err error) {
	return instance.SetProperty("Credential", value)
}

// GetCredential gets the value of Credential for the instance
func (instance *MSFT_FileResourceManager) GetPropertyCredential() (value MSFT_Credential, err error) {
	retValue, err := instance.GetProperty("Credential")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Credential)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourcePath sets the value of SourcePath for the instance
func (instance *MSFT_FileResourceManager) SetPropertySourcePath(value string) (err error) {
	return instance.SetProperty("SourcePath", value)
}

// GetSourcePath gets the value of SourcePath for the instance
func (instance *MSFT_FileResourceManager) GetPropertySourcePath() (value string, err error) {
	retValue, err := instance.GetProperty("SourcePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
