package egoscale

import (
	"encoding/json"
	"net/url"
	"fmt"
)

func (exo *Client) CreateEgressRule(rule SecurityGroupRule) (*AuthorizeSecurityGroupEgressResponse, error) {

	params := url.Values{}
	params.Set("securitygroupid", rule.SecurityGroupId)
	params.Set("cidrlist", rule.Cidr)
	params.Set("protocol", rule.Protocol)

	if (rule.Protocol == "icmp") {
		params.Set("icmpcode", fmt.Sprintf("%d", rule.IcmpCode))
		params.Set("icmptype", fmt.Sprintf("%d", rule.IcmpType))
	} else {
		params.Set("startport", fmt.Sprintf("%d", rule.Port))
		params.Set("endport", fmt.Sprintf("%d", rule.Port))
	}

	resp, err := exo.Request("authorizeSecurityGroupEgress", params)
	if err != nil {
		return nil, err
	}

	var r AuthorizeSecurityGroupEgressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (exo *Client) CreateIngressRule(rule SecurityGroupRule) (*AuthorizeSecurityGroupIngressResponse, error) {

	params := url.Values{}
	params.Set("securitygroupid", rule.SecurityGroupId)
	params.Set("cidrlist", rule.Cidr)
	params.Set("protocol", rule.Protocol)

	if (rule.Protocol == "icmp") {
		params.Set("icmpcode", fmt.Sprintf("%d", rule.IcmpCode))
		params.Set("icmptype", fmt.Sprintf("%d", rule.IcmpType))
	} else {
		params.Set("startport", fmt.Sprintf("%d", rule.Port))
		params.Set("endport", fmt.Sprintf("%d", rule.Port))
	}

	resp, err := exo.Request("authorizeSecurityGroupIngress", params)

	if err != nil {
		return nil, err
	}

	var r AuthorizeSecurityGroupIngressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (exo *Client) CreateSecurityGroupWithRules(name string, egress []SecurityGroupRule, ingress []SecurityGroupRule) (*CreateSecurityGroupResponse, error) {

	params := url.Values{}
	params.Set("name", name)

	resp, err := exo.Request("createSecurityGroup", params)

	var r CreateSecurityGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	for _, erule := range(egress) {
		_, err = exo.CreateEgressRule(erule)
		if (err != nil) {
			return nil, err
		}
	}

	for _, inrule := range(ingress) {
		_, err = exo.CreateIngressRule(inrule)
		if (err != nil) {
			return nil, err
		}
	}

	return &r, nil
}
