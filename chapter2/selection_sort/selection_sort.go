package selection_sort

type ArtistInfo struct {
	name string
	playCount int
}

func selectionSort(artists []ArtistInfo) []ArtistInfo {
	artistsCopy := make([]ArtistInfo, len(artists))
	copy(artistsCopy, artists)
	for index, value := range artistsCopy {
		lessPlayedArtist := value
		indexToReplace := index
		for auxIndex := index + 1; auxIndex < len(artistsCopy); auxIndex++ {
			if lessPlayedArtist.playCount > artistsCopy[auxIndex].playCount {
				lessPlayedArtist = artistsCopy[auxIndex]
				indexToReplace = auxIndex
			}
		}

		aux := artistsCopy[indexToReplace]
		artistsCopy[indexToReplace] = artistsCopy[index]
		artistsCopy[index] = aux
	}

	return artistsCopy
}