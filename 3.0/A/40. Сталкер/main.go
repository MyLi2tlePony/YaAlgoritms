package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type Phone struct {
	mapsEdges    [][][]int16
	changePoints [][]int16
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var buildingsNumber, mapsNumber int16
	_, _ = fmt.Sscanf(line, "%d %d", &buildingsNumber, &mapsNumber)

	p := Phone{
		mapsEdges:    make([][][]int16, mapsNumber),    // карта - здание1 - здание 2
		changePoints: make([][]int16, buildingsNumber), // здание - карта
	}

	changePointsHash := make([][]bool, buildingsNumber)
	for i := range changePointsHash {
		changePointsHash[i] = make([]bool, mapsNumber)
	}

	for i := int16(0); i < mapsNumber; i++ {
		line, _ = reader.ReadString('\n')

		var edgeNumber int
		_, _ = fmt.Sscanf(line, "%d", &edgeNumber)

		p.mapsEdges[i] = make([][]int16, buildingsNumber)

		for j := 0; j < edgeNumber; j++ {
			line, _ = reader.ReadString('\n')

			var start, end int16
			_, _ = fmt.Sscanf(line, "%d %d", &start, &end)

			start--
			end--

			if !changePointsHash[start][i] {
				changePointsHash[start][i] = true
				p.changePoints[start] = append(p.changePoints[start], i)
			}

			if !changePointsHash[end][i] {
				changePointsHash[end][i] = true
				p.changePoints[end] = append(p.changePoints[end], i)
			}

			p.mapsEdges[i][start] = append(p.mapsEdges[i][start], end)
			p.mapsEdges[i][end] = append(p.mapsEdges[i][end], start)
		}
	}

	hash := make([][]bool, buildingsNumber)
	for j := range hash {
		hash[j] = make([]bool, mapsNumber)
	}

	commonHash := make([][]bool, buildingsNumber)
	for j := range hash {
		commonHash[j] = make([]bool, mapsNumber)
	}

	checked := make([][]bool, buildingsNumber)
	for j := range hash {
		checked[j] = make([]bool, mapsNumber)
	}

	steps := list.New()
	for i := range p.changePoints[0] {
		hash[0][p.changePoints[0][i]] = true
		steps.PushBack([2]int16{0, p.changePoints[0][i]})
	}

	cost := FuncMinPathCost(buildingsNumber-1, &p, steps, checked, commonHash, hash)
	fmt.Println(cost)
}

func FuncMinPathCost(endBuilding int16, p *Phone, steps *list.List, checked, commonHash, hash [][]bool) int {
	cost, minCost, currentCostLen := 0, -1, steps.Len()

	for minCost == -1 && steps.Len() > 0 {
		cost++

		for currentCostLen > 0 {
			current := steps.Front()
			steps.Remove(current)
			currentCostLen--

			bm := current.Value.([2]int16)
			maps := bm[1]
			building := bm[0]

			if checked[building][maps] {
				continue
			}
			checked[building][maps] = true

			for _, toBuilding := range p.mapsEdges[maps][building] {
				if toBuilding == endBuilding {
					minCost = cost
				}

				if !hash[toBuilding][maps] {
					hash[toBuilding][maps] = true

					steps.PushFront([2]int16{toBuilding, maps})
					currentCostLen++
				}
			}

			for _, toMap := range p.changePoints[building] {
				if hash[building][toMap] || commonHash[building][toMap] {
					continue
				}

				commonHash[building][toMap] = true
				steps.PushBack([2]int16{building, toMap})
			}
		}

		currentCostLen = steps.Len()
	}

	return minCost
}

func FindDeepBuildings(building, maps int16, p *Phone, hash [][]bool, action func(int16)) {
	if hash[building][maps] {
		return
	}

	hash[building][maps] = true
	action(building)

	for _, toBuilding := range p.mapsEdges[maps][building] {
		FindDeepBuildings(toBuilding, maps, p, hash, action)
	}
}
