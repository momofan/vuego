package router

import (
   "demo/handlers"
   "github.com/gin-gonic/gin"
)
         
func Init() {
   // Creates a default gin router
   r := gin.Default()    // Grouping routes
   // group： v1
   v1 := r.Group("/v1")
   {
            v1.GET("/hello", handlers.HelloPage)
   }
   r.Run(":8000") // listen and serve on 0.0.0.0:8000
}