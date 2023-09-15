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
	BD := e.Group("/BD")
	RL := e.Group("/RL")

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
	mn.DELETE("/delete-menu", controllers.DeleteMenu)
	//upload_foto_menu
	mn.POST("/upload-foto-menu", controllers.UploadFotoMenu)

	//Pengantar
	//Sign-UP-Pengantar
	PA.POST("/su-pengantar", controllers.SignUpPengantar)
	//Read-Pengantar
	PA.GET("/read-pengantar", controllers.ReadPengantar)
	//Update-Maps-Pengantar
	PA.PUT("/update-maps-pengantar", controllers.UpdateMaps)
	//Read-Maps-Pengantar
	PA.GET("/read-maps-pengantar", controllers.ReadMapsPengantar)

	//Order
	//input
	ORD.POST("/input-order", controllers.InputOrder)
	//Read_Order
	ORD.GET("/read-order", controllers.ReadOrder)
	//Read_Detail_order
	ORD.GET("/read-detail-order", controllers.ReadDetailOrder)
	//Read_Order_Menu
	ORD.GET("/show-order-menu", controllers.ShowOrderMenu)
	//Set_Pengantar
	ORD.PUT("/set-pengantar", controllers.SetPegantar)
	//Set_Pengantar
	ORD.PUT("/confirm-order", controllers.ConfirmOrder)
	//Order_Detail_User
	ORD.GET("/order_detail_user", controllers.OrderDetailUser)

	//Pembayaran
	//Read_Pembayaran
	PBR.GET("/read-pembayaran", controllers.ReadPembayaran)
	//Upload_Foto_Pembayaran
	PBR.POST("/upload-foto", controllers.UploadFotoPembayaran)
	//Confirm_Pembayaran
	PBR.PUT("/confirm-pembayaran", controllers.ConfirmPembayaran)

	//foto
	//get image foto
	e.GET("/read-foto", controllers.ReadFoto)

	//MAPS
	//Input_Maps
	MP.POST("/input-maps", controllers.InputMaps)
	//Read_Maps
	MP.GET("/read-maps", controllers.ReadMaps)

	//Budgeting
	//Input_Budgeting
	BD.POST("/input-budgeting", controllers.InputBudgeting)
	//Read_Awal_Budgeting
	BD.GET("/read-awal-budgeting", controllers.ReadAwalBudgeting)
	//Read_Budgeting
	BD.GET("/read-budgeting", controllers.ReadBudgeting)

	//Realisasi
	//Input_Realisai
	RL.POST("/input-realisasi", controllers.InputRealisasi)
	//Read_Realisasi
	RL.GET("/read-realisasi", controllers.ReadRealisasi)

	//Rating

	return e
}
