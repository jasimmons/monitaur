package monitaur

type Handlers struct {
	m map[string]Handler
}

func (h *Handlers) Add(key string, handler Handler) {
	h.m[key] = handler
}

func (h *Handlers) Get(key string) (Handler, bool) {
	handler, ok := h.m[key]
	return handler, ok
}
