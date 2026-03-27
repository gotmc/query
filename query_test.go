// Copyright (c) 2020–2026 The query developers. All rights reserved.
// Project site: https://github.com/gotmc/query
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package query

import (
	"context"
	"errors"
	"testing"
)

type query struct {
	data map[string]string
}

func (q query) Query(_ context.Context, cmd string) (string, error) {
	return q.data[cmd], nil
}

type errQuerier struct{}

func (q errQuerier) Query(_ context.Context, cmd string) (string, error) {
	return "", errors.New("query failed")
}

func TestBool(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1":  "0",
			"cmd2":  "OFF",
			"cmd3":  "1",
			"cmd4":  "ON",
			"cmd5":  "0\n",
			"cmd6":  "1\n",
			"cmd7":  "off",
			"cmd8":  "on",
			"cmd9":  "Off",
			"cmd10": "On",
			"cmd11": "FALSE",
			"cmd12": "TRUE",
			"cmd13": "false",
			"cmd14": "true",
			"cmd15": "False",
			"cmd16": "True",
		},
	}
	testCases := []struct {
		name     string
		cmd      string
		expected bool
		wantErr  bool
	}{
		{"zero", "cmd1", false, false},
		{"OFF", "cmd2", false, false},
		{"one", "cmd3", true, false},
		{"ON", "cmd4", true, false},
		{"zero with newline", "cmd5", false, false},
		{"one with newline", "cmd6", true, false},
		{"off lowercase", "cmd7", false, false},
		{"on lowercase", "cmd8", true, false},
		{"Off mixed case", "cmd9", false, false},
		{"On mixed case", "cmd10", true, false},
		{"FALSE", "cmd11", false, false},
		{"TRUE", "cmd12", true, false},
		{"false lowercase", "cmd13", false, false},
		{"true lowercase", "cmd14", true, false},
		{"False mixed case", "cmd15", false, false},
		{"True mixed case", "cmd16", true, false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Bool(context.Background(), q, tc.cmd)
			if (err != nil) != tc.wantErr {
				t.Errorf("Bool() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.expected {
				t.Errorf("Bool() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestBool_invalidValue(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1": "YES",
			"cmd2": "2",
			"cmd3": "maybe",
		},
	}
	for _, cmd := range []string{"cmd1", "cmd2", "cmd3"} {
		t.Run(cmd, func(t *testing.T) {
			_, err := Bool(context.Background(), q, cmd)
			if err == nil {
				t.Error("Bool() expected error for unrecognized value, got nil")
			}
		})
	}
}

func TestBool_queryError(t *testing.T) {
	_, err := Bool(context.Background(), errQuerier{}, "cmd")
	if err == nil {
		t.Error("Bool() expected error from querier, got nil")
	}
}

func TestBoolf(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1": "0",
			"cmd2": "OFF",
			"cmd3": "1",
			"cmd4": "ON",
		},
	}
	testCases := []struct {
		name     string
		cmdNum   int
		expected bool
	}{
		{"zero", 1, false},
		{"OFF", 2, false},
		{"one", 3, true},
		{"ON", 4, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Boolf(context.Background(), q, "cmd%d", tc.cmdNum)
			if err != nil {
				t.Errorf("Boolf() unexpected error: %v", err)
				return
			}
			if got != tc.expected {
				t.Errorf("Boolf() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestInt(t *testing.T) {
	q := query{
		data: map[string]string{
			"one":   "+1.2300000000000E+02",
			"two":   "123",
			"three": "  456\n",
		},
	}
	testCases := []struct {
		name     string
		cmd      string
		expected int
	}{
		{"scientific notation", "one", 123},
		{"plain integer", "two", 123},
		{"whitespace trimmed", "three", 456},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Int(context.Background(), q, tc.cmd)
			if err != nil {
				t.Errorf("Int() unexpected error: %v", err)
				return
			}
			if got != tc.expected {
				t.Errorf("Int() = %d, want %d", got, tc.expected)
			}
		})
	}
}

func TestInt_invalidValue(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1": "abc",
		},
	}
	_, err := Int(context.Background(), q, "cmd1")
	if err == nil {
		t.Error("Int() expected error for non-numeric string, got nil")
	}
}

