package controllers

type CourseController struct {
	BaseController
}

// 最新课程
func (c *CourseController) Index() {
	c.Display("views/course/course.html")
}
