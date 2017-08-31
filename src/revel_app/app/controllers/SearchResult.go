package controllers

// importは()でやるっぽい。きもいな。{}と間違えないように気をつけないといけない
import (
	dbconnector "revel_app/app/dbconnector"
	models "revel_app/app/models"

	"github.com/revel/revel"
)

func (c App) SearchResult(name string) revel.Result {
	// 検索結果を格納　search.htmlに記載した"value"パラメーターの所？
	inputDatas := new(models.SearchInput)
	inputDatas.Name = name

	// 社員名でdbを検索
	// Goは戻りに二つセットできるからこれで結果もリザルトも一発で引っ張ってこれる。返す側も片方値ありのもう片方nilとか舐めたことできる
	results, error := dbconnector.Search(name)
	// if error != nilじゃなくてこれじゃダメかな・・・
	if error != nil {
		c.RenderError(error)
	}

	return c.Render(inputDatas, results)
}
