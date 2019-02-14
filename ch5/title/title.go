package title

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func title(url string)error{
	resp,err := http.Get(url)
	if err!=nil{
		return err
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	_,err = html.Parse(resp.Body)
	if err!=nil{
		return fmt.Errorf("parse %s as HTML:%v",url,err)
	}
	return err
}