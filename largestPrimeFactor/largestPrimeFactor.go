// Brief one line overview of package
package largestPrimeFactor

func Problem3(n uint) uint {

	ps := []uint{} // Prime numbers
	pf := []uint{} // Prime Factors we have found so far
	guard := uint(1)

	i := uint(2)

	for guard < n {
		for _, p := range ps {
			if i%p == 0 { // i is not prime
				i++
				break
			}
		}

		ps = append(ps, i)

		if n%i == 0 {
			pf = append(pf, i)
			guard *= i
		}
		i++
	}
	return pf[len(pf)-1]
}
