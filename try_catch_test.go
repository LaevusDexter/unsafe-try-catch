package utc

import (
	"testing"
	"errors"
)

func getError() error {
	return errors.New("Test passed")
}

func testTryCatch1() error {
	var tc TryCatcher

	if err := tc.Catch(); err != nil {
		return err
	} else {
		err = getError()
		tc.Try(err)
	}

	return nil
}

func testTryCatch2() error {
	var tc TryCatcher

	if err := tc.Catch(); err != nil {
		return err
	}

	err := error(nil)
	tc.Try(err)

	err = getError()
	tc.Try(err)

	return nil
}

func TestTryCatch(t *testing.T) {
	if err := testTryCatch1(); err != nil {
		t.Log(err)
	} else {
		t.FailNow()
	}

	if err := testTryCatch2(); err != nil {
		t.Log(err)
	} else {
		t.FailNow()
	}

	var tc TryCatcher
	if err := tc.Catch(); err != nil {
		t.FailNow()
	}

	tc.Try(error(nil))
}

func BenchmarkTryCatch(b *testing.B) {
	var tc TryCatcher

	for i := 0; i < b.N; i++ {

		if err := tc.Catch(); err != nil {
			continue
		}

		tc.Try(getError())
	}
}