package ports

type RestServer interface {
	Run()
	SetupRoute()
	Shutdown()
}
