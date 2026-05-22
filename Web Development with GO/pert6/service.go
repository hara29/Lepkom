package main

import "errors"

type Student struct {
	Name  string
	UTS   float64
	UAS   float64
	Tugas float64
}

func (s *Student) CalculateFinalScore() (float64, error) {
	if s.UTS < 0 || s.UTS > 100 {
		return 0, errors.New("nilai UTS harus berada di antara 0 dan 100")
	}

	if s.UAS < 0 || s.UAS > 100 {
		return 0, errors.New("nilai UAS harus berada di antara 0 dan 100")
	}

	if s.Tugas < 0 || s.Tugas > 100 {
		return 0, errors.New("nilai Tugas harus berada di antara 0 dan 100")
	}

	finalScore := (s.UTS * 0.30) + (s.UAS * 0.40) + (s.Tugas * 0.30)

	return finalScore, nil
}

func (s *Student) CalculateGrade() (string, error) {
	finalScore, err := s.CalculateFinalScore()
	if err != nil {
		return "", err
	}

	switch {
	case finalScore >= 85:
		return "A", nil
	case finalScore >= 70:
		return "B", nil
	case finalScore >= 60:
		return "C", nil
	case finalScore >= 50:
		return "D", nil
	default:
		return "E", nil
	}
}

func (s *Student) IsPassed() (bool, error) {
	grade, err := s.CalculateGrade()
	if err != nil {
		return false, err
	}

	return grade == "A" || grade == "B" || grade == "C", nil
}
