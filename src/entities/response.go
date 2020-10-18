package entities

type ResponseBag struct {
}

func (bag ResponseBag) New(bag_data map[string]interface{}) map[string]interface{} {

	data := map[string]interface{}{"status": "success", "data": bag_data}

	return data
}

func (bag ResponseBag) NewError(bag_data ErrorBag) map[string]interface{} {

	data := map[string]interface{}{"status": "failure", "error": bag_data}

	return data
}

type ResponseData struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}
