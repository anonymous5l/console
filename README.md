Console
==========

golang colorful format print 

Example:
  ```go
	func main() {
		// for stop loop chan
		defer console.Abort()

	    console.Log("format %s", "test")
	    console.Ok("format %s", "test")
	    console.Err("format %s", "test")
	    console.Warn("format %s", "test")
	    console.Debug("format %s", "test")
	}
  ```
