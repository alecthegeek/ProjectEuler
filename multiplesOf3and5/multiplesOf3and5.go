// Brief one line overview of package
package multiplesOf3and5

func Problem1(n uint) uint {

	var sum uint
	for i := uint(3); i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
