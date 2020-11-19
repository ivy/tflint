// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsALBTargetGroupInvalidProtocolRule checks the pattern is valid
type AwsALBTargetGroupInvalidProtocolRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsALBTargetGroupInvalidProtocolRule returns new rule with default attributes
func NewAwsALBTargetGroupInvalidProtocolRule() *AwsALBTargetGroupInvalidProtocolRule {
	return &AwsALBTargetGroupInvalidProtocolRule{
		resourceType:  "aws_alb_target_group",
		attributeName: "protocol",
		enum: []string{
			"HTTP",
			"HTTPS",
			"TCP",
			"TLS",
			"UDP",
			"TCP_UDP",
			"GENEVE",
		},
	}
}

// Name returns the rule name
func (r *AwsALBTargetGroupInvalidProtocolRule) Name() string {
	return "aws_alb_target_group_invalid_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsALBTargetGroupInvalidProtocolRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsALBTargetGroupInvalidProtocolRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsALBTargetGroupInvalidProtocolRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsALBTargetGroupInvalidProtocolRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as protocol`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
