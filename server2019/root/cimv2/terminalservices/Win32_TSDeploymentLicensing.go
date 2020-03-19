// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_TSDeploymentLicensing struct
type Win32_TSDeploymentLicensing struct {
	*CIM_LogicalElement

	// License Servers to use
	LicenseServers []string

	//  Licensing Mode
	LicensingType TSDeploymentLicensing_LicensingType

	// Use deployment-wide licensing settings, as opposed to setting them per-server.  If this is set to false, all other licensing settings are ignored.
	UseCentralLicensingSettings bool
}

func NewWin32_TSDeploymentLicensingEx1(instance *cim.WmiInstance) (newInstance *Win32_TSDeploymentLicensing, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSDeploymentLicensing{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewWin32_TSDeploymentLicensingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSDeploymentLicensing, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSDeploymentLicensing{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetLicenseServers sets the value of LicenseServers for the instance
func (instance *Win32_TSDeploymentLicensing) SetPropertyLicenseServers(value []string) (err error) {
	return instance.SetProperty("LicenseServers", value)
}

// GetLicenseServers gets the value of LicenseServers for the instance
func (instance *Win32_TSDeploymentLicensing) GetPropertyLicenseServers() (value []string, err error) {
	retValue, err := instance.GetProperty("LicenseServers")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLicensingType sets the value of LicensingType for the instance
func (instance *Win32_TSDeploymentLicensing) SetPropertyLicensingType(value TSDeploymentLicensing_LicensingType) (err error) {
	return instance.SetProperty("LicensingType", value)
}

// GetLicensingType gets the value of LicensingType for the instance
func (instance *Win32_TSDeploymentLicensing) GetPropertyLicensingType() (value TSDeploymentLicensing_LicensingType, err error) {
	retValue, err := instance.GetProperty("LicensingType")
	if err != nil {
		return
	}
	value, ok := retValue.(TSDeploymentLicensing_LicensingType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseCentralLicensingSettings sets the value of UseCentralLicensingSettings for the instance
func (instance *Win32_TSDeploymentLicensing) SetPropertyUseCentralLicensingSettings(value bool) (err error) {
	return instance.SetProperty("UseCentralLicensingSettings", value)
}

// GetUseCentralLicensingSettings gets the value of UseCentralLicensingSettings for the instance
func (instance *Win32_TSDeploymentLicensing) GetPropertyUseCentralLicensingSettings() (value bool, err error) {
	retValue, err := instance.GetProperty("UseCentralLicensingSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
