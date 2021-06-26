package main

import "testing"

func TestMessageParse(t *testing.T) {
	message := "hello"
	data := createmessage(MessageTypeText, message)
	mtype, mlen, msg := readmessage(data)
	if mtype != MessageTypeText {
		t.Errorf("invald message type. got: %d, want: %d", mtype, MessageTypeText)
	}

	if mlen != uint32(len(message)) {
		t.Errorf("invalid length. got: %d, want: %d", mlen, len(message))
	}

	if msg != message {
		t.Errorf("invalid message. got: %s, want: %s", msg, message)
	}
}
