package main

import (
	"fmt"
	"lancer-gen/constant"
	"os"
	"strings"
	"time"
)

func main() {
	// ./bin/exec 1 2 3 4

	args := os.Args
	if len(args) < constant.ArgsLength {
		fmt.Println("./exe Mode CSName FunName")
		return
	}

	mode := args[1]

	if mode != constant.CS && mode != constant.RR && mode != constant.MIGRATE && mode != constant.CURD {
		fmt.Println("params err")
		return
	}

	//mode := args[1]
	CsName := args[2]
	FunName := args[3]

	switch mode {
	case constant.CS:
		GenerateControllerService(CsName)
	case constant.RR:
		GenerateReqRes(CsName, FunName)
	case constant.MIGRATE:
		GenerateMigrateFile(CsName)
	case constant.CURD:
		GenerateCurd(CsName)
	default:
		fmt.Println("params error")
	}

}

//gen controller file and service file

func GenerateControllerService(name string) {

	//titleName := strings.Title(strings.Replace(name, "_", "", -1))

	var titleName string
	split := strings.Split(name, "_")
	for _, s := range split {
		titleName = titleName + strings.Title(s)
	}

	// gen service file
	serviceFileName := fmt.Sprintf("./service/%v_service.go", name)
	serviceFileContentTemplate := `package service

type %sService struct {
}

func New%sService() *%sService {
	return &%sService{}
}

`

	// replace
	serviceFileContent := fmt.Sprintf(serviceFileContentTemplate, titleName, titleName, titleName, titleName)

	// create file
	if _, err := os.Stat(serviceFileName); os.IsNotExist(err) {
		//file no exists
		err := os.WriteFile(serviceFileName, []byte(serviceFileContent), 0644)
		if err != nil {
			fmt.Println("Error creating service file:", err)
			return
		}
		fmt.Println("File", serviceFileName, "created successfully!")
	} else {
		fmt.Println("File", serviceFileName, "already exists!")
	}

	// gen controller file
	controllerFileName := fmt.Sprintf("./controller/%v_controller.go", name)
	controllerFileContentTemplate := `package controller

type %sController struct {
}

func New%sController() *%sController {
	return &%sController{}
}

`

	// replace
	controllerFileContent := fmt.Sprintf(controllerFileContentTemplate, titleName, titleName, titleName, titleName)

	// create file
	if _, err := os.Stat(controllerFileName); os.IsNotExist(err) {
		err := os.WriteFile(controllerFileName, []byte(controllerFileContent), 0644)
		if err != nil {
			fmt.Println("Error creating controller file:", err)
			return
		}
		fmt.Println("File", controllerFileName, "created and successfully!")
	} else {
		fmt.Println("File", controllerFileName, "already exists!")
	}

}

//gen request file and response file

func GenerateReqRes(name string, funName string) {
	//judge request exist
	requestDir := "./request"
	if _, err := os.Stat(requestDir); os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("%v no exists!", requestDir))
		return
	}
	responseDir := "./response"
	if _, err := os.Stat(responseDir); os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("%v no exists!", responseDir))
		return
	}

	var titleName string
	split := strings.Split(name+"_"+funName, "_")
	for _, s := range split {
		titleName = titleName + strings.Title(s)
	}

	//judge request name dir exist
	funcReqDir := requestDir + "/" + name
	if _, err := os.Stat(funcReqDir); os.IsNotExist(err) {
		// no exists, create at
		err = os.Mkdir(funcReqDir, 0777)
		if err != nil {
			fmt.Println(fmt.Sprintf("creat dir err:%v", err))
			return
		}
	}

	//judge response name dir exist
	funcResDir := responseDir + "/" + name
	if _, err := os.Stat(funcResDir); os.IsNotExist(err) {
		// no exists, create at
		err = os.Mkdir(funcResDir, 0777)
		if err != nil {
			fmt.Println(fmt.Sprintf("creat dir err:%v", err))
			return
		}
	}

	//gen request file
	reqFileName := fmt.Sprintf("%v/%v/%v.go", requestDir, name, name+"_"+funName)
	reqFileContentTemplate := `package request

type %sRequest struct {
}
`

	// replace
	reqFileContent := fmt.Sprintf(reqFileContentTemplate, titleName)

	// create file
	if _, err := os.Stat(reqFileName); os.IsNotExist(err) {
		//file no exists
		err := os.WriteFile(reqFileName, []byte(reqFileContent), 0644)
		if err != nil {
			fmt.Println("Error creating request file:", err)
			return
		}
		fmt.Println("File", reqFileName, "created and successfully!")
	} else {
		fmt.Println("File", reqFileName, "already exists!")
	}

	//gen response file
	resFileName := fmt.Sprintf("%v/%v/%v.go", responseDir, name, name+"_"+funName)
	resFileContentTemplate := `package response

type %sResponse struct {
}
`

	// replace
	resFileContent := fmt.Sprintf(resFileContentTemplate, titleName)

	// create file
	if _, err := os.Stat(resFileName); os.IsNotExist(err) {
		//file no exists
		err := os.WriteFile(resFileName, []byte(resFileContent), 0644)
		if err != nil {
			fmt.Println("Error creating response file:", err)
			return
		}
		fmt.Println("File", resFileName, "created successfully!")
	} else {
		fmt.Println("File", resFileName, "already exists!")
	}

}

