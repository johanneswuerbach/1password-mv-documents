package op

// Example JSON
// {
// 	"uuid": "UUID",
// 	"templateUuid": "006",
// 	"trashed": "N",
// 	"createdAt": "2016-08-03T18:45:47Z",
// 	"updatedAt": "2019-12-27T13:09:51Z",
// 	"changerUuid": "UUID",
// 	"itemVersion": 3,
// 	"vaultUuid": "UUID",
// 	"overview": {
// 		"ainfo": "513 bytes",
// 		"ps": 0,
// 		"tags": [
// 			"nested/tag",
// 			"test"
// 		],
// 		"title": "backup-codes.txt - Some website"
// 	}
// },

type Overview struct {
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
}

type Document struct {
	UUID     string   `json:"uuid"`
	Overview Overview `json:"overview"`
}

type Item struct {
	Details ItemDetails `json:"details"`
}

type ItemDetails struct {
	DocumentAttributes DocumentAttributes `json:"documentAttributes"`
}

type DocumentAttributes struct {
	FileName string `json:"fileName"`
}
