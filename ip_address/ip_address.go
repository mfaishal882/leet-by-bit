   

func defangIPaddr(address string) string {
	//return strings.Replace(address, ".", "[.]", -1)

	// replace using byte
	var result []byte
	for _, value := range address {
		if value == '.' {
			result = append(result, "[.]"...)
		} else {
			result = append(result, byte(value))
		}
	}
	return string(result)
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}
