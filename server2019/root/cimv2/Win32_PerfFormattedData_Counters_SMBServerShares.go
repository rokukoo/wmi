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

// Win32_PerfFormattedData_Counters_SMBServerShares struct
type Win32_PerfFormattedData_Counters_SMBServerShares struct {
	*Win32_PerfFormattedData

	//
	AvgBytesPerRead uint64

	//
	AvgBytesPerWrite uint64

	//
	AvgDataBytesPerRequest uint64

	//
	AvgDataQueueLength uint64

	//
	AvgReadQueueLength uint64

	//
	AvgsecPerDataRequest uint32

	//
	AvgsecPerRead uint32

	//
	AvgsecPerRequest uint32

	//
	AvgsecPerWrite uint32

	//
	AvgWriteQueueLength uint64

	//
	BytesCompressedPersec uint64

	//
	CompressedRequestsPersec uint64

	//
	CompressedResponsesPersec uint64

	//
	CurrentBypassOpenFileCount uint64

	//
	CurrentDataQueueLength uint64

	//
	CurrentDurableOpenFileCount uint64

	//
	CurrentOpenFileCount uint64

	//
	CurrentPendingRequests uint64

	//
	DataBytesPersec uint64

	//
	DataRequestsPersec uint32

	//
	FilesOpenedPersec uint64

	//
	MetadataRequestsPersec uint64

	//
	PercentPersistentHandles uint64

	//
	PercentResilientHandles uint64

	//
	ReadBytesPersec uint64

	//
	ReadBytestransmittedByPassCSVPersec uint64

	//
	ReadBytestransmittedviaSMBDirectPersec uint64

	//
	ReadRequestsPersec uint32

	//
	ReadRequeststransmittedviaBypassCSVPersec uint32

	//
	ReadRequeststransmittedviaSMBDirectPersec uint32

	//
	ReceivedBytesPersec uint64

	//
	RequestsPersec uint64

	//
	SentBytesPersec uint64

	//
	TotalDurableHandleReopenCount uint64

	//
	TotalFailedDurableHandleReopenCount uint64

	//
	TotalFailedPersistentHandleReopenCount uint64

	//
	TotalFailedResilientHandleReopenCount uint64

	//
	TotalFileOpenCount uint64

	//
	TotalPersistentHandleReopenCount uint64

	//
	TotalResilientHandleReopenCount uint64

	//
	TransferredBytesPersec uint64

	//
	TreeConnectCount uint64

	//
	WriteBytesPersec uint64

	//
	WriteBytestransmittedByPassCSVPersec uint64

	//
	WriteBytestransmittedviaSMBDirectPersec uint64

	//
	WriteRequestsPersec uint32

	//
	WriteRequeststransmittedviaBypassCSVPersec uint32

	//
	WriteRequeststransmittedviaSMBDirectPersec uint32
}

func NewWin32_PerfFormattedData_Counters_SMBServerSharesEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_SMBServerShares, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_SMBServerShares{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_SMBServerSharesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_SMBServerShares, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_SMBServerShares{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetAvgBytesPerRead sets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyAvgBytesPerRead(value uint64) (err error) {
	return instance.SetProperty("AvgBytesPerRead", value)
}

// GetAvgBytesPerRead gets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyAvgBytesPerRead() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgBytesPerRead")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgBytesPerWrite sets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyAvgBytesPerWrite(value uint64) (err error) {
	return instance.SetProperty("AvgBytesPerWrite", value)
}

// GetAvgBytesPerWrite gets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyAvgBytesPerWrite() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgBytesPerWrite")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgDataBytesPerRequest sets the value of AvgDataBytesPerRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyAvgDataBytesPerRequest(value uint64) (err error) {
	return instance.SetProperty("AvgDataBytesPerRequest", value)
}

// GetAvgDataBytesPerRequest gets the value of AvgDataBytesPerRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyAvgDataBytesPerRequest() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgDataBytesPerRequest")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgDataQueueLength sets the value of AvgDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyAvgDataQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgDataQueueLength", value)
}

// GetAvgDataQueueLength gets the value of AvgDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyAvgDataQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgDataQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgReadQueueLength sets the value of AvgReadQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyAvgReadQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgReadQueueLength", value)
}

