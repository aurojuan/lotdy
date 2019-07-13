package main

import (
	"fmt"
	"math/rand"
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

// 拉出來計算base game裡沒中獎的case  要注意這裡沒寫scatter 進函數simulation()時 要自行加條件篩
// size自行寫死為3 4 5連線
func dealWithGetNothing(win [][]string, pictureArray []string) bool {
	dealZero := 0
	for _, v := range pictureArray {
		dealZero += sumOfBoolArray(judgeSizeAndSymbol(win, 3, v)) + sumOfBoolArray(judgeSizeAndSymbol(win, 4, v)) + sumOfBoolArray(judgeSizeAndSymbol(win, 5, v))
	}
	if dealZero == 0 {
		return true
	}
	return false
}

// 製作隨機進free game的權重機制
func destinyOfFG() int {
	chanceOfFG := rand.Intn(5)

	var caseOfGame int

	if chanceOfFG == 0 {
		caseOfGame = 3 // 20倍
	} else if chanceOfFG == 1 || chanceOfFG == 2 || chanceOfFG == 3 {
		caseOfGame = 1 // 10倍
	} else if chanceOfFG == 4 || chanceOfFG == 5 {
		caseOfGame = 2 // 15倍
	}

	return caseOfGame
}

func freeGamePlay(refr1 *reel, refr2 *reel, refr3 *reel, refr4 *reel, refr5 *reel, numOfReels int, numOfPicture int, caseOfGame int, numOfMock int, singleBet float64) float64 {

	fgBalance := float64(0)

	var fcountOfA3, fcountOfA4, fcountOfA5, fcountOfB3, fcountOfB4, fcountOfB5, fcountOfC3, fcountOfC4, fcountOfC5, fcountOfD3, fcountOfD4, fcountOfD5,
		fcountOfE3, fcountOfE4, fcountOfE5, fcountOfF3, fcountOfF4, fcountOfF5, fcountOfG3, fcountOfG4, fcountOfG5, fcountOfH3, fcountOfH4, fcountOfH5,
		fcountOfS3, fcountOfS4, fcountOfS5 int64

	fcountOfA3, fcountOfA4, fcountOfA5, fcountOfB3, fcountOfB4, fcountOfB5, fcountOfC3, fcountOfC4, fcountOfC5, fcountOfD3, fcountOfD4, fcountOfD5,
		fcountOfE3, fcountOfE4, fcountOfE5, fcountOfF3, fcountOfF4, fcountOfF5, fcountOfG3, fcountOfG4, fcountOfG5, fcountOfH3, fcountOfH4, fcountOfH5,
		fcountOfS3, fcountOfS4, fcountOfS5 = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0

	// 前面5個參數應該不是必要  之後再改
	payset := payTable(refr1, refr2, refr3, refr4, refr5, numOfPicture, caseOfGame)

	for mock := 0; mock < numOfMock; mock++ {

		//fmt.Println(balance)
		// 從這裡開始隨機生成每個輪子的起頭
		rand.Seed(int64(time.Now().UnixNano()))
		chance1f := rand.Intn(refr1.n)
		chance2f := rand.Intn(refr2.n)
		chance3f := rand.Intn(refr3.n)
		chance4f := rand.Intn(refr4.n)
		chance5f := rand.Intn(refr5.n)

		win := reelsToMatrix(refr1, refr2, refr3, refr4, refr5, numOfReels, chance1f, chance2f, chance3f, chance4f, chance5f)
		// print2DMatrix(win, 3)
		// fmt.Println("===========================================")

		fcountOfA5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "A")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "A"))) * payset[0][2] * singleBet
		fcountOfA4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "A")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "A"))) * payset[0][1] * singleBet
		fcountOfA3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "A")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "A"))) * payset[0][0] * singleBet

		fcountOfB5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "B")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "B"))) * payset[1][2] * singleBet
		fcountOfB4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "B")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "B"))) * payset[1][1] * singleBet
		fcountOfB3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "B")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "B"))) * payset[1][0] * singleBet

		fcountOfC5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "C")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "C"))) * payset[2][2] * singleBet
		fcountOfC4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "C")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "C"))) * payset[2][1] * singleBet
		fcountOfC3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "C")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "C"))) * payset[2][0] * singleBet

		fcountOfD5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "D")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "D"))) * payset[3][2] * singleBet
		fcountOfD4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "D")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "D"))) * payset[3][1] * singleBet
		fcountOfD3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "D")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "D"))) * payset[3][0] * singleBet

		fcountOfE5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "E")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "E"))) * payset[4][2] * singleBet
		fcountOfE4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "E")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "E"))) * payset[4][1] * singleBet
		fcountOfE3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "E")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "E"))) * payset[4][0] * singleBet

		fcountOfF5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "F")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "F"))) * payset[5][2] * singleBet
		fcountOfF4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "F")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "F"))) * payset[5][1] * singleBet
		fcountOfF3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "F")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "F"))) * payset[5][0] * singleBet

		fcountOfG5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "G")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "G"))) * payset[6][2] * singleBet
		fcountOfG4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "G")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "G"))) * payset[6][1] * singleBet
		fcountOfG3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "G")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "G"))) * payset[6][0] * singleBet

		fcountOfH5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "H")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "H"))) * payset[7][2] * singleBet
		fcountOfH4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "H")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "H"))) * payset[7][1] * singleBet
		fcountOfH3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "H")))
		fgBalance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "H"))) * payset[7][0] * singleBet

		fcriS5 := countOf5Scatter(win)
		// S S S S #S
		fcriS4 := countOf4Scatter(win)
		// 暫定S S S #S any
		fcriS3 := countOf3Scatter(win)

		if fcriS5 {
			fcountOfS5++
			//print2DMatrix(win, 3)
			//fmt.Println("===========================================")
		}

		// S S S S #S
		if fcriS4 {
			fcountOfS4++
			//print2DMatrix(win, 3)
			//fmt.Println("===========================================")
		}

		// 暫定S S S #S any
		if fcriS3 {
			fcountOfS3++
			//print2DMatrix(win, 3)
			//fmt.Println("===========================================")
		}

	}
	//調試用
	/* fmt.Println("fcountOfA3, fcountOfA4, fcountOfA5, fcountOfB3, fcountOfB4, fcountOfB5, fcountOfC3, fcountOfC4, fcountOfC5, fcountOfD3, fcountOfD4, fcountOfD5",
		"fcountOfE3, fcountOfE4, fcountOfE5, fcountOfF3, fcountOfF4, fcountOfF5, fcountOfG3, fcountOfG4, fcountOfG5, fcountOfH3, fcountOfH4, fcountOfH5",
		"fcountOfS3, fcountOfS4, fcountOfS5", fcountOfA3, fcountOfA4, fcountOfA5, fcountOfB3, fcountOfB4, fcountOfB5, fcountOfC3, fcountOfC4, fcountOfC5, fcountOfD3, fcountOfD4, fcountOfD5,
		fcountOfE3, fcountOfE4, fcountOfE5, fcountOfF3, fcountOfF4, fcountOfF5, fcountOfG3, fcountOfG4, fcountOfG5, fcountOfH3, fcountOfH4, fcountOfH5,
		fcountOfS3, fcountOfS4, fcountOfS5)

	// 調試用
	fmt.Println("fgBalance:", fgBalance) */
	return fgBalance

}

