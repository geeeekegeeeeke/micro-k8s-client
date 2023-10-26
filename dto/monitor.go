package dto

import "time"

type MonitorSearch struct {
	Param     string    `form:"param" json:"param" validate:"required,oneof=all cpu memory load io network"`
	Info      string    `form:"info" json:"info"`
	StartTime time.Time `form:"startTime" json:"startTime"`
	EndTime   time.Time `form:"endTime" json:"endTime"`
}

type MonitorData struct {
	Param string        `json:"param" validate:"required,oneof=cpu memory load io network"`
	Date  []time.Time   `json:"date"`
	Value []interface{} `json:"value"`
}
