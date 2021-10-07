
package entities 

import "github.com/google/uuid"

	type Movie struct {
	Id string 
	Title string
	Genre []string
	Description string 
	Director string
	Actors []string
	Rating float64 
    }


	func (m *Movie) GetId(){
		m.Id = uuid.New().String()
	}
