// Brief one line overview of package
package evenFibonacciNumbers

func Problem2(l uint) uint {

	if l < 2 {
		return 0
	}
	var (
		efs uint = 0
		cf  uint = 1 // last fib calcalated (fib(1)
		of  uint = 1 // Previous fib calculated (fibo(0)
	)

	for cf < l {
		nf := of + cf
		if nf < l && nf%2 == 0 {
			efs += nf
		}
		of, cf = cf, nf
	}

	return efs
}
