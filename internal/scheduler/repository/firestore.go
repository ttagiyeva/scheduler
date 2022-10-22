package repository

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/firestore"
	"github.com/ttagiyeva/scheduler/internal/config"
	"github.com/ttagiyeva/scheduler/internal/scheduler/domain"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

const (
	dronePath = "drone_name"
)

type Firestore struct {
	log    *zap.SugaredLogger
	config *config.Config
	client *firestore.Client
}

//New creates Firestore Client and Firestore instance
func New(lc fx.Lifecycle, log *zap.SugaredLogger, config *config.Config) (*Firestore, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.FirestoreConfig.ProjectName)

	if err != nil {
		log.Fatal(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return &Firestore{
		log:    log,
		config: config,
		client: client,
	}, nil
}

//Save creates a scheduler document in forestore collection
func (f *Firestore) Save(ctx context.Context, s *domain.Scheduler) error {

	_, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(s.OrderName).Set(ctx, s)
	if err != nil {
		f.log.Error(err)
		return err
	}

	return nil
}

//Get reterieves a scheduler document of given id
func (f *Firestore) Get(ctx context.Context, orderName string) (*domain.Scheduler, error) {

	doc, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(orderName).Get(ctx)
	if err != nil {
		f.log.Error(err)
		return nil, err
	}

	body, err := json.Marshal(doc.Data())
	if err != nil {
		f.log.Error(err)
		return nil, err
	}

	scheduler := &domain.Scheduler{}
	if err := json.Unmarshal(body, scheduler); err != nil {
		f.log.Error(err)
		return nil, err
	}

	return scheduler, nil
}

//GetAll retrieves queried documents
func (f *Firestore) GetAll(ctx context.Context, path string, op string, value interface{}) ([]*domain.Scheduler, error) {

	var datas = make([]*domain.Scheduler, 0)
	iter := f.client.Collection(f.config.FirestoreConfig.CollectionName).Where(path, op, value).Documents(ctx)

	defer iter.Stop()

	for {

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			f.log.Error(err)
			return nil, err
		}

		var data *domain.Scheduler
		if err := doc.DataTo(data); err != nil {
			f.log.Error(err)
			return nil, err
		}

		datas = append(datas, data)
	}

	return datas, nil
}

//Update updates dronePath field of a scheduler document
func (f *Firestore) Update(ctx context.Context, s *domain.Scheduler) error {

	_, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(s.OrderName).Update(ctx, []firestore.Update{
		{
			Path:  dronePath,
			Value: s.DroneName,
		},
	})
	if err != nil {
		f.log.Error(err)
		return err
	}

	return nil
}

//Delete deletes document of given id
func (f *Firestore) Delete(ctx context.Context, orderName string) error {

	_, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(orderName).Delete(ctx)
	if err != nil {
		f.log.Error(err)
		return err
	}

	return nil
}
