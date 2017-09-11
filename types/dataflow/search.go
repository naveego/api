package dataflow

import "time"

// SearchRequest represents a request to search the logs
type SearchRequest struct {
	TenantID string    `json:"-"`
	Type     string    `json:"type"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Query    string    `json:"query"`
	Size     int32     `json:"size"`
	Aggs     bool      `json:"aggs"`
}

// SearchResult represents the output of a search operation
type SearchResult struct {
	Hits      int64         `json:"hits"`
	Logs      []interface{} `json:"logs"`
	Query     string        `json:"query"`
	StartDate time.Time     `json:"start_date"`
	EndDate   time.Time     `json:"end_date"`
	Buckets   []Bucket      `json:"buckets"`
}

// Bucket is an aggregation of log metrics over time
type Bucket struct {
	Key         string       `json:"key"`
	Count       int32        `json:"count"`
	SizeInBytes int64        `json:"size_in_bytes"`
	LevelCounts []LevelCount `json:"level_counts"`
}

// LevelCount is a break out of the documents included by level
type LevelCount struct {
	Level string `json:"level"`
	Count int32  `json:"count"`
}
