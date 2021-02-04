/*
* Copyright (c) 2018,2021 ADLINK Technology Inc.
*
* This program and the accompanying materials are made available under the
* terms of the Eclipse Public License 2.0 which is available at
* http://www.eclipse.org/legal/epl-2.0, or the Apache Software License 2.0
* which is available at https://www.apache.org/licenses/LICENSE-2.0.
*
* SPDX-License-Identifier: EPL-2.0 OR Apache-2.0
* Contributors:
*   ADLINK fog05 team, <fog05@adlink-labs.tech>
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/atolab/fog05-manager/pkg/client/clientset/versioned/typed/fdu/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeFog05V1 struct {
	*testing.Fake
}

func (c *FakeFog05V1) FDUs(namespace string) v1.FDUInterface {
	return &FakeFDUs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeFog05V1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}