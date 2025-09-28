# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

## Part 1: Requirements Clarification

1. **Empty lists:**  
   If a student has no grades for a particular type such as zero essays, should the average for that type be treated as `0`, or should it be excluded from the weighted average calculation? Please specify the expected behavior for each category (assignments, exams, essays).

2. **Boundaries & rounding:**  
   Do grade thresholds include the boundary values (is `90.0` considered an `A` and `80.0` a `B`)? Should any rounding be applied (such as rounding to one decimal place) before converting a numeric grade into a letter grade?

3. **Weight consistency:**  
   If the provided weights do not sum exactly to 100% due to data or configuration errors, should they be normalized to sum to 100%, or should the calculator return an error? Additionally, can weights vary for each class, or should the defaults (Assignments 50%, Exams 35%, Essays 15%) always be enforced?
