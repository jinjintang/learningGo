// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"qingguo-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/advertise",
			beego.NSInclude(
				&controllers.AdvertiseController{},
			),
		),

		beego.NSNamespace("/area",
			beego.NSInclude(
				&controllers.AreaController{},
			),
		),

		beego.NSNamespace("/attendance",
			beego.NSInclude(
				&controllers.AttendanceController{},
			),
		),

		beego.NSNamespace("/classs",
			beego.NSInclude(
				&controllers.ClasssController{},
			),
		),

		beego.NSNamespace("/course",
			beego.NSInclude(
				&controllers.CourseController{},
			),
		),

		beego.NSNamespace("/course_advertise",
			beego.NSInclude(
				&controllers.CourseAdvertiseController{},
			),
		),

		beego.NSNamespace("/course_category",
			beego.NSInclude(
				&controllers.CourseCategoryController{},
			),
		),

		beego.NSNamespace("/course_rank",
			beego.NSInclude(
				&controllers.CourseRankController{},
			),
		),

		beego.NSNamespace("/refund",
			beego.NSInclude(
				&controllers.RefundController{},
			),
		),

		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{},
			),
		),

		beego.NSNamespace("/student",
			beego.NSInclude(
				&controllers.StudentController{},
			),
		),

		beego.NSNamespace("/suser",
			beego.NSInclude(
				&controllers.SuserController{},
			),
		),

		beego.NSNamespace("/teacher_course",
			beego.NSInclude(
				&controllers.TeacherCourseController{},
			),
		),

		beego.NSNamespace("/trade",
			beego.NSInclude(
				&controllers.TradeController{},
			),
		),

		beego.NSNamespace("/trade_detail",
			beego.NSInclude(
				&controllers.TradeDetailController{},
			),
		),

		beego.NSNamespace("/transfer",
			beego.NSInclude(
				&controllers.TransferController{},
			),
		),

		beego.NSNamespace("/user_role",
			beego.NSInclude(
				&controllers.UserRoleController{},
			),
		),

		beego.NSNamespace("/withdraw",
			beego.NSInclude(
				&controllers.WithdrawController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
