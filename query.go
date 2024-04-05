// Copyright (c) 2020–2024 The query developers. All rights reserved.
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
	Query(cmd string) (string, error)
}

// Bool queries a Querier with the given command and returns a bool.
func Bool(q Querier, cmd string) (bool, error) {
	s, err := q.Query(cmd)
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

// Boolf queries the Querier according to a format specifier and returns a
// bool.
func Boolf(q Querier, format string, a ...interface{}) (bool, error) {
	return Bool(q, fmt.Sprintf(format, a...))
}

// Float64 queries the Querier with the given command and returns a float64.
func Float64(q Querier, cmd string) (float64, error) {
	s, err := q.Query(cmd)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

// Float64f queries the querier according to a format specifier and returns a
// float.
func Float64f(q Querier, format string, a ...interface{}) (float64, error) {
	return Float64(q, fmt.Sprintf(format, a...))
}

// Int queries the querier with the given command and returns an int.
func Int(q Querier, cmd string) (int, error) {
	s, err := q.Query(cmd)
	if err != nil {
		return 0, err
	}
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		// The string might be formatted to in scientific format.
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}
		return int(f), nil
	}
	return i, nil
}

// Intf queries the querier according to a format specifier and returns a
// int.
func Intf(q Querier, format string, a ...interface{}) (int, error) {
	return Int(q, fmt.Sprintf(format, a...))
}

// String queries the querier with the given command and returns a string.
func String(q Querier, cmd string) (string, error) {
	return q.Query(cmd)
}

// Stringf queries the querier according to a format specifier and returns a
// string.
func Stringf(q Querier, format string, a ...interface{}) (string, error) {
	return String(q, fmt.Sprintf(format, a...))
}
