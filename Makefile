# Warning: profiling is broken on OSX
cpu.pprof:
	go test -v -bench=. -run='Benchmark*' -cpuprofile cpu.pprof

profile-txt: cpu.pprof
	go tool pprof --text ./go-uasparser.test cpu.pprof

# requires graphviz and ghostscript
profile-cg: cpu.pprof
	go tool pprof --pdf ./go-uasparser.test cpu.pprof > callgraph.pdf
	open callgraph.pdf

