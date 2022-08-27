// STEP04: gachaパッケージにバージョンを付けよう

package main

import (
	"fmt"
	"os"
	"strconv"
	// TODO: インポートパスを公開したものに変更する
	"github.com/kose-yusuke/gogacha/gacha"
)

func main() {
	p := gacha.NewPlayer(10, 100)

	n := inputN(p)
	results, summary := gacha.DrawN(p, n)

	saveResults(results)
	saveSummary(summary)
}

func inputN(p *gacha.Player) int {

	max := p.DrawableNum()
	fmt.Printf("ガチャを引く回数を入力してください（最大:%d回）\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		if n > 0 && n <= max {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", max)
	}

	return n
}

func initialTickets() int {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "ガチャチケットの枚数を入力してください")
		os.Exit(1)
	}

	// TODO: プログラム引数の1つめを取得し
	// strconv.Atoi関数でint型に変換する
	// 第1戻り値は変数num、第2戻り値は変数errに代入する
	num, err :=strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return num
}

func saveResults(results []*gacha.Card) {
	f, err := os.Create("results.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for _, result := range results {
		fmt.Fprintln(f, result)
	}
}

func saveSummary(summary map[gacha.Rarity]int) {
	f, err := os.Create("summary.txt")
	if err != nil {
		fmt.Println(os.Stderr,err)
		os.Exit(1)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	for rarity, count := range summary {
		fmt.Fprintf(f, "%s %d\n", rarity.String(), count)
	}
}