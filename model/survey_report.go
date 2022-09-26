package model

import "time"

type SurveyReport struct {

	// Preset values
	StudentName   string    `db:"studentname"`
	SchoolName    string    `db:"schoolname"`
	DateOfSurvey  time.Time `db:"dateofsurvey"`
	SCAPercentage float64   `db:"scapercentage"`
	DValue        float64   `db:"dvalue"`
	MValue        float64   `db:"mvalue"`
	FValue        float64   `db:"fvalue"`

	// Calculated values
	RiskProfile                 string
	OperatorSuggestionRecurring string
	OperatorSuggestionFluoride  string
	OperatorSuggestionDiet      string
	OperatorSuggestionSealant   string
	OperatorSuggestionART       string
	ParentReminder              []string
	ParentGuidance              []string
	ParentSupervision           []string
	TeacherReminder             []string
	TeacherGuidance             []string
}

func (sr *SurveyReport) Setup() {
	sr.RiskProfile = "low"

	if sr.SCAPercentage > 66 {
		sr.RiskProfile = "high"
	} else if sr.SCAPercentage > 33 && sr.SCAPercentage <= 66 {
		sr.RiskProfile = "medium"
	}

	switch sr.RiskProfile {

	case "low":
		sr.ParentReminder = []string{
			"Orang tua mengingatkan agar kontrol ke dokter gigi setiap 6 bulan sekali",
		}
		sr.ParentGuidance = []string{
			"Orang tua mengajarkan cara menyikat gigi yang benar",
			"Orang tua mengingatkan agar menyikat gigi 2x sehari dengan pasta gigi ber fluoride",
		}
		sr.ParentSupervision = []string{
			"Orang tua memberikan pengawasan terhadap makanan manis dan lengket yang dikonsumsi sehari - hari",
		}
		sr.TeacherReminder = []string{
			"Guru mengingatkan agar kontrol ke dokter gigi setiap 6 bulan sekali",
		}
		sr.TeacherGuidance = []string{
			"Guru mengajarkan cara menyikat gigi yang benar",
			"Guru mengingatkan agar menyikat gigi 2x sehari dengan pasta gigi ber fluoride",
		}
		sr.OperatorSuggestionRecurring = "setiap 6-12 bulan"
		sr.OperatorSuggestionFluoride = "pasta gigi 2x sehari"
		sr.OperatorSuggestionDiet = "pemeliharaan asupan diet"
		sr.OperatorSuggestionSealant = "fissure sealant dilakukan jika diperlukan"
		sr.OperatorSuggestionART = "pengawasan karies baru"

	case "medium":
		sr.ParentReminder = []string{
			"Orang tua mengingatkan agar kontrol ke dokter gigi setiap 4-6 bulan sekali",
		}
		sr.ParentGuidance = []string{
			"Orang tua mengajarkan cara menyikat gigi yang benar",
			"Orang tua mengingatkan agar menyikat gigi 2x sehari dengan pasta gigi ber fluoride",
			"Orang tua mengingatkan agar dilakukan perawatan topical aplikasi fluoride",
		}
		sr.ParentSupervision = []string{
			"Orang tua melakukan diet makanan manis dan lengket yang dikonsumsi sehari- hari",
		}
		sr.TeacherReminder = []string{
			"Guru mengingatkan agar kontrol ke dokter gigi setiap 4-6 bulan sekali",
		}
		sr.TeacherGuidance = []string{
			"Guru mengajarkan cara menyikat gigi yang benar",
			"Guru mengingatkan agar menyikat gigi 2x sehari dengan pasta gigi ber fluoride",
			"Guru mengingatkan agar dilakukan perawatan topical aplikasi fluoride",
		}
		sr.OperatorSuggestionRecurring = "setiap 4-6 bulan"
		sr.OperatorSuggestionFluoride = "pasta gigi 2x sehari + Topikal aplikasi"
		sr.OperatorSuggestionDiet = "diet dengan pengawasan"
		sr.OperatorSuggestionSealant = "fissure sealant dilakukan jika diperlukan"
		sr.OperatorSuggestionART = "pengawasan karies baru + restorasi dari kavitas baru"

	case "high":
		sr.ParentReminder = []string{
			"Orang tua mengingatkan agar kontrol ke dokter gigi setiap 3-4 bulan sekali",
		}
		sr.ParentGuidance = []string{
			"Orang tua mengajarkan cara menyikat gigi yang benar",
			"Orang tua mengingatkan agar menyikat gigi 2x sehari dengan pasta gigi ber fluoride",
			"Orang tua mengingatkan agar dilakukan perawatan topical aplikasi fluoride",
		}
		sr.ParentSupervision = []string{
			"Orang tua melakukan diet makanan manis dan lengket yang dikonsumsi sehari- hari",
			"Orang tua mengganti konsumsi permen yang manis dengan permen xylitol",
		}
		sr.TeacherReminder = []string{
			"Guru mengingatkan agar kontrol ke dokter gigi setiap 3-4 bulan sekali",
		}
		sr.TeacherGuidance = []string{
			"Guru mengajarkan cara menyikat gigi yang benar",
			"Guru mengingatkan agar menyikat gigi 2x sehari dengan pasta gigi ber fluoride",
			"Guru mengingatkan agar dilakukan perawatan topical aplikasi fluoride",
		}
		sr.OperatorSuggestionRecurring = "setiap 3-4 bulan"
		sr.OperatorSuggestionFluoride = "topikal aplikasi + pasta gigi 2x sehari"
		sr.OperatorSuggestionDiet = "diet dengan pengawasan + xylitol"
		sr.OperatorSuggestionSealant = "direkomendasikan fissure sealant"
		sr.OperatorSuggestionSealant = "pengawasan karies baru + restorasi dari kavitas baru"

	}
}
