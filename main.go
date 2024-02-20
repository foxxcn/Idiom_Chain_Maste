package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"math/rand"

	"github.com/mozillazg/go-pinyin"
)

type Idiom struct {
	Text   string
	Pinyin string
}

// 加载成语库并处理拼音
func loadIdioms(filePath string) ([]Idiom, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var idioms []Idiom
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		pinyinArgs := pinyin.NewArgs()
		py := pinyin.Pinyin(text, pinyinArgs)
		// 修正：确保py数组不为空且py[0]也不为空
		if len(py) > 0 && len(py[0]) > 0 {
			flatPinyin := strings.Join(py[0], "")
			idioms = append(idioms, Idiom{
				Text:   text,
				Pinyin: flatPinyin,
			})
		}
	}
	return idioms, scanner.Err()
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数生成器
	idioms, err := loadIdioms("idioms.txt")
	if err != nil {
		fmt.Printf("Failed to load idioms: %v\n", err)
		return
	}

	fmt.Println("游戏开始，请输入成语：")
	scanner := bufio.NewScanner(os.Stdin)

	usedIdioms := make(map[string]bool) // 用于跟踪已使用的成语
	var lastIdiom Idiom
	score := 0.0
	startTime := time.Now()
	rounds := 0

	for rounds < 10 {
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		if input == "quit" {
			break
		}

		if input == "tip" {
			// 提示逻辑（简化实现）
			for _, idiom := range idioms {
				if !usedIdioms[idiom.Text] && (lastIdiom.Text == "" || strings.HasPrefix(idiom.Text, string(lastIdiom.Text[len(lastIdiom.Text)-1]))) {
					fmt.Printf("提示：%s\n", idiom.Text)
					score -= 0.08
					break
				}
			}
			continue
		}

		found := false
		for _, idiom := range idioms {
			if idiom.Text == input {
				if lastIdiom.Text == "" || strings.HasPrefix(input, string(lastIdiom.Text[len(lastIdiom.Text)-1])) {
					if usedIdioms[input] {
						fmt.Println("成语重复，扣0.1分。")
						score -= 0.1
					} else {
						fmt.Println("接龙成功，请继续：")
						score += 1
						rounds++
						usedIdioms[input] = true
						lastIdiom = idiom
					}
					found = true
					break
				} else {
					fmt.Println("接龙失败，请重新输入：")
					score -= 0.3
					found = true
					break
				}
			}
		}

		if !found {
			fmt.Println("输入的成语不在词典中，请重新输入：")
			score -= 0.3
		}

		if score <= 0 {
			fmt.Println("分数为0，游戏结束。")
			break
		}
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime).Minutes()
	finalScore := (score * 6) + (float64(rounds) * 3) // 根据规则计算最终得分

	fmt.Printf("游戏结束。总分：%.2f，游戏时长：%.2f分钟。\n", finalScore, duration)
}

