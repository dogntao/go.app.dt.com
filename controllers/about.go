package controllers

type AboutController struct {
	BaseController
}

// 关于我们
func (a *AboutController) About() {
	a.Display("views/about/about.html")
}

// 团队介绍
func (a *AboutController) Team() {
	a.Display("views/about/team.html")
}

// 联系我们
func (a *AboutController) Contact() {
	a.Display("views/about/contact.html")
}
