// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package instance

import (
	"fmt"
	"strings"
	"sync"

	"github.com/microsoft/wmi/pkg/base/credential"
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	wmisession "github.com/microsoft/wmi/pkg/base/session"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

var (
	instanceManagerMap map[string]*WmiInstanceManager
	mutex              sync.Mutex
)

type WmiInstanceManager struct {
	Host      *host.WmiHost
	session   *wmi.WmiSession
	Namespace string
}

func init() {
	instanceManagerMap = map[string]*WmiInstanceManager{}
}

func newWmiInstanceManager(hostname, namespaceName, userName, password, domainName string) (*WmiInstanceManager, error) {
	im := &WmiInstanceManager{
		Host:      host.NewWmiHostWithCredential(hostname, userName, password, domainName),
		Namespace: namespaceName,
	}

	wsession, err := wmisession.GetHostSession(namespaceName, im.Host)
	if err != nil {
		return nil, err
	}
	im.session = wsession

	return im, nil

}

func GetWmiInstanceManager(hostname, namespaceName, userName, password, domainName string) (*WmiInstanceManager, error) {
	mapId := strings.Join([]string{hostname, namespaceName, domainName}, "_")
	if val, ok := instanceManagerMap[mapId]; ok {
		return val, nil
	}

	mutex.Lock()
	defer mutex.Unlock()
	var err error
	instanceManagerMap[mapId], err = newWmiInstanceManager(hostname, namespaceName, userName, password, domainName)
	if err != nil {
		return nil, err
	}
	return instanceManagerMap[mapId], nil

}

func (im *WmiInstanceManager) QueryInstances(queryString string) ([]*wmi.WmiInstance, error) {
	return im.session.QueryInstances(queryString)
}

func (im *WmiInstanceManager) QueryInstanceEx(queryString string) (*wmi.WmiInstance, error) {
	instances, err := im.QueryInstances(queryString)
	if err != nil {
		return nil, err
	}

	if len(instances) == 0 {
		return nil, fmt.Errorf("No Instance Found")
	}

	return instances[0], nil
}

func (im *WmiInstanceManager) QueryInstance(inquery *query.WmiQuery) (*wmi.WmiInstance, error) {
	return im.QueryInstanceEx(inquery.String())
}

func GetWmiInstanceByName(whost *host.WmiHost, namespaceName, className, instanceName string) (*wmi.WmiInstance, error) {
	return GetWmiInstanceEx(whost, namespaceName, query.NewWmiQuery(className, "Name", instanceName))
}

func GetWmiInstanceEx2(hostName string, cred credential.WmiCredential, namespaceName string, inquery *query.WmiQuery) (*wmi.WmiInstance, error) {
	return GetWmiInstance(hostName, namespaceName, cred.UserName, cred.Password, cred.Domain, inquery)
}

func GetWmiInstanceEx(whost *host.WmiHost, namespaceName string, inquery *query.WmiQuery) (*wmi.WmiInstance, error) {
	cred := whost.GetCredential()
	return GetWmiInstance(whost.HostName, namespaceName, cred.UserName, cred.Password, cred.Domain, inquery)
}

func GetWmiInstance(hostname, namespaceName, userName, password, domainName string, inquery *query.WmiQuery) (*wmi.WmiInstance, error) {
	im, err := GetWmiInstanceManager(hostname, namespaceName, userName, password, domainName)
	if err != nil {
		return nil, err
	}
	return im.QueryInstance(inquery)
}