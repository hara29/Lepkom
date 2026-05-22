package main

import "testing"

func TestAllLogic_Coverage(t *testing.T) {
	scenarios := []struct {
		name           string
		student        Student
		expectedScore  float64
		expectedGrade  string
		expectedPassed bool
	}{
		{
			name:           "Grade A - Lulus",
			student:        Student{Name: "Cindy", UTS: 90, UAS: 90, Tugas: 90},
			expectedScore:  90,
			expectedGrade:  "A",
			expectedPassed: true,
		},
		{
			name:           "Grade B - Lulus",
			student:        Student{Name: "Budi", UTS: 75, UAS: 75, Tugas: 75},
			expectedScore:  75,
			expectedGrade:  "B",
			expectedPassed: true,
		},
		{
			name:           "Grade C - Lulus",
			student:        Student{Name: "Ani", UTS: 65, UAS: 65, Tugas: 65},
			expectedScore:  65,
			expectedGrade:  "C",
			expectedPassed: true,
		},
		{
			name:           "Grade D - Tidak Lulus",
			student:        Student{Name: "Rina", UTS: 55, UAS: 55, Tugas: 55},
			expectedScore:  55,
			expectedGrade:  "D",
			expectedPassed: false,
		},
		{
			name:           "Grade E - Tidak Lulus",
			student:        Student{Name: "Doni", UTS: 40, UAS: 40, Tugas: 40},
			expectedScore:  40,
			expectedGrade:  "E",
			expectedPassed: false,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			score, err := scenario.student.CalculateFinalScore()
			if err != nil {
				t.Errorf("tidak seharusnya error: %v", err)
			}

			if score != scenario.expectedScore {
				t.Errorf("expected score %.2f but got %.2f", scenario.expectedScore, score)
			}

			grade, err := scenario.student.CalculateGrade()
			if err != nil {
				t.Errorf("tidak seharusnya error: %v", err)
			}

			if grade != scenario.expectedGrade {
				t.Errorf("expected grade %s but got %s", scenario.expectedGrade, grade)
			}

			passed, err := scenario.student.IsPassed()
			if err != nil {
				t.Errorf("tidak seharusnya error: %v", err)
			}

			if passed != scenario.expectedPassed {
				t.Errorf("expected passed %v but got %v", scenario.expectedPassed, passed)
			}
		})
	}

	invalidStudents := []Student{
		{Name: "Invalid UTS Minus", UTS: -5, UAS: 80, Tugas: 80},
		{Name: "Invalid UTS More", UTS: 101, UAS: 80, Tugas: 80},
		{Name: "Invalid UAS Minus", UTS: 80, UAS: -1, Tugas: 80},
		{Name: "Invalid UAS More", UTS: 80, UAS: 101, Tugas: 80},
		{Name: "Invalid Tugas Minus", UTS: 80, UAS: 80, Tugas: -1},
		{Name: "Invalid Tugas More", UTS: 80, UAS: 80, Tugas: 101},
	}

	for _, student := range invalidStudents {
		t.Run(student.Name, func(t *testing.T) {
			_, err := student.CalculateFinalScore()
			if err == nil {
				t.Error("expected error untuk nilai tidak valid")
			}

			_, err = student.CalculateGrade()
			if err == nil {
				t.Error("expected error pada CalculateGrade")
			}

			_, err = student.IsPassed()
			if err == nil {
				t.Error("expected error pada IsPassed")
			}
		})
	}
}
