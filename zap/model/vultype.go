package model

var vulType = map[int]string{
	4:   "Response Without X-Content-Type-Options Header",
	23:  "sql-injection",
	28:  "path-traversal",
	40:  "reflected-xss",
	82:  "cmd-injection",
	154: "Thymeleaf模版注入",
	158: "Thymeleaf模版注入",
}

func GetVulType(id int, name string) string {
	value, ok := vulType[id]
	if ok {
		return value
	} else {
		return name
	}
}
