package repo 

import (
    "encoding/json"
    "io/ioutil"
    "Downloads/MovieDbProj/entities"
    "log"

)

type GoMovieDb struct {
	Movies []entities.Movie 

}


type TheFile struct {
    Filename string
}

func NewRepo(filename string) TheFile {
    return TheFile{
        Filename : filename, 
    }
}



func (f TheFile) NewMovie(movie entities.Movie)  error {

    movieDbSlice := GoMovieDb{}

    movie.GetId()

   file, err := ioutil.ReadFile(f.Filename)
   if err != nil {
       return err 
   }
    
    
    err = json.Unmarshal(file, &movieDbSlice)
    if err != nil {
        return err 
    }


	movieDbSlice.Movies = append(movieDbSlice.Movies, movie)




    ourNewMovie, err := json.MarshalIndent(movieDbSlice, "", " ")
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(f.Filename, ourNewMovie, 0644)
    if err != nil {
        return err
    }

	return nil
}




func (f TheFile) FindById(id string) (entities.Movie, error) {
    file, err := ioutil.ReadFile(f.Filename)
    if err != nil {
        log.Fatalln(err)
    }

    theMovies := GoMovieDb{}

    err = json.Unmarshal(file, &theMovies)


    getResult := entities.Movie{}

    for _, v := range theMovies.Movies {
        if v.Id == id {
                getResult = v
                return getResult, nil 
        }
    }
    return entities.Movie{}, err 
}



func (f TheFile) DeleteById(id string) error {
    file, err := ioutil.ReadFile(f.Filename)
    if err != nil {
        log.Fatalln(err)
    }

    allMovies := GoMovieDb{} 

    err = json.Unmarshal(file, &allMovies)
    if err != nil {
        return err 
    }
    
        for i, v := range allMovies.Movies {
          if v.Id == id {                   
            allMovies.Movies = append(allMovies.Movies[:i], allMovies.Movies[i+1:]...)   
            ourNewDb, err := json.MarshalIndent(&allMovies, "", " ")     
             if err != nil {
                log.Fatalln(err)
             }
            
            err = ioutil.WriteFile(f.Filename, ourNewDb, 0644)
               if err != nil {
                   return err
                }
          }
          
        
            
        }
    return nil 

}



func (f TheFile) UpdateById(id string, movie entities.Movie) error {
    file, err := ioutil.ReadFile(f.Filename) 
    if err != nil {
        return err
    }

    allMovies := GoMovieDb{} 

    err = json.Unmarshal(file, &allMovies) 
    if err != nil {
        return err 
    }

    
    for i, v := range allMovies.Movies {
        if v.Id == id {
            allMovies.Movies = append(allMovies.Movies[:i], allMovies.Movies[i+1:]...)
            v.Title = movie.Title 
            v.Genre = movie.Genre
            v.Description = movie.Description
            v.Director = movie.Director 
            v.Actors =  movie.Actors
            allMovies.Movies = append(allMovies.Movies, v)
        }
    }
    
          

    
      ourUpdatedMovie, err := json.MarshalIndent(&allMovies, "", " ")
      if err != nil {
          return err 
      }

      
      err = ioutil.WriteFile(f.Filename, ourUpdatedMovie, 0644)
      if err != nil {
          return err 
      }


  return nil 

}