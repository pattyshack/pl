package cfg

type Config struct {
	BuildTags map[string]struct{}
}

func NewConfig(additionalTags []string) *Config {
	// TODO: populate platform arch / os tags, etc.
	tags := map[string]struct{}{}
	for _, tag := range additionalTags {
		if tag == "" || tag == "_" {
			continue
		}

		tags[tag] = struct{}{}
	}

	return &Config{
		BuildTags: tags,
	}
}
