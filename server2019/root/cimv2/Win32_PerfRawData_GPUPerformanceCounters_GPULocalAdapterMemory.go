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

// Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory struct
type Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory struct {
	*Win32_PerfRawData

	//
	LocalUsage uint64
}

func NewWin32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemoryEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetLocalUsage sets the value of LocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory) SetPropertyLocalUsage(value uint64) (err error) {
	return instance.SetProperty("LocalUsage", value)
}

// GetLocalUsage gets the value of LocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory) GetPropertyLocalUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("LocalUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}