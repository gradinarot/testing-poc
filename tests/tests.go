package tests

var

func TestMain() {
	cfg := config.Load()

	fsClient := filestorage.NewClient(cfg.Filestorage.Endpoint)
	tClient := task.NewClient(cfg.Tasks.Endpoint)
}
