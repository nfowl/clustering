// SPDX-License-Identifier: Apache-2.0

package clustering

type ClusterPoint interface {
	Distance(ClusterPoint) float64
	Id() string
}

type label int

const (
	Unknown label = 0
	Noise   label = 1
	Cluster       = 2
)

// DBScan Implements a reference implementation of the DBScan clustering algorithm see https://en.wikipedia.org/wiki/DBSCAN
// - eps is the epsilon value for 2 points to be considered to be neighbours
// - minPoints is the minimum number of points needed to form a cluster
// - points is the list of ClusterPoint's that the algorithm is being run on
//
// This function can use optimisation improvements
func DBScan(minPoints int, eps float64, points ...ClusterPoint) ([][]ClusterPoint, []ClusterPoint) {
	clusters := make([][]ClusterPoint, 0)
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
		for c := range currentCluster {
			if pointLabel := visited[c.Id()]; pointLabel == Noise {
				visited[c.Id()] = Cluster
			}
			if pointLabel := visited[c.Id()]; pointLabel != Unknown {
				continue SeedLoop
			}
			visited[c.Id()] = Cluster
			newNeighbours := getNeighbours(eps, points, c)
			if len(newNeighbours) >= minPoints {
				for newNeighbour := range newNeighbours {
					currentCluster[newNeighbour] = struct{}{}
				}
			}
		}
		clusterSlice := make([]ClusterPoint, 0)
		for i := range currentCluster {
			clusterSlice = append(clusterSlice, i)
		}
		clusters = append(clusters, clusterSlice)
	}
	noise := make([]ClusterPoint, 0)
	for _, v := range points {
		if visited[v.Id()] == Noise {
			noise = append(noise, v)
		}
	}
	return clusters, noise
}

func getNeighbours(eps float64, points []ClusterPoint, current ClusterPoint) map[ClusterPoint]struct{} {
	neighbours := make(map[ClusterPoint]struct{})
	for _, p := range points {
		if current.Distance(p) <= eps {
			neighbours[p] = struct{}{}
		}
	}
	return neighbours
}
