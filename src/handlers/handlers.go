package handlers

/*func Registration(context *gin.Context) {
	var user *models.User

	decode := json.NewDecoder(context.Request.Body).Decode(&user)
	if decode != nil {
		context.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
		return
	}

	isExist := database.IsExistUserByLogin(user.Login)
	if isExist {
		context.JSON(http.StatusOK, gin.H{
			"response": "User with this login already exists.",
		})
	} else {
		user := &models.User{ID: user.ID, Login: user.Login, Password: user.Password}

		isSuccessAdd := database.Add(user)
		if isSuccessAdd != nil {
			context.JSON(http.StatusOK, gin.H{
				"response": "Something go wrong. Please, try again later.",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"response": gin.H{
					"code":    http.StatusOK,
					"message": "Successfull.",
				},
			})
		}
	}
}*/

/*func Login(context *gin.Context) {

}*/
