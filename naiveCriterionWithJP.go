package main

import (
	"fmt"
	"math"
	"time"
)

type reel struct {
	wheel []string // 存輪子
	n     int      // 單輪圖量
	m     int      // 單輪上畫面量
}

// 將指定好的各輪狀態 送去特定matrix  方便印出 存查(驗所有排列用)
func reelsToMatrix(re1 *reel, re2 *reel, re3 *reel, re4 *reel, re5 *reel, numOfReels int, num1 int, num2 int, num3 int, num4 int, num5 int) [][]string {
	// var showMatrix = [3][5]string
	// 初始化2D matrix
	showMatrix := make([][]string, re1.m) // row
	for index := range showMatrix {
		showMatrix[index] = make([]string, numOfReels)
	}

	tempReel1 := make([]string, re1.m)
	tempReel2 := make([]string, re2.m)
	tempReel3 := make([]string, re3.m)
	tempReel4 := make([]string, re4.m)
	tempReel5 := make([]string, re5.m)

	for i := 0; i < re1.m; i++ {
		tempReel1[i] = re1.wheel[(num1+i)%(re1.n)]
		tempReel2[i] = re2.wheel[(num2+i)%(re2.n)]
		tempReel3[i] = re3.wheel[(num3+i)%(re3.n)]
		tempReel4[i] = re4.wheel[(num4+i)%(re4.n)]
		tempReel5[i] = re5.wheel[(num5+i)%(re5.n)]
	}

	tempReels := [][]string{tempReel1, tempReel2, tempReel3, tempReel4, tempReel5}

	// 非方陣要調整
	for i := 0; i < re1.m; i++ {
		for j := 0; j < numOfReels; j++ {
			showMatrix[i][j] = tempReels[j][i]
		}
	}

	return showMatrix
}

func print2DMatrix(twoDimMatrix [][]string, totalRow int) {
	// 非方陣要調整
	// row i寫死成“上畫面數”
	for i := 0; i < totalRow; i++ {
		fmt.Println(twoDimMatrix[i][:])
	}
}

//  根據DRY 和上式之後必須抽出來重構
func printStatisticsMatrix(twoDimMatrix [][]int64, totalRow int) {
	// 非方陣要調整
	// row i寫死成“上畫面數”
	for i := 0; i < totalRow; i++ {
		fmt.Println(twoDimMatrix[i][:])
	}
}

// 連線判斷單元工具
// 為了方便 把Wild的判斷也一次在此寫入
func isEntryHasIt(Mat [][]string, row int, col int, symbol string) bool {
	if Mat[row][col] == symbol || Mat[row][col] == "W" {
		return true
	}

	return false
}

// 組合成scatter 連線判準上集  包含包含 你是要包 還是要含
func isColumnIncludeIt(Mat [][]string, col int, sca string) bool {
	// 先寫死為3x5 matrix的case
	for index := 0; index < 3; index++ {
		if Mat[index][col] == sca {
			return true
		}
	}
	return false
}

// 組合成scatter 連線判準下集  The 排擠
func isColumnExcludeIt(Mat [][]string, col int, sca string) bool {
	// 先寫死為3x5 matrix的case
	for index := 0; index < 3; index++ {
		if Mat[index][col] == sca {
			return false
		}
	}
	return true
}

// size in [3,4,5]  "W"預設為Wild
func verifyRow(win [][]string, row int, size int, picture string) bool {

	switch size {
	case 3:
		straightThree := (win[row][0] == "W" || win[row][0] == picture) && (win[row][1] == "W" || win[row][1] == picture) && (win[row][2] == "W" || win[row][2] == picture) && (win[row][3] != "W" && win[row][3] != picture)
		straight3W := win[row][0] == "W" && win[row][1] == "W" && win[row][2] == "W"
		if straightThree && !straight3W {
			return true
		}

		return false

	case 4:
		straightFour := (win[row][0] == "W" || win[row][0] == picture) && (win[row][1] == "W" || win[row][1] == picture) && (win[row][2] == "W" || win[row][2] == picture) && (win[row][3] == "W" || win[row][3] == picture) && (win[row][4] != "W" && win[row][4] != picture)
		straight4W := win[row][0] == "W" && win[row][1] == "W" && win[row][2] == "W" && win[row][3] == "W"
		if straightFour && !straight4W {
			return true
		}

		return false

	default:
		straightFive := (win[row][0] == "W" || win[row][0] == picture) && (win[row][1] == "W" || win[row][1] == picture) && (win[row][2] == "W" || win[row][2] == picture) && (win[row][3] == "W" || win[row][3] == picture) && (win[row][4] == "W" || win[row][4] == picture)
		straight5W := win[row][0] == "W" && win[row][1] == "W" && win[row][2] == "W" && win[row][3] == "W" && win[row][4] == "W"
		if straightFive && !straight5W {
			return true
		}

		return false
	}

}

func countOf3Scatter(win [][]string) bool {
	s3 := (isColumnIncludeIt(win, 0, "S")) && (isColumnIncludeIt(win, 1, "S")) && (isColumnIncludeIt(win, 2, "S")) && (isColumnExcludeIt(win, 3, "S"))
	if s3 {
		return true
	}
	return false
}

func countOf4Scatter(win [][]string) bool {
	s4 := (isColumnIncludeIt(win, 0, "S")) && (isColumnIncludeIt(win, 1, "S")) && (isColumnIncludeIt(win, 2, "S")) && (isColumnIncludeIt(win, 3, "S")) && (isColumnExcludeIt(win, 4, "S"))
	if s4 {
		return true
	}
	return false
}

func countOf5Scatter(win [][]string) bool {
	s5 := (isColumnIncludeIt(win, 0, "S")) && (isColumnIncludeIt(win, 1, "S")) && (isColumnIncludeIt(win, 2, "S")) && (isColumnIncludeIt(win, 3, "S")) && (isColumnIncludeIt(win, 4, "S"))
	if s5 {
		return true
	}
	return false
}

