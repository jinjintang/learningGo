package controllers

import (
	"encoding/json"
	"errors"
	"qingguo-api/models"

	"strings"

	"github.com/astaxie/beego"
)

// CourseCategoryController operations for CourseCategory
type CourseCategoryController struct {
	beego.Controller
}

// URLMapping ...
func (c *CourseCategoryController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create CourseCategory
// @Param	body		body 	models.CourseCategory	true		"body for CourseCategory content"
// @Success 201 {int} models.CourseCategory
// @Failure 403 body is empty
// @router / [post]
func (c *CourseCategoryController) Post() {
	var v models.CourseCategory
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddCourseCategory(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get CourseCategory by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.CourseCategory
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CourseCategoryController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	
	v, err := models.GetCourseCategoryById(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get CourseCategory
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.CourseCategory
// @Failure 403
// @router / [get]
func (c *CourseCategoryController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllCourseCategory(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the CourseCategory
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.CourseCategory	true		"body for CourseCategory content"
// @Success 200 {object} models.CourseCategory
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CourseCategoryController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	
	v := models.CourseCategory{Id: idStr}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateCourseCategoryById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the CourseCategory
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CourseCategoryController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	
	if err := models.DeleteCourseCategory(idStr); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
