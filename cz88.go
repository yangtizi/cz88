package cz88

import (
	"github.com/yangtizi/go/sysutils"
)

// 来自 http://www.cz88.net/ip/

// GetAddress 通过IP获取地址
func GetAddress(strIP string) string {
	//
	// fmt.Println(strIP)
	n := sysutils.IPStrToInt(strIP)
	// fmt.Println(n)
	nLen := uint32(len(AddressList))
	return check(0, nLen, n)
}

// 二分法
func check(nLeft uint32, nRight uint32, nValue uint32) string {
	nMid := (nLeft + nRight) / 2

	// fmt.Println("mid = ", nMid)

	// 值比查询的左边还小, 说明值在更左边
	if AddressList[nMid].Left > nValue {
		return check(nLeft, nMid, nValue)

	}

	if AddressList[nMid].Right < nValue {
		return check(nMid, nRight, nValue)

	}

	//
	// fmt.Println("查到了", nMid,  )
	return AddressList[nMid].S
}

// 获取IP地址缩写
func GetAddressShort(strIP string) string {
	//
	// fmt.Println(strIP)
	n := sysutils.IPStrToInt(strIP)
	// fmt.Println(n)
	nLen := uint32(len(AddressList2))
	return check2(0, nLen, n)
}

// 二分法
func check2(nLeft uint32, nRight uint32, nValue uint32) string {
	nMid := (nLeft + nRight) / 2

	// fmt.Println("mid = ", nMid)

	// 值比查询的左边还小, 说明值在更左边
	if AddressList2[nMid].Left > nValue {
		return check2(nLeft, nMid, nValue)

	}

	if AddressList2[nMid].Right < nValue {
		return check2(nMid, nRight, nValue)

	}

	return AddressList2[nMid].S
}
