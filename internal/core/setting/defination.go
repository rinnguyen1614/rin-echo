package setting

import "errors"

type SettingDefinition struct {
	Name             string
	DisplayName      string
	Description      string
	DefaultValue     string
	VisibleToClients bool
	AllowedProviders []string
	IsEncrypted      bool
}

type SettingDefinitionManager struct {
	definitionsByName map[string]SettingDefinition
	definitions       []SettingDefinition
}

func NewSettingDefinitionManager(definitions []SettingDefinition) *SettingDefinitionManager {
	if definitions == nil {
		panic("requires definitions")
	}
	var mSettings = make(map[string]SettingDefinition)
	for _, setting := range definitions {
		mSettings[setting.Name] = setting
	}

	return &SettingDefinitionManager{
		definitions:       definitions,
		definitionsByName: mSettings,
	}
}

func (s SettingDefinitionManager) Get(name string) (SettingDefinition, error) {
	if name == "" {
		panic("SettingDefinnationManager_Get requires name")
	}

	setting, ok := s.definitionsByName[name]
	if !ok {
		return SettingDefinition{}, errors.New("Undefined setting: " + name)
	}
	return setting, nil
}

func (s SettingDefinitionManager) GetMulti(names []string) []SettingDefinition {
	var definitions []SettingDefinition
	for _, name := range names {
		if definition, ok := s.definitionsByName[name]; ok {
			definitions = append(definitions, definition)
		}
	}

	return definitions
}

func (s SettingDefinitionManager) GetAll() []SettingDefinition {
	return s.definitions
}

func (s *SettingDefinitionManager) Add(definition SettingDefinition) error {
	var name = definition.Name
	if _, ok := s.definitionsByName[name]; ok {
		return errors.New("SettingDefinition has existed")
	}

	s.definitions = append(s.definitions, definition)
	s.definitionsByName[name] = definition

	return nil
}

func (s *SettingDefinitionManager) Remove(name string) error {
	if _, ok := s.definitionsByName[name]; !ok {
		return errors.New("SettingDefinition doesn't exist")
	}

	delete(s.definitionsByName, name)
	// clone from map instead of spliptting slice.
	s.definitions = make([]SettingDefinition, len(s.definitionsByName))
	i := 0
	for _, pro := range s.definitionsByName {
		s.definitions[i] = pro
		i++
	}

	return nil
}
