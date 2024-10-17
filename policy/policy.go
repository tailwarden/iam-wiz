package policy

import (
	"encoding/json"
	"fmt"
)

// IAMPolicy represents the structure of an AWS IAM policy
type IAMPolicy struct {
	Version   string         `json:"Version"`
	Statement []IAMStatement `json:"Statement"`
}

// IAMStatement represents an individual statement in an IAM policy
type IAMStatement struct {
	Effect    string                 `json:"Effect"`
	Action    interface{}            `json:"Action"`
	Resource  interface{}            `json:"Resource"`
	Condition map[string]interface{} `json:"Condition,omitempty"`
}

// PrettyPrint formats the policy into an indented JSON string
func (p *IAMPolicy) PrettyPrint() (string, error) {
	policyJSON, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshalling policy: %v", err)
	}
	return string(policyJSON), nil
}
