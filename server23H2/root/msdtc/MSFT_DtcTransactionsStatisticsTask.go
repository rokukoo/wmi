// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.msdtc
//
// ////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DtcTransactionsStatisticsTask struct
type MSFT_DtcTransactionsStatisticsTask struct {
	*cim.WmiInstance
}

func NewMSFT_DtcTransactionsStatisticsTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_DtcTransactionsStatisticsTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcTransactionsStatisticsTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DtcTransactionsStatisticsTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DtcTransactionsStatisticsTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcTransactionsStatisticsTask{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="DtcName" type="string "></param>

// <param name="cmdletOutput" type="DtcTransactionsStatistics "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionsStatisticsTask) Get( /* IN */ DtcName string,
	/* OUT */ cmdletOutput DtcTransactionsStatistics) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", DtcName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
