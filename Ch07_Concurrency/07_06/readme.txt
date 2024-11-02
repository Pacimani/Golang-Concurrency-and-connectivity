running benchmark only for profiling
  go test -bench . -run None

 
saving the profiling in a text file called cpu.prof
go test -bench . -run None -cpuprofile=cpu.prof

viewing the profile in the browser
go tool pprof -http=:8080 cpu.prof