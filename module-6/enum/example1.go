package main;

import "fmt";

type status string;

const (
	Approved status="Approved"
	Pending status="Pending"
	Declined status="Declined"
	Draft status="Draft"
)

func changeStatus(s status)(string){
	switch (s){
	case Approved:
	return "Approveds";
	break;
	case Pending:
	return "Pendings";
	break;
	case Declined:
	return "Declineds";
	break;
	default:
	return "Drafts";
	};
	return "not matches";
	
}

func main(){

fmt.Println(changeStatus(Pending))
	
}