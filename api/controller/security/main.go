package security

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"../../config"
	"../../model"
	"../../util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"encoding/base64"
	"os"
	"io"
)

func Register(c *gin.Context) {

	cloudantUrl := config.StrNoSQLDrive()
	//cloudant := util.CloudantDefault()
	//
	////ensure db exists
	////if the db exists the db will be returned anyway
	//dbName := config.StrNoSQLDBname()
	// cloudant.CreateDB(dbName)

	if cloudantUrl == "" {
		c.JSON(200, gin.H{})
		return
	}

	var (
		// result model.AlldocsResult
		// doc    model.Document
		// docs   []model.Document
		datas util.Response
		user  model.User
	)

	if c.BindJSON(&user) == nil {
		var arrKey = []string{"users", "_id", "_rev"}

		query := model.QuerySelectorAll{
			Selector: map[string]interface{}{
				"meta":        arrKey[0],
				"users.email": user.Email,
			},
			Fields: arrKey,
		}

		respText := util.FindDataAll(query)

		jsonToString := (respText)
		decode := []byte(jsonToString)
		var result model.UserDocumentsArray
		json.Unmarshal(decode, &result)

		var us []model.User
		for i := range result.Doc {

			fmt.Println("##i: ", result.Doc[i])
			us = append(us, model.User{
				IDs:      result.Doc[i].ID,
				Rev:      result.Doc[i].Rev,
				ID:       result.Doc[i].ID,
				Username: result.Doc[i].User.Username,
				Email:    result.Doc[i].User.Email,
				Password: result.Doc[i].User.Password,
			})

		}

		if len(us) != 0 {
			datas = util.Response{
				false,
				"error_exception",
				nil,
			}
			//c.JSON(200, datas)
		} else {
			if user.Avatar != "" {
				user.Avatar = util.B64ToImage(user.Avatar)
			}

			byteP := []byte(user.Password)
			p := util.HashAndSalt(byteP)
			tt := model.User{
				Avatar:   user.Avatar,
				Username: user.Username,
				Email:    user.Email,
				Password: p,
			}

			if user.Username == "" && user.Email == "" && user.Password == "" {
				datas = util.Response{
					false,
					"error_exception",
					nil,
				}
				//c.JSON(200, datas)
			} else {
				var arrKey = []string{"users"}
				//cloudant.DB(dbName).Post(map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "users": tt})
				util.PostCouchDB(map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "users": tt})
				//c.JSON(200, user)
				datas = util.Response{
					true,
					"ok",
					user,
				}
			}
		}
	} else {
		datas = util.Response{
			false,
			"error_exception",
			nil,
		}
		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func Login(c *gin.Context) {
	var (
		datas util.Response
		user  model.User
	)
	if c.BindJSON(&user) == nil {

		if user.Username == "" && user.Email == "" && user.Password == "" {
			datas = util.Response{
				false,
				"error_exception",
				nil,
			}
			//c.JSON(200, datas)
		} else {
			var arrKey = []string{"users", "_id", "_rev"}

			query := model.QuerySelectorAll{
				Selector: map[string]interface{}{
					"meta":        arrKey[0],
					"users.email": user.Email,
				},
				Fields: arrKey,
			}

			respText := util.FindDataAll(query)

			jsonToString := (respText)
			decode := []byte(jsonToString)
			var results model.UserDocumentsArray
			json.Unmarshal(decode, &results)

			var us []model.User
			for i := range results.Doc {
				fmt.Println("##i: ", results.Doc[i])
				us = append(us, model.User{
					IDs:      results.Doc[i].ID,
					Rev:      results.Doc[i].Rev,
					ID:       results.Doc[i].ID,
					Username: results.Doc[i].User.Username,
					Email:    results.Doc[i].User.Email,
					Password: results.Doc[i].User.Password,
				})
			}

			if len(results.Doc) == 0 {
				datas = util.Response{
					false,
					"error_exception",
					nil,
				}
				//c.JSON(200, datas)
			} else {

				find := us[0]
				byteP := []byte(user.Password)

				if util.ComparePasswords(find.Password, byteP) {
					id := find.ID
					claims := util.CustomClaims{
						ID:    id,
						Email: find.Email,
						StandardClaims: jwt.StandardClaims{
							ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
						},
					}
					tokenStr := util.SetToken(claims)

					strTo := util.Token{Token: tokenStr}

					datas = util.Response{
						true,
						"ok",
						strTo,
					}
				} else {
					datas = util.Response{
						false,
						"error_email_password_mismatch",
						nil,
					}
				}
			}
			//c.JSON(200, datas)
		}
	} else {
		datas = util.Response{
			false,
			"error_exception",
			nil,
		}
		//c.JSON(200, datas)
	}
	c.JSON(200, datas)

}

func EditUser(c *gin.Context) {

	var (
		// result model.AlldocsResult
		// doc    model.Document
		// docs   []model.Document
		datas util.Response
		user  model.User
	)

	if c.BindJSON(&user) == nil {

		var arrKey = []string{"users", "_id", "_rev"}

		query := model.QuerySelectorAll{
			Selector: map[string]interface{}{
				"meta":        arrKey[0],
				"users.email": user.Email,
			},
			Fields: arrKey,
		}

		respText := util.FindDataAll(query)

		jsonToString := (respText)
		decode := []byte(jsonToString)
		var results model.UserDocumentsArray
		json.Unmarshal(decode, &results)

		var us []model.User
		for i := range results.Doc {

			fmt.Println("##i: ", results.Doc[i])
			us = append(us, model.User{
				IDs:      results.Doc[i].ID,
				Rev:      results.Doc[i].Rev,
				ID:       results.Doc[i].ID,
				Username: results.Doc[i].User.Username,
				Email:    results.Doc[i].User.Email,
				Password: results.Doc[i].User.Password,
			})

		}

		if len(results.Doc) == 0 {
			datas = util.Response{
				false,
				"error_exception",
				nil,
			}
			//c.JSON(200, datas)
		} else if len(us) == 0 {
			datas = util.Response{
				false,
				"error_exception",
				nil,
			}
			//c.JSON(200, datas)
		} else {
			find := us[0]

			datas = util.Response{
				true,
				"ok",
				find,
			}

			//c.JSON(200, datas)
		}
	} else {
		datas = util.Response{
			false,
			"error_exception",
			nil,
		}
		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func PermissionFaker(c *gin.Context) {
	var datas util.Response
	cloudantUrl := config.StrNoSQLDrive()
	//cloudant := util.CloudantDefault()
	//
	////ensure db exists
	////if the db exists the db will be returned anyway
	//dbName := config.StrNoSQLDBname()
	//// cloudant.CreateDB(dbName)
	//
	//var result model.AlldocsResult
	if cloudantUrl == "" {
		c.JSON(200, gin.H{})
		return
	}

	data0, err := ioutil.ReadFile("FileSystem/permissionFaker.json")
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("#strData: ", string(data))

	data1, err1 := ioutil.ReadFile("FileSystem/permissionFakerOld.json")
	if err1 != nil {
		fmt.Println(err1)
	}

	var arrKey1 = []string{"permissions", "_id", "_rev"}

	query1 := model.QuerySelectorAll{
		Selector: map[string]interface{}{
			"meta": arrKey1[0],
		},
		Fields: arrKey1,
	}

	respText1 := util.FindDataAll(query1)

	jsonToString1 := (respText1)
	decode0x := []byte(jsonToString1)
	var results1 model.PermissionDocumentsArray
	json.Unmarshal(decode0x, &results1)

	//var arrKey = []string{"permissions"}
	//errs := cloudant.DB(dbName).AllDocs(&result, couchdb.Options{"include_docs": true, "startkey": arrKey})
	//if errs != nil {
	//	// c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch docs"})
	//	datas = util.Response{
	//		false,
	//		"error_exception",
	//		nil,
	//	}
	//	c.JSON(200, datas)
	//}
	//
	//jsonToString, _ := json.Marshal(result.Rows)
	//// fmt.Println("##jsonToString: ", string(jsonToString))
	//
	//decode := []byte(jsonToString)
	//var results []model.PermissionDocument
	//json.Unmarshal(decode, &results)

	var ts []model.Permission

	for i := range results1.Doc {
		// fmt.Println("##doc: ", results[i].Doc)
		// fmt.Println("##i: ", results[i].Doc.Test)

		ts = append(ts, model.Permission{
			IDs:                    results1.Doc[i].ID,
			Rev:                    results1.Doc[i].Rev,
			ID:                     results1.Doc[i].ID,
			ObjectModulePermission: results1.Doc[i].Permission.ObjectModulePermission,
			CreatedAt:              results1.Doc[i].Permission.CreatedAt,
			UpdatedAt:              results1.Doc[i].Permission.UpdatedAt,
		})

	}

	// var perm []model.Permission
	// count := "*"
	// db.Debug().Find(&perm).Count(&count)
	perm := ts

	decode0 := []byte(string(data0))
	var result0 []interface{}
	json.Unmarshal(decode0, &result0)

	decode1 := []byte(string(data1))
	var result1 []interface{}
	json.Unmarshal(decode1, &result1)

	if len(perm) == 0 {
		for i := range result0 {
			fmt.Println("##i: ", result0[i])
			js, _ := json.Marshal(result0[i])
			fmt.Println("##js: ", string(js))

			t := model.Permission{
				ObjectModulePermission: string(js),
			}

			// db.Debug().Create(&t)
			var arrKey = []string{"permissions"}
			//cloudant.DB(dbName).Post(map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "permissions": t})
			util.PostCouchDB(map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "permissions": t})
		}
	} else {
		for i := range result1 {
			js1, _ := json.Marshal(result1[i])
			// var perms model.Permission
			// db.Debug().Where(map[string]interface{}{"object_module_permission": string(js1)}).First(&perms)
			permissionx := result1
			for j := range permissionx {
				fmt.Println(j)
				fmt.Println("##perms. : ", string(js1))
				// fmt.Println("##perms.id: ", perms.ID)
				// perms.ObjectModulePermission = string(js1)
			}

			// db.Save(&perms)
			arrayData, _ := json.Marshal(permissionx)
			byteArrayData := []byte(arrayData)

			err := ioutil.WriteFile("FileSystem/permissionFakerOld.json", byteArrayData, 0777)
			if err != nil {
				fmt.Println("err file: ", err)
			}
		}

	}

	datas = util.Response{
		true,
		"ok",
		"",
	}

	c.JSON(http.StatusAccepted, datas)
}

func PermissionAll(c *gin.Context) {
	var datas util.Response

	var arrKey = []string{"permissions", "_id", "_rev"}

	query := model.QuerySelectorAll{
		Selector: map[string]interface{}{
			"meta": arrKey[0],
		},
		Fields: arrKey,
	}

	respText := util.FindDataAll(query)

	jsonToString := (respText)
	decode := []byte(jsonToString)
	var results model.PermissionDocumentsArray
	json.Unmarshal(decode, &results)

	var ts []model.Permission

	for i := range results.Doc {
		fmt.Println("##docs: ", results.Doc[i])
		// fmt.Println("##i: ", results[i].Doc.Permission)
		if results.Doc[i].Permission.ObjectModulePermission != "" {

			// fmt.Println("##ObjectModulePermission: ", results[i].Doc.Permission.ObjectModulePermission)

			decodeModules := []byte(results.Doc[i].Permission.ObjectModulePermission)
			// var modules map[string]interface{}
			var modules interface{}
			json.Unmarshal(decodeModules, &modules)

			ts = append(ts, model.Permission{
				IDs:                    results.Doc[i].ID,
				Rev:                    results.Doc[i].Rev,
				ID:                     results.Doc[i].ID,
				ObjectModulePermission: results.Doc[i].Permission.ObjectModulePermission,
				Modules:                modules,
				CreatedAt:              results.Doc[i].Permission.CreatedAt,
				UpdatedAt:              results.Doc[i].Permission.UpdatedAt,
			})
		}
	}

	if len(results.Doc) == 0 {
		datas = util.Response{
			false,
			"error_exception",
			nil,
		}
		//c.JSON(200, datas)
	} else {

		datas = util.Response{
			true,
			"ok",
			ts,
		}

		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func AssignRoles(c *gin.Context) {
	cloudantUrl := config.StrNoSQLDrive()
	//cloudant := util.CloudantDefault()

	//ensure db exists
	//if the db exists the db will be returned anyway
	//dbName := config.StrNoSQLDBname()
	// cloudant.CreateDB(dbName)

	if cloudantUrl == "" {
		c.JSON(200, gin.H{})
		return
	}

	var (
		datas util.Response
		t     model.Profile
	)

	if c.BindJSON(&t) == nil {
		// cloudant.DB(dbName).Put(id, map[string]interface{}{"tests": t}, rev

		if t.UserID == "" && t.PermissionID == "" {
			datas = util.Response{
				false,
				"error_exception",
				nil,
			}
			//c.JSON(200, datas)
		} else {
			var arrKey = []string{"profiles"}
			//cloudant.DB(dbName).Post(map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "profiles": t})
			util.PostCouchDB(map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "profiles": t})
			datas = util.Response{
				true,
				"ok",
				t,
			}
			//c.JSON(200, datas)
		}
	} else {
		datas = util.Response{
			false,
			"error_exception",
			nil,
		}
		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func EditRoles(c *gin.Context) {
	id := c.Param("id")

	cloudantUrl := config.StrNoSQLDrive()
	//cloudant := util.CloudantDefault()
	//
	////ensure db exists
	////if the db exists the db will be returned anyway
	//dbName := config.StrNoSQLDBname()
	// cloudant.CreateDB(dbName)
	var datas util.Response

	if cloudantUrl == "" {
		c.JSON(200, gin.H{})
		return
	}

	var arrKeyz = []string{"profiles", "_id", "_rev"}

	queryz := model.QuerySelectorAll{
		Selector: map[string]interface{}{
			"meta": arrKeyz[0],
			"_id":  id,
		},
		Fields: arrKeyz,
	}

	respTextz := util.FindDataAll(queryz)

	jsonToStringz := (respTextz)
	decodez := []byte(jsonToStringz)
	var resultz model.ProfileDocumentsArray
	json.Unmarshal(decodez, &resultz)

	var arrKeyz1 = []string{"users", "_id", "_rev"}

	queryz1 := model.QuerySelectorAll{
		Selector: map[string]interface{}{
			"meta": arrKeyz1[0],
			"_id":  resultz.Doc[0].Profile.UserID,
		},
		Fields: arrKeyz1,
	}

	respTextz1 := util.FindDataAll(queryz1)

	jsonToStringz1 := (respTextz1)
	decodez1 := []byte(jsonToStringz1)
	var resultz1 model.UserDocumentsArray
	json.Unmarshal(decodez1, &resultz1)

	var arrKeyz2 = []string{"permissions", "_id", "_rev"}

	queryz2 := model.QuerySelectorAll{
		Selector: map[string]interface{}{
			"meta": arrKeyz2[0],
			"_id":  resultz.Doc[0].Profile.PermissionID,
		},
		Fields: arrKeyz2,
	}

	respTextz2 := util.FindDataAll(queryz2)

	jsonToStringz2 := (respTextz2)
	decodez2 := []byte(jsonToStringz2)
	var resultz2 model.PermissionDocumentsArray
	json.Unmarshal(decodez2, &resultz2)

	// var result model.AlldocsResult

	decodeModules := []byte(resultz2.Doc[0].Permission.ObjectModulePermission)
	// var modules map[string]interface{}
	var modules interface{}
	json.Unmarshal(decodeModules, &modules)

	perms := model.Permission{
		IDs:                    resultz2.Doc[0].ID,
		Rev:                    resultz2.Doc[0].Rev,
		ID:                     resultz2.Doc[0].ID,
		ObjectModulePermission: resultz2.Doc[0].Permission.ObjectModulePermission,
		Modules:                modules,
	}

	//find := us
	find1 := perms
	//
	//fmt.Println("##find: ", (find))
	//fmt.Println("##find2: ", (find1))

	t := model.Profile{
		IDs:        resultz.Doc[0].ID,
		Rev:        resultz.Doc[0].Rev,
		ID:         resultz.Doc[0].ID,
		User:       resultz1.Doc[0].User,
		Permission: find1,
		CreatedAt:  resultz.Doc[0].Profile.CreatedAt,
		UpdatedAt:  resultz.Doc[0].Profile.UpdatedAt,
	}

	datas = util.Response{
		true,
		"ok",
		t,
	}

	c.JSON(200, datas)

	// c.JSON(200, t)

}

func GetRolByUser(c *gin.Context) {
	id := c.Param("user_id")

	var rol util.Rol
	rol.Acl = "test"
	var datas util.Response

	if util.IsEdit(c, rol).Success == true {

		var arrKey = []string{"profiles", "_id", "_rev"}

		query := model.QuerySelectorAll{
			Selector: map[string]interface{}{
				"meta":             arrKey[0],
				"profiles.user_id": id,
			},
			Fields: arrKey,
		}

		respText := util.FindDataAll(query)

		jsonToString := (respText)
		decode := []byte(jsonToString)
		var result model.ProfileDocumentsArray
		json.Unmarshal(decode, &result)

		results := result.Doc[0]

		t := model.Profile{
			IDs:          results.ID,
			Rev:          results.Rev,
			ID:           results.ID,
			UserID:       results.Profile.UserID,
			PermissionID: results.Profile.PermissionID,
		}

		datas = util.Response{
			true,
			"ok",
			t,
		}

		//c.JSON(200, datas)

	} else {
		datas = util.IsEdit(c, rol)

		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func UpdateRoles(c *gin.Context) {
	id := c.Param("id")
	rev := c.Param("rev")

	cloudantUrl := config.StrNoSQLDrive()
	//cloudant := util.CloudantDefault()
	//
	////ensure db exists
	////if the db exists the db will be returned anyway
	//dbName := config.StrNoSQLDBname()
	// cloudant.CreateDB(dbName)

	if cloudantUrl == "" {
		c.JSON(200, gin.H{})
		return
	}

	var (
		datas util.Response
		t     model.Profile
	)

	if c.BindJSON(&t) == nil {

		if t.UserID == "" && t.PermissionID == "" {
			datas = util.Response{
				false,
				"error_exception",
				nil,
			}
			//c.JSON(200, datas)
		} else {
			var arrKey = []string{"profiles"}
			//cloudant.DB(dbName).Put(id, map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "profiles": t}, rev)
			util.PutCouchDBByID(id, map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "profiles": t, "_id": id, "_rev": rev})
			datas = util.Response{
				true,
				"ok",
				t,
			}
			//c.JSON(200, datas)
		}
	} else {
		datas = util.Response{
			false,
			"error_exception",
			nil,
		}
		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func DeleteRoles(c *gin.Context) {
	id := c.Param("id")
	rev := c.Param("rev")

	cloudantUrl := config.StrNoSQLDrive()
	cloudant := util.CloudantDefault()

	//ensure db exists
	//if the db exists the db will be returned anyway
	dbName := config.StrNoSQLDBname()
	// cloudant.CreateDB(dbName)

	if cloudantUrl == "" {
		c.JSON(200, gin.H{})
		return
	}

	cloudant.DB(dbName).Delete(id, rev)
	var datas util.Response
	datas = util.Response{
		true,
		"ok",
		nil,
	}
	c.JSON(200, datas)
}

func Menu(c *gin.Context) {
	var datas util.Response

	menu := util.ArrayPermMenu(c)
	fmt.Println("#menu: ", menu)
	if len(menu) != 0 {
		datas = util.Response{
			true,
			"ok",
			menu,
		}
	} else {
		datas = util.Response{
			false,
			"menu empty",
			"",
		}
	}

	c.JSON(http.StatusAccepted, datas)
}

func SetupPermission(c *gin.Context) {
	var datas util.Response
	strResponse0 := util.CurlGet("/api/testdb/all")
	fmt.Println("##strResponse0: ", string(strResponse0))

	strResponse := util.CurlGet("/api/permission/faker")
	fmt.Println("##strResponse: ", string(strResponse))
	strTk, _ := json.Marshal(strResponse)
	decode := []byte(strTk)
	var results interface{}
	json.Unmarshal(decode, &results)

	datas = util.Response{
		true,
		"ok",
		results,
	}

	c.JSON(http.StatusAccepted, datas)
}

func SetupAdmin(c *gin.Context) {

	var (
		rol   util.Rol
		datas util.Response
		t     model.User
	)
	rol.Rol = "admin"

	if c.BindJSON(&t) == nil {

		strResponse0 := util.CurlPost(t, "/api/87780FA5DE684E87CB92B279F0BC07B14F572851E73B8943A097C1770A5F38E6")
		fmt.Println("##strResponse0: ", string(strResponse0))

		objUser := map[string]interface{}{
			"email": t.Email,
		}

		g := util.GetIDsByRol(c, rol)
		p := g[0]
		fmt.Println("##GetIDsByRol: ", (p.IDs))

		strResponse1 := util.CurlPost(objUser, "/api/edit-user")
		fmt.Println("##strResponse1: ", string(strResponse1))

		decode := []byte(strResponse1)
		var results util.ResponseU
		json.Unmarshal(decode, &results)

		// fmt.Println("##strResponse data1: ", (results))
		// fmt.Println("##strResponse   1: ", (results.Data.IDs))

		objUserRoles := map[string]interface{}{
			"user_id":       results.Data.IDs,
			"permission_id": p.IDs,
		}

		strResponse2 := util.CurlPost(objUserRoles, "/api/roles/assign")
		fmt.Println("##strResponse2: ", string(strResponse2))

		datas = util.Response{
			true,
			"ok",
			t,
		}
		//c.JSON(200, datas)
	} else {
		datas = util.Response{
			false,
			"error_exception",
			nil,
		}
		//c.JSON(200, datas)
	}
	c.JSON(200, datas)

}

func SetupUser(c *gin.Context) {
	var (
		rol   util.Rol
		datas util.Response
		t     model.User
	)
	rol.Rol = "users"

	if c.BindJSON(&t) == nil {

		strResponse0 := util.CurlPost(t, "/api/87780FA5DE684E87CB92B279F0BC07B14F572851E73B8943A097C1770A5F38E6")
		fmt.Println("##strResponse0: ", string(strResponse0))

		objUser := map[string]interface{}{
			"email": t.Email,
		}

		g := util.GetIDsByRol(c, rol)
		p := g[0]
		fmt.Println("##GetIDsByRol: ", (p.IDs))

		strResponse1 := util.CurlPost(objUser, "/api/edit-user")
		fmt.Println("##strResponse1: ", string(strResponse1))

		decode := []byte(strResponse1)
		var results util.ResponseU
		json.Unmarshal(decode, &results)

		// fmt.Println("##strResponse data1: ", (results))
		// fmt.Println("##strResponse   1: ", (results.Data.IDs))

		objUserRoles := map[string]interface{}{
			"user_id":       results.Data.IDs,
			"permission_id": p.IDs,
		}

		strResponse2 := util.CurlPost(objUserRoles, "/api/roles/assign")
		fmt.Println("##strResponse2: ", string(strResponse2))

		datas = util.Response{
			true,
			"ok",
			t,
		}
		//c.JSON(200, datas)
	} else {
		datas = util.Response{
			false,
			"error_exception",
			nil,
		}
		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func AddUsers(c *gin.Context) {
	var (
		rol   util.Rol
		role  util.Rol
		datas util.Response
		t     model.PayloadUserRoles
		u     model.User
	)
	// rol.Rol = "users"
	///////////////////////
	role.Rol = "admin"
	role.Acl = "register_users"
	if util.IsRead(c, role).Success == true {

		if c.BindJSON(&t) == nil {
			if t.Avatar!=""{
				t.Avatar=util.B64ToImage(t.Avatar)
			}
			u = model.User{
				Avatar:   t.Avatar,
				Username: t.Username,
				Email:    t.Email,
				Password: t.Password,
			}
			strResponse0 := util.CurlPost(u, "/api/87780FA5DE684E87CB92B279F0BC07B14F572851E73B8943A097C1770A5F38E6")
			fmt.Println("##strResponse0: ", string(strResponse0))

			objUser := map[string]interface{}{
				"email": t.Email,
			}

			rol.Rol = t.Rol
			g := util.GetIDsByRol(c, rol)
			p := g[0]
			fmt.Println("##GetIDsByRol: ", (p.IDs))

			strResponse1 := util.CurlPost(objUser, "/api/edit-user")
			fmt.Println("##strResponse1: ", string(strResponse1))

			decode := []byte(strResponse1)
			var results util.ResponseU
			json.Unmarshal(decode, &results)

			// fmt.Println("##strResponse data1: ", (results))
			// fmt.Println("##strResponse   1: ", (results.Data.IDs))

			objUserRoles := map[string]interface{}{
				"user_id":       results.Data.IDs,
				"permission_id": p.IDs,
			}

			strResponse2 := util.CurlPost(objUserRoles, "/api/roles/assign")
			fmt.Println("##strResponse2: ", string(strResponse2))

			datas = util.Response{
				true,
				"ok",
				t,
			}
			//c.JSON(200, datas)
		} else {
			datas = util.Response{
				false,
				"error_exception",
				nil,
			}
			//c.JSON(200, datas)
		}
	} else {
		datas = util.IsRead(c, role)

		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func SearchPaginateUsers(c *gin.Context) {
	skip := c.Param("skip")
	limit := c.Param("limit")

	skips, errk := strconv.ParseUint(skip, 10, 32)
	if errk != nil {
		fmt.Println(errk)
	}

	limits, errl := strconv.ParseUint(limit, 10, 32)
	if errl != nil {
		fmt.Println(errl)
	}

	sk := uint(skips)
	lm := uint(limits)

	fmt.Println("##sk: ", sk)
	fmt.Println("##lm: ", lm)

	var rol util.Rol
	rol.Rol = "admin"
	rol.Acl = "register_users"
	var datas util.Response
	var t model.User

	if util.IsRead(c, rol).Success == true {
		if c.BindJSON(&t) == nil {
			var arrKey = []string{"users", "_id", "_rev", "meta", "tag"}
			var arrTag = []string{"users"}
			input := []byte(`[{"_id": "desc"}]`)
			var sort []interface{}
			errj := json.Unmarshal(input, &sort)
			if errj != nil {
				log.Fatal(errj)
			}

			query := map[string]interface{}{
				"selector": map[string]interface{}{
					"meta": arrKey[0],
					"users.username": map[string]interface{}{
						"$gt": t.Username,
					},
					"tag": arrTag,
				},
				"limit":  lm,
				"skip":   sk,
				"fields": arrKey,
				"sort":   sort,
			}

			fmt.Println(query)

			str, _ := json.Marshal(query)
			decodex := []byte(str)

			var resultsx interface{}
			json.Unmarshal(decodex, &resultsx)

			fmt.Println("##str: ", resultsx)

			respText := util.FindDataInterface(resultsx)

			jsonToString := (respText)
			decode := []byte(jsonToString)
			var results model.UserDocumentsArray
			json.Unmarshal(decode, &results)

			var ts []model.User

			for i := range results.Doc {
				ts = append(ts, model.User{
					IDs:       results.Doc[i].ID,
					Rev:       results.Doc[i].Rev,
					ID:        results.Doc[i].ID,
					Username:  results.Doc[i].User.Username,
					Email:     results.Doc[i].User.Email,
					CreatedAt: results.Doc[i].User.CreatedAt,
					UpdatedAt: results.Doc[i].User.UpdatedAt,
				})
			}

			if len(results.Doc) == 0 {
				datas = util.Response{
					true,
					"empty_data",
					results.Doc,
				}
				//c.JSON(200, datas)
				//return
			} else {

				datas = util.Response{
					true,
					"ok",
					ts,
				}
				//c.JSON(200, datas)
				//return
			}
		} else {
			datas = util.Response{
				false,
				"empty_data",
				nil,
			}
			//c.JSON(200, datas)
			//return
		}

	} else {
		datas = util.IsRead(c, rol)

		//c.JSON(200, datas)
		// return
	}
	c.JSON(200, datas)
}

func PaginateUsers(c *gin.Context) {
	skip := c.Param("skip")
	limit := c.Param("limit")

	skips, errk := strconv.ParseUint(skip, 10, 32)
	if errk != nil {
		fmt.Println(errk)
	}

	limits, errl := strconv.ParseUint(limit, 10, 32)
	if errl != nil {
		fmt.Println(errl)
	}

	sk := uint(skips)
	lm := uint(limits)

	fmt.Println("##sk: ", sk)
	fmt.Println("##lm: ", lm)

	var rol util.Rol
	rol.Rol = "admin"
	rol.Acl = "register_users"
	var datas util.Response

	if util.IsRead(c, rol).Success == true {

		var arrKey = []string{"users", "_id", "_rev", "meta", "tag"}
		var arrTag = []string{"users"}
		input := []byte(`[{"_id": "desc"}]`)
		var sort []interface{}
		errj := json.Unmarshal(input, &sort)
		if errj != nil {
			log.Fatal(errj)
		}

		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"meta": arrKey[0],
				"tag":  arrTag,
			},
			"limit":  lm,
			"skip":   sk,
			"fields": arrKey,
			"sort":   sort,
		}

		fmt.Println(query)

		str, _ := json.Marshal(query)
		decodex := []byte(str)

		var resultsx interface{}
		json.Unmarshal(decodex, &resultsx)

		fmt.Println("##str: ", resultsx)

		respText := util.FindDataInterface(resultsx)

		jsonToString := (respText)
		decode := []byte(jsonToString)
		var results model.UserDocumentsArray
		json.Unmarshal(decode, &results)

		var ts []model.User

		for i := range results.Doc {
			ts = append(ts, model.User{
				IDs:       results.Doc[i].ID,
				Rev:       results.Doc[i].Rev,
				ID:        results.Doc[i].ID,
				Username:  results.Doc[i].User.Username,
				Email:     results.Doc[i].User.Email,
				CreatedAt: results.Doc[i].User.CreatedAt,
				UpdatedAt: results.Doc[i].User.UpdatedAt,
			})
		}

		if len(results.Doc) == 0 {
			datas = util.Response{
				true,
				"empty_data",
				results.Doc,
			}
			//c.JSON(200, datas)
			// return
		} else {

			datas = util.Response{
				true,
				"ok",
				ts,
			}
			//c.JSON(200, datas)
		}

	} else {
		datas = util.IsRead(c, rol)

		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func EditUsers(c *gin.Context) {
	id := c.Param("id")

	var rol util.Rol
	rol.Rol = "admin"
	rol.Acl = "register_users"
	var datas util.Response

	if util.IsEdit(c, rol).Success == true {

		var arrKey = []string{"users", "_id", "_rev"}

		query := model.QuerySelectorAll{
			Selector: map[string]interface{}{
				"meta": arrKey[0],
				"_id":  id,
			},
			Fields: arrKey,
		}

		respText := util.FindDataAll(query)

		jsonToString := (respText)
		decode := []byte(jsonToString)
		var result model.UserDocumentsArray
		json.Unmarshal(decode, &result)

		results := result.Doc[0]

		t := model.User{
			IDs:       results.ID,
			Rev:       results.Rev,
			ID:        results.ID,
			Avatar:    results.User.Avatar,
			Username:  results.User.Username,
			Email:     results.User.Email,
			CreatedAt: results.User.CreatedAt,
			UpdatedAt: results.User.UpdatedAt,
		}

		datas = util.Response{
			true,
			"ok",
			t,
		}

		//c.JSON(200, datas)

	} else {
		datas = util.IsEdit(c, rol)

		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func PutUsers(c *gin.Context) {
	id := c.Param("id")
	rev := c.Param("rev")
	var rol util.Rol
	var roles util.Rol
	rol.Rol = "admin"
	rol.Acl = "register_users"
	var datas util.Response

	if util.IsUpdate(c, rol).Success == true {

		cloudantUrl := config.StrNoSQLDrive()
		//cloudant := util.CloudantDefault()
		//
		////ensure db exists
		////if the db exists the db will be returned anyway
		//dbName := config.StrNoSQLDBname()
		// cloudant.CreateDB(dbName)

		if cloudantUrl == "" {
			c.JSON(200, gin.H{})
			return
		}

		var (
			// datas util.Response
			// t1 model.User
			t model.PayloadUserRoles
			// u  model.User
		)

		if c.BindJSON(&t) == nil {

			if t.Username == "" && t.Email == "" {
				datas = util.Response{
					false,
					"error_exception",
					nil,
				}
				//c.JSON(200, datas)
			} else {

				roles.Rol = t.Rol
				objUser := map[string]interface{}{
					"email": t.Email,
				}

				if t.Avatar != "" {
					t.Avatar = util.B64ToImage(t.Avatar)
				}

				byteP := []byte(t.Password)
				pass := util.HashAndSalt(byteP)
				payloadUser := map[string]interface{}{
					"avatar":   t.Avatar,
					"username": t.Username,
					"email":    t.Email,
					"password": pass,
				}

				g := util.GetIDsByRol(c, roles)
				p := g[0]
				fmt.Println("##GetIDsByRol: ", (p.IDs))

				strResponse1 := util.CurlPost(objUser, "/api/edit-user")
				fmt.Println("##strResponse1: ", string(strResponse1))

				decode := []byte(strResponse1)
				var results util.ResponseU
				json.Unmarshal(decode, &results)

				objUserRoles := map[string]interface{}{
					"user_id":       results.Data.IDs,
					"permission_id": p.IDs,
				}

				strResponse2 := util.CurlGet("/api/roles/by-user/" + results.Data.IDs)
				fmt.Println("##strResponse2: ", string(strResponse2))

				decode1 := []byte(strResponse2)
				var results1 util.ResponseProfile
				json.Unmarshal(decode1, &results1)
				strResponse3 := util.CurlPost(objUserRoles, "/api/roles/update/"+results1.Data.IDs+"/"+results1.Data.Rev)
				// roles/update/:id/:rev
				fmt.Println("##strResponse3: ", string(strResponse3))

				var arrKey = []string{"users"}

				//cloudant.DB(dbName).Put(id, map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "users": payloadUser}, rev)
				util.PutCouchDBByID(id, map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "users": payloadUser, "_id": id, "_rev": rev})
				datas = util.Response{
					true,
					"ok",
					t,
				}
				//c.JSON(200, datas)
			}
		} else {
			datas = util.Response{
				false,
				"error_exception",
				nil,
			}
			//c.JSON(200, datas)
		}
	} else {
		datas = util.IsUpdate(c, rol)

		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	rev := c.Param("rev")

	var rol util.Rol
	rol.Rol = "admin"
	rol.Acl = "register_users"
	var datas util.Response

	if util.IsDelete(c, rol).Success == true {

		cloudantUrl := config.StrNoSQLDrive()
		//cloudant := util.CloudantDefault()
		//
		////ensure db exists
		////if the db exists the db will be returned anyway
		//dbName := config.StrNoSQLDBname()
		// cloudant.CreateDB(dbName)

		if cloudantUrl == "" {
			c.JSON(200, gin.H{})
			return
		}

		//cloudant.DB(dbName).Delete(id, rev)
		util.DeleteCouchDBByID(id, rev)
		//var datas util.Response
		datas = util.Response{
			true,
			"ok",
			nil,
		}
		//c.JSON(200, datas)
	} else {
		datas = util.IsDelete(c, rol)

		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func ListPermissions(c *gin.Context) {
	var datas util.Response
	var arrKey = []string{"permissions", "_id", "_rev", "meta", "tag"}
	var arrTag = []string{"permissions"}
	input := []byte(`[{"_id": "desc"}]`)
	var sort []interface{}
	errj := json.Unmarshal(input, &sort)
	if errj != nil {
		log.Fatal(errj)
	}

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"meta": arrKey[0],
			"tag":  arrTag,
		},
		"fields": arrKey,
		"sort":   sort,
	}

	fmt.Println(query)

	str, _ := json.Marshal(query)
	decodex := []byte(str)

	var resultsx interface{}
	json.Unmarshal(decodex, &resultsx)

	fmt.Println("##str: ", resultsx)

	respText := util.FindDataInterface(resultsx)

	jsonToString := (respText)
	decode := []byte(jsonToString)
	var results model.PermissionDocumentsArray
	json.Unmarshal(decode, &results)

	var ts []model.ObjectModulePermission
	for i := range results.Doc {
		// perm:=results.Doc[i].Permission.ObjectModulePermission
		decodeModules := []byte(results.Doc[i].Permission.ObjectModulePermission)
		var modules model.ObjectModulePermission
		json.Unmarshal(decodeModules, &modules)
		ts = append(ts, model.ObjectModulePermission{
			IDs:    results.Doc[i].ID,
			Rev:    results.Doc[i].Rev,
			Rol:    modules.Rol,
			Module: modules.Module,
		})
	}

	if len(results.Doc) == 0 {
		datas = util.Response{
			false,
			"error_exception",
			nil,
		}
		//c.JSON(200, datas)
	} else {

		datas = util.Response{
			true,
			"ok",
			ts,
		}

		//c.JSON(200, datas)
	}
	c.JSON(200, datas)
}

func InfoUser(c *gin.Context) {
	var datas util.Response
	raw := util.Auth(c)
	if raw.Email == "" {
		datas = util.Response{
			false,
			"error",
			nil,
		}
	} else {

		datas = util.Response{
			true,
			"ok",
			raw,
		}
	}
	c.JSON(200, datas)
}

func B64ToImage(c *gin.Context) {

	/* let splitFile = img[i].split(',');
        let pos0, pos1;
        pos0 = splitFile[0];
        pos0 = pos0.split('/')
        let ext = pos0[1];
        pos1 = splitFile[1];
        let dataimg = await this.base64Decode(pos1);
        let dirImg = dir + '/img/';
        let imgRes = dirImg + nameFile + '.' + ext;*/

	img := "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAYABgAAD/4TSYRXhpZgAATU0AKgAAAAgABgALAAIAAAAmAAAIYgESAAMAAAABAAEAAAExAAIAAAAmAAAIiAEyAAIAAAAUAAAIrodpAAQAAAABAAAIwuocAAcAAAgMAAAAVgAAEUYc6gAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFdpbmRvd3MgUGhvdG8gRWRpdG9yIDEwLjAuMTAwMTEuMTYzODQAV2luZG93cyBQaG90byBFZGl0b3IgMTAuMC4xMDAxMS4xNjM4NAAyMDE5OjA2OjIwIDIzOjU2OjQzAAAGkAMAAgAAABQAABEckAQAAgAAABQAABEwkpEAAgAAAAMwMAAAkpIAAgAAAAMwMAAAoAEAAwAAAAEAAQAA6hwABwAACAwAAAkQAAAAABzqAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMjAxOTowNjoyMCAyMzo1NToxNgAyMDE5OjA2OjIwIDIzOjU1OjE2AAAAAAYBAwADAAAAAQAGAAABGgAFAAAAAQAAEZQBGwAFAAAAAQAAEZwBKAADAAAAAQAC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YWRvYmUuY29tL3hhcC8xLjAvADw/eHBhY2tldCBiZWdpbj0n77u/JyBpZD0nVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkJz8+DQo8eDp4bXBtZXRhIHhtbG5zOng9ImFkb2JlOm5zOm1ldGEvIj48cmRmOlJERiB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiPjxyZGY6RGVzY3JpcHRpb24gcmRmOmFib3V0PSJ1dWlkOmZhZjViZGQ1LWJhM2QtMTFkYS1hZDMxLWQzM2Q3NTE4MmYxYiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIj48eG1wOkNyZWF0b3JUb29sPldpbmRvd3MgUGhvdG8gRWRpdG9yIDEwLjAuMTAwMTEuMTYzODQ8L3htcDpDcmVhdG9yVG9vbD48eG1wOkNyZWF0ZURhdGU+MjAxOS0wNi0yMFQyMzo1NToxNjwveG1wOkNyZWF0ZURhdGU+PC9yZGY6RGVzY3JpcHRpb24+PC9yZGY6UkRGPjwveDp4bXBtZXRhPg0KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgICAgICAg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"
	fmt.Println(img)
	splitFile := strings.Split(img, ",")
	pos0 := splitFile[0]
	pos00 := strings.Split(pos0, "/")
	ext0 := pos00[1]
	ext1 := strings.Split(ext0, ";")
	ext := ext1[0]
	fmt.Println(ext)
	pos1 := splitFile[1]

	dec, err := base64.StdEncoding.DecodeString(pos1)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("Images/gallery." + ext)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	// go to begginng of file
	f.Seek(0, 0)

	// output file contents
	io.Copy(os.Stdout, f)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Jokes handler not implemented yet",
	})

}
