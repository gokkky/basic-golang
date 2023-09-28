package struc

import "time"

type Student struct {
	ID        int       `json:id`
	Firstname string    `json:first_name`
	Lastname  string    `json:last_name`
	JoinDate  time.Time `json:join_date`
}

func (stu *Student) SetFirstname(firstname string) {
	stu.Firstname = firstname
}

func (stu Student) GetFirstname() string {
	return stu.Firstname
}

func NewStudent(id int, firstname string, lastname string, join time.Time) Student {
	return Student{
		ID:        id,
		Firstname: firstname,
		Lastname:  lastname,
		JoinDate:  join,
	}
}
