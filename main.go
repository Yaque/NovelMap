package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type RequestData struct {
	Name string
}

type ResponseData struct {
	Message string
	Status  int
}

type RequestDataPW struct {
	PW string
}

type ResponseDataPW struct {
	Data PaperList
}

type PaperList struct {
	Names []string
	Count int
}

type PenPos struct {
	X float32
	Y float32
}

type OnePen struct {
	OnePen    []PenPos
	Scale     float32
	Color     string
	LineWidth int
}

type AllPenData struct {
	Name string
	Data []OnePen
}

// func pingServer(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
// 	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
// 	w.Header().Set("content-type", "application/json")             //返回数据格式是json

// 	var requestData RequestData
// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Println(requestData.Name)

// 	resData := ResponseData{Message: "OK", Status: 1}
// 	resJsonData, _ := json.Marshal(resData)
// 	w.Write(resJsonData)
// }

func pingServer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	r.ParseForm()
	fmt.Println("收到客户端请求: ", r.Form)
}

// 创建一张纸
func createPaper(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(requestData.Name)

	writerJsonFile(AllPenData{Name: requestData.Name})

	w.Header().Set("Content-Type", "application/json")

	resData := ResponseData{Message: "OK", Status: 1}
	resJsonData, _ := json.Marshal(resData)
	w.Write(resJsonData)
}

func getPaper(w http.ResponseWriter, r *http.Request) {
	var requestDataPW RequestDataPW
	err := json.NewDecoder(r.Body).Decode(&requestDataPW)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(requestDataPW)

	data := getDirNumber()

	w.Header().Set("Content-Type", "application/json")

	resJsonData, _ := json.Marshal(ResponseDataPW{Data: data})

	w.Write(resJsonData)
}

// 创建一张纸
func deletePaper(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(requestData.Name)

	os.Remove("./data/" + requestData.Name)

	w.Header().Set("Content-Type", "application/json")

	resData := ResponseData{Message: "OK", Status: 1}
	resJsonData, _ := json.Marshal(resData)
	w.Write(resJsonData)
}

// 数据写入纸张
func uploadAllPenData(w http.ResponseWriter, r *http.Request) {
	var allPenData AllPenData
	err := json.NewDecoder(r.Body).Decode(&allPenData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(allPenData.Name)

	writerJsonFile(allPenData)

	w.Header().Set("Content-Type", "application/json")

	resData := ResponseData{Message: "OK", Status: 1}
	resJsonData, _ := json.Marshal(resData)
	w.Write(resJsonData)
}

// 数据写入纸张
func downloadAllPenData(w http.ResponseWriter, r *http.Request) {
	var allPenData AllPenData
	err := json.NewDecoder(r.Body).Decode(&allPenData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(allPenData.Name)

	allPenData = readJsonFile(allPenData.Name)

	w.Header().Set("Content-Type", "application/json")

	resJsonData, _ := json.Marshal(allPenData)
	w.Write(resJsonData)
}

func main() {

	fmt.Println("http://127.0.0.1/ui")

	fs := http.FileServer(http.Dir("ui/"))
	http.Handle("/ui/", http.StripPrefix("/ui/", fs))

	http.HandleFunc("/pingServer", pingServer)

	http.HandleFunc("/createPaper", createPaper)
	http.HandleFunc("/getPaper", getPaper)
	http.HandleFunc("/deletePaper", deletePaper)

	http.HandleFunc("/uploadAllPenData", uploadAllPenData)

	http.HandleFunc("/downloadAllPenData", downloadAllPenData)

	//http.HandleFunc("/ungo",myHandler2 )
	// addr：监听的地址
	// handler：回调函数
	http.ListenAndServe(":80", nil)
}

// 读json
func readJsonFile(filename string) AllPenData {
	var allPenData AllPenData
	file, err := os.OpenFile("data/"+filename+".json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//创建map，用于接收解码好的数据
	// dataMap1 := make(map[string]interface{})
	//创建文件的解码器
	decoder := json.NewDecoder(file)
	//解码文件中的数据，丢入dataMap所在的内存
	err8 := decoder.Decode(&allPenData)
	if err8 != nil {
		fmt.Println(err8)
		return allPenData
	}
	fmt.Println("解码成功", allPenData)
	return allPenData

}

// 写json
func writerJsonFile(allPenData AllPenData) {
	//打开文件

	file, _ := os.OpenFile("data/"+allPenData.Name+".json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file.Close()
	//创建encoder 数据输出到file中
	encoder := json.NewEncoder(file)
	//把dataMap的数据encode到file中
	err := encoder.Encode(allPenData)
	//异常处理
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("编码成功")
}

func getDirNumber() PaperList {
	count := 0 // 计数器
	var data []string

	files, err := os.ReadDir("data")
	// var count int
	for _, f := range files {
		fmt.Println(f.IsDir())
		if !f.IsDir() {

			data = append(data, f.Name())
			count++
		}

	}
	fmt.Println(count)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("文件夹%s下共有%d个文件\n", "data", count)
	}
	return PaperList{Names: data, Count: count}
}
