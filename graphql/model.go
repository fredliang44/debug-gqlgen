package graphql

type Message struct {
	ID   string `json:"id" gorm:"primary_key:true;type:char(36);column:id"`
	Text string `json:"text"`

	TestMessage *string `json:"test_message"`
}
