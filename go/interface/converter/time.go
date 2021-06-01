package converter

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// NOTE: UTCからInメソッドを使ってJSTに変えようとしたけど、
// timestamppb.Newメソッドを使うとUTCに書き換えられるので、
// UTCとJSTの時差9時間を直接プラスする。要リファクタ
func FromUTCtoJSTinProto(date time.Time) *timestamppb.Timestamp {
	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// fmt.Println(jst.String())
	// fmt.Println(time.Now().In(jst))
	// jjjj := date.In(jst)
	// fmt.Println(jjjj)
	// fmt.Println(timestamppb.New(date))
	// fmt.Println(timestamppb.New(jjjj))

	return timestamppb.New(date.Add(9 * time.Hour))
}
