// Package routes (Setup Routes Group)
package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

// Setup >>>
func Setup() {
	r := gin.Default()

	// Config
	r.LoadHTMLGlob("assets/pdf_templates/*")

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
	cvs.GET("/basics_information/image/serving/:img", controllers.ServingBasicsInformationsImage)
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
	main.GET("/image/:img", controllers.ServingImage)
	main.GET("/generate_cv/:id", controllers.GenerateCv)
	main.GET("/serving_cv/:id", controllers.ServingCV)
	main.GET("/servingHtml", controllers.ServingHTML)
	main.GET("/serving_caption_html/:img", controllers.ServingTemplateCaption)

	// AdminPanel
	admin := r.Group("/admin")
	admin.POST("/templates/store", controllers.StoreTemplate)

	// --------- Run ------- //
	r.Run(":8082")

}
