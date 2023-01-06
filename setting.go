package models

import ()

//const (
//	SETTING_TYPE_TEXT     = 1
//	SETTING_TYPE_MULTIPLE = 2
//	SETTING_TYPE_BOOL     = 3
//)

type Setting struct {
	ID    int
	Title string
	Key   string
	Desc  string
	Value string
	//	Type  int
}
