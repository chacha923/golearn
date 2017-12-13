package graph

import (
	"strconv"
	"fmt"
)

const (
	INF = 999999
)

//迪杰斯特拉算法
// 求图的源顶点到某顶点的最短路径
//1)      适用条件&范围：
//a)   单源最短路径(从源点s到其它所有顶点v);
//b)   有向图&无向图(无向图可以看作(u,v),(v,u)同属于边集E的有向图)
//c)   所有边权非负(任取(i,j)∈E都有Wij≥0);
//2)      算法描述：
//在带权图中最常遇到的问题就是，寻找两点间的最短路径问题。
//求0点到n-1点的最短路径
//[i][j]tu  用邻接矩阵表示图的路径, 表示i与j点之间有路径
// 是一种贪心算法, 只能解决正权重, 单源节点
func dijkstra( tu [][]int) {
	//假设源节点的index为0
	//假设目标节点的index为count-1
	count := len(tu)
	dis := make([]int, 0, count)	//保存源节点到其他节点的最短路径长度
	book := make([]int, 0, count)	//将节点分为两部分, 如果节点index已知最短路径, 那么book[index]=1, 否则0
	book[0] = 1
	for i:=0; i<count; i++{
		dis = append(dis, tu[0][i])	//初始化dis数组, 此时的dis数组的值为最短路径的估计值
	}

	shortPathStr := make([]string, 0, count)	//保存最短路径上的节点
	shortPathStr = append(shortPathStr, "0")

	var u int = 0

	for i:=0;i<count;i++{
		min := INF
		for j:=0;j<count;j++{
			//更新book[], 找出当前dis[]中最小的值, 那么0节点该值的下标节点的最短路径被确定
			if book[j] == 0 && dis[j] < min {
				min = dis[j]
				u = j
			}
		}
		book[u] = 1

		//松弛 , 以u为中间点，修正从start到未访问各点的距离
		for v:=0;v < count; v++ {
			if tu[u][v] < INF {		//如果存在边uv
				if dis[v] >= dis[u] + tu[u][v] {		//如果0直接到v的最短路径估计值, 大于0到u的最短路径 + uv的边长, 更新dis[v]
					dis[v] = dis[u] + tu[u][v]
					shortPathStr[v] = shortPathStr[u] + "-->" + strconv.Itoa(v)
				}
			}
		}
	}

	fmt.Println("从0 ~ "+strconv.Itoa(count-1)+"节点的最小路径为 ", shortPathStr[count-1])
}
