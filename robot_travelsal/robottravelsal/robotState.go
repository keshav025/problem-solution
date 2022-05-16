package robottravelsal

type IRobotState interface {
	Move()
	ChangeDirection(m string) IRobotState
	GetCordinate() Corditnate
	NextCordinate() Corditnate
	SetCordinate(c Corditnate)
	GetDirection() string
}

type Corditnate struct {
	XCorditnate int
	YCordinate  int
}

type RobotNorthState struct {
	Corditnate Corditnate
	Direction  string
}

func (r *RobotNorthState) NextCordinate() Corditnate {
	cord := r.Corditnate
	cord.YCordinate += 1
	return cord
}

func (r *RobotNorthState) Move() {
	r.Corditnate.YCordinate += 1
}

func (r *RobotNorthState) GetCordinate() Corditnate {
	return r.Corditnate
}

func (r *RobotNorthState) SetCordinate(c Corditnate) {
	r.Corditnate = c
}

func (r *RobotNorthState) GetDirection() string {
	return r.Direction
}

func (r *RobotNorthState) ChangeDirection(m string) IRobotState {
	if m == "R" {
		return &RobotEastState{r.Corditnate, "E"}
	} else if m == "S" {
		return &RobotWestState{r.Corditnate, "W"}
	}
	return nil
}

type RobotSouthState struct {
	Corditnate Corditnate
	Direction  string
}

func (r *RobotSouthState) NextCordinate() Corditnate {
	cord := r.Corditnate
	cord.YCordinate += 1
	return cord
}

func (r *RobotSouthState) GetCordinate() Corditnate {
	return r.Corditnate
}

func (r *RobotSouthState) SetCordinate(c Corditnate) {
	r.Corditnate = c
}

func (r *RobotSouthState) GetDirection() string {
	return r.Direction
}

func (r *RobotSouthState) ChangeDirection(m string) IRobotState {
	if m == "R" {
		return &RobotWestState{r.Corditnate, "W"}
	} else if m == "S" {
		return &RobotEastState{r.Corditnate, "E"}
	}
	return nil
}

func (r *RobotSouthState) Move() {
	r.Corditnate.YCordinate -= 1
}

type RobotEastState struct {
	Corditnate Corditnate
	Direction  string
}

func (r *RobotEastState) NextCordinate() Corditnate {
	cord := r.Corditnate
	cord.YCordinate += 1
	return cord
}

func (r *RobotEastState) SetCordinate(c Corditnate) {
	r.Corditnate = c
}

func (r *RobotEastState) GetDirection() string {
	return r.Direction
}

func (r *RobotEastState) GetCordinate() Corditnate {
	return r.Corditnate
}

func (r *RobotEastState) ChangeDirection(m string) IRobotState {
	if m == "R" {
		return &RobotSouthState{r.Corditnate, "S"}
	} else if m == "L" {
		return &RobotNorthState{r.Corditnate, "N"}
	}
	return nil
}

func (r *RobotEastState) Move() {
	r.Corditnate.XCorditnate += 1
}

type RobotWestState struct {
	Corditnate Corditnate
	Direction  string
}

func (r *RobotWestState) NextCordinate() Corditnate {
	cord := r.Corditnate
	cord.YCordinate += 1
	return cord
}

func (r *RobotWestState) SetCordinate(c Corditnate) {
	r.Corditnate = c
}

func (r *RobotWestState) GetDirection() string {
	return r.Direction
}

func (r *RobotWestState) ChangeDirection(m string) IRobotState {
	if m == "R" {
		return &RobotNorthState{r.Corditnate, "N"}
	} else if m == "L" {
		return &RobotSouthState{r.Corditnate, "S"}
	}
	return nil
}
func (r *RobotWestState) GetCordinate() Corditnate {
	return r.Corditnate
}
func (r *RobotWestState) Move() {
	r.Corditnate.XCorditnate -= 1
}

func GetRobotNorthState() *RobotNorthState {
	return &RobotNorthState{Corditnate{}, "N"}
}

func GetRobotSouthState() *RobotSouthState {
	return &RobotSouthState{Corditnate{}, "S"}
}
func GetRobotEastState() *RobotEastState {
	return &RobotEastState{Corditnate{}, "E"}
}
func GetRobotWestState() *RobotWestState {
	return &RobotWestState{Corditnate{}, "W"}
}
