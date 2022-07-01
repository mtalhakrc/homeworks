package main

import "fmt"

func main() {
	stations := make(map[string]set, 1)
	stations["kone"] = newset("id", "nv", "ut")
	stations["ktwo"] = newset("wa", "id", "mt")
	stations["kthree"] = newset("or", "nv", "ca")
	stations["kfour"] = newset("nv", "ut")
	stations["kfive"] = newset("ca", "az")

	statesNeeded := set{}
	statesNeeded.add("mt", "wa", "or", "id", "nv", "ut", "ca", "az")

	index := 0
	bestStation := ""
	finalStations := make([]string, 5)
	for len(statesNeeded) != 0 {
		statesCovered := newset()
		for station, states := range stations {
			covered := intersectionofSets(statesNeeded, states)
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
			for i := 0; i < len(statesCovered); i++ {
				delete(statesNeeded, statesCovered.getStates()[i])
			}
		}
		finalStations[index] = bestStation
		index++
	}
	fmt.Println(finalStations)
}

type set map[string]bool

func newset(keys ...string) set {
	set := set{}
	for i := 0; i < len(keys); i++ {
		set[keys[i]] = true
	}
	return set
}

func (s set) add(keys ...string) {
	for i := 0; i < len(keys); i++ {
		s[keys[i]] = true
	}
}

func (s set) delete(keys ...string) {
	for i := 0; i < len(keys); i++ {
		delete(s, keys[i])
	}
}

func (s set) printset() string {
	str := ""
	for k, _ := range s {
		str = str + k + "  "
	}
	return str
}

func (s set) getStates() []string {
	result := make([]string, len(s))
	count := 0
	for state, _ := range s {
		result[count] = state
		count++
	}
	return result
}

func intersectionofSets(s1 set, s2 set) set {
	result := newset()
	for index, _ := range s2 {
		for a, _ := range s1 {
			if a == index {
				result[index] = true
			}
		}
	}
	return result
}
