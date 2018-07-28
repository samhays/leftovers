/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

import (
	"github.com/vmware/go-vmware-nsxt/common"
)

type DneRuleStats struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Aggregated number of incoming bytes processed by the rule.
	BytesIn int64 `json:"bytes_in,omitempty"`

	// Aggregated number of outgoing bytes processed by the rule.
	BytesOut int64 `json:"bytes_out,omitempty"`

	// Aggregated number of incoming packets processed by the rule.
	PacketsIn int64 `json:"packets_in,omitempty"`

	// Aggregated number of outgoing packets processed by the rule.
	PacketsOut int64 `json:"packets_out,omitempty"`

	// Rule Identifier of the DNE rule. This is a globally unique number.
	RuleId string `json:"rule_id,omitempty"`
}