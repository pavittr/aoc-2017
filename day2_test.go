package aoc17

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayTwoPuzzleOne(t *testing.T) {
	//5 1 9 5
	//7 5 3
	//2 4 6 8
	//The first row's largest and smallest values are 9 and 1, and their difference is 8.
	//The second row's largest and smallest values are 7 and 3, and their difference is 4.
	//The third row's difference is 6.
	//In this example, the spreadsheet's checksum would be 8 + 4 + 6 = 18.
	arraySmmer := func(array string) int {
		totalCount := 0
		// First split by line
		for _, line := range strings.Split(array, "\n") {
			high := 0
			low := 10000
			if line == "" {
				continue
			}
			for _, number := range strings.Split(strings.Replace(line, "\t", " ", -1), " ") {
				currentNumber, err := strconv.Atoi(number)
				if err != nil {
					t.Logf("Got err %+v\n", err)
					return -1
				}
				if low > currentNumber {
					low = currentNumber
				}
				if high < currentNumber {
					high = currentNumber
				}

			}
			totalCount += (high - low)
		}
		return totalCount
	}
	testArray := `
5 1 9 5
7 5 3
2 4 6 8`
	assert.Equal(t, 18, arraySmmer(testArray))

	myInput := `
798	1976	1866	1862	559	1797	1129	747	85	1108	104	2000	248	131	87	95
201	419	336	65	208	57	74	433	68	360	390	412	355	209	330	135
967	84	492	1425	1502	1324	1268	1113	1259	81	310	1360	773	69	68	290
169	264	107	298	38	149	56	126	276	45	305	403	89	179	394	172
3069	387	2914	2748	1294	1143	3099	152	2867	3082	113	145	2827	2545	134	469
3885	1098	2638	5806	4655	4787	186	4024	2286	5585	5590	215	5336	2738	218	266
661	789	393	159	172	355	820	891	196	831	345	784	65	971	396	234
4095	191	4333	161	3184	193	4830	4153	2070	3759	1207	3222	185	176	2914	4152
131	298	279	304	118	135	300	74	269	96	366	341	139	159	17	149
1155	5131	373	136	103	5168	3424	5126	122	5046	4315	126	236	4668	4595	4959
664	635	588	673	354	656	70	86	211	139	95	40	84	413	618	31
2163	127	957	2500	2370	2344	2224	1432	125	1984	2392	379	2292	98	456	154
271	4026	2960	6444	2896	228	819	676	6612	6987	265	2231	2565	6603	207	6236
91	683	1736	1998	1960	1727	84	1992	1072	1588	1768	74	58	1956	1627	893
3591	1843	3448	1775	3564	2632	1002	3065	77	3579	78	99	1668	98	2963	3553
2155	225	2856	3061	105	204	1269	171	2505	2852	977	1377	181	1856	2952	2262
`
	assert.Equal(t, 41919, arraySmmer(myInput))

}

func splitArray(array string) [][]int {
	returnedArray := make([][]int, 0)

	for _, line := range strings.Split(array, "\n") {
		if line == "" {
			continue
		}
		neArray := make([]int, 0)

		for _, number := range strings.Split(strings.Replace(line, "\t", " ", -1), " ") {
			currentNumber, err := strconv.Atoi(number)
			if err != nil {
				return nil
			}
			neArray = append(neArray, currentNumber)

		}
		returnedArray = append(returnedArray, neArray)
	}
	return returnedArray

}

func TestArraySplitter(t *testing.T) {
	testArray := `
5 9 2 8
9 4 7 3
3 8 6 5`

	returnedArray := splitArray(testArray)
	assert.NotNil(t, returnedArray)
	assert.Equal(t, 5, returnedArray[0][0])
	assert.Equal(t, 9, returnedArray[0][1])
	assert.Equal(t, 2, returnedArray[0][2])
	assert.Equal(t, 8, returnedArray[0][3])

	assert.Equal(t, 9, returnedArray[1][0])
	assert.Equal(t, 4, returnedArray[1][1])
	assert.Equal(t, 7, returnedArray[1][2])
	assert.Equal(t, 3, returnedArray[1][3])

	assert.Equal(t, 3, returnedArray[2][0])
	assert.Equal(t, 8, returnedArray[2][1])
	assert.Equal(t, 6, returnedArray[2][2])
	assert.Equal(t, 5, returnedArray[2][3])

}

func TestDayTwoPuzzleTwo(t *testing.T) {
	//5 9 2 8
	//9 4 7 3
	//3 8 6 5
	//In the first row, the only two numbers that evenly divide are 8 and 2; the result of this division is 4.
	//In the second row, the two numbers are 9 and 3; the result is 3.
	//In the third row, the result is 2.
	//In this example, the sum of the results would be 4 + 3 + 2 = 9
	arraySummByDivisor := func(array string) int {
		totalCount := 0
		returnedArray := splitArray(array)
		assert.NotNil(t, returnedArray)
		for _, sortedLine := range returnedArray {
			// Find a pair of divisible numbers
			sort.Ints(sortedLine)
			lineValue := 0
			for i := len(sortedLine) - 1; i >= 0; i-- {
				elementToProcess := sortedLine[i]
				for j := i - 1; j >= 0; j-- {
					nextElement := sortedLine[j]
					if elementToProcess%nextElement == 0 {
						lineValue = elementToProcess / nextElement
						break
					}
				}
				if lineValue > 0 {
					break
				}
			}
			totalCount += lineValue
		}
		return totalCount
	}
	testArray := `
5 9 2 8
9 4 7 3
3 8 6 5`
	assert.Equal(t, 9, arraySummByDivisor(testArray))

	myInput := `
798	1976	1866	1862	559	1797	1129	747	85	1108	104	2000	248	131	87	95
201	419	336	65	208	57	74	433	68	360	390	412	355	209	330	135
967	84	492	1425	1502	1324	1268	1113	1259	81	310	1360	773	69	68	290
169	264	107	298	38	149	56	126	276	45	305	403	89	179	394	172
3069	387	2914	2748	1294	1143	3099	152	2867	3082	113	145	2827	2545	134	469
3885	1098	2638	5806	4655	4787	186	4024	2286	5585	5590	215	5336	2738	218	266
661	789	393	159	172	355	820	891	196	831	345	784	65	971	396	234
4095	191	4333	161	3184	193	4830	4153	2070	3759	1207	3222	185	176	2914	4152
131	298	279	304	118	135	300	74	269	96	366	341	139	159	17	149
1155	5131	373	136	103	5168	3424	5126	122	5046	4315	126	236	4668	4595	4959
664	635	588	673	354	656	70	86	211	139	95	40	84	413	618	31
2163	127	957	2500	2370	2344	2224	1432	125	1984	2392	379	2292	98	456	154
271	4026	2960	6444	2896	228	819	676	6612	6987	265	2231	2565	6603	207	6236
91	683	1736	1998	1960	1727	84	1992	1072	1588	1768	74	58	1956	1627	893
3591	1843	3448	1775	3564	2632	1002	3065	77	3579	78	99	1668	98	2963	3553
2155	225	2856	3061	105	204	1269	171	2505	2852	977	1377	181	1856	2952	2262
`
	assert.Equal(t, 303, arraySummByDivisor(myInput))

}
