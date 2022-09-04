package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	bNum1 := []byte(num1)
	bNum2 := []byte(num2)
	tmpNum := make([]int, len(bNum1)+len(bNum2))

	for i := 0; i < len(bNum1); i++ {
		for j := 0; j < len(bNum2); j++ {
			tmpNum[i+j+1] += int(bNum1[i]-'0') * int(bNum2[j]-'0')
		}
	}

	for i := len(tmpNum) - 1; i > 0; i-- {
		tmpNum[i-1] += tmpNum[i] / 10
		tmpNum[i] = tmpNum[i] % 10
	}

	if tmpNum[0] == 0 {
		tmpNum = tmpNum[1:]
	}

	result := make([]byte, len(tmpNum))
	for i := 0; i < len(tmpNum); i++ {
		result[i] = byte(tmpNum[i]) + '0'
	}

	return string(result)
}

func main() {
	multiply("123", "65")
}
