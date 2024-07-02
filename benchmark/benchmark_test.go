package benchmark

import (
	"testing"
)

func init() {

}

func BenchmarkApi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var err = benchmarkApi()
		if err {
			b.Fatal("api error")
		}
	}
}
