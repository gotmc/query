// Copyright (c) 2020 The query developers. All rights reserved.
// Project site: https://github.com/gotmc/query
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package query

import (
	"fmt"
	"strconv"
	"strings"
)

// Querier provides the interface to query using a given command and provide
// the resultant string. The command string should include the appropriate
// terminator for the instrument.
type Querier interface {
	Query(cmd string) (value string, err error)
}

// Bool is used to query a Querier interface and return a bool.
func Bool(q Querier, query string) (bool, error) {
	s, err := q.Query(query)
	if err != nil {
		return false, err
	}
	switch s {
	case "OFF", "0":
		return false, nil
	case "ON", "1":
		return true, nil
	default:
		return false, fmt.Errorf("could not determine boolean status from %s", s)
	}
}

// Float64 is used to query a Querier interface and return a float64.
func Float64(q Querier, query string) (float64, error) {
	s, err := q.Query(query)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

// Int is used to query a Querier interface and return an int.
func Int(q Querier, query string) (int, error) {
	s, err := q.Query(query)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 32)
	return int(i), err
}

// String is used to query a Querier interface and return a string.
func String(q Querier, query string) (string, error) {
	return q.Query(query)
}
