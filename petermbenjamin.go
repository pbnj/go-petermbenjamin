package petermbenjamin

const (
	name  = "Peter M. Benjamin"
	title = "Senior Security Software Engineer"
)

// maps cannot be const'ed
var (
	contactInfo = map[string]string{
		"twitter":  "https://twitter.com/pmbenjamin",
		"github":   "https://github.com/petermbenjamin",
		"blog":     "https://petermbenjamin.github.io/",
		"linkedIn": "https://linkedin.com/in/pmbenjamin",
		"npm":      "https://www.npmjs.com/~pmbenjamin",
		"email":    "petermbenjamin@gmail.com",
		"keybase":  "https://keybase.io/pbenj",
	}
)

// FullName returns full name as a string
func FullName() string {
	return name
}

// Title returns title as a string
func Title() string {
	return title
}

// Contact returns contact info as a map
func Contact() map[string]string {
	return contactInfo
}
