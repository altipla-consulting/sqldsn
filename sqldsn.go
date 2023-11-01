package sqldsn

import (
	"fmt"
	"net/url"
	"strings"
)

// Deprecated: Use FromURL instead.
func PrismaToGo(dsn string) string {
	return FromURL(dsn)
}

// FromURL takes a URL format and converts it to the specific format of
// each kind of SQL driver.
func FromURL(dsn string) string {
	switch {
	case strings.HasPrefix(dsn, "mysql://"):
		u, err := url.Parse(dsn)
		if err != nil {
			panic(err)
		}
		pass, _ := u.User.Password()

		switch {
		case u.Host == "cloudsql":
			return fmt.Sprintf("%s:%s@cloudsql-mysql(%s)/%s?parseTime=true", u.User.Username(), pass, u.Query().Get("socket"), u.Path[1:])

		case u.Query().Has("socket"):
			return fmt.Sprintf("%s:%s@unix(%s)/%s?parseTime=true", u.User.Username(), pass, u.Query().Get("socket"), u.Path[1:])

		default:
			if u.Port() == "" {
				return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", u.User.Username(), pass, u.Hostname(), u.Path[1:])
			} else {
				return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", u.User.Username(), pass, u.Hostname(), u.Port(), u.Path[1:])
			}
		}

	case strings.HasPrefix(dsn, "sqlserver://"):
		segments := strings.Split(dsn[len("sqlserver://"):], ";")
		for i, segment := range segments {
			if !strings.Contains(segment, "=") {
				segments[i] = "server=" + strings.SplitN(segment, ":", 2)[0]
			}
		}
		return strings.Join(segments, ";")
	}

	panic(fmt.Sprintf("unsupported dsn transformation: %s", dsn))
}