func countCase4(win [][]string, size int, picture string) bool {
	switch size {
	case 3:
		lineThree := isEntryHasIt(win, 1, 0, picture) && isEntryHasIt(win, 0, 1, picture) && isEntryHasIt(win, 0, 2, picture) && !isEntryHasIt(win, 0, 3, picture)
		lineThreeW := isEntryHasIt(win, 1, 0, "W") && isEntryHasIt(win, 0, 1, "W") && isEntryHasIt(win, 0, 2, "W")
		if lineThree && !lineThreeW {
			return true
		}
		return false

	case 4:
		lineFour := isEntryHasIt(win, 1, 0, picture) && isEntryHasIt(win, 0, 1, picture) && isEntryHasIt(win, 0, 2, picture) && isEntryHasIt(win, 0, 3, picture) && !isEntryHasIt(win, 1, 4, picture)
		lineFourW := isEntryHasIt(win, 1, 0, "W") && isEntryHasIt(win, 0, 1, "W") && isEntryHasIt(win, 0, 2, "W") && isEntryHasIt(win, 0, 3, "W")
		if lineFour && !lineFourW {
			return true
		}
		return false

	default:
		lineFive := isEntryHasIt(win, 1, 0, picture) && isEntryHasIt(win, 0, 1, picture) && isEntryHasIt(win, 0, 2, picture) && isEntryHasIt(win, 0, 3, picture) && isEntryHasIt(win, 1, 4, picture)
		lineFiveW := isEntryHasIt(win, 1, 0, "W") && isEntryHasIt(win, 0, 1, "W") && isEntryHasIt(win, 0, 2, "W") && isEntryHasIt(win, 0, 3, "W") && isEntryHasIt(win, 1, 4, "W")
		if lineFive && !lineFiveW {
			return true
		}
		return false
	}

}

func countCase5(win [][]string, size int, picture string) bool {
	switch size {
	case 3:
		lineThree := isEntryHasIt(win, 1, 0, picture) && isEntryHasIt(win, 2, 1, picture) && isEntryHasIt(win, 2, 2, picture) && !isEntryHasIt(win, 2, 3, picture)
		lineThreeW := isEntryHasIt(win, 1, 0, "W") && isEntryHasIt(win, 2, 1, "W") && isEntryHasIt(win, 2, 2, "W")
		if lineThree && !lineThreeW {
			return true
		}
		return false

	case 4:
		lineFour := isEntryHasIt(win, 1, 0, picture) && isEntryHasIt(win, 2, 1, picture) && isEntryHasIt(win, 2, 2, picture) && isEntryHasIt(win, 2, 3, picture) && !isEntryHasIt(win, 1, 4, picture)
		lineFourW := isEntryHasIt(win, 1, 0, "W") && isEntryHasIt(win, 2, 1, "W") && isEntryHasIt(win, 2, 2, "W") && isEntryHasIt(win, 2, 3, "W")
		if lineFour && !lineFourW {
			return true
		}
		return false

	default:
		lineFive := isEntryHasIt(win, 1, 0, picture) && isEntryHasIt(win, 2, 1, picture) && isEntryHasIt(win, 2, 2, picture) && isEntryHasIt(win, 2, 3, picture) && isEntryHasIt(win, 1, 4, picture)
		lineFiveW := isEntryHasIt(win, 1, 0, "W") && isEntryHasIt(win, 2, 1, "W") && isEntryHasIt(win, 2, 2, "W") && isEntryHasIt(win, 2, 3, "W") && isEntryHasIt(win, 1, 4, "W")
		if lineFive && !lineFiveW {
			return true
		}
		return false
	}
}

func countCase6(win [][]string, size int, picture string) bool {
	switch size {
	case 3:
		lineThree := isEntryHasIt(win, 0, 0, picture) && isEntryHasIt(win, 0, 1, picture) && isEntryHasIt(win, 1, 2, picture) && !isEntryHasIt(win, 0, 3, picture)
		lineThreeW := isEntryHasIt(win, 0, 0, "W") && isEntryHasIt(win, 0, 1, "W") && isEntryHasIt(win, 1, 2, "W")
		if lineThree && !lineThreeW {
			return true
		}
		return false

	case 4:
		lineFour := isEntryHasIt(win, 0, 0, picture) && isEntryHasIt(win, 0, 1, picture) && isEntryHasIt(win, 1, 2, picture) && isEntryHasIt(win, 0, 3, picture) && !isEntryHasIt(win, 0, 4, picture)
		lineFourW := isEntryHasIt(win, 0, 0, "W") && isEntryHasIt(win, 0, 1, "W") && isEntryHasIt(win, 1, 2, "W") && isEntryHasIt(win, 0, 3, "W")
		if lineFour && !lineFourW {
			return true
		}
		return false

	default:
		lineFive := isEntryHasIt(win, 0, 0, picture) && isEntryHasIt(win, 0, 1, picture) && isEntryHasIt(win, 1, 2, picture) && isEntryHasIt(win, 0, 3, picture) && isEntryHasIt(win, 0, 4, picture)
		lineFiveW := isEntryHasIt(win, 0, 0, "W") && isEntryHasIt(win, 0, 1, "W") && isEntryHasIt(win, 1, 2, "W") && isEntryHasIt(win, 0, 3, "W") && isEntryHasIt(win, 0, 4, "W")
		if lineFive && !lineFiveW {
			return true
		}
		return false
	}
}

