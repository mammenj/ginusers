package main

import (
	"log"

	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"github.com/mammenj/ginusers/users/config"
	"github.com/mammenj/ginusers/users/controllers"
	"github.com/mammenj/ginusers/users/security"

	"github.com/casbin/casbin"
	//"github.com/gin-contrib/authz"
)

func main() {
	log.Println("Starting server...")
	r := gin.Default()
	e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv", true)

	v1 := r.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", controllers.GetUsers)
			users.PUT("/:id", controllers.UpdateUser)
			users.GET("/:id", controllers.GetUser)
			users.POST("/", controllers.CreateUser)
		}
	}

	public := r.Group("/login")
	public.POST("/", security.Login)
	public.Use(security.NewJwtAuthorizer(e))
	private := r.Group("/v1/auth")
	{
		users := private.Group("/users")
		{
			config, err := config.GetConfiguration("config.json")
			if err != nil {
				log.Fatal(err)
			}
			users.Use(jwt.Auth(config.Jwtsecret))
			users.Use(security.NewJwtAuthorizer(e))
			users.GET("/", controllers.GetUsers)
			users.DELETE("/:id", controllers.DeleteUser)
			users.POST("/", controllers.CreateUser)
			users.PUT("/:id", controllers.UpdateUser)
		}
	}

	r.RunTLS(":8443", "interns2019.pem", "interns2019.key")
}
