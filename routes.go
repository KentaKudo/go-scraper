package main

func (s *server) routes() {
	s.router.HandleFunc("/healthz", s.handleHealthz())
	s.router.HandleFunc("/", s.handleIndex())
}
