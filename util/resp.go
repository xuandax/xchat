package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RespOk(w http.ResponseWriter, data interface{}) {
	Resp(w, 0, "", data)
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, msg, nil)
}

func Resp(w http.ResponseWriter, code int, msg string, data interface{}) {
	r := R{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ret, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ret)
}
