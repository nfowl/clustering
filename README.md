# DBScan

Reference Implementation of the [DBScan Clustering Algorithm](https://en.wikipedia.org/wiki/DBSCAN)

## Usage

to use this you will need to make a type that implements the ClusterPoint interface that defines 2 functions:

- `Distance(ClusterPoint) float64`: The distance inbetween two points
- `Id() string`: An Id to use as a map index internally

```golang
type Point struct {
  loc   int
  name  string
}

func (p Point) Distance(other ClusterPoint) {
  return p.loc - other.(Point).loc
}

func (p Point) Id() string {
  return p.name
}

cluster, noise := DBScan()
```
