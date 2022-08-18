package compress

import (
	"testing"

	"github.com/smallpdf/testing-poc/pkg/config"
	"github.com/smallpdf/testing-poc/pkg/filestorage"
	"github.com/smallpdf/testing-poc/pkg/task"
)

type handler struct {
	filestorage *filestorage.Client
	tasks       *task.Client
}

func new() *handler {
	cfg := config.Load()

	fsClient := filestorage.NewClient(cfg.Filestorage.Endpoint)
	tClient := task.NewClient(cfg.Tasks.Endpoint)

	return &handler{
		filestorage: fsClient,
		tasks:       tClient,
	}
}

func TestCompress(t *testing.T) {
	h := new()

	// t.Fail()

	t.Log("testing compress")
}


type Test struct {
	task Task
	validate(args ...func() error )
}


var []testcase = {
	//here
}
