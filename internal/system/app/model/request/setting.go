package request

type SettingSecurity struct {
	RequireEmailConfirmationForLogin bool `json:"require_email_confirmation_for_login"`
}

type SettingApp struct {
	LogoPath string
}

type SettingEmail struct {
}

type Setting struct {
	Security SettingSecurity `json:"security"`
}
