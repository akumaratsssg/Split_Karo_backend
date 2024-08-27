package routes

import (
	"Expense_Management/backend/controllers"
	"Expense_Management/backend/middlewares"
	"Expense_Management/backend/repositories"
	"Expense_Management/backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(middlewares.Logger())

	// Create the GroupRepository
	groupRepository := repositories.NewGroupRepository(db)

	// Initialize services and controllers
	groupService := services.NewGroupService(groupRepository)
	groupController := controllers.NewGroupController(groupService)

	auth := router.Group("/auth")
	{
		auth.GET("/", controllers.Home)
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/logout", controllers.Logout)
		auth.GET("/search_user", middlewares.JWTAuthentication(), controllers.SearchUserByUsername)
	}

	u_group := router.Group("/group")
	{
		u_group.POST("/create_group", middlewares.JWTAuthentication(), controllers.Group)
		u_group.GET("/get_groups", middlewares.JWTAuthentication(), groupController.GetGroups)
		u_group.POST("/add_member", middlewares.JWTAuthentication(), controllers.AddMemberByGroupNameAndEmail)
		u_group.DELETE("/remove_member", middlewares.JWTAuthentication(), controllers.RemoveMemberByGroupNameAndEmail)
		u_group.GET("/get_members", middlewares.JWTAuthentication(), controllers.GetMembersByGroupName)

	}

	exp := router.Group("/expense")
	{

		exp.POST("create_expense", middlewares.JWTAuthentication(), controllers.AddExpenseToGroup)
		exp.POST("create_balance", middlewares.JWTAuthentication(), controllers.CreateBalance)
		exp.GET("get_balance", middlewares.JWTAuthentication(), controllers.GetGroupBalances)
	}

	return router
}
