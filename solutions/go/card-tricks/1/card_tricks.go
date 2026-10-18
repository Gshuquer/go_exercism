package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
//	panic("Please implement the FavoriteCards function")
    num := []int{2,6,9}
    return num
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
//	panic("Please implement the GetItem function")
    if index >= len(slice) || index < 0 {
        return -1
    }else{
        return slice[index]
    }
}

// SetItem writes an item to a slice at a given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
//	panic("Please implement the SetItem function")
    if index < len(slice) && index >= 0 {
        slice[index] = value
    }else{
        slice = append(slice,value)
    }
    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
//	panic("Please implement the PrependItems function")
    slice = append(values,slice ...)
    return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
//	panic("Please implement the RemoveItem function")
    if index < len(slice) && index >= 0{
        slice1 := slice[:index]
        slice2 := slice[index+1:]
        slice = append(slice1,slice2...)
        return slice
    }else{
        return slice
    }
}
