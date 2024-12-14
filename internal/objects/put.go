package objects

type Put struct {
	Bucket string
	Body   interface{}
	Config ObjectDetails
}
