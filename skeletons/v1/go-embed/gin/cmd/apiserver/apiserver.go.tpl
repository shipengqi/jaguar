package main

import "{{ .App.ModuleName }}/internal"

func main() {
	internal.NewApp().Run()
}
