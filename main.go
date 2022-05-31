package main

import "TikTok/db"

func main() {
	err := db.InitDb()
	if err != nil {
		panic(err)
	}
}
