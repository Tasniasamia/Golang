package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)
var db*pgx.Conn;
var err error;
type user struct{
	Id int `json:"id"`;
	Name string `json:"name`;
	Age int `json:"age"`;
	Cgpa float64 `json:"cgpa`;
}
type course struct{
	Id int `json:"id"`;
	Title string `json:"title"`;
	Credit int `json:credit`;
}
type enroll struct{
	Id int `json:"id"`;
	UserId int `json:"userid"`;
	CourseId int `json:"courseid"`;
	users user
	courses course
}
func main(){
	func(){
		db,err=pgx.Connect(context.Background(),"postgres://postgres:123456@localhost:5432/goDatabase");
		if err != nil {
			panic(err);
		}
		fmt.Println("Database Connected Successfully");
	}()

   mux :=http.NewServeMux();
   mux.HandleFunc("GET /courses",getCourses);
   mux.HandleFunc("POST /course",postCourse);
   mux.HandleFunc("DELETE /course/{id}",deleteCourse);
   mux.HandleFunc("POST /enroll",createEnroll);
   mux.HandleFunc("GET /user/{id}/courses",getCourseByUser);

   http.ListenAndServe(":5055",mux);

}

func getCourses(res http.ResponseWriter,req*http.Request){
	query :=`SELECT * FROM courses`;
	rows,err :=db.Query(context.Background(),query);
    if err != nil {
		res.WriteHeader(http.StatusInternalServerError);
		fmt.Fprintln(res,err);
		return;
	}
	var courseList[]course;
	for rows.Next(){
		var singleCourse course;
		err =rows.Scan(&singleCourse.Id,&singleCourse.Title,&singleCourse.Credit);
		if err != nil{
			res.WriteHeader(http.StatusInternalServerError);
			fmt.Println("error from here ",err)
			fmt.Fprintln(res,err);
			return;
		}
		courseList =append(courseList, singleCourse);

	}
	res.Header().Set("Content-Type","application/json");
	res.WriteHeader(http.StatusOK);
	json.NewEncoder(res).Encode(courseList);
}

func postCourse(res http.ResponseWriter,req*http.Request){
	decoder :=json.NewDecoder(req.Body);
	decoder.DisallowUnknownFields();
	var newCourse course;
	err:=decoder.Decode(&newCourse);
	if err !=nil{
		res.WriteHeader(http.StatusBadRequest);
		fmt.Fprintln(res,err);
		return;
	}

	query :=`INSERT INTO courses (title,credit) VALUES($1,$2) RETURNING *`;
	

	err=db.QueryRow(context.Background(),query,newCourse.Title,newCourse.Credit).Scan(&newCourse.Id,&newCourse.Title,&newCourse.Credit);
    if err != nil{
		res.WriteHeader(http.StatusInternalServerError);
		fmt.Println("error while inserting course ",err);
		fmt.Fprintln(res,err);
		return;
	}
   
	res.Header().Set("Content-Type","application/json");
	res.WriteHeader(http.StatusCreated);
	json.NewEncoder(res).Encode(newCourse)
}

func deleteCourse(res http.ResponseWriter,req*http.Request){
	paramsId :=req.PathValue("id");
	id,err :=strconv.Atoi(paramsId);

	if err!= nil {
		res.WriteHeader(http.StatusBadRequest);
		fmt.Fprintln(res,err);
		return;
	}
    query :=`SELECT id FROM courses WHERE id=$1`;
	err=db.QueryRow(context.Background(),query,id).Scan(&id)
	if err != nil{
		res.WriteHeader(http.StatusInternalServerError);
		fmt.Fprintln(res,err);
		return;
	}
	query=`DELETE FROM courses WHERE id=$1`;
	_,err=db.Exec(context.Background(),query,id);
	if err != nil {
        res.WriteHeader(http.StatusInternalServerError);
		fmt.Fprintln(res,err);
		return;
	}
	//res.Header().Set("Content-Type","application/json");
	res.WriteHeader(http.StatusOK);
	fmt.Fprintln(res,"Course Deleted Successfully");
	return;
}


func createEnroll(res http.ResponseWriter,req*http.Request){
	var newEnroll enroll;
	decoder :=json.NewDecoder(req.Body);
	decoder.DisallowUnknownFields();
	decoder.Decode(&newEnroll);
	query :=`INSERT INTO enrollment (userId,courseId) VALUES($1,$2) RETURNING *`;
	err :=db.QueryRow(context.Background(),query,newEnroll.UserId,newEnroll.CourseId).Scan(&newEnroll.Id,&newEnroll.UserId,&newEnroll.CourseId);
	if err != nil{
		res.WriteHeader(http.StatusInternalServerError);
		fmt.Fprintln(res,err);
		return;
	}
	res.Header().Set("Content-Type","application/json");
	res.WriteHeader(http.StatusCreated);
	json.NewEncoder(res).Encode(newEnroll)
}

func getCourseByUser(res http.ResponseWriter,req*http.Request){
	paramId :=req.PathValue("id");
	id,err :=strconv.Atoi(paramId);
	if err != nil{
		res.WriteHeader(http.StatusNotAcceptable);
		fmt.Fprintln(res,err);
		return;
	}

	query := `
    SELECT e.userid, e.courseid, u.id, u.name, c.id, c.title, c.credit 
    FROM enrollment e
    JOIN users u ON e.userid = u.id
    JOIN courses c ON e.courseid = c.id
    WHERE e.userid = $1
`;
  rows,err2:=db.Query(context.Background(),query,id);
  if err2 != nil{
		res.WriteHeader(http.StatusNotAcceptable);
		fmt.Fprintln(res,err);
		return;
  }
  var courseList [] enroll;
  for rows.Next(){
	var singleUser enroll;
	err=rows.Scan(&singleUser.Id,&singleUser.CourseId,&singleUser.users.Id,&singleUser.users.Name,&singleUser.courses.Id,&singleUser.courses.Title,&singleUser.courses.Credit);
	if err != nil {
    res.WriteHeader(http.StatusInternalServerError);
	fmt.Fprintln(res,err);
	return;
	}
   courseList=append(courseList, singleUser);
  }
  res.Header().Set("Content-Type","application/json");
	res.WriteHeader(http.StatusOK);
	json.NewEncoder(res).Encode(courseList)
}