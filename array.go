package goutils

// StringInSlice 检查字符串是否存在于列表中，返回是否存在标识，若存在则返回所在索引，https://stackoverflow.com/questions/15323767/does-golang-have-if-x-in-construct-similar-to-python
func StringInSlice(a string, list []string) (bool, int) {
	for i, b := range list {
		if b == a {
			return true, i
		}
	}
	return false, -1
}
