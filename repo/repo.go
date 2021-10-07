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

   output, err := ioutil.ReadFile(f.Filename)
   if err != nil {
       return err 
   }
    

    //Unmarshaling the movie in json form from our request into our GoMovie database 
    
    err = json.Unmarshal(output, &movieDbSlice)
    if err != nil {
        return err 
    }

    
	//Appends the passed in movie to our movies slice in our GoMovie database 
	movieDbSlice.Movies = append(movieDbSlice.Movies, movie)




	//Insert functionality to marshal back into json file 
    input, err := json.MarshalIndent(movieDbSlice, "", " ")
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(f.Filename, input, 0644)
    if err != nil {
        return err
    }

	return nil
}



//Method on our json file that takes in the id and returns a movie back from our Get request
//Read the json file and unmarshal its contents into an instance of our movie database(theMovie)
//We range over the movies now to see which one matches the id we passed in 
//Line 104 resets the value of our instance of Movie created on line 100
//We then return that movie on line 109

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



//Ranging over all of the movies and finding the id that matches what is passed in 
//If the id in our db matches the id passed in, we iterate over it with continue
//Our loop defaults to setting the value of our deleteResult db to the movies that were iterated over- ignoring the movie w/ the given id 

func (f TheFile) DeleteById(id string) error {
    file, err := ioutil.ReadFile(f.Filename)
    if err != nil {
        log.Fatalln(err)
    }

    allMovies := GoMovieDb{} //Use this to unmarshal all movies into this slice to range over them

    err = json.Unmarshal(file, &allMovies)
    if err != nil {
        return err 
    }

    
    deleteResult := GoMovieDb{} //Use this to new slice to get our database without the movie who's id was entered

    
        for _, v := range allMovies.Movies {
          if v.Id == id {                   
            continue                 
          } else {                             
            deleteResult.Movies = append(deleteResult.Movies, v)
        }

        }

    ourNewDb, err := json.MarshalIndent(deleteResult, "", " ")
    if err != nil {
        log.Fatalln(err)
    }

    err = ioutil.WriteFile(f.Filename, ourNewDb, 0644)
    if err != nil {
        return err
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
            v.Id = movie.Id
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