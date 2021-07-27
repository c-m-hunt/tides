package tides

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func Test_createDate(t *testing.T) {
	now := time.Now()
	london, err := time.LoadLocation("Europe/London")
	if err != nil {
		log.Fatal(err)
	}
	type args struct {
		dateStr string
		timeStr string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"Test am", args{"Tues, 27th", "09:45am"}, time.Date(now.Year(), now.Month(), 27, 9, 45, 0, 0, london)},
		{"Test pm", args{"Tues, 27th", "09:45pm"}, time.Date(now.Year(), now.Month(), 27, 21, 45, 0, 0, london)},
		{"Test next month", args{"Tues, 1st", "09:45pm"}, time.Date(now.Year(), now.Month()+1, 1, 21, 45, 0, 0, london)},
		{"Test midday", args{"Tues, 1st", "12:45pm"}, time.Date(now.Year(), now.Month()+1, 1, 12, 45, 0, 0, london)},
		{"Test midnight", args{"Tues, 1st", "12:45am"}, time.Date(now.Year(), now.Month()+1, 1, 0, 45, 0, 0, london)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createDate(tt.args.dateStr, tt.args.timeStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createHeight(t *testing.T) {
	type args struct {
		heightString string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test height 1", args{"5.66m"}, 5.66},
		{"Test height 2", args{"0.8m"}, 0.8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createHeight(tt.args.heightString); got != tt.want {
				t.Errorf("createHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
