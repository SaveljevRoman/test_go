package main

import (
	"fmt"
	"time"
)

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// начало решения
	res := make([]any, len(funcs))
	for k, f := range funcs {
		res[k] = f()
	}
	return res
	// конец решения
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