func countCase7(win [][]string, size int, picture string) bool {
	switch size {
	case 3:
		lineThree := isEntryHasIt(win, 2, 0, picture) && isEntryHasIt(win, 2, 1, picture) && isEntryHasIt(win, 1, 2, picture) && !isEntryHasIt(win, 2, 3, picture)
		lineThreeW := isEntryHasIt(win, 2, 0, "W") && isEntryHasIt(win, 2, 1, "W") && isEntryHasIt(win, 1, 2, "W")
		if lineThree && !lineThreeW {
			return true
		}
		return false

	case 4:
		lineFour := isEntryHasIt(win, 2, 0, picture) && isEntryHasIt(win, 2, 1, picture) && isEntryHasIt(win, 1, 2, picture) && isEntryHasIt(win, 2, 3, picture) && !isEntryHasIt(win, 2, 4, picture)
		lineFourW := isEntryHasIt(win, 2, 0, "W") && isEntryHasIt(win, 2, 1, "W") && isEntryHasIt(win, 1, 2, "W") && isEntryHasIt(win, 2, 3, "W")
		if lineFour && !lineFourW {
			return true
		}
		return false

	default:
		lineFive := isEntryHasIt(win, 2, 0, picture) && isEntryHasIt(win, 2, 1, picture) && isEntryHasIt(win, 1, 2, picture) && isEntryHasIt(win, 2, 3, picture) && isEntryHasIt(win, 2, 4, picture)
		lineFiveW := isEntryHasIt(win, 2, 0, "W") && isEntryHasIt(win, 2, 1, "W") && isEntryHasIt(win, 1, 2, "W") && isEntryHasIt(win, 2, 3, "W") && isEntryHasIt(win, 2, 4, "W")
		if lineFive && !lineFiveW {
			return true
		}
		return false
	}
}

func countCase8(win [][]string, size int, picture string) bool {
	switch size {
	case 3:
		lineThree := isEntryHasIt(win, 0, 0, picture) && isEntryHasIt(win, 1, 1, picture) && isEntryHasIt(win, 0, 2, picture) && !isEntryHasIt(win, 1, 3, picture)
		lineThreeW := isEntryHasIt(win, 0, 0, "W") && isEntryHasIt(win, 1, 1, "W") && isEntryHasIt(win, 0, 2, "W")
		if lineThree && !lineThreeW {
			return true
		}
		return false

	case 4:
		lineFour := isEntryHasIt(win, 0, 0, picture) && isEntryHasIt(win, 1, 1, picture) && isEntryHasIt(win, 0, 2, picture) && isEntryHasIt(win, 1, 3, picture) && !isEntryHasIt(win, 0, 4, picture)
		lineFourW := isEntryHasIt(win, 0, 0, "W") && isEntryHasIt(win, 1, 1, "W") && isEntryHasIt(win, 0, 2, "W") && isEntryHasIt(win, 1, 3, "W")
		if lineFour && !lineFourW {
			return true
		}
		return false

	default:
		lineFive := isEntryHasIt(win, 0, 0, picture) && isEntryHasIt(win, 1, 1, picture) && isEntryHasIt(win, 0, 2, picture) && isEntryHasIt(win, 1, 3, picture) && isEntryHasIt(win, 0, 4, picture)
		lineFiveW := isEntryHasIt(win, 0, 0, "W") && isEntryHasIt(win, 1, 1, "W") && isEntryHasIt(win, 0, 2, "W") && isEntryHasIt(win, 1, 3, "W") && isEntryHasIt(win, 0, 4, "W")
		if lineFive && !lineFiveW {
			return true
		}
		return false
	}
}

func countCase9(win [][]string, size int, picture string) bool {
	switch size {
	case 3:
		lineThree := isEntryHasIt(win, 2, 0, picture) && isEntryHasIt(win, 1, 1, picture) && isEntryHasIt(win, 2, 2, picture) && !isEntryHasIt(win, 1, 3, picture)
		lineThreeW := isEntryHasIt(win, 2, 0, "W") && isEntryHasIt(win, 1, 1, "W") && isEntryHasIt(win, 2, 2, "W")
		if lineThree && !lineThreeW {
			return true
		}
		return false

	case 4:
		lineFour := isEntryHasIt(win, 2, 0, picture) && isEntryHasIt(win, 1, 1, picture) && isEntryHasIt(win, 2, 2, picture) && isEntryHasIt(win, 1, 3, picture) && !isEntryHasIt(win, 2, 4, picture)
		lineFourW := isEntryHasIt(win, 2, 0, "W") && isEntryHasIt(win, 1, 1, "W") && isEntryHasIt(win, 2, 2, "W") && isEntryHasIt(win, 1, 3, "W")
		if lineFour && !lineFourW {
			return true
		}
		return false

	default:
		lineFive := isEntryHasIt(win, 2, 0, picture) && isEntryHasIt(win, 1, 1, picture) && isEntryHasIt(win, 2, 2, picture) && isEntryHasIt(win, 1, 3, picture) && isEntryHasIt(win, 2, 4, picture)
		lineFiveW := isEntryHasIt(win, 2, 0, "W") && isEntryHasIt(win, 1, 1, "W") && isEntryHasIt(win, 2, 2, "W") && isEntryHasIt(win, 1, 3, "W") && isEntryHasIt(win, 2, 4, "W")
		if lineFive && !lineFiveW {
			return true
		}
		return false
	}
}

// 每個分量去判定fix size and symbol之下 不同case是否發生  這是上集  下集  瓜哥：請收看下面一位 func
// 想print中獎matrix的話 用iterate 對true的分量 print matrix
func judgeSizeAndSymbol(win [][]string, size int, picture string) []bool {
	judge := make([]bool, 9) // 現在先算9種case 暫寫死為9  之後擴充要記得改
	judge = []bool{verifyRow(win, 0, size, picture), verifyRow(win, 1, size, picture), verifyRow(win, 2, size, picture), countCase4(win, size, picture), countCase5(win, size, picture), countCase6(win, size, picture), countCase7(win, size, picture), countCase8(win, size, picture), countCase9(win, size, picture)}

	return judge
}

// 這是下集  有發生的case才計數
func sumOfBoolArray(bla []bool) int {
	sum := 0
	for _, v := range bla {
		if v {
			sum++
		}
	}
	return sum
}

