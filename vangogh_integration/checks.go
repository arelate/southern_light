package vangogh_integration

func IsGOGPagedProduct(pt ProductType) bool {
	return containsProductType(GOGPagedProducts(), pt)
}

func IsArrayProduct(pt ProductType) bool {
	return IsGOGArrayProduct(pt) ||
		IsHltbArrayProduct(pt)
}

func IsGOGArrayProduct(pt ProductType) bool {
	return containsProductType(GOGArrayProducts(), pt)
}

func IsHltbArrayProduct(pt ProductType) bool {
	return containsProductType(HltbArrayProducts(), pt)
}

func IsGOGDetailProduct(pt ProductType) bool {
	return containsProductType(GOGDetailProducts(), pt)
}

func IsGOGProduct(pt ProductType) bool {
	return IsGOGPagedProduct(pt) ||
		IsGOGArrayProduct(pt) ||
		IsGOGDetailProduct(pt)
}

func IsSteamDetailProduct(pt ProductType) bool {
	return containsProductType(SteamDetailProducts(), pt)
}

func IsSteamProduct(pt ProductType) bool {
	return IsSteamDetailProduct(pt)
}

func IsPcgwDetailProduct(pt ProductType) bool {
	return containsProductType(PcgwDetailProducts(), pt)
}

func IsHltbDetailProduct(pt ProductType) bool {
	return containsProductType(HltbDetailProducts(), pt)
}

func IsPcgwProduct(pt ProductType) bool {
	return IsPcgwDetailProduct(pt)
}

func IsHltbProduct(pt ProductType) bool {
	return IsHltbArrayProduct(pt) ||
		IsHltbDetailProduct(pt)
}

func IsProtonDBDetailProduct(pt ProductType) bool {
	return containsProductType(ProtonDBDetailProducts(), pt)
}

func IsProtonDBProduct(pt ProductType) bool {
	return IsProtonDBDetailProduct(pt)
}

func IsFastPageFetchProduct(pt ProductType) bool {
	return containsProductType(FastPageFetchProducts(), pt)
}

func IsProductRequiresAuth(pt ProductType) bool {
	return containsProductType(requireAuth, pt)
}

func IsImageRequiresAuth(it ImageType) bool {
	for _, itra := range imageTypeRequiresAuth {
		if itra == it {
			return true
		}
	}
	return false
}

func IsGetItemsSupported(pt ProductType) bool {
	return containsProductType(supportsGetItems, pt)
}

func IsImageTypeSupported(pt ProductType, it ImageType) bool {
	if !IsValidProductType(pt) ||
		!IsValidImageType(it) {
		return false
	}

	supportedIts, ok := supportedImageTypes[pt]
	if !ok {
		return false
	}

	for _, sit := range supportedIts {
		if sit == it {
			return true
		}
	}

	return false
}

func containsProductType(all []ProductType, pt ProductType) bool {
	for _, apt := range all {
		if apt == pt {
			return true
		}
	}
	return false
}

func IsSupportedProperty(pt ProductType, property string) bool {
	for _, supportedProperty := range supportedProperties[pt] {
		if property == supportedProperty {
			return true
		}
	}
	return false
}
