package controllers 
 
import ( 
         "github.com/gin-gonic/gin" 
         "net/http" 
) 
func HelloPage(c *gin.Context) { 
         c.JSON(http.StatusOK, gin.H{ 
                  "message": "welcome", 
         }) 
} 