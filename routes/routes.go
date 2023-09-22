package routes

import (
	"KevinCatering/controllers"
	"KevinCatering/controllers/Rating"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Skripsi Kevin Catering")
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
	RT := e.Group("/RT")

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
	//Get_QR_Catering
	cat.GET("/get-QR-catering", controllers.GetQRCatering)

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
	PA.PUT("/update-Maps-pengantar", controllers.UpdateMaps)
	//Read-Maps-Pengantar
	PA.GET("/read-Maps-pengantar", controllers.ReadMapsPengantar)

	//Order
	//input
	ORD.POST("/input-order", controllers.InputOrder)
	//Read_Order_Menu
	ORD.GET("/show-order-menu", controllers.ShowOrderMenu)
	//Set_Pengantar
	ORD.PUT("/set-pengantar", controllers.SetPegantar)
	//Confirm_Order
	ORD.PUT("/confirm-order", controllers.ConfirmOrder)
	//Order_Detail_User
	ORD.GET("/order-detail-user", controllers.OrderDetailUser)
	//History_Order
	ORD.GET("/history-order", controllers.HistoryOrder)

	//Pembayaran
	//Read_Pembayaran
	PBR.GET("/read-pembayaran", controllers.ReadPembayaran)
	//Upload_Foto_Pembayaran
	PBR.POST("/upload-foto", controllers.UploadFotoPembayaran)
	//Confirm_Pembayaran
	PBR.PUT("/confirm-pembayaran", controllers.ConfirmPembayaran)
	//Read_Recipe_Order
	PBR.GET("/read-recipe-order", controllers.ReadOrderRecipe)
	//Read_Detail_Rescipe_order
	PBR.GET("/read-detail-rescipe-order", controllers.ReadDetailOrderRecipe)
	//Read_Notif_Pembayaran
	PBR.GET("/read-notif-pembayaran", controllers.NotifPembayaran)

	//foto
	//get image foto
	e.GET("/read-foto", controllers.ReadFoto)

	//MAPS
	//Input_Maps
	MP.POST("/input-Maps", controllers.InputMaps)
	//Read_Maps
	MP.GET("/read-Maps", controllers.ReadMaps)

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
	RT.POST("/input-rating", Rating.InputRating)

	return e
}