func TestInt_truncated(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1": "1.9",
			"cmd2": "3.14",
			"cmd3": "-2.7",
		},
	}
	testCases := []struct {
		name     string
		cmd      string
		expected int
	}{
		{"1.9 truncates to 1", "cmd1", 1},
		{"3.14 truncates to 3", "cmd2", 3},
		{"-2.7 truncates to -2", "cmd3", -2},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Int(context.Background(), q, tc.cmd)
			var truncErr *TruncationError
			if !errors.As(err, &truncErr) {
				t.Fatalf("Int() expected TruncationError, got %v", err)
			}
			if got != tc.expected {
				t.Errorf("Int() = %d, want %d", got, tc.expected)
			}
			if truncErr.Value != tc.expected {
				t.Errorf("TruncationError.Value = %d, want %d", truncErr.Value, tc.expected)
			}
		})
	}
}

func TestInt_queryError(t *testing.T) {
	_, err := Int(context.Background(), errQuerier{}, "cmd")
	if err == nil {
		t.Error("Int() expected error from querier, got nil")
	}
}

func TestIntf(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1": "+1.2300000000000E+02",
			"cmd2": "123",
		},
	}
	testCases := []struct {
		name     string
		cmdNum   int
		expected int
	}{
		{"scientific notation", 1, 123},
		{"plain integer", 2, 123},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Intf(context.Background(), q, "cmd%d", tc.cmdNum)
			if err != nil {
				t.Errorf("Intf() unexpected error: %v", err)
				return
			}
			if got != tc.expected {
				t.Errorf("Intf() = %d, want %d", got, tc.expected)
			}
		})
	}
}

func TestString(t *testing.T) {
	q := query{
		data: map[string]string{
			"get_name": "MyName",
			"get_sn":   "MySerialNumber",
			"get_ws":   "  trimmed  \n",
		},
	}
	testCases := []struct {
		name     string
		cmd      string
		expected string
	}{
		{"name", "get_name", "MyName"},
		{"serial number", "get_sn", "MySerialNumber"},
		{"whitespace trimmed", "get_ws", "trimmed"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := String(context.Background(), q, tc.cmd)
			if err != nil {
				t.Errorf("String() unexpected error: %v", err)
				return
			}
			if got != tc.expected {
				t.Errorf("String() = %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestString_queryError(t *testing.T) {
	_, err := String(context.Background(), errQuerier{}, "cmd")
	if err == nil {
		t.Error("String() expected error from querier, got nil")
	}
}

func TestStringf(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1": "MyName",
			"cmd2": "MySerialNumber",
		},
	}
	testCases := []struct {
		name     string
		cmdNum   int
		expected string
	}{
		{"name", 1, "MyName"},
		{"serial number", 2, "MySerialNumber"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Stringf(context.Background(), q, "cmd%d", tc.cmdNum)
			if err != nil {
				t.Errorf("Stringf() unexpected error: %v", err)
				return
			}
			if got != tc.expected {
				t.Errorf("Stringf() = %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1": "1.2345",
			"cmd2": "3.14159",
			"cmd3": "  2.718\n",
		},
	}
	testCases := []struct {
		name     string
		cmd      string
		expected float64
	}{
		{"basic float", "cmd1", 1.2345},
		{"pi", "cmd2", 3.14159},
		{"whitespace trimmed", "cmd3", 2.718},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Float64(context.Background(), q, tc.cmd)
			if err != nil {
				t.Errorf("Float64() unexpected error: %v", err)
				return
			}
			if got != tc.expected {
				t.Errorf("Float64() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestFloat64_invalidValue(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1": "not_a_number",
		},
	}
	_, err := Float64(context.Background(), q, "cmd1")
	if err == nil {
		t.Error("Float64() expected error for non-numeric string, got nil")
	}
}

func TestFloat64_queryError(t *testing.T) {
	_, err := Float64(context.Background(), errQuerier{}, "cmd")
	if err == nil {
		t.Error("Float64() expected error from querier, got nil")
	}
}

func TestFloat64f(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1": "1.2345",
			"cmd2": "3.14159",
		},
	}
	testCases := []struct {
		name     string
		cmdNum   int
		expected float64
	}{
		{"basic float", 1, 1.2345},
		{"pi", 2, 3.14159},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Float64f(context.Background(), q, "cmd%d", tc.cmdNum)
			if err != nil {
				t.Errorf("Float64f() unexpected error: %v", err)
				return
			}
			if got != tc.expected {
				t.Errorf("Float64f() = %v, want %v", got, tc.expected)
			}
		})
	}
}
