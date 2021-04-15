package slide_window

import (
	"math/rand"
	"testing"
	"time"
)

func TestWindow(t *testing.T) {
	win, err := NewWindow(10, 10*time.Second)
	if err != nil {
		t.Errorf("init window error %v\n", err)
	}
	for i := 0; i < 50; i++ {
		if rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10)%2 == 1 {
			win.Roll(IncrOutSuccess)
			t.Logf("request success %d\n", i)
		} else {
			win.Roll(IncrOutFail)
			t.Logf("request fail %d\n", i)
		}
		time.Sleep(time.Millisecond * 10)
	}
	t.Logf("window success:%d - fail:%d \n", win.Sum("Success"), win.Sum("Fail"))

}
