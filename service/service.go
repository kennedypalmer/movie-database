package service 

import (
"Downloads/MovieDbProj/entities"

)

type Repository interface {
    NewMovie(movie entities.Movie) error 
    FindById(id string) (entities.Movie, error)
    DeleteById(id string) error
    UpdateById(id string, movie entities.Movie) error
}

type Service struct {
	Repo Repository 
}


func NewService(f Repository) Service {
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