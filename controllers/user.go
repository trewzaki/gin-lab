package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/trewzaki/gin-lab/models"
)

func CreateUser(c *gin.Context) {
	userModel := models.User{}
	c.ShouldBindWith(&userModel, binding.JSON)

	if createUserErr := models.DB.Create(&userModel).Error; createUserErr != nil {
		resMap := map[string]interface{}{
			"success": false,
			"message": "Create user error.",
		}

		resByte, _ := json.Marshal(resMap)

		c.Data(200, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	}

	resMap := map[string]interface{}{
		"success": true,
	}

	resByte, _ := json.Marshal(resMap)

	c.Data(200, "application/json; charset=utf-8", resByte)
	c.Abort()

	return
}

func commandExample() {
	/* Start Map Example */
	mapExample := map[string]interface{}{
		"stringja": "eueu",
		"intja":    2,
	}

	tmpString := mapExample["stringja"].(string)
	tmpInt := mapExample["intja"].(int)
	fmt.Println("tmpString: ", tmpString)

	fmt.Println("tmpInt: ", tmpInt)
	/* End Map Example */

	/* Start Struct Example */
	userModel := models.User{}

	userModel.FirstName = "Tananut"
	userModel.LastName = "Panyagosa"
	userModel.Email = "tongtananut@gmail.com"

	fmt.Println("userModel: ", userModel)
	/* End Struct Example */

	/* Start Marshal Example */
	mapExample2 := map[string]interface{}{
		"first_name": "Sam",
		"last_name":  "Za",
		"email":      "samza55plus@gmail.com",
	}

	tmpByte, marshalErr := json.Marshal(mapExample2)
	if marshalErr != nil {
		fmt.Println("marshal err: ", marshalErr)
	}
	/* End Marshal Example */

	/* Start Unmarshal Example */
	inputMap := map[string]interface{}{}
	json.Unmarshal(tmpByte, &inputMap)

	userModel2 := models.User{}
	json.Unmarshal(tmpByte, &userModel2)

	fmt.Println("tmpByte: ", string(tmpByte))
	fmt.Println("inputMap: ", inputMap)
	fmt.Println("inputMap: ", userModel2)
	/* End Unmarshal Example */
}
