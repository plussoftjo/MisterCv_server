// Package routes (Setup Routes Group)
package routes

import (
	"server/config"
	"server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Setup >>>
func Setup() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "authorization", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))

	// Config
	r.LoadHTMLGlob(config.ServerInfo.PublicPath + "public/Templates/CvTemplates/*")

	// gin.SetMode(gin.ReleaseMode)

	r.Use(static.Serve("/public", static.LocalFile(config.ServerInfo.PublicPath+"public", true)))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success",
		})
	})

	// -------- Auth Groups ----------//

	// ~~~ Auth Group ~~~ //
	auth := r.Group("/auth")
	// Auth Routes
	auth.POST("/login", controllers.LoginController)
	auth.POST("/register", controllers.RegisterController)
	auth.GET("/auth", controllers.Auth)
	auth.POST("/update", controllers.UpdateUser)

	// ~~~ Cvs Group ~~~ //
	cvs := r.Group("/cvs")
	cvs.POST("/store", controllers.StoreCv)
	cvs.POST("/update", controllers.UpdateCV)
	cvs.POST("/update_template_id/:id", controllers.UpdateCvTemplate)
	cvs.POST("/search_cv", controllers.SearchInCvs)
	cvs.GET("/index", controllers.IndexCv)
	cvs.POST("/basics_information/store", controllers.StoreBasicsInformations)
	cvs.POST("/basics_information/update", controllers.UpdateBasicsInformations)
	cvs.POST("/basics_information/image/store", controllers.BasicsInformationsUploadImage)
	cvs.POST("/summary/store", controllers.StoreSummary)
	cvs.POST("/summary/update", controllers.UpdateSummary)
	cvs.POST("/skills/store", controllers.StoreSkills)
	cvs.POST("/skills/update", controllers.UpdateSkills)
	cvs.GET("/skills/delete/:id", controllers.DeleteSkill)
	cvs.POST("/work_experience/store", controllers.StoreWorkExperience)
	cvs.POST("/work_experience/update", controllers.UpdateWorkExperience)
	cvs.GET("/work_experience/delete/:id", controllers.DeleteWorkExperience)
	cvs.POST("/educations/store", controllers.StoreEducations)
	cvs.POST("/educations/update", controllers.UpdateEducations)
	cvs.GET("/educations/delete/:id", controllers.DeleteEducation)
	cvs.POST("/certifications/store", controllers.StoreCertifications)
	cvs.POST("/certifications/update", controllers.UpdateCertifications)
	cvs.GET("/certifications/delete/:id", controllers.DeleteCertifications)

	// ~~~ Categories ~~~ //
	categories := r.Group("/categories")
	categories.POST("/store", controllers.StoreCategories)
	categories.POST("/update", controllers.UpdateCategories)
	categories.GET("/index", controllers.IndexCategories)

	// ~~~ Main Controller //
	main := r.Group("/main")
	main.GET("/index", controllers.MainIndex)
	main.GET("/index/without_auth", controllers.IndexWithoutAuth)
	main.GET("/generate_cv/:id", controllers.GenerateCv)

	// AdminPanel
	admin := r.Group("/admin")
	admin.POST("/templates/store", controllers.StoreTemplate)

	// --------- Run ------- //

	r.Run(":8082")

}
