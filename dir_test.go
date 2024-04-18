package main

import (
	"os"
	"path"
	"strings"
	"testing"
)

func TestDir(t *testing.T) {
	ListDir("/home/eric/Downloads/2024", "xlsx")
}

