package main

import (
	"e-learning/auth"
	"e-learning/courses"
	"e-learning/handler"
	"e-learning/membership"
	"e-learning/users"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=root password=123qweasdzxc dbname=e-learning port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// repository
	userRepository := users.Repositories(db)
	membershipRepositories := membership.Repositories(db)
	courseRepository := courses.Repositories(db)

	// service
	userService := users.Services(userRepository)
	authService := auth.AuthServices([]byte("123qweasdzxczxcasdqwe123123qweasdzxc"))
	membershipServices := membership.Services(membershipRepositories)
	courseServices := courses.Services(courseRepository)

	// middleware
	userMiddleware := auth.UsersMidlleware(authService, userService)
	adminMiddleware := auth.AdministratorMidlleware(authService, userService)

	// handler
	userHandler := handler.UserHandler(userService, authService, membershipServices)
	courseHandler := handler.CourseHandlers(courseServices)

	// membership domain
	app := fiber.New()

	v1 := app.Group("/api/v1")

	v1.Post("/user", userHandler.SignUpUser)
	v1.Post("/user/login", userHandler.LoginUser)
	v1.Put("/user", userMiddleware, userHandler.UpdateUserInfo)
	v1.Post("/user/password", userMiddleware, userHandler.UpdateUserPassword)

	v1.Post("/course", adminMiddleware, courseHandler.CreateNewCourse)
	v1.Put("/course/:id", adminMiddleware, courseHandler.UpdateCourse)
	v1.Get("/course/:id", courseHandler.GetCourseById)
	v1.Post("/course/:id", adminMiddleware, courseHandler.CreateSubCourseByCourseId)
	v1.Put("/course/:id/sub-course/:sub_id", adminMiddleware, courseHandler.UpdateSubCourse)

	// `uri:"course_id"`
	// SubCourseId int `uri:"sub_course_id"`
	log.Fatal(app.Listen(":8181"))
}
