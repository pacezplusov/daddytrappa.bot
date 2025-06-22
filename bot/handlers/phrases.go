package handlers

import "strings"

func HandlePhrase(text string) string {
	switch {
	case strings.Contains(text, "шляпа"):
		return "увижу хуйню — будет пизда, мне насрано"
	case strings.Contains(text, "долг"):
		return "нет кайфа ща закидывать"
	case strings.Contains(text, "91к"):
		return "совсем пидорасы?"
	case strings.Contains(text, "трек"):
		return "ушки пердят"
	case strings.Contains(text, "кайф"):
		return "да не это кумарит"
	case strings.Contains(text, "газ"):
		return "тапочка в пол))"
	case strings.Contains(text, "борян"):
		return "боряе ты меня пздц доебать решил"
	case strings.Contains(text, "суббота"):
		return "она помогла мне не занами"
	default:
		return ""
	}
}
