package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	universeSize    uint8
	randomnessValue int64
	randValue       uint8
	generations     = 10
)

func userInput() {
	for universeSize < 1 {
		fmt.Scan(&universeSize)
	}

}

func createWorld() [][]string {
	world := make([][]string, universeSize)

	for i := range world {
		world[i] = make([]string, universeSize)
		for j := range world[i] {
			randValue = uint8(rand.Intn(2))
			if randValue != 0 {
				world[i][j] = "O"
			} else {
				world[i][j] = " "
			}
		}

	}
	return world
}

func warpMatrix(index, universeSize int) int {
	return (index%universeSize + universeSize) % universeSize
}

func calculateFuture(world [][]string, generations int) {
	N := len(world)

	directions := []struct{ di, dj int }{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
	for g := 0; g < generations; g++ {
		currentPrint := 0
		newWorld := make([][]string, N)
		for i := range newWorld {
			newWorld[i] = make([]string, N)
			for j := range newWorld[i] {
				newWorld[i][j] = " " // Initialize all cells as dead
			}
		}
		for i, row := range world {
			for j, _ := range row {
				countNeighbors := 0

				for _, dir := range directions {
					ni := warpMatrix(i+dir.di, N)
					nj := warpMatrix(j+dir.dj, N)
					if world[ni][nj] == "O" {
						countNeighbors++
					}

				}
				if world[i][j] == "O" {
					if countNeighbors < 2 || countNeighbors > 3 {
						newWorld[i][j] = " "
					} else {
						newWorld[i][j] = "O"
					}
				} else {
					if countNeighbors == 3 {
						newWorld[i][j] = "O"
					}
				}

			}

		}

		world = newWorld
		alive := 0
		for i := range world {
			for j := range world[i] {
				if world[i][j] == "O" {
					alive++
				}
			}
		}
		currentPrint++
		if currentPrint < generations {
			fmt.Printf("Generation: #%d\nAlive: %d\n", g+1, alive)
			for _, row := range world {

				line := strings.Join(row, "")
				fmt.Println(line)
			}

		}

	}

}

/*
	func presentWorld() {
		worldCurrent := createWorld()

		for _, row := range worldCurrent {
			line := strings.Join(row, "")
			fmt.Println(line)


		worldFuture := calculateFuture(worldCurrent, generations)

}
*/

func main() {

	userInput()
	rand.Seed(randomnessValue)
	calculateFuture(createWorld(), generations)

}
