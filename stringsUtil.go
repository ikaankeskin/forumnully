package main

import "strings"

func stringToSlice(data string, sep string) []string {
	sl := strings.Split(data, sep)
	result := []string{}
	for _, str := range sl {
		s := strings.TrimSpace(str)
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
func contains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}
