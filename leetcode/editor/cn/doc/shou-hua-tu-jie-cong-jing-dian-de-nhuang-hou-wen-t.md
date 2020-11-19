#### 起初思路
以 4 皇后为例，我画出一个搜索树，初始时棋盘的格子都是`"."`。
![image.png](https://pic.leetcode-cn.com/1599090972-ffZdFD-image.png)

- 一行选一个格子置为`"Q"`，一行行地往下选，第一行有四种选择。
- 在选下一行的皇后时，为了不发生列的冲突，有三种选择。
- 继续选下去，构建完整解，可能会遇到对角线冲突。
- 遇到冲突，继续选下去没有意义，得不出合法的解。需要回溯。

回溯的套路（很重要）：
1. 遍历枚举出所有可能的选择。
2. 依次尝试这些选择：作出一种选择，并往下递归。
3. 如果这个选择产生不出正确的解，要撤销这个选择（将当前置为 "Q" 的格子恢复为 "."），回到之前的状态，并作出下一个可用的选择。

这是一个「选择、探索、撤销选择」的过程。识别出死胡同，就回溯，尝试下一个点，不做无效的搜索。

#### 思路修正

- 我在枚举选择时，规避了行和列的冲突，前面画的搜索树已经是修剪后的。

- 但我还是想复杂了，对角线冲突和行列冲突一起作为约束就好，直接充分剪枝。如下图。

  - 剪枝做法是遍历之前的行，如果当前的格子和皇后们 同列或同对角线，则跳过该点
  - （这需要优化，后面会讲）。

![image.png](https://pic.leetcode-cn.com/1599102658-ubFHgg-image.png)

- 你看上图左边两个叶子节点，下一行怎么放都冲突，可选的选项都被剪完了，

- 当所有可选的选择迭代完，当前递归分支就结束，撤销最后的选择，回到上一层，切入另一个分支。

- 当填完第四行，如上图的绿钩，生成了一个解，加入解集，并返回（这里不返回也行，因为已经做了充分的剪枝，不返回就会走一遍迭代，递归也结束），开始回溯，继续寻找完整解。



#### 回溯的三要点

1. 选择，决定了搜索空间，决定了搜索空间有哪些节点。
2. 约束，用来剪枝，避免进入无效的分支。
3. 目标，决定了什么时候捕获有效的解，提前结束递归，开始回溯。



#### 代码
Runtime: 4 ms, faster than 92.71% of Go online submissions for N-Queens.
```js []
const solveNQueens = (n) => {
  const board = new Array(n);
  for (let i = 0; i < n; i++) {     // 棋盘的初始化
    board[i] = new Array(n).fill('.');
  }
  const res = [];
  const isValid = (row, col) => {  
    for (let i = 0; i < row; i++) { // 之前的行
      for (let j = 0; j < n; j++) { // 所有的列
        if (board[i][j] == 'Q' &&   // 发现了皇后，并且和自己同列/对角线
          (j == col || i + j === row + col || i - j === row - col)) {
          return false;             // 不是合法的选择
        }
      }
    }
    return true;
  };
  const helper = (row) => {   // 放置当前行的皇后
    if (row == n) {           // 递归的出口，超出了最后一行
      const stringsBoard = board.slice(); // 拷贝一份board
      for (let i = 0; i < n; i++) {
        stringsBoard[i] = stringsBoard[i].join(''); // 将每一行拼成字符串
      }
      res.push(stringsBoard); // 推入res数组
      return;
    }
    for (let col = 0; col < n; col++) { // 枚举出所有选择
      if (isValid(row, col)) {          // 剪掉无效的选择
        board[row][col] = "Q";          // 作出选择，放置皇后
        helper(row + 1);                // 继续选择，往下递归
        board[row][col] = '.';          // 撤销当前选择
      }
    }
  };
  helper(0);  // 从第0行开始放置
  return res;
};
```
```go []
func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	res := [][]string{}
	helper(0, bd, &res, n)
	return res
}
func helper(r int, bd [][]string, res *[][]string, n int) {
	if r == n {
		temp := make([]string, len(bd))
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, temp)
		return
	}
	for c := 0; c < n; c++ {
		if isValid(r, c, n, bd) {
			bd[r][c] = "Q"
			helper(r+1, bd, res, n)
			bd[r][c] = "."
		}
	}

}
func isValid(r, c int, n int, bd [][]string) bool {
	for i := 0; i < r; i++ {
		for j := 0; j < n; j++ {
			if bd[i][j] == "Q" && (j == c || i+j == r+c || i-j == r-c) {
				return false
			}
		}
	}
	return true
}
```

#### 用空间换时间，进行优化
这道题需要记录之前放皇后的位置，才能结合约束条件去做剪枝。

感谢weiwei哥提醒：每次都调用 isValid 遍历一遍前面的格子，效率是不优的。

最好是用三个数组或 Set 去记录出现过皇后的列们、正对角线们、反对角线们，用空间换取时间。

#### 优化后的代码
Runtime: 4 ms, faster than 92.71% of Go online submissions for N-Queens.
```js []
const solveNQueens = (n) => {
  const board = new Array(n);
  for (let i = 0; i < n; i++) {
    board[i] = new Array(n).fill('.');
  }

  const cols = new Set();  // 列集，记录出现过皇后的列
  const diag1 = new Set(); // 正对角线集
  const diag2 = new Set(); // 反对角线集
  const res = [];

  const helper = (row) => {
    if (row == n) {
      const stringsBoard = board.slice();
      for (let i = 0; i < n; i++) {
        stringsBoard[i] = stringsBoard[i].join('');
      }
      res.push(stringsBoard);
      return;
    }
    for (let col = 0; col < n; col++) {
      // 如果当前点的所在的列，所在的对角线都没有皇后，即可选择，否则，跳过
      if (!cols.has(col) && !diag1.has(row + col) && !diag2.has(row - col)) { 
        board[row][col] = 'Q';  // 放置皇后
        cols.add(col);          // 记录放了皇后的列
        diag1.add(row + col);   // 记录放了皇后的正对角线
        diag2.add(row - col);   // 记录放了皇后的负对角线
        helper(row + 1);
        board[row][col] = '.';  // 撤销该点的皇后
        cols.delete(col);       // 对应的记录也删一下
        diag1.delete(row + col);
        diag2.delete(row - col);
      }
    }
  };
  helper(0);
  return res;
};
```
```go []
func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	cols := map[int]bool{}
	diag1 := map[int]bool{}
	diag2 := map[int]bool{}

	res := [][]string{}
	helper(0, bd, &res, n, &cols, &diag1, &diag2)
	return res
}
func helper(r int, bd [][]string, res *[][]string, n int, cols, diag1, diag2 *map[int]bool) {
	if r == n {
		temp := make([]string, len(bd))
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, temp)
		return
	}
	for c := 0; c < n; c++ {
		if !(*cols)[c] && !(*diag1)[r+c] && !(*diag2)[r-c] {
			bd[r][c] = "Q"
			(*cols)[c] = true
			(*diag1)[r+c] = true
			(*diag2)[r-c] = true
			helper(r+1, bd, res, n, cols, diag1, diag2)
			bd[r][c] = "."
			(*cols)[c] = false
			(*diag1)[r+c] = false
			(*diag2)[r-c] = false
		}
	}
}
```

#### 文字经过了多次推敲，尽量表达出我的理解，如果有不准确的地方，欢迎友善的反馈，一起完善题解，往后我也会根据大家的反馈，不断调整表述方式和思路。