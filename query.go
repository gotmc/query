// Copyright (c) 2020–2026 The query developers. All rights reserved.
// Project site: https://github.com/gotmc/query
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package query provides convenience functions for parsing the string results
// from a Querier interface into Go types such as bool, int, float64, and
// string.
package query

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// Querier provides the interface to query using a given command and provide
// the resultant string. The command string should include the appropriate
// terminator for the instrument.
type Querier interface {
	Query(ctx context.Context, cmd string) (string, error)
}

// Bool queries a Querier with the given command and returns a bool.
func Bool(ctx context.Context, q Querier, cmd string) (bool, error) {
	s, err := q.Query(ctx, cmd)
	if err != nil {
		return false, err
	}
	switch strings.TrimSpace(s) {
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
func Boolf(ctx context.Context, q Querier, format string, a ...any) (bool, error) {
	return Bool(ctx, q, fmt.Sprintf(format, a...))
}

// Float64 queries the Querier with the given command and returns a float64.
func Float64(ctx context.Context, q Querier, cmd string) (float64, error) {
	s, err := q.Query(ctx, cmd)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

// Float64f queries the querier according to a format specifier and returns a
// float.
func Float64f(ctx context.Context, q Querier, format string, a ...any) (float64, error) {
	return Float64(ctx, q, fmt.Sprintf(format, a...))
}

// Int queries the querier with the given command and returns an int.
func Int(ctx context.Context, q Querier, cmd string) (int, error) {
	s, err := q.Query(ctx, cmd)
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
func Intf(ctx context.Context, q Querier, format string, a ...any) (int, error) {
	return Int(ctx, q, fmt.Sprintf(format, a...))
}

// String queries the querier with the given command and returns a string
// trimming any whitespace.
func String(ctx context.Context, q Querier, cmd string) (string, error) {
	s, err := q.Query(ctx, cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(s), nil
}

// Stringf queries the querier according to a format specifier and returns a
// string.
func Stringf(ctx context.Context, q Querier, format string, a ...any) (string, error) {
	return String(ctx, q, fmt.Sprintf(format, a...))
}
