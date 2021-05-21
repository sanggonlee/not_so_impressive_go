package asyncwriter_test

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/sanggonlee/not_so_impressive_go_concurrency/asyncwriter"
)

const tempDir = "tmp"

func TestMain(m *testing.M) {
	cleanup := setupTempDir()
	exitCode := m.Run()
	cleanup()
	os.Exit(exitCode)
}

func BenchmarkAsyncMultiWriter_Write_3_writers(b *testing.B) {
	f1, _ := os.Create(filepath.Join(tempDir, filepath.Base("1.txt")))
	f2, _ := os.Create(filepath.Join(tempDir, filepath.Base("2.txt")))
	f3, _ := os.Create(filepath.Join(tempDir, filepath.Base("3.txt")))
	writer := asyncwriter.AsyncMultiWriter(
		f1, f2, f3,
	)
	for i := 0; i < b.N; i++ {
		_, err := writer.Write(testInputBytes)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMultiWriter_Write_3_writers(b *testing.B) {
	f1, _ := os.Create(filepath.Join(tempDir, filepath.Base("1.txt")))
	f2, _ := os.Create(filepath.Join(tempDir, filepath.Base("2.txt")))
	f3, _ := os.Create(filepath.Join(tempDir, filepath.Base("3.txt")))
	writer := io.MultiWriter(
		f1, f2, f3,
	)
	for i := 0; i < b.N; i++ {
		_, err := writer.Write(testInputBytes)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkAsyncMultiWriter_Write_18_writers(b *testing.B) {
	f1, _ := os.Create(filepath.Join(tempDir, filepath.Base("1.txt")))
	f2, _ := os.Create(filepath.Join(tempDir, filepath.Base("2.txt")))
	f3, _ := os.Create(filepath.Join(tempDir, filepath.Base("3.txt")))
	f4, _ := os.Create(filepath.Join(tempDir, filepath.Base("4.txt")))
	f5, _ := os.Create(filepath.Join(tempDir, filepath.Base("5.txt")))
	f6, _ := os.Create(filepath.Join(tempDir, filepath.Base("6.txt")))
	f7, _ := os.Create(filepath.Join(tempDir, filepath.Base("7.txt")))
	f8, _ := os.Create(filepath.Join(tempDir, filepath.Base("8.txt")))
	f9, _ := os.Create(filepath.Join(tempDir, filepath.Base("9.txt")))
	writer := asyncwriter.AsyncMultiWriter(
		f1, f2, f3, f4, f5, f6, f7, f8, f9,
		f1, f2, f3, f4, f5, f6, f7, f8, f9,
	)
	for i := 0; i < b.N; i++ {
		_, err := writer.Write(testInputBytes)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMultiWriter_Write_18_writers(b *testing.B) {
	f1, _ := os.Create(filepath.Join(tempDir, filepath.Base("1.txt")))
	f2, _ := os.Create(filepath.Join(tempDir, filepath.Base("2.txt")))
	f3, _ := os.Create(filepath.Join(tempDir, filepath.Base("3.txt")))
	f4, _ := os.Create(filepath.Join(tempDir, filepath.Base("4.txt")))
	f5, _ := os.Create(filepath.Join(tempDir, filepath.Base("5.txt")))
	f6, _ := os.Create(filepath.Join(tempDir, filepath.Base("6.txt")))
	f7, _ := os.Create(filepath.Join(tempDir, filepath.Base("7.txt")))
	f8, _ := os.Create(filepath.Join(tempDir, filepath.Base("8.txt")))
	f9, _ := os.Create(filepath.Join(tempDir, filepath.Base("9.txt")))
	writer := io.MultiWriter(
		f1, f2, f3, f4, f5, f6, f7, f8, f9,
		f1, f2, f3, f4, f5, f6, f7, f8, f9,
	)
	for i := 0; i < b.N; i++ {
		_, err := writer.Write(testInputBytes)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkAsyncMultiWriter_Write_20_writers(b *testing.B) {
	f1, _ := os.Create(filepath.Join(tempDir, filepath.Base("1.txt")))
	f2, _ := os.Create(filepath.Join(tempDir, filepath.Base("2.txt")))
	f3, _ := os.Create(filepath.Join(tempDir, filepath.Base("3.txt")))
	f4, _ := os.Create(filepath.Join(tempDir, filepath.Base("4.txt")))
	f5, _ := os.Create(filepath.Join(tempDir, filepath.Base("5.txt")))
	f6, _ := os.Create(filepath.Join(tempDir, filepath.Base("6.txt")))
	f7, _ := os.Create(filepath.Join(tempDir, filepath.Base("7.txt")))
	f8, _ := os.Create(filepath.Join(tempDir, filepath.Base("8.txt")))
	f9, _ := os.Create(filepath.Join(tempDir, filepath.Base("9.txt")))
	f10, _ := os.Create(filepath.Join(tempDir, filepath.Base("10.txt")))
	writer := asyncwriter.AsyncMultiWriter(
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10,
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10,
	)
	for i := 0; i < b.N; i++ {
		_, err := writer.Write(testInputBytes)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMultiWriter_Write_20_writers(b *testing.B) {
	f1, _ := os.Create(filepath.Join(tempDir, filepath.Base("1.txt")))
	f2, _ := os.Create(filepath.Join(tempDir, filepath.Base("2.txt")))
	f3, _ := os.Create(filepath.Join(tempDir, filepath.Base("3.txt")))
	f4, _ := os.Create(filepath.Join(tempDir, filepath.Base("4.txt")))
	f5, _ := os.Create(filepath.Join(tempDir, filepath.Base("5.txt")))
	f6, _ := os.Create(filepath.Join(tempDir, filepath.Base("6.txt")))
	f7, _ := os.Create(filepath.Join(tempDir, filepath.Base("7.txt")))
	f8, _ := os.Create(filepath.Join(tempDir, filepath.Base("8.txt")))
	f9, _ := os.Create(filepath.Join(tempDir, filepath.Base("9.txt")))
	f10, _ := os.Create(filepath.Join(tempDir, filepath.Base("10.txt")))
	writer := io.MultiWriter(
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10,
		f1, f2, f3, f4, f5, f6, f7, f8, f9, f10,
	)
	for i := 0; i < b.N; i++ {
		_, err := writer.Write(testInputBytes)
		if err != nil {
			panic(err)
		}
	}
}

func setupTempDir() func() {
	// Workaround for b.TempDir() bug when used in benchmarks: https://github.com/golang/go/issues/41062
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		fmt.Println(err)
		if err = os.Mkdir(tempDir, 0700); err != nil {
			fmt.Println(err)
		}
	}

	return func() {
		os.RemoveAll(tempDir)
	}
}

var testInputBytes = []byte(`asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
asdfasdfdsafsdafsdfasdfdsfdsafsda
`)
