package models

type SearchInput struct {
	Name string
}

// この米印は何だ？キャスト？
// っぽいwwww
// ってことは上で定義したタイミングでセッター的なの作れないのかな？
// とりあえず予測変換で"set"や"didset"的なのは引っかからない
// Goの拡張入れたら勝手に詰めたりするようになったからこの書き方が正規っぽい？（＊の横詰めるやつ）
// TODO: 今度調べるリマインダー
func (input *SearchInput) setInput(name string) {
	input.Name = name
}