// 先對每個symbol暴搜 each case的3 4 5連線  再依此暴搜集其他symbol
func criterion(re1 *reel, re2 *reel, re3 *reel, re4 *reel, re5 *reel, numOfReels int) []int64 {
	var countOfA3, countOfA4, countOfA5, countOfB3, countOfB4, countOfB5, countOfC3, countOfC4, countOfC5, countOfD3, countOfD4, countOfD5,
		countOfE3, countOfE4, countOfE5, countOfF3, countOfF4, countOfF5, countOfG3, countOfG4, countOfG5, countOfH3, countOfH4, countOfH5,
		countOfS3, countOfS4, countOfS5 int64 // 方法數可能爆炸大

	countOfA3, countOfA4, countOfA5, countOfB3, countOfB4, countOfB5, countOfC3, countOfC4, countOfC5, countOfD3, countOfD4, countOfD5,
		countOfE3, countOfE4, countOfE5, countOfF3, countOfF4, countOfF5, countOfG3, countOfG4, countOfG5, countOfH3, countOfH4, countOfH5,
		countOfS3, countOfS4, countOfS5 = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0

	for i := 0; i < re1.n; i++ {
		for j := 0; j < re2.n; j++ {
			for k := 0; k < re3.n; k++ {
				for l := 0; l < re4.n; l++ {
					for m := 0; m < re5.n; m++ {
						win := reelsToMatrix(re1, re2, re3, re4, re5, numOfReels, i, j, k, l, m) // win[3][5]

						countOfA5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "A")))
						countOfA4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "A")))
						countOfA3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "A")))

						countOfB5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "B")))
						countOfB4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "B")))
						countOfB3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "B")))

						countOfC5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "C")))
						countOfC4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "C")))
						countOfC3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "C")))

						countOfD5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "D")))
						countOfD4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "D")))
						countOfD3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "D")))

						countOfE5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "E")))
						countOfE4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "E")))
						countOfE3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "E")))

						countOfF5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "F")))
						countOfF4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "F")))
						countOfF3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "F")))

						countOfG5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "G")))
						countOfG4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "G")))
						countOfG3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "G")))

						countOfH5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "H")))
						countOfH4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "H")))
						countOfH3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "H")))

						if countOf5Scatter(win) {
							countOfS5++
							//print2DMatrix(win, 3)
							//fmt.Println("===========================================")
						}

						// S S S S #S
						if countOf4Scatter(win) {
							countOfS4++
							//print2DMatrix(win,3)
							//fmt.Println("===========================================")
						}

						// 暫定S S S #S any
						if countOf3Scatter(win) {
							countOfS3++
							//print2DMatrix(win,3)
							//fmt.Println("===========================================")
						}

					}
				}
			}
		}
	}

	sols := []int64{countOfA3, countOfA4, countOfA5, countOfB3, countOfB4, countOfB5, countOfC3, countOfC4, countOfC5, countOfD3, countOfD4, countOfD5, countOfE3, countOfE4, countOfE5, countOfF3, countOfF4, countOfF5, countOfG3, countOfG4, countOfG5, countOfH3, countOfH4, countOfH5, countOfS3, countOfS4, countOfS5}

	return sols
}

func runPermutationsOfReels(re1 *reel, re2 *reel, re3 *reel, re4 *reel, re5 *reel) int64 {
	countOfPermutations := int64(re1.n) * int64(re2.n) * int64(re3.n) * int64(re4.n) * int64(re5.n)

	/*var countOfPermutations int64 // 方法數可能爆炸大
	countOfPermutations = 0
	// 非方陣要調整
	for i := 0; i < re1.n; i++ {
		for j := 0; j < re2.n; j++ {
			for k := 0; k < re3.n; k++ {
				for l := 0; l < re4.n; l++ {
					for m := 0; m < re5.n; m++ {

						// print2DMatrix(reelsToMatrix(re1, re2, re3, re4, re5, numOfReels, i, j, k, l, m),3) // reels參數可寫死為5
						// fmt.Println("=============================================")
						countOfPermutations++
					}
				}
			}
		}
	}

	// fmt.Println("排列總數:", countOfPermutations)*/
	return countOfPermutations
}

// 原始資料統計
// numOfPicture含Scatter 不含Wild
func statisticOfHits(re1 *reel, re2 *reel, re3 *reel, re4 *reel, re5 *reel, numOfReels int, numOfPicture int) [][]int64 {

	ArrayOfHits := make([][]int64, numOfPicture) // row
	for index := range ArrayOfHits {
		ArrayOfHits[index] = make([]int64, 3) // for3 4 5連線
	}

	// A_i, j = B_3i+j   數學好方便啊  棒棒der~~~  肥宅救星
	cri := criterion(re1, re2, re3, re4, re5, numOfReels)
	for i := 0; i < numOfPicture; i++ {
		for j := 0; j < 3; j++ {
			ArrayOfHits[i][j] = cri[3*i+j]
		}
	}

	// printStatisticsMatrix(ArrayOfHits, numOfPicture)
	fmt.Println("A       3 4 5:", ArrayOfHits[0][:])
	fmt.Println("B       3 4 5:", ArrayOfHits[1][:])
	fmt.Println("C       3 4 5:", ArrayOfHits[2][:])
	fmt.Println("D       3 4 5:", ArrayOfHits[3][:])
	fmt.Println("E       3 4 5:", ArrayOfHits[4][:])
	fmt.Println("F       3 4 5:", ArrayOfHits[5][:])
	fmt.Println("G       3 4 5:", ArrayOfHits[6][:])
	fmt.Println("H       3 4 5:", ArrayOfHits[7][:])
	fmt.Println("Scatter 3 4 5:", ArrayOfHits[8][:])
	return ArrayOfHits
}

func statisticOfProbability(re1 *reel, re2 *reel, re3 *reel, re4 *reel, re5 *reel, numOfReels int, numOfPicture int) [][]float64 {

	ArrayOfProbability := make([][]float64, numOfPicture) // row
	for index := range ArrayOfProbability {
		ArrayOfProbability[index] = make([]float64, 3) // for3 4 5連線
	}

	stat := statisticOfHits(re1, re2, re3, re4, re5, numOfReels, numOfPicture)
	allPermutation := float64(runPermutationsOfReels(re1, re2, re3, re4, re5))

	for i := 0; i < numOfPicture; i++ {
		for j := 0; j < 3; j++ {
			ArrayOfProbability[i][j] = float64(stat[i][j]) / allPermutation
		}
	}

	fmt.Println("prob of A       3 4 5:", ArrayOfProbability[0][:])
	fmt.Println("prob of B       3 4 5:", ArrayOfProbability[1][:])
	fmt.Println("prob of C       3 4 5:", ArrayOfProbability[2][:])
	fmt.Println("prob of D       3 4 5:", ArrayOfProbability[3][:])
	fmt.Println("prob of E       3 4 5:", ArrayOfProbability[4][:])
	fmt.Println("prob of F       3 4 5:", ArrayOfProbability[5][:])
	fmt.Println("prob of G       3 4 5:", ArrayOfProbability[6][:])
	fmt.Println("prob of H       3 4 5:", ArrayOfProbability[7][:])
	fmt.Println("prob of Scatter 3 4 5:", ArrayOfProbability[8][:])
	return ArrayOfProbability

}

