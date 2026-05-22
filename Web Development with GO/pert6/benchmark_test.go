package main

import "testing"

func BenchmarkCalculateFinalScore(b *testing.B) {
	student := Student{Name: "Cindy", UTS: 80, UAS: 85, Tugas: 80}

	for i := 0; i < b.N; i++ {
		student.CalculateFinalScore()
	}
}
