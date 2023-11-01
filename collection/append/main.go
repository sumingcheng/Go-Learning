package main

func main() {
	//intSlice := make([]int, 3, 5)
	//
	//intSlice[0] = 1
	//intSlice[1] = 2
	//intSlice[2] = 3
	//
	//newIntSlice := append(intSlice, 4, 5, 6)
	//
	//// intSlice=[1 2 3], &intSlice=0xc00000e4b8,len=3, cap=5
	//fmt.Printf("intSlice=%v, &intSlice=%p,len=%d, cap=%d\n", intSlice, &intSlice[1], len(intSlice), cap(intSlice))
	//
	//// newIntSlice=[1 2 3 4 5 6],&newIntSlice=0xc0000121e8, len=6, cap=10
	//fmt.Printf("newIntSlice=%v,&newIntSlice=%p, len=%d, cap=%d\n", newIntSlice, &newIntSlice[1], len(newIntSlice), cap(newIntSlice))

	intSlice := []int{1, 2, 3, 4, 5}
	newSlice := append(intSlice[:2], intSlice[3:]...)
	println(newSlice)
}
