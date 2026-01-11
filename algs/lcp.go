package algs

import (
	"math"
)

func LCP() int {
	return 0
}

func LCA() {
}

type RMQPreproc struct {
	blen           int
	blocks         []int
	arr            []int
	sparseTable    [][]int
	blockMinTable  [][][]int
	blockTypeCache []int
}

// RMQ 0 <= i, j < N, arr[i] = arr[i-1] +(-) 1
// return index of the minimal element
func RMQ(arr []int, i, j int) int {
	return 0
}

func genRMQPreproc(arr []int) RMQPreproc {
	blen := int(math.Log2(float64(len(arr))) / 2)
	blocks := minBlockSplit(arr, blen)
	blockTypeCache := make([]int, 0, len(blocks))
	index := 0
	for i := range len(arr)-1 {
		if (i+1) % blen == 0 {
			blockTypeCache = append(blockTypeCache, index)
			index = 0	
			continue
		}
		if arr[i+1] > arr[i] {
			index <<= 1
			index |= 1
		} else {
			index <<= 1
		}
	}
	return RMQPreproc{
		blen:           blen,
		arr:            arr,
		blocks:         blocks,
		sparseTable:    sparseTable(blocks),
		blockMinTable:  blockMinTable(blen - 1),
		blockTypeCache: blockTypeCache,
	}
}

func blockMinTable(blen int) [][][]int {
	n := 1 << blen
	res := make([][][]int, 0, n)
	bt := make([]bool, blen)
	for range n {
		example := make([]int8, blen+1)
		for i := range blen {
			if bt[i] {
				example[i+1] = example[i] + 1
			} else {
				example[i+1] = example[i] - 1
			}
		}
		lr := make([][]int, 0, blen+1)
		for l := 0; l < blen+1; l++ {
			rrow := make([]int, blen+1)
			localMin := int8(127)
			minIndex := -1
			for r := l; r < blen+1; r++ {
				if example[r] < localMin {
					localMin = example[r]
					minIndex = r
				}
				rrow[r] = minIndex
			}
			lr = append(lr, rrow)
		}
		res = append(res, lr)
		btNext(bt)
	}
	return res
}

func btNext(bt []bool) {
	i := len(bt) - 1
	for i > 0 && bt[i] {
		bt[i] = false
		i--
	}
	bt[i] = true
}

// keep indexes of the min in arr
// n log n
func sparseTable(arr []int) [][]int {
	n := int(math.Log2(float64(len(arr)))) + 1
	st := make([][]int, len(arr))
	for i := range len(arr) {
		st[i] = make([]int, n)
		st[i][0] = i
	}
	pj2 := 1
	for j := 1; j < n; j++ {
		for i := range len(arr) {
			if i+pj2 < len(arr) && arr[st[i][j-1]] >= arr[st[i+pj2][j-1]] {
				st[i][j] = st[i+pj2][j-1]
			} else {
				st[i][j] = st[i][j-1]
			}
		}
		pj2 <<= 1
	}
	return st
}

// index of min in range
func minFromST(arr []int, st [][]int, l, r int) int {
	j := int(math.Log2(float64(r - l + 1)))
	pj2 := 1 << j
	if arr[st[l][j]] < arr[st[r-pj2+1][j]] {
		return st[l][j]
	}
	return st[r-pj2+1][j]
}

// every block has index of the minimum element in array in that block
// O(n)
func minBlockSplit(arr []int, blen int) []int {
	blocks := make([]int, len(arr)/blen+1)
	blockMin := math.MaxInt
	for i := range len(arr) {
		if arr[i] < blockMin {
			blockMin = arr[i]
			blocks[i/blen] = i
		}
		if (i+1)%blen == 0 {
			blockMin = math.MaxInt
		}
	}
	return blocks
}

func eulerSearch(node *STreeNode, nodes *[]*STreeNode, depths *[]int, depth int) {
	*nodes = append(*nodes, node)
	*depths = append(*depths, depth)
	for _, child := range node.childs {
		eulerSearch(child, nodes, depths, depth+1)
		*nodes = append(*nodes, node)
		*depths = append(*depths, depth)
	}
}
