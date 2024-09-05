package main

import "testing"

// 10 ns
func BenchmarkIndex(b *testing.B) {
	b.StopTimer()

	column := []string{"id", "name", "age", "sex", "phone"}
	index := []string{"id", "name", "phone", "name", "sex"} //first 3 fields is cluster index
	userField := []string{"id", "name", "phone"}

	f := NewIndexBuilder(column, index)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		f.Choice(userField)
	}
}
