package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thirumalaivasan95/bookstore_users-api/domain/users"
	"github.com/thirumalaivasan95/bookstore_users-api/services"
	"github.com/thirumalaivasan95/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO: handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	//TODO: handle json error
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		// fmt.Println(err)
		//TODO: return bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: user creation error
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
	// fmt.Println(string(bytes))
	// fmt.Println(err)
	// c.String(http.StatusNotImplemented, "Implement me!\n\n")
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User ID should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
	// c.String(http.StatusNotImplemented, "Implement me!\n\n")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!\n\n")
}
