package routes

import (
	auth "Diskusiaza-BE/app/http/controllers"
	profile "Diskusiaza-BE/app/http/controllers/profile"
	therad "Diskusiaza-BE/app/http/controllers/therad"
	"Diskusiaza-BE/constants"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

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
	withToken := e.Group("")
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

	//-------------------------------------------
	// THERAD GROUP
	//--------------------------------------------
	// get all kategori therad
	withToken.GET("/therad/kategori", therad.GetKategoriTheradController)

	return e
}
