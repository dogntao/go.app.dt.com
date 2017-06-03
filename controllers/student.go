package controllers

type StudentController struct {
	BaseController
}

// 首页
func (s *StudentController) Index() {
	s.Display("views/student/student.html")
}
