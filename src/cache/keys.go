package cache

import (
	"fmt"
	"strconv"
)

const (
	DailyRankKey = "rank:daily"
)

func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}