// 暫時用手動在function內寫死的方式
// 之後再補array沒填滿或非法值的宣告
// caseOfGame  0:base game    1:free game with each pay*10  2:free game with each pay*15  3:free game with each pay*20
// 之後應該要補不合理input的報錯機制
func payTable(re1 *reel, re2 *reel, re3 *reel, re4 *reel, re5 *reel, numOfPicture int, caseOfGame int) [][]float64 {

	ArrayOfPayTable := make([][]float64, numOfPicture) // row
	for index := range ArrayOfPayTable {
		ArrayOfPayTable[index] = make([]float64, 3) // for3 4 5連線
	}

	// 9x3 matrix   pay nothing for all scatter
	// 改動調整的重點處之一
	ArrayOfPayTable = [][]float64{[]float64{2, 6, 30}, []float64{2, 4, 9}, []float64{2, 4, 9},
		[]float64{2, 4, 9}, []float64{2, 4, 9}, []float64{2, 4, 9},
		[]float64{2, 6, 15}, []float64{2, 4, 15}, []float64{0, 0, 0}}

	switch caseOfGame {
	case 0:
		ArrayOfPayTable = ArrayOfPayTable
	case 1:
		for i := 0; i < numOfPicture; i++ {
			for j := 0; j < 3; j++ {
				ArrayOfPayTable[i][j] = ArrayOfPayTable[i][j] * 10
			}
		}
	case 2:
		for i := 0; i < numOfPicture; i++ {
			for j := 0; j < 3; j++ {
				ArrayOfPayTable[i][j] = ArrayOfPayTable[i][j] * 15
			}
		}
	case 3:
		for i := 0; i < numOfPicture; i++ {
			for j := 0; j < 3; j++ {
				ArrayOfPayTable[i][j] = ArrayOfPayTable[i][j] * 20
			}
		}

	}

	fmt.Println("pay of A       3 4 5:", ArrayOfPayTable[0][:])
	fmt.Println("pay of B       3 4 5:", ArrayOfPayTable[1][:])
	fmt.Println("pay of C       3 4 5:", ArrayOfPayTable[2][:])
	fmt.Println("pay of D       3 4 5:", ArrayOfPayTable[3][:])
	fmt.Println("pay of E       3 4 5:", ArrayOfPayTable[4][:])
	fmt.Println("pay of F       3 4 5:", ArrayOfPayTable[5][:])
	fmt.Println("pay of G       3 4 5:", ArrayOfPayTable[6][:])
	fmt.Println("pay of H       3 4 5:", ArrayOfPayTable[7][:])
	fmt.Println("pay of Scatter 3 4 5:", ArrayOfPayTable[8][:])

	return ArrayOfPayTable

}

