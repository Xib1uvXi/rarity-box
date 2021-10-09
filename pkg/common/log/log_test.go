package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"go.uber.org/zap"
)

type Tst struct {
	A      string    `json:"a"`
	Number uint      `json:"number"`
	Time   time.Time `json:"time"`
}

func TestInitLogger(t *testing.T) {
	Logger.Debug("test init debug log", zap.String("a", "asss"))

	tst := &Tst{A: "fff", Number: 44}
	Logger.Debug("test debug log struct", zap.Any("tst", tst))

	tst1 := &Tst{A: "fff", Number: 44, Time: time.Now()}
	Logger.Debug("test debug log struct", zap.Any("tst", tst1))

	Logger.Debug("ddddddd")

	Logger.Info("iiiiiiiiii")

	Logger.Warn("wwwwwwwwwww")

	Logger.Error("eeeeeee")

}



func Test_newLogOptions(t *testing.T) {
	assert.True(t, 1 == len(newLogOptions()))
}


func TestTrace(t *testing.T) {
	Logger.Info("aaaaaaa")
}

func BenchmarkLogCaller(b *testing.B) {
	//  119269	      8697 ns/op
	for i := 0; i < b.N; i++ {
		Logger.Info("aaaaaaa")
	}
}

func BenchmarkNoCaller(b *testing.B) {
	// 219262	      5611 ns/op
	for i := 0; i < b.N; i++ {
		Logger.Info("aaaaaa")
	}
}