// GetAvgReadQueueLength gets the value of AvgReadQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyAvgReadQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgReadQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgsecPerDataRequest sets the value of AvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyAvgsecPerDataRequest(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerDataRequest", value)
}

// GetAvgsecPerDataRequest gets the value of AvgsecPerDataRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyAvgsecPerDataRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerDataRequest")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgsecPerRead sets the value of AvgsecPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyAvgsecPerRead(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerRead", value)
}

// GetAvgsecPerRead gets the value of AvgsecPerRead for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyAvgsecPerRead() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerRead")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgsecPerRequest sets the value of AvgsecPerRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyAvgsecPerRequest(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerRequest", value)
}

// GetAvgsecPerRequest gets the value of AvgsecPerRequest for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyAvgsecPerRequest() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerRequest")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgsecPerWrite sets the value of AvgsecPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyAvgsecPerWrite(value uint32) (err error) {
	return instance.SetProperty("AvgsecPerWrite", value)
}

// GetAvgsecPerWrite gets the value of AvgsecPerWrite for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyAvgsecPerWrite() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgsecPerWrite")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvgWriteQueueLength sets the value of AvgWriteQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyAvgWriteQueueLength(value uint64) (err error) {
	return instance.SetProperty("AvgWriteQueueLength", value)
}

// GetAvgWriteQueueLength gets the value of AvgWriteQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyAvgWriteQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgWriteQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesCompressedPersec sets the value of BytesCompressedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyBytesCompressedPersec(value uint64) (err error) {
	return instance.SetProperty("BytesCompressedPersec", value)
}

// GetBytesCompressedPersec gets the value of BytesCompressedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyBytesCompressedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesCompressedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompressedRequestsPersec sets the value of CompressedRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyCompressedRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("CompressedRequestsPersec", value)
}

// GetCompressedRequestsPersec gets the value of CompressedRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyCompressedRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompressedRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompressedResponsesPersec sets the value of CompressedResponsesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyCompressedResponsesPersec(value uint64) (err error) {
	return instance.SetProperty("CompressedResponsesPersec", value)
}

// GetCompressedResponsesPersec gets the value of CompressedResponsesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyCompressedResponsesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompressedResponsesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentBypassOpenFileCount sets the value of CurrentBypassOpenFileCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyCurrentBypassOpenFileCount(value uint64) (err error) {
	return instance.SetProperty("CurrentBypassOpenFileCount", value)
}

// GetCurrentBypassOpenFileCount gets the value of CurrentBypassOpenFileCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyCurrentBypassOpenFileCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentBypassOpenFileCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentDataQueueLength sets the value of CurrentDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyCurrentDataQueueLength(value uint64) (err error) {
	return instance.SetProperty("CurrentDataQueueLength", value)
}

// GetCurrentDataQueueLength gets the value of CurrentDataQueueLength for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyCurrentDataQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentDataQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentDurableOpenFileCount sets the value of CurrentDurableOpenFileCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyCurrentDurableOpenFileCount(value uint64) (err error) {
	return instance.SetProperty("CurrentDurableOpenFileCount", value)
}

// GetCurrentDurableOpenFileCount gets the value of CurrentDurableOpenFileCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyCurrentDurableOpenFileCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentDurableOpenFileCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentOpenFileCount sets the value of CurrentOpenFileCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyCurrentOpenFileCount(value uint64) (err error) {
	return instance.SetProperty("CurrentOpenFileCount", value)
}

// GetCurrentOpenFileCount gets the value of CurrentOpenFileCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyCurrentOpenFileCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentOpenFileCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentPendingRequests sets the value of CurrentPendingRequests for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyCurrentPendingRequests(value uint64) (err error) {
	return instance.SetProperty("CurrentPendingRequests", value)
}

// GetCurrentPendingRequests gets the value of CurrentPendingRequests for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyCurrentPendingRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentPendingRequests")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataBytesPersec sets the value of DataBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyDataBytesPersec(value uint64) (err error) {
	return instance.SetProperty("DataBytesPersec", value)
}

// GetDataBytesPersec gets the value of DataBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyDataBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DataBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataRequestsPersec sets the value of DataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyDataRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("DataRequestsPersec", value)
}

