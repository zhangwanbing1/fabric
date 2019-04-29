/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package peer

import (
	"github.com/zhangwanbing1/fabric/common/channelconfig"
	"github.com/zhangwanbing1/fabric/core/common/ccprovider"
	"github.com/zhangwanbing1/fabric/core/peer"
)

type MockSupportImpl struct {
	GetApplicationConfigRv     channelconfig.Application
	GetApplicationConfigBoolRv bool
	ChaincodeByNameRv          *ccprovider.ChaincodeData
	ChaincodeByNameBoolRv      bool
}

func (s *MockSupportImpl) GetApplicationConfig(cid string) (channelconfig.Application, bool) {
	return s.GetApplicationConfigRv, s.GetApplicationConfigBoolRv
}

func (s *MockSupportImpl) ChaincodeByName(chainname, ccname string) (*ccprovider.ChaincodeData, bool) {
	return s.ChaincodeByNameRv, s.ChaincodeByNameBoolRv
}

type MockSupportFactoryImpl struct {
	NewSupportRv *MockSupportImpl
}

func (c *MockSupportFactoryImpl) NewSupport() peer.Support {
	return c.NewSupportRv
}
