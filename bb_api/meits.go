import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

req, err := http.NewRequest("GET", "https://gitit.post.ch/rest/api/1.0/projects/proxyservices/repos", nil)
if err != nil {
	// handle err
}
req.SetBasicAuth("user", "pass")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()