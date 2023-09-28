package struc

type ResponseJson struct {
	Status string      `json:s`
	Data   interface{} `json:d`
}
