package monitaur

type Config struct {
	ConfigDir string
}

var (
	TYPE_CHECK   = "check"
	TYPE_HANDLER = "handler"
)
