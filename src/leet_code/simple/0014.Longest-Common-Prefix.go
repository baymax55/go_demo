package leetcode

//func main() {
//	//str := []string{"aflower", "sflow", "dflight"}
//	str := []string{"flower", "flow", "flight"}
//
//	println(longestCommonPrefix(str))
//}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}

	return prefix
}
