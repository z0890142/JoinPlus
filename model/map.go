package model

type FindPlaceResponse struct {
	Candidates []Place
	Status     string
}

type Place struct {
	Place_id string
}

type PlaceDetailResponse struct {
	Html_attributions []string
	result            string
}

type PlaceDetailResult struct {
	Html_attributions []string
	result            string
}