func AppendContent() {
	fileName := "./file/generated.go"

	// 要添加的新代码内容
	// 要添加的新代码内容
	newCode := `
func additionalFunction() {
	fmt.Println("This is additional code.")
}
`

	// 读取原有文件内容
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 将新代码添加到文件末尾
	modifiedContent := append(fileContent, []byte(newCode)...)

	// 将修改后的内容写回文件
	err = os.WriteFile(fileName, modifiedContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Code added successfully to", fileName)

}

func GenerateMigrateFile(tableName string) {
	migrateDir := "./migrate/migratefile"

	if _, err := os.Stat(migrateDir); os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("%v no exists!", migrateDir))
		return
	}

	var titleName string
	split := strings.Split(tableName, "_")
	for _, s := range split {
		titleName = titleName + strings.Title(s)
	}

	timeStr := time.Now().Format(constant.TimeYmdHis)
	fileName := fmt.Sprintf("%v_%v", timeStr, tableName)
	migrateFileName := fmt.Sprintf("%v/%v.go", migrateDir, fileName)
	typeName := fmt.Sprintf("%v%v", strings.Title(titleName), timeStr)

	migrateFileContentTemplate := `package migratefile

type %s struct {
}

func (migrate *%s) Before() {
}

func (migrate *%s) After() {
}

func (migrate *%s) Run() {
}
`
	// replace
	migrateFileContent := fmt.Sprintf(migrateFileContentTemplate, typeName, typeName, typeName, typeName)

	// create file
	if _, err := os.Stat(migrateFileName); os.IsNotExist(err) {
		//file no exists
		err := os.WriteFile(migrateFileName, []byte(migrateFileContent), 0644)
		if err != nil {
			fmt.Println("Error creating response file:", err)
			return
		}
		fmt.Println("File", migrateFileName, "created successfully!")
	} else {
		fmt.Println("File", migrateFileName, "already exists!")
	}
}

func GenerateCurd(curdName string) {

	var titleName string
	split := strings.Split(curdName, "_")
	for _, s := range split {
		titleName = titleName + strings.Title(s)
	}

	//generate service
	serviceFileName := fmt.Sprintf("./service/%v_service.go", curdName)
	serviceFileContentTemplate := `package service

import (
	request "lancer/request/%s"
	response "lancer/response/%s"
)

type %sService struct {
}

func New%sService() *%sService {
	return &%sService{}
}

func (service *%sService) List(req *request.%sListRequest) (ret *response.%sListResponse, err error) {
	ret = new(response.%sListResponse)

	return
}

func (service *%sService) Create(req *request.%sCreateRequest) (ret *response.%sCreateResponse, err error) {
	ret = new(response.%sCreateResponse)

	return
}

func (service *%sService) Update(req *request.%sUpdateRequest) (ret *response.%sUpdateResponse, err error) {
	ret = new(response.%sUpdateResponse)

	return
}

func (service *%sService) Delete(req *request.%sDeleteRequest) (ret *response.%sDeleteResponse, err error) {
	ret = new(response.%sCreateResponse)

	return
}

`
	// replace
	serviceFileContent := fmt.Sprintf(serviceFileContentTemplate, curdName, curdName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName)

	// create file
	if _, err := os.Stat(serviceFileName); os.IsNotExist(err) {
		//file no exists
		err := os.WriteFile(serviceFileName, []byte(serviceFileContent), 0644)
		if err != nil {
			fmt.Println("Error creating service file:", err)
			return
		}
		fmt.Println("File", serviceFileName, "created successfully!")
	} else {
		fmt.Println("File", serviceFileName, "already exists!")
	}

	// generate controller
	controllerFileName := fmt.Sprintf("./controller/%v_controller.go", curdName)
	controllerFileContentTemplate := `package controller

import (
	"github.com/gin-gonic/gin"
	"lancer/constant"
	"lancer/plugin/translate"
	request "lancer/request/%s"
	response2 "lancer/response"
	response "lancer/response/%s"
	"lancer/service"
	"net/http"
)

type %sController struct {
}

func New%sController() *%sController {
	return &%sController{}
}

func (controller *%sController) List(c *gin.Context) {
	req := new(request.%sListRequest)
	ret := new(response.%sListResponse)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.New%sService()
	ret, err = srv.List(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

func (controller *%sController) Create(c *gin.Context) {
	req := new(request.%sCreateRequest)
	ret := new(response.%sCreateResponse)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.New%sService()
	ret, err = srv.Create(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

func (controller *%sController) Update(c *gin.Context) {
	req := new(request.%sUpdateRequest)
	ret := new(response.%sUpdateResponse)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.New%sService()
	ret, err = srv.Update(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

func (controller *%sController) Delete(c *gin.Context) {
	req := new(request.%sDeleteRequest)
	ret := new(response.%sDeleteResponse)
	err := c.ShouldBind(req)

	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.New%sService()
	ret, err = srv.Delete(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

`

	// replace
	controllerFileContent := fmt.Sprintf(controllerFileContentTemplate, curdName, curdName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName, titleName)
	if _, err := os.Stat(controllerFileName); os.IsNotExist(err) {
		err := os.WriteFile(controllerFileName, []byte(controllerFileContent), 0644)
		if err != nil {
			fmt.Println("Error creating controller file:", err)
			return
		}
		fmt.Println("File", controllerFileName, "created and successfully!")
	} else {
		fmt.Println("File", controllerFileName, "already exists!")
	}

	//generate request response
	GenerateReqRes(curdName, "list")
	GenerateReqRes(curdName, "create")
	GenerateReqRes(curdName, "update")
	GenerateReqRes(curdName, "delete")
	return
}
