// Code generaTed by fileb0x at "2019-06-18 18:15:57.157932567 +0200 CEST m=+0.117198615" from config file "b0x.yaml" DO NOT EDIT.
// modified(2019-06-18 17:41:42.127453346 +0200 CEST)
// original path: swagger-ui/dist/swagger-ui-bundle.js

package swaggerFiles

import (
  
  "os"
)

// FileB0xfileSwaggerUIBundleJs is "b0xfile__/swagger-ui-bundle.js"

func init() {
  

  f, err := FS.OpenFile(CTX, "b0xfile__/swagger-ui-bundle.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    panic(err)
  }

  
  _, err = f.Write(FileB0xfileSwaggerUIBundleJs)
  if err != nil {
    panic(err)
  }
  

  err = f.Close()
  if err != nil {
    panic(err)
  }
}
