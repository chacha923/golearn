package dp

import "fmt"

//求最长公共子序列

//import java.util.Random;
//public class LCS {
//
//public static int[][] lengthofLCS(char[] X, char[] Y){
///* 构造二维数组c[][]记录X[i]和Y[j]的LCS长度 (i,j)是前缀
// * c[i][j]=0; 当 i = j = 0;
// * c[i][j]=c[i-1][j-1]+1; 当 i = j > 0; Xi == Y[i]
// * c[i][j]=max(c[i-1][j],c[i][j+1]); 当 i = j > 0; X[i] != Y[i]
// * 需要计算 m*n 个子问题的长度 即 任意c[i][j]的长度
// * -- 填表过程
// */
//int[][]c = new int[X.length+1][Y.length+1];
//
//// 动态规划计算所有子问题
//for(int i=1;i<=X.length;i++){
//for (int j=1;j<=Y.length;j++){
//if(X[i-1]==Y[j-1]){
//c[i][j] = c[i-1][j-1]+1;
//}
//else if(c[i-1][j] >= c[i][j-1]){
//c[i][j] = c[i-1][j];
//}
//else{
//c[i][j] = c[i][j-1];
//}
//}
//}

////打印C数组
//for(int i=0;i<=X.length;i++){
//for (int j=0;j<=Y.length;j++){
//System.out.print(c[i][j]+" ");
//}
//System.out.println();
//}
//return c;
//}
//// 输出LCS序列
//public static void print(int[][] arr, char[] X, char[] Y, int i, int j) {
//if(i == 0 || j == 0)
//return;
//if(X[i-1] == Y[j-1]) {
//System.out.print("element " + X[i-1] + " ");
//// 寻找的
//print(arr, X, Y, i-1, j-1);
//}else if(arr[i-1][j] >= arr[i][j-1]) {
//print(arr, X, Y, i-1, j);
//}else{
//print(arr, X, Y, i, j-1);
//}
//}
//public static void main(String[] args) {
//// TODO Auto-generated method stub
//char[] x ={'A','B','C','B','D','A','B'};
//char[] y ={'B','D','C','A','B','A'};
//int[][] c = lengthofLCS(x,y);
//print(c, x, y, x.length, y.length);
//}
//}

func lcs(x, y string) {
	var c [len(x) + 1][len(y) + 1]int
	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			if x[i-1] == y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else if x[i-1] >= y[j-1] {
				c[i][j] = c[i-1][j]
			} else {
				c[i][j] = c[i][j-1]
			}
		}
	}

	//打印 c数组
	for i := 0; i <= len(x); i++ {
		for j := 0; j <= len(y); j++ {
			fmt.Println(c[i][j])
		}
		fmt.Println()
	}

}

func LCSLength(x []byte, y []byte) int {
	m := len(x)
	n := len(y)
	lookup := make([][]int, m+1)
	for i := range lookup {
		lookup[i] = make([]int, n+1)
	}
	var i, j int
	for i = 0; i <= m; i++ {
		lookup[i][0] = 0
	}
	for j = 0; j <= n; j++ {
		lookup[0][j] = 0
	}

	for i = 1; i <= m; i++ {
		for j = 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				lookup[i][j] = lookup[i-1][j-1] + 1
			} else {
				len1 := lookup[i-1][j]
				len2 := lookup[i][j-1]
				if len1 > len2 {
					lookup[i][j] = len1
				} else {
					lookup[i][j] = len2
				}
			}
		}
	}
	return lookup[m][n]
}
