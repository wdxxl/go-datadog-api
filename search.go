/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import (
	"fmt"
)

// reqSearch is the container for receiving search results.
type reqSearch struct {
	Results struct {
		Hosts   []string `json:"hosts,omitempty"`
		Metrics []string `json:"metrics,omitempty"`
	} `json:"results"`
}

// SearchHosts searches through the hosts facet, returning matching hostnames.
func (client *Client) SearchHosts(search string) ([]string, error) {
	var out reqSearch
	if err := client.doJsonRequest("GET", "/v1/search?q=hosts:"+search, nil, &out); err != nil {
		return nil, err
	}
	return out.Results.Hosts, nil
}

// SearchMetrics searches through the metrics facet, returning matching ones.
func (client *Client) SearchMetrics(search string) ([]string, error) {
	var out reqSearch
	if err := client.doJsonRequest("GET", "/v1/search?q=metrics:"+search, nil, &out); err != nil {
		return nil, err
	}
	return out.Results.Metrics, nil
}

type NewReqSearch struct {
	Monitors []NewMonitor `json:"monitors,omitempty"`
	Metadata Metadata     `json:"metadata,omitempty"`
}

type NewMonitor struct {
	Status               string         `json:"status,omitempty"`
	Scopes               []string       `json:"scopes,omitempty"`
	Classification       string         `json:"classification,omitempty"`
	Creator              Creator        `json:"creator,omitempty"`
	OverallStateModified int            `json:"overall_state_modified,omitempty"`
	Metrics              []string       `json:"metrics,omitempty"`
	Notifications        []Notification `json:"notifications,omitempty"`
	LastTriggeredTs      int            `json:"last_triggered_ts,omitempty"`
	Query                string         `json:"query,omitempty"`
	Id                   int            `json:"id,omitempty"`
	Name                 string         `json:"name,omitempty"`
	Tags                 []string       `json:"tags,omitempty"`
	OrgID                int            `json:"org_id,omitempty"`
	Type                 string         `json:"type,omitempty"`
}

type Notification struct {
	Handle string `json:"handle,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Metadata struct {
	TotalCount int `json:"total_count,omitempty"`
	PageCount  int `json:"page_count,omitempty"`
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
}

// SearchMetrics searches through the metrics facet, returning matching ones.
// https://docs.datadoghq.com/api/?lang=bash#monitors-search
func (client *Client) SearchNotification(search string, page, perPage int) ([]NewMonitor, error) {
	var out NewReqSearch
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/monitor/search?query=notification:%s&page=%d&per_page=%d", search, page, perPage), nil, &out); err != nil {
		return nil, err
	}
	return out.Monitors, nil
}
