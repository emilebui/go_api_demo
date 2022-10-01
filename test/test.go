package main

import "go_api/service/scan"

func main() {
	path, _ := scan.DownloadRepo("https://github.com/emilebui/Minigames-on-Processing.git", "test")
	_ = scan.ScanRepo(path, "test")
}
