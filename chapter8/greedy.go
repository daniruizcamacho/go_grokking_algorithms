package greedy

type station struct {
	name   string
	places []string
}

type stations struct {
	stations []station
}

func (s *stations) findBestStations(places []string) []string {
	bestStations := make([]string, 0)
	coveredPlaces := make([]string, 0)

	for len(places) > 0 && len(bestStations) < len(s.stations) {
		station := s.findBestStation(bestStations)
		bestStations = append(bestStations, station.name)
		coveredPlaces = append(coveredPlaces, station.places...)
		places = cleanPlacesList(places, coveredPlaces)
	}

	return bestStations
}

func (s *stations) findBestStation(visitedStations []string) station {
	bestStation := station{}

	for _, station := range s.stations {
		if !contains(visitedStations, station.name) {
			stationPlaces := station.places
			if len(stationPlaces) > len(bestStation.places) {
				bestStation = station
			}
		}

	}

	return bestStation
}

func cleanPlacesList(places []string, coveredPlaces []string) []string {
	cleanedPlaces := make([]string, 0)

	for _, place := range places {
		if !contains(coveredPlaces, place) {
			cleanedPlaces = append(cleanedPlaces, place)
		}
	}

	return cleanedPlaces
}

func contains(list []string, element string) bool {
	for _, e := range list {
		if e == element {
			return true
		}
	}
	return false
}
