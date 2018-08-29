package hashing

import "testing"

func TestHashBase64(t *testing.T) {
	table := [][]string{
		// input,		expected hash
		{"angryMonkey", `ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==`},
		//
		// TODO: more test cases
	}

	for _, testCase := range table {
		if result := HashBase64([]byte(testCase[0])); result != testCase[1] {
			t.Fatalf("Expected hash %q doesn't match calculated hash: %q", testCase[1], result)
		}
	}
}
