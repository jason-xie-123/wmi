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

// MDM_Firewall_PublicProfile02 struct
type MDM_Firewall_PublicProfile02 struct {
	*cim.WmiInstance

	//
	AllowLocalIpsecPolicyMerge bool

	//
	AllowLocalPolicyMerge bool

	//
	AuthAppsAllowUserPrefMerge bool

	//
	DefaultInboundAction int32

	//
	DefaultOutboundAction int32

	//
	DisableInboundNotifications bool

	//
	DisableStealthMode bool

	//
	DisableStealthModeIpsecSecuredPacketExemption bool

	//
	DisableUnicastResponsesToMulticastBroadcast bool

	//
	EnableFirewall bool

	//
	GlobalPortsAllowUserPrefMerge bool

	//
	InstanceID string

	//
	ParentID string

	//
	Shielded bool
}

func NewMDM_Firewall_PublicProfile02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Firewall_PublicProfile02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Firewall_PublicProfile02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Firewall_PublicProfile02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Firewall_PublicProfile02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Firewall_PublicProfile02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowLocalIpsecPolicyMerge sets the value of AllowLocalIpsecPolicyMerge for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyAllowLocalIpsecPolicyMerge(value bool) (err error) {
	return instance.SetProperty("AllowLocalIpsecPolicyMerge", value)
}

// GetAllowLocalIpsecPolicyMerge gets the value of AllowLocalIpsecPolicyMerge for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyAllowLocalIpsecPolicyMerge() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowLocalIpsecPolicyMerge")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowLocalPolicyMerge sets the value of AllowLocalPolicyMerge for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyAllowLocalPolicyMerge(value bool) (err error) {
	return instance.SetProperty("AllowLocalPolicyMerge", value)
}

// GetAllowLocalPolicyMerge gets the value of AllowLocalPolicyMerge for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyAllowLocalPolicyMerge() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowLocalPolicyMerge")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAuthAppsAllowUserPrefMerge sets the value of AuthAppsAllowUserPrefMerge for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyAuthAppsAllowUserPrefMerge(value bool) (err error) {
	return instance.SetProperty("AuthAppsAllowUserPrefMerge", value)
}

// GetAuthAppsAllowUserPrefMerge gets the value of AuthAppsAllowUserPrefMerge for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyAuthAppsAllowUserPrefMerge() (value bool, err error) {
	retValue, err := instance.GetProperty("AuthAppsAllowUserPrefMerge")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultInboundAction sets the value of DefaultInboundAction for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyDefaultInboundAction(value int32) (err error) {
	return instance.SetProperty("DefaultInboundAction", value)
}

// GetDefaultInboundAction gets the value of DefaultInboundAction for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyDefaultInboundAction() (value int32, err error) {
	retValue, err := instance.GetProperty("DefaultInboundAction")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultOutboundAction sets the value of DefaultOutboundAction for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyDefaultOutboundAction(value int32) (err error) {
	return instance.SetProperty("DefaultOutboundAction", value)
}

// GetDefaultOutboundAction gets the value of DefaultOutboundAction for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyDefaultOutboundAction() (value int32, err error) {
	retValue, err := instance.GetProperty("DefaultOutboundAction")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableInboundNotifications sets the value of DisableInboundNotifications for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyDisableInboundNotifications(value bool) (err error) {
	return instance.SetProperty("DisableInboundNotifications", value)
}

// GetDisableInboundNotifications gets the value of DisableInboundNotifications for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyDisableInboundNotifications() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableInboundNotifications")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableStealthMode sets the value of DisableStealthMode for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyDisableStealthMode(value bool) (err error) {
	return instance.SetProperty("DisableStealthMode", value)
}

// GetDisableStealthMode gets the value of DisableStealthMode for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyDisableStealthMode() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableStealthMode")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableStealthModeIpsecSecuredPacketExemption sets the value of DisableStealthModeIpsecSecuredPacketExemption for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyDisableStealthModeIpsecSecuredPacketExemption(value bool) (err error) {
	return instance.SetProperty("DisableStealthModeIpsecSecuredPacketExemption", value)
}

// GetDisableStealthModeIpsecSecuredPacketExemption gets the value of DisableStealthModeIpsecSecuredPacketExemption for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyDisableStealthModeIpsecSecuredPacketExemption() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableStealthModeIpsecSecuredPacketExemption")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableUnicastResponsesToMulticastBroadcast sets the value of DisableUnicastResponsesToMulticastBroadcast for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyDisableUnicastResponsesToMulticastBroadcast(value bool) (err error) {
	return instance.SetProperty("DisableUnicastResponsesToMulticastBroadcast", value)
}

// GetDisableUnicastResponsesToMulticastBroadcast gets the value of DisableUnicastResponsesToMulticastBroadcast for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyDisableUnicastResponsesToMulticastBroadcast() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableUnicastResponsesToMulticastBroadcast")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableFirewall sets the value of EnableFirewall for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyEnableFirewall(value bool) (err error) {
	return instance.SetProperty("EnableFirewall", value)
}

// GetEnableFirewall gets the value of EnableFirewall for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyEnableFirewall() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableFirewall")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGlobalPortsAllowUserPrefMerge sets the value of GlobalPortsAllowUserPrefMerge for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyGlobalPortsAllowUserPrefMerge(value bool) (err error) {
	return instance.SetProperty("GlobalPortsAllowUserPrefMerge", value)
}

// GetGlobalPortsAllowUserPrefMerge gets the value of GlobalPortsAllowUserPrefMerge for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyGlobalPortsAllowUserPrefMerge() (value bool, err error) {
	retValue, err := instance.GetProperty("GlobalPortsAllowUserPrefMerge")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Firewall_PublicProfile02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyParentID() (value string, err error) {
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

// SetShielded sets the value of Shielded for the instance
func (instance *MDM_Firewall_PublicProfile02) SetPropertyShielded(value bool) (err error) {
	return instance.SetProperty("Shielded", value)
}

// GetShielded gets the value of Shielded for the instance
func (instance *MDM_Firewall_PublicProfile02) GetPropertyShielded() (value bool, err error) {
	retValue, err := instance.GetProperty("Shielded")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}