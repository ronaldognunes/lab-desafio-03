package auction_test

import (
	"context"
	"testing"
	"time"

	"fullcycle-auction_go/configuration/database/mongodb"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database/auction"

	"github.com/stretchr/testify/assert"
)

func TestCreateAuction(t *testing.T) {
	ctx := context.Background()
	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)

	assert.Nil(t, err)

	auctionRepository := auction.NewAuctionRepository(databaseConnection)

	t.Setenv("AUCTION_INTERVAL", "10s")

	auctionEntity, err := auction_entity.CreateAuction("Carro", "Automóvel", "carro esportivo ", 0)
	assert.Nil(t, err)

	err = auctionRepository.CreateAuction(ctx, auctionEntity)
	assert.Nil(t, err)

	time.Sleep(11 * time.Second)

	auctionEntityConsulta, err := auctionRepository.FindAuctionById(ctx, auctionEntity.Id)

	assert.Nil(t, err)
	assert.Equal(t, auction_entity.Completed, auctionEntityConsulta.Status)

}

func TestGetAuctionInterval(t *testing.T) {
	t.Run("valid auction interval", func(t *testing.T) {
		t.Setenv("AUCTION_INTERVAL", "10m")
		duration := auction.GetAuctionInterval()
		assert.Equal(t, 10*time.Minute, duration)
	})

	t.Run("invalid auction interval", func(t *testing.T) {
		t.Setenv("AUCTION_INTERVAL", "invalid")
		duration := auction.GetAuctionInterval()
		assert.Equal(t, 5*time.Minute, duration)
	})
}
