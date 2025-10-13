package router

import (
	model "BockWeb/Model"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const jwtSecurityStr string = "BockWeb2Login8Security0String5"

func jwtString(id int, uname string) (tokenStr string, err error) {
	clains := jwt.StandardClaims{
		Issuer:    uname,
		Id:        strconv.Itoa(id),
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Unix() + 60*60*24,
	}
	tobj := jwt.NewWithClaims(jwt.SigningMethodHS256, clains)
	tokenStr, err = tobj.SignedString([]byte(jwtSecurityStr))
	return
}

func jwtParse(tokenS string) (uname *string, uid *string) {
	token, e := jwt.ParseWithClaims(tokenS, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecurityStr), nil
	})
	if e != nil {
		return
	}
	data := token.Claims.(*jwt.StandardClaims)
	return &data.Issuer, &data.Id
}

func cookieTokenVerify(ctx *gin.Context) (uname *string, uid uint) {
	tStr, e := ctx.Cookie("token")
	if e != nil || tStr == "" {
		return
	}
	uname, uidStr := jwtParse(tStr)
	id, _ := strconv.ParseUint(*uidStr, 10, 64)
	uid = uint(id)
	return
}

func register(ctx *gin.Context) {
	user := model.User{
		Name:      ctx.PostForm("name"),
		Password:  ctx.PostForm("password"),
		Email:     ctx.PostForm("email"),
		PostCount: 0,
	}

	if user.Name == "" || user.Password == "" || user.Email == "" {
		ctx.String(http.StatusOK, "数据异常")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.String(http.StatusBadRequest, "密码字符异常")
		return
	}
	user.Password = string(hashedPassword)
	if user.Register() {
		ctx.String(http.StatusOK, "成功注册")
	} else {
		ctx.String(http.StatusOK, "用户已存在")
	}
}

func loginVerify(ctx *gin.Context) {
	user := model.User{Name: ctx.PostForm("user")}
	user.LoginVerify()
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ctx.PostForm("psd")))
	if err != nil {
		ctx.String(http.StatusBadRequest, "用户名或密码错误")
		return
	}
	tokenS, err := jwtString(int(user.ID), user.Name)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "登录数据异常")
	} else {
		ctx.SetCookie("token", tokenS, 20000, "/home", "", false, true)
		ctx.Redirect(http.StatusPermanentRedirect, "/home")
	}
}

func directLogin(ctx *gin.Context) {
	c, e := ctx.Cookie("token")
	if e != nil || c == "" {
		ctx.Redirect(http.StatusPermanentRedirect, "static/login.html")
		return
	}
	user, _ := jwtParse(c)
	if user == nil {
		ctx.SetCookie("token", "", 0, "/", "/home", false, true)
		ctx.Redirect(http.StatusPermanentRedirect, "static/login.html")
		return
	}
	u := model.User{Name: *user}
	if !u.Exist() {
		ctx.Redirect(http.StatusFound, "/static/login.html")
		return
	}
	posts := u.GetPosts()
	ctx.HTML(http.StatusOK, "home.html", map[string]interface{}{
		"Name":  *user,
		"Posts": posts,
		"Home":  "true",
	})
	//ctx.String(http.StatusOK, *user+"的主页")
}

func userRouter(ser *gin.Engine) {
	//用户注册
	ser.POST("/register", register)
	//用户登录
	ser.POST("/loginVerify", loginVerify)

	//直接访问主页
	ser.GET("/home", directLogin)
	ser.POST("/home", directLogin)

}
