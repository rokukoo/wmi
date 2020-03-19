// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_IPsecSAEndpoint struct
type CIM_IPsecSAEndpoint struct {
	*CIM_SecurityAssociationEndpoint

	//
	DFHandling uint16

	//
	EncapsulationMode uint16

	//
	InboundDirection bool

	//
	PFSInUse bool

	//
	SPI uint32
}

func NewCIM_IPsecSAEndpointEx1(instance *cim.WmiInstance) (newInstance *CIM_IPsecSAEndpoint, err error) {
	tmp, err := NewCIM_SecurityAssociationEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_IPsecSAEndpoint{
		CIM_SecurityAssociationEndpoint: tmp,
	}
	return
}

func NewCIM_IPsecSAEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_IPsecSAEndpoint, err error) {
	tmp, err := NewCIM_SecurityAssociationEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_IPsecSAEndpoint{
		CIM_SecurityAssociationEndpoint: tmp,
	}
	return
}

// SetDFHandling sets the value of DFHandling for the instance
func (instance *CIM_IPsecSAEndpoint) SetPropertyDFHandling(value uint16) (err error) {
	return instance.SetProperty("DFHandling", value)
}

// GetDFHandling gets the value of DFHandling for the instance
func (instance *CIM_IPsecSAEndpoint) GetPropertyDFHandling() (value uint16, err error) {
	retValue, err := instance.GetProperty("DFHandling")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncapsulationMode sets the value of EncapsulationMode for the instance
func (instance *CIM_IPsecSAEndpoint) SetPropertyEncapsulationMode(value uint16) (err error) {
	return instance.SetProperty("EncapsulationMode", value)
}

// GetEncapsulationMode gets the value of EncapsulationMode for the instance
func (instance *CIM_IPsecSAEndpoint) GetPropertyEncapsulationMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("EncapsulationMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInboundDirection sets the value of InboundDirection for the instance
func (instance *CIM_IPsecSAEndpoint) SetPropertyInboundDirection(value bool) (err error) {
	return instance.SetProperty("InboundDirection", value)
}

// GetInboundDirection gets the value of InboundDirection for the instance
func (instance *CIM_IPsecSAEndpoint) GetPropertyInboundDirection() (value bool, err error) {
	retValue, err := instance.GetProperty("InboundDirection")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPFSInUse sets the value of PFSInUse for the instance
func (instance *CIM_IPsecSAEndpoint) SetPropertyPFSInUse(value bool) (err error) {
	return instance.SetProperty("PFSInUse", value)
}

// GetPFSInUse gets the value of PFSInUse for the instance
func (instance *CIM_IPsecSAEndpoint) GetPropertyPFSInUse() (value bool, err error) {
	retValue, err := instance.GetProperty("PFSInUse")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSPI sets the value of SPI for the instance
func (instance *CIM_IPsecSAEndpoint) SetPropertySPI(value uint32) (err error) {
	return instance.SetProperty("SPI", value)
}

// GetSPI gets the value of SPI for the instance
func (instance *CIM_IPsecSAEndpoint) GetPropertySPI() (value uint32, err error) {
	retValue, err := instance.GetProperty("SPI")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
