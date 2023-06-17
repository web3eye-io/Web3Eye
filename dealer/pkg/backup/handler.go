package backup

type Handler struct {
	Index uint64
}

func NewHandler(options ...func(*Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithIndex(index uint64) func(*Handler) error {
	return func(h *Handler) error {
		h.Index = index
		return nil
	}
}
