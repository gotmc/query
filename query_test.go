// Copyright (c) 2020–2023 The query developers. All rights reserved.
// Project site: https://github.com/gotmc/query
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package query

import (
	"testing"
)

type query struct {
	data map[string]string
}

func (q query) Query(cmd string) (string, error) {
	return q.data[cmd], nil
}

func TestBool(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1": "0",
			"cmd2": "OFF",
			"cmd3": "1",
			"cmd4": "ON",
		},
	}
	testCases := []struct {
		cmd         string
		expected    bool
		expectedErr error
	}{
		{"cmd1", false, nil},
		{"cmd2", false, nil},
		{"cmd3", true, nil},
		{"cmd4", true, nil},
	}
	for _, testCase := range testCases {
		got, err := Bool(q, testCase.cmd)
		if err != testCase.expectedErr {
			t.Errorf("wanted err %s / got err %s", testCase.expectedErr, err)
		}
		if got != testCase.expected {
			t.Errorf("wanted %v / got %v", testCase.expected, got)
		}
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
		cmdNum      int
		expected    bool
		expectedErr error
	}{
		{1, false, nil},
		{2, false, nil},
		{3, true, nil},
		{4, true, nil},
	}
	for _, testCase := range testCases {
		got, err := Boolf(q, "cmd%d", testCase.cmdNum)
		if err != testCase.expectedErr {
			t.Errorf("wanted err %s / got err %s", testCase.expectedErr, err)
		}
		if got != testCase.expected {
			t.Errorf("wanted %v / got %v", testCase.expected, got)
		}
	}
}
func TestInt(t *testing.T) {
	q := query{
		data: map[string]string{
			"one": "+1.2300000000000E+02",
			"two": "123",
		},
	}
	testCases := []struct {
		cmd         string
		expected    int
		expectedErr error
	}{
		{"one", 123, nil},
		{"two", 123, nil},
	}
	for _, testCase := range testCases {
		got, err := Int(q, testCase.cmd)
		if err != testCase.expectedErr {
			t.Errorf("wanted err %s / got err %s", testCase.expectedErr, err)
		}
		if got != testCase.expected {
			t.Errorf("wanted %d / got %d", testCase.expected, got)
		}
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
		cmdNum      int
		expected    int
		expectedErr error
	}{
		{1, 123, nil},
		{2, 123, nil},
	}
	for _, testCase := range testCases {
		got, err := Intf(q, "cmd%d", testCase.cmdNum)
		if err != testCase.expectedErr {
			t.Errorf("wanted err %s / got err %s", testCase.expectedErr, err)
		}
		if got != testCase.expected {
			t.Errorf("wanted %d / got %d", testCase.expected, got)
		}
	}
}

func TestString(t *testing.T) {
	q := query{
		data: map[string]string{
			"get_name": "MyName",
			"get_sn":   "MySerialNumber",
		},
	}
	testCases := []struct {
		cmd         string
		expected    string
		expectedErr error
	}{
		{"get_name", "MyName", nil},
		{"get_sn", "MySerialNumber", nil},
	}
	for _, testCase := range testCases {
		got, err := String(q, testCase.cmd)
		if err != testCase.expectedErr {
			t.Errorf("wanted err %s / got err %s", testCase.expectedErr, err)
		}
		if got != testCase.expected {
			t.Errorf("wanted %s / got %s", testCase.expected, got)
		}
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
		cmdNum      int
		expected    string
		expectedErr error
	}{
		{1, "MyName", nil},
		{2, "MySerialNumber", nil},
	}
	for _, testCase := range testCases {
		got, err := Stringf(q, "cmd%d", testCase.cmdNum)
		if err != testCase.expectedErr {
			t.Errorf("wanted err %s / got err %s", testCase.expectedErr, err)
		}
		if got != testCase.expected {
			t.Errorf("wanted %s / got %s", testCase.expected, got)
		}
	}
}

func TestFloat64(t *testing.T) {
	q := query{
		data: map[string]string{
			"cmd1": "1.2345",
			"cmd2": "3.14159",
		},
	}
	testCases := []struct {
		cmd         string
		expected    float64
		expectedErr error
	}{
		{"cmd1", 1.2345, nil},
		{"cmd2", 3.14159, nil},
	}
	for _, testCase := range testCases {
		got, err := Float64(q, testCase.cmd)
		if err != testCase.expectedErr {
			t.Errorf("wanted err %s / got err %s", testCase.expectedErr, err)
		}
		if got != testCase.expected {
			t.Errorf("wanted %v / got %v", testCase.expected, got)
		}
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
		cmdNum      int
		expected    float64
		expectedErr error
	}{
		{1, 1.2345, nil},
		{2, 3.14159, nil},
	}
	for _, testCase := range testCases {
		got, err := Float64f(q, "cmd%d", testCase.cmdNum)
		if err != testCase.expectedErr {
			t.Errorf("wanted err %s / got err %s", testCase.expectedErr, err)
		}
		if got != testCase.expected {
			t.Errorf("wanted %v / got %v", testCase.expected, got)
		}
	}
}
