package _billion_sort

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	tempFiles, err := splitAndSortChunks()
	if err != nil {
		log.Fatal("分块失败:", err)
	}

	err = mergeSortedFiles(tempFiles, "bigdata.sorted")
	if err != nil {
		log.Fatal("归并失败:", err)
	}

	log.Println("排序完成，输出文件：bigdata.sorted")
}

func TestGenerateBigDataFile(t *testing.T) {
	file, err := os.Create("bigdata")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10000; i++ {
		num := rand.Intn(1_000_000_000) // 生成0~10亿之间的随机数
		fmt.Fprintf(file, "%d\n", num)
	}
}