// let x be random variable
// exp(x)=sigma_all i (xi * pi)  ; var(x)=sigma_all i [(x-exp(x))^2 * pi=exp(x^2)-exp^2(x) ; std = var^(1/2)]
// 算free game平均次數時 各Scatter觸發的free game次數是事先知道  可以寫死
// 為了之後算含free game的統計數據方便  func可改return type of array
// 承上 free game多return 調整係數
// 目前設計上有綁特定pay table  調整caseOfGame的參數
func baseEV(re1 *reel, re2 *reel, re3 *reel, re4 *reel, re5 *reel, numOfReels int, numOfPicture int) []float64 {

	prob := statisticOfProbability(re1, re2, re3, re4, re5, numOfReels, numOfPicture)
	pay := payTable(re1, re2, re3, re4, re5, numOfPicture, 0) // caseOfGame = 0
	expectedValue := float64(0)
	sqrtnorm := float64(0)

	// 注意規則是分9個case押注  我們假想單次總下注為each case x 9
	for i := 0; i < numOfPicture; i++ {
		for j := 0; j < 3; j++ {
			expectedValue += prob[i][j] * pay[i][j]
			sqrtnorm += pay[i][j] * pay[i][j] * prob[i][j]
		}
	}

	variance := sqrtnorm - (expectedValue * expectedValue)
	std := math.Sqrt(variance)

	// Scatter 3連4連5連
	// triggerOfFG := []float64{10, 15, 25}
	triggerOfFG := []float64{5, 10, 20}
	/* // 原理:sum of infinite series
	triggerFactor := 1 / (1 - (10*prob[8][0] + 15*prob[8][1] + 25*prob[8][2]))

	// weighted of free game
	adjustOfFG := make([]float64, 3)
	for i := range adjustOfFG {
		adjustOfFG[i] = triggerOfFG[i] * triggerFactor
	}

	// free game平均次
	// 也可用a=sigma(pi*si)  b=triggerFactor  直接a*b得出結果
	expOfFG := adjustOfFG[0]*prob[4][0] + adjustOfFG[1]*prob[8][1] + adjustOfFG[2]*prob[8][2] */

	// weitedFG := 10*prob[8][0] + 15*prob[8][1] + 25*prob[8][2]
	weitedFG := float64(0)
	for i, v := range triggerOfFG {
		weitedFG += v * prob[8][i]
	}

	triggerFactor := 1 / (1 - weitedFG)
	expOfFG := weitedFG * triggerFactor

	// 算總中獎機率
	chanceOFPrize := float64(0)
	for i := 0; i < numOfPicture; i++ {
		for j := 0; j < 3; j++ {
			chanceOFPrize += prob[i][j]
		}
	}

	fmt.Println("base game EV值:", expectedValue, " base game Var:", variance, " base game std:", std)
	fmt.Println("\n")
	fmt.Println("according to normal distribution, the estimation of theoretical values of base game as the following:\n")
	fmt.Println("base game RTP is:", expectedValue*100, "%")
	fmt.Println("我現在宣布 天下第一 base game波動程度理論值估計大會開始")
	fmt.Println("玩100次", "90%信賴水準 base game期望值落在：", "(", expectedValue-(1.64*std/math.Sqrt(100)), ",", expectedValue+(1.64*std/math.Sqrt(100)), ")")
	fmt.Println("玩100次", "95%信賴水準 base game期望值落在：", "(", expectedValue-(1.96*std/math.Sqrt(100)), ",", expectedValue+(1.96*std/math.Sqrt(100)), ")")
	fmt.Println("玩100次", "99%信賴水準 base game期望值落在：", "(", expectedValue-(2.58*std/math.Sqrt(100)), ",", expectedValue+(2.58*std/math.Sqrt(100)), ")")
	fmt.Println("我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我")
	fmt.Println("玩10000次", "90%信賴水準 base game期望值落在：", "(", expectedValue-(1.64*std/math.Sqrt(10000)), ",", expectedValue+(1.64*std/math.Sqrt(10000)), ")")
	fmt.Println("玩10000次", "95%信賴水準 base game期望值落在：", "(", expectedValue-(1.96*std/math.Sqrt(10000)), ",", expectedValue+(1.96*std/math.Sqrt(10000)), ")")
	fmt.Println("玩10000次", "99%信賴水準 base game期望值落在：", "(", expectedValue-(2.58*std/math.Sqrt(10000)), ",", expectedValue+(2.58*std/math.Sqrt(10000)), ")")
	fmt.Println("我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我")
	fmt.Println("玩1000000次", "90%信賴水準 base game期望值落在：", "(", expectedValue-(1.64*std/math.Sqrt(1000000)), ",", expectedValue+(1.64*std/math.Sqrt(1000000)), ")")
	fmt.Println("玩1000000次", "95%信賴水準 base game期望值落在：", "(", expectedValue-(1.96*std/math.Sqrt(1000000)), ",", expectedValue+(1.96*std/math.Sqrt(1000000)), ")")
	fmt.Println("玩1000000次", "99%信賴水準 base game期望值落在：", "(", expectedValue-(2.58*std/math.Sqrt(1000000)), ",", expectedValue+(2.58*std/math.Sqrt(1000000)), ")")
	fmt.Println("\n")
	fmt.Println("3連線 4連線 5連線 可玩的free game:", triggerOfFG)
	fmt.Println("trigger因子", triggerFactor)
	// fmt.Println("weighted of free game:", adjustOfFG)
	fmt.Println("可玩free game平均次數:", expOfFG)
	fmt.Println("base game總中獎機率:", chanceOFPrize)
	fmt.Println("===========================================")

	// return expOfFG 方便算總RTP
	arrayOfEv := []float64{expectedValue, variance, std, expOfFG}
	return arrayOfEv

}

