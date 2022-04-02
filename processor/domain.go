package processor

type Replier interface {
	Apply(input string) bool
	Reply() string
}

type CountryReplier struct {
	Id string
}

type FallbackReplier struct {
	Id string
}
