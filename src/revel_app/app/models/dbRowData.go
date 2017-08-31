// このパッケージってのがひとかたまり的な扱いで、こう書くことでこいつは一括で他が"models"をインポートすれば使えちゃうって寸法
// その代わり同じパッケージ内で同名のメソッドは使えなくなる
// サーチモデルなど作るときはちゃんと”何”のサーチかってのをメソッド名につけてやらんと全部サーチとかになっちゃうかも。。。。？
package models

// struct=const?
type DBRowData struct {
	MemberType string
	Name       string
}
