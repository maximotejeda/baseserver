package main

import (
	"os"
	"testing"
)

func TestPreRun(t *testing.T) {
	os.Setenv("ADDR", "")
	os.Setenv("PORT", "1010")
	t.Run("test index exists", func(t *testing.T) {
		val := preRun()
		if _, ok := val["address"]; !ok {

			t.Fatal("address not found on map")
		}
		if _, ok := val["port"]; !ok {

			t.Fatal("PORT not found on map")
		}
	})
}
