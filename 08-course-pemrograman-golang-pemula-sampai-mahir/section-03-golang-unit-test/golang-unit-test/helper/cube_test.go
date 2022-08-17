package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/* go test -v -run=TestSuccess */
func TestSuccess(t *testing.T) {
	fmt.Println("Start TestSuccess")
	var cube CubeInterface
	cube = &Cube{10, 9, 8}
	if cube.Volume() != 720 {
		t.Fail()
	}
	fmt.Println("Finish TestSuccess")
}

/* go test -v -run=TestFail */
func TestFail(t *testing.T) {
	fmt.Println("Start TestFail")
	var cube CubeInterface
	cube = &Cube{10, 9, 8}
	if cube.Volume() != 721 {
		t.Fail()
	}
	fmt.Println("Finish TestFail")
}

/* go test -v -run=TestFailNow */
func TestFailNow(t *testing.T) {
	fmt.Println("Start TestFailNow")
	var cube CubeInterface
	cube = &Cube{10, 9, 8}
	if cube.Volume() != 725 {
		t.FailNow()
	}
	fmt.Println("Finish TestFailNow")
}

/* go test -v -run=TestError */
func TestError(t *testing.T) {
	fmt.Println("Start TestError")
	var cube CubeInterface
	cube = &Cube{10, 9, 8}
	if cube.Volume() != 700 {
		t.Error("TestError: value must be 700")
	}
	fmt.Println("Finish TestError")
}

/* go test -v -run=TestFatal */
func TestFatal(t *testing.T) {
	fmt.Println("Start TestFatal")
	var cube CubeInterface
	cube = &Cube{10, 9, 8}
	if cube.Volume() != 700 {
		t.Fatal("TestFatal: value must be 700")
	}
	fmt.Println("Finish TestFatal")
}

/* go test -v -run=TestAssert */
func TestAssert(t *testing.T) {
	fmt.Print("Start TestAssert")
	var cube CubeInterface
	cube = &Cube{12, 5, 5}
	assert.Equal(t, float64(100), cube.SurfaceArea(), "Test Assert: Expected value is 100")
	fmt.Println("Finish TestAssert")
}

/* go test -v -run=TestRequire */
func TestRequire(t *testing.T) {
	fmt.Println("Start TestRequire")
	var cube CubeInterface
	cube = &Cube{12, 5, 5}
	require.Equal(t, float64(100), cube.SurfaceArea(), "Test Required: Expected value is 100")
	fmt.Println("Finish TestRequired")
}

/* go test -v -run=TestSubTestCube */
func TestSubTestCube(t *testing.T) {
	fmt.Println("Start TestSubtest")
	var cube CubeInterface
	cube = &Cube{10, 5, 3}
	t.Run("volume", func(t *testing.T) {
		assert.Equal(t, float64(150), cube.Volume(), "TestSubTest: Expected value is 151")
	})
	t.Run("surface", func(t *testing.T) {
		assert.Equal(t, float64(200), cube.SurfaceArea(), "TestSubTest: Expected value is 200")
	})
	fmt.Println("Finish Test Sub")
}

/* go test -v -run=TestTableTestCube */
func TestTableTestCube(t *testing.T) {
	fmt.Println("Start TestTableTestCube")
	tests := []struct {
		Name          string
		ExpectedValue float64
		Type          string
		Cube          Cube
	}{
		{"Test1", 500, "volume", Cube{10, 10, 5}},
		{"Test2", 174, "surface", Cube{9, 5, 3}},
		{"Test3", 104, "surface", Cube{5, 6, 2}},
		{"Test4", 270, "volume", Cube{9, 5, 6}},
		{"Test5", 6, "volume", Cube{1, 1, 6}},
	}
	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {
			var cube CubeInterface
			cube = &test.Cube
			switch test.Type {
			case "volume":
				{
					assert.Equal(t, test.ExpectedValue, cube.Volume(), fmt.Sprintf("Expected value must be: %v", test.ExpectedValue))
				}
			case "surface":
				{
					assert.Equal(t, test.ExpectedValue, cube.SurfaceArea(), fmt.Sprintf("Expected value must be: %v", test.ExpectedValue))
				}
			default:
				{
					//
				}
			}
		})
	}
	fmt.Println("Finish TestTableTestCube")
}

/* go test -v -run=NoTest -bench=BenchmarkTestCube */
func BenchmarkTestCube(b *testing.B) {
	var cube CubeInterface
	cube = &Cube{15, 8, 12}
	for i := 0; i < b.N; i++ {
		cube.SurfaceArea()
	}
}

/* go test -v -run=NoTest -bench=BenchmarkSubTestCube */
func BenchmarkSubTestCube(b *testing.B) {
	var cube CubeInterface
	cube = &Cube{15, 8, 12}
	b.Run("volume", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cube.Volume()
		}
	})
	b.Run("surface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cube.SurfaceArea()
		}
	})
}

/* go test -v=NoTest -bench=BenchmarkTableTestCube */
func BenchmarkTableTestCube(b *testing.B) {
	benchmarks := []struct {
		Name string
		Type string
		Cube Cube
	}{
		{"bench1", "volume", Cube{10, 9, 8}},
		{"bench2", "surface", Cube{12, 5, 8}},
		{"bench3", "volume", Cube{9, 4, 6}},
		{"bench4", "surface", Cube{12, 3, 5}},
		{"bench5", "volume", Cube{8, 5, 9}},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.Name, func(b *testing.B) {
			var cube CubeInterface
			cube = &benchmark.Cube
			switch benchmark.Type {
			case "volume":
				{
					for i := 0; i < b.N; i++ {
						cube.Volume()
					}
				}
			case "surface":
				{
					for i := 0; i < b.N; i++ {
						cube.SurfaceArea()
					}
				}
			default:
				{
					//
				}
			}
		})
	}
}