// GetDataRequestsPersec gets the value of DataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyDataRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DataRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFilesOpenedPersec sets the value of FilesOpenedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyFilesOpenedPersec(value uint64) (err error) {
	return instance.SetProperty("FilesOpenedPersec", value)
}

// GetFilesOpenedPersec gets the value of FilesOpenedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyFilesOpenedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FilesOpenedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMetadataRequestsPersec sets the value of MetadataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyMetadataRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("MetadataRequestsPersec", value)
}

// GetMetadataRequestsPersec gets the value of MetadataRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyMetadataRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MetadataRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentPersistentHandles sets the value of PercentPersistentHandles for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyPercentPersistentHandles(value uint64) (err error) {
	return instance.SetProperty("PercentPersistentHandles", value)
}

// GetPercentPersistentHandles gets the value of PercentPersistentHandles for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyPercentPersistentHandles() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentPersistentHandles")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentResilientHandles sets the value of PercentResilientHandles for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyPercentResilientHandles(value uint64) (err error) {
	return instance.SetProperty("PercentResilientHandles", value)
}

// GetPercentResilientHandles gets the value of PercentResilientHandles for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyPercentResilientHandles() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentResilientHandles")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadBytesPersec sets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytesPersec", value)
}

// GetReadBytesPersec gets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadBytestransmittedByPassCSVPersec sets the value of ReadBytestransmittedByPassCSVPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyReadBytestransmittedByPassCSVPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytestransmittedByPassCSVPersec", value)
}

// GetReadBytestransmittedByPassCSVPersec gets the value of ReadBytestransmittedByPassCSVPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyReadBytestransmittedByPassCSVPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytestransmittedByPassCSVPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadBytestransmittedviaSMBDirectPersec sets the value of ReadBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyReadBytestransmittedviaSMBDirectPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytestransmittedviaSMBDirectPersec", value)
}

// GetReadBytestransmittedviaSMBDirectPersec gets the value of ReadBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyReadBytestransmittedviaSMBDirectPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytestransmittedviaSMBDirectPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadRequestsPersec sets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyReadRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("ReadRequestsPersec", value)
}

// GetReadRequestsPersec gets the value of ReadRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyReadRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReadRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadRequeststransmittedviaBypassCSVPersec sets the value of ReadRequeststransmittedviaBypassCSVPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyReadRequeststransmittedviaBypassCSVPersec(value uint32) (err error) {
	return instance.SetProperty("ReadRequeststransmittedviaBypassCSVPersec", value)
}

// GetReadRequeststransmittedviaBypassCSVPersec gets the value of ReadRequeststransmittedviaBypassCSVPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyReadRequeststransmittedviaBypassCSVPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReadRequeststransmittedviaBypassCSVPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadRequeststransmittedviaSMBDirectPersec sets the value of ReadRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyReadRequeststransmittedviaSMBDirectPersec(value uint32) (err error) {
	return instance.SetProperty("ReadRequeststransmittedviaSMBDirectPersec", value)
}

// GetReadRequeststransmittedviaSMBDirectPersec gets the value of ReadRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyReadRequeststransmittedviaSMBDirectPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReadRequeststransmittedviaSMBDirectPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedBytesPersec sets the value of ReceivedBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyReceivedBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReceivedBytesPersec", value)
}

// GetReceivedBytesPersec gets the value of ReceivedBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyReceivedBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestsPersec sets the value of RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("RequestsPersec", value)
}

// GetRequestsPersec gets the value of RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSentBytesPersec sets the value of SentBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertySentBytesPersec(value uint64) (err error) {
	return instance.SetProperty("SentBytesPersec", value)
}

// GetSentBytesPersec gets the value of SentBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertySentBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalDurableHandleReopenCount sets the value of TotalDurableHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyTotalDurableHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalDurableHandleReopenCount", value)
}

// GetTotalDurableHandleReopenCount gets the value of TotalDurableHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyTotalDurableHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalDurableHandleReopenCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalFailedDurableHandleReopenCount sets the value of TotalFailedDurableHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyTotalFailedDurableHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalFailedDurableHandleReopenCount", value)
}

// GetTotalFailedDurableHandleReopenCount gets the value of TotalFailedDurableHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyTotalFailedDurableHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalFailedDurableHandleReopenCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalFailedPersistentHandleReopenCount sets the value of TotalFailedPersistentHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyTotalFailedPersistentHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalFailedPersistentHandleReopenCount", value)
}

