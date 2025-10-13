# README
项目使用vscode打开Task4文件夹 运行  选择Launch Package项进行调试

dbConfig.txt    中配置数据库连接信息
日志记录在ginlog文件中
static目录下是前端页面模板

 登录页面
localhost:8080/static/login.html

注册页面
localhost:8080/static/register.html

用户主页 页面中有文章的增删改查功能
localhost:8080/home

展示指定:id的文章内容及评论  在页面中可对文章进行评论
localhost:8080/post/:id

展示用户:username对应的文章列表
localhost:8080/Posts/:username  