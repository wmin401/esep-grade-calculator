package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	// Adjust to weighted average < 60
	gradeCalculator.AddGrade("open source assignment", 10, Assignment)
	gradeCalculator.AddGrade("exam 1", 20, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 30, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

//(Part 4) add function 1
//func TestGradeTypeString(t *testing.T) {
//	if Assignment.String() != "assignment" {
//		t.Fatalf("Assignment.String()=%q", Assignment.String())
//	}
//	if Exam.String() != "exam" {
//		t.Fatalf("Exam.String()=%q", Exam.String())
//	}
//	if Essay.String() != "essay" {
//		t.Fatalf("Essay.String()=%q", Essay.String())
//	}
//}

//func TestGetGradeD(t *testing.T) {
//	gc := NewGradeCalculator()
//	gc.AddGrade("assign", 60, Assignment)
//	gc.AddGrade("exam1", 60, Exam)
//	gc.AddGrade("essay1", 60, Essay)
//
//	if got := gc.GetFinalGrade(); got != "D" {
//		t.Fatalf("GetFinalGrade()=%s; want D", got)
//	}
//}
//
//func TestGetGradeReallyF(t *testing.T) {
//	gc := NewGradeCalculator()
//	gc.AddGrade("assign", 0, Assignment)
//	gc.AddGrade("exam1", 0, Exam)
//	gc.AddGrade("essay1", 0, Essay)
//
//	if got := gc.GetFinalGrade(); got != "F" {
//		t.Fatalf("GetFinalGrade()=%s; want F", got)
//	}
//}
//
//func TestComputeAverageEmpty(t *testing.T) {
//	if got := computeAverage(nil); got != 0 {
//		t.Fatalf("computeAverage(nil)=%d; want 0", got)
//	}
//	if got := computeAverage([]Grade{}); got != 0 {
//		t.Fatalf("computeAverage(empty)=%d; want 0", got)
//	}
//}
//
//func TestGetGradeC(t *testing.T) {
//	gc := NewGradeCalculator()
//	gc.AddGrade("assign", 70, Assignment)
//	gc.AddGrade("exam1", 70, Exam)
//	gc.AddGrade("essay1", 70, Essay)
//
//	if got := gc.GetFinalGrade(); got != "C" {
//		t.Fatalf("GetFinalGrade()=%s; want C", got)
//	}
//}
