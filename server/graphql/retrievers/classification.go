package retrievers

import (
	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database"
	"github.com/olyop/objectgraph/objectgraph"
)

type RetrieveClassification struct{}

func (*RetrieveClassification) ByID(args objectgraph.RetrieverArgs) (*database.Classification, error) {
	classificationID := args.GetPrimary().(uuid.UUID)

	classification, err := database.SelectClassificationByID(classificationID)
	if err != nil {
		return nil, err
	}

	return classification, nil
}

func (*RetrieveClassification) ByIDs(args objectgraph.RetrieverArgs) ([]*database.Classification, error) {
	classificationIDs := args.GetPrimary().([]uuid.UUID)

	classifications, err := database.SelectClassificationsByIDs(classificationIDs)
	if err != nil {
		return nil, err
	}

	return classifications, nil
}
