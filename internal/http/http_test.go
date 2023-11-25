package http

func TestSetupHTTPServer(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		want *gin.Engine
	}{

		{
			name: "TestSetupHTTPServer",
			want: &gin.Engine{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetupHTTPServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetupHTTPServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
