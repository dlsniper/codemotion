package bench3_test

import (
	"encoding/json"
	"testing"

	"github.com/dlsniper/codemotion/code/bench/3"
)

var (
	byteEncoded   []byte
	stringEncoded string
)

func TestInjectIndent(t *testing.T) {
	val := bench3.NewBench("demo", 42, json.Marshal)
	res, err := val.ToJSONString2()
	if err != nil {
		t.Logf("Test failed with error: %q", err)
		t.Fail()
	}
	if res != "{\"field_a\":\"demo\",\"field_b\":42}" {
		t.Logf("Test failed, got unexpected result.\nExpecting %q\nGot %q\n", "{\"field_a\":\"demo\",\"field_b\":42}", res)
		t.Fail()
	}

	encoder := func(v interface{}) ([]byte, error) {
		return json.MarshalIndent(v, "", "    ")
	}

	val.SetEncoder(encoder)
	res, err = val.ToJSONString2()
	if err != nil {
		t.Logf("Test failed with error: %q", err)
		t.Fail()
	}
	if res != "{\n    \"field_a\": \"demo\",\n    \"field_b\": 42\n}" {
		t.Logf("Test failed, got unexpected result.\nExpecting %q\nGot %q\n", "{\n    \"field_a\": \"demo\",\n    \"field_b\": 42\n}", res)
		t.Fail()
	}
}

func BenchmarkToJSON(b *testing.B) {
	var (
		vals = make([]*bench3.Bench, b.N)
		res  = []byte{}
		err  error
	)

	for i := 0; i < b.N; i++ {
		vals[i] = bench3.NewPointerBench("demo", 42, json.Marshal)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err = vals[i].ToJSON()
		if err != nil {
			b.Logf("Failed test: %d with error: %q", i, err)
			b.Fail()
		}
	}

	b.StopTimer()

	byteEncoded = res
}

func BenchmarkToJSONString(b *testing.B) {
	var (
		vals = make([]*bench3.Bench, b.N)
		res  = ""
		err  error
	)

	for i := 0; i < b.N; i++ {
		vals[i] = bench3.NewPointerBench("demo", 42, json.Marshal)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err = vals[i].ToJSONString()
		if err != nil {
			b.Logf("Failed test: %d with error: %q", i, err)
			b.Fail()
		}
	}

	b.StopTimer()

	stringEncoded = res
}

func BenchmarkToJSON2(b *testing.B) {
	var (
		vals = make([]bench3.Bench, b.N)
		res  = []byte{}
		err  error
	)

	for i := 0; i < b.N; i++ {
		vals[i] = bench3.NewBench("demo", 42, json.Marshal)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err = vals[i].ToJSON()
		if err != nil {
			b.Logf("Failed test: %d with error: %q", i, err)
			b.Fail()
		}
	}

	b.StopTimer()

	byteEncoded = res
}

func BenchmarkToJSONString2(b *testing.B) {
	var (
		vals = make([]bench3.Bench, b.N)
		res  = ""
		err  error
	)

	for i := 0; i < b.N; i++ {
		vals[i] = bench3.NewBench("demo", 42, json.Marshal)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err = vals[i].ToJSONString2()
		if err != nil {
			b.Logf("Failed test: %d with error: %q", i, err)
			b.Fail()
		}
	}

	b.StopTimer()

	stringEncoded = res
}
