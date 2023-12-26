package main

import (
	"github.com/oriental1212/scriptTools/fileTools"
)

func main() {
	// 数据库
	//db := sqlTools.ConnectDevMySQL()
	//var irsCmsSync sqlTools.IrsCmsSync
	//result := db.Where("id < ?", id).Limit(1000).Delete(&irsCmsSync)
	//for result.RowsAffected == 1000 {
	//	result = db.Where("id < ?", id).Limit(1000).Delete(&irsCmsSync)
	//	nowTime = time.Now()
	//	fmt.Printf("当前时间:%02d:%02d:%02d ", nowTime.Hour(), nowTime.Minute(), nowTime.Second())
	//	fmt.Println(",删除数量:", result.RowsAffected)
	//}

	// 文件去重
	fileTools.RegularMatchFile()

}
