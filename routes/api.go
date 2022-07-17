package routes

import (
	auth "Diskusiaza-BE/app/http/controllers"
	"Diskusiaza-BE/app/http/controllers/follow"
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
		AllowOrigins:     []string{"https://diskusiaza.netlify.app", "*", "http://localhost:3000"},
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
	withToken.GET("/user/profile", profile.GetUsersByTokenController)
	// update profile
	withToken.PUT("/user/profile", profile.UpdateUsersDataController)

	//---------------------------------------------
	// USER CREATE THERAD NORMAL
	//---------------------------------------------
	// user create therad (question)
	withToken.POST("/user/therad", therad.CreateTheradNormalController)
	// user create therad in ruang
	withToken.POST("/user/therad/ruang", therad.CreateTheradInRuangController)
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
	// get all therad by user id
	withToken.GET("/therad/user/:user_id", therad.GetListTheradNormalByUsersIdController)
	// get search therad
	withToken.GET("/therad/search", therad.GetSearchTheradController)
	// get therad by kategori
	withToken.GET("/therad/kategori/:kategori_id", therad.GetTheradByKategoriIdController)
	// get favorite therad
	withToken.GET("/therad/trending", therad.GetTrendingTheradController)
	// get by id therad
	withToken.GET("/therad/:therad_id", therad.GetByIdTheradController)
	// like therad
	withToken.POST("/therad/like", therad.LikeTheradController)
	// unlike therad
	withToken.DELETE("/therad/like", therad.UnLikeTheradController)

	// ------------------
	// USER GROUP
	//-------------------
	// get all users
	withToken.GET("/user", follow.GetAllUsersController)
	// get users detail by id
	withToken.GET("/user/search/:user_id", follow.GetUsersByIdController)
	// search users
	withToken.GET("/user/search", follow.GetSearchUsersByNameController)
	// get all followers
	withToken.GET("/user/followers", follow.GetDetailFollowersController)
	// get all followed
	withToken.GET("/user/followed", follow.GetDetailFollowedController)

	//---------------------
	// FOLLOW USERS
	//---------------------
	// follow users
	withToken.POST("/follow", follow.FollowedUserController)
	// unfollow users
	withToken.DELETE("/follow", follow.UnFollowedUserController)
	// remove followers
	withToken.DELETE("/followers", follow.RemoveFollowersController)

	//-----------------------
	// RUANG
	//----------------------
	// create ruang
	withToken.POST("/ruang", ruang.CreateRuangController)
	// delete ruang
	withToken.DELETE("/ruang/:ruang_id", ruang.DeleteRuangController)
	// update ruang
	withToken.PUT("/ruang/:ruang_id", ruang.UpdateRuangController)
	// get ruang by user
	withToken.GET("/ruang/user", ruang.GetRuangByUserIdController)
	// get all ruang
	withToken.GET("/ruang", ruang.GetAllRuangController)
	// Get detail ruang by id
	withToken.GET("/ruang/detail/:ruang_id", ruang.GetRuangByIdController)
	// search ruang
	withToken.GET("/ruang/search", ruang.GetSearchRuangController)
	// get all member in ruang
	withToken.GET("/ruang/member/:ruang_id", ruang.GetAllMemberInRuangController)
	// get all therad in ruang
	withToken.GET("/ruang/therad/:ruang_id", therad.GetAllTheradInRuangController)

	//------------------------
	// USER JOIN TO RUANG
	//------------------------
	// join to ruang
	withToken.POST("/ruang/join", ruang.CreateMemberRuangController)
	// leave from ruang
	withToken.DELETE("/ruang/leave/:ruang_id", ruang.DeleteMemberRuangController)

	return e
}
