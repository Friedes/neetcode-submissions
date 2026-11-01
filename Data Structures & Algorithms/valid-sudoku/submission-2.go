func isValidSudoku(board [][]byte) bool {
	row := len(board)
	col := len(board[0])
	//Check row && Check col
	for i:=0;i<row;i++{
		mR := make(map[byte]bool)
		mL := make(map[byte]bool)
		for j:=0;j<col;j++{
			if board[i][j] != '.' {
				if _, ok := mR[board[i][j]]; ok {
					return false
				} else {
					mR[board[i][j]] = true
				}
			}

			if board[j][i] != '.' {
				if _, ok := mL[board[j][i]]; ok {
					return false
				} else {
					mL[board[j][i]] = true
				}
			}
		}
	}

	arr := make([][4]int, 9)
	arr[0] = [4]int{0,0,2,2}
	arr[1] = [4]int{0,3,2,5}
	arr[2] = [4]int{0,6,2,8}
	arr[3] = [4]int{3,0,5,2}
	arr[4] = [4]int{3,3,5,5}
	arr[5] = [4]int{3,6,5,8}
	arr[6] = [4]int{6,0,8,2}
	arr[7] = [4]int{6,3,8,5}
	arr[8] = [4]int{6,6,8,8}

	//Check box
	for _, ar := range arr {
		if !CheckBox(ar, board) {
			return false
		}
	}

	return true
}

func CheckBox(bound [4]int, board [][]byte) bool {
	m := make(map[byte]bool)
	for i:=bound[0];i<=bound[2];i++{
		for j:=bound[1];j<=bound[3];j++{
			if board[i][j] != '.' {
				if _, ok := m[board[i][j]]; ok {
					return false
				} else {
					m[board[i][j]] = true
				}
			}
		}
	}

	return true
}
