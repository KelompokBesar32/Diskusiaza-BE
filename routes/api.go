package routes

import (
	auth "Diskusiaza-BE/app/http/controllers"
	profile "Diskusiaza-BE/app/http/controllers/profile"
	ruang "Diskusiaza-BE/app/http/controllers/ruang"
	therad "Diskusiaza-BE/app/http/controllers/therad"
	"Diskusiaza-BE/constants"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://diskusiaza.netlify.app"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))
	// application routes
	e.GET("/", auth.TestController)

	//---------------------------------
	// AUTH GROUP
	//--------------------------------
	e.POST("/auth/register", auth.RegisterController)
	e.POST("/auth/login", auth.LoginController)

	//--------------------------------
	// WIDTH MIDDLEWARE
	//--------------------------------
	withToken := e.Group("t")
	withToken.Use(middleware.JWT([]byte(constants.ScreetJwtForUser)))
	//------------------------------------------------------
	// logout
	withToken.GET("/auth/logout", auth.LogoutController)

	//-----------------------------------------------------
	// PROFILE GROUP
	//----------------------------------------------------
	// get users profile
	withToken.GET("/user/profile", profile.GetUsersByToken)
	// update profile
	withToken.PUT("/user/profile", profile.UpdateUsersData)

	//---------------------------------------------
	// USER CREATE THERAD NORMAL
	//---------------------------------------------
	// user create therad (question)
	withToken.POST("/user/therad", therad.CreateTheradNormalController)
	// get therad by user
	withToken.GET("/user/therad", therad.GetListTheradNormalByUsersController)
	// delete therad
	withToken.DELETE("/user/therad/:therad_id", therad.DeleteTheradController)
	// update therad
	withToken.PUT("/user/therad/:therad_id", therad.UpdateTheradNormalController)

	//-------------------------------------------
	// THERAD GROUP
	//--------------------------------------------
	// get all kategori therad
	withToken.GET("/therad/kategori", therad.GetKategoriTheradController)
	// get all therad
	withToken.GET("/therad", therad.GetAllTheradController)
	// get by id therad
	withToken.GET("/therad/:therad_id", therad.GetByIdTheradController)

	//-----------------------
	// ruang
	//----------------------
	// Create Ruang
	withToken.POST("/user/ruang", ruang.CreateRuangController)
	// Delete Ruang
	withToken.DELETE("/user/ruang", ruang.DeleteRuangController)
	// Update Ruang
	withToken.POST("/user/ruang", ruang.EditRuangController)
	// get Ruang
	withToken.GET("/user/ruang", ruang.GetRuangController)
	return e
}
