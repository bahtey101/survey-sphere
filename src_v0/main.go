package src

import (
	"net/http"
	"src/controllers"
	"src/database"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	database.Init()

	database.CreateTables()

	var new_user = database.User{Login: "NikitOS", Password: "Dimon4eg"}
	database.CreateUser(&new_user)
	database.GetUser(&new_user)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.GET("/users", controllers.FindUsers)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.FindUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.Run()
	// handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	// 	resp := []byte(`{"status": "ok"}`)
	// 	rw.Header().Set("Content-Type", "application/json")
	// 	rw.Header().Set("Content-Length", fmt.Sprint(len(resp)))
	// 	rw.Write(resp)
	// })

	// log.Println("Server is available at http://localhost:8000")
	// log.Fatal(http.ListenAndServe(":8000", handler))
	// tests.TEST_1()

	//database.DeleteUser(user1)
	//database.DeleteUser(user2)
}

func loginHandler(c *gin.Context) {
	// Здесь вы можете выполнить проверку имени пользователя и пароля
	// В нашем примере мы предполагаем, что пользователь успешно прошел аутентификацию

	// Генерируем токен JWT
	token := generateToken()

	// Возвращаем токен как ответ
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func generateToken() string {
	// Генерация токена JWT
	// В реальном приложении вам нужно будет адаптировать это под свои нужды
	// Здесь приведен простейший пример
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	tokenString, _ := token.SignedString([]byte("your_secret_key"))
	return tokenString
}
