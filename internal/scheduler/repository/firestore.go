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
	documentId  = "document_id"
	orderPath   = "order_name"
	kitchenPath = "kitchen_name"
	dronePath   = "drone_name"
)

//Firestore is a struct for firestore repository
type Firestore struct {
	log    *zap.SugaredLogger
	config *config.Config
	client *firestore.Client
}

//New returns a new instance of Firestore
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

//Save saves a scheduler document
func (f *Firestore) Save(ctx context.Context, s *domain.Scheduler) error {
	_, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(s.DocumentId).Set(ctx, s)
	if err != nil {
		f.log.Error(err)
		return err
	}

	return nil
}

//Get retrieves a scheduler document
func (f *Firestore) Get(ctx context.Context, documentId string) (*domain.Scheduler, error) {
	doc, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(documentId).Get(ctx)
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

//GetAll returns a list of all scheduler documents
func (f *Firestore) GetAll(ctx context.Context) ([]*domain.Scheduler, error) {

	var datas = make([]*domain.Scheduler, 0)
	iter := f.client.Collection(f.config.FirestoreConfig.CollectionName).Documents(ctx)

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

		data := &domain.Scheduler{}

		if err := doc.DataTo(data); err != nil {
			f.log.Error(err)
			return nil, err
		}

		datas = append(datas, data)
	}

	return datas, nil
}

//GetShiped returns a list of all shiped scheduler documents
func (f *Firestore) GetShiped(ctx context.Context) ([]*domain.Scheduler, error) {

	var datas = make([]*domain.Scheduler, 0)
	iter := f.client.Collection(f.config.FirestoreConfig.CollectionName).Where(dronePath, "!=", "").Documents(ctx)

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

		data := &domain.Scheduler{}

		if err := doc.DataTo(data); err != nil {
			f.log.Error(err)
			return nil, err
		}

		datas = append(datas, data)
	}

	return datas, nil
}

//GetNotShiped returns a list of all not shiped scheduler documents
func (f *Firestore) GetNotShiped(ctx context.Context) ([]*domain.Scheduler, error) {

	var datas = make([]*domain.Scheduler, 0)
	iter := f.client.Collection(f.config.FirestoreConfig.CollectionName).Where(dronePath, "==", "").Documents(ctx)

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

		data := &domain.Scheduler{}

		if err := doc.DataTo(data); err != nil {
			f.log.Error(err)
			return nil, err
		}

		datas = append(datas, data)
	}

	return datas, nil
}

//Update updates a scheduler document
func (f *Firestore) Update(ctx context.Context, s *domain.Scheduler) error {
	_, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(s.DocumentId).Update(ctx, []firestore.Update{
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

//Delete deletes a scheduler document
func (f *Firestore) Delete(ctx context.Context, documentId string) error {
	_, err := f.client.Collection(f.config.FirestoreConfig.CollectionName).Doc(documentId).Delete(ctx)
	if err != nil {
		f.log.Error(err)
		return err
	}

	return nil
}
