package processor

func StartReply(input string) (bool, string) {
	if input == "start" {
		return true, "š   Welcome to Lumos!\nš±   Raise awareness.\nš   Spread the word."
	}
	return false, ""
}