// 目前設計上有綁特定pay table
func simulation(re1 *reel, re2 *reel, re3 *reel, re4 *reel, re5 *reel, refr1 *reel, refr2 *reel, refr3 *reel, refr4 *reel, refr5 *reel, numOfReels int, numOfPicture int, numOfMock int, initialBalace float64, singleBet float64) {
	var countOfA3, countOfA4, countOfA5, countOfB3, countOfB4, countOfB5, countOfC3, countOfC4, countOfC5, countOfD3, countOfD4, countOfD5,
		countOfE3, countOfE4, countOfE5, countOfF3, countOfF4, countOfF5, countOfG3, countOfG4, countOfG5, countOfH3, countOfH4, countOfH5,
		countOfS3, countOfS4, countOfS5 int64 // 方法數可能爆炸大

	countOfA3, countOfA4, countOfA5, countOfB3, countOfB4, countOfB5, countOfC3, countOfC4, countOfC5, countOfD3, countOfD4, countOfD5,
		countOfE3, countOfE4, countOfE5, countOfF3, countOfF4, countOfF5, countOfG3, countOfG4, countOfG5, countOfH3, countOfH4, countOfH5,
		countOfS3, countOfS4, countOfS5 = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0

	balance := initialBalace
	payset := payTable(re1, re2, re3, re4, re5, numOfPicture, 0)
	loseTimes := int64(0)
	playTimes := int64(0)
	// 另外加的變數  配合原本沒scatter的設計  追蹤做最後算RTP用
	playFGBalance := float64(0)

	// 計量JP用
	countOfJP := int64(0)

	for mock := 0; mock < numOfMock; mock++ {

		if balance <= float64(0) {
			break
		}

		if balance < singleBet {
			fmt.Println("總金額不足 遊戲結束 請降低投注額或加值!!  懂")
			break
		}
		//fmt.Println(balance)
		// 從這裡開始隨機生成每個輪子的起頭
		rand.Seed(int64(time.Now().UnixNano()))
		chance1 := rand.Intn(re1.n)
		chance2 := rand.Intn(re2.n)
		chance3 := rand.Intn(re3.n)
		chance4 := rand.Intn(re4.n)
		chance5 := rand.Intn(re5.n)

		win := reelsToMatrix(re1, re2, re3, re4, re5, numOfReels, chance1, chance2, chance3, chance4, chance5)
		// print2DMatrix(win, 3)
		// fmt.Println("===========================================")

		// 每一回合 計量scatter以外  各symbol fix size時 中線數  並根據中線數給獎
		// 記得處理都沒中  要扣款的case
		countOfA5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "A")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "A"))) * payset[0][2] * singleBet
		countOfA4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "A")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "A"))) * payset[0][1] * singleBet
		countOfA3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "A")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "A"))) * payset[0][0] * singleBet

		countOfB5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "B")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "B"))) * payset[1][2] * singleBet
		countOfB4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "B")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "B"))) * payset[1][1] * singleBet
		countOfB3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "B")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "B"))) * payset[1][0] * singleBet

		countOfC5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "C")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "C"))) * payset[2][2] * singleBet
		countOfC4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "C")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "C"))) * payset[2][1] * singleBet
		countOfC3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "C")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "C"))) * payset[2][0] * singleBet

		countOfD5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "D")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "D"))) * payset[3][2] * singleBet
		countOfD4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "D")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "D"))) * payset[3][1] * singleBet
		countOfD3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "D")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "D"))) * payset[3][0] * singleBet

		countOfE5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "E")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "E"))) * payset[4][2] * singleBet
		countOfE4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "E")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "E"))) * payset[4][1] * singleBet
		countOfE3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "E")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "E"))) * payset[4][0] * singleBet

		countOfF5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "F")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "F"))) * payset[5][2] * singleBet
		countOfF4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "F")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "F"))) * payset[5][1] * singleBet
		countOfF3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "F")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "F"))) * payset[5][0] * singleBet

		countOfG5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "G")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "G"))) * payset[6][2] * singleBet
		countOfG4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "G")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "G"))) * payset[6][1] * singleBet
		countOfG3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "G")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "G"))) * payset[6][0] * singleBet

		countOfH5 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "H")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 5, "H"))) * payset[7][2] * singleBet
		countOfH4 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "H")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 4, "H"))) * payset[7][1] * singleBet
		countOfH3 += int64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "H")))
		balance += float64(sumOfBoolArray(judgeSizeAndSymbol(win, 3, "H"))) * payset[7][0] * singleBet

		criS5 := countOf5Scatter(win)
		// S S S S #S
		criS4 := countOf4Scatter(win)
		// 暫定S S S #S any
		criS3 := countOf3Scatter(win)

		if criS5 {
			countOfS5++
			caseOfGameS5 := destinyOfFG()
			//fmt.Println(caseOfGameS5)
			// Scatter 由左到右不同column共5連線  可玩free game 20次
			numOfMockS5 := 20
			freeGameBalance := freeGamePlay(refr1, refr2, refr3, refr4, refr5, numOfReels, numOfPicture, caseOfGameS5, numOfMockS5, singleBet)
			balance += freeGameBalance
			playFGBalance += freeGameBalance
			//print2DMatrix(win, 3)
			//fmt.Println("===========================================")
		}

		// S S S S #S
		if criS4 {
			countOfS4++
			caseOfGameS4 := destinyOfFG()
			//fmt.Println(caseOfGameS4)
			// Scatter 由左到右不同column共4連線  可玩free game 10次
			numOfMockS4 := 10
			freeGameBalance := freeGamePlay(refr1, refr2, refr3, refr4, refr5, numOfReels, numOfPicture, caseOfGameS4, numOfMockS4, singleBet)
			balance += freeGameBalance
			playFGBalance += freeGameBalance
			//print2DMatrix(win, 3)
			//fmt.Println("===========================================")
		}

		// 暫定S S S #S any
		if criS3 {
			countOfS3++
			caseOfGameS3 := destinyOfFG()
			//fmt.Println(caseOfGameS3)
			// Scatter 由左到右不同column共3連線  可玩free game 5次
			numOfMockS3 := 5
			freeGameBalance := freeGamePlay(refr1, refr2, refr3, refr4, refr5, numOfReels, numOfPicture, caseOfGameS3, numOfMockS3, singleBet)
			balance += freeGameBalance
			playFGBalance += freeGameBalance
			//print2DMatrix(win, 3)
			//fmt.Println("===========================================")
		}

		if dealWithGetNothing(win, []string{"A", "B", "C", "D", "E", "F", "G", "H"}) && !criS5 && !criS4 && !criS3 {
			balance -= singleBet // 不同case數自行調整
			loseTimes++
		}

		// 測JP
		if win[1][0] == "JP" && win[1][1] == "JP" && win[1][2] == "JP" && win[1][3] == "JP" && win[1][4] == "JP" {
			countOfJP++
			print2DMatrix(win, 3)
			fmt.Println("===========================================")
		}

		playTimes++

	}

	totalState := int64(0)
	// 圖案數有變更時  記得自行修改
	totalChanceOfState := make([]int64, 27)
	totalChanceOfState = []int64{countOfA3, countOfA4, countOfA5, countOfB3, countOfB4, countOfB5, countOfC3, countOfC4, countOfC5, countOfD3, countOfD4, countOfD5,
		countOfE3, countOfE4, countOfE5, countOfF3, countOfF4, countOfF5, countOfG3, countOfG4, countOfG5, countOfH3, countOfH4, countOfH5,
		countOfS3, countOfS4, countOfS5}

	for _, v := range totalChanceOfState {
		totalState += v
	}

	returnAmount := float64(0)
	for i := 0; i < numOfPicture; i++ {
		for j := 0; j < 3; j++ {
			returnAmount += float64(totalChanceOfState[3*i+j]) * payset[i][j] * singleBet
		}
	}

	// 一般玩法拿到的 + 進free game拿到的
	returnAmount = returnAmount + playFGBalance

	// totalChance := float64(totalState) / float64(numOfMock)

	var loseRatio float64

	fgCall := countOfS3 + countOfS4 + countOfS5
	var fgRatio float64

	// playTimes must equal or less than nomOfMock
	if playTimes != int64(numOfMock) {
		loseRatio = float64(loseTimes) / float64(playTimes)
		fgRatio = float64(fgCall) / float64(playTimes)
	} else {
		loseRatio = float64(loseTimes) / float64(numOfMock)
		fgRatio = float64(fgCall) / float64(numOfMock)
	}

	totalIn := singleBet * float64(playTimes) // 現在是9個case押注  不同case數自行調整 更新:又調回來了

	rtp := returnAmount / totalIn

	// 統計各symbol出現的相對頻率 圖案數有變更時  記得自行修改
	frequencyOfSymbols := make([]float64, 27)
	var freqA3, freqA4, freqA5, freqB3, freqB4, freqB5, freqC3, freqC4, freqC5, freqD3, freqD4, freqD5, freqE3,
		freqE4, freqE5, freqF3, freqF4, freqF5, freqG3, freqG4, freqG5, freqH3, freqH4, freqH5, freqS3, freqS4, freqS5 float64
	frequencyOfSymbols = []float64{freqA3, freqA4, freqA5, freqB3, freqB4, freqB5, freqC3, freqC4, freqC5, freqD3,
		freqD4, freqD5, freqE3, freqE4, freqE5, freqF3, freqF4, freqF5, freqG3, freqG4, freqG5, freqH3, freqH4,
		freqH5, freqS3, freqS4, freqS5}

	for i, v := range totalChanceOfState {
		frequencyOfSymbols[i] = float64(v) / float64(playTimes)
	}

	fmt.Println("模擬次數, 初始餘額, 單次下注:", numOfMock, initialBalace, singleBet)
	fmt.Println("實際玩的次數:", playTimes)
	fmt.Println("最終餘額:", balance)
	fmt.Println("餘額變化:", balance-initialBalace)
	fmt.Println("總投入金額:", totalIn)

	fmt.Println("return amount:", returnAmount)
	fmt.Println("RTP:", rtp*100, "%")

	fmt.Println("以下為各symbol hits紀錄")
	fmt.Println("中A 3連 4連 5連:", countOfA3, countOfA4, countOfA5)
	fmt.Println("中B 3連 4連 5連:", countOfB3, countOfB4, countOfB5)
	fmt.Println("中C 3連 4連 5連:", countOfC3, countOfC4, countOfC5)
	fmt.Println("中D 3連 4連 5連:", countOfD3, countOfD4, countOfD5)
	fmt.Println("中E 3連 4連 5連:", countOfE3, countOfE4, countOfE5)
	fmt.Println("中F 3連 4連 5連:", countOfF3, countOfF4, countOfF5)
	fmt.Println("中G 3連 4連 5連:", countOfG3, countOfG4, countOfG5)
	fmt.Println("中H 3連 4連 5連:", countOfH3, countOfH4, countOfH5)
	fmt.Println("中Scatter 3連 4連 5連", countOfS3, countOfS4, countOfS5)

	fmt.Println("以下為各symbol發生頻率統計")
	fmt.Println("中A 3連 4連 5連:", frequencyOfSymbols[0], frequencyOfSymbols[1], frequencyOfSymbols[2])
	fmt.Println("中B 3連 4連 5連:", frequencyOfSymbols[3], frequencyOfSymbols[4], frequencyOfSymbols[5])
	fmt.Println("中C 3連 4連 5連:", frequencyOfSymbols[6], frequencyOfSymbols[7], frequencyOfSymbols[8])
	fmt.Println("中D 3連 4連 5連:", frequencyOfSymbols[9], frequencyOfSymbols[10], frequencyOfSymbols[11])
	fmt.Println("中E 3連 4連 5連:", frequencyOfSymbols[12], frequencyOfSymbols[13], frequencyOfSymbols[14])
	fmt.Println("中F 3連 4連 5連:", frequencyOfSymbols[15], frequencyOfSymbols[16], frequencyOfSymbols[17])
	fmt.Println("中G 3連 4連 5連:", frequencyOfSymbols[18], frequencyOfSymbols[19], frequencyOfSymbols[20])
	fmt.Println("中H 3連 4連 5連:", frequencyOfSymbols[21], frequencyOfSymbols[22], frequencyOfSymbols[23])
	fmt.Println("中Scatter 3連 4連 5連", frequencyOfSymbols[24], frequencyOfSymbols[25], frequencyOfSymbols[26])

	fmt.Println("總中線次數:", totalState)
	// fmt.Println("總中比率:", totalChance)
	fmt.Println("總輸次數:", loseTimes)
	fmt.Println("輸率:", loseRatio)
	fmt.Println("贏率:", 1-loseRatio)
	fmt.Println("中free game次數:", fgCall)
	fmt.Println("中free game比率:", fgRatio)
	fmt.Println("=================JP分隔線===================")
	fmt.Println("中JP次數:", countOfJP)
	fmt.Println("中JP比率:", float64(countOfJP)/float64(playTimes))

}

