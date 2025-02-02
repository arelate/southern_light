package vangogh_integration

import (
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
)

func NewReduxWriter(properties ...string) (redux.Writeable, error) {
	rdp, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}
	return redux.NewWriter(rdp, properties...)
}

func NewReduxReader(properties ...string) (redux.Readable, error) {
	rdp, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}
	return redux.NewReader(rdp, properties...)
}
