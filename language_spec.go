package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println(">> DAY 10: GO LANGUAGE SPECIFICATION - SLICE INTERNALS")

	// We define a slice
	// According to the Spec, a slice header consists of:
	// 1. A pointer to the underlying array
	// 2. Length (len)
	// 3. Capacity (cap)
	we_slice := []int{10, 20, 30}

	// We use reflect to look at the SliceHeader structure
	header := (*reflect.SliceHeader)(unsafe.Pointer(&we_slice))

	fmt.Printf("Slice Content: %v\n", we_slice)
	fmt.Printf("Pointer to Array: %X\n", header.Data)
	fmt.Printf("Length: %d\n", header.Len)
	fmt.Printf("Capacity: %d\n", header.Cap)

	// Pattern: Reslicing
	// When we reslice, we don't copy the data; we just move the pointer.
	we_resliced := we_slice[1:]
	new_header := (*reflect.SliceHeader)(unsafe.Pointer(&we_resliced))

	fmt.Println("\n>> AFTER RESLICING [1:]")
	fmt.Printf("New Pointer (Moved forward): %X\n", new_header.Data)
	fmt.Printf("New Length: %d\n", new_header.Len)
}
