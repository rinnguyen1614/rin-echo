package initdata

import (
	"errors"
	"fmt"
	"os"

	"github.com/rinnguyen1614/rin-echo/internal/system/adapters/repository"

	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils/file"
	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
	"gopkg.in/yaml.v3"
)

type Resource struct {
	Name        string     `json:"name" validate:"required,min=5"`
	Slug        string     `json:"slug" validate:"required,min=6"`
	Object      string     `json:"object"`
	Action      string     `json:"action"`
	Description string     `json:"description"`
	Children    []Resource `json:"children"`
}

func initResources(uow iuow.UnitOfWork, path string) error {
	if !file.IsFile(path) {
		return errors.New("resource file doesn't exist")
	}
	fs, err := os.Open(path)
	if err != nil {
		return err
	}

	defer fs.Close()

	var (
		repo      = repository.NewResourceRepository(uow.DB())
		resources []Resource
	)
	decoder := yaml.NewDecoder(fs)
	if err = decoder.Decode(&resources); err != nil {
		return fmt.Errorf("resource file decodes error: %v", err)
	}
	for _, re := range resources {
		if _, err := createResource(repo, re, nil); err != nil {
			return err
		}
	}
	return nil
}

func createResource(repo domain.ResourceRepository, cmd Resource, parent *domain.Resource) (uint, error) {
	resource, err := domain.NewResource(cmd.Name, cmd.Slug, cmd.Object, cmd.Action, cmd.Description, parent)
	if err != nil {
		return 0, err
	}

	if err = repo.Create(resource); err != nil {
		return 0, err
	}

	for _, mc := range cmd.Children {
		if _, err := createResource(repo, mc, resource); err != nil {
			return 0, err
		}
	}

	return resource.ID, nil
}
