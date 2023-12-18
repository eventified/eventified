package dao

type NotFoundError struct {
	err string
}

func (err NotFoundError) Error() string {
	return err.err
}
