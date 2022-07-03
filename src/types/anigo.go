package types

type Anigo struct {
	Anime, Episode, Image, Url string
}

func (anigo *Anigo) UseService(svc string, mod string, args ...string) ([]interface{}, bool) {
	for _, s := range Services {
		if s.Solver != svc {

			continue
		}

		if hand, ok := s.Handler[mod]; ok {

			return hand(args...), true
		}

	}

	return []any{}, false
}

func (anigo *Anigo) DirectUrl(u string) (string, bool) {
	for _, p := range Providers {
		if resp, ok := p.Handler(u); ok {

			return resp, true
		}

	}

	return u, false
}

func (anigo *Anigo) GetServices() (svc []string) {
	for _, p := range Services {
		svc = append(svc, p.Solver)

	}

	return
}
