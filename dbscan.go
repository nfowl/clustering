// SPDX-License-Identifier: Apache-2.0

package clustering

// ClusterPoint is the interface that Users are expected to implement on their
// type
type ClusterPoint interface {
	// Distance defines the distance calculation between two points
	Distance(ClusterPoint) float64
	// Id defines the name of the point/object when being indexed
	Id() string
}

type label int

const (
	Unknown label = iota
	Noise
	Cluster
)

// DBScan Implements a reference implementation of the DBScan clustering algorithm see https://en.wikipedia.org/wiki/DBSCAN
// 	eps: The epsilon value for 2 points to be considered to be neighbours
// 	minPoints: The minimum number of points needed to form a cluster
//	points: The list of ClusterPoints that the algorithm is being run on
//
// This function can use optimisation improvements
func DBScan[T ClusterPoint](minPoints int, eps float64, points ...T) ([][]T, []T) {
	clusters := make([][]T, 0)
	visited := make(map[string]label, len(points))

MainLoop:
	for _, p := range points {
		if val := visited[p.Id()]; val != Unknown {
			//Already visited
			continue MainLoop
		}
		neighbours := getNeighbours(eps, points, p)
		if len(neighbours) < minPoints {
			visited[p.Id()] = Noise
			continue MainLoop
		}
		currentCluster := neighbours
		visited[p.Id()] = Cluster

	SeedLoop:
		for _, c := range currentCluster {
			if pointLabel := visited[c.Id()]; pointLabel == Noise {
				visited[c.Id()] = Cluster
			}
			if pointLabel := visited[c.Id()]; pointLabel != Unknown {
				continue SeedLoop
			}
			visited[c.Id()] = Cluster
			newNeighbours := getNeighbours(eps, points, c)
			if len(newNeighbours) >= minPoints {
				for _, newNeighbour := range newNeighbours {
					currentCluster[newNeighbour.Id()] = newNeighbour
				}
			}
		}
		//Translate current cluster to slice for return to caller
		clusterSlice := make([]T, 0)
		for _, val := range currentCluster {
			clusterSlice = append(clusterSlice, val)
		}
		clusters = append(clusters, clusterSlice)
	}
	//Construct Noise slice
	noise := make([]T, 0)
	for _, v := range points {
		if visited[v.Id()] == Noise {
			noise = append(noise, v)
		}
	}
	return clusters, noise
}

func getNeighbours[T ClusterPoint](eps float64, points []T, current T) map[string]T {
	neighbours := make(map[string]T)
	for _, p := range points {
		if current.Distance(p) <= eps {
			neighbours[p.Id()] = p
		}
	}
	return neighbours
}
