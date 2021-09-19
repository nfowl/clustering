# DBScan

Reference Implementation of the [DBScan Clustering Algorithm](https://en.wikipedia.org/wiki/DBSCAN)

## Usage

to use this you will need to make a type that implements the ClusterPoint interface that defines 2 functions:

- `Distance(ClusterPoint) float64`: The distance inbetween two points
- `Id() string`: An Id to use as a map index internally

You can then use that structure to call the `clustering.DBScan` method with your desired parameters.

See the [examples](examples) folder for more info.