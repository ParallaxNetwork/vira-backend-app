package enums

type Status string

const (
	Ongoing 				Status = "ongoing"
	Crowdfunding		Status = "crowd funding"
	Finished				Status = "finished"
	Failed					Status = "failed"
	Harvested				Status = "harvested"
	HarvestFailed		Status = "harvest failed"
)