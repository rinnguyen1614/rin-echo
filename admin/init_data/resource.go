package initdata

import (
	"errors"
	"fmt"
	"os"
	"rin-echo/admin/app/command"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils/file"

	"gopkg.in/yaml.v3"
)

func initResources(uow iuow.UnitOfWork, path string) error {
	if !file.IsFile(path) {
		return errors.New("resource file doesn't exist")
	}
	fs, err := os.Open(path)
	if err != nil {
		return err
	}

	defer fs.Close()

	var resources command.CreateResources
	decoder := yaml.NewDecoder(fs)
	if err = decoder.Decode(&resources); err != nil {
		return fmt.Errorf("resource file decodes error: %v", err)
	}

	for _, r := range resources {
		if err = r.CreateRecursive(uow, nil); err != nil {
			return err
		}
	}

	return nil
}
