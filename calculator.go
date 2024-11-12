package main

import (
	"fmt"
	"log"
	"os"
)

// Fungsi untuk penjumlahan
func Add(a, b float64) float64 {
	return a + b
}

// Fungsi untuk pengurangan
func Subtract(a, b float64) float64 {
	return a - b
}

// Fungsi untuk perkalian
func Multiply(a, b float64) float64 {
	return a * b
}

// Fungsi untuk pembagian
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	// Menampilkan menu
	fmt.Println("Simple Calculator")
	fmt.Println("------------------")
	fmt.Println("Select an operation:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")

	// Membaca pilihan operasi
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		log.Fatal(err)
	}

	// Membaca angka pertama
	fmt.Print("Enter the first number: ")
	var num1 float64
	_, err = fmt.Scanln(&num1)
	if err != nil {
		log.Fatal(err)
	}

	// Membaca angka kedua
	fmt.Print("Enter the second number: ")
	var num2 float64
	_, err = fmt.Scanln(&num2)
	if err != nil {
		log.Fatal(err)
	}

	// Proses berdasarkan pilihan
	var result float64
	var errDiv error

	switch choice {
	case 1:
		result = Add(num1, num2)
		fmt.Printf("Result: %.2f + %.2f = %.2f\n", num1, num2, result)
	case 2:
		result = Subtract(num1, num2)
		fmt.Printf("Result: %.2f - %.2f = %.2f\n", num1, num2, result)
	case 3:
		result = Multiply(num1, num2)
		fmt.Printf("Result: %.2f * %.2f = %.2f\n", num1, num2, result)
	case 4:
		result, errDiv = Divide(num1, num2)
		if errDiv != nil {
			fmt.Println("Error:", errDiv)
			os.Exit(1)
		}
		fmt.Printf("Result: %.2f / %.2f = %.2f\n", num1, num2, result)
	default:
		fmt.Println("Invalid choice. Please select a valid operation.")
	}
}