// 前五輪為base game的reel 參數  後五輪為free game的reel 參數
// 以之後會獨立設計base game  free game的情境來寫此func
// 暫時設計為所有game 進各free game的機率有不同權重
// 1:free game with each pay*10  2:free game with each pay*15  3:free game with each pay*20
func baseAndfreeEV(re1 *reel, re2 *reel, re3 *reel, re4 *reel, re5 *reel, refr1 *reel, refr2 *reel, refr3 *reel, refr4 *reel, refr5 *reel, numOfReels int, numOfPicture int) {
	// 弄base game部分
	baseArray := baseEV(re1, re2, re3, re4, re5, numOfReels, numOfPicture)

	// 先分開重算free game的基本數據
	prob := statisticOfProbability(refr1, refr2, refr3, refr4, refr5, numOfReels, numOfPicture)

	payOfFG1 := payTable(re1, re2, re3, re4, re5, numOfPicture, 1)
	payOfFG2 := payTable(re1, re2, re3, re4, re5, numOfPicture, 2)
	payOfFG3 := payTable(re1, re2, re3, re4, re5, numOfPicture, 3)

	expectedValueOfFG1 := float64(0)
	expectedValueOfFG2 := float64(0)
	expectedValueOfFG3 := float64(0)

	sqrtnormOfFG1 := float64(0)
	sqrtnormOfFG2 := float64(0)
	sqrtnormOfFG3 := float64(0)

	// 注意規則是分9個case押注  我們假想單次總下注為each case x 9
	// 最新:已調回 忽略上面
	for i := 0; i < numOfPicture; i++ {
		for j := 0; j < 3; j++ {
			expectedValueOfFG1 += prob[i][j] * payOfFG1[i][j]
			sqrtnormOfFG1 += payOfFG1[i][j] * payOfFG1[i][j] * prob[i][j]

			expectedValueOfFG2 += prob[i][j] * payOfFG2[i][j]
			sqrtnormOfFG2 += payOfFG2[i][j] * payOfFG2[i][j] * prob[i][j]

			expectedValueOfFG3 += prob[i][j] * payOfFG3[i][j]
			sqrtnormOfFG3 += payOfFG3[i][j] * payOfFG3[i][j] * prob[i][j]
		}
	}

	varianceOfFG1 := sqrtnormOfFG1 - (expectedValueOfFG1 * expectedValueOfFG1)
	// stdOfFG1 := math.Sqrt(varianceOfFG1)

	varianceOfFG2 := sqrtnormOfFG2 - (expectedValueOfFG2 * expectedValueOfFG2)
	// stdOfFG2 := math.Sqrt(varianceOfFG2)

	varianceOfFG3 := sqrtnormOfFG3 - (expectedValueOfFG3 * expectedValueOfFG3)
	// stdOfFG3 := math.Sqrt(varianceOfFG3)

	// 算總中獎機率
	chanceOFPrize := float64(0)
	for i := 0; i < numOfPicture; i++ {
		for j := 0; j < 3; j++ {
			chanceOFPrize += prob[i][j]
		}
	}

	// 在baseEV裡  我們精油  經由啦  幹  by 收斂的無窮級數和之計算 得出free game平均次數
	// 現在進一步設計  三種free game的隨機權重機制  權重由rand.Perm of 對稱群S_3給出
	// 輸出identity  則進20倍free game  3種奇置換進10倍free game 最後2種偶置換進15倍free game
	// var自行手動調整權重
	totalExpected := baseArray[0] + baseArray[3]*(expectedValueOfFG1*3+expectedValueOfFG2*2+expectedValueOfFG3)/6
	//totalExpected := baseArray[0] + baseArray[3]*(expectedValueOfFG1+expectedValueOfFG2+expectedValueOfFG3)/3
	totalVariance := baseArray[1] + baseArray[3]*baseArray[3]*(varianceOfFG1*9+varianceOfFG2*4+varianceOfFG3)/36
	//totalVariance := baseArray[1] + baseArray[3]*baseArray[3]*(varianceOfFG1+varianceOfFG2+varianceOfFG3)/9
	totalStd := math.Sqrt(totalVariance)

	fmt.Println("3種free game的EV (10 15 20):", expectedValueOfFG1, expectedValueOfFG2, expectedValueOfFG3)
	fmt.Println("3種free game的var (10 15 20):", varianceOfFG1, varianceOfFG2, varianceOfFG3)
	fmt.Println(baseArray[0], baseArray[3])
	fmt.Println("總EV值:", totalExpected, "Var:", totalVariance, "std:", totalStd)
	fmt.Println("\n")
	fmt.Println("according to normal distribution, the estimation of theoretical values as the following:\n")
	fmt.Println("RTP is:", totalExpected*100, "%")
	fmt.Println("我現在宣布 天下第一 波動程度理論值估計大會開始")
	fmt.Println("玩100次", "90%信賴水準 期望值落在：", "(", totalExpected-(1.64*totalStd/math.Sqrt(100)), ",", totalExpected+(1.64*totalStd/math.Sqrt(100)), ")")
	fmt.Println("玩100次", "95%信賴水準 期望值落在：", "(", totalExpected-(1.96*totalStd/math.Sqrt(100)), ",", totalExpected+(1.96*totalStd/math.Sqrt(100)), ")")
	fmt.Println("玩100次", "99%信賴水準 期望值落在：", "(", totalExpected-(2.58*totalStd/math.Sqrt(100)), ",", totalExpected+(2.58*totalStd/math.Sqrt(100)), ")")

	fmt.Println("我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我")
	fmt.Println("玩10000次", "90%信賴水準 期望值落在：", "(", totalExpected-(1.64*totalStd/math.Sqrt(10000)), ",", totalExpected+(1.64*totalStd/math.Sqrt(10000)), ")")
	fmt.Println("玩10000次", "95%信賴水準 期望值落在：", "(", totalExpected-(1.96*totalStd/math.Sqrt(10000)), ",", totalExpected+(1.96*totalStd/math.Sqrt(10000)), ")")
	fmt.Println("玩10000次", "99%信賴水準 期望值落在：", "(", totalExpected-(2.58*totalStd/math.Sqrt(10000)), ",", totalExpected+(2.58*totalStd/math.Sqrt(10000)), ")")
	fmt.Println("我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我 我是分隔線 分隔線是我")
	fmt.Println("玩1000000次", "90%信賴水準 期望值落在：", "(", totalExpected-(1.64*totalStd/math.Sqrt(1000000)), ",", totalExpected+(1.64*totalStd/math.Sqrt(1000000)), ")")
	fmt.Println("玩1000000次", "95%信賴水準 期望值落在：", "(", totalExpected-(1.96*totalStd/math.Sqrt(1000000)), ",", totalExpected+(1.96*totalStd/math.Sqrt(1000000)), ")")
	fmt.Println("玩1000000次", "99%信賴水準 期望值落在：", "(", totalExpected-(2.58*totalStd/math.Sqrt(1000000)), ",", totalExpected+(2.58*totalStd/math.Sqrt(1000000)), ")")
	fmt.Println("\n")
	fmt.Println("free game總中獎機率:", chanceOFPrize)

}

// 後續懶得轉型 所以先一次弄成float64
func factorial(k float64) float64 {
	if k == 0 {
		return 1
	}
	return k * factorial(k-1)

}

// A discrete random variable X  is said to have a Poisson distribution with parameter lamda > 0,
// if, for k = 0, 1, 2, ..., the probability mass function of X  is given by
// Prob(X=k) = (lamda^k)* (e^(-lamda)) / k!
func poissonDistribution(k float64, lamda float64) float64 {
	p := (math.Pow(lamda, k)) * (math.Exp(-1 * lamda)) / factorial(k)
	return p
}

// 故意寫成每輪size不一樣大時  同樣可apply
// 暫定估base game就好  進free game機率本身就夠低了 先略去不計
func estimationOfJP(re1 *reel, re2 *reel, re3 *reel, re4 *reel, re5 *reel, ntimes float64) {

	anchorFreq := float64(1) / (float64(re1.n) * float64(re2.n) * float64(re3.n) * float64(re4.n) * float64(re5.n))
	// 拉的次數
	// ntimes := 100000000
	// lamda = 理論機率x拉的次數
	lamda := anchorFreq * ntimes

	fmt.Println("==================================")
	fmt.Println("玩拉霸次數:", ntimes)
	fmt.Println("參數lamda:", lamda)
	fmt.Println("大獎完全沒開出機率:", poissonDistribution(0, lamda))

	fmt.Println("大獎至少發生一次的機率:", 1-poissonDistribution(0, lamda))
	fmt.Println("大獎發生1次的機率:", poissonDistribution(1, lamda))
	fmt.Println("大獎發生2次的機率:", poissonDistribution(2, lamda))
	fmt.Println("大獎發生3次的機率:", poissonDistribution(3, lamda))
	fmt.Println("大獎發生4次的機率:", poissonDistribution(4, lamda))
	fmt.Println("大獎發生5次的機率:", poissonDistribution(5, lamda))
	fmt.Println("大獎發生6次的機率:", poissonDistribution(6, lamda))
	fmt.Println("大獎發生7次的機率:", poissonDistribution(7, lamda))
	fmt.Println("大獎發生8次的機率:", poissonDistribution(8, lamda))
	fmt.Println("大獎發生9次的機率:", poissonDistribution(9, lamda))
	fmt.Println("大獎發生10次的機率:", poissonDistribution(10, lamda))

	aboveFive := float64(0)
	for i := 0; i < 5; i++ {
		aboveFive += poissonDistribution(float64(i), lamda)
	}

	fmt.Println("大獎發生5次(含)以上的機率:", 1-aboveFive)
	fmt.Println("==================================")
}

