package utils

import (
	//	"fmt"
	"bytes"
	"strings"
)

type ArrayUtil struct {
}

//传入InterfaceArray 暂时用处不多
func (this *ArrayUtil) InArray(need interface{}, needArr []interface{}) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

//忽略空格去对比
func (this *ArrayUtil) InStringArray(value string, targetStringSlice []string, ignoreSpace bool) bool {
	for _, v := range targetStringSlice {
		if ignoreSpace {
			value = strings.TrimSpace(value)
			v = strings.TrimSpace(v)
		}

		if value == v {
			return true
		}

	}
	return false
}

func (this *ArrayUtil) InUintArray(need uint, needArr []uint) bool {

	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

//取交集
func (this *ArrayUtil) Intersect(slice1 []string, slice2 []string) []string {
	diffSlice := []string{}
	for _, v := range slice1 {
		for _, v2 := range slice2 {
			if v == v2 {
				diffSlice = append(diffSlice, v)
				break
			}
		}
	}

	return diffSlice
}

//合并byteArray
func (this *ArrayUtil) BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
