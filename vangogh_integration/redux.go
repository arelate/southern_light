package vangogh_integration

import (
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/pathways"
)

func NewReduxWriter(properties ...string) (kevlar.WriteableRedux, error) {
	rdp, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}
	return kevlar.NewReduxWriter(rdp, properties...)
}

func NewReduxReader(properties ...string) (kevlar.ReadableRedux, error) {
	rdp, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}
	return kevlar.NewReduxReader(rdp, properties...)
}