func main() {

	/*chanceOne1 := &reel{[]string{"W", "H", "JP", "C", "S", "F", "H", "A", "W", "D", "G", "F", "C", "D", "B", "E", "A", "G", "A", "E", "B", "F"}, 22, 3}
	chanceTwo2 := &reel{[]string{"S", "E", "E", "G", "W", "A", "D", "H", "B", "D", "C", "S", "B", "G", "JP", "C", "F", "H", "B", "F", "G", "H"}, 22, 3}
	chanceThree3 := &reel{[]string{"E", "A", "F", "D", "A", "B", "H", "C", "H", "W", "B", "G", "F", "D", "D", "W", "S", "G", "W", "C", "E", "JP"}, 22, 3}
	chanceFour4 := &reel{[]string{"E", "B", "JP", "H", "E", "B", "C", "D", "W", "F", "W", "E", "F", "D", "G", "A", "A", "C", "G", "S", "H", "G"}, 22, 3}
	chanceFive5 := &reel{[]string{"F", "H", "E", "W", "F", "G", "G", "C", "C", "D", "B", "D", "C", "D", "E", "S", "W", "A", "B", "JP", "A", "H"}, 22, 3}*/

	/*chanceOne1a := &reel{[]string{"S", "F", "C", "JP", "E", "C", "D", "A", "E", "D", "F", "H", "C", "E", "G", "B", "G", "H", "W", "A", "D", "A", "F", "W", "B"}, 25, 3}
	chanceTwo2a := &reel{[]string{"F", "C", "H", "H", "W", "D", "S", "C", "G", "F", "JP", "D", "D", "B", "E", "B", "B", "G", "E", "E", "S", "A", "H", "F", "C"}, 25, 3}
	chanceThree3a := &reel{[]string{"H", "A", "F", "B", "S", "E", "B", "C", "D", "F", "A", "H", "G", "G", "C", "B", "E", "G", "JP", "D", "H", "W", "W", "W", "A"}, 25, 3}
	chanceFour4a := &reel{[]string{"E", "D", "W", "H", "E", "JP", "G", "F", "G", "G", "H", "D", "C", "C", "A", "A", "W", "E", "A", "B", "S", "F", "B", "B", "F"}, 25, 3}
	chanceFive5a := &reel{[]string{"B", "G", "C", "W", "JP", "H", "E", "F", "A", "S", "D", "E", "A", "D", "W", "D", "B", "B", "F", "G", "C", "E", "G", "H", "C"}, 25, 3}*/

	chanceOne := &reel{[]string{"S", "F", "C", "G", "JP", "E", "C", "D", "F", "A", "E", "D", "F", "H", "C", "E", "G", "D", "B", "G", "H", "W", "E", "C", "A", "D", "A", "E", "F", "W", "B", "C", "F", "H", "D"}, 35, 3}
	chanceTwo := &reel{[]string{"F", "C", "H", "H", "W", "E", "B", "D", "F", "S", "C", "G", "F", "JP", "F", "E", "D", "D", "B", "E", "B", "B", "G", "E", "E", "H", "D", "S", "A", "H", "F", "B", "C", "H", "D"}, 35, 3}
	chanceThree := &reel{[]string{"H", "E", "A", "C", "F", "B", "S", "H", "E", "B", "C", "D", "F", "A", "H", "G", "E", "H", "G", "C", "B", "F", "E", "B", "G", "JP", "D", "H", "F", "W", "W", "W", "A", "C", "B"}, 35, 3}
	chanceFour := &reel{[]string{"C", "E", "H", "B", "D", "W", "H", "F", "E", "JP", "G", "F", "C", "G", "G", "H", "D", "D", "C", "C", "A", "A", "W", "E", "A", "B", "S", "B", "F", "B", "B", "F", "H", "F", "D"}, 35, 3}
	chanceFive := &reel{[]string{"B", "G", "E", "C", "W", "JP", "H", "E", "B", "F", "C", "F", "A", "S", "B", "D", "G", "D", "E", "A", "D", "W", "E", "C", "D", "B", "B", "F", "G", "C", "E", "G", "H", "C", "D"}, 35, 3}
	fmt.Println("===========================================")

	s1 := time.Now()
	//fmt.Println(criterion(chanceOne1, chanceTwo2, chanceThree3, chanceFour4, chanceFive5, 5))

	fmt.Println("排列總數:", runPermutationsOfReels(chanceOne, chanceTwo, chanceThree, chanceFour, chanceFive))
	fmt.Println("===========================================")

	baseAndfreeEV(chanceOne, chanceTwo, chanceThree, chanceFour, chanceFive, chanceOne, chanceTwo, chanceThree, chanceFour, chanceFive, 5, 9)

	estimationOfJP(chanceOne, chanceTwo, chanceThree, chanceFour, chanceFive, 1000)
	estimationOfJP(chanceOne, chanceTwo, chanceThree, chanceFour, chanceFive, 10000)
	estimationOfJP(chanceOne, chanceTwo, chanceThree, chanceFour, chanceFive, 100000)
	estimationOfJP(chanceOne, chanceTwo, chanceThree, chanceFour, chanceFive, 1000000)
	estimationOfJP(chanceOne, chanceTwo, chanceThree, chanceFour, chanceFive, 10000000)

	fmt.Println("總共耗時:", time.Since(s1))

}
