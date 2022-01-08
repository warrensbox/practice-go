package main

func main() {

}

func canTwoMoviesFillFlight(movieLengths []int, flightLength int) bool {

	movieLengthsSeen := make(map[int]bool)

	for _, firstMovieLength := range movieLengths {

		matchingSecondMovieLength := flightLength - firstMovieLength
		if _, ok := movieLengthsSeen[matchingSecondMovieLength]; ok {
			return true
		}

		movieLengthsSeen[firstMovieLength] = true
	}

	return false
}
