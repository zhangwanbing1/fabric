/*
 * Copyright Greg Haskins All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * See github.com/zhangwanbing1/fabric/test/chaincodes/AutoVendor/chaincode/main.go for details
 */
package directdep

import (
	"chaincodes/AutoVendor/indirectdep"
)

func PointlessFunction() {
	// delegate to our indirect dependency
	indirectdep.PointlessFunction()
}
