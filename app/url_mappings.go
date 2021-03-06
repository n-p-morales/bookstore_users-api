package app

import(
  "github.com/n-p-morales/bookstore_users-api/controllers/users"
  "github.com/n-p-morales/bookstore_users-api/controllers/ping"
)

func mapUrls()  {
  router.GET("/ping", ping.Ping)
  router.GET("/users/:user_id", users.GetUser)
  router.POST("/users", users.CreateUser)
}
