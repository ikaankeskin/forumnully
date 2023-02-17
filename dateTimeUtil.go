package main

import "time"

func getCurrentMilli() int64 {
	return time.Now().UnixMilli()
}
func formatMilli(date int) string {
	t := time.Unix(int64(date), 0)
	return t.Format("02-Jan-2006 15:04:05")
}
