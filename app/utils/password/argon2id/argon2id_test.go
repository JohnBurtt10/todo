package argon2id

import (
	"testing"
)

type generateHashTest struct {
	arg1, arg2 string
}

var generateHashTests = []generateHashTest{
	generateHashTest{"pa$$word", "otherPa$$word"},
	generateHashTest{"$ecretpa$$word", "not$ecretpa$$word"},
	generateHashTest{"1234", "12345"},
}

func TestComparePasswordAndHash(t *testing.T) {
	argon2ID := New()
	for _, test := range generateHashTests {
		hash, err := argon2ID.GenerateFromPassword(test.arg1)
		if err != nil {
			t.Fatal(err)
		}

		err = argon2ID.ComparePasswordAndHash(hash, test.arg1)
		if err != nil {
			t.Error("expected password and hash to match")
		}

		err = argon2ID.ComparePasswordAndHash(hash, test.arg2)
		if err == nil {
			t.Error("expected password and hash to not match")
		}
	}
}
