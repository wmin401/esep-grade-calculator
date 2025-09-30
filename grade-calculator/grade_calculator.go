package esepunittests

type GradeCalculator struct {
	// updated it to a single list
	gradeList []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		// updated variable name
		gradeList: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {

	numerical := gc.calculateNumericalGrade()
	//updated it to using switch
	switch {
	case numerical >= 90:
		return "A"
	case numerical >= 80:
		return "B"
	case numerical >= 70:
		return "C"
	case numerical >= 60:
		return "D"
	default:
		return "F"
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	// updated variable name
	gc.gradeList = append(gc.gradeList, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	var sumA, cntA, sumE, cntE, sumS, cntS int
	// updated variable name
	for _, g := range gc.gradeList {
		switch g.Type {
		case Assignment:
			sumA += g.Grade
			cntA++
		case Exam:
			sumE += g.Grade
			cntE++
		case Essay:
			sumS += g.Grade
			cntS++
		}
	}
	assignmentAvg := avgInt(sumA, cntA)
	examAvg := avgInt(sumE, cntE)
	essayAvg := avgInt(sumS, cntS)

	weighted := float64(assignmentAvg)*0.50 +
		float64(examAvg)*0.35 +
		float64(essayAvg)*0.15

	return int(weighted)
}

// add function avgInt
func avgInt(sum, cnt int) int {
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}

func computeAverage(grades []Grade) int {
	//bugfix 3
	if len(grades) == 0 {
		return 0
	}
	sum := 0
	//bugfix 2
	for _, grade := range grades {
		sum += grade.Grade
	}

	return sum / len(grades)
}

// add function to get Pass or Fail
func (gc *GradeCalculator) GetPassFail() string {
	numeric := gc.calculateNumericalGrade()
	if numeric >= 70 {
		return "Pass"
	}
	return "Fail"
}
