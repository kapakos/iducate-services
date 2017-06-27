package model

type CourseraCourseCollection struct {
	Courses []struct {
		Key string `json:"id"`
		Title string `json:"name"`
		Summary string `json:"description"`
		Image string `json:"photoUrl"`
		/*DomainTypes []struct {
		CourseType string `json:"courseType"`
		PartnerLogo string `json:"partnerLogo"`
			SubdomainID string `json:"subdomainId"`
			DomainID string `json:"domainId"`
		} `json:"domainTypes"`
		Categories []string `json:"categories"`
		Slug string `json:"slug"`
		InstructorIds []string `json:"instructorIds"`
		Specializations []interface{} `json:"specializations"`
		Workload string `json:"workload"`
		PrimaryLanguages []string `json:"primaryLanguages"`
		PartnerIds []string `json:"partnerIds"`
		Certificates []string `json:"certificates"`
		SubtitleLanguages []string `json:"subtitleLanguages"`
		S12NIds []interface{} `json:"s12nIds"`
		StartDate int64 `json:"startDate"`*/
	} `json:"elements"`
	Paging struct {
		Next string `json:"next"`
		Total int `json:"total"`
	} `json:"paging"`
	Linked struct {
	} `json:"linked"`
}

func (c *CourseraCourseCollection) GetCourses() *CourseraCourseCollection {
	return c;
}
