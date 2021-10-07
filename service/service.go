package service 

import (
"Downloads/MovieDbProj/entities"
"Downloads/MovieDbProj/repo"
)



type Service struct {
	Repo repo.TheFile
}


func NewService(f repo.TheFile) Service {
	return Service{
		Repo: f,
	}
}


func (s Service) PostMovie(movie entities.Movie)  error {
	 
	 err := s.Repo.NewMovie(movie)
	 if err != nil {
		 return err
	 }

	 return nil 
}



func (s Service) FindMovieById(id string) (entities.Movie, error) {
	getMovie, err := s.Repo.FindById(id)
	if err != nil {
		return getMovie, err
	}

  return  getMovie, nil 
}





func (s Service) DeleteMovieById(id string) error {
	 err := s.Repo.DeleteById(id)
	if err != nil {
		return err
	}
  return nil 
}

func (s Service) UpdateMovieById(id string, movie entities.Movie) error {

	err := s.Repo.UpdateById(id, movie)
	if err != nil {
		return err 
	}

	return nil 
}