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

// (Part 4) add function gradeType.String
func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatalf("Assignment.String()=%q", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Fatalf("Exam.String()=%q", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Fatalf("Essay.String()=%q", Essay.String())
	}
}

// add function getGradeD
func TestGetGradeD(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assign", 60, Assignment)
	gc.AddGrade("exam1", 60, Exam)
	gc.AddGrade("essay1", 60, Essay)

	if got := gc.GetFinalGrade(); got != "D" {
		t.Fatalf("GetFinalGrade()=%s; want D", got)
	}
}

// add helper function to test F
func TestGetGradeReallyF(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assign", 0, Assignment)
	gc.AddGrade("exam1", 0, Exam)
	gc.AddGrade("essay1", 0, Essay)

	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("GetFinalGrade()=%s; want F", got)
	}
}

// add function to check if computeAverage is zero
func TestComputeAverageEmpty(t *testing.T) {
	if got := computeAverage(nil); got != 0 {
		t.Fatalf("computeAverage(nil)=%d; want 0", got)
	}
	if got := computeAverage([]Grade{}); got != 0 {
		t.Fatalf("computeAverage(empty)=%d; want 0", got)
	}
}

// add function getGradeC
func TestGetGradeC(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assign", 70, Assignment)
	gc.AddGrade("exam1", 70, Exam)
	gc.AddGrade("essay1", 70, Essay)

	if got := gc.GetFinalGrade(); got != "C" {
		t.Fatalf("GetFinalGrade()=%s; want C", got)
	}
}

// add function to test only assignments
func TestGetGradeOnlyAssignments(t *testing.T) {
	gc := NewGradeCalculator()

	// Only assignments -> exams, essays count=0
	gc.AddGrade("assign1", 100, Assignment)

	// Weighted: 100 * 0.5 = 50 -> "F"
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("GetFinalGrade()=%s; want F", got)
	}
}

// add function to test average-non-empty
func TestComputeAverageNonEmpty(t *testing.T) {
	in := []Grade{
		{Name: "a", Grade: 80, Type: Assignment},
		{Name: "b", Grade: 100, Type: Assignment},
	}
	if got := computeAverage(in); got != 90 {
		t.Fatalf("computeAverage(non-empty)=%d; want 90", got)
	}
}

// add function to test pass/fail above C
func TestPassFail_PassAtCAndAbove(t *testing.T) {
	gc := NewGradeCalculator()

	// C (70) -> Pass
	gc.AddGrade("assign", 70, Assignment)
	gc.AddGrade("exam", 70, Exam)
	gc.AddGrade("essay", 70, Essay)
	if got := gc.GetPassFail(); got != "Pass" {
		t.Fatalf("GetPassFail()=%s; want Pass (C threshold)", got)
	}

	// B (80) -> Pass
	gc = NewGradeCalculator()
	gc.AddGrade("assign", 80, Assignment)
	gc.AddGrade("exam", 80, Exam)
	gc.AddGrade("essay", 80, Essay)
	if got := gc.GetPassFail(); got != "Pass" {
		t.Fatalf("GetPassFail()=%s; want Pass (B)", got)
	}

	// A (90) -> Pass
	gc = NewGradeCalculator()
	gc.AddGrade("assign", 90, Assignment)
	gc.AddGrade("exam", 90, Exam)
	gc.AddGrade("essay", 90, Essay)
	if got := gc.GetPassFail(); got != "Pass" {
		t.Fatalf("GetPassFail()=%s; want Pass (A)", got)
	}
}

// add function to test pass/fail below C
func TestPassFail_FailBelowC(t *testing.T) {
	// D (60) -> Fail
	gc := NewGradeCalculator()
	gc.AddGrade("assign", 60, Assignment)
	gc.AddGrade("exam", 60, Exam)
	gc.AddGrade("essay", 60, Essay)
	if got := gc.GetPassFail(); got != "Fail" {
		t.Fatalf("GetPassFail()=%s; want Fail (D)", got)
	}

	// F (<60) -> Fail
	gc = NewGradeCalculator()
	gc.AddGrade("assign", 10, Assignment)
	gc.AddGrade("exam", 20, Exam)
	gc.AddGrade("essay", 30, Essay)
	if got := gc.GetPassFail(); got != "Fail" {
		t.Fatalf("GetPassFail()=%s; want Fail (F)", got)
	}
}
