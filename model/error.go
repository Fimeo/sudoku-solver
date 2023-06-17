package model

type Error struct {
	m string
	e error
}

func (err Error) Error() string {
	err.m = "model error : " + err.m
	if err.e != nil {
		return err.m + err.e.Error()
	}
	return err.m
}
