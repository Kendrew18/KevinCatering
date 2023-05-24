package routes

import (
	"KevinCatering/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Project-NDL")
	})

	user := e.Group("/user")
	cat := e.Group("/cat")
	mn := e.Group("/mn")
	PA := e.Group("/PA")
	ORD := e.Group("/ORD")
	PBR := e.Group("/PBR")
	MP := e.Group("/MP")

	//user
	//Sign_UP
	user.POST("/sign-up", controllers.SignUP)
	//Login
	user.GET("/login", controllers.LoginUser)
	//Read_profile
	user.GET("/get-profile", controllers.Read_Profile)
	//Edit_profile
	user.PUT("/edit-profile", controllers.EditProfile)

	//Catering
	cat.POST("/input-catering", controllers.InputCatering)
	//Read_Profile_Catering
	cat.GET("/read-profile-cat", controllers.ReadProfileCatering)
	//Read_Catering
	cat.GET("/read-cat", controllers.ReadCatering)
	//Read_Awal_Catering
	cat.GET("/read-awal-cat", controllers.ReadAwalCatering)
	//Edit_Profile_Catering
	cat.PUT("/edit-prof-cat", controllers.EditProfileCatering)

	//Menu
	//input_menu
	mn.POST("/input-menu", controllers.InputMenu)
	//read_menu
	mn.GET("/read-menu", controllers.ReadMenu)
	//Edit_menu
	mn.PUT("/update-menu", controllers.EditMenu)
	//Delete_menu
	mn.PUT("/Delete_Menu", controllers.DeleteMenu)

	//Pengantar
	//Sign-UP-Penganta
	PA.POST("/su-pengantar", controllers.SignUpPengantar)
	//Read-Pengantar
	PA.GET("/read-pengantar", controllers.ReadPengantar)
	//Update-Maps-Pengantar
	PA.PUT("/update-maps-pengantar", controllers.UpdateMaps)
	//Read-Maps-Pengantar
	PA.PUT("/read-maps-pengantar", controllers.ReadMapsPengantar)

	//Order
	//input
	ORD.POST("/input-order", controllers.InputOrder)
	//Read_Order
	ORD.GET("/read-order", controllers.ReadOrder)
	//Read_Detail_order
	ORD.GET("/read-detail-order", controllers.ReadDetailOrder)

	//Pembayaran
	//Read_Pembayaran
	PBR.GET("/read-pembayaran", controllers.ReadPembayaran)
	//Upload_Foto_Pembayaran
	PBR.POST("/upload-foto", controllers.UploadFotoPembayaran)

	//foto
	//get image foto
	PBR.GET("/read-foto", controllers.ReadFoto)

	//MAPS
	MP.POST("/input-maps", controllers.InputMaps)

	MP.GET("/read-maps", controllers.ReadMaps)

	//Rating

	return e
}
