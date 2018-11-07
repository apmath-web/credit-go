package viewModels

type ViewModelInterface interface {
	Fill(JsonData interface{}) (bool, error)
	Fetch() (interface{}, error)
	Validate() (bool, error)
	getValidation() (interface{}, error)
}
