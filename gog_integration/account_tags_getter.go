package gog_integration

func getTagIds(atg accountTagsGetter) []string {
	at := atg.getAccountTags()
	tagIds := make([]string, 0, len(at))
	for _, tt := range at {
		tagIds = append(tagIds, tt.Id)
	}
	return tagIds
}

func getTagNames(tagIds []string, atg accountTagsGetter) map[string]string {
	at := atg.getAccountTags()
	tagNames := make(map[string]string, 0)
	for _, tagId := range tagIds {
		for _, tt := range at {
			if tt.Id == tagId {
				tagNames[tt.Id] = tt.Name
			}
		}
	}
	return tagNames
}
