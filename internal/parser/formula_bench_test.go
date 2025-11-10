package parser

import "testing"

func BenchmarkSimpleFormula(b *testing.B) {
	formula := "1 + 2 * 3"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Parser.ParseString("", formula)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkComplexFormula(b *testing.B) {
	formula := "(a + b) * (c - d) / e >= 10 AND x < 20 OR NOT y"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Parser.ParseString("", formula)
		if err != nil {
			b.Fatal(err)
		}
	}
}
