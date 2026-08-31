package mathematicsfordsa

func gcd(a int ,b int)int {
	if b == 0{
		return a
	}
	return gcd(b, a%b)
}

func Lcm(a,b int)int{
	res := (a*b) / gcd(a,b)
	return res
}

func gcdOfOddEvenSums(n int) int {
	var sumOdd int
	var sumEven int 

    for i := 1; i <= n; i++ {
		sumEven += 2 * i        
        sumOdd += (2 * i) - 1    
    }
	return gcd(sumOdd,sumEven)
}