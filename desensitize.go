package utils

// DesensitizePhone 脱敏手机号
func DesensitizePhone(phone string) string {
	return DesensitizePhoneScope(phone, 3, 7)
}

// DesensitizePhoneScope 脱敏手机号指定范围
func DesensitizePhoneScope(phone string, start, end int) string {
	if end <= start {
		return phone
	}
	if len(phone) < 11 {
		return phone
	}
	runePhone := []rune(phone)
	for i := start; i < end; i++ {
		runePhone[i] = '*'
	}
	return string(runePhone)
}

// DesensitizeEmail 脱敏邮箱
func DesensitizeEmail(email string) string {
	if len(email) < 6 {
		return email
	}
	runeEmail := []rune(email)
	for i := 1; i < len(runeEmail); i++ {
		if runeEmail[i] == '@' {
			break
		}
		runeEmail[i] = '*'
	}
	return string(runeEmail)
}

// DesensitizeName 脱敏姓名
func DesensitizeName(name string) string {
	if len(name) < 2 {
		return name
	}
	runeName := []rune(name)
	if len(runeName) == 2 {
		runeName[1] = '*'
	} else {
		for i := 1; i < len(runeName)-1; i++ {
			runeName[i] = '*'
		}
	}
	return string(runeName)
}

// DesensitizeIDCard 脱敏身份证号
func DesensitizeIDCard(idCard string) string {
	if len(idCard) != 15 && len(idCard) != 18 {
		return idCard
	}
	runeIDCard := []rune(idCard)
	for i := 6; i < len(runeIDCard)-3; i++ {
		runeIDCard[i] = '*'
	}
	return string(runeIDCard)
}
