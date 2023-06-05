package lib

// Returns the index of an element.
// If not found, -1 is returned.
func Index[T comparable](haystack []T, needle T) int {
	for i, curr := range haystack {
		if curr == needle {
			return i
		}
	}
	return -1
}

// Returns true if the element is found in the slice.
// Returns false if not found.
func Contains[T comparable](haystack []T, needle T) bool {
	return Index(haystack, needle) > -1
}

// Removes the first occurrence of an element from a slice.
// The length of the returned slice will be one less if the element was found.
// The original slice will be modified if the element was found,
// but use the returned slice on the caller side.
// The returned boolean is true if the element was removed, and false otherwise.
func Remove[T comparable](slice []T, element T) ([]T, bool) {
	found := false
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			found = true
			// Remove the element by shifting the remaining elements
			copy(slice[i:], slice[i+1:])
			slice = slice[:len(slice)-1]
			break
		}
	}
	return slice, found
}
