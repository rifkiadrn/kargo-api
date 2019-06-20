package processor

import (
	"testing"
	"time"

	. "github.com/golang/mock/gomock"
	"github.com/kargo-api/mocks"
	"github.com/kargo-api/models"
	"github.com/stretchr/testify/assert"
)

func TestGetSortedJobsReturnSortedJobs(t *testing.T) {
	ctrl := NewController(t)
	mock := mocks.NewMockRepository(ctrl)

	proc := &processor{
		repo: mock,
	}

	t1, _ := time.Parse(time.RFC3339, "2019-07-20T13:00:00Z")
	t2, _ := time.Parse(time.RFC3339, "2019-06-21T13:00:00Z")
	t3, _ := time.Parse(time.RFC3339, "2019-08-21T14:00:00Z")
	t4, _ := time.Parse(time.RFC3339, "2019-05-21T12:00:00Z")
	t5, _ := time.Parse(time.RFC3339, "2019-03-21T13:00:00Z")

	data := []models.Job{
		models.Job{
			ID:           1,
			Origin:       "Jakarta",
			Destination:  "Surabaya",
			ShipmentDate: t1,
			Description:  "",
			Budget:       1000,
			ShipperID:    1,
		},
		models.Job{
			ID:           2,
			Origin:       "Jakarta",
			Destination:  "Solo",
			ShipmentDate: t2,
			Description:  "",
			Budget:       1000,
			ShipperID:    1,
		},
		models.Job{
			ID:           3,
			Origin:       "Jakarta",
			Destination:  "Papua",
			ShipmentDate: t3,
			Description:  "",
			Budget:       1000,
			ShipperID:    1,
		},
		models.Job{
			ID:           4,
			Origin:       "Padang",
			Destination:  "Palembang",
			ShipmentDate: t4,
			Description:  "",
			Budget:       1000,
			ShipperID:    2,
		},
		models.Job{
			ID:           5,
			Origin:       "Medan",
			Destination:  "Bandung",
			ShipmentDate: t5,
			Description:  "",
			Budget:       1000,
			ShipperID:    3,
		},
	}

	mock.EXPECT().GetJobs().Return(data, nil)

	sortedJobs, err := proc.GetSortedJobs()
	for index := 1; index < len(sortedJobs); index++ {
		assert.True(t, sortedJobs[index].ShipmentDate.After(sortedJobs[index-1].ShipmentDate))
	}
	assert.Nil(t, err)
}

func TestGetSortedBidsReturnSortedBids(t *testing.T) {
	ctrl := NewController(t)
	mock := mocks.NewMockRepository(ctrl)

	proc := &processor{
		repo: mock,
	}

	data := []models.Bid{
		models.Bid{
			ID:            1,
			Price:         1000,
			Vehicle:       "pickup",
			Description:   "",
			JobID:         1,
			TransporterID: 1,
		},
		models.Bid{
			ID:            2,
			Price:         2000,
			Vehicle:       "cdd",
			Description:   "",
			JobID:         1,
			TransporterID: 2,
		},
		models.Bid{
			ID:            3,
			Price:         1000,
			Vehicle:       "pickup",
			Description:   "",
			JobID:         1,
			TransporterID: 3,
		},
		models.Bid{
			ID:            4,
			Price:         1000,
			Vehicle:       "pickup",
			Description:   "",
			JobID:         1,
			TransporterID: 3,
		},
		models.Bid{
			ID:            5,
			Price:         1000,
			Vehicle:       "pickup",
			Description:   "",
			JobID:         1,
			TransporterID: 3,
		},
	}

	mock.EXPECT().GetBids(int64(1)).Return(data, nil)

	sortedBids, err := proc.GetSortedBids(1)
	for index := 1; index < len(sortedBids); index++ {
		assert.True(t, sortedBids[index].Price >= sortedBids[index-1].Price)
	}
	assert.Nil(t, err)
}
