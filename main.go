package main

import "fmt"

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
	PB = TB * 1024
)

type Bytesz float64

func ( b Bytesz ) String() string {
	switch {
	case b >= PB:
	return "PetaByte Big!!"
	case b >= TB:
	return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
	return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
	return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
	return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%dB", b)
}

func main(){
	fmt.Println(Bytesz(2048)) // Outputs: 2.00KB
	fmt.Println(Bytesz(3292528.64)) // Outputs: 3.14MB
}

