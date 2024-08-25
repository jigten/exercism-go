package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	var encrypted []rune

	for _, c := range plain {
		if c >= 'a' && c <= 'z' {
			encrypted = append(encrypted, 'a'+(c-'a'+rune(shiftKey))%26)
		} else if c >= 'A' && c <= 'Z' {
			encrypted = append(encrypted, 'A'+(c-'A'+rune(shiftKey))%26)
		} else {
			encrypted = append(encrypted, c)
		}

	}

	return string(encrypted)
}
