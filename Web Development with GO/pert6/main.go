package main

import "fmt"

func main() {
	student := Student{Name: "Cindy", UTS: 80, UAS: 85, Tugas: 80}

	score, err := student.CalculateFinalScore()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	grade, err := student.CalculateGrade()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	passed, err := student.IsPassed()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Nama Mahasiswa :", student.Name)
	fmt.Println("Nilai Akhir    :", score)
	fmt.Println("Grade          :", grade)
	fmt.Println("Status Lulus   :", passed)
}
