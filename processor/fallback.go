package processor

func (r FallbackReplier) Apply(_ string) bool {
	return true
}

func (r FallbackReplier) Reply() string {
	return "🪄   Happiness can be found, even in the darkest of times, if one only remembers to turn on the light.\n" +
		"🌎   We do not need magic to transform our world."
}
