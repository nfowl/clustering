package main

import (
	"fmt"
	"math"

	"github.com/nfowl/clustering"
)

type TestPoint struct {
	X    float64
	Y    float64
	Name string
}

func (t TestPoint) Distance(value clustering.ClusterPoint) float64 {
	other := value.(TestPoint)
	return math.Sqrt(math.Pow((t.X-other.X), 2) + math.Pow((t.Y-other.Y), 2))
}

func (t TestPoint) Id() string {
	return t.Name
}

func main() {
	clusters, noise := clustering.DBScan(2, 1,
		TestPoint{0, 0, "test"},
		TestPoint{0, 1, "test1"},
		TestPoint{2, 1, "test2"},
		TestPoint{10, 0, "test3"},
		TestPoint{10, 1, "test4"},
		TestPoint{9, 0, "test5"},
	)
	fmt.Printf("Clusters: %v\n", clusters)
	fmt.Printf("Noise: %v", noise)
}
