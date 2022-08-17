package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Start unit test")
	m.Run()
	fmt.Println("Finish unit test")
}

func TestHelloNoval(t *testing.T) {
	result := Hello("noval")
	if result != "Hello novalx" {
		t.Error("Result must be 'Hello noval'")
	}
	fmt.Println("Finish TestHelloNoval")
}

func TestHelloNovalAssert(t *testing.T) {
	result := Hello("noval")
	assert.Equal(t, "Hello novalx", result, "Result must be 'Hello noval'")
	fmt.Println("Finish TestHelloNovalAssert")
}

func TestHelloWardhana(t *testing.T) {
	result := Hello("Kusuma")
	if result != "Hello Kusumax" {
		t.Fatal("Result must be 'Hello Kusuma'")
	}
	fmt.Println("Finish TestHelloWardhana")
}

func TestHelloWardhanaRequire(t *testing.T) {
	result := Hello("Wardhana")
	require.Equal(t, "Hello Wardhana", result, "Result must be 'Hello Wardhana'")
	fmt.Println("Finish TestHelloWardhanaRequire")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip()
	}
	result := Hello("Wardhana")
	assert.Equal(t, "Hello Wardhana", result)
	fmt.Println("Finish TestSkip")
}

func TestSubTest(t *testing.T) {
	t.Run("Noval", func(t *testing.T) {
		result := Hello("Noval")
		require.Equal(t, "Hello Noval", result)
	})
	t.Run("Kusuma", func(t *testing.T) {
		result := Hello("Kusumaa")
		require.Equal(t, "Hello Kusuma", result)
	})
	t.Run("Wardhana", func(t *testing.T) {
		result := Hello("Wardhana")
		require.Equal(t, "Hello Wardhana", result)
	})
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		Name           string
		Request        string
		ExpectedResult string
	}{
		{"Noval", "Noval", "Hello Noval"},
		{"Kusuma", "Kusuma", "Hello Kusuma"},
		{"Wardhana", "Wardhana", "Hello Wardhana"},
	}

	for _, test := range tests {
		result := Hello(test.Request)
		t.Run(test.Name, func(t *testing.T) {
			require.Equal(t, test.ExpectedResult, result)
		})
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("noval")
	}
}

func BenchmarkSubTest(b *testing.B) {
	b.Run("Noval", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Noval")
		}
	})
	b.Run("Kusuma", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Kusuma")
		}
	})
	b.Run("Wardhana", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Wardhana")
		}
	})
}

func BenchmarkTableTest(b *testing.B) {
	tests := []struct {
		Name    string
		Request string
	}{
		{"Noval", "Noval"},
		{"Kusuma", "Kusuma"},
		{"Wardhana", "Wardhana"},
	}

	for _, test := range tests {
		b.Run(test.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Hello(test.Request)
			}
		})
	}
}
