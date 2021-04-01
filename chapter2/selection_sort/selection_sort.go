package selection_sort

type ArtistInfo struct {
	name string
	playCount int
}

func selectionSort(artists []ArtistInfo) []ArtistInfo {
	for index, value := range artists {
		lessPlayedArtist := value
		indexToReplace := index
		for auxIndex := index + 1; auxIndex < len(artists); auxIndex++ {
			if lessPlayedArtist.playCount > artists[auxIndex].playCount {
				lessPlayedArtist = artists[auxIndex]
				indexToReplace = auxIndex
			}
		}

		aux := artists[indexToReplace]
		artists[indexToReplace] = artists[index]
		artists[index] = aux
	}

	return artists
}