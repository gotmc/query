// Copyright (c) 2020–2023 The query developers. All rights reserved.
// Project site: https://github.com/gotmc/query
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package query

import (
	"testing"
)

type intQuery struct {
	data map[string]string
}

func (q intQuery) Query(cmd string) (string, error) {
	return q.data[cmd], nil
}

func TestInt(t *testing.T) {
	q := intQuery{
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
