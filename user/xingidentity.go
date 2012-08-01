package user

import (
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/view"
)

///////////////////////////////////////////////////////////////////////////////
// XingIdentity

type XingIdentity struct {
	ID model.String
}

func (self *XingIdentity) ProfileURL() string {
	if self.ID == "" {
		return ""
	}
	return "http://www.xing.com/profile/" + self.ID.Get()
}

func (self *XingIdentity) URL(args ...string) string {
	return self.ProfileURL()
}

func (self *XingIdentity) LinkContent(urlArgs ...string) view.View {
	return view.Escape(self.LinkTitle(urlArgs...))
}

func (self *XingIdentity) LinkTitle(urlArgs ...string) string {
	return self.ID.Get()
}

func (self *XingIdentity) LinkRel() string {
	return ""
}
