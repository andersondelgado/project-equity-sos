package userKyc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	//"strconv"

	"../../config"
	"../../model"
	"../../util"
	"github.com/gin-gonic/gin"
	//"github.com/timjacobi/go-couchdb"
)

// Documents KYC //
func SelectKyc(c *gin.Context) {
	var rol util.Rol
	rol.Acl = "kyc"
	var datas util.Response
	//datas = util.Response{
	//	false,
	//	"error_exception",
	//	nil,
	//}

	if util.IsRead(c, rol).Success == true {
		var arrKey = []string{"kyc", "_id", "_rev"}

		query := model.QuerySelectorAll{
			Selector: map[string]interface{}{
				"meta": arrKey[0],
			},
			Fields: arrKey,
		}

		//fmt.Println(query)
		respText := util.FindDataAll(query)

		/*fmt.Println(respText)
		fmt.Println("*************************\n")*/

		jsonToString := (respText)
		decode := []byte(jsonToString)
		var results model.KYCDocumentsArray
		json.Unmarshal(decode, &results)
		var ts []model.KYC

		for i := range results.Doc {
			if results.Doc[i].KYC.Name != "" {
				ts = append(ts, model.KYC{
					IDs:         results.Doc[i].ID,
					Rev:         results.Doc[i].Rev,
					ID:          results.Doc[i].ID,
					Name:        results.Doc[i].KYC.Name,
					Description: results.Doc[i].KYC.Description,
					CreatedAt:   results.Doc[i].KYC.CreatedAt,
					UpdatedAt:   results.Doc[i].KYC.UpdatedAt,
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
		datas = util.IsRead(c, rol)
		//c.JSON(200, datas)
		//return
	}
	c.JSON(200, datas)
}

// Documents KYC //
func SelectKycUser(c *gin.Context) {

	var rol util.Rol
	rol.Acl = "test"
	var datas util.Response

	if util.IsRead(c, rol).Success == true {

		var arrKey = []string{"kyc_user", "_id", "_rev"}
		UserID := c.Param("id")
		if (UserID == "") {
			user := util.Auth(c)
			UserID = user.ID
		}
		fmt.Println(UserID)

		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"meta":             arrKey[0],
				"kyc_user.user_id": UserID,
			},
		}

		str, _ := json.Marshal(query)
		decodex := []byte(str)

		var resultsx interface{}
		json.Unmarshal(decodex, &resultsx)

		respText := util.FindDataInterface(resultsx)
		jsonToString := (respText)
		decode := []byte(jsonToString)
		var results model.KYCUserDocumentsArray
		json.Unmarshal(decode, &results)

		var ts model.KycUser
		for i := range results.Doc {
			ts = model.KycUser{
				IDs:              results.Doc[i].ID,
				Rev:              results.Doc[i].Rev,
				ID:               results.Doc[i].ID,
				UserID:           results.Doc[i].KycUser.UserID,
				CompanyName:      results.Doc[i].KycUser.CompanyName,
				CompanyCountryID: results.Doc[i].KycUser.CompanyCountryID,
				CountryID:        results.Doc[i].KycUser.CountryID,
				Name:             results.Doc[i].KycUser.Name,
				LastName:         results.Doc[i].KycUser.LastName,
				City:             results.Doc[i].KycUser.City,
				Address:          results.Doc[i].KycUser.Address,
				AboutPerson:      results.Doc[i].KycUser.AboutPerson,
				PostalCode:       results.Doc[i].KycUser.PostalCode,
				Dob:              results.Doc[i].KycUser.Dob,
				Attachment:       results.Doc[i].KycUser.Attachment,
				CreatedAt:        results.Doc[i].KycUser.CreatedAt,
				UpdatedAt:        results.Doc[i].KycUser.UpdatedAt,
			}
		}

		if len(results.Doc) == 0 {
			datas = util.Response{
				false,
				"error_exception",
				nil,
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
		datas = util.IsRead(c, rol)
		//c.JSON(200, datas)
		//return
	}
	c.JSON(200, datas)

}

func SelectKycUserByID(c *gin.Context) {

	var rol util.Rol
	rol.Acl = "test"
	var datas util.Response

	if util.IsRead(c, rol).Success == true {

		var arrKey = []string{"kyc_user", "_id", "_rev"}
		UserID := c.Param("id")
		//if (UserID == "") {
		//	user := util.Auth(c)
		//	UserID = user.ID
		//}
		fmt.Println(UserID)

		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"meta":             arrKey[0],
				"kyc_user.user_id": UserID,
			},
		}

		str, _ := json.Marshal(query)
		decodex := []byte(str)

		var resultsx interface{}
		json.Unmarshal(decodex, &resultsx)

		respText := util.FindDataInterface(resultsx)
		jsonToString := (respText)
		decode := []byte(jsonToString)
		var results model.KYCUserDocumentsArray
		json.Unmarshal(decode, &results)

		var ts model.KycUser
		for i := range results.Doc {
			ts = model.KycUser{
				IDs:              results.Doc[i].ID,
				Rev:              results.Doc[i].Rev,
				ID:               results.Doc[i].ID,
				UserID:           results.Doc[i].KycUser.UserID,
				CompanyName:      results.Doc[i].KycUser.CompanyName,
				CompanyCountryID: results.Doc[i].KycUser.CompanyCountryID,
				CountryID:        results.Doc[i].KycUser.CountryID,
				Name:             results.Doc[i].KycUser.Name,
				LastName:         results.Doc[i].KycUser.LastName,
				City:             results.Doc[i].KycUser.City,
				Address:          results.Doc[i].KycUser.Address,
				AboutPerson:      results.Doc[i].KycUser.AboutPerson,
				PostalCode:       results.Doc[i].KycUser.PostalCode,
				Dob:              results.Doc[i].KycUser.Dob,
				Attachment:       results.Doc[i].KycUser.Attachment,
				CreatedAt:        results.Doc[i].KycUser.CreatedAt,
				UpdatedAt:        results.Doc[i].KycUser.UpdatedAt,
			}
		}

		if len(results.Doc) == 0 {
			datas = util.Response{
				false,
				"error_exception",
				nil,
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
		datas = util.IsRead(c, rol)
		//c.JSON(200, datas)
		//return
	}
	c.JSON(200, datas)

}

func AddKycUser(c *gin.Context) {
	var rol util.Rol
	rol.Acl = "test"
	var datas util.Response

	if util.IsCreate(c, rol).Success == true {

		cloudantUrl := config.StrNoSQLDrive()
		//cloudant := util.CloudantDefault()
		//dbName := config.StrNoSQLDBname()

		if cloudantUrl == "" {
			c.JSON(200, gin.H{})
			return
		}

		var (
			//datas util.Response
			t model.KycUser
		)
		if c.BindJSON(&t) == nil {
			if t.Name == "" {
				datas = util.Response{
					false,
					"error_exception",
					nil,
				}
			} else {
				user := util.Auth(c)
				date := time.Now()

				for i := range t.Attachment {
					if t.Attachment[i].AttachmentName != "" {
						t.Attachment[i].AttachmentName = config.EnviromentsRaw().RemoteHost[0].Name + "/" + util.B64ToImage(t.Attachment[i].AttachmentName)
					}
				}

				t = model.KycUser{
					UserID:           user.ID,
					CompanyName:      t.CompanyName,
					CompanyCountryID: t.CompanyCountryID,
					CountryID:        t.CountryID,
					Name:             t.Name,
					LastName:         t.LastName,
					City:             t.City,
					Address:          t.Address,
					AboutPerson:      t.AboutPerson,
					PostalCode:       t.PostalCode,
					Dob:              t.Dob,
					Attachment:       t.Attachment,
					CreatedAt:        date,
					UpdatedAt:        date,
				}

				var arrKey = []string{"kyc_user"}

				//cloudant.DB(dbName).Post(map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "kyc_user": t})
				util.PostCouchDB(map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "kyc_user": t})
				datas = util.Response{
					true,
					"ok",
					t,
				}
			}
		} else {
			datas = util.Response{
				false,
				"error_exception",
				nil,
			}
		}
	} else {
		datas = util.IsCreate(c, rol)
	}
	c.JSON(200, datas)
}

func EditKycUser(c *gin.Context) {
	var rol util.Rol
	rol.Acl = "test"
	var datas util.Response

	if util.IsEdit(c, rol).Success == true {

		var (
			t model.KycUser
		)
		if c.BindJSON(&t) == nil {
			if t.Name == "" {
				datas = util.Response{
					false,
					"error_exception",
					nil,
				}
			} else {

				//cloudant := util.CloudantDefault()
				//dbName := config.StrNoSQLDBname()
				var attachmentDB []model.KycAttachment
				var arrKey = []string{"kyc_user", "_id", "_rev"}

				query := model.QuerySelectorAll{
					Selector: map[string]interface{}{
						"meta": arrKey[0],
						"_id":  t.ID,
					},
					Fields: arrKey,
				}

				respText := util.FindDataAll(query)
				jsonToString := (respText)
				decode := []byte(jsonToString)
				var result model.KYCUserDocumentsArray
				json.Unmarshal(decode, &result)

				if len(result.Doc) > 0 {
					attachmentDoc := result.Doc[0].KycUser.Attachment
					//Put Data
					if len(t.Attachment) > 0 {
						for j := range t.Attachment {
							//Query
							for i := range attachmentDoc {
								validateExist := false
								if t.Attachment[j].KycID != attachmentDoc[i].KycID {
									for z := range attachmentDB {
										if attachmentDB[z].KycID == attachmentDoc[i].KycID {
											validateExist = true
											break
										}
									}

									if validateExist {
										break
									} else {
										attachmentDB = append(attachmentDB, attachmentDoc[i])
									}
								}
							}
						}

						fmt.Println(len(attachmentDB))
						for i := range attachmentDB {
							t.Attachment = append(t.Attachment, attachmentDB[i])
						}
					} else {
						t.Attachment = attachmentDoc
					}
				}

				for i := range t.Attachment {
					if t.Attachment[i].AttachmentName != "" {
						t.Attachment[i].AttachmentName = config.EnviromentsRaw().RemoteHost[0].Name + "/" + util.B64ToImage(t.Attachment[i].AttachmentName)
					}
				}

				//CloudantDB PUT
				rev := t.Rev
				id := t.ID
				var arrKeyPost = []string{"kyc_user"}
				//cloudant.DB(dbName).Put(id, map[string]interface{}{"meta": arrKeyPost[0], "tag": arrKeyPost, "kyc_user": t}, rev)
				util.PutCouchDBByID(id, map[string]interface{}{"meta": arrKeyPost[0], "tag": arrKeyPost, "kyc_user": t, "_id": id, "_rev": rev})
				datas = util.Response{
					true,
					"ok",
					t,
				}
			}
		}
	} else {
		datas = util.IsEdit(c, rol)
	}
	c.JSON(200, datas)
}

func KycFaker(c *gin.Context) {
	data0, err := ioutil.ReadFile("FileSystem/kyc.json")
	if err != nil {
		fmt.Println(err)
	}
	decode0 := []byte(string(data0))

	var result0 []model.KYC

	json.Unmarshal(decode0, &result0)

	for i := range result0 {
		t := model.KYC{
			Name:        result0[i].Name,
			Description: result0[i].Description,
			CreatedAt:   result0[i].CreatedAt,
			UpdatedAt:   result0[i].UpdatedAt,
		}
		strResponse0 := util.CurlPost(t, "/api/kyc/bulk")
		fmt.Println("##strResponse0: ", string(strResponse0))
	}

	datas := util.Response{
		true,
		"ok",
		result0,
	}
	c.JSON(200, datas)
}

func BulkKyc(c *gin.Context) {
	cloudantUrl := config.StrNoSQLDrive()
	//cloudant := util.CloudantDefault()
	//
	//dbName := config.StrNoSQLDBname()

	if cloudantUrl == "" {
		c.JSON(200, gin.H{})
		return
	}

	var (
		datas util.Response
		t     model.KYC
	)

	if c.BindJSON(&t) == nil {
		// cloudant.DB(dbName).Post(t)
		if t.Name == "" && t.Description == "" {
			datas = util.Response{
				false,
				"error_exception",
				nil,
			}
			//c.JSON(200, datas)
		} else {
			var arrKey = []string{"kyc"}
			//cloudant.DB(dbName).Post(map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "kyc": t})
			util.PostCouchDB(map[string]interface{}{"meta": arrKey[0], "tag": arrKey, "kyc": t})
			datas = util.Response{
				true,
				"ok",
				t,
			}
			//c.JSON(200, datas)
		}
		// c.JSON(200, t)
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

// Documents KYC //
