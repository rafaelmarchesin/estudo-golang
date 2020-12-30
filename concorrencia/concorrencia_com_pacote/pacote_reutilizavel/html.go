//Criei este arquivo para funcionar como um pacote para ser chamado facilmende dentro dos outros arguivos que vou criar
//Vamos ver se dá certo, né?

package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//Titulo ...
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}