// 暫時用手動在function內寫死的方式
// 之後再補array沒填滿或非法值的宣告
// caseOfGame  0:base game    1:free game with each pay*10  2:free game with each pay*15  3:free game with each pay*20
// 之後應該要補不合理input的報錯機制
func payTable(re1 *reel, re2 *reel, re3 *reel, re4 *reel, re5 *reel, numOfPicture int, caseOfGame int) [][]float64 {

	ArrayOfPayTable := make([][]float64, numOfPicture) // row
	for index := range ArrayOfPayTable {
		ArrayOfPayTable[index] = make([]float64, 3) // for 3 4 5連線
	}

	// 9x3 matrix   pay nothing for all scatter
	// 改動調整RTP的重點處之一
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

	/*fmt.Println("pay of A       3 4 5:", ArrayOfPayTable[0][:])
	fmt.Println("pay of B       3 4 5:", ArrayOfPayTable[1][:])
	fmt.Println("pay of C       3 4 5:", ArrayOfPayTable[2][:])
	fmt.Println("pay of D       3 4 5:", ArrayOfPayTable[3][:])
	fmt.Println("pay of E       3 4 5:", ArrayOfPayTable[4][:])
	fmt.Println("pay of F       3 4 5:", ArrayOfPayTable[5][:])
	fmt.Println("pay of G       3 4 5:", ArrayOfPayTable[6][:])
	fmt.Println("pay of H       3 4 5:", ArrayOfPayTable[7][:])
	fmt.Println("pay of Scatter 3 4 5:", ArrayOfPayTable[8][:])*/

	return ArrayOfPayTable

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

	/*chanceOne := &reel{[]string{"S", "F", "C", "G", "JP", "S", "C", "D", "F", "A", "E", "D", "F", "H", "C", "E", "G", "D", "B", "G", "H", "W", "E", "C", "A", "D", "A", "E", "F", "W", "B", "C", "F", "H", "D"}, 35, 3}
	chanceTwo := &reel{[]string{"F", "C", "H", "H", "W", "E", "B", "D", "F", "S", "C", "G", "F", "JP", "S", "E", "D", "D", "B", "E", "B", "B", "G", "E", "E", "H", "D", "S", "A", "H", "F", "B", "C", "H", "D"}, 35, 3}
	chanceThree := &reel{[]string{"H", "E", "A", "C", "F", "B", "S", "H", "E", "B", "C", "D", "F", "A", "H", "G", "E", "H", "G", "C", "B", "F", "E", "B", "G", "JP", "D", "S", "F", "W", "W", "W", "A", "C", "B"}, 35, 3}
	chanceFour := &reel{[]string{"C", "S", "H", "B", "D", "W", "H", "F", "E", "JP", "G", "F", "C", "G", "G", "H", "D", "D", "C", "C", "A", "A", "W", "E", "A", "B", "S", "B", "F", "B", "B", "F", "H", "F", "D"}, 35, 3}
	chanceFive := &reel{[]string{"B", "G", "E", "C", "W", "JP", "H", "E", "B", "F", "C", "F", "A", "S", "B", "D", "G", "D", "E", "A", "D", "W", "E", "C", "D", "B", "B", "S", "G", "C", "E", "G", "H", "C", "D"}, 35, 3}*/

	zzz1 := &reel{[]string{"S", "F", "C", "G", "JP", "E", "C", "D", "F", "A", "S", "D", "F", "H", "C", "S", "G", "D", "B", "G", "H", "W", "S", "C", "A", "D", "A", "E", "S", "W", "B", "C", "S", "H", "D"}, 35, 3}
	zzz2 := &reel{[]string{"F", "C", "H", "H", "W", "E", "B", "S", "F", "S", "C", "G", "F", "JP", "F", "E", "D", "D", "B", "E", "B", "B", "G", "E", "E", "H", "D", "S", "A", "S", "F", "B", "C", "S", "D"}, 35, 3}
	zzz3 := &reel{[]string{"H", "E", "A", "C", "F", "B", "S", "H", "E", "B", "C", "D", "F", "A", "H", "G", "E", "H", "G", "C", "B", "F", "E", "B", "G", "JP", "D", "H", "F", "W", "W", "W", "A", "C", "B"}, 35, 3}
	zzz4 := &reel{[]string{"C", "E", "H", "B", "D", "W", "H", "F", "E", "JP", "G", "F", "C", "G", "G", "H", "D", "D", "C", "C", "A", "A", "W", "E", "A", "B", "S", "B", "F", "B", "B", "F", "H", "F", "D"}, 35, 3}
	zzz5 := &reel{[]string{"B", "G", "E", "C", "W", "JP", "H", "E", "B", "F", "C", "F", "A", "S", "B", "D", "G", "D", "E", "A", "D", "W", "E", "C", "D", "B", "B", "F", "G", "C", "E", "G", "H", "C", "D"}, 35, 3}

	s1 := time.Now()

	fmt.Println("===========================================")

	simulation(chanceOne, chanceTwo, chanceThree, chanceFour, chanceFive, chanceOne, chanceTwo, chanceThree, chanceFour, chanceFive, 5, 9, 10000, 20, 4.7)

	simulation(zzz1, zzz2, zzz3, zzz4, zzz5, zzz1, zzz2, zzz3, zzz4, zzz5, 5, 9, 1000, 1000, 1)
	fmt.Println("總共耗時:", time.Since(s1))

}
