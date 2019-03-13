package filelisting

import (
	"os"
	"net/http"
	"io/ioutil"
	"strings"
)

const prefix = "/list/"



type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {

	if strings.Index(request.URL.Path, prefix) != 0 {
		// 返回自定义的 error
		return userError("path must start with " + prefix)
	}

	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	writer.Write(all)

	return nil
}
