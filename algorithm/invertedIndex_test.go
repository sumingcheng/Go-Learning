package algorithm

import (
	"fmt"
	"testing"
)

func TestBuildInvertedIndex(t *testing.T) {
	docs := []Doc{
		{Id: 1, Keywords: []string{"苹果", "橘子", "手机"}},
		{Id: 2, Keywords: []string{"苹果", "电脑", "手机"}},
		{Id: 3, Keywords: []string{"橘子", "电脑", "手机"}},
	}

	index := BuildInvertedIndex(docs)

	for k, v := range index {
		fmt.Println(k, v)
	}
}
