package robottravelsal

import "fmt"

type Robot struct {
	MaxX  int
	MaxY  int
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
			cordinate := r.State.NextCordinate()
			if cordinate.XCorditnate > r.MaxX || cordinate.YCordinate > r.MaxY {
				break
			}
			if _, ok := visited[cordinate]; ok {
				break
			}
			r.State.Move()

		} else {
			fmt.Println("Not Suported")
			break
		}
	}
}
