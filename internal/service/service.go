package services

// service defines any Music service recognized by song.link
// Since most services have multiple url version (long, short, etc...),
// it is allowd to have multiple ones.
// This allows for seamless song link detection
type service struct {
	name     string
	patterns []string
}

func GetServices() []service {
	return []service{
		*New(
			"Apple Music",
			[]string{"https://music.apple.com/"},
		),
		*New(
			"Spotify",
			[]string{"https://open.spotify.com/track/", "https://open.spotify.com/album/", "https://open.spotify.com/playlist/"},
		),
		*New(
			"Youtube Music",
			[]string{"https://music.youtube.com/"},
		),
		// *New(
		// 	"Amazon Music",
		// 	[]string{"https://music.amazon.com/",},
		// ),
		*New(
			"Tidal",
			[]string{"https://tidal.com/"},
		),
		*New(
			"Deezer",
			[]string{"https://www.deezer.com/"},
		),
		*New(
			"Pandora",
			[]string{"https://www.pandora.com/"},
		),
	}
}

func New(name string, patterns []string) *service {
	return &service{name, patterns}
}

func (s *service) Name() string {
	return s.name
}

func (s *service) Owns(url string) bool {
	// If the url begins the same as any of the patterns, it is owned by the service
	for _, pattern := range s.patterns {
		if len(url) >= len(pattern) && url[:len(pattern)] == pattern {
			return true
		}
	}
	return false
}

// Resolve returns the song.link url for the given service url
func (s *service) Resolve(url string) string {
	return "https://song.link/" + url
}
