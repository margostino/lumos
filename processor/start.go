package processor

func StartReply(input string) (bool, string) {
	if input == "start" {
		return true, "👋   Welcome to Lumos!\n🌱   Raise awareness.\n🔊   Spread the word."
	}
	return false, ""
}
