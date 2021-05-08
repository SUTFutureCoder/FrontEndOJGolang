package v1

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/routers/api/v1/admin/contest"
	"FrontEndOJGolang/routers/api/v1/admin/contest_lab"
	"FrontEndOJGolang/routers/api/v1/admin/contest_user"
	"FrontEndOJGolang/routers/api/v1/admin/lab"
	"FrontEndOJGolang/routers/api/v1/admin/lab_testcase"
	"FrontEndOJGolang/routers/api/v1/admin/testfield"
	"FrontEndOJGolang/routers/api/v1/admin/user"
	"github.com/gin-gonic/gin"
)

func InitAdminRouter(r *gin.Engine) {

	// 测试区
	admin := r.Group("/admin")

	// 前置权限检测
	admin.Use(app.CheckUserAdmin)

	test := admin.Group("/test")
	test.POST("/screenshot", testfield.ScreenShot)

	// admin管理区
	adminLab := admin.Group("/lab")
	adminLab.POST("/add", lab.AddLab)
	adminLab.POST("/modify", lab.ModifyLab)
	adminLab.POST("/enable", lab.EnableLab)
	adminLab.POST("/disable", lab.DisableLab)
	adminLab.POST("/constructing", lab.ConstructingLab)

	labTestcaseGroup := admin.Group("/lab_testcase")
	labTestcaseGroup.POST("/set", lab_testcase.Set)
	labTestcaseGroup.POST("/add", lab_testcase.Add)
	labTestcaseGroup.POST("/list", lab_testcase.List)
	labTestcaseGroup.POST("/modify", lab_testcase.Modify)
	labTestcaseGroup.POST("/run", lab_testcase.Run)

	// user管理区
	userGroup := admin.Group("/user")
	userGroup.POST("/disable", user.DisableUser)
	userGroup.POST("/enable", user.EnableUser)
	userGroup.POST("/list", user.List)
	userGroup.POST("/create", user.CreateUser)
	userGroup.POST("/change_pwd", user.ChangePasswd)
	userGroup.POST("/grant", user.GrantPermission)
	userGroup.POST("/modify", user.ModifyUser)

	// contest管理区
	contestGroup := admin.Group("/contest")
	contestGroup.POST("/create", contest.CreateContest)
	contestGroup.POST("/modify", contest.ModifyContest)
	contestGroup.POST("/status_modify", contest.ModifyContestStatus)

	contestGroup.POST("/get_labs", contest_lab.GetLabs)
	contestGroup.POST("/manage_labs", contest_lab.ManageLabs)

	contestGroup.POST("/add_users", contest_user.AddUsers)
	contestGroup.POST("/users_status_modify", contest_user.ModifyUsersStatus)


}
