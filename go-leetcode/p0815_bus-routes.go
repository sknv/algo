package leetcode

// https://leetcode.com/problems/bus-routes/description/

type (
	busesSet map[int]struct{}

	visitedStop struct {
		stop  int
		buses int
	}
)

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	// Build a graph of stops to sets of buses
	graph := make(map[int]busesSet)
	for bus, route := range routes {
		for _, stop := range route {
			if buses, ok := graph[stop]; ok {
				buses[bus] = struct{}{}
			} else {
				buses = make(busesSet)
				buses[bus] = struct{}{}
				graph[stop] = buses
			}
		}
	}

	visitedStops := make(map[int]struct{})
	visitedBuses := make(map[int]struct{})

	q := NewDeque[visitedStop](1)
	q.PushBack(visitedStop{
		stop:  source,
		buses: 0,
	})

	for !q.IsEmpty() {
		curStop, _ := q.PopFront()

		if curStop.stop == target {
			return curStop.buses
		}

		for bus := range graph[curStop.stop] {
			if _, ok := visitedBuses[bus]; ok {
				continue
			}

			visitedBuses[bus] = struct{}{}

			for _, stop := range routes[bus] {
				if _, ok := visitedStops[stop]; ok {
					continue
				}

				visitedStops[stop] = struct{}{}
				q.PushBack(visitedStop{
					stop:  stop,
					buses: curStop.buses + 1,
				})
			}
		}
	}

	return -1
}
