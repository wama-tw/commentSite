package controllers

import "OSProject1/src/models"

func merge(left, right []models.Post) (merged []models.Post) {
	for leftIndex, rightIndex := 0, 0; leftIndex < len(left) || rightIndex < len(right); {
		if rightIndex == len(right) {
			for ; leftIndex < len(left); leftIndex++ {
				merged = append(merged, left[leftIndex])
			}
			break
		} else if leftIndex == len(left) {
			for ; rightIndex < len(right); rightIndex++ {
				merged = append(merged, right[rightIndex])
			}
			break
		}

		if right[rightIndex].CreatedAt.UnixNano() > left[leftIndex].CreatedAt.UnixNano() {
			merged = append(merged, right[rightIndex])
			rightIndex++
		} else {
			merged = append(merged, left[leftIndex])
			leftIndex++
		}
	}
	return merged
}

func mergeSort(unsorted []models.Post, c chan []models.Post) {
	if len(unsorted) <= 1 {
		c <- unsorted
		return
	}

	cLeft := make(chan []models.Post)
	cRight := make(chan []models.Post)

	go mergeSort(unsorted[:len(unsorted)/2], cLeft)
	go mergeSort(unsorted[len(unsorted)/2:], cRight)
	c <- merge(<-cLeft, <-cRight)
	return
}
