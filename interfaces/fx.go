// A module that provides a HTTP server.
newHTTPServer := func() fx.Option {
	port := os.Getenv("PORT")
	if port == "" {
		return port = "80"
	}
	return fx.Provide(&http.Server{
		Addr: fmt.Sprintf(":%s", port),
	})
}

app := fx.New(
	newHTTPServer(),
	fx.Invoke(func(s *http.Server) error { return s.ListenAndServe() }),
)

fmt.Println(app.Err())
