package _02510

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	set := map[string]bool{}
	equationMap := map[string][]string{}
	reverseEquationMap := map[string][]string{}
	valMap := map[string]float64{}

	for i, equation := range equations {
		set[equation[0]] = true
		set[equation[1]] = true
		valMap[equation[0]+"_"+equation[1]] = values[i]
		if _, ok := equationMap[equation[0]]; ok {
			equationMap[equation[0]] = append(equationMap[equation[0]], equation[1])
		} else {
			equationMap[equation[0]] = []string{equation[1]}
		}
		if _, ok := reverseEquationMap[equation[1]]; ok {
			reverseEquationMap[equation[1]] = append(reverseEquationMap[equation[1]], equation[0])
		} else {
			reverseEquationMap[equation[1]] = []string{equation[0]}
		}
	}
	/*var paths []float64
	visited := map[string]bool{}
	var calculate func(a, b string)
	calculate = func(a, b string) {
		if v, ok := valMap[a+"_"+b]; ok {
			paths = append(paths, v)
			return
		}
		if list, ok:= equationMap[a];ok{

		}else{
			return
		}

	}*/

	return []float64{}
}
