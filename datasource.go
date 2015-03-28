// Copyright 2015 tsuru-autoscale authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// dataSource represents a data source.
type dataSource interface {
	// Get gets data from data source and
	// parses the JSON-encoded data and stores the result
	// in the value pointed to by v.
	Get(v interface{}) error
}

type dataSourceFactory func(conf map[string]interface{}) (dataSource, error)

var dataSources = make(map[string]dataSourceFactory)

// Register registers a new dataSource.
func Register(name string, ds dataSourceFactory) {
	dataSources[name] = ds
}

// NewDataSource creates a new data source instance.
func NewDataSource(name string, conf map[string]interface{}) (dataSource, error) {
	return dataSources[name](conf)
}
