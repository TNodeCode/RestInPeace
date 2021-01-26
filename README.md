# Rest In Peace

### A CPU Benchmark tool build with Go

#### Howto

```bash
# Calculating all prime numbers up to 1 Million with 8 parallel goroutines

$ time ./rip -t 8 -n 1000000

Number of Threads:  8
Maximum Number:  1000000
---
 78498 prime numbers found

real    1m22,570s
user    0m0,031s
sys     0m0,031s
```

#### Executables

You find the executable files in the /dist folder

|         | AMD64 | 386 | ARM64 | ARM |
| ------- | ----- | --- | ----- | --- |
| Windows | x     | x   |       |     |
| Linux   | x     | x   | x     | x   |
| MacOS   | x     | x   | x     | x   |
