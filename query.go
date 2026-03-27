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
	"math"
	"strconv"
	"strings"
)

// Querier provides the interface to query using a given command and provide
// the resultant string. The command string should include the appropriate
// terminator for the instrument.
type Querier interface {
	Query(ctx context.Context, cmd string) (string, error)
}

// TruncationError is returned by Int when a float value is truncated to an
// integer. The truncated value is still returned alongside this error.
type TruncationError struct {
	Value    int
	Original string
}

func (e *TruncationError) Error() string {
	return fmt.Sprintf("value %s truncated to %d", e.Original, e.Value)
}

// Bool queries a Querier with the given command and returns a bool.
func Bool(ctx context.Context, q Querier, cmd string) (bool, error) {
	s, err := q.Query(ctx, cmd)
	if err != nil {
		return false, fmt.Errorf("querying bool %q: %w", cmd, err)
	}
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "OFF", "0", "FALSE":
		return false, nil
	case "ON", "1", "TRUE":
		return true, nil
	default:
		return false, fmt.Errorf("could not determine boolean status from %q", s)
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
		return 0, fmt.Errorf("querying float64 %q: %w", cmd, err)
	}
	f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0, fmt.Errorf("parsing float64 for %q: %w", cmd, err)
	}
	return f, nil
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
		return 0, fmt.Errorf("querying int %q: %w", cmd, err)
	}
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		// The string might be in scientific notation (e.g., "+1.23E+02").
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, fmt.Errorf("parsing int for %q: %w", cmd, err)
		}
		if f > float64(math.MaxInt) || f < float64(math.MinInt) {
			return 0, fmt.Errorf("value %s overflows int for %q", s, cmd)
		}
		if f != math.Trunc(f) {
			return int(f), &TruncationError{Value: int(f), Original: s}
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
		return "", fmt.Errorf("querying string %q: %w", cmd, err)
	}
	return strings.TrimSpace(s), nil
}

// Stringf queries the querier according to a format specifier and returns a
// string.
func Stringf(ctx context.Context, q Querier, format string, a ...any) (string, error) {
	return String(ctx, q, fmt.Sprintf(format, a...))
}
