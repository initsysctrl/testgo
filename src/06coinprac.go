package main

import "fmt"

/**
 *@Author: yepeng
 *@Date: 2021/8/30 1:55 上午
 *@Name: 06coinprac
 *@Desc:分金币游戏，并计算剩余,
 * 规则，包含a加10，包含b+20，包含c+30，包含d+40，其余不变
 */
var coins = 10
var users []string
var dis = make(map[string]int)

func disfunc() {
	users := append(users, "abc")
	users = append(users, "bbb")
	users = append(users, "ccc")
	users = append(users, "defd")
	users = append(users, "oooo")
	fmt.Println(users)
	for _, u := range users {
		for _, c := range u {
			if coins <= 0 {
				break
			}
			switch c {
			case 'a':
				coins -= 1
				dis[u] += 1
			case 'b':
				coins -= 2
				dis[u] += 2
			case 'c':
				coins -= 3
				dis[u] += 3
			case 'd':
				coins -= 4
				dis[u] += 4
			default:
				dis[u] += 0
			}
		}
	}

}
func main() {
	disfunc()
	fmt.Println(dis)

}
