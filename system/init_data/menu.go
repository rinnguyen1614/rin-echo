package initdata

import (
	"errors"
	"fmt"
	"os"
	"rin-echo/system/adapters/repository"
	"rin-echo/system/domain"

	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo-core/utils/file"

	"gopkg.in/yaml.v3"
)

type Menu struct {
	Name      string `json:"name" `
	Slug      string `json:"slug"`
	ParentID  uint   `json:"parent_id" `
	Path      string `json:"path"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      int    `json:"sort"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Children  []Menu `json:"children"`
}

func initMenus(uow iuow.UnitOfWork, path string) error {
	if !file.IsFile(path) {
		return errors.New("menu file doesn't exist")
	}
	fs, err := os.Open(path)
	if err != nil {
		return err
	}

	defer fs.Close()

	var (
		menus        []Menu
		repo         = repository.NewMenuRepository(uow.DB())
		repoResource = repository.NewResourceRepository(uow.DB())
	)
	decoder := yaml.NewDecoder(fs)
	if err = decoder.Decode(&menus); err != nil {
		return fmt.Errorf("resource file decodes error: %v", err)
	}

	for _, re := range menus {
		if _, err := createMenu(repo, repoResource, re, nil); err != nil {
			return err
		}
	}
	return nil
}

func createMenu(repo domain.MenuRepository, repoResource domain.ResourceRepository, cmd Menu, parent *domain.Menu) (uint, error) {

	menu, err := domain.NewMenu(cmd.Name, cmd.Slug, cmd.Path, cmd.Hidden, cmd.Component, cmd.Sort, cmd.Type, cmd.Icon, cmd.Title, parent)
	if err != nil {
		return 0, err
	}

	if err = repo.Create(menu); err != nil {
		return 0, err
	}

	for _, mc := range cmd.Children {
		if _, err := createMenu(repo, repoResource, mc, menu); err != nil {
			return 0, err
		}
	}

	return menu.ID, nil
}
