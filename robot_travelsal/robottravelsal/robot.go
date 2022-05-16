package robottravelsal

import "fmt"

type Robot struct {
	State IRobotState
}

func (r *Robot) Move(m string) {
	visited := make(map[Corditnate]struct{})

	for _, s1 := range m {
		s := string(s1)
		if s == "R" {
			r.State = r.State.ChangeDirection(s)
		} else if s == "L" {
			r.State = r.State.ChangeDirection(s)
		} else if s == "M" {
			cordinate := r.State.GetCordinate()
			r.State.Move()
			if _, ok := visited[r.State.GetCordinate()]; ok {
				r.State.SetCordinate(cordinate)
				break
			}
		} else {
			fmt.Println("Not Suported")
			break
		}
	}
}
