# Structs Vs Pointers

This repository tries to compare the performance of structs and pointers across various scenarios.

**Note** - This test is run against Go version `1.21.3` with 1 cpu.

## Benchmarking Results

- **[No Stack Without Return](#no-stack-without-return)** - This is the most common scenario where the object, either struct or its pointer is used within a single method with the object being withing the function itself i.e. no return of object.
- **[No Stack With Return](#no-stack-with-return)** - This is the most common scenario where the object, either struct or its pointer is used within a single method along with returning the value.
- **[Basic Stack](#basic-stack)** - This is the scenario where the object, either struct or its pointer is passed down multiple functions.
- **[Arrays](#arrays)** - This is the scenario where the object, either struct or its pointer is part of an array and is passed down multiple functions.
- **[Maps](#maps)** - This is the scenario where the object, either struct or its pointer is part of a map and is passed down multiple functions.

### No Stack Without Return

This is the most common scenario where the object, either struct or its pointer is used within a single method with the object being withing the function itself i.e. no return of object.

```text
/Users/shubham.sinha/sdk/go1.21.3/bin/go test -v ./... -bench . -run ^$ -benchmem -cpu 1
goos: darwin
goarch: amd64
pkg: github.com/sinhashubham95/structs-vs-pointers/nostackwithoutreturn
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkS10fByValue
BenchmarkS10fByValue   	1000000000	         0.3053 ns/op	       0 B/op	       0 allocs/op
BenchmarkS10fByPointer
BenchmarkS10fByPointer 	1000000000	         0.2872 ns/op	       0 B/op	       0 allocs/op
BenchmarkS2fByValue
BenchmarkS2fByValue    	1000000000	         0.3149 ns/op	       0 B/op	       0 allocs/op
BenchmarkS2fByPointer
BenchmarkS2fByPointer  	1000000000	         0.2905 ns/op	       0 B/op	       0 allocs/op
BenchmarkS5fByValue
BenchmarkS5fByValue    	1000000000	         0.3412 ns/op	       0 B/op	       0 allocs/op
BenchmarkS5fByPointer
BenchmarkS5fByPointer  	1000000000	         0.2850 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/sinhashubham95/structs-vs-pointers/nostackwithoutreturn	       2.367s
```

### No Stack With Return

This is the most common scenario where the object, either struct or its pointer is used within a single method along with returning the value.

```text
/Users/shubham.sinha/sdk/go1.21.3/bin/go test -v ./... -bench . -run ^$ -benchmem -cpu 1
goos: darwin
goarch: amd64
pkg: github.com/sinhashubham95/structs-vs-pointers/nostackwithreturn
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkS10fByValue
BenchmarkS10fByValue   	134305704	         9.561 ns/op	       0 B/op	       0 allocs/op
BenchmarkS10fByPointer
BenchmarkS10fByPointer 	1000000000	         0.3030 ns/op	       0 B/op	       0 allocs/op
BenchmarkS2fByValue
BenchmarkS2fByValue    	1000000000	         0.2883 ns/op	       0 B/op	       0 allocs/op
BenchmarkS2fByPointer
BenchmarkS2fByPointer  	1000000000	         0.2863 ns/op	       0 B/op	       0 allocs/op
BenchmarkS5fByValue
BenchmarkS5fByValue    	237990206	         5.632 ns/op	       0 B/op	       0 allocs/op
BenchmarkS5fByPointer
BenchmarkS5fByPointer  	1000000000	         0.2560 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/sinhashubham95/structs-vs-pointers/nostackwithreturn	               5.625s
```

### Basic Stack

This is the scenario where the object, either struct or its pointer is passed down multiple functions.

```text
/Users/shubham.sinha/sdk/go1.21.3/bin/go test -v ./... -bench . -run ^$ -benchmem -cpu 1
goos: darwin
goarch: amd64
pkg: github.com/sinhashubham95/structs-vs-pointers/stack
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkS10fByValue
BenchmarkS10fByValue   	63777744	        18.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkS10fByPointer
BenchmarkS10fByPointer 	1000000000	        0.2970 ns/op	       0 B/op	       0 allocs/op
BenchmarkS2fByValue
BenchmarkS2fByValue    	1000000000	        0.2998 ns/op	       0 B/op	       0 allocs/op
BenchmarkS2fByPointer
BenchmarkS2fByPointer  	1000000000	        0.2845 ns/op	       0 B/op	       0 allocs/op
BenchmarkS5fByValue
BenchmarkS5fByValue    	89151668	        11.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkS5fByPointer
BenchmarkS5fByPointer  	1000000000	        0.2558 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/sinhashubham95/structs-vs-pointers/stack	       5.856s
```

### Arrays

This is the scenario where the object, either struct or its pointer is part of an array and is passed down multiple functions.

```text
/Users/shubham.sinha/sdk/go1.21.3/bin/go test -v ./... -bench . -run ^$ -benchmem -cpu 1
goos: darwin
goarch: amd64
pkg: github.com/sinhashubham95/structs-vs-pointers/arrays
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkS10fByValue
BenchmarkS10fByValue   	54680168	        21.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkS10fByPointer
BenchmarkS10fByPointer 	40205076	        32.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkS2fByValue
BenchmarkS2fByValue    	1000000000	        0.2780 ns/op	       0 B/op	       0 allocs/op
BenchmarkS2fByPointer
BenchmarkS2fByPointer  	157222998	        6.792 ns/op	       0 B/op	       0 allocs/op
BenchmarkS5fByValue
BenchmarkS5fByValue    	1000000000	        0.3046 ns/op	       0 B/op	       0 allocs/op
BenchmarkS5fByPointer
BenchmarkS5fByPointer  	64820839	        17.53 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/sinhashubham95/structs-vs-pointers/arrays           7.374s
```

### Maps

This is the scenario where the object, either struct or its pointer is part of a map and is passed down multiple functions.

```text
/Users/shubham.sinha/sdk/go1.21.3/bin/go test -v ./... -bench . -run ^$ -benchmem -cpu 1
goos: darwin
goarch: amd64
pkg: github.com/sinhashubham95/structs-vs-pointers/maps
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkS10fByValue
BenchmarkS10fByValue   	 1296307	       948.6 ns/op	     320 B/op	       2 allocs/op
BenchmarkS10fByPointer
BenchmarkS10fByPointer 	 3258984	       365.7 ns/op	     320 B/op	       2 allocs/op
BenchmarkS2fByValue
BenchmarkS2fByValue    	 6066308	       207.0 ns/op	     0 B/op	       0 allocs/op
BenchmarkS2fByPointer
BenchmarkS2fByPointer  	 4198629	       294.0 ns/op	     64 B/op	       2 allocs/op
BenchmarkS5fByValue
BenchmarkS5fByValue    	 4554530	       288.8 ns/op	     0 B/op	       0 allocs/op
BenchmarkS5fByPointer
BenchmarkS5fByPointer  	 2706584	       394.4 ns/op	     160 B/op	       2 allocs/op
PASS
ok  	github.com/sinhashubham95/structs-vs-pointers/maps	     10.182s
```
