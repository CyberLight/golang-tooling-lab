##### Tools
* ```https://github.com/google/pprof``` (standard tool)
* https://github.com/tsliwowicz/go-wrk
* https://github.com/brendangregg/FlameGraph/
* https://github.com/uber/go-torch

##### Commands
* go run .\main.go
* go-wrk.exe -d 5 http://localhost:8080/cyberlight@gopher.org
* go-torch -t 5
 * ```or without torch``` go tool pprof --seconds=5 http://localhost:8080/debug/pprof/profile 

* go test -bench .
* go test -bench . -cpuprofile perfs\prof.cpu.0
* go test -bench . -cpuprofile perfs\prof.cpu.1
* go test -bench . -cpuprofile perfs\prof.cpu.2

* go-torch.exe --binaryname .\golang-tooling-lab.test.exe -b .\perfs\prof.cpu.0
* go-torch.exe --binaryname .\golang-tooling-lab.test.exe -b .\perfs\prof.cpu.1
* go-torch.exe --binaryname .\golang-tooling-lab.test.exe -b .\perfs\prof.cpu.2

*  go tool pprof .\golang-tooling-lab.test.exe .\perfs\prof.cpu.1
   * list handler
   
```
go tool pprof .\golang-tooling-lab.test.exe .\perfs\prof.cpu.1
Entering interactive mode (type "help" for commands)
(pprof) list handler
Total: 1.38s
ROUTINE ======================== github.com/cyberlight/golang-tooling-lab.handler in c:/go-workspace/src/github.com/cyberlight/golang-tooli
ng-lab/main.go
      10ms      690ms (flat, cum) 50.00% of Total
         .          .     18:   // match := re.FindAllStringSubmatch(path, -1)
         .          .     19:   // if match != nil {
         .          .     20:   //      fmt.Fprintf(w, "Hello, gopher %s!\n", match[0][1])
         .          .     21:   //      return
         .          .     22:   // }
      10ms       10ms     23:   path := r.URL.Path[1:]
         .          .     24:   if strings.HasSuffix(path, GoperSuffix) {
         .      600ms     25:           fmt.Fprintf(w, "Hello, gopher %s!\n", strings.TrimSuffix(path, GoperSuffix))
         .          .     26:   }
         .       80ms     27:   fmt.Fprintf(w, "Hello, %s!\n", path)
         .          .     28:}
         .          .     29:
         .          .     30:func main() {
         .          .     31:   http.HandleFunc("/", handler)
         .          .     32:   err := http.ListenAndServe(":8080", nil)
```

*  go tool pprof .\golang-tooling-lab.test.exe .\perfs\prof.cpu.2
   * list handler

```
go tool pprof .\golang-tooling-lab.test.exe .\perfs\pro
f.cpu.2
Entering interactive mode (type "help" for commands)
(pprof) list handler
Total: 1.94s
ROUTINE ======================== github.com/cyberlight/golang-tooling-lab.handler in c:/go-workspace/src/github.com/cybe
rlight/golang-tooling-lab/main.go
      10ms      850ms (flat, cum) 43.81% of Total
         .          .     18:   // match := re.FindAllStringSubmatch(path, -1)
         .          .     19:   // if match != nil {
         .          .     20:   //      fmt.Fprintf(w, "Hello, gopher %s!\n", match[0][1])
         .          .     21:   //      return
         .          .     22:   // }
         .      100ms     23:   w.Header().Set("Content-Type", "text/plain")
         .          .     24:   path := r.URL.Path[1:]
         .       10ms     25:   if strings.HasSuffix(path, GoperSuffix) {
      10ms      620ms     26:           fmt.Fprintf(w, "Hello, gopher %s!\n", strings.TrimSuffix(path, GoperSuffix))
         .          .     27:   }
         .      120ms     28:   fmt.Fprintf(w, "Hello, %s!\n", path)
         .          .     29:}
         .          .     30:
         .          .     31:func main() {
         .          .     32:   http.HandleFunc("/", handler)
         .          .     33:   err := http.ListenAndServe(":8080", nil)
```

##### Final bench
```
PS C:\go-workspace\src\github.com\cyberlight\golang-tooling-lab> go test -bench .
BenchmarkHandler-4        500000              3388 ns/op
PASS
ok      github.com/cyberlight/golang-tooling-lab        1.806s
```