package utils

type Test_message struct {

}

func (tm *Test_message) Reset()  {
	*tm = Test_message{}
}
func (tm *Test_message) String() string {
	return "dasd"
}
func (tm *Test_message) ProtoMessage()  {

}
func NewMessage() *Test_message {
   testMessage := new(Test_message)
   return testMessage
}

