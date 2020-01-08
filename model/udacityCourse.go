package model

type CourseCollection interface {
	getModel() []Courses
}

func (c UdacityCourseCollection) getModel() UdacityCourseCollection {
	return c;
}

type UdacityCourseCollection struct {
	Courses []struct {
		Key                  string `json:"key"`
		Title                string `json:"title"`
		Summary              string `json:"summary"`
		Image                string `json:"image"`
		/*Featured             bool `json:"featured"`
		Homepage             string `json:"homepage"`
		ProjectName          string `json:"project_name"`
		Subtitle             string `json:"subtitle"`
		ShortSummary         string `json:"short_summary"`
		Level                string `json:"level"`
		ExpectedLearning     string `json:"expected_learning"`
		TeaserVideo          struct {
					     YoutubeURL string `json:"youtube_url"`
				     } `json:"teaser_video"`
		Instructors          []interface{} `json:"instructors"`
		RelatedDegreeKeys    []interface{} `json:"related_degree_keys"`
		RequiredKnowledge    string `json:"required_knowledge"`
		Syllabus             string `json:"syllabus"`
		NewRelease           bool `json:"new_release"`
		ProjectDescription   string `json:"project_description"`
		FullCourseAvailable  bool `json:"full_course_available"`
		Faq                  string `json:"faq"`
		Affiliates           []interface{} `json:"affiliates"`
		Tracks               []interface{} `json:"tracks"`
		BannerImage          string `json:"banner_image"`
		Slug                 string `json:"slug"`
		Starter              bool `json:"starter"`
		ExpectedDurationUnit string `json:"expected_duration_unit"`
		ExpectedDuration     int `json:"expected_duration"`*/
	} `json:"courses"`
	Tracks  []struct {
		Courses     []string `json:"courses"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"tracks"`
	Degrees []struct {
		Instructors          []interface{} `json:"instructors"`
		Subtitle             string `json:"subtitle"`
		Key                  string `json:"key"`
		Image                string `json:"image"`
		ExpectedLearning     string `json:"expected_learning"`
		Featured             bool `json:"featured"`
		ProjectName          string `json:"project_name"`
		TeaserVideo          struct {
					     YoutubeURL string `json:"youtube_url"`
				     } `json:"teaser_video"`
		Title                string `json:"title"`
		RelatedDegreeKeys    []interface{} `json:"related_degree_keys"`
		RequiredKnowledge    string `json:"required_knowledge"`
		Syllabus             string `json:"syllabus"`
		NewRelease           bool `json:"new_release"`
		Homepage             string `json:"homepage"`
		ProjectDescription   string `json:"project_description"`
		FullCourseAvailable  bool `json:"full_course_available"`
		Faq                  string `json:"faq"`
		Affiliates           []interface{} `json:"affiliates"`
		Tracks               []interface{} `json:"tracks"`
		BannerImage          string `json:"banner_image"`
		ShortSummary         string `json:"short_summary"`
		Slug                 string `json:"slug"`
		Starter              bool `json:"starter"`
		Level                string `json:"level"`
		ExpectedDurationUnit string `json:"expected_duration_unit"`
		Summary              string `json:"summary"`
		ExpectedDuration     int `json:"expected_duration"`
	} `json:"degrees"`
}
