package e2e

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func Test_GetAllEvent(t *testing.T) {
	client := http.Client{}

	resp, err := client.Get("http://localhost:5000/bmstu-stud-web/api/events/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
