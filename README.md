# unsafe-try-catch

Unsafe try-catch solution for Go. It does not work on newer versions.

# usage

```go
func CopyFile(src, dst string) error {
	var tc, tcr TryCatcher
	var err error

	if err = tc.Catch(); err != nil {
		return fmt.Errorf("copy %s %s: %v", src, dst, err)
	}
	
	r, err := os.Open(src)
	tc.Try(err)

	defer r.Close()

	if err = tcr.Catch(); err != nil {
		os.Remove(dst)
		return fmt.Errorf("copy %s %s: %v", src, dst, err)
	}
	
	/* tc may be reused */
	if err = tc.Catch(); err != nil {
		w.Close()
		os.Remove(dst)
		return fmt.Errorf("copy %s %s: %v", src, dst, err)
	}

	w, err := check os.Create(dst)
	tcr.Try(err) // remove and return err

	_, err = io.Copy(w, r)
	tc.Try(err) // close, remove and return err

	err = w.Close()
	tcr.Try(err) // remove and return err

	return nil
}
```
