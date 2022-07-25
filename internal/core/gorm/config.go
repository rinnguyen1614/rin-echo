package gorm

import "fmt"

type Database struct {
	Host        string            `json:"host,omitempty" yaml:"host,omitempty" ini:"host,omitempty"`
	Port        string            `json:"port,omitempty" yaml:"port,omitempty" ini:"port,omitempty"`
	User        string            `json:"user,omitempty" yaml:"user,omitempty" ini:"user,omitempty"`
	Pwd         string            `json:"pwd,omitempty" yaml:"pwd,omitempty" ini:"pwd,omitempty"`
	Name        string            `json:"name,omitempty" yaml:"name,omitempty" ini:"name,omitempty"`
	Driver      string            `json:"driver,omitempty" yaml:"driver,omitempty" ini:"driver,omitempty"`
	DNS         string            `json:"dns,omitempty" yaml:"dns,omitempty" ini:"dns,omitempty"`
	Params      map[string]string `json:"params,omitempty" yaml:"params,omitempty" ini:"params,omitempty"`
	MaxIdleCon  int               `json:"max_idle_con,omitempty" yaml:"max_idle_con,omitempty" ini:"max_idle_con,omitempty"`
	MaxOpenCon  int               `json:"max_open_con,omitempty" yaml:"max_open_con,omitempty" ini:"max_open_con,omitempty"`
	BatchSize   int               `json:"batch_size,omitempty" yaml:"batch_size,omitempty" ini:"batch_size,omitempty"`
	PrepareStmt bool              `json:"prepare_smt,omitempty" yaml:"prepare_smt,omitempty" ini:"prepare_smt,omitempty"`
}

func (d Database) GetDNS() string {
	if d.DNS != "" {
		return d.DNS
	}
	if d.Driver == DriverPostgresql {
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s"+d.ParamStr(), d.Host, d.Port, d.User, d.Pwd, d.Name)
	}
	return ""
}

func (d Database) ParamStr() string {
	p := ""
	fJoinKeyValue := func(params map[string]string, strjoin string, sep string) string {
		str := sep
		for k, v := range params {
			str += k + strjoin + v + sep
		}
		return str[:len(str)-len(sep)]
	}
	if d.Params == nil {
		d.Params = make(map[string]string)
	}

	if d.Driver == DriverPostgresql {
		if _, ok := d.Params["sslmode"]; !ok {
			d.Params["sslmode"] = "disable"
		}
		fJoinKeyValue(d.Params, "=", " ")
	}
	return p
}
