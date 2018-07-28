/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

// Logical port attachment
type LogicalPortAttachment struct {

	// Indicates the type of logical port attachment. By default it is Virtual Machine interface (VIF)
	AttachmentType string `json:"attachment_type,omitempty"`

	// Extra context data for the attachment
	Context *AttachmentContext `json:"context,omitempty"`

	// Identifier of the interface attached to the logical port
	Id string `json:"id"`
}