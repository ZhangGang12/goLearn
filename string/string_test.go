package string_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var stringls = []string{
	"hellow",
	"word",
	"test",
}

func concatStringByOperator(stringls []string) string {
	var res string
	for _, v := range stringls {
		res = res + v
	}
	return res
}

func BenchmarkConcatStringByOperator(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByOperator(stringls)
	}
}

func concatStringBySprintf(stringls []string) string {
	var res string
	for _, v := range stringls {
		res = fmt.Sprintf("%s%s", res, v)
	}
	return res
}

func BenchmarkConcatStringBySprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringBySprintf(stringls)
	}
}

func concatStringByJoin(stringls []string) string {
	res := strings.Join(stringls, "")
	return res
}

func BenchmarkConcatStringByJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByJoin(stringls)
	}
}

func concatStringByStringBuilder(stringls []string) string {
	var b strings.Builder
	for _, v := range stringls {
		b.WriteString(v)
	}
	return b.String()
}

func BenchmarkConcatStringByStringBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByStringBuilder(stringls)
	}
}

func concatStringByStringBuilderWithSize(stringls []string) string {
	var b strings.Builder
	b.Grow(64)
	for _, v := range stringls {
		b.WriteString(v)
	}
	return b.String()
}

func BenchmarkConcatStringByStringBuilderWithSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByStringBuilderWithSize(stringls)
	}
}

func concatStringByBytesBuffer(stringls []string) string {
	var b bytes.Buffer
	for _, v := range stringls {
		b.WriteString(v)
	}
	return b.String()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByBytesBuffer(stringls)
	}
}

func concatStringByBytesBufferWithSize(stringls []string) string {
	buf := make([]byte, 0, 64)
	b := bytes.NewBuffer(buf)
	for _, v := range stringls {
		b.WriteString(v)
	}
	return b.String()
}

func BenchmarkConcatStringByBytesBufferWithSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByBytesBufferWithSize(stringls)
	}
}
