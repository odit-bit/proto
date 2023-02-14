package resp

//resp is protocol to communicate via network
//https://redis.io/docs/reference/protocol-spec/
const (
	//resp-type header
	SIMPLE  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'

	//suffix or prefix
	cr   = '\r'
	lf   = '\n'
	crlf = "\r\n"
)

// represent valid resp-type
type Typ struct {
	prefix byte
	body   []byte
	array  []Typ
}

// if exist it will return the value of array
func (t Typ) Array() []Typ {
	if t.prefix == ARRAY {
		return t.array
	}
	return []Typ{}
}

// convert the body field value into string
func (t Typ) String() string {
	if t.prefix == BULK || t.prefix == SIMPLE || t.prefix == INTEGER || t.prefix == ERROR {
		return string(t.body)
	}
	return ""
}
