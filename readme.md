# DC (Dependency Container)

## A Golang Simple Dependency Container !!! 

```go
package main

import "github.com/aKorsi/dc"

func main(){
	container := dc.NewDC()

	// set a dependency
	container.SetDependency("sms.service", func() interface{} {
        return sms.NewService() 
    })

	// get a dependency
	myService := container.GetDependency("sms.service").(services.ISMSService)

	// delete a dependency
	container.DeleteDependency("sms.service")

	// clear container
	container.DeleteAll()
}
```