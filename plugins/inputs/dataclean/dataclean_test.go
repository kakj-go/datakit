package dataclean

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"

	"github.com/gin-gonic/gin"
)

const (
	pt1 = `point01,t1=tags10,t2=tags20 f1=11i,f2=true,f3="hello" 1602581410306591000`
	pt2 = `point02,t1=tags10,t2=tags20 f1=11i,f2=true,f3="hello" 1602581410306591000`
	ob1 = `{"source":"dk1", "status":200}`
)

func TestMain(t *testing.T) {
	io.TestOutput()

	tmpfile, err := ioutil.TempFile("", "example.*.lua")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if err := ioutil.WriteFile(tmpfile.Name(), []byte(luaCode), 0644); err != nil {
		t.Fatal(err)
	}

	dataclean := DataClean{
		Path: "/dataclean",
		// PointsLuaFiles: []string{tmpfile.Name()},
		ObjectLuaFiles: []string{tmpfile.Name()},
	}

	{
		router := gin.New()
		router.POST(dataclean.Path, handle)
		httpsrv := &http.Server{
			Addr:    "0.0.0.0:9999",
			Handler: router,
		}
		go func() {
			if err := httpsrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				t.Error(err)
			}
		}()
		defer httpsrv.Close()
	}

	time.Sleep(1 * time.Second)

	go dataclean.Run()

	http.Post("http://127.0.0.1:9999/dataclean?category=/v1/write/object", "application/json; charset=utf-8",
		strings.NewReader(ob1))

	// http.Post("http://127.0.0.1:9999/dataclean?category=/v1/write/metric", "text/html; charset=utf-8",
	// 	strings.NewReader(pt1))

	// http.Post("http://127.0.0.1:9999/dataclean?category=/v1/write/logging", "text/html; charset=utf-8",
	// 	strings.NewReader(pt2))

	time.Sleep(3 * time.Second)

	datakit.Exit.Close()
}

// var luaCode = `
// function handle(points)
// 	for _, pt in pairs(points) do
// 		print("name", pt.name)
// 		print("time", pt.time)
// 		print("-------\ntags:")
// 		for k, v in pairs(pt.tags) do
// 			print(k, v)
// 		end
// 		print("-------\nfields:")
// 		for k, v in pairs(pt.fields) do
// 			print(k, v)
// 		end
// 		print("------------------------")
// 	end
// 	return points
// end
// `

var luaCode = `
function handle(xxx)
	for key, element in pairs(xxx) do
		print("key:     ", key)
		print("element: ", element)
		element="modify data"
	end
	print("------------------------")
	return xxx
end
`
