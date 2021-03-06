package main

import (
	"errors"

	"github.com/pingcap/failpoint"
)

func demo() (int, error) {
	failpoint.Inject("failpoint-name", func() {
		failpoint.Return(2, errors.New("hahah"))
	})
	return 1, nil
}

func demo2() (int, error) {
	failpoint.Inject("failpoint-name", func(val failpoint.Value) {
		if val.(bool) {
			failpoint.Return(0, errors.New("hehe"))
		}
	})
	return 1, nil
}
