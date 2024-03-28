package request

import dto "micro-k8s-client/dto"

type RuleSearch struct {
	dto.PageInfo
	Info string `json:"info"`
	Type string `json:"type" validate:"required"`
}

// Port
type PortRuleOperate struct {
	Operation string `json:"operation" validate:"required,oneof=add remove"`
	Address   string `json:"address"`
	Port      string `json:"port" validate:"required"`
	Protocol  string `json:"protocol" validate:"required,oneof=tcp udp tcp/udp"`
	Strategy  string `json:"strategy" validate:"required,oneof=accept drop"`
}

type PortRuleUpdate struct {
	OldRule PortRuleOperate `json:"oldRule"`
	NewRule PortRuleOperate `json:"newRule"`
}

type BatchRuleOperate struct {
	Type  string            `json:"type" validate:"required"`
	Rules []PortRuleOperate `json:"rules"`
}

// address
type AddrRuleOperate struct {
	Operation string `json:"operation" validate:"required,oneof=add remove"`
	Address   string `json:"address"  validate:"required"`
	Strategy  string `json:"strategy" validate:"required,oneof=accept drop"`
}

type AddrRuleUpdate struct {
	OldRule AddrRuleOperate `json:"oldRule"`
	NewRule AddrRuleOperate `json:"newRule"`
}

type FirewallOperation struct {
	Operation string `json:"operation" validate:"required,oneof=start stop disablePing enablePing"`
}
