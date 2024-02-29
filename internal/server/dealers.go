package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Dealer struct {
	Id                string `json:"id"`
	DealerName        string `json:"dealer"`
	StreetAddress     string `json:"street_address"`
	State             string `json:"state"`
	Zip               string `json:"zip"`
	GeneralPhone      string `json:"general_phone"`
	DirectorOfService string `json:"director_of_service"`
	LogoURL           any    `json:"logo_url"`
}

type DealerTable struct {
	TableHeaders []string
	Dealers      []Dealer
}

var tblHdrs = []string{"dealer", "address", "state", "zip", "phone", "director of service"}

func (s *Server) DealerHandler(rw http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Dealers", Endpoint: "dealers"}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) DealerGetTable(rw http.ResponseWriter, r *http.Request) {
	var sampleListString = "[{\"id\":\"ad9b0818-c551-42b8-a106-c9c59d881e69\",\"dealer\":\"Gabrielli \",\"street_address\":\"239 Bergen Turnpike, Ridgefield Park\",\"state\":\"NJ\",\"zip\":\"07660\",\"general_phone\":\"201-641-4440\",\"director_of_service\":\"Anthony Roy\",\"logo_url\":null}, \n {\"id\":\"2963e249-1267-4635-9157-90db76970958\",\"dealer\":\"Greatwest Kenworth\",\"street_address\":\"5909 6 St SE, Calgary\",\"state\":\"AB\",\"zip\":\"T2H 1M4\",\"general_phone\":\"403-253-7555\",\"director_of_service\":\"Mike Allen\",\"logo_url\":null}, \n {\"id\":\"2671aaea-358c-4bd1-bb0f-bed7c71e3cb1\",\"dealer\":\"Inland-Canada\",\"street_address\":\" Gloucester St, Burnaby\",\"state\":\"BC\",\"zip\":\"V3M\",\"general_phone\":\"604-291-6431\",\"director_of_service\":\"Trevor Matthews / Jerome Wasilieff / Barry Duncan\",\"logo_url\":null}, \n {\"id\":\"0814a2d4-3a9a-442b-b964-910cbb096210\",\"dealer\":\"Inland-USA\",\"street_address\":\"9730 Cherry Ave, Phoenix\",\"state\":\"CA\",\"zip\":\"92335\",\"general_phone\":\"619-755-9956\",\"director_of_service\":\"Matt Allen\",\"logo_url\":null}, \n {\"id\":\"f8ca0d46-9f60-4e08-824d-f92b8bf62a36\",\"dealer\":\"Kenworth Northeast Group,Inc.\",\"street_address\":\"100 Commerce Dr, Stoughton\",\"state\":\"NY\",\"zip\":\"14218\",\"general_phone\":\"800-688-3380\",\"director_of_service\":\"Dan Ray\",\"logo_url\":null}, \n {\"id\":\"155ac1c2-f786-4dea-9f47-7b645f2340fd\",\"dealer\":\"Kenworth of Louisiana\",\"street_address\":\"1365 Northwest Dr, Port Allen\",\"state\":\"LA\",\"zip\":\"70767\",\"general_phone\":\"225-303-0440\",\"director_of_service\":\"Jeramie Thibodeaux\",\"logo_url\":null}, \n {\"id\":\"79cc73e5-7f44-4d85-835f-8356da8dcf24\",\"dealer\":\"Kenworth of Pennsylvania\",\"street_address\":\"198 Kost Rd, Mount Joy\",\"state\":\"PA\",\"zip\":\"17015\",\"general_phone\":\"717-766-8000\",\"director_of_service\":\"Tony Wiser\",\"logo_url\":null}, \n {\"id\":\"c2188ea0-1768-484e-a5df-2a6cd9974488\",\"dealer\":\"Kenworth of Richfield\",\"street_address\":\"2890 Brecksville Rd, Richfield\",\"state\":\"OH\",\"zip\":\"44286\",\"general_phone\":\"330-659-4123\",\"director_of_service\":\"Denise Vargo\",\"logo_url\":null}, \n {\"id\":\"ad31f81e-2d5d-468c-a964-d29c20678320\",\"dealer\":\"Kenworth Sales\",\"street_address\":\"2125 Constitution Blvd, Salt Lake County\",\"state\":\"UT\",\"zip\":\"84119\",\"general_phone\":\"801-487-4161\",\"director_of_service\":\"Colby Shaw\",\"logo_url\":null}, \n {\"id\":\"05b044c8-500d-4e27-b458-57affa81c9e4\",\"dealer\":\"Kenworth Truck Centres\",\"street_address\":\"500 Creditstone Rd, Vaughan\",\"state\":\"ON\",\"zip\":\"L4K 5P9\",\"general_phone\":\"905-695-0740\",\"director_of_service\":\"Sean Warren\",\"logo_url\":null}, \n {\"id\":\"847da7cb-094d-4289-8ba9-563d3a5f9415\",\"dealer\":\"MHC\",\"street_address\":\"11120 Tomahawk Creek Pkwy, Tupelo\",\"state\":\"KS\",\"zip\":\"66211\",\"general_phone\":\"816-483-6444\",\"director_of_service\":\"Brandon Parsons\",\"logo_url\":null}, \n {\"id\":\"c76cfbb9-d5cb-4575-bdac-52f6b71d3010\",\"dealer\":\"Palmer Trucks\",\"street_address\":\"2929 S Holt Rd, Indianapolis\",\"state\":\"IN\",\"zip\":\"46241\",\"general_phone\":\"317-247-8421\",\"director_of_service\":\"Joe Carter\",\"logo_url\":null}, \n {\"id\":\"b38f6f60-c16c-4160-9003-85c62bca5d39\",\"dealer\":\"Pape\",\"street_address\":\"355 Goodpasture Island Rd, Lane County\",\"state\":\"OR\",\"zip\":\"97401\",\"general_phone\":\"541-683-5073\",\"director_of_service\":\"Wes Sage\",\"logo_url\":null}, \n {\"id\":\"a3939231-2889-4b48-b440-105802090ea7\",\"dealer\":\"Papé Kenworth Alaska\",\"street_address\":\"355 Goodpasture Island Rd, Eugene\",\"state\":\"OR\",\"zip\":\"97401\",\"general_phone\":\"541-683-5073\",\"director_of_service\":\"Wes Sage\",\"logo_url\":null}, \n {\"id\":\"a0ad22d0-0919-4091-945a-5805c9e4847e\",\"dealer\":\"Papé Kenworth Northwest\",\"street_address\":\"355 Goodpasture Island Rd, Eugene\",\"state\":\"OR\",\"zip\":\"97401\",\"general_phone\":\"541-683-5073\",\"director_of_service\":\"Wes Sage\",\"logo_url\":null}, \n {\"id\":\"f44defac-93e1-4692-85c9-a9097da48679\",\"dealer\":\"Rihm Kenworth\",\"street_address\":\"425 Concord St S, St Paul\",\"state\":\"MN\",\"zip\":\"55075\",\"general_phone\":\"800-988-8235\",\"director_of_service\":\"Kenny Huff (North) / Doug Conniff (South)\",\"logo_url\":null}, \n {\"id\":\"4c988662-b324-4926-bfbf-4f19e4b09b29\",\"dealer\":\"Tri-State Kenworth\",\"street_address\":\"1 Depot Hill Rd, Enfield\",\"state\":\"CT\",\"zip\":\"06082\",\"general_phone\":\"860-627-8030\",\"director_of_service\":\"Don Longyhore\",\"logo_url\":null}, \n {\"id\":\"02e2639f-3d4e-4720-b4df-535d449df956\",\"dealer\":\"Truck Enterprises, Inc.\",\"street_address\":\"3440 S Main St, Harrisonburg\",\"state\":\"VA\",\"zip\":\"22801\",\"general_phone\":\"804-317-8449\",\"director_of_service\":\"Lee Hawkins\",\"logo_url\":null}, \n {\"id\":\"dd66344f-9adf-4934-8d1a-b16e41fda751\",\"dealer\":\"Truckworx Kenworth\",\"street_address\":\"2220 Finley Blvd, Birmingham\",\"state\":\"AL\",\"zip\":\"35234\",\"general_phone\":\"800-444-6170\",\"director_of_service\":\"(Director of Customer Service) Jim Taylor\",\"logo_url\":null}, \n {\"id\":\"6464d27f-44f1-45e2-928b-cd5be8efaf1b\",\"dealer\":\"Wallwork Trucks\",\"street_address\":\"900 35th St N, Fargo\",\"state\":\"ND\",\"zip\":\"58102\",\"general_phone\":\"701-476-7000\",\"director_of_service\":\"Brian Meier\",\"logo_url\":null}, \n {\"id\":\"f2750be3-79d3-4c7b-adfe-dcfba0194d4a\",\"dealer\":\"Whiteford Kenworth\",\"street_address\":\"4625 Western Ave, South Bend\",\"state\":\"IN\",\"zip\":\"46619\",\"general_phone\":\"574-234-9007\",\"director_of_service\":\"Doug Bennitt\",\"logo_url\":null}, \n {\"id\":\"2147ab14-8af0-46f5-b93d-400a4ce12632\",\"dealer\":\"Worldwide Equipment\",\"street_address\":\"18285 Lee Hwy, Abingdon\",\"state\":\"VA\",\"zip\":\"24210\",\"general_phone\":\"555-555-5555\",\"director_of_service\":\"Gary Coffman\",\"logo_url\":null}]"
	dt := DealerTable{
		TableHeaders: tblHdrs,
		Dealers:      []Dealer{},
	}
	var err = json.Unmarshal([]byte(sampleListString), &dt.Dealers)
	if err != nil {
		fmt.Println(err)
	}

	err = s.tmpl.ExecuteTemplate(rw, "dealer-content", dt)
	if err != nil {
		log.Print(err)
	}
}

// DB services and what not
