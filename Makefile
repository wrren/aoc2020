INTERNAL_SOURCES=$(find internal/ -name "*.go")

all: day1 day2 day3 day4 day5 day6 day7 day8 day9 day10

day1: $(INTERNAL_SOURCES) cmd/day1/main.go
	go build -o build/day1 cmd/day1/main.go

day2: $(INTERNAL_SOURCES) cmd/day2/main.go
	go build -o build/day2 cmd/day2/main.go

day3: $(INTERNAL_SOURCES) cmd/day3/main.go
	go build -o build/day3 cmd/day3/main.go

day4: $(INTERNAL_SOURCES) cmd/day4/main.go
	go build -o build/day4 cmd/day4/main.go

day5: $(INTERNAL_SOURCES) cmd/day5/main.go
	go build -o build/day5 cmd/day5/main.go

day6: $(INTERNAL_SOURCES) cmd/day6/main.go
	go build -o build/day6 cmd/day6/main.go

day7: $(INTERNAL_SOURCES) cmd/day7/main.go
	go build -o build/day7 cmd/day7/main.go

day8: $(INTERNAL_SOURCES) cmd/day8/main.go
	go build -o build/day8 cmd/day8/main.go

day9: $(INTERNAL_SOURCES) cmd/day9/main.go
	go build -o build/day9 cmd/day9/main.go

day10: $(INTERNAL_SOURCES) cmd/day10/main.go
	go build -o build/day10 cmd/day10/main.go