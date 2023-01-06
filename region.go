package models

import ()

type Region struct {
	ID               int    `json:"-"`
	RegionId         string `json:"region_id"`
	EnName           string `json:"region_name"`
	CnName           string `json:"-"`
	IsShowInRecharge bool   `json:"-"`
	CurrencyCode     string `json:"currency_code"`
}
