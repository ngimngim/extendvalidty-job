package main

import "extendvalidity-job/Routes"

func main()  {
	r:=Routes.SetupRoutes()
	r.Run()
}