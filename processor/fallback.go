package processor

func FallbackReply(_ string) (bool, string) {
	return true, "🪄   Happiness can be found, even in the darkest of times, if one only remembers to turn on the light.\n" +
		"🌎   We do not need magic to transform our world."
}
