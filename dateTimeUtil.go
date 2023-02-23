package main

import "time"

//
// Get current time in milliseconds
func getCurrentMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

// Convert milliseconds into this format "02-Jan-2006 15:04:05"
func formatMilli(date int) string {
	time := time.Unix(int64(date), 0)
	return time.Format("2023-02-25 21:09:56")
}