// GetTotalFailedPersistentHandleReopenCount gets the value of TotalFailedPersistentHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyTotalFailedPersistentHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalFailedPersistentHandleReopenCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalFailedResilientHandleReopenCount sets the value of TotalFailedResilientHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyTotalFailedResilientHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalFailedResilientHandleReopenCount", value)
}

// GetTotalFailedResilientHandleReopenCount gets the value of TotalFailedResilientHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyTotalFailedResilientHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalFailedResilientHandleReopenCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalFileOpenCount sets the value of TotalFileOpenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyTotalFileOpenCount(value uint64) (err error) {
	return instance.SetProperty("TotalFileOpenCount", value)
}

// GetTotalFileOpenCount gets the value of TotalFileOpenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyTotalFileOpenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalFileOpenCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalPersistentHandleReopenCount sets the value of TotalPersistentHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyTotalPersistentHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalPersistentHandleReopenCount", value)
}

// GetTotalPersistentHandleReopenCount gets the value of TotalPersistentHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyTotalPersistentHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalPersistentHandleReopenCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalResilientHandleReopenCount sets the value of TotalResilientHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyTotalResilientHandleReopenCount(value uint64) (err error) {
	return instance.SetProperty("TotalResilientHandleReopenCount", value)
}

// GetTotalResilientHandleReopenCount gets the value of TotalResilientHandleReopenCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyTotalResilientHandleReopenCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalResilientHandleReopenCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTransferredBytesPersec sets the value of TransferredBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyTransferredBytesPersec(value uint64) (err error) {
	return instance.SetProperty("TransferredBytesPersec", value)
}

// GetTransferredBytesPersec gets the value of TransferredBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyTransferredBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferredBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTreeConnectCount sets the value of TreeConnectCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyTreeConnectCount(value uint64) (err error) {
	return instance.SetProperty("TreeConnectCount", value)
}

// GetTreeConnectCount gets the value of TreeConnectCount for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyTreeConnectCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TreeConnectCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteBytesPersec sets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytesPersec", value)
}

// GetWriteBytesPersec gets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteBytestransmittedByPassCSVPersec sets the value of WriteBytestransmittedByPassCSVPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyWriteBytestransmittedByPassCSVPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytestransmittedByPassCSVPersec", value)
}

// GetWriteBytestransmittedByPassCSVPersec gets the value of WriteBytestransmittedByPassCSVPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyWriteBytestransmittedByPassCSVPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteBytestransmittedByPassCSVPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteBytestransmittedviaSMBDirectPersec sets the value of WriteBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyWriteBytestransmittedviaSMBDirectPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytestransmittedviaSMBDirectPersec", value)
}

// GetWriteBytestransmittedviaSMBDirectPersec gets the value of WriteBytestransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyWriteBytestransmittedviaSMBDirectPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteBytestransmittedviaSMBDirectPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteRequestsPersec sets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyWriteRequestsPersec(value uint32) (err error) {
	return instance.SetProperty("WriteRequestsPersec", value)
}

// GetWriteRequestsPersec gets the value of WriteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyWriteRequestsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WriteRequestsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteRequeststransmittedviaBypassCSVPersec sets the value of WriteRequeststransmittedviaBypassCSVPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyWriteRequeststransmittedviaBypassCSVPersec(value uint32) (err error) {
	return instance.SetProperty("WriteRequeststransmittedviaBypassCSVPersec", value)
}

// GetWriteRequeststransmittedviaBypassCSVPersec gets the value of WriteRequeststransmittedviaBypassCSVPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyWriteRequeststransmittedviaBypassCSVPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WriteRequeststransmittedviaBypassCSVPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteRequeststransmittedviaSMBDirectPersec sets the value of WriteRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) SetPropertyWriteRequeststransmittedviaSMBDirectPersec(value uint32) (err error) {
	return instance.SetProperty("WriteRequeststransmittedviaSMBDirectPersec", value)
}

// GetWriteRequeststransmittedviaSMBDirectPersec gets the value of WriteRequeststransmittedviaSMBDirectPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_SMBServerShares) GetPropertyWriteRequeststransmittedviaSMBDirectPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WriteRequeststransmittedviaSMBDirectPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
