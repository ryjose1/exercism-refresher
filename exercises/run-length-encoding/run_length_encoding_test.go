package encode

import "testing"

func TestRunLengthEncode(t *testing.T) {
	for _, test := range encodeTests {
		if actual := RunLengthEncode(test.input); actual != test.expected {
			t.Errorf("FAIL %s - RunLengthEncode(%s) = %q, expected %q.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS RunLengthEncode - %s", test.description)
	}
}
func TestRunLengthDecode(t *testing.T) {
	for _, test := range decodeTests {
		if actual := RunLengthDecode(test.input); actual != test.expected {
			t.Errorf("FAIL %s - RunLengthDecode(%s) = %q, expected %q.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS RunLengthDecode - %s", test.description)
	}
}
func TestRunLengthEncodeDecode(t *testing.T) {
	for _, test := range encodeDecodeTests {
		if actual := RunLengthDecode(RunLengthEncode(test.input)); actual != test.expected {
			t.Errorf("FAIL %s - RunLengthDecode(RunLengthEncode(%s)) = %q, expected %q.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS %s", test.description)
	}
}

func TestNextDataChunk(t *testing.T) {
	var nextDataChunkTests = []struct {
		input        string
		expectedPref string
		expectedSuff string
		description  string
	}{
		{"12A2B", "12A", "2B", ""},
		{"AAA2B", "AAA", "2B", ""},
		{"AAAB2C", "AAA", "B2C", ""},
		{"ABC2B", "ABC", "2B", ""},
		{"ABCDDB2C", "ABC", "DDB2C", ""},
		{"ABC", "ABC", "", ""},
		{"AAA", "AAA", "", ""},
		{"2B", "2B", "", ""},
	}
	for _, test := range nextDataChunkTests {
		actualPref, actualSuff, _ := NextDataChunk(test.input)
		if actualPref != test.expectedPref {
			t.Errorf("FAIL %s - NextDataChunk(%s) = %q, expected %q.",
				test.description, test.input, actualPref, test.expectedPref)
		}
		if actualSuff != test.expectedSuff {
			t.Errorf("FAIL %s - NextDataChunk(%s) = %q, expected %q.",
				test.description, test.input, actualSuff, test.expectedSuff)
		}
		t.Logf("PASS NextDataChunk - %s", test.description)

	}
}
