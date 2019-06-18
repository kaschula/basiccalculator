package testing

import "testing"

func BakeStringEquals(t *testing.T) func(label, expected, actual string) {
	return func(label, expected, actual string) {
		if expected != actual {
			t.Fatalf("Error asserting the %v property. Expected %v to be %v", label, expected, actual)

			return
		}

		t.Logf("%v property successful", label)

		return
	}
}

func BakeIntEquals(t *testing.T) func(label string, expected, actual int) {
	return func(label string, expected, actual int) {
		if expected != actual {
			t.Fatalf("Error asserting the %v property. Expected %v to be %v", label, expected, actual)

			return
		}

		t.Logf("%v property successful", label)

		return
	}
}

func BakeInt64Equals(t *testing.T) func(label string, expected, actual int64) {
	return func(label string, expected, actual int64) {
		if expected != actual {
			t.Fatalf("Error asserting the %v property. Expected %v to be %v", label, expected, actual)

			return
		}

		t.Logf("%v property successful", label)

		return
	}
}

func BakeFloat64Equals(t *testing.T) func(label string, expected, actual float64) {
	return func(label string, expected, actual float64) {
		if expected != actual {
			t.Fatalf("Error asserting the %v property. Expected %v to be %v", label, expected, actual)

			return
		}

		t.Logf("%v property successful", label)

		return
	}
}

func BakeBoolEquals(t *testing.T) func(label string, expected, actual bool) {
	return func(label string, expected, actual bool) {
		if expected != actual {
			t.Fatalf("Error asserting the %v property. Expected %v to be %v", label, expected, actual)

			return
		}

		t.Logf("%v property successful", label)

		return
	}
}
