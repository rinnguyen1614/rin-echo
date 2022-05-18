package handler

import echox "rin-echo/common/echo"

func (h AccountHandler) Profile(c echox.Context) error {
	session := c.MustSession()
	profile, err := h.service.WithContext(c).Profile(session.UserID())
	if err != nil {
		return err
	}

	echox.OKWithData(c, profile)
	return nil
}

func (h AccountHandler) UpdateProfile(c echox.Context) error {
	session := c.MustSession()
	profile, err := h.service.WithContext(c).Profile(session.UserID())
	if err != nil {
		return err
	}

	echox.OKWithData(c, profile)
	return nil
}

func (h AccountHandler) UpdateAvatar(c echox.Context) error {
	session := c.MustSession()
	profile, err := h.service.WithContext(c).Profile(session.UserID())
	if err != nil {
		return err
	}

	echox.OKWithData(c, profile)
	return nil
}
