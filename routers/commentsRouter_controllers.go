package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["qingguo-api/controllers:AdvertiseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AdvertiseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AdvertiseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AdvertiseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AdvertiseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AdvertiseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AdvertiseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AdvertiseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AdvertiseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AdvertiseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AreaController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AreaController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AreaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AreaController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AreaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AreaController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AreaController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AttendanceController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AttendanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AttendanceController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AttendanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AttendanceController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AttendanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AttendanceController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AttendanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:AttendanceController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:AttendanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:ClasssController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:ClasssController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:ClasssController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:ClasssController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:ClasssController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:ClasssController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:ClasssController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:ClasssController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:ClasssController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:ClasssController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseAdvertiseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseAdvertiseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseAdvertiseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseAdvertiseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseAdvertiseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseAdvertiseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseAdvertiseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseAdvertiseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseAdvertiseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseAdvertiseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseCategoryController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseCategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseCategoryController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseCategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseCategoryController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseCategoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseCategoryController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseCategoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseCategoryController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseCategoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseRankController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseRankController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseRankController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseRankController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseRankController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseRankController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseRankController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseRankController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:CourseRankController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:CourseRankController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:RefundController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:RefundController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:RefundController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:RefundController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:RefundController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:RefundController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:RefundController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:RefundController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:RefundController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:RefundController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:RoleController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:RoleController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:RoleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:RoleController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:RoleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:RoleController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:RoleController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:SuserController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:SuserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:SuserController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:SuserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:SuserController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:SuserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:SuserController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:SuserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:SuserController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:SuserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TeacherCourseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TeacherCourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TeacherCourseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TeacherCourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TeacherCourseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TeacherCourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TeacherCourseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TeacherCourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TeacherCourseController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TeacherCourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TradeController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TradeController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TradeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TradeController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TradeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TradeController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TradeController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TradeDetailController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TradeDetailController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TradeDetailController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TradeDetailController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TradeDetailController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TradeDetailController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TradeDetailController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TradeDetailController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TradeDetailController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TradeDetailController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TransferController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TransferController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TransferController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TransferController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TransferController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TransferController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TransferController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TransferController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:TransferController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:TransferController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:UserRoleController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:UserRoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:UserRoleController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:UserRoleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:UserRoleController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:UserRoleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:UserRoleController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:UserRoleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:UserRoleController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:UserRoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:WithdrawController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:WithdrawController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:WithdrawController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:WithdrawController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:WithdrawController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:WithdrawController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:WithdrawController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:WithdrawController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["qingguo-api/controllers:WithdrawController"] = append(beego.GlobalControllerRouter["qingguo-api/controllers:WithdrawController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
