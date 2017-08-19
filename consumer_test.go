package main

// func TestFeedBackConsumer(t *testing.T) {
// 	type args struct {
// 		cfg      Configuration
// 		messages chan []byte
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{name: "Default",
// 			args: args{cfg: Configuration{Broker: "127.0.0.1:9092", GroupID: "feedback_to_hipchat_test", Topic: "feedback_test"}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			FeedBackConsumer(tt.args.cfg, tt.args.messages)
// 		})
// 	}
// }
