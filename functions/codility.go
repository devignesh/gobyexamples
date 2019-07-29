package solution

import (
	"math"
	"strconv"
	"strings"
)

func Solution(A []int) string {
	response := processSolution(A)

	if A[0]%2 != 0 && response == "NO SOLUTION" && len(A) > 1 {
		lastindex := strconv.Itoa(len(A) - 1)
		return "1," + lastindex
	} else {
		return response
	}
}

func processSolution(A []int) string {
	// write your code in Go 1.4
	var (
		evenCount,
		oddCount,
		index,
		indexCount int
	)
	res := ""
	if len(A) == 0 {
		return "NO SOLUTION"
	}

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	if len(A) == 1 && oddCount > 0 {
		return "NO SOLUTION"
	} else if oddCount%2 == 0 {
		lastindex := strconv.Itoa(len(A) - 1)
		return "0," + lastindex
	} else {
		if A[len(A)-1]%2 != 0 {
			lastindex := strconv.Itoa(len(A) - 2)
			return "0," + lastindex
		} else if A[0]%2 != 0 && oddCount == 1 {
			lastindex := strconv.Itoa(len(A) - 1)
			return "1," + lastindex
		}
	}
	//
	var (
		rightEven,
		leftEven,
		leftOddEven,
		tempIndex,
		tempIndex1,
		j int
	)
	rightOddEven := 1

	for p := len(A) - 2; p > 0; p-- {
		if A[p]%2 == 0 {
			rightOddEven++
		} else if A[p]%2 != 0 {
			tempIndex = p
			for w := p - 1; w > 0; w-- {
				if A[w]%2 == 0 {
					tempIndex1 = w
					leftOddEven++
				} else if A[w]%2 != 0 {
					w = 0
					p = 0
				}
			}
		}
	}
	if leftOddEven > 0 {
		for p := 0; p < tempIndex1; p++ {
			if A[p]%2 == 0 {
				leftOddEven++
			} else {
				p = tempIndex1
			}
		}
		if leftOddEven >= rightOddEven {
			j = tempIndex
			res = check(j, A)
		}
	}

	if A[len(A)-2]%2 != 0 && leftOddEven < rightOddEven {
		j = len(A) - 2
		res = check(j, A)
	} else if oddCount%2 != 0 && leftOddEven < rightOddEven {
		var (
			proceed,
			rightRight,
			inside int
		)

		for i1 := len(A) - 2; i1 > 1; i1-- {
			if A[i1]%2 == 0 && inside == 0 {
				rightRight++
			} else if A[i1]%2 != 0 {
				inside = 1
				for {
					if rightRight > 0 {
						if i1-rightRight > 0 {
							if A[i1-rightRight]%2 == 0 {
								proceed = 2
							} else {
								proceed = 1
								rightRight = 0
							}
						}
						rightRight--
					} else {
						break
					}
				}

				if proceed == 2 {
					j = i1
					res = check(j, A)
				}
				i1 = 0
			}
		}

		z := len(A) - 2
		if proceed == 0 || proceed == 1 {
			for i := 1; i < len(A)-1/2; i++ {
				if A[z]%2 != 0 {
					j = z
					res = check(j, A)
					i = len(A)
				} else if A[i]%2 != 0 {
					j = i
					res = check(j, A)
					i = len(A)
				}
				z--
			}
		}
	} // elseif

	// explode the res
	explodeRes := strings.Split(res, ",")
	if len(explodeRes) > 1 {
		leftEven, _ = strconv.Atoi(explodeRes[0])
		rightEven, _ = strconv.Atoi(explodeRes[1])
	}
	// find diff
	diff := 0
	diff = rightEven - leftEven
	if diff == 0 {
		return "NO SOLUTION"
	} else if diff > 0 {
		temp := diff
		odd := 0
		index = j
		loop := j
		indexCount = 0
		//--while
		for {
			if j < len(A)-1 {
				indexCount++
				if A[loop]%2 == 0 {
					temp = temp - 1
				} else {
					odd++
					if odd == 2 {
						temp = temp - 1
						odd = 0
					}
				}
				loop++
				if temp == 0 && odd == 0 {
					indexCount = index + indexCount - 1
					result := stringConcat(index, indexCount)
					return result
				}
				if temp == 0 && odd != 0 {
					j = j + 1
					index = j
					loop = j
					indexCount = 0
					odd = 0
					temp = diff
				}

			} else {
				break
			}
		}

		if j == len(A)-1 {
			return "NO SOLUTION"
		}
		indexCount = index + indexCount - 1

		result := stringConcat(index, indexCount)
		return result
	} else if diff < 0 {
		diff1 := math.Abs(float64(diff))

		loopLimit := j
		j = 0
		temp := diff1
		odd := 0
		index = j
		loop := j
		indexCount = 0
		// while
		for {
			if j < loopLimit && loop < len(A) {
				indexCount++
				if A[loop]%2 == 0 {
					temp = temp - 1
				} else {
					odd++
					if odd == 2 {
						temp = temp - 1
						odd = 0
					}
				}
				loop++
				if temp == 0 && odd == 0 {
					indexCount = index + indexCount - 1
					tm := loop
					// while
					for {
						if tm < loopLimit {
							if A[tm]%2 != 0 {
								temp = 0
								odd = 1
							}
							tm++
						} else {
							break
						}
					}

					if (A[indexCount+1]%2 == 0 || indexCount+1 == loopLimit) && odd == 0 {
						result := stringConcat(index, indexCount)
						return result
					} else {
						temp = 0
						odd = 1
					}
				}

				if temp == 0 && odd != 0 {
					j = j + 1
					index = j
					loop = j
					indexCount = 0
					odd = 0
					temp = diff1
				}

			} else {
				break
			}
		}

		if j == len(A)-1 {
			return "NO SOLUTION"
		}

		indexCount = index + indexCount - 1

		result := stringConcat(index, indexCount)
		return result
	}

	result := stringConcat(index, indexCount)
	return result
}

func check(j int, a []int) string {
	rightEven := 0
	leftEven := 0
	odd := 0

	for i := j; i < len(a); i++ {
		if a[i]%2 == 0 {
			rightEven++
		} else {
			odd++
		}
	}

	odd = odd / 2
	rightEven = rightEven + odd
	odd = 0
	for i := 0; i < j; i++ {
		if a[i]%2 == 0 {
			leftEven++
		} else {
			odd++
		}
	}
	odd = odd / 2
	leftEven = leftEven + odd
	result := stringConcat(leftEven, rightEven)

	return result
}

func stringConcat(left int, right int) string {
	resLeftEven := strconv.Itoa(left)
	resRightEven := strconv.Itoa(right)
	result := resLeftEven + "," + resRightEven

	return result
}
