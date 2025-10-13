package router

import (
	model "BockWeb/Model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func publishPost(ctx *gin.Context) {
	uname, uid := cookieTokenVerify(ctx)
	if uname != nil && uid > 0 {
		tiltle := ctx.PostForm("title")
		content := ctx.PostForm("content")
		if tiltle == "" || content == "" {
			ctx.String(http.StatusBadRequest, "标题或文章不可为空")
			return
		}
		post := model.Post{Title: tiltle, Content: content, UserID: uid, CmtCount: 0}
		if e := post.Publish(); e != nil {
			ctx.String(http.StatusInternalServerError, e.Error())
			return
		}
		ctx.Redirect(http.StatusFound, "/home")
		return
	}
	ctx.String(http.StatusBadRequest, "权限不足")
}

func postDisplay(ctx *gin.Context) {
	var post_id int64
	if !intParamCheck(ctx, Param, "id", &post_id) {
		ctx.String(http.StatusBadRequest, "无效参数")
		return
	}
	p := model.Post{}
	p.ID = uint(post_id)
	var author string
	exist, err := p.LoadWithCommnets(&author)
	if !exist || err != nil {
		ctx.String(http.StatusBadRequest, "数据异常:"+err.Error())
		return
	}
	cmtInfos := make([]model.CmtInfo, 0, len(p.Comments))
	for _, v := range p.Comments {
		cmtInfos = append(cmtInfos, v.CreateInfor())
	}
	ctx.HTML(http.StatusOK, "displayPost.html", map[string]interface{}{
		"Title":    p.Title,
		"Content":  p.Content,
		"Author":   author,
		"Comments": cmtInfos,
		"PostId":   post_id,
	})
}

func editPost(ctx *gin.Context) {
	_, uid := cookieTokenVerify(ctx)
	var post_id int64
	if !intParamCheck(ctx, Param, "id", &post_id) {
		ctx.String(http.StatusBadRequest, "未知参数")
		return
	}

	post := model.Post{}
	post.ID = uint(post_id)
	post.LoadData()
	if post.UserID != uid {
		ctx.String(http.StatusBadRequest, "权限不足")
		return
	}

	ctx.HTML(http.StatusOK, "editPost.html", map[string]interface{}{
		"Title":   post.Title,
		"Content": post.Content,
		"PostId":  post.ID,
	})
}

func updatePost(ctx *gin.Context) {
	uname, uid := cookieTokenVerify(ctx)
	var id int64
	if !intParamCheck(ctx, PostForm, "post_id", &id) {
		ctx.String(http.StatusBadRequest, "post_id参数异常")
		return
	}
	if uname != nil && uid > 0 {
		c := ctx.PostForm("content")
		p := model.Post{Content: c, UserID: uid}
		p.ID = uint(id)
		if p.UpdateContent() { //文章归属正确 更新成功
			ctx.String(http.StatusOK, "已保存")
			return
		}
	}
	ctx.String(http.StatusNotFound, "权限不足")
}

func deletePost(ctx *gin.Context) {
	_, uid := cookieTokenVerify(ctx)
	idStr := ctx.Param("id")
	id, e := strconv.ParseUint(idStr, 10, 64)
	if e != nil {
		ctx.String(http.StatusBadRequest, "参数异常")
		return
	}
	p := model.Post{UserID: uid}
	p.ID = uint(id)
	if p.Delete() {
		ctx.String(http.StatusOK, "已删除")
		return
	}
	ctx.String(http.StatusNotFound, "权限不足")
}

func comment(ctx *gin.Context) {
	uname, uid := cookieTokenVerify(ctx)
	var post_id int64
	if !intParamCheck(ctx, Param, "post_id", &post_id) {
		ctx.String(http.StatusBadRequest, "参数异常")
		return
	}
	if uname != nil && uid > 0 {
		cntStr := ctx.PostForm("content")
		if cntStr == "" {
			ctx.String(http.StatusBadRequest, "评论不能为空")
			return
		}
		cnt := model.Comment{Content: cntStr, PostID: uint(post_id), UserID: uid}
		cnt.Publish()
		ctx.Redirect(http.StatusFound, "/post/"+strconv.Itoa(int(post_id)))
		return
	}
	ctx.String(http.StatusBadRequest, "权限不足")
}

func postsListDisplay(ctx *gin.Context) {
	user := model.User{Name: ctx.Param("uname")}
	posts := user.GetPosts()
	ctx.HTML(http.StatusOK, "postList.html", map[string]interface{}{
		"Name":  user.Name,
		"Posts": posts,
		"Home":  "false",
	})
}

func postRouter(ser *gin.Engine) {
	//发布文章
	ser.POST("/home/publishPost", publishPost)

	//浏览目标文章及评论
	ser.GET("/post/:id", postDisplay)

	//编辑修改目标文章
	ser.POST("/home/editPost/:id", editPost)

	//保存更新目标文章
	ser.POST("/home/updatePost", updatePost)

	//删除文章
	ser.POST("/home/deletePost/:id", deletePost)

	//对发表发表评论
	ser.POST("/home/comment/:post_id", comment)

	//文章列表
	ser.GET("/Posts/:uname", postsListDisplay)

}
