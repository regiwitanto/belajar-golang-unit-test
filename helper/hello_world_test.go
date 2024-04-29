package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Regi",
			request: "Regi",
		},
		{
			name:    "Witanto",
			request: "Witanto",
		},
		{
			name:    "RegiWitantoJohn",
			request: "Regi Witanto John",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Regi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Regi")
		}
	})
	b.Run("Witanto", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Witanto")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Regi")
	}
}

func BenchmarkHelloWorldWitanto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Witanto")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Regi",
			request:  "Regi",
			expected: "Hello Regi",
		},
		{
			name:     "Witanto",
			request:  "Witanto",
			expected: "Hello Witanto",
		},
		{
			name:     "John",
			request:  "John",
			expected: "Hello John",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Regi", func(t *testing.T) {
		result := HelloWorld("Regi")
		require.Equal(t, "Hello Regi", result, "Result must be 'Hello Regi'")
	})
	t.Run("Witanto", func(t *testing.T) {
		result := HelloWorld("Witanto")
		require.Equal(t, "Hello Witanto", result, "Result must be 'Hello Witanto'")
	})
	t.Run("John", func(t *testing.T) {
		result := HelloWorld("John")
		require.Equal(t, "Hello John", result, "Result must be 'Hello John'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Regi")
	require.Equal(t, "Hello Regi", result, "Result must be 'Hello Regi'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Regi")
	require.Equal(t, "Hello Regi", result, "Result must be 'Hello Regi'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Regi")
	assert.Equal(t, "Hello Regi", result, "Result must be 'Hello Regi'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRegi(t *testing.T) {
	result := HelloWorld("Regi")

	if result != "Hello Regi" {
		// error
		t.Error("Result must be 'Hello Regi'")
	}

	fmt.Println("TestHelloWorldRegi Done")
}

func TestHelloWorldJohn(t *testing.T) {
	result := HelloWorld("John")

	if result != "Hello John" {
		// error
		t.Fatal("Result must be 'Hello John'")
	}

	fmt.Println("TestHelloWorldJohn Done")
}

func TestHelloWorldWitanto(t *testing.T) {
	result := HelloWorld("Witanto")

	if result != "Hello Witanto" {
		// error
		t.Fatal("Result must be 'Hello Witanto'")
	}

	fmt.Println("TestHelloWorldWitanto Done")
}
