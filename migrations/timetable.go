package platform

import (
	"fmt"
	"time"
)

type Lesson struct {
	Theme    string
	Comment  string
	Date     time.Time
	Duration time.Duration
	Cancel   bool
	Finish   bool
}

type Direction struct {
	Name string
}

type Group struct {
	Name      string
	Direction Direction
	Lessons   []Lesson
}

func isholidays(t time.Time, excludeSummer, excludeHolidays bool) bool {

	if excludeHolidays {
		if (t.Day() == 31 && t.Month() == time.December) || (t.Day() == 1 && t.Month() == time.January) || (t.Day() == 2 && t.Month() == time.January) || (t.Day() == 3 && t.Month() == time.January) || (t.Day() == 4 && t.Month() == time.January) || (t.Day() == 5 && t.Month() == time.January) || (t.Day() == 6 && t.Month() == time.January) || (t.Day() == 7 && t.Month() == time.January) || (t.Day() == 8 && t.Month() == time.January) || (t.Day() == 9 && t.Month() == time.January) || (t.Day() == 10 && t.Month() == time.January) {
			return true
		}

		if t.Day() == 23 && t.Month() == time.February {
			return true
		}

		if (t.Day() == 9 && t.Month() == time.May) || (t.Day() == 1 && t.Month() == time.May) {
			return true
		}

	}
	if excludeSummer {
		if t.Month() == time.June || t.Month() == time.July || t.Month() == time.August {
			return true
		}
	}

	return false
}

func (g *Group) Generate(startdate time.Time, weekday []time.Weekday, countLessions int, duration time.Duration, excludeSummer, excludeHolidays bool) {

	t := startdate
	for {
		if len(g.Lessons) == countLessions {
			break
		}
		for i := 0; i < len(weekday); i++ {
			if t.Equal(startdate) {
				lesson := new(Lesson)
				lesson.Date = t
				lesson.Duration = duration
				g.Lessons = append(g.Lessons, *lesson)
				fmt.Println(lesson.Date.Format("January 2, 2006"))
				break
			}
			if t.Weekday() == weekday[i] {
				if isholidays(t, excludeSummer, excludeHolidays) {
					fmt.Println(t.Format("January 2, 2006") + " Is Holidays")
					break
				}
				lesson := new(Lesson)
				lesson.Date = t
				lesson.Duration = duration
				g.Lessons = append(g.Lessons, *lesson)
				fmt.Println(lesson.Date.Format("January 2, 2006"))
			}
		}
		t = t.Add(time.Hour * 24)
	}
}
