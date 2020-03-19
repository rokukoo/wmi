// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msft_WmiProvider_PutClassAsyncEvent_Post struct
type Msft_WmiProvider_PutClassAsyncEvent_Post struct {
	*Msft_WmiProvider_OperationEvent_Post

	//
	ClassObject interface{}

	//
	Flags uint32

	//
	ObjectParameter interface{}

	//
	ResultCode uint32

	//
	StringParameter string
}

func NewMsft_WmiProvider_PutClassAsyncEvent_PostEx1(instance *cim.WmiInstance) (newInstance *Msft_WmiProvider_PutClassAsyncEvent_Post, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PostEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_PutClassAsyncEvent_Post{
		Msft_WmiProvider_OperationEvent_Post: tmp,
	}
	return
}

func NewMsft_WmiProvider_PutClassAsyncEvent_PostEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_WmiProvider_PutClassAsyncEvent_Post, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PostEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_PutClassAsyncEvent_Post{
		Msft_WmiProvider_OperationEvent_Post: tmp,
	}
	return
}

// SetClassObject sets the value of ClassObject for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Post) SetPropertyClassObject(value interface{}) (err error) {
	return instance.SetProperty("ClassObject", value)
}

// GetClassObject gets the value of ClassObject for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Post) GetPropertyClassObject() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ClassObject")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Post) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Post) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetObjectParameter sets the value of ObjectParameter for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Post) SetPropertyObjectParameter(value interface{}) (err error) {
	return instance.SetProperty("ObjectParameter", value)
}

// GetObjectParameter gets the value of ObjectParameter for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Post) GetPropertyObjectParameter() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ObjectParameter")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResultCode sets the value of ResultCode for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Post) SetPropertyResultCode(value uint32) (err error) {
	return instance.SetProperty("ResultCode", value)
}

// GetResultCode gets the value of ResultCode for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Post) GetPropertyResultCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResultCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStringParameter sets the value of StringParameter for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Post) SetPropertyStringParameter(value string) (err error) {
	return instance.SetProperty("StringParameter", value)
}

// GetStringParameter gets the value of StringParameter for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Post) GetPropertyStringParameter() (value string, err error) {
	retValue, err := instance.GetProperty("StringParameter")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
