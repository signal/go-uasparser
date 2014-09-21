profile:
	go test -v -bench=. -run='Benchmark*' -cpuprofile cpu.pprof
	go tool pprof --text ./go-uasparser.test cpu.pprof
